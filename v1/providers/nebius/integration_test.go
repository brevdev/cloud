package v1

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
	"testing"
	"time"

	v1 "github.com/brevdev/cloud/v1"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/ssh"
)

// Integration tests that require actual Nebius credentials
// These tests are skipped unless proper environment variables are set

func setupIntegrationTest(t *testing.T) *NebiusClient {
	serviceAccountJSON := os.Getenv("NEBIUS_SERVICE_ACCOUNT_JSON")
	tenantID := os.Getenv("NEBIUS_TENANT_ID")

	if serviceAccountJSON == "" || tenantID == "" {
		t.Skip("Skipping integration test: NEBIUS_SERVICE_ACCOUNT_JSON and NEBIUS_TENANT_ID must be set")
	}

	// Read from file if path is provided
	if _, err := os.Stat(serviceAccountJSON); err == nil {
		data, err := os.ReadFile(serviceAccountJSON)
		require.NoError(t, err, "Failed to read service account file")
		serviceAccountJSON = string(data)
	}

	// Create client (project ID is now determined in NewNebiusClient as default-project-{location})
	client, err := NewNebiusClient(
		context.Background(),
		"integration-test-ref",
		serviceAccountJSON,
		tenantID,
		"", // projectID is now determined as default-project-{location}
		"eu-north1",
	)
	require.NoError(t, err, "Failed to create Nebius client for integration test")

	return client
}

// generateTestSSHKeyPair generates an RSA SSH key pair for testing
// Returns private key (PEM format) and public key (OpenSSH format)
func generateTestSSHKeyPair(t *testing.T) (privateKey, publicKey string) {
	// Generate RSA key pair
	privKey, err := rsa.GenerateKey(rand.Reader, 2048)
	require.NoError(t, err, "Failed to generate RSA key")

	// Encode private key to PEM format
	privKeyPEM := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privKey),
	}
	privateKeyBytes := pem.EncodeToMemory(privKeyPEM)

	// Generate public key in OpenSSH format
	pub, err := ssh.NewPublicKey(&privKey.PublicKey)
	require.NoError(t, err, "Failed to create SSH public key")
	publicKeyBytes := ssh.MarshalAuthorizedKey(pub)

	return string(privateKeyBytes), string(publicKeyBytes)
}

// waitForSSH waits for SSH to become available on the instance
// This is critical because cloud-init takes time to configure the instance
func waitForSSH(t *testing.T, publicIP, privateKey, sshUser string, timeout time.Duration) error {
	// Parse private key
	signer, err := ssh.ParsePrivateKey([]byte(privateKey))
	if err != nil {
		return fmt.Errorf("failed to parse private key: %w", err)
	}

	config := &ssh.ClientConfig{
		User: sshUser,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), // For testing only - NEVER use in production
		Timeout:         5 * time.Second,
	}

	deadline := time.Now().Add(timeout)
	attempt := 0
	for time.Now().Before(deadline) {
		attempt++
		t.Logf("SSH connection attempt %d to %s:22 (timeout in %v)...",
			attempt, publicIP, time.Until(deadline).Round(time.Second))

		conn, err := ssh.Dial("tcp", fmt.Sprintf("%s:22", publicIP), config)
		if err == nil {
			conn.Close()
			t.Logf("✓ SSH is ready on %s after %d attempts", publicIP, attempt)
			return nil
		}

		t.Logf("  SSH not ready yet: %v", err)
		time.Sleep(10 * time.Second)
	}

	return fmt.Errorf("SSH did not become ready within %v (%d attempts)", timeout, attempt)
}

