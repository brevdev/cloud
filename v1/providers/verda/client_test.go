package verda

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/alecthomas/units"
	v1 "github.com/brevdev/cloud/v1"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	verdago "github.com/verda-cloud/verdacloud-sdk-go/pkg/verda"
	"golang.org/x/crypto/ssh"
)

func TestVerdaCredential(t *testing.T) {
	credential := NewVerdaCredential("credential-ref", "client-id", "client-secret")

	assert.Equal(t, DefaultAPIURL, credential.APIURL)
	assert.Equal(t, v1.CloudProviderID(CloudProviderID), credential.GetCloudProviderID())
	assert.Equal(t, v1.APITypeGlobal, credential.GetAPIType())
	assert.Equal(t, "credential-ref", credential.GetReferenceID())
	require.NoError(t, credential.Validate())

	tenantID, err := credential.GetTenantID()
	require.NoError(t, err)
	assert.NotEmpty(t, tenantID)

	rotatedSecret := NewVerdaCredential("credential-ref", "client-id", "new-secret")
	rotatedTenantID, err := rotatedSecret.GetTenantID()
	require.NoError(t, err)
	assert.Equal(t, tenantID, rotatedTenantID)

	invalid := NewVerdaCredential("credential-ref", "", "")
	require.Error(t, invalid.Validate())
}

func TestWrapVerdaInsufficientResourcesError(t *testing.T) {
	err := wrapVerdaError(&verdago.APIError{
		StatusCode: http.StatusServiceUnavailable,
		Message:    "Not enough resources to deploy a 1 GPU instance type 1A100.22V in FIN-01",
	})

	require.ErrorIs(t, err, v1.ErrInsufficientResources)
	assert.False(t, errors.Is(err, v1.ErrServiceUnavailable))
}

