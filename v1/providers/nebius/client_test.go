package v1

import (
	"context"
	"encoding/json"
	"runtime/debug"
	"sync"
	"testing"

	v1 "github.com/brevdev/cloud/v1"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNebiusCredential(t *testing.T) {
	tests := []struct {
		name        string
		refID       string
		serviceKey  string
		tenantID    string
		expectError bool
	}{
		{
			name:  "valid credentials",
			refID: "test-ref-id",
			serviceKey: `{
				"subject-credentials": {
					"type": "JWT",
					"alg": "RS256",
					"private-key": "-----BEGIN PRIVATE KEY-----\nMIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQC7test\n-----END PRIVATE KEY-----\n",
					"kid": "publickey-test123",
					"iss": "serviceaccount-test456",
					"sub": "serviceaccount-test456"
				}
			}`,
			tenantID: "test-tenant-id",
		},
		{
			name:  "empty tenant ID",
			refID: "test-ref",
			serviceKey: `{
				"subject-credentials": {
					"type": "JWT",
					"alg": "RS256",
					"private-key": "-----BEGIN PRIVATE KEY-----\ntest\n-----END PRIVATE KEY-----\n",
					"kid": "publickey-test123",
					"iss": "serviceaccount-test456",
					"sub": "serviceaccount-test456"
				}
			}`,
			tenantID:    "",
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cred := NewNebiusCredential(tt.refID, tt.serviceKey, tt.tenantID)

			assert.Equal(t, tt.refID, cred.GetReferenceID())
			assert.Equal(t, v1.CloudProviderID("nebius"), cred.GetCloudProviderID())
			assert.Equal(t, v1.APITypeLocational, cred.GetAPIType())

			tenantID, err := cred.GetTenantID()
			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				// tenantID should now just return the tenant ID (not a project ID)
				assert.Equal(t, tt.tenantID, tenantID)
			}
		})
	}
}

func TestNebiusCredential_GetCapabilities(t *testing.T) {
	serviceKey := `{
		"subject-credentials": {
			"type": "JWT",
			"alg": "RS256",
			"private-key": "-----BEGIN PRIVATE KEY-----\ntest\n-----END PRIVATE KEY-----\n",
			"kid": "publickey-test123",
			"iss": "serviceaccount-test456",
			"sub": "serviceaccount-test456"
		}
	}`
	cred := NewNebiusCredential("test", serviceKey, "tenant-id")

	capabilities, err := cred.GetCapabilities(context.Background())
	require.NoError(t, err)

	expectedCapabilities := []v1.Capability{
		v1.CapabilityCreateInstance,
		v1.CapabilityTerminateInstance,
		v1.CapabilityCreateTerminateInstance,
		v1.CapabilityRebootInstance,
		v1.CapabilityStopStartInstance,
		v1.CapabilityResizeInstanceVolume,
		v1.CapabilityMachineImage,
		v1.CapabilityTags,
		v1.CapabilityInstanceUserData,
		v1.CapabilityVPC,
		v1.CapabilityManagedKubernetes,
	}

	assert.ElementsMatch(t, expectedCapabilities, capabilities)
}

func TestNebiusClient_Creation(t *testing.T) {
	tests := []struct {
		name          string
		serviceKey    string
		expectError   bool
		errorContains string
	}{
		{
			name: "valid service account JSON",
			serviceKey: `{
				"subject-credentials": {
					"type": "JWT",
					"alg": "RS256",
					"private-key": "-----BEGIN PRIVATE KEY-----\nMIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQC7test\n-----END PRIVATE KEY-----\n",
					"kid": "publickey-test123",
					"iss": "serviceaccount-test456",
					"sub": "serviceaccount-test456"
				}
			}`,
		},
		{
			name:          "invalid JSON",
			serviceKey:    `invalid json`,
			expectError:   true,
			errorContains: "failed to parse service account key JSON",
		},
		{
			name:          "empty JSON object",
			serviceKey:    `{}`,
			expectError:   true,
			errorContains: "invalid service account algorithm",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client, err := NewNebiusClient(
				context.Background(),
				"test-ref",
				tt.serviceKey,
				"test-tenant-id",
				"test-project-id",
				"eu-north1",
			)

			if tt.expectError {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.errorContains)
				assert.Nil(t, client)
			} else if err != nil {
				// Note: This will likely fail due to invalid credentials
				// but we're testing the JSON parsing part
				// Check if it's a JSON parsing error vs SDK initialization error
				assert.NotContains(t, err.Error(), "failed to parse service account key JSON")
			}
		})
	}
}

