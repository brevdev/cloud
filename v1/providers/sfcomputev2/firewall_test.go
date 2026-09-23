package v2

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	v1 "github.com/brevdev/cloud/v1"
	"github.com/stretchr/testify/require"
)

const (
	terminateTestInstanceRoute = "POST /integrations/brev/v1/instances/inst_test/terminate"
	createInstanceRoute        = "POST /integrations/brev/v1/instances"
	listInstancesRoute         = "GET /integrations/brev/v1/instances"
	getTestInstanceRoute       = "GET /integrations/brev/v1/instances/inst_test"
	createFirewallRoute        = "POST /integrations/brev/v1/firewalls"
	deleteTestFirewallRoute    = "DELETE /integrations/brev/v1/firewalls/frwl_test"
	getTestPoolRoute           = "GET /integrations/brev/v1/pools/sfc:pool:account:workspace:default"
)

func networkingTestClient(t *testing.T, enabled bool, handler http.HandlerFunc) *SFCClientV2 {
	t.Helper()
	server := httptest.NewServer(handler)
	t.Cleanup(server.Close)
	credential := NewSFCCredentialV2("cred", "key", "account", "workspace")
	credential.EnableConfigurableFirewall = enabled
	client, err := credential.MakeClient(context.Background(), sfcLocation)
	require.NoError(t, err)
	sfc, ok := client.(*SFCClientV2)
	require.True(t, ok)
	sfc.client.baseURL = server.URL
	return sfc
}

func networkingPool() poolResponse {
	return poolResponse{
		PublicIPv4SKUs: pointerTo([]string{"is_public"}),
		AllocationSchedule: allocationSchedule{ByInstanceSKU: map[string][]scheduleEntry{
			"is_legacy": {{StartAt: 0, NodeCount: 2}},
			"is_public": {{StartAt: 0, NodeCount: 2}},
		}},
	}
}

func TestNetworkingRequiresOptInAndServerSupport(t *testing.T) {
	t.Parallel()
	for _, enabled := range []bool{false, true} {
		t.Run(fmt.Sprint(enabled), func(t *testing.T) {
			t.Parallel()
			created := false
			client := networkingTestClient(t, enabled, func(w http.ResponseWriter, r *http.Request) {
				switch r.Method + " " + r.URL.Path {
				case getTestPoolRoute:
					pool := networkingPool()
					pool.PublicIPv4SKUs = nil
					writeJSON(t, w, pool)
				case listInstancesRoute:
					writeJSON(t, w, listInstancesResponse{})
				case createInstanceRoute:
					created = true
					var body map[string]any
					require.NoError(t, json.NewDecoder(r.Body).Decode(&body))
					require.NotContains(t, body, "enable_public_ipv4")
					require.NotContains(t, body, "firewall")
					writeJSON(t, w, instanceResponse{ID: "inst_test", Status: instanceStatusAwaitingAllocation})
				default:
					t.Errorf("unexpected request: %s %s", r.Method, r.URL.Path)
					http.NotFound(w, r)
				}
			})
			_, err := client.CreateInstance(context.Background(), v1.CreateInstanceAttrs{RefID: "test"})
			if enabled {
				require.ErrorContains(t, err, "not enabled")
				require.False(t, created)
			} else {
				require.NoError(t, err)
				require.True(t, created)
			}
		})
	}
}

func TestPublicInstanceLifecycle(t *testing.T) {
	t.Parallel()
	firewall := firewallResponse{ID: "frwl_test"}
	var instanceTags map[string]string
	deleted := false
	client := networkingTestClient(t, true, func(w http.ResponseWriter, r *http.Request) {
		switch r.Method + " " + r.URL.Path {
		case getTestPoolRoute:
			writeJSON(t, w, networkingPool())
		case listInstancesRoute:
			writeJSON(t, w, listInstancesResponse{})
		case createFirewallRoute:
			var body struct {
				Workspace string         `json:"workspace"`
				Rules     []firewallRule `json:"rules"`
			}
			require.NoError(t, json.NewDecoder(r.Body).Decode(&body))
			require.Equal(t, "sfc:workspace:account:workspace", body.Workspace)
			firewall.Rules = body.Rules
			require.Len(t, firewall.Rules, 4)
			require.Contains(t, firewall.Rules, firewallRule{"ingress", "udp", "8000-8010", "8.8.8.0/24"})
			writeJSON(t, w, firewall)
		case createInstanceRoute:
			var body createInstanceRequest
			require.NoError(t, json.NewDecoder(r.Body).Decode(&body))
			require.Equal(t, "is_public", body.InstanceSKU)
			require.True(t, body.EnablePublicIPv4)
			require.Equal(t, firewall.ID, body.Firewall)
			require.Equal(t, firewall.ID, body.Tags[tagKeyFirewallID])
			instanceTags = body.Tags
			writeJSON(t, w, instanceResponse{ID: "inst_test", EnablePublicIPv4: true, Firewall: firewall.ID, Tags: instanceTags})
		case getTestInstanceRoute:
			writeJSON(t, w, instanceResponse{
				ID: "inst_test", Status: instanceStatusRunning,
				EnablePublicIPv4: true, Firewall: firewall.ID, PublicIP: "192.0.2.10",
				Tags: instanceTags,
			})
		case "GET /integrations/brev/v1/firewalls/frwl_test":
			writeJSON(t, w, firewall)
		case terminateTestInstanceRoute:
			writeJSON(t, w, instanceResponse{
				ID: "inst_test", Status: instanceStatusTerminated,
				EnablePublicIPv4: true, Firewall: firewall.ID, Tags: instanceTags,
			})
		case deleteTestFirewallRoute:
			deleted = true
			w.WriteHeader(http.StatusNoContent)
		default:
			t.Errorf("unexpected request: %s %s", r.Method, r.URL.Path)
			http.NotFound(w, r)
		}
	})
	ctx := context.Background()
	capabilities, err := client.GetCapabilities(ctx)
	require.NoError(t, err)
	require.Contains(t, capabilities, v1.CapabilityModifyFirewall)
	_, err = client.CreateInstance(ctx, v1.CreateInstanceAttrs{RefID: "test", FirewallRules: v1.FirewallRules{
		IngressRules: []v1.FirewallRule{{FromPort: 8000, ToPort: 8010, IPRanges: []string{"8.8.8.1/24"}}},
	}})
	require.NoError(t, err)
	instance, err := client.GetInstance(ctx, "inst_test")
	require.NoError(t, err)
	require.Equal(t, "192.0.2.10", instance.PublicIP)
	require.Equal(t, 22, instance.SSHPort)
	require.Len(t, instance.FirewallRules.IngressRules, 1)
	require.NotEmpty(t, instance.FirewallRules.IngressRules[0].ID)
	require.NotContains(t, instance.Tags, tagKeyFirewallID)
	require.NoError(t, client.TerminateInstance(ctx, "inst_test"))
	require.True(t, deleted)
}

