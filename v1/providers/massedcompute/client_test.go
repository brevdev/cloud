package massedcompute

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/alecthomas/units"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	v1 "github.com/brevdev/cloud/v1"
	openapi "github.com/brevdev/cloud/v1/providers/massedcompute/gen/massedcompute"
)

const testSSHPublicKey = "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIDBdptDTzJ2cOmdyryG1B7yb1YssiCQs6SWu4HlbZXGE"

func TestMassedComputeCredential(t *testing.T) {
	credential := NewMassedComputeCredential("credential-ref", "api-token")

	assert.Equal(t, DefaultAPIURL, credential.APIURL)
	assert.Equal(t, v1.CloudProviderID(CloudProviderID), credential.GetCloudProviderID())
	assert.Equal(t, v1.APITypeGlobal, credential.GetAPIType())
	assert.Equal(t, "credential-ref", credential.GetReferenceID())
	require.NoError(t, credential.Validate())

	tenantID, err := credential.GetTenantID()
	require.NoError(t, err)
	assert.NotEmpty(t, tenantID)

	invalid := NewMassedComputeCredential("credential-ref", "")
	require.Error(t, invalid.Validate())
}

func TestGetInstanceTypesAndLocations(t *testing.T) { //nolint:funlen // test ok
	server := newMassedComputeTestServer(t, func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, "/gpu-inventory", r.URL.Path)
		writeMassedComputeJSON(t, w, map[string]any{
			"gpu_inventory": map[string]any{
				"gpu_1x_l40": map[string]any{
					"instance_type": map[string]any{
						"name":                 "gpu_1x_l40",
						"description":          "1x L40 (48GB)",
						"price_cents_per_hour": 86,
						"specs": map[string]any{
							"vcpu_count": 14,
							"memory_gib": 72,
							"storage_gb": 625,
						},
					},
					"regions_with_capacity_available": []map[string]any{{
						"name":        "us-central-1",
						"description": "Wichita, KS",
					}},
					"capacity_available": 19,
				},
				"gpu_4x_h200_nvl_nvlink_discount": map[string]any{
					"instance_type": map[string]any{
						"name":                 "gpu_4x_h200_nvl_nvlink_discount",
						"description":          "4x H200 NVL (141GB) NVLink [Spot]",
						"price_cents_per_hour": 78,
						"specs": map[string]any{
							"vcpu_count": 14,
							"memory_gib": 72,
							"storage_gb": 625,
						},
					},
					"regions_with_capacity_available": []map[string]any{{"name": "spot-only-1"}},
					"capacity_available":              19,
				},
				"gpu_1x_a6000_high_ram": map[string]any{
					"instance_type": map[string]any{
						"name":                 "gpu_1x_a6000_high_ram",
						"description":          "1x RTX A6000 (48GB) [Premium]",
						"price_cents_per_hour": 125,
						"specs": map[string]any{
							"vcpu_count": 32,
							"memory_gib": 256,
							"storage_gb": 1000,
						},
					},
					"regions_with_capacity_available": []map[string]any{{"name": "us-central-1"}},
				},
				"gpu_1x_a6000_low_ram": map[string]any{
					"instance_type": map[string]any{
						"name":                 "gpu_1x_a6000_low_ram",
						"description":          "1x RTX A6000 (48GB) [ALT Config]",
						"price_cents_per_hour": 95,
						"specs": map[string]any{
							"vcpu_count": 20,
							"memory_gib": 128,
							"storage_gb": 750,
						},
					},
					"regions_with_capacity_available": []map[string]any{{"name": "us-central-1"}},
				},
				"gpu_4x_h100_nvl_nvlink": map[string]any{
					"instance_type": map[string]any{
						"name":        "gpu_4x_h100_nvl_nvlink",
						"description": "4x H100 NVL",
						"specs":       map[string]any{},
					},
					"regions_with_capacity_available": []map[string]any{{"name": "us-central-1"}},
				},
				"gpu_8x_DGX_A100": map[string]any{
					"instance_type": map[string]any{
						"name":        "gpu_8x_DGX_A100",
						"description": "8x DGX A100 (80GB)",
						"specs":       map[string]any{},
					},
					"regions_with_capacity_available": []map[string]any{{"name": "us-central-1"}},
				},
				"gpu_1x_A100_SXM4": map[string]any{
					"instance_type": map[string]any{
						"name":        "gpu_1x_A100_SXM4",
						"description": "1x A100 SXM4 (80GB)",
						"specs":       map[string]any{},
					},
					"regions_with_capacity_available": []map[string]any{{"name": "us-central-1"}},
				},
				"cpu_small_amd_epyc": map[string]any{
					"instance_type": map[string]any{
						"name":        "cpu_small_amd_epyc",
						"description": "CPU-only instance",
						"specs": map[string]any{
							"vcpu_count": 4,
							"memory_gib": 16,
							"storage_gb": 250,
						},
					},
					"regions_with_capacity_available": []map[string]any{{"name": "us-central-1"}},
				},
				"gpu_2x_a6000_nvlink": map[string]any{
					"instance_type": map[string]any{
						"name":        "gpu_2x_a6000_nvlink",
						"description": "2x RTX A6000 (48GB) NVLink",
						"specs":       map[string]any{},
					},
					"regions_with_capacity_available": []map[string]any{{"name": "us-central-1"}},
				},
			},
		})
	})
	defer server.Close()

	client := newMassedComputeTestClient(t, server, "")
	ctx := context.Background()
	instanceTypes, err := client.GetInstanceTypes(ctx, v1.GetInstanceTypeArgs{})
	require.NoError(t, err)
	require.Len(t, instanceTypes, 8)
	instanceTypesByName := make(map[string]v1.InstanceType, len(instanceTypes))
	for _, instanceType := range instanceTypes {
		instanceTypesByName[instanceType.Type] = instanceType
	}

	instanceType, ok := instanceTypesByName["gpu_1x_l40"]
	require.True(t, ok)
	assert.Equal(t, "gpu_1x_l40", instanceType.Type)
	assert.Equal(t, "us-central-1", instanceType.Location)
	assert.Equal(t, v1.NewBytes(72, v1.Gibibyte), instanceType.MemoryBytes)
	assertLegacyBytesMatch(t, instanceType.Memory, instanceType.MemoryBytes)
	require.Len(t, instanceType.SupportedStorage, 1)
	assert.Equal(t, v1.NewBytes(625, v1.Gigabyte), instanceType.SupportedStorage[0].SizeBytes)
	assertLegacyBytesMatch(t, instanceType.SupportedStorage[0].Size, instanceType.SupportedStorage[0].SizeBytes)
	require.Len(t, instanceType.SupportedGPUs, 1)
	assert.Equal(t, int32(1), instanceType.SupportedGPUs[0].Count)
	assert.Equal(t, "L40", instanceType.SupportedGPUs[0].Name)
	assert.Equal(t, v1.NewBytes(48, v1.Gigabyte), instanceType.SupportedGPUs[0].MemoryBytes)
	assert.Equal(t, []string{"on-demand"}, instanceType.SupportedUsageClasses)
	assert.False(t, instanceType.Preemptible)
	assert.Equal(t, "0.86", instanceType.BasePrice.Number())

	nvlInstanceType, ok := instanceTypesByName["gpu_4x_h100_nvl_nvlink"]
	require.True(t, ok)
	require.Len(t, nvlInstanceType.SupportedGPUs, 1)
	assert.Equal(t, int32(4), nvlInstanceType.SupportedGPUs[0].Count)
	assert.Equal(t, "H100", nvlInstanceType.SupportedGPUs[0].Name)
	assert.Equal(t, "H100", nvlInstanceType.SupportedGPUs[0].Type)
	assert.Equal(t, v1.NewBytes(94, v1.Gigabyte), nvlInstanceType.SupportedGPUs[0].MemoryBytes)
	assert.Equal(t, "NVLink", nvlInstanceType.SupportedGPUs[0].NetworkDetails)

	for _, test := range []struct {
		typeName  string
		memoryGiB v1.BytesValue
		vcpus     int32
		storageGB v1.BytesValue
	}{
		{typeName: "gpu_1x_a6000_high_ram", memoryGiB: 256, vcpus: 32, storageGB: 1000},
		{typeName: "gpu_1x_a6000_low_ram", memoryGiB: 128, vcpus: 20, storageGB: 750},
	} {
		variant, ok := instanceTypesByName[test.typeName]
		require.True(t, ok)
		require.Len(t, variant.SupportedGPUs, 1)
		assert.Equal(t, "RTX A6000", variant.SupportedGPUs[0].Name)
		assert.Equal(t, v1.NewBytes(test.memoryGiB, v1.Gibibyte), variant.MemoryBytes)
		assert.Equal(t, test.vcpus, variant.VCPU)
		require.Len(t, variant.SupportedStorage, 1)
		assert.Equal(t, v1.NewBytes(test.storageGB, v1.Gigabyte), variant.SupportedStorage[0].SizeBytes)
	}

	dgxInstanceType, ok := instanceTypesByName["gpu_8x_DGX_A100"]
	require.True(t, ok)
	require.Len(t, dgxInstanceType.SupportedGPUs, 1)
	assert.Equal(t, int32(8), dgxInstanceType.SupportedGPUs[0].Count)
	assert.Equal(t, "DGX A100", dgxInstanceType.SupportedGPUs[0].Name)
	assert.Equal(t, "DGX A100", dgxInstanceType.SupportedGPUs[0].Type)

	sxmInstanceType, ok := instanceTypesByName["gpu_1x_A100_SXM4"]
	require.True(t, ok)
	require.Len(t, sxmInstanceType.SupportedGPUs, 1)
	assert.Equal(t, "A100", sxmInstanceType.SupportedGPUs[0].Name)
	assert.Equal(t, "A100", sxmInstanceType.SupportedGPUs[0].Type)
	assert.Equal(t, "SXM4", sxmInstanceType.SupportedGPUs[0].NetworkDetails)

	cpuInstanceType, ok := instanceTypesByName["cpu_small_amd_epyc"]
	require.True(t, ok)
	assert.Empty(t, cpuInstanceType.SupportedGPUs)
	assert.Equal(t, v1.NewBytes(16, v1.Gibibyte), cpuInstanceType.MemoryBytes)
	assert.Equal(t, int32(4), cpuInstanceType.VCPU)

	nvlinkInstanceType, ok := instanceTypesByName["gpu_2x_a6000_nvlink"]
	require.True(t, ok)
	require.Len(t, nvlinkInstanceType.SupportedGPUs, 1)
	assert.Equal(t, int32(2), nvlinkInstanceType.SupportedGPUs[0].Count)
	assert.Equal(t, "RTX A6000", nvlinkInstanceType.SupportedGPUs[0].Name)
	assert.Equal(t, "NVLink", nvlinkInstanceType.SupportedGPUs[0].NetworkDetails)

	locations, err := client.GetLocations(ctx, v1.GetLocationsArgs{})
	require.NoError(t, err)
	require.Len(t, locations, 1)
	assert.Equal(t, "us-central-1", locations[0].Name)
	assert.Equal(t, "Wichita, KS", locations[0].Description)
	assert.True(t, locations[0].Available)

	require.NoError(t, v1.ValidateGetLocations(ctx, client))
	require.NoError(t, v1.ValidateGetInstanceTypes(ctx, client))
	require.NoError(t, v1.ValidateLocationalInstanceTypes(ctx, client))
}