func TestNebiusClient_BasicMethods(t *testing.T) {
	// Create a client with mock credentials (will fail SDK initialization but that's OK for basic tests)
	client := &NebiusClient{
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

	t.Run("GetAPIType", func(t *testing.T) {
		assert.Equal(t, v1.APITypeLocational, client.GetAPIType())
	})

	t.Run("GetCloudProviderID", func(t *testing.T) {
		assert.Equal(t, v1.CloudProviderID("nebius"), client.GetCloudProviderID())
	})

	t.Run("GetReferenceID", func(t *testing.T) {
		assert.Equal(t, "test-ref", client.GetReferenceID())
	})

	t.Run("GetTenantID", func(t *testing.T) {
		tenantID, err := client.GetTenantID()
		assert.NoError(t, err)
		assert.Equal(t, "test-project", tenantID)
	})

	t.Run("GetMaxCreateRequestsPerMinute", func(t *testing.T) {
		assert.Equal(t, 10, client.GetMaxCreateRequestsPerMinute())
	})
}

func TestNebiusClient_GetCapabilities(t *testing.T) {
	client := &NebiusClient{
		projectID: "test-project",
	}

	capabilities, err := client.GetCapabilities(context.Background())
	require.NoError(t, err)

	expectedCapabilities := []v1.Capability{
		v1.CapabilityCreateInstance,
		v1.CapabilityTerminateInstance,
		v1.CapabilityCreateTerminateInstance,
		v1.CapabilityRebootInstance,
		v1.CapabilityStopStartInstance,
		v1.CapabilityResizeInstanceVolume,
		v1.CapabilityMachineImage,
		v1.CapabilityTags,
		v1.CapabilityInstanceUserData,
		v1.CapabilityVPC,
		v1.CapabilityManagedKubernetes,
	}

	assert.ElementsMatch(t, expectedCapabilities, capabilities)
}

func TestValidServiceAccountJSON(t *testing.T) {
	tests := []struct {
		name    string
		jsonStr string
		isValid bool
	}{
		{
			name: "valid nebius service account",
			jsonStr: `{
				"id": "service-account-key-id",
				"service_account_id": "your-service-account-id",
				"created_at": "2024-01-01T00:00:00Z",
				"key_algorithm": "RSA_2048",
				"public_key": "-----BEGIN PUBLIC KEY-----\ntest\n-----END PUBLIC KEY-----\n",
				"private_key": "-----BEGIN PRIVATE KEY-----\ntest\n-----END PRIVATE KEY-----\n"
			}`,
			isValid: true,
		},
		{
			name: "minimal valid JSON",
			jsonStr: `{
				"service_account_id": "test-sa",
				"private_key": "test-key"
			}`,
			isValid: true,
		},
		{
			name:    "invalid JSON",
			jsonStr: `{invalid}`,
			isValid: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result map[string]interface{}
			err := json.Unmarshal([]byte(tt.jsonStr), &result)

			if tt.isValid {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
			}
		})
	}
}

