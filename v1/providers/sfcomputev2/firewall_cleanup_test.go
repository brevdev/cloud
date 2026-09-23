package v2

import (
	"context"
	"net/http"
	"testing"

	v1 "github.com/brevdev/cloud/v1"
	"github.com/stretchr/testify/require"
)

func TestNetworkingCleansUpRejectedOrUnacknowledgedCreates(t *testing.T) {
	t.Parallel()
	for _, rejected := range []bool{false, true} {
		name := "unacknowledged"
		if rejected {
			name = "rejected"
		}
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			terminated, deleted := false, false
			client := networkingTestClient(t, true, func(w http.ResponseWriter, r *http.Request) {
				switch r.Method + " " + r.URL.Path {
				case getTestPoolRoute:
					writeJSON(t, w, networkingPool())
				case listInstancesRoute:
					writeJSON(t, w, listInstancesResponse{})
				case createFirewallRoute:
					writeJSON(t, w, firewallResponse{ID: "frwl_test"})
				case createInstanceRoute:
					if rejected {
						w.WriteHeader(http.StatusUnprocessableEntity)
						return
					}
					writeJSON(t, w, instanceResponse{ID: "inst_test", Status: instanceStatusAwaitingAllocation})
				case terminateTestInstanceRoute:
					terminated = true
					writeJSON(t, w, instanceResponse{ID: "inst_test", Status: instanceStatusTerminated})
				case deleteTestFirewallRoute:
					deleted = true
					w.WriteHeader(http.StatusNoContent)
				default:
					t.Errorf("unexpected request: %s %s", r.Method, r.URL.Path)
					http.NotFound(w, r)
				}
			})
			_, err := client.CreateInstance(context.Background(), v1.CreateInstanceAttrs{RefID: "test"})
			require.Error(t, err)
			require.True(t, deleted)
			require.Equal(t, !rejected, terminated)
		})
	}
}

func TestLegacyTerminationDoesNotRequireFirewallPermissions(t *testing.T) {
	t.Parallel()
	client := networkingTestClient(t, false, func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodPost, r.Method)
		require.Equal(t, "/integrations/brev/v1/instances/inst_test/terminate", r.URL.Path)
		writeJSON(t, w, instanceResponse{ID: "inst_test", Status: instanceStatusTerminated})
	})
	require.NoError(t, client.TerminateInstance(context.Background(), "inst_test"))
}

func TestCanceledCreateStillCleansUpFirewall(t *testing.T) {
	t.Parallel()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	deleted := false
	client := networkingTestClient(t, true, func(w http.ResponseWriter, r *http.Request) {
		switch r.Method + " " + r.URL.Path {
		case getTestPoolRoute:
			writeJSON(t, w, networkingPool())
		case listInstancesRoute:
			writeJSON(t, w, listInstancesResponse{})
		case createFirewallRoute:
			writeJSON(t, w, firewallResponse{ID: "frwl_test"})
		case createInstanceRoute:
			cancel()
			w.WriteHeader(http.StatusUnprocessableEntity)
		case deleteTestFirewallRoute:
			deleted = true
			w.WriteHeader(http.StatusNoContent)
		default:
			t.Errorf("unexpected request: %s %s", r.Method, r.URL.Path)
			http.NotFound(w, r)
		}
	})
	_, err := client.CreateInstance(ctx, v1.CreateInstanceAttrs{RefID: "test"})
	require.Error(t, err)
	require.True(t, deleted, "cleanup must survive cancellation of the create request")
}

func TestTerminationRetriesFailedFirewallCleanup(t *testing.T) {
	t.Parallel()
	deletes := 0
	client := networkingTestClient(t, false, func(w http.ResponseWriter, r *http.Request) {
		switch r.Method + " " + r.URL.Path {
		case terminateTestInstanceRoute:
			writeJSON(t, w, instanceResponse{
				ID: "inst_test", Status: instanceStatusTerminated,
				EnablePublicIPv4: true, Firewall: "frwl_test",
				Tags: map[string]string{tagKeyFirewallID: "frwl_test"},
			})
		case deleteTestFirewallRoute:
			deletes++
			if deletes == 1 {
				w.WriteHeader(http.StatusServiceUnavailable)
			} else {
				w.WriteHeader(http.StatusNoContent)
			}
		default:
			t.Errorf("unexpected request: %s %s", r.Method, r.URL.Path)
			http.NotFound(w, r)
		}
	})
	require.ErrorContains(t, client.TerminateInstance(context.Background(), "inst_test"), "frwl_test")
	require.NoError(t, client.TerminateInstance(context.Background(), "inst_test"))
	require.Equal(t, 2, deletes)
}

func TestTerminationPreservesUnmanagedFirewalls(t *testing.T) {
	t.Parallel()
	for _, ownedFirewall := range []string{"", "frwl_previous"} {
		t.Run("owned="+ownedFirewall, func(t *testing.T) {
			t.Parallel()
			client := networkingTestClient(t, false, func(w http.ResponseWriter, r *http.Request) {
				if r.Method+" "+r.URL.Path != terminateTestInstanceRoute {
					t.Errorf("must not delete an unmanaged firewall: %s %s", r.Method, r.URL.Path)
					http.NotFound(w, r)
					return
				}
				writeJSON(t, w, instanceResponse{
					ID: "inst_test", Status: instanceStatusTerminated,
					EnablePublicIPv4: true, Firewall: "frwl_test",
					Tags: map[string]string{tagKeyFirewallID: ownedFirewall},
				})
			})
			require.NoError(t, client.TerminateInstance(context.Background(), "inst_test"))
		})
	}
}