func TestGPUVRAMFallback(t *testing.T) {
	for _, test := range []struct {
		description    string
		count          int32
		model          string
		networkDetails string
		memoryGB       v1.BytesValue
	}{
		{description: "1x H100 NVL", count: 1, model: "H100", networkDetails: "NVLink", memoryGB: 94},
		{description: "8x B200 SXM6", count: 8, model: "B200", networkDetails: "SXM6", memoryGB: 180},
		{description: "8x B300 SXM6", count: 8, model: "B300", networkDetails: "SXM6", memoryGB: 288},
		{description: "4x H100 NVL (80GB)", count: 4, model: "H100", networkDetails: "NVLink", memoryGB: 80},
	} {
		t.Run(test.description, func(t *testing.T) {
			gpus := massedComputeGPUs(test.description)
			require.Len(t, gpus, 1)
			assert.Equal(t, test.count, gpus[0].Count)
			assert.Equal(t, test.model, gpus[0].Name)
			assert.Equal(t, test.networkDetails, gpus[0].NetworkDetails)
			assert.Equal(t, v1.NewBytes(test.memoryGB, v1.Gigabyte), gpus[0].MemoryBytes)
		})
	}

	assert.Empty(t, massedComputeGPUs("CPU-only instance"))
}

func TestCreateInstanceHonorsRequestedLocation(t *testing.T) {
	var createdKey openapi.SshKeysPostRequest
	var launchRequest openapi.InstanceLaunchPostRequest
	server := newMassedComputeTestServer(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "Bearer api-token", r.Header.Get("Authorization"))
		switch {
		case r.Method == http.MethodGet && r.URL.Path == "/images":
			writeMassedComputeJSON(t, w, map[string]any{"images": []map[string]any{{
				"vm_image_id": 42, "vm_image_name": "Ubuntu Server 22.04 w/ drivers",
			}}})
		case r.Method == http.MethodGet && r.URL.Path == "/ssh-keys":
			writeMassedComputeJSON(t, w, map[string]any{"sshKeys": []any{}})
		case r.Method == http.MethodPost && r.URL.Path == "/ssh-keys":
			require.NoError(t, json.NewDecoder(r.Body).Decode(&createdKey))
			writeMassedComputeJSON(t, w, map[string]any{"sshKey": map[string]any{
				"id": "key-id", "name": createdKey.Name,
			}})
		case r.Method == http.MethodPost && r.URL.Path == "/instance/launch":
			require.NoError(t, json.NewDecoder(r.Body).Decode(&launchRequest))
			writeMassedComputeJSON(t, w, map[string]any{"response": "instance-id"})
		case r.Method == http.MethodGet && r.URL.Path == "/instance/instance-id":
			writeMassedComputeJSON(t, w, map[string]any{"runningInstances": []map[string]any{{
				"uuid":     "instance-id",
				"name":     "dev_tagged-credential_ref-123",
				"status":   "rented",
				"username": "ubuntu",
				"region":   map[string]any{"name": "requested-region"},
				"image":    map[string]any{"id": 42, "name": "Ubuntu Server 22.04 w/ drivers"},
				"product":  map[string]any{"name": "gpu_1x_l40"},
			}}})
		default:
			http.NotFound(w, r)
		}
	})
	defer server.Close()

	client := newMassedComputeTestClient(t, server, "client-region")
	instance, err := client.CreateInstance(context.Background(), v1.CreateInstanceAttrs{
		RefID:        "ref-123",
		Name:         "dev-environment",
		InstanceType: "gpu_1x_l40",
		Location:     "requested-region",
		PublicKey:    testSSHPublicKey,
		Tags: v1.Tags{
			"dev-plane-stage":         "dev",
			"dev-plane-x-cloudCredId": "tagged-credential",
		},
	})
	require.NoError(t, err)

	assert.Equal(t, "brevkey ref123", createdKey.Name)
	assert.Equal(t, testSSHPublicKey, createdKey.PublicKey)
	assert.Equal(t, "requested-region", launchRequest.RegionName)
	assert.Equal(t, "gpu_1x_l40", launchRequest.ProductName)
	assert.Equal(t, []string{"brevkey ref123"}, launchRequest.SshKeys)
	assert.Equal(t, int32(42), launchRequest.ImageId)
	require.NotNil(t, launchRequest.Command)
	assert.Contains(t, *launchRequest.Command, "base64 --decode | sudo -n bash")
	require.NotNil(t, launchRequest.InstanceName)
	assert.Equal(t, "dev_tagged-credential_ref-123", *launchRequest.InstanceName)

	assert.Equal(t, v1.CloudProviderInstanceID("instance-id"), instance.CloudID)
	assert.Equal(t, "requested-region", instance.Location)
	assert.Equal(t, "Ubuntu Server 22.04 w/ drivers", instance.ImageID)
	assert.Equal(t, "ref-123", instance.RefID)
	assert.Equal(t, "ref-123", instance.Name)
	assert.Equal(t, "tagged-credential", instance.CloudCredRefID)
	assert.Equal(t, "dev", instance.Tags["dev-plane-stage"])
	assert.Equal(t, "tagged-credential", instance.Tags["dev-plane-x-cloudCredId"])
	assert.False(t, instance.Spot)
}

