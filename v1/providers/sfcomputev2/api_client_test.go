package v2

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAPIClientUsesBrevContract(t *testing.T) {
	t.Parallel()

	handler := http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		require.Equal(t, "Bearer api-key", request.Header.Get("Authorization"))

		switch request.Method + " " + request.URL.Path {
		case "POST /integrations/brev/v1/instances":
			var body map[string]any
			require.NoError(t, json.NewDecoder(request.Body).Decode(&body))
			require.Equal(t, "sfc:pool:account:workspace:default", body["pool"])
			require.Equal(t, "sfc:image:sfcompute:public:ubuntu", body["image"])
			require.Equal(t, "is_sku", body["instance_sku"])
			require.Equal(t, "cloud-init", body["cloud_init_user_data"])
			tags, ok := body["tags"].(map[string]any)
			require.True(t, ok)
			require.Equal(t, "brev-ref", tags[tagKeyRefID])
			require.Equal(t, false, body["_preview_enable_infiniband"])
			writeJSON(t, writer, instanceResponse{ID: "inst_created", Status: instanceStatusAwaitingAllocation})
		case "GET /integrations/brev/v1/instances":
			require.Equal(t, "sfc:workspace:account:workspace", request.URL.Query().Get("workspace"))
			require.Equal(t, []string{"sfc:pool:account:workspace:default"}, request.URL.Query()["pool"])
			require.Equal(t, "200", request.URL.Query().Get("limit"))
			if request.URL.Query().Get("starting_after") == "" {
				writeJSON(t, writer, listInstancesResponse{
					Cursor:  pointerTo("next-page"),
					HasMore: true,
					Data:    []instanceResponse{{ID: "inst_listed_1"}},
				})
				return
			}
			require.Equal(t, "next-page", request.URL.Query().Get("starting_after"))
			writeJSON(t, writer, listInstancesResponse{Data: []instanceResponse{{ID: "inst_listed_2"}}})
		case "GET /integrations/brev/v1/instances/inst_test":
			writeJSON(t, writer, instanceResponse{ID: "inst_test", Status: instanceStatusRunning})
		case "GET /integrations/brev/v1/instances/inst_test/ssh":
			writeJSON(t, writer, instanceSSHInfo{Hostname: "192.0.2.1", Port: 22})
		case "POST /integrations/brev/v1/instances/inst_test/terminate":
			writeJSON(t, writer, instanceResponse{ID: "inst_test", Status: instanceStatusTerminated})
		case "GET /integrations/brev/v1/pools/sfc:pool:account:workspace:default":
			writeJSON(t, writer, poolResponse{AllocationSchedule: allocationSchedule{
				ByInstanceSKU: map[string][]scheduleEntry{"is_sku": {{StartAt: 0, NodeCount: 1}}},
			}})
		default:
			http.NotFound(writer, request)
		}
	})
	server := httptest.NewServer(handler)
	t.Cleanup(server.Close)

	client := newAPIClient("api-key")
	client.baseURL = server.URL
	ctx := context.Background()

	created, err := client.createInstance(ctx, createInstanceRequest{
		Pool:              "sfc:pool:account:workspace:default",
		Image:             "sfc:image:sfcompute:public:ubuntu",
		InstanceSKU:       "is_sku",
		CloudInitUserData: pointerTo("cloud-init"),
		Tags:              map[string]string{tagKeyRefID: "brev-ref"},
	})
	require.NoError(t, err)
	require.Equal(t, "inst_created", created.ID)

	listed, err := client.listInstances(
		ctx,
		"sfc:workspace:account:workspace",
		"sfc:pool:account:workspace:default",
	)
	require.NoError(t, err)
	require.Equal(t, []instanceResponse{{ID: "inst_listed_1"}, {ID: "inst_listed_2"}}, listed.Data)

	instance, err := client.getInstance(ctx, "inst_test")
	require.NoError(t, err)
	require.Equal(t, instanceStatusRunning, instance.Status)

	sshInfo, err := client.getSSHInfo(ctx, "inst_test")
	require.NoError(t, err)
	require.Equal(t, "192.0.2.1", sshInfo.Hostname)

	require.NoError(t, client.terminateInstance(ctx, "inst_test"))

	pool, err := client.getPool(ctx, "sfc:pool:account:workspace:default")
	require.NoError(t, err)
	require.Equal(t, 1, pool.AllocationSchedule.ByInstanceSKU["is_sku"][0].NodeCount)
}

func TestAPIClientReturnsResponseErrors(t *testing.T) {
	t.Parallel()

	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, _ *http.Request) {
		http.Error(writer, `{"error":"not found"}`, http.StatusNotFound)
	}))
	t.Cleanup(server.Close)

	client := newAPIClient("api-key")
	client.baseURL = server.URL

	_, err := client.getInstance(context.Background(), "inst_missing")
	var responseError *apiError
	require.ErrorAs(t, err, &responseError)
	require.Equal(t, http.StatusNotFound, responseError.statusCode)
	require.Contains(t, responseError.body, "not found")
}

func writeJSON(t *testing.T, writer http.ResponseWriter, value any) {
	t.Helper()
	writer.Header().Set("Content-Type", "application/json")
	require.NoError(t, json.NewEncoder(writer).Encode(value))
}

func pointerTo[T any](value T) *T {
	return &value
}