func TestGetInstanceTypesAndLocations(t *testing.T) { //nolint:funlen // One catalog fixture exercises all shared validations.
	server := newVerdaTestServer(t, func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/instance-types":
			assert.Equal(t, "usd", r.URL.Query().Get("currency"))
			writeJSON(t, w, []verdago.InstanceTypeInfo{
				{
					ID:           "native-gb300",
					InstanceType: "1GB300.32V",
					Model:        "GB300",
					CPU: verdago.InstanceCPU{
						NumberOfCores: 32,
					},
					GPU: verdago.InstanceGPU{
						NumberOfGPUs: 1,
					},
					GPUMemory: verdago.InstanceMemory{
						SizeInGigabytes: 288,
					},
					Memory: verdago.InstanceMemory{
						SizeInGigabytes: 225,
					},
					Storage: verdago.InstanceStorage{
						Description: "dynamic",
					},
					PricePerHour: 8.62,
					Currency:     "usd",
					Manufacturer: "NVIDIA",
				},
				{
					ID:           "native-h100",
					InstanceType: "1H100.80S.22V",
					Model:        "H100 80GB",
					CPU: verdago.InstanceCPU{
						NumberOfCores: 22,
					},
					GPU: verdago.InstanceGPU{
						NumberOfGPUs: 1,
					},
					GPUMemory: verdago.InstanceMemory{
						SizeInGigabytes: 80,
					},
					Memory: verdago.InstanceMemory{
						SizeInGigabytes: 128,
					},
					Storage: verdago.InstanceStorage{
						Description: "1TiB NVMe SSD",
					},
					PricePerHour: 3.17,
					SpotPrice:    2.54,
					Currency:     "usd",
					Manufacturer: "NVIDIA",
					P2P:          "SXM",
				},
				{
					ID:           "native-cpu",
					InstanceType: "CPU.4V.16G",
					CPU: verdago.InstanceCPU{
						NumberOfCores: 4,
					},
					Memory: verdago.InstanceMemory{
						SizeInGigabytes: 16,
					},
					Storage: verdago.InstanceStorage{
						Description: "100GiB SSD",
					},
					PricePerHour: 0.25,
					Currency:     "usd",
				},
			})
		case "/instance-availability":
			writeJSON(t, w, []verdago.LocationAvailability{
				{
					LocationCode: "FIN-03",
					Availabilities: []string{
						"1H100.80S.22V",
						"1GB300.32V",
					},
				},
				{LocationCode: "NOR-01", Availabilities: []string{"CPU.4V.16G"}},
			})
		case "/locations":
			writeJSON(t, w, []verdago.Location{
				{Code: "NOR-01", Name: "Norway 01", CountryCode: "NO"},
				{Code: "FIN-03", Name: "Finland 03", CountryCode: "FI"},
				{Code: "FIN-01", Name: "Finland 01", CountryCode: "FI"},
			})
		default:
			http.NotFound(w, r)
		}
	})
	defer server.Close()

	client := newTestVerdaClient(t, server)
	ctx := context.Background()

	instanceTypes, err := client.GetInstanceTypes(ctx, v1.GetInstanceTypeArgs{})
	require.NoError(t, err)
	require.Len(t, instanceTypes, 3)

	h100 := findInstanceType(t, instanceTypes, "1H100.80S.22V")
	assert.Equal(t, v1.ManufacturerNVIDIA, h100.SupportedGPUs[0].Manufacturer)
	assert.Equal(t, int32(1), h100.SupportedGPUs[0].Count)
	assert.Equal(t, "H100", h100.SupportedGPUs[0].Name)
	assert.Equal(t, "H100 80GB", h100.SupportedGPUs[0].Type)
	assert.Equal(t, "NVMe", h100.SupportedStorage[0].Type)
	assert.Equal(t, v1.Tebibyte, h100.SupportedStorage[0].SizeBytes.Unit())
	assert.True(t, h100.Preemptible)
	assert.True(t, h100.ElasticRootVolume)
	assert.Equal(t, CloudProviderID, h100.Provider)
	assert.Equal(t, []v1.Architecture{v1.ArchitectureX86_64}, h100.SupportedArchitectures)
	assert.Equal(t, v1.Gigabyte, h100.MemoryBytes.Unit())
	assertLegacyBytesMatch(t, h100.Memory, h100.MemoryBytes)
	assertLegacyBytesMatch(t, h100.SupportedGPUs[0].Memory, h100.SupportedGPUs[0].MemoryBytes)
	assertLegacyBytesMatch(t, h100.SupportedStorage[0].Size, h100.SupportedStorage[0].SizeBytes)

	armTypes, err := client.GetInstanceTypes(ctx, v1.GetInstanceTypeArgs{
		ArchitectureFilter: &v1.ArchitectureFilter{
			IncludeArchitectures: []v1.Architecture{v1.ArchitectureARM64},
		},
	})
	require.NoError(t, err)
	require.Len(t, armTypes, 1)
	assert.Equal(t, "1GB300.32V", armTypes[0].Type)
	assert.Equal(t, []v1.Architecture{v1.ArchitectureARM64}, armTypes[0].SupportedArchitectures)

	filtered, err := client.GetInstanceTypes(ctx, v1.GetInstanceTypeArgs{
		Locations:     v1.LocationsFilter{"NOR-01"},
		InstanceTypes: []string{"CPU.4V.16G"},
	})
	require.NoError(t, err)
	require.Len(t, filtered, 1)
	assert.Equal(t, "CPU.4V.16G", filtered[0].Type)
	assert.Empty(t, filtered[0].SupportedGPUs)

	locations, err := client.GetLocations(ctx, v1.GetLocationsArgs{})
	require.NoError(t, err)
	require.Len(t, locations, 2)
	assert.Equal(t, "FIN-03", locations[0].Name)
	assert.Equal(t, "FIN", locations[0].Country)
	assert.Equal(t, "NOR", locations[1].Country)

	allLocations, err := client.GetLocations(ctx, v1.GetLocationsArgs{IncludeUnavailable: true})
	require.NoError(t, err)
	require.Len(t, allLocations, 3)
	assert.False(t, allLocations[0].Available)

	require.NoError(t, v1.ValidateGetLocations(ctx, client))
	require.NoError(t, v1.ValidateGetInstanceTypes(ctx, client))
	require.NoError(t, v1.ValidateLocationalInstanceTypes(ctx, client))
	require.NoError(t, v1.ValidateStableInstanceTypeIDs(ctx, client, []v1.InstanceTypeID{
		"FIN-03-noSub-1GB300.32V",
		"FIN-03-noSub-1H100.80S.22V",
		"NOR-01-noSub-CPU.4V.16G",
	}))
}