func TestProviderInstanceName(t *testing.T) {
	providerName := makeProviderInstanceName("dev", "credential-ref", "ref-123")
	assert.Equal(t, "dev_credential-ref_ref-123", providerName)

	stage, cloudCredRefID, refID, err := parseProviderInstanceName(providerName)
	require.NoError(t, err)
	assert.Equal(t, "dev", stage)
	assert.Equal(t, "credential-ref", cloudCredRefID)
	assert.Equal(t, "ref-123", refID)
}

func TestResolveImageIDHonorsNumericOverride(t *testing.T) {
	server := newMassedComputeTestServer(t, func(http.ResponseWriter, *http.Request) {
		t.Error("numeric image IDs should not require an image catalog request")
	})
	defer server.Close()

	client := newMassedComputeTestClient(t, server, "")
	imageID, err := client.resolveImageID(context.Background(), "73")
	require.NoError(t, err)
	assert.Equal(t, int32(73), imageID)
}

func TestResolveImageName(t *testing.T) {
	server := newMassedComputeTestServer(t, func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)
		require.Equal(t, "/images", r.URL.Path)
		writeMassedComputeJSON(t, w, map[string]any{"images": []map[string]any{{
			"vm_image_id": 73, "vm_image_name": "Ubuntu Server 22.04 w/ drivers",
		}}})
	})
	defer server.Close()

	client := newMassedComputeTestClient(t, server, "")
	imageName, err := client.resolveImageName(context.Background(), 73)
	require.NoError(t, err)
	assert.Equal(t, "Ubuntu Server 22.04 w/ drivers", imageName)
}