func TestFirewallUpdateRetriesWithoutLosingConcurrentRules(t *testing.T) {
	t.Parallel()
	firewall := firewallResponse{ID: "frwl_test", Rules: sshFirewallRules()}
	version := 1
	updates := 0
	concurrentRule := firewallRule{"ingress", "tcp", "9000", "8.8.8.0/24"}
	client := networkingTestClient(t, true, func(w http.ResponseWriter, r *http.Request) {
		switch r.Method + " " + r.URL.Path {
		case getTestInstanceRoute:
			writeJSON(t, w, instanceResponse{ID: "inst_test", EnablePublicIPv4: true, Firewall: firewall.ID})
		case "GET /integrations/brev/v1/firewalls/frwl_test":
			w.Header().Set(firewallVersionHeader, fmt.Sprint(version))
			writeJSON(t, w, firewall)
		case "PUT /integrations/brev/v1/firewalls/frwl_test":
			updates++
			require.Equal(t, fmt.Sprint(version), r.Header.Get(firewallVersionHeader))
			if updates == 1 {
				firewall.Rules = append(firewall.Rules, concurrentRule)
				version++
				w.WriteHeader(http.StatusConflict)
				return
			}
			var body firewallResponse
			require.NoError(t, json.NewDecoder(r.Body).Decode(&body))
			firewall.Rules = body.Rules
			version++
			writeJSON(t, w, firewall)
		default:
			t.Errorf("unexpected request: %s %s", r.Method, r.URL.Path)
			http.NotFound(w, r)
		}
	})
	ctx := context.Background()
	require.NoError(t, client.AddFirewallRulesToInstance(ctx, v1.AddFirewallRulesToInstanceArgs{
		InstanceID: "inst_test", FirewallRules: v1.FirewallRules{IngressRules: []v1.FirewallRule{{FromPort: 8080, ToPort: 8080}}},
	}))
	require.Equal(t, 2, updates)
	require.Contains(t, firewall.Rules, concurrentRule)
	require.Len(t, firewall.Rules, 5)
	id := ingressRuleID(firewallRule{Port: "8080", Source: "0.0.0.0/0"})
	require.NoError(t, client.RevokeSecurityGroupRules(ctx, v1.RevokeSecurityGroupRuleArgs{
		InstanceID: "inst_test", SecurityGroupRuleIDs: []string{id},
	}))
	require.Len(t, firewall.Rules, 3)
	require.Contains(t, firewall.Rules, concurrentRule)
	for _, rule := range sshFirewallRules() {
		require.Contains(t, firewall.Rules, rule)
	}
}

func TestNetworkingRejectsInvalidRules(t *testing.T) {
	t.Parallel()
	for _, rule := range []v1.FirewallRule{
		{FromPort: -1, ToPort: 22},
		{FromPort: 20, ToPort: 10},
		{FromPort: 1, ToPort: 65536},
		{FromPort: 80, ToPort: 80, IPRanges: []string{"::/0"}},
		{FromPort: 80, ToPort: 80, IPRanges: []string{"192.0.2.1"}},
	} {
		_, err := expandFirewallRules(v1.FirewallRules{IngressRules: []v1.FirewallRule{rule}})
		require.Error(t, err)
	}
	_, err := expandFirewallRules(v1.FirewallRules{EgressRules: []v1.FirewallRule{{FromPort: 80, ToPort: 80}}})
	require.ErrorContains(t, err, "outbound")
}

func TestFirewallUpdateRequiresServerVersion(t *testing.T) {
	t.Parallel()
	client := networkingTestClient(t, true, func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)
		if r.URL.Path == "/integrations/brev/v1/instances/inst_test" {
			writeJSON(t, w, instanceResponse{ID: "inst_test", EnablePublicIPv4: true, Firewall: "frwl_test"})
		} else {
			writeJSON(t, w, firewallResponse{ID: "frwl_test", Rules: sshFirewallRules()})
		}
	})
	err := client.AddFirewallRulesToInstance(context.Background(), v1.AddFirewallRulesToInstanceArgs{
		InstanceID: "inst_test", FirewallRules: v1.FirewallRules{IngressRules: []v1.FirewallRule{{FromPort: 8080, ToPort: 8080}}},
	})
	require.ErrorContains(t, err, "conditional firewall updates")
}
