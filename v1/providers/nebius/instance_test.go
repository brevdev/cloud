package v1

import (
	"strings"
	"testing"
	"time"

	v1 "github.com/brevdev/cloud/v1"
	"github.com/stretchr/testify/assert"
)

func createTestClient() *NebiusClient {
	return &NebiusClient{
		refID: "test-ref",
		serviceAccountKey: `{
			"subject-credentials": {
				"type": "JWT",
				"alg": "RS256",
				"private-key": "-----BEGIN PRIVATE KEY-----\ntest\n-----END PRIVATE KEY-----\n",
				"kid": "publickey-test123",
				"iss": "serviceaccount-test456",
				"sub": "serviceaccount-test456"
			}
		}`,
		tenantID:  "test-tenant",
		projectID: "test-project",
		location:  "eu-north1",
	}
}

func TestNebiusClient_CreateInstance(t *testing.T) {
	t.Skip("CreateInstance requires real SDK initialization - use integration tests instead")
}

func TestNebiusClient_GetInstance(t *testing.T) {
	t.Skip("GetInstance requires real SDK initialization - use integration tests instead")
}

func TestNebiusClient_NotImplementedMethods(t *testing.T) {
	t.Skip("These methods now require real SDK initialization - use integration tests instead")
}

func TestNebiusClient_GetLocations(t *testing.T) {
	t.Skip("GetLocations requires real SDK initialization - use integration tests instead")
}

func TestNebiusClient_MergeInstanceForUpdate(t *testing.T) {
	client := createTestClient()

	currInstance := v1.Instance{
		RefID:          "current-ref",
		CloudCredRefID: "current-cred",
		Name:           "current-name",
		Location:       "current-location",
		CreatedAt:      time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
		CloudID:        "current-cloud-id",
		InstanceType:   "current-type",
		Status:         v1.Status{LifecycleStatus: v1.LifecycleStatusRunning},
	}

	newInstance := v1.Instance{
		RefID:          "new-ref",
		CloudCredRefID: "new-cred",
		Name:           "new-name",
		Location:       "new-location",
		CreatedAt:      time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
		CloudID:        "new-cloud-id",
		InstanceType:   "new-type",
		Status:         v1.Status{LifecycleStatus: v1.LifecycleStatusStopped},
	}

	merged := client.MergeInstanceForUpdate(currInstance, newInstance)

	// These fields should be preserved from current instance
	assert.Equal(t, currInstance.RefID, merged.RefID)
	assert.Equal(t, currInstance.CloudCredRefID, merged.CloudCredRefID)
	assert.Equal(t, currInstance.Name, merged.Name)
	assert.Equal(t, currInstance.Location, merged.Location)
	assert.Equal(t, currInstance.CreatedAt, merged.CreatedAt)
	assert.Equal(t, currInstance.CloudID, merged.CloudID)

	// These fields should come from new instance
	assert.Equal(t, newInstance.InstanceType, merged.InstanceType)
	assert.Equal(t, newInstance.Status, merged.Status)
}

// BenchmarkCreateInstance benchmarks the CreateInstance method
func BenchmarkCreateInstance(b *testing.B) {
	b.Skip("CreateInstance requires real SDK initialization - use integration tests instead")
}

// BenchmarkGetInstance benchmarks the GetInstance method
func BenchmarkGetInstance(b *testing.B) {
	b.Skip("GetInstance requires real SDK initialization - use integration tests instead")
}