func TestListGetAndTerminateInstance(t *testing.T) {
	var terminateRequest openapi.InstanceRestartPostRequest
	server := newMassedComputeTestServer(t, func(w http.ResponseWriter, r *http.Request) {
		switch {
		case r.Method == http.MethodGet && r.URL.Path == "/instance":
			writeMassedComputeJSON(t, w, map[string]any{"runningInstances": []map[string]any{{
				"uuid":     "instance-id",
				"name":     "dev_creator-credential_ref-123",
				"ip":       "192.0.2.10",
				"status":   "rented",
				"username": "ubuntu",
				"created":  "2026-08-25T12:34:56.000Z",
				"region":   map[string]any{"name": "us-central-1"},
				"image":    map[string]any{"id": 42, "name": "Ubuntu Server 22.04 w/ drivers"},
				"product": map[string]any{
					"name": "gpu_1x_l40", "vcpu": 14, "ram": 72, "storage": 625,
				},
			}}})
		case r.Method == http.MethodGet && r.URL.Path == "/instance/instance-id":
			writeMassedComputeJSON(t, w, map[string]any{"runningInstances": []map[string]any{{
				"uuid":     "instance-id",
				"name":     "dev_creator-credential_ref-123",
				"ip":       "192.0.2.10",
				"status":   "rented",
				"username": "ubuntu",
				"created":  "2026-08-25T12:34:56.000Z",
				"region":   map[string]any{"name": "us-central-1"},
				"image":    map[string]any{"id": 42, "name": "Ubuntu Server 22.04 w/ drivers"},
				"product": map[string]any{
					"name": "gpu_1x_l40", "vcpu": 14, "ram": 72, "storage": 625,
				},
			}}})
		case r.Method == http.MethodPost && r.URL.Path == "/instance/terminate":
			require.NoError(t, json.NewDecoder(r.Body).Decode(&terminateRequest))
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusAccepted)
			require.NoError(t, json.NewEncoder(w).Encode(map[string]any{"response": map[string]any{"data": map[string]any{"terminated_instances": []any{}}}}))
		default:
			http.NotFound(w, r)
		}
	})
	defer server.Close()

	client := newMassedComputeTestClient(t, server, "")
	ctx := context.Background()
	instances, err := client.ListInstances(ctx, v1.ListInstancesArgs{
		InstanceIDs: []v1.CloudProviderInstanceID{"instance-id"},
		Locations:   v1.LocationsFilter{"us-central-1"},
	})
	require.NoError(t, err)
	require.Len(t, instances, 1)
	instance := instances[0]
	assert.Equal(t, "ref-123", instance.RefID)
	assert.Equal(t, "ref-123", instance.Name)
	assert.Equal(t, "creator-credential", instance.CloudCredRefID)
	assert.Equal(t, v1.LifecycleStatusRunning, instance.Status.LifecycleStatus)
	assert.Equal(t, "ubuntu", instance.SSHUser)
	assert.Equal(t, "Ubuntu Server 22.04 w/ drivers", instance.ImageID)
	assert.Equal(t, "us-central-1", instance.Location)
	assert.Equal(t, 22, instance.SSHPort)
	assert.Equal(t, "ssd", instance.VolumeType)
	assert.Equal(t, v1.NewBytes(625, v1.Gigabyte), instance.DiskSizeBytes)
	assertLegacyBytesMatch(t, instance.DiskSize, instance.DiskSizeBytes)

	got, err := client.GetInstance(ctx, "instance-id")
	require.NoError(t, err)
	assert.Equal(t, instance.CloudID, got.CloudID)
	assert.Equal(t, instance.DiskSizeBytes, got.DiskSizeBytes)
	assert.Equal(t, instance.Location, got.Location)

	require.NoError(t, client.TerminateInstance(ctx, "instance-id"))
	assert.Equal(t, []string{"instance-id"}, terminateRequest.InstanceUuids)
}