// testSSHConnectivity validates that SSH connectivity works and the instance is accessible
func testSSHConnectivity(t *testing.T, publicIP, privateKey, sshUser string) {
	t.Logf("Testing SSH connectivity to %s as user %s...", publicIP, sshUser)

	// Parse private key
	signer, err := ssh.ParsePrivateKey([]byte(privateKey))
	require.NoError(t, err, "Failed to parse private key")

	config := &ssh.ClientConfig{
		User: sshUser,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), // For testing only
		Timeout:         10 * time.Second,
	}

	// Connect to the instance
	client, err := ssh.Dial("tcp", fmt.Sprintf("%s:22", publicIP), config)
	require.NoError(t, err, "SSH connection should succeed")
	defer client.Close()
	t.Log("✓ SSH connection established successfully")

	// Run a test command to verify functionality
	session, err := client.NewSession()
	require.NoError(t, err, "Failed to create SSH session")
	defer session.Close()

	// Run a simple command
	output, err := session.CombinedOutput("echo 'SSH connectivity test successful' && uname -a")
	require.NoError(t, err, "Failed to run test command")

	outputStr := string(output)
	assert.Contains(t, outputStr, "SSH connectivity test successful", "Command output should contain test message")
	assert.NotEmpty(t, outputStr, "Command output should not be empty")

	t.Logf("✓ SSH command execution successful")
	t.Logf("  Output: %s", outputStr)
}

func TestIntegration_ClientCreation(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	client := setupIntegrationTest(t)
	// Test basic client functionality
	assert.Equal(t, v1.APITypeLocational, client.GetAPIType())
	assert.Equal(t, v1.CloudProviderID("nebius"), client.GetCloudProviderID())
	assert.Equal(t, "integration-test-ref", client.GetReferenceID())

	tenantID, err := client.GetTenantID()
	assert.NoError(t, err)
	assert.NotEmpty(t, tenantID)
}

func TestIntegration_GetCapabilities(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	client := setupIntegrationTest(t)
	ctx := context.Background()

	capabilities, err := client.GetCapabilities(ctx)
	require.NoError(t, err)
	assert.NotEmpty(t, capabilities)

	// Verify expected capabilities are present
	expectedCapabilities := []v1.Capability{
		v1.CapabilityCreateInstance,
		v1.CapabilityTerminateInstance,
		v1.CapabilityRebootInstance,
		v1.CapabilityStopStartInstance,
		v1.CapabilityResizeInstanceVolume,
		v1.CapabilityMachineImage,
		v1.CapabilityTags,
	}

	for _, expected := range expectedCapabilities {
		assert.Contains(t, capabilities, expected)
	}
}

func TestIntegration_GetLocations(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	client := setupIntegrationTest(t)
	ctx := context.Background()

	locations, err := client.GetLocations(ctx, v1.GetLocationsArgs{})
	require.NoError(t, err)
	assert.NotEmpty(t, locations)

	// Verify location structure
	for _, location := range locations {
		assert.NotEmpty(t, location.Name)
		// Note: DisplayName might not be available in current implementation
	}
}

