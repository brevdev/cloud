package v1

import (
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