func TestBuildStartupScript(t *testing.T) {
	script, err := buildStartupScript(v1.FirewallRules{
		IngressRules: []v1.FirewallRule{{
			FromPort: 8080,
			ToPort:   8081,
			IPRanges: []string{"192.0.2.7/24"},
		}},
	})
	require.NoError(t, err)

	assert.Contains(t, script, "passwd --lock ubuntu")
	assert.NotContains(t, script, "authorized_keys")
	assert.NotContains(t, script, "useradd")
	assert.Contains(t, script, "ufw allow from 192.0.2.0/24 to any port 8080:8081 proto tcp")
	assert.Contains(t, script, "iptables -A DOCKER-USER -s 192.0.2.0/24 -p tcp --dport 8080:8081 -j ACCEPT")
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

func TestSpotIsNotSupported(t *testing.T) {
	server := newMassedComputeTestServer(t, func(http.ResponseWriter, *http.Request) {
		t.Error("spot validation should fail before making an API request")
	})
	defer server.Close()

	client := newMassedComputeTestClient(t, server, "us-central-1")
	_, err := client.CreateInstance(context.Background(), v1.CreateInstanceAttrs{
		RefID:        "ref-123",
		InstanceType: "gpu_1x_l40_spot",
		PublicKey:    testSSHPublicKey,
	})
	require.ErrorContains(t, err, "spot instances are not supported")

	_, err = client.CreateInstance(context.Background(), v1.CreateInstanceAttrs{
		RefID:        "ref-123",
		InstanceType: "gpu_1x_l40",
		PublicKey:    testSSHPublicKey,
		UseSpot:      true,
	})
	require.ErrorContains(t, err, "spot instances are not supported")

	capabilities, err := client.GetCapabilities(context.Background())
	require.NoError(t, err)
	assert.False(t, capabilities.IsCapable(v1.CapabilityStopStartInstance))
}

func newMassedComputeTestServer(t *testing.T, handler http.HandlerFunc) *httptest.Server {
	t.Helper()
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "Bearer api-token", r.Header.Get("Authorization"))
		handler(w, r)
	}))
}