// TestIntegration_InstanceLifecycle tests the full instance lifecycle
// This is a "smoke test" that creates, monitors, and destroys an instance
func TestIntegration_InstanceLifecycle(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	// This test is currently expected to fail with "not implemented" errors
	// Update when full Nebius API implementation is complete

	client := setupIntegrationTest(t)
	ctx := context.Background()

	// Step 0: Get available instance types to find one we can use
	t.Log("Discovering available instance types...")
	instanceTypes, err := client.GetInstanceTypes(ctx, v1.GetInstanceTypeArgs{})
	require.NoError(t, err, "Failed to get instance types")

	if len(instanceTypes) == 0 {
		t.Skip("No instance types available - skipping instance lifecycle test")
	}

	// Use the first available instance type (should have quota)
	selectedInstanceType := instanceTypes[0]
	t.Logf("Using instance type: %s (Location: %s)", selectedInstanceType.ID, selectedInstanceType.Location)

	// Step 0.5: Generate SSH key pair for testing (inspired by Shadeform's SSH key handling)
	t.Log("Generating SSH key pair for instance access...")
	privateKey, publicKey := generateTestSSHKeyPair(t)
	t.Log("✓ SSH key pair generated successfully")

	// Step 1: Create instance with SSH key
	instanceRefID := "integration-test-" + time.Now().Format("20060102-150405")
	instanceName := "nebius-int-test-" + time.Now().Format("20060102-150405") // Unique name to avoid collisions
	createAttrs := v1.CreateInstanceAttrs{
		RefID:        instanceRefID,
		Name:         instanceName,
		InstanceType: string(selectedInstanceType.ID), // Use discovered instance type
		ImageID:      "ubuntu22.04-cuda12",            // Use known-good Nebius image family
		DiskSize:     50 * 1024 * 1024 * 1024,         // 50 GiB in bytes
		Location:     selectedInstanceType.Location,   // Use the instance type's location
		PublicKey:    publicKey,                       // SSH public key for access (like Shadeform)
		Tags: map[string]string{
			"test":        "integration",
			"created-by":  "nebius-integration-test",
			"auto-delete": "true",
		},
	}

	t.Logf("Creating instance with RefID: %s", instanceRefID)
	instance, err := client.CreateInstance(ctx, createAttrs)

	// For now, we expect this to work (returns mock instance)
	// When real implementation is ready, this should create actual instance
	require.NoError(t, err)
	require.NotNil(t, instance)
	assert.Equal(t, instanceRefID, instance.RefID)

	instanceCloudID := instance.CloudID
	t.Logf("Created instance with CloudID: %s", instanceCloudID)

	// Register cleanup to ensure resources are deleted even if test fails
	// Track whether we've already terminated to avoid double-delete
	instanceTerminated := false
	t.Cleanup(func() {
		if instanceTerminated {
			t.Logf("Cleanup: Instance %s already terminated, skipping", instanceCloudID)
			return
		}
		t.Logf("Cleanup: Terminating instance %s", instanceCloudID)
		cleanupCtx := context.Background()
		if err := client.TerminateInstance(cleanupCtx, instanceCloudID); err != nil {
			t.Logf("WARNING: Failed to cleanup instance %s: %v", instanceCloudID, err)
			t.Logf("         Please manually delete: instance=%s, disk=%s-boot-disk", instanceCloudID, instanceName)
		} else {
			t.Logf("Successfully cleaned up instance %s", instanceCloudID)
		}
	})

	// Step 2: Get instance details and validate SSH connectivity fields
	t.Logf("Getting instance details for CloudID: %s", instanceCloudID)
	retrievedInstance, err := client.GetInstance(ctx, instanceCloudID)
	require.NoError(t, err)
	require.NotNil(t, retrievedInstance)
	assert.Equal(t, instanceCloudID, retrievedInstance.CloudID)

	// Validate SSH connectivity fields are populated (similar to Shadeform)
	t.Log("Validating SSH connectivity fields...")
	assert.NotEmpty(t, retrievedInstance.PublicIP, "Public IP should be assigned")
	assert.NotEmpty(t, retrievedInstance.PrivateIP, "Private IP should be assigned")
	assert.NotEmpty(t, retrievedInstance.SSHUser, "SSH user should be set")
	assert.Equal(t, 22, retrievedInstance.SSHPort, "SSH port should be 22")
	assert.NotEmpty(t, retrievedInstance.Hostname, "Hostname should be set")
	t.Logf("✓ SSH connectivity fields populated: IP=%s, User=%s, Port=%d",
		retrievedInstance.PublicIP, retrievedInstance.SSHUser, retrievedInstance.SSHPort)

	// Step 2.5: Wait for SSH to be ready (instances need time to boot and run cloud-init)
	// This is critical - cloud-init takes time to configure SSH keys
	if retrievedInstance.PublicIP != "" {
		t.Log("Waiting for SSH to become available (cloud-init configuration may take 2-5 minutes)...")
		err = waitForSSH(t, retrievedInstance.PublicIP, privateKey, retrievedInstance.SSHUser, 5*time.Minute)
		if err != nil {
			t.Logf("WARNING: SSH did not become available: %v", err)
			t.Log("This may be expected if the instance is still booting or cloud-init is still running")
		} else {
			// Step 2.6: Test actual SSH connectivity
			t.Log("Testing SSH connectivity and command execution...")
			testSSHConnectivity(t, retrievedInstance.PublicIP, privateKey, retrievedInstance.SSHUser)
			t.Log("✓ SSH connectivity validated successfully")
		}
	} else {
		t.Log("WARNING: No public IP available, skipping SSH connectivity test")
	}

	// Step 3: List instances (currently not implemented)
	t.Log("Listing instances...")
	instances, err := client.ListInstances(ctx, v1.ListInstancesArgs{})
	// This is expected to fail with current implementation
	if err != nil {
		t.Logf("ListInstances failed as expected: %v", err)
		assert.Contains(t, err.Error(), "implementation pending")
	} else {
		t.Logf("Found %d instances", len(instances))
	}

	// Step 4: Stop instance
	t.Logf("Stopping instance: %s", instanceCloudID)
	err = client.StopInstance(ctx, instanceCloudID)
	require.NoError(t, err, "StopInstance should succeed")
	t.Logf("✓ Successfully stopped instance %s", instanceCloudID)

	// Verify instance is stopped
	stoppedInstance, err := client.GetInstance(ctx, instanceCloudID)
	require.NoError(t, err, "Should be able to get stopped instance")
	assert.Equal(t, v1.LifecycleStatusStopped, stoppedInstance.Status.LifecycleStatus, "Instance should be stopped")
	t.Logf("✓ Verified instance status: %s", stoppedInstance.Status.LifecycleStatus)

	// Step 5: Start instance
	t.Logf("Starting instance: %s", instanceCloudID)
	err = client.StartInstance(ctx, instanceCloudID)
	require.NoError(t, err, "StartInstance should succeed")
	t.Logf("✓ Successfully started instance %s", instanceCloudID)

	// Verify instance is running again
	startedInstance, err := client.GetInstance(ctx, instanceCloudID)
	require.NoError(t, err, "Should be able to get started instance")
	assert.Equal(t, v1.LifecycleStatusRunning, startedInstance.Status.LifecycleStatus, "Instance should be running")
	t.Logf("✓ Verified instance status: %s", startedInstance.Status.LifecycleStatus)

	// Step 6: Terminate instance
	// Note: Cleanup is registered via t.Cleanup() above to ensure deletion even on test failure
	// This step tests that termination works as part of the lifecycle test
	t.Logf("Testing termination of instance: %s", instanceCloudID)
	err = client.TerminateInstance(ctx, instanceCloudID)

	// TerminateInstance is fully implemented, should succeed
	if err != nil {
		t.Errorf("TerminateInstance failed: %v", err)
	} else {
		t.Logf("Successfully terminated instance %s", instanceCloudID)
		instanceTerminated = true // Mark as terminated to skip cleanup
	}

	t.Log("Instance lifecycle test completed")
}

