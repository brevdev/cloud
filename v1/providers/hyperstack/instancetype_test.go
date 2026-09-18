package hyperstack

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	v1 "github.com/brevdev/cloud/v1"
)

func TestGetInstanceTypesAndLocations(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, request *http.Request) {
		assert.Equal(t, "api-key", request.Header.Get("api_key"))
		assert.Equal(t, "brev-cloud", request.Header.Get("User-Agent"))
		switch request.URL.Path {
		case "/v1/core/flavors":
			writeJSON(t, w, map[string]any{
				"status": true,
				"data": []map[string]any{
					{
						"region_name": "CANADA-1",
						"gpu":         "H100-80G-PCIe",
						"flavors": []map[string]any{{
							"id":              1,
							"name":            "n3-H100x2",
							"region_name":     "CANADA-1",
							"cpu":             56,
							"ram":             360,
							"disk":            100,
							"ephemeral":       1500,
							"gpu":             "H100-80G-PCIe",
							"gpu_count":       2,
							"stock_available": true,
						}},
					},
					{
						"region_name": "NORWAY-1",
						"gpu":         "",
						"flavors": []map[string]any{{
							"id":              2,
							"name":            "n1-cpu-small",
							"region_name":     "NORWAY-1",
							"cpu":             4,
							"ram":             4,
							"disk":            100,
							"ephemeral":       0,
							"gpu":             "",
							"gpu_count":       0,
							"stock_available": false,
						}},
					},
				},
			})
		case "/v1/pricebook":
			writeJSON(t, w, []map[string]any{
				{"name": "H100-80G-PCIe", "value": "2.5"},
				{"name": "vCPU (cpu-only-flavors)", "value": "0.01"},
				{"name": "RAM (cpu-only-flavors)", "value": "0.02"},
				{"name": "hypervisor-local-storage (cpu-only-flavors)", "value": "0.001"},
			})
		case "/v1/core/regions":
			writeJSON(t, w, map[string]any{
				"status": true,
				"regions": []map[string]any{
					{"id": 1, "name": "CANADA-1", "country": "CA", "features": map[string]any{"floating_ip": true}},
					{"id": 2, "name": "NORWAY-1", "country": "NO", "features": map[string]any{"floating_ip": true}},
					{"id": 3, "name": "CANADA-2", "country": "CA", "features": map[string]any{"floating_ip": false}},
				},
			})
		default:
			http.NotFound(w, request)
		}
	}))
	defer server.Close()

	client := newTestClient(t, server.URL+"/v1")
	instanceTypes, err := client.GetInstanceTypes(context.Background(), v1.GetInstanceTypeArgs{})
	require.NoError(t, err)
	require.Len(t, instanceTypes, 2)

	gpuType := instanceTypes[0]
	assert.Equal(t, v1.InstanceTypeID("CANADA-1-noSub-n3-H100x2"), gpuType.ID)
	assert.Equal(t, v1.NewBytes(360, v1.Gigabyte), gpuType.MemoryBytes)
	assert.Equal(t, "5.0", gpuType.BasePrice.Number())
	require.Len(t, gpuType.SupportedGPUs, 1)
	assert.Equal(t, int32(2), gpuType.SupportedGPUs[0].Count)
	assert.Equal(t, v1.NewBytes(80, v1.Gigabyte), gpuType.SupportedGPUs[0].MemoryBytes)
	assert.Equal(t, "H100", gpuType.SupportedGPUs[0].Name)
	assert.Equal(t, "PCIe", gpuType.SupportedGPUs[0].NetworkDetails)
	require.Len(t, gpuType.SupportedStorage, 1)
	assert.Equal(t, "ephemeral", gpuType.SupportedStorage[0].Type)
	assert.Equal(t, v1.NewBytes(1500, v1.Gigabyte), gpuType.SupportedStorage[0].SizeBytes)
	assert.True(t, gpuType.SupportedStorage[0].IsEphemeral)

	cpuType := instanceTypes[1]
	assert.Equal(t, "0.220", cpuType.BasePrice.Number())
	assert.False(t, cpuType.IsAvailable)
	require.Len(t, cpuType.SupportedStorage, 1)
	assert.Equal(t, "ssd", cpuType.SupportedStorage[0].Type)
	assert.Equal(t, v1.NewBytes(100, v1.Gigabyte), cpuType.SupportedStorage[0].SizeBytes)
	assert.False(t, cpuType.SupportedStorage[0].IsEphemeral)

	filteredTypes, err := client.GetInstanceTypes(context.Background(), v1.GetInstanceTypeArgs{
		Locations:     v1.LocationsFilter{"NORWAY-1"},
		InstanceTypes: []string{"n1-cpu-small"},
	})
	require.NoError(t, err)
	require.Len(t, filteredTypes, 1)
	assert.Equal(t, "NORWAY-1", filteredTypes[0].Location)

	locations, err := client.GetLocations(context.Background(), v1.GetLocationsArgs{IncludeUnavailable: true})
	require.NoError(t, err)
	require.Len(t, locations, 3)
	assert.Equal(t, "CANADA-1", locations[0].Name)
	assert.True(t, locations[0].Available)
	assert.Equal(t, "CAN", locations[0].Country)
	assert.False(t, locations[1].Available)
	assert.False(t, locations[2].Available)
}

func newTestClient(t *testing.T, apiURL string) *HyperstackClient {
	t.Helper()
	credential := NewHyperstackCredential("credential-ref", "api-key")
	credential.APIURL = apiURL
	client, err := NewHyperstackClient(*credential, "")
	require.NoError(t, err)
	return client
}

func writeJSON(t *testing.T, writer http.ResponseWriter, value any) {
	t.Helper()
	writer.Header().Set("Content-Type", "application/json")
	require.NoError(t, json.NewEncoder(writer).Encode(value))
}