func TestInstanceLifecycle(t *testing.T) { //nolint:gocyclo,funlen // One fixture shows the complete API lifecycle.
	var mu sync.Mutex
	var createdRequest verdago.CreateInstanceRequest
	var createdInstance verdago.Instance
	var actionRequest verdago.InstanceActionRequest
	var createdSSHKey verdago.SSHKey
	var createdScript verdago.StartupScript
	deletedResources := make(map[string]bool)

	server := newVerdaTestServer(t, func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		defer mu.Unlock()

		switch {
		case r.Method == http.MethodGet && r.URL.Path == "/images":
			assert.Equal(t, "1H100.80S.22V", r.URL.Query().Get("instance_type"))
			writeJSON(t, w, []verdago.Image{
				{
					ID:        "image-id",
					ImageType: "ubuntu-24.04-cuda-12.8-open-docker",
					IsDefault: true,
				},
			})
		case r.Method == http.MethodGet && r.URL.Path == "/ssh-keys":
			keys := []verdago.SSHKey{}
			if createdSSHKey.ID != "" {
				keys = append(keys, createdSSHKey)
			}
			writeJSON(t, w, keys)
		case r.Method == http.MethodPost && r.URL.Path == "/ssh-keys":
			var request verdago.CreateSSHKeyRequest
			require.NoError(t, json.NewDecoder(r.Body).Decode(&request))
			createdSSHKey = verdago.SSHKey{
				ID:        "ssh-key-1",
				Name:      request.Name,
				PublicKey: request.PublicKey,
			}
			w.WriteHeader(http.StatusCreated)
			_, _ = io.WriteString(w, "ssh-key-1")
		case r.Method == http.MethodGet && r.URL.Path == "/ssh-keys/ssh-key-1":
			writeJSON(t, w, []verdago.SSHKey{createdSSHKey})
		case r.Method == http.MethodGet && r.URL.Path == "/scripts":
			scripts := []verdago.StartupScript{}
			if createdScript.ID != "" {
				scripts = append(scripts, createdScript)
			}
			writeJSON(t, w, scripts)
		case r.Method == http.MethodPost && r.URL.Path == "/scripts":
			var request verdago.CreateStartupScriptRequest
			require.NoError(t, json.NewDecoder(r.Body).Decode(&request))
			assert.Contains(t, request.Script, "ufw default deny incoming")
			assert.Contains(t, request.Script, "DOCKER-USER")
			createdScript = verdago.StartupScript{
				ID:     "script-1",
				Name:   request.Name,
				Script: request.Script,
			}
			writeJSON(t, w, createdScript)
		case r.Method == http.MethodPost && r.URL.Path == "/instances":
			require.NoError(t, json.NewDecoder(r.Body).Decode(&createdRequest))
			ip := "203.0.113.10"
			osVolumeID := "os-volume-1"
			createdInstance = verdago.Instance{
				ID:           "instance-1",
				IP:           &ip,
				Status:       verdago.StatusPending,
				CreatedAt:    time.Now(),
				Hostname:     createdRequest.Hostname,
				Description:  createdRequest.Description,
				Location:     createdRequest.LocationCode,
				IsSpot:       createdRequest.IsSpot,
				InstanceType: createdRequest.InstanceType,
				Image:        createdRequest.Image,
				Storage: verdago.InstanceStorage{
					Description: "100GiB NVMe SSD",
				},
				SSHKeyIDs:  createdRequest.SSHKeyIDs,
				OSVolumeID: &osVolumeID,
				VolumeIDs:  []string{"data-volume-1"},
			}
			w.WriteHeader(http.StatusCreated)
			writeJSON(t, w, createdInstance)
		case r.Method == http.MethodGet && r.URL.Path == "/instances/instance-1":
			writeJSON(t, w, createdInstance)
		case r.Method == http.MethodPut && r.URL.Path == "/instances":
			require.NoError(t, json.NewDecoder(r.Body).Decode(&actionRequest))
			w.WriteHeader(http.StatusAccepted)
			writeJSON(t, w, []verdago.InstanceActionResult{{
				Action:     actionRequest.Action,
				InstanceID: "instance-1",
				Status:     "success",
			}})
		case r.Method == http.MethodDelete && r.URL.Path == "/scripts/script-1":
			deletedResources[r.URL.Path] = true
			w.WriteHeader(http.StatusNoContent)
		case r.Method == http.MethodDelete && r.URL.Path == "/ssh-keys/ssh-key-1":
			deletedResources[r.URL.Path] = true
			w.WriteHeader(http.StatusNoContent)
		default:
			http.NotFound(w, r)
		}
	})
	defer server.Close()

	client := newTestVerdaClient(t, server)
	authorizedKey, _, _ := newTestPublicKeys(t)
	instance, err := client.CreateInstance(context.Background(), v1.CreateInstanceAttrs{
		Name:          "Brev Validation VM",
		RefID:         "ref-123",
		Location:      "FIN-03",
		PublicKey:     authorizedKey + " test@example.com",
		InstanceType:  "1H100.80S.22V",
		DiskSizeBytes: v1.NewBytes(100, v1.Gibibyte),
	})
	require.NoError(t, err)
	assert.Equal(t, "Brev Validation VM", instance.Name)
	assert.Equal(t, "Brev Validation VM", instance.Hostname)
	assert.Equal(t, "ref-123", instance.RefID)
	assert.Equal(t, "credential-ref", instance.CloudCredRefID)
	assert.Equal(t, "ubuntu-24.04-cuda-12.8-open-docker", instance.ImageID)
	assert.Equal(t, v1.LifecycleStatusPending, instance.Status.LifecycleStatus)
	assert.Equal(t, v1.InstanceTypeID("FIN-03-noSub-1H100.80S.22V"), instance.InstanceTypeID)
	assertLegacyBytesMatch(t, instance.DiskSize, instance.DiskSizeBytes)

	assert.Equal(t, "ref-123_credential-ref", createdRequest.Description)
	assert.LessOrEqual(t, len(createdRequest.Description), maxInstanceDescriptionLength)
	assert.Equal(t, "brev-key-ref-123", createdSSHKey.Name)
	assert.Equal(t, authorizedKey, createdSSHKey.PublicKey)
	assert.Equal(t, "brev-firewall-ref-123", createdScript.Name)
	assert.Equal(t, []string{"ssh-key-1"}, createdRequest.SSHKeyIDs)
	require.NotNil(t, createdRequest.OSVolume)
	assert.Equal(t, 100, createdRequest.OSVolume.Size)

	require.NoError(t, client.TerminateInstance(context.Background(), instance.CloudID))
	assert.Equal(t, verdago.ActionDelete, actionRequest.Action)
	assert.ElementsMatch(t, []string{"data-volume-1", "os-volume-1"}, actionRequest.VolumeIDs)
	assert.True(t, actionRequest.DeletePermanently)
	assert.True(t, deletedResources["/scripts/script-1"])
	assert.True(t, deletedResources["/ssh-keys/ssh-key-1"])
}