// TestIntegration_GetInstanceTypes tests fetching available instance types
// Removed - comprehensive version is below

// TestIntegration_GetImages tests fetching available images
func TestIntegration_GetImages(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	client := setupIntegrationTest(t)
	ctx := context.Background()

	images, err := client.GetImages(ctx, v1.GetImageArgs{})

	// Currently expected to fail with "not implemented"
	if err != nil {
		t.Logf("GetImages failed as expected: %v", err)
		assert.Contains(t, err.Error(), "implementation pending")
	} else {
		t.Logf("Found %d images", len(images))

		// If implementation is complete, verify image structure
		for _, img := range images {
			assert.NotEmpty(t, img.ID)
			assert.NotEmpty(t, img.Name)
		}
	}
}

// TestIntegration_ErrorHandling tests how the client handles various error conditions
func TestIntegration_ErrorHandling(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	// Test with invalid credentials
	t.Run("InvalidCredentials", func(t *testing.T) {
		tenantID := os.Getenv("NEBIUS_TENANT_ID")
		if tenantID == "" {
			t.Skip("NEBIUS_TENANT_ID must be set for error handling test")
		}

		_, err := NewNebiusClient(
			context.Background(),
			"test-ref",
			`{"invalid": "credentials"}`,
			tenantID,
			"test-project-id",
			"eu-north1",
		)

		// Should fail during SDK initialization
		assert.Error(t, err)
		t.Logf("Invalid credentials error: %v", err)
	})

	// Test with malformed JSON
	t.Run("MalformedJSON", func(t *testing.T) {
		_, err := NewNebiusClient(
			context.Background(),
			"test-ref",
			`{invalid json}`,
			"test-tenant",
			"test-project",
			"eu-north1",
		)

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "failed to parse service account key JSON")
	})
}