func TestExtractRegionFromProjectName(t *testing.T) {
	tests := []struct {
		name           string
		projectName    string
		expectedRegion string
	}{
		{
			name:           "default-project pattern with eu-north1",
			projectName:    "default-project-eu-north1",
			expectedRegion: "eu-north1",
		},
		{
			name:           "default-project pattern with us-central1",
			projectName:    "default-project-us-central1",
			expectedRegion: "us-central1",
		},
		{
			name:           "default pattern with region",
			projectName:    "default-eu-west1",
			expectedRegion: "eu-west1",
		},
		{
			name:           "project name containing region",
			projectName:    "my-project-eu-north1-test",
			expectedRegion: "eu-north1",
		},
		{
			name:           "just region name",
			projectName:    "eu-north1",
			expectedRegion: "eu-north1",
		},
		{
			name:           "uppercase project name",
			projectName:    "DEFAULT-PROJECT-US-EAST1",
			expectedRegion: "us-east1",
		},
		{
			name:           "project name without known region",
			projectName:    "my-custom-project",
			expectedRegion: "",
		},
		{
			name:           "empty project name",
			projectName:    "",
			expectedRegion: "",
		},
		{
			name:           "project name with partial region match",
			projectName:    "eu-project", // contains "eu-" but not full region
			expectedRegion: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := extractRegionFromProjectName(tt.projectName)
			assert.Equal(t, tt.expectedRegion, result,
				"extractRegionFromProjectName(%q) = %q, want %q",
				tt.projectName, result, tt.expectedRegion)
		})
	}
}

func TestMainVersionFromLDFlags(t *testing.T) {
	tests := []struct {
		name      string
		ldflags   string
		wantValue string
	}{
		{
			name:      "main version in standard format",
			ldflags:   "-X main.Version=v1.2.3 -X main.BuildTime=2026-03-04_12:00:00",
			wantValue: "v1.2.3",
		},
		{
			name:      "main version in quoted token",
			ldflags:   `-s -w -X "main.Version=2026.03.04"`,
			wantValue: "2026.03.04",
		},
		{
			name:      "main version after other -X values",
			ldflags:   "-X main.BuildTime=2026-03-04_12:00:00 -X main.Version=abc123",
			wantValue: "abc123",
		},
		{
			name:      "main version not present",
			ldflags:   "-s -w -X main.BuildTime=2026-03-04_12:00:00",
			wantValue: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.wantValue, mainVersionFromLDFlags(tt.ldflags))
		})
	}
}

func TestFormatNebiusUserAgentPrefix(t *testing.T) {
	tests := []struct {
		name      string
		version   string
		wantValue string
	}{
		{
			name:      "no version",
			version:   "",
			wantValue: nebiusSDKUserAgentPrefix,
		},
		{
			name:      "version set",
			version:   "v1.2.3",
			wantValue: "brev-cloud-sdk-nebius/v1.2.3",
		},
		{
			name:      "leading slash in version",
			version:   "/v1.2.3",
			wantValue: "brev-cloud-sdk-nebius/v1.2.3",
		},
		{
			name:      "whitespace version",
			version:   "   ",
			wantValue: nebiusSDKUserAgentPrefix,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.wantValue, formatNebiusUserAgentPrefix(tt.version))
		})
	}
}

func TestMainVersionFromBuildInfo_IsCached(t *testing.T) {
	originalReadBuildInfo := readBuildInfo
	t.Cleanup(func() {
		readBuildInfo = originalReadBuildInfo
		mainVersionCacheOnce = sync.Once{}
		mainVersionCache = ""
	})

	mainVersionCacheOnce = sync.Once{}
	mainVersionCache = ""

	calls := 0
	readBuildInfo = func() (*debug.BuildInfo, bool) {
		calls++
		return &debug.BuildInfo{
			Settings: []debug.BuildSetting{
				{
					Key:   "-ldflags",
					Value: "-X main.Version=v9.9.9",
				},
			},
		}, true
	}

	assert.Equal(t, "v9.9.9", mainVersionFromBuildInfo())

	// Change the source to prove the second call comes from cache, not a new read.
	readBuildInfo = func() (*debug.BuildInfo, bool) {
		calls++
		return &debug.BuildInfo{
			Settings: []debug.BuildSetting{
				{
					Key:   "-ldflags",
					Value: "-X main.Version=v0.0.1",
				},
			},
		}, true
	}

	assert.Equal(t, "v9.9.9", mainVersionFromBuildInfo())
	assert.Equal(t, 1, calls)
}