func TestNormalizeSSHPublicKey(t *testing.T) {
	authorizedKey, pkixPEMKey, pkcs1PEMKey := newTestPublicKeys(t)

	for _, test := range []struct {
		name      string
		publicKey string
		want      string
		wantError bool
	}{
		{
			name:      "OpenSSH",
			publicKey: authorizedKey + " test@example.com",
			want:      authorizedKey,
		},
		{
			name:      "PKIX PEM",
			publicKey: pkixPEMKey,
			want:      authorizedKey,
		},
		{
			name:      "PKCS1 PEM",
			publicKey: pkcs1PEMKey,
			want:      authorizedKey,
		},
		{
			name:      "invalid",
			publicKey: "not-a-public-key",
			wantError: true,
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			got, err := normalizeSSHPublicKey(test.publicKey)
			if test.wantError {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, test.want, got)
		})
	}
}

func newTestPublicKeys(t *testing.T) (authorizedKey string, pkixPEMKey string, pkcs1PEMKey string) {
	t.Helper()

	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	require.NoError(t, err)
	publicKey := &privateKey.PublicKey
	sshPublicKey, err := ssh.NewPublicKey(publicKey)
	require.NoError(t, err)
	pkixDER, err := x509.MarshalPKIXPublicKey(publicKey)
	require.NoError(t, err)

	return strings.TrimSpace(string(ssh.MarshalAuthorizedKey(sshPublicKey))),
		string(pem.EncodeToMemory(&pem.Block{Type: "PUBLIC KEY", Bytes: pkixDER})),
		string(pem.EncodeToMemory(&pem.Block{Type: "RSA PUBLIC KEY", Bytes: x509.MarshalPKCS1PublicKey(publicKey)}))
}