func newMassedComputeTestClient(t *testing.T, server *httptest.Server, location string) *MassedComputeClient {
	t.Helper()
	credential := NewMassedComputeCredential("credential-ref", "api-token")
	credential.APIURL = server.URL
	client, err := NewMassedComputeClient(*credential, location, WithHTTPClient(server.Client()))
	require.NoError(t, err)
	return client
}

func writeMassedComputeJSON(t *testing.T, w http.ResponseWriter, value any) {
	t.Helper()
	w.Header().Set("Content-Type", "application/json")
	require.NoError(t, json.NewEncoder(w).Encode(value))
}

func assertLegacyBytesMatch(t *testing.T, legacy units.Base2Bytes, size v1.Bytes) {
	t.Helper()
	assert.Equal(t, size.ByteCount().Int64(), int64(legacy))
}

func TestWrapMassedComputeError(t *testing.T) {
	err := wrapMassedComputeError(errors.New("no capacity available"), &http.Response{StatusCode: http.StatusConflict})
	require.ErrorIs(t, err, v1.ErrInsufficientResources)

	err = wrapMassedComputeError(errors.New("temporary failure"), &http.Response{StatusCode: http.StatusServiceUnavailable})
	require.ErrorIs(t, err, v1.ErrServiceUnavailable)
}