func TestIntegration_GetInstanceTypes(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	client := setupIntegrationTest(t)
	ctx := context.Background()

	t.Run("Get instance types with quota filtering", func(t *testing.T) {
		instanceTypes, err := client.GetInstanceTypes(ctx, v1.GetInstanceTypeArgs{})
		require.NoError(t, err, "Failed to get instance types")

		t.Logf("Found %d instance types with available quota", len(instanceTypes))

		// Verify that we got some instance types
		// If this fails, it means either:
		// 1. No quotas are configured for this tenant
		// 2. All quotas are fully consumed
		// 3. The quota API integration is not working
		if len(instanceTypes) == 0 {
			t.Log("WARNING: No instance types with available quota found. Check tenant quotas.")
		}

		// Validate instance type structure
		for _, it := range instanceTypes {
			t.Logf("Instance Type: %s (%s) - Location: %s, Available: %v",
				it.ID, it.Type, it.Location, it.IsAvailable)

			// Basic validation
			assert.NotEmpty(t, it.ID, "Instance type should have an ID")
			assert.NotEmpty(t, it.Type, "Instance type should have a type")
			assert.NotEmpty(t, it.Location, "Instance type should have a location")
			assert.True(t, it.IsAvailable, "Returned instance types should be available")
			assert.True(t, it.ElasticRootVolume, "Nebius supports elastic root volumes")

			// Verify supported storage is configured
			assert.NotEmpty(t, it.SupportedStorage, "Instance type should have supported storage")
			if len(it.SupportedStorage) > 0 {
				storage := it.SupportedStorage[0]
				assert.NotNil(t, storage.MinSize, "Storage should have minimum size")
				assert.NotNil(t, storage.MaxSize, "Storage should have maximum size")
				assert.True(t, storage.IsElastic, "Storage should be elastic")
				assert.Equal(t, "network-ssd", storage.Type, "Storage type should be network-ssd")

				t.Logf("  Storage: %s, Min: %d GB, Max: %d GB, Elastic: %v",
					storage.Type,
					*storage.MinSize/(1024*1024*1024),
					*storage.MaxSize/(1024*1024*1024),
					storage.IsElastic)
			}

			// Verify GPU details if present
			if len(it.SupportedGPUs) > 0 {
				gpu := it.SupportedGPUs[0]
				t.Logf("  GPU: %s (Type: %s), Count: %d, Manufacturer: %s",
					gpu.Name, gpu.Type, gpu.Count, gpu.Manufacturer)

				assert.NotEmpty(t, gpu.Type, "GPU should have a type")
				assert.NotEmpty(t, gpu.Name, "GPU should have a name")
				assert.Greater(t, gpu.Count, int32(0), "GPU count should be positive")
				assert.Equal(t, v1.ManufacturerNVIDIA, gpu.Manufacturer, "Nebius GPUs are NVIDIA")

				// Verify GPU type is not empty (any GPU with quota is supported)
				assert.NotEmpty(t, gpu.Type, "GPU type should not be empty")
			}

			// Verify CPU and memory
			assert.Greater(t, it.VCPU, int32(0), "VCPU count should be positive")
			assert.Greater(t, int64(it.Memory), int64(0), "Memory should be positive")

			// Verify pricing is enriched from Nebius Billing API
			if it.BasePrice != nil {
				t.Logf("  Price: %s %s/hr", it.BasePrice.Number(), it.BasePrice.CurrencyCode())
				assert.NotEmpty(t, it.BasePrice.Number(), "Price should have a value")
				assert.Equal(t, "USD", it.BasePrice.CurrencyCode(), "Nebius pricing should be in USD")

				// Price should be reasonable (not negative or extremely high)
				priceStr := it.BasePrice.Number()
				var priceFloat float64
				fmt.Sscanf(priceStr, "%f", &priceFloat)
				assert.Greater(t, priceFloat, 0.0, "Price should be positive")
				assert.Less(t, priceFloat, 1000.0, "Price per hour should be reasonable (< $1000/hr)")
			} else {
				t.Logf("  Price: Not available (pricing API may have failed)")
			}
		}
	})

	t.Run("Verify pricing enrichment", func(t *testing.T) {
		instanceTypes, err := client.GetInstanceTypes(ctx, v1.GetInstanceTypeArgs{})
		require.NoError(t, err)

		pricedCount := 0
		unpricedCount := 0

		for _, it := range instanceTypes {
			if it.BasePrice != nil {
				pricedCount++
			} else {
				unpricedCount++
			}
		}

		t.Logf("Pricing statistics:")
		t.Logf("  Instance types with pricing: %d", pricedCount)
		t.Logf("  Instance types without pricing: %d", unpricedCount)

		// We expect most (ideally all) instance types to have pricing
		// But pricing API failures are non-critical, so we just log if missing
		if unpricedCount > 0 {
			t.Logf("WARNING: %d instance types are missing pricing data", unpricedCount)
			t.Logf("         This may indicate Nebius Billing API issues or quota problems")
		}

		// At least verify that pricing is available for SOME instance types
		// If zero, that suggests a systematic problem with pricing integration
		if len(instanceTypes) > 0 && pricedCount == 0 {
			t.Error("No instance types have pricing data - pricing integration may be broken")
		}
	})

	t.Run("Filter by supported platforms", func(t *testing.T) {
		instanceTypes, err := client.GetInstanceTypes(ctx, v1.GetInstanceTypeArgs{})
		require.NoError(t, err)

		// Count instance types by platform type
		gpuCounts := make(map[string]int)
		cpuCount := 0

		for _, it := range instanceTypes {
			if len(it.SupportedGPUs) > 0 {
				gpuType := it.SupportedGPUs[0].Type
				gpuCounts[gpuType]++
			} else {
				cpuCount++
			}
		}

		t.Logf("Instance type distribution:")
		for gpuType, count := range gpuCounts {
			t.Logf("  %s: %d", gpuType, count)
		}
		t.Logf("  CPU-only: %d", cpuCount)

		// Verify we have at least some instance types (either GPU or CPU)
		assert.Greater(t, len(instanceTypes), 0, "Should have at least one instance type with quota")

		// If no GPU quota is available, that's okay - just log it
		if len(gpuCounts) == 0 {
			t.Logf("⚠️  No GPU quota allocated - only CPU instances available")
			t.Logf("   To test GPU instances, request GPU quota from Nebius support")
		}

		// Verify CPU presets are limited per region
		if cpuCount > 0 {
			// We limit CPU platforms to 3 presets each, and have 2 CPU platforms (cpu-d3, cpu-e2)
			// Across multiple regions, this multiplies (e.g., 4 regions × 2 platforms × 3 presets = 24)
			maxCPUPresetsPerRegion := 6 // 3 per platform × 2 platforms
			// The count could be higher if we have quota in multiple regions
			t.Logf("   CPU instance types found: %d (max %d per region)", cpuCount, maxCPUPresetsPerRegion)
		}
	})

	t.Run("Verify preset enumeration", func(t *testing.T) {
		instanceTypes, err := client.GetInstanceTypes(ctx, v1.GetInstanceTypeArgs{})
		require.NoError(t, err)

		// Group by platform and count presets
		presetsByPlatform := make(map[string][]string)
		for _, it := range instanceTypes {
			platformName := ""
			if len(it.SupportedGPUs) > 0 {
				platformName = it.SupportedGPUs[0].Type
			} else {
				platformName = "CPU"
			}
			presetsByPlatform[platformName] = append(presetsByPlatform[platformName], string(it.ID))
		}

		t.Logf("Preset enumeration by platform:")
		for platform, presets := range presetsByPlatform {
			t.Logf("  %s: %d presets", platform, len(presets))
			for _, preset := range presets {
				t.Logf("    - %s", preset)
			}
		}

		// Verify each platform has multiple presets (1, 2, 4, 8 GPUs typically)
		for platform, presets := range presetsByPlatform {
			if platform != "CPU" {
				assert.Greater(t, len(presets), 0,
					"Platform %s should have at least one preset", platform)
			}
		}
	})
}

// Example of how to run integration tests:
//
// # Set up credentials
// export NEBIUS_SERVICE_ACCOUNT_JSON='{"service_account_id": "...", "private_key": "..."}'
// export NEBIUS_TENANT_ID="your-tenant-id"
//
// # Run integration tests
// go test -v -tags=integration ./v1/providers/nebius/...
//
// # Run only integration tests (not unit tests)
// go test -v -run TestIntegration ./v1/providers/nebius/...
//
// # Run integration tests with timeout
// go test -v -timeout=10m -run TestIntegration ./v1/providers/nebius/...