func TestBuildStartupScriptRejectsUnsafeRules(t *testing.T) {
	_, err := buildStartupScript(v1.FirewallRules{
		IngressRules: []v1.FirewallRule{{
			FromPort: 9999,
			ToPort:   9999,
			IPRanges: []string{"not-a-cidr"},
		}},
	})
	require.Error(t, err)
}

func TestRequestedDiskSizeGiB(t *testing.T) {
	tests := []struct {
		name  string
		attrs v1.CreateInstanceAttrs
		want  int
	}{
		{
			name:  "new bytes field",
			attrs: v1.CreateInstanceAttrs{DiskSizeBytes: v1.NewBytes(100, v1.Gibibyte)},
			want:  100,
		},
		{
			name:  "new bytes field rounds up",
			attrs: v1.CreateInstanceAttrs{DiskSizeBytes: v1.NewBytes(100, v1.Gigabyte)},
			want:  94,
		},
		{
			name:  "legacy field",
			attrs: v1.CreateInstanceAttrs{DiskSize: 100 * (1 << 30)},
			want:  100,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.want, requestedDiskSizeGiB(test.attrs))
		})
	}
}

func TestVerdaGPUName(t *testing.T) {
	tests := map[string]string{
		"A100 80GB":         "A100",
		"A100 22GB":         "A100",
		"RTX PRO 6000":      "RTX PRO 6000",
		"RTX PRO 6000 96GB": "RTX PRO 6000",
	}

	for model, want := range tests {
		t.Run(model, func(t *testing.T) {
			assert.Equal(t, want, verdaGPUName(model))
		})
	}
}

func TestInstanceDescription(t *testing.T) {
	description, err := makeInstanceDescription("ref-123", "credential-ref")
	require.NoError(t, err)
	assert.Equal(t, "ref-123_credential-ref", description)

	refID, cloudCredRefID := parseInstanceDescription(description)
	assert.Equal(t, "ref-123", refID)
	assert.Equal(t, "credential-ref", cloudCredRefID)

	_, err = makeInstanceDescription(strings.Repeat("a", maxInstanceDescriptionLength), "credential-ref")
	require.Error(t, err)
}

func findInstanceType(t *testing.T, instanceTypes []v1.InstanceType, typeName string) v1.InstanceType {
	t.Helper()
	for _, instanceType := range instanceTypes {
		if instanceType.Type == typeName {
			return instanceType
		}
	}
	require.FailNow(t, "instance type not found", typeName)
	return v1.InstanceType{}
}

func assertLegacyBytesMatch(t *testing.T, legacy units.Base2Bytes, size v1.Bytes) {
	t.Helper()
	assert.Equal(t, int64(legacy), size.ByteCount().Int64())
}

func newVerdaTestServer(
	t *testing.T,
	handler func(http.ResponseWriter, *http.Request),
) *httptest.Server {
	t.Helper()
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/oauth2/token" {
			writeJSON(t, w, verdago.TokenResponse{
				AccessToken: "test-token",
				TokenType:   "Bearer",
				ExpiresIn:   3600,
			})
			return
		}
		assert.Equal(t, "Bearer test-token", r.Header.Get("Authorization"))
		handler(w, r)
	}))
}

func newTestVerdaClient(t *testing.T, server *httptest.Server) *VerdaClient {
	t.Helper()
	credential := NewVerdaCredential("credential-ref", "client-id", "client-secret")
	credential.APIURL = server.URL
	client, err := NewVerdaClient(
		*credential,
		"",
		WithHTTPClient(server.Client()),
	)
	require.NoError(t, err)
	return client
}

func writeJSON(t *testing.T, w http.ResponseWriter, value any) {
	t.Helper()
	w.Header().Set("Content-Type", "application/json")
	require.NoError(t, json.NewEncoder(w).Encode(value))
}