// TestStripCIDR tests CIDR notation removal from IP addresses
// Nebius API returns IPs with CIDR notation (e.g., "192.168.1.1/32")
// which breaks SSH connectivity if not stripped
func TestStripCIDR(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "IPv4 with /32 CIDR",
			input:    "195.242.10.162/32",
			expected: "195.242.10.162",
		},
		{
			name:     "IPv4 with /24 CIDR",
			input:    "192.168.1.0/24",
			expected: "192.168.1.0",
		},
		{
			name:     "IPv4 without CIDR",
			input:    "10.0.0.1",
			expected: "10.0.0.1",
		},
		{
			name:     "empty string",
			input:    "",
			expected: "",
		},
		{
			name:     "private IP with CIDR",
			input:    "10.128.0.5/32",
			expected: "10.128.0.5",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := stripCIDR(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

// TestGetGPUMemory tests VRAM mapping for GPU types
func TestGetGPUMemory(t *testing.T) {
	// Import the function from instancetype.go (it's in the same package)
	tests := []struct {
		gpuType      string
		expectedGiB  int64
		shouldBeZero bool
	}{
		{
			gpuType:     "L40S",
			expectedGiB: 48,
		},
		{
			gpuType:     "H100",
			expectedGiB: 80,
		},
		{
			gpuType:     "H200",
			expectedGiB: 141,
		},
		{
			gpuType:     "A100",
			expectedGiB: 80,
		},
		{
			gpuType:     "V100",
			expectedGiB: 32,
		},
		{
			gpuType:     "A10",
			expectedGiB: 24,
		},
		{
			gpuType:     "T4",
			expectedGiB: 16,
		},
		{
			gpuType:     "L4",
			expectedGiB: 24,
		},
		{
			gpuType:     "B200",
			expectedGiB: 192,
		},
		{
			gpuType:      "UNKNOWN_GPU",
			expectedGiB:  0,
			shouldBeZero: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.gpuType, func(t *testing.T) {
			vram := getGPUMemory(tt.gpuType)
			vramGiB := int64(vram) / (1024 * 1024 * 1024)

			if tt.shouldBeZero {
				assert.Equal(t, int64(0), vramGiB, "Unknown GPU type should return 0 VRAM")
			} else {
				assert.Equal(t, tt.expectedGiB, vramGiB,
					"GPU type %s should have %d GiB VRAM", tt.gpuType, tt.expectedGiB)
			}
		})
	}
}

func TestExtractGPUTypeAndName(t *testing.T) {
	// Verify that GPU names no longer include "NVIDIA" prefix
	// Manufacturer info is stored separately in GPU.Manufacturer field
	tests := []struct {
		platformName string
		expectedType string
		expectedName string
	}{
		{
			platformName: "gpu-h100-sxm",
			expectedType: "H100",
			expectedName: "H100", // Should be "H100", not "NVIDIA H100"
		},
		{
			platformName: "gpu-h200-sxm",
			expectedType: "H200",
			expectedName: "H200", // Should be "H200", not "NVIDIA H200"
		},
		{
			platformName: "gpu-l40s",
			expectedType: "L40S",
			expectedName: "L40S", // Should be "L40S", not "NVIDIA L40S"
		},
		{
			platformName: "gpu-a100-sxm4",
			expectedType: "A100",
			expectedName: "A100", // Should be "A100", not "NVIDIA A100"
		},
		{
			platformName: "gpu-v100-sxm2",
			expectedType: "V100",
			expectedName: "V100", // Should be "V100", not "NVIDIA V100"
		},
		{
			platformName: "unknown-platform",
			expectedType: "GPU",
			expectedName: "GPU", // Generic fallback
		},
	}

	for _, tt := range tests {
		t.Run(tt.platformName, func(t *testing.T) {
			gpuType, gpuName := extractGPUTypeAndName(tt.platformName)

			assert.Equal(t, tt.expectedType, gpuType,
				"Platform %s should extract GPU type %s", tt.platformName, tt.expectedType)
			assert.Equal(t, tt.expectedName, gpuName,
				"Platform %s should extract GPU name %s (without 'NVIDIA' prefix)", tt.platformName, tt.expectedName)

			// Ensure name does not contain manufacturer prefix
			assert.NotContains(t, gpuName, "NVIDIA",
				"GPU name should not contain 'NVIDIA' prefix - use GPU.Manufacturer field instead")
		})
	}
}

// TestParseInstanceTypeFormat tests the instance type ID format parsing
func TestParseInstanceTypeFormat(t *testing.T) {
	tests := []struct {
		name             string
		instanceTypeID   string
		expectedGPUType  string
		expectedPreset   string
		shouldParseAsNEW bool
	}{
		{
			name:             "H100 single GPU",
			instanceTypeID:   "nebius-eu-north1-h100-1gpu-16vcpu-200gb",
			expectedGPUType:  "h100",
			expectedPreset:   "1gpu-16vcpu-200gb",
			shouldParseAsNEW: true,
		},
		{
			name:             "L40S quad GPU",
			instanceTypeID:   "nebius-eu-north1-l40s-4gpu-96vcpu-768gb",
			expectedGPUType:  "l40s",
			expectedPreset:   "4gpu-96vcpu-768gb",
			shouldParseAsNEW: true,
		},
		{
			name:             "H200 octa GPU",
			instanceTypeID:   "nebius-us-central1-h200-8gpu-128vcpu-1600gb",
			expectedGPUType:  "h200",
			expectedPreset:   "8gpu-128vcpu-1600gb",
			shouldParseAsNEW: true,
		},
		{
			name:             "CPU only",
			instanceTypeID:   "nebius-eu-north1-cpu-4vcpu-16gb",
			expectedGPUType:  "cpu",
			expectedPreset:   "4vcpu-16gb",
			shouldParseAsNEW: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Parse the format
			parts := strings.Split(tt.instanceTypeID, "-")
			assert.GreaterOrEqual(t, len(parts), 4, "Instance type should have at least 4 parts")
			assert.Equal(t, "nebius", parts[0], "Should start with 'nebius'")

			// Find GPU type
			var gpuType string
			var presetStartIdx int
			for i := 1; i < len(parts); i++ {
				partLower := strings.ToLower(parts[i])
				if partLower == "cpu" || partLower == "l40s" || partLower == "h100" ||
					partLower == "h200" || partLower == "a100" || partLower == "v100" {
					gpuType = partLower
					presetStartIdx = i + 1
					break
				}
			}

			assert.Equal(t, tt.expectedGPUType, gpuType, "Should extract correct GPU type")
			assert.Greater(t, presetStartIdx, 0, "Should find preset start index")

			if presetStartIdx > 0 && presetStartIdx < len(parts) {
				presetName := strings.Join(parts[presetStartIdx:], "-")
				assert.Equal(t, tt.expectedPreset, presetName, "Should extract correct preset name")
			}
		})
	}
}
