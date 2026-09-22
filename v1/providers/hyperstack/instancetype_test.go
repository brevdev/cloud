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

func TestGetInstanceTypesMapsFlavor(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, request *http.Request) {
		assert.Equal(t, "api-key", request.Header.Get("api_key"))
		assert.Equal(t, "brev-cloud", request.Header.Get("User-Agent"))
		switch request.URL.Path {
		case "/v1/core/flavors":
			writeJSON(t, w, map[string]any{
				"status": true,
				"data": []map[string]any{{
					"region_name": "CANADA-1",
					"flavors": []map[string]any{{
						"name":            "n3-H100x2",
						"cpu":             56,
						"ram":             360,
						"disk":            100,
						"ephemeral":       1500,
						"gpu":             "H100-80G-PCIe",
						"gpu_count":       2,
						"stock_available": true,
					}},
				}},
			})
		case "/v1/pricebook":
			writeJSON(t, w, []map[string]any{{"name": "H100-80G-PCIe", "value": "2.5"}})
		default:
			http.NotFound(w, request)
		}
	}))
	defer server.Close()

	client := newTestClient(t, server.URL+"/v1")
	instanceTypes, err := client.GetInstanceTypes(context.Background(), v1.GetInstanceTypeArgs{})
	require.NoError(t, err)
	require.Len(t, instanceTypes, 1)

	instanceType := instanceTypes[0]
	assert.Equal(t, v1.InstanceTypeID("CANADA-1-noSub-n3-H100x2"), instanceType.ID)
	assert.Equal(t, "5.0", instanceType.BasePrice.Number())
	require.Len(t, instanceType.SupportedGPUs, 1)
	assert.Equal(t, "H100", instanceType.SupportedGPUs[0].Name)
	assert.Equal(t, int32(2), instanceType.SupportedGPUs[0].Count)
	require.Len(t, instanceType.SupportedStorage, 2)
	assert.Equal(t, "ssd", instanceType.SupportedStorage[0].Type)
	assert.Equal(t, "ephemeral", instanceType.SupportedStorage[1].Type)
}

func TestGetLocationsMapsRegions(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, request *http.Request) {
		switch request.URL.Path {
		case "/v1/core/regions":
			writeJSON(t, w, map[string]any{
				"status": true,
				"regions": []map[string]any{
					{"name": "CANADA-1", "country": "CA", "features": map[string]any{"floating_ip": true}},
					{"name": "NORWAY-1", "country": "NO", "features": map[string]any{"floating_ip": true}},
				},
			})
		case "/v1/core/flavors":
			writeJSON(t, w, map[string]any{
				"status": true,
				"data": []map[string]any{{
					"region_name": "CANADA-1",
					"flavors":     []map[string]any{{"name": "n1-cpu-small", "stock_available": true}},
				}},
			})
		default:
			http.NotFound(w, request)
		}
	}))
	defer server.Close()

	client := newTestClient(t, server.URL+"/v1")
	locations, err := client.GetLocations(context.Background(), v1.GetLocationsArgs{IncludeUnavailable: true})
	require.NoError(t, err)
	require.Len(t, locations, 2)
	assert.Equal(t, "CANADA-1", locations[0].Name)
	assert.True(t, locations[0].Available)
	assert.Equal(t, "CAN", locations[0].Country)
	assert.False(t, locations[1].Available)
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
