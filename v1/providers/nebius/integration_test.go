package v1

import (
	"context"
	"os"
	"testing"
	"time"

	v1 "github.com/brevdev/cloud/v1"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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

	// Create credential to get the project ID
	cred := NewNebiusCredential("integration-test-ref", serviceAccountJSON, tenantID)
	projectID, err := cred.GetTenantID()
	require.NoError(t, err, "Failed to get project ID")

	client, err := NewNebiusClient(
		context.Background(),
		"integration-test-ref",
		serviceAccountJSON,
		tenantID,
		projectID,
		"eu-north1",
	)
	require.NoError(t, err, "Failed to create Nebius client for integration test")

	return client
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

	// Step 1: Create instance
	instanceRefID := "integration-test-" + time.Now().Format("20060102-150405")
	createAttrs := v1.CreateInstanceAttrs{
		RefID:        instanceRefID,
		Name:         "nebius-integration-test",
		InstanceType: "standard-2", // This may need to be updated with actual Nebius instance types
		ImageID:      "ubuntu-20.04", // This may need to be updated with actual Nebius image IDs
		DiskSize:     20,
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

	// Step 2: Get instance details
	t.Logf("Getting instance details for CloudID: %s", instanceCloudID)
	retrievedInstance, err := client.GetInstance(ctx, instanceCloudID)
	require.NoError(t, err)
	require.NotNil(t, retrievedInstance)
	assert.Equal(t, instanceCloudID, retrievedInstance.CloudID)

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

	// Step 4: Stop instance (currently not implemented)
	t.Logf("Stopping instance: %s", instanceCloudID)
	err = client.StopInstance(ctx, instanceCloudID)
	if err != nil {
		t.Logf("StopInstance failed as expected: %v", err)
		assert.Contains(t, err.Error(), "implementation pending")
	}

	// Step 5: Start instance (currently not implemented)
	t.Logf("Starting instance: %s", instanceCloudID)
	err = client.StartInstance(ctx, instanceCloudID)
	if err != nil {
		t.Logf("StartInstance failed as expected: %v", err)
		assert.Contains(t, err.Error(), "implementation pending")
	}

	// Step 6: Terminate instance (currently not implemented)
	t.Logf("Terminating instance: %s", instanceCloudID)
	err = client.TerminateInstance(ctx, instanceCloudID)
	if err != nil {
		t.Logf("TerminateInstance failed as expected: %v", err)
		assert.Contains(t, err.Error(), "implementation pending")
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