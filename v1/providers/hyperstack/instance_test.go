package hyperstack

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	virtualmachine "github.com/NexGenCloud/hyperstack-sdk-go/lib/virtual_machine"
	"github.com/alecthomas/units"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	v1 "github.com/brevdev/cloud/v1"
)

const testSSHPublicKey = "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIDBdptDTzJ2cOmdyryG1B7yb1YssiCQs6SWu4HlbZXGE"

func TestInstanceLifecycleRequests(t *testing.T) { //nolint:funlen // one stateful server makes the lifecycle easy to verify
	var labels []string
	deleted := false
	keyPairDeleteCount := 0
	stopCount := 0
	startCount := 0
	providerStatus := "ACTIVE"
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		assert.Equal(t, "api-key", request.Header.Get("api_key"))
		switch {
		case request.URL.Path == "/v1/core/environments" && request.Method == http.MethodGet:
			assert.Equal(t, "default-CANADA-1", request.URL.Query().Get("search"))
			writeJSON(t, writer, map[string]any{"status": true, "environments": []map[string]any{{
				"id": 5, "name": "default-CANADA-1", "region": "CANADA-1",
			}}})
		case request.URL.Path == "/v1/core/keypairs" && request.Method == http.MethodGet:
			writeJSON(t, writer, map[string]any{"status": true, "keypairs": []any{}})
		case request.URL.Path == "/v1/core/keypairs" && request.Method == http.MethodPost:
			var payload map[string]any
			require.NoError(t, json.NewDecoder(request.Body).Decode(&payload))
			assert.Equal(t, "default-CANADA-1", payload["environment_name"])
			assert.Equal(t, "ref-123", payload["name"])
			assert.Equal(t, testSSHPublicKey, payload["public_key"])
			writeJSON(t, writer, map[string]any{"status": true, "keypair": map[string]any{"id": 7, "name": payload["name"]}})
		case request.URL.Path == "/v1/core/keypair/7" && request.Method == http.MethodDelete:
			keyPairDeleteCount++
			writeJSON(t, writer, map[string]any{"status": true})
		case request.URL.Path == "/v1/core/virtual-machines" && request.Method == http.MethodPost:
			var payload virtualmachine.CreateInstancesPayload
			require.NoError(t, json.NewDecoder(request.Body).Decode(&payload))
			assert.Equal(t, "ref-123", payload.Name)
			assert.Equal(t, "default-CANADA-1", payload.EnvironmentName)
			assert.Equal(t, "n3-H100x1", payload.FlavorName)
			assert.Equal(t, defaultImageName, stringValue(payload.ImageName))
			require.NotNil(t, payload.AssignFloatingIp)
			assert.True(t, *payload.AssignFloatingIp)
			require.NotNil(t, payload.EnablePortRandomization)
			assert.False(t, *payload.EnablePortRandomization)
			require.NotNil(t, payload.EnhancedMonitoringEnabled)
			assert.False(t, *payload.EnhancedMonitoringEnabled)
			require.NotNil(t, payload.SecurityRules)
			require.Len(t, *payload.SecurityRules, 1)
			assertSecurityRule(t, (*payload.SecurityRules)[0], "10.0.0.0/8", 8080, 8080)
			require.NotNil(t, payload.Labels)
			labels = *payload.Labels
			require.NotNil(t, payload.UserData)
			assert.Contains(t, *payload.UserData, readinessMarker)
			assert.Contains(t, *payload.UserData, "WantedBy=multi-user.target")
			assert.Contains(t, *payload.UserData, "systemctl, enable, brev-cloud-ready.service")
			writeJSON(t, writer, map[string]any{"status": true, "instances": []map[string]any{{"id": 42}}})
		case request.URL.Path == "/v1/core/virtual-machines/42/logs" && request.Method == http.MethodPost:
			var payload virtualmachine.RequestInstanceLogsPayload
			require.NoError(t, json.NewDecoder(request.Body).Decode(&payload))
			require.NotNil(t, payload.Length)
			assert.Equal(t, consoleLogLineCount, *payload.Length)
			writeJSON(t, writer, map[string]any{"request_id": 99})
		case request.URL.Path == "/v1/core/virtual-machines/42/logs" && request.Method == http.MethodGet:
			assert.Equal(t, "99", request.URL.Query().Get("request_id"))
			writeJSON(t, writer, map[string]any{"logs": "boot output\n" + readinessMarker + "\n"})
		case request.URL.Path == "/v1/core/virtual-machines/42" && request.Method == http.MethodGet:
			if deleted {
				writer.WriteHeader(http.StatusNotFound)
				writeJSON(t, writer, map[string]any{"status": false, "message": "not found"})
				return
			}
			writeJSON(t, writer, map[string]any{
				"status": true,
				"instance": map[string]any{
					"id":          42,
					"name":        "ref-123",
					"status":      providerStatus,
					"created_at":  "2026-09-04T12:00:00",
					"floating_ip": "203.0.113.42",
					"fixed_ip":    "10.0.0.42",
					"environment": map[string]any{"name": "default-CANADA-1", "region": "CANADA-1"},
					"image":       map[string]any{"name": defaultImageName},
					"flavor":      map[string]any{"name": "n3-H100x1", "disk": 100},
					"labels":      labels,
					"security_rules": []map[string]any{{
						"id": 8, "direction": "ingress", "protocol": "tcp", "port_range_min": 8080,
						"port_range_max": 8080, "remote_ip_prefix": "10.0.0.0/8",
					}},
				},
			})
		case request.URL.Path == "/v1/core/virtual-machines" && request.Method == http.MethodGet:
			writeJSON(t, writer, map[string]any{
				"status": true,
				"instances": []map[string]any{{
					"id": 42, "name": "ref-123", "status": "ACTIVE", "created_at": "2026-09-04T12:00:00",
					"environment": map[string]any{"region": "CANADA-1"},
					"flavor":      map[string]any{"name": "n3-H100x1", "disk": 100}, "image": map[string]any{"name": defaultImageName},
					"labels": labels,
				}},
			})
		case request.URL.Path == "/v1/core/virtual-machines/42" && request.Method == http.MethodDelete:
			if deleted {
				writer.WriteHeader(http.StatusNotFound)
				writeJSON(t, writer, map[string]any{"status": false, "message": "not found"})
				return
			}
			deleted = true
			writeJSON(t, writer, map[string]any{"status": true})
		case request.URL.Path == "/v1/core/virtual-machines/42/stop" && request.Method == http.MethodGet:
			stopCount++
			providerStatus = "SHUTOFF"
			writeJSON(t, writer, map[string]any{"status": true})
		case request.URL.Path == "/v1/core/virtual-machines/42/start" && request.Method == http.MethodGet:
			startCount++
			providerStatus = "ACTIVE"
			writeJSON(t, writer, map[string]any{"status": true})
		default:
			http.NotFound(writer, request)
		}
	}))
	defer server.Close()

	client := newTestClient(t, server.URL+"/v1")
	instance, err := client.CreateInstance(context.Background(), v1.CreateInstanceAttrs{
		Location:      "CANADA-1",
		Name:          "display-name",
		RefID:         "ref-123",
		PublicKey:     testSSHPublicKey,
		InstanceType:  "n3-H100x1",
		DiskSize:      256 * units.Gibibyte,
		DiskSizeBytes: v1.NewBytes(256, v1.Gibibyte),
		Tags: v1.Tags{
			"dev-plane-managedBy":    "dev-plane",
			"dev-plane-x-instanceId": "instance-123",
			"team":                   "compute",
		},
		FirewallRules: v1.FirewallRules{IngressRules: []v1.FirewallRule{{
			FromPort: 8080,
			ToPort:   8080,
			IPRanges: []string{"10.0.0.0/8"},
		}}},
	})
	require.NoError(t, err)
	assert.Equal(t, v1.CloudProviderInstanceID("42"), instance.CloudID)
	assert.Equal(t, "ref-123", instance.Name)
	assert.Equal(t, "ref-123", instance.RefID)
	assert.Equal(t, "credential-ref", instance.CloudCredRefID)
	assert.Equal(t, "dev-plane", instance.Tags["dev-plane-managedBy"])
	assert.Equal(t, "instance-123", instance.Tags["dev-plane-x-instanceId"])
	assert.NotContains(t, instance.Tags, "team")
	assert.Equal(t, v1.LifecycleStatusRunning, instance.Status.LifecycleStatus)
	assert.Equal(t, "203.0.113.42", instance.PublicIP)
	assert.Equal(t, v1.InstanceTypeID("CANADA-1-noSub-n3-H100x1"), instance.InstanceTypeID)
	assert.True(t, instance.Stoppable)

	instances, err := client.ListInstances(context.Background(), v1.ListInstancesArgs{
		InstanceIDs: []v1.CloudProviderInstanceID{"42"},
		Locations:   v1.LocationsFilter{"CANADA-1"},
		TagFilters:  map[string][]string{"dev-plane-managedBy": {"dev-plane"}},
	})
	require.NoError(t, err)
	require.Len(t, instances, 1)
	assert.Equal(t, instance.RefID, instances[0].RefID)
	require.NoError(t, client.StopInstance(context.Background(), "42"))
	require.NoError(t, client.StartInstance(context.Background(), "42"))
	require.NoError(t, client.StartInstance(context.Background(), "42"))
	assert.Equal(t, 1, stopCount)
	assert.Equal(t, 1, startCount)

	require.NoError(t, client.TerminateInstance(context.Background(), "42"))
	require.NoError(t, client.TerminateInstance(context.Background(), "42"))
	assert.Equal(t, 1, keyPairDeleteCount)
	_, err = client.GetInstance(context.Background(), "42")
	require.Error(t, err)
	assert.True(t, errors.Is(err, v1.ErrInstanceNotFound))
}

func TestResolveKeyPairSearchesEveryPage(t *testing.T) {
	listCalls := 0
	importCalls := 0
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		require.Equal(t, "/v1/core/keypairs", request.URL.Path)
		switch request.Method {
		case http.MethodGet:
			listCalls++
			assert.Equal(t, "100", request.URL.Query().Get("pageSize"))
			assert.Equal(t, "ref", request.URL.Query().Get("search"))
			if request.URL.Query().Get("page") == "1" {
				keys := make([]map[string]any, defaultPageSize)
				for index := range keys {
					keys[index] = map[string]any{"name": "unrelated-key"}
				}
				writeJSON(t, writer, map[string]any{"status": true, "keypairs": keys})
				return
			}
			writeJSON(t, writer, map[string]any{"status": true, "keypairs": []map[string]any{{
				"id": 7, "name": "ref", "public_key": testSSHPublicKey,
				"environment": map[string]any{"name": "default-CANADA-1"},
			}}})
		case http.MethodPost:
			importCalls++
			writeJSON(t, writer, map[string]any{"status": true})
		default:
			http.NotFound(writer, request)
		}
	}))
	defer server.Close()

	client := newTestClient(t, server.URL+"/v1")
	keyPair, err := client.resolveKeyPair(context.Background(), v1.CreateInstanceAttrs{
		RefID:     "ref",
		PublicKey: testSSHPublicKey,
	}, "default-CANADA-1")
	require.NoError(t, err)
	assert.Equal(t, "ref", keyPair.name)
	assert.Equal(t, 7, keyPair.managedID)
	assert.Equal(t, 2, listCalls)
	assert.Zero(t, importCalls)
}

func TestLabelsRoundTripPlainValues(t *testing.T) {
	const refID = "82d299a7-dfd9-40e6-8707-3c477374b2a6"
	tags := v1.Tags{
		"dev-plane-managedBy":       "dev-plane",
		"dev-plane-x-instanceId":    "instance-id",
		"dev-plane-x-environmentId": "environment-id",
		"dev-plane-x-userId":        "user-id",
		"dev-plane-x-launchableId":  "launchable-id",
		"dev-plane-x-cloudCredId":   "cloud-cred-id",
		"dev-plane-stage":           "dev",
		"team":                      "gpu-workers",
	}
	labels := makeLabels(refID, tags)
	labels = append(labels, managedKeyIDLabelPrefix+"7")
	assert.Contains(t, labels, refIDLabelPrefix+refID)
	assert.Contains(t, labels, "brev-tag-dev-plane-managedby_dev-plane")
	assert.Contains(t, labels, "brev-tag-dev-plane-x-instanceid_instance-id")
	assert.NotContains(t, labels, tagLabelPrefix+"team=gpu-workers")

	parsedRefID, parsedTags := parseLabels(&labels)
	assert.Equal(t, refID, parsedRefID)
	for _, key := range instanceTagLabelKeys {
		assert.Equal(t, tags[key], parsedTags[key])
	}
	assert.NotContains(t, parsedTags, "team")
}

func TestParseLabelsOnlyDecodesKnownTagPrefixes(t *testing.T) {
	labels := []string{
		"brev-tag-team_gpu-workers",
		"brev-tag-dev-plane-x-instanceid_instance-123",
	}

	_, tags := parseLabels(&labels)

	assert.Equal(t, "instance-123", tags["dev-plane-x-instanceId"])
	assert.Equal(t, "", tags["brev-tag-team_gpu-workers"])
	assert.NotContains(t, tags, "team")
}

func TestCallerKeyPairIsNotManaged(t *testing.T) {
	keyName := "customer-key"
	keyPair, err := (&HyperstackClient{}).resolveKeyPair(context.Background(), v1.CreateInstanceAttrs{
		KeyPairName: &keyName,
	}, "default-CANADA-1")
	require.NoError(t, err)
	assert.Equal(t, keyName, keyPair.name)
	assert.Zero(t, keyPair.managedID)
}

func TestActiveInstanceWaitsForFloatingIP(t *testing.T) {
	status := "ACTIVE"
	client := &HyperstackClient{}

	withoutIP := client.convertInstance(virtualmachine.InstanceFields{Status: &status}, false)
	assert.Equal(t, v1.LifecycleStatusPending, withoutIP.Status.LifecycleStatus)

	publicIP := "203.0.113.42"
	attaching := "ATTACHING"
	withAttachingIP := client.convertInstance(virtualmachine.InstanceFields{
		Status: &status, FloatingIp: &publicIP, FloatingIpStatus: &attaching,
	}, false)
	assert.Equal(t, v1.LifecycleStatusPending, withAttachingIP.Status.LifecycleStatus)
}

func TestActiveInstanceWaitsForGuestBoot(t *testing.T) {
	client := &HyperstackClient{}
	status := "ACTIVE"
	publicIP := "203.0.113.42"
	attached := "ATTACHED"
	running := "RUNNING"
	providerInstance := virtualmachine.InstanceFields{
		Status: &status, FloatingIp: &publicIP, FloatingIpStatus: &attached,
		VmState: &status, PowerState: &running,
	}

	booting := client.convertInstance(providerInstance, false)
	assert.Equal(t, v1.LifecycleStatusPending, booting.Status.LifecycleStatus)

	ready := client.convertInstance(providerInstance, true)
	assert.Equal(t, v1.LifecycleStatusRunning, ready.Status.LifecycleStatus)
}

func TestConsoleReadyPollsAsyncLogRequest(t *testing.T) {
	requestCalls := 0
	getCalls := 0
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		switch request.Method {
		case http.MethodPost:
			requestCalls++
			writer.Header().Set("Content-Type", "application/json")
			writer.WriteHeader(http.StatusAccepted)
			require.NoError(t, json.NewEncoder(writer).Encode(map[string]any{"request_id": 17}))
		case http.MethodGet:
			getCalls++
			assert.Equal(t, "17", request.URL.Query().Get("request_id"))
			if getCalls == 1 {
				writer.WriteHeader(http.StatusBadRequest)
				return
			}
			writeJSON(t, writer, map[string]any{"logs": readinessMarker})
		default:
			http.NotFound(writer, request)
		}
	}))
	defer server.Close()

	client := newTestClient(t, server.URL)
	ready, err := client.vmOperatingSystemReportsReady(context.Background(), 42)
	require.NoError(t, err)
	assert.True(t, ready)
	assert.Equal(t, 1, requestCalls)
	assert.Equal(t, 2, getCalls)
}

func TestConsoleReadyRetainsRequestWhileLogsAreProcessing(t *testing.T) {
	requestCalls := 0
	getCalls := 0
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		switch request.Method {
		case http.MethodPost:
			requestCalls++
			writeJSON(t, writer, map[string]any{"request_id": 17})
		case http.MethodGet:
			getCalls++
			assert.Equal(t, "17", request.URL.Query().Get("request_id"))
			if getCalls == 1 {
				writeJSON(t, writer, map[string]any{
					"status": true, "message": "request is still processing",
				})
				return
			}
			writeJSON(t, writer, map[string]any{"logs": readinessMarker})
		default:
			http.NotFound(writer, request)
		}
	}))
	defer server.Close()

	client := newTestClient(t, server.URL)
	ready, err := client.vmOperatingSystemReportsReady(context.Background(), 42)
	require.NoError(t, err)
	assert.True(t, ready)
	assert.Equal(t, 1, requestCalls)
	assert.Equal(t, 2, getCalls)
}

func assertSecurityRule(t *testing.T, rule virtualmachine.CreateSecurityRulePayload, cidr string, fromPort, toPort int) {
	t.Helper()
	assert.Equal(t, "ingress", rule.Direction)
	assert.Equal(t, "IPv4", rule.Ethertype)
	assert.Equal(t, virtualmachine.Tcp, rule.Protocol)
	assert.Equal(t, cidr, rule.RemoteIpPrefix)
	require.NotNil(t, rule.PortRangeMin)
	require.NotNil(t, rule.PortRangeMax)
	assert.Equal(t, fromPort, *rule.PortRangeMin)
	assert.Equal(t, toPort, *rule.PortRangeMax)
}

func TestMakeSecurityRules(t *testing.T) {
	rules, err := makeSecurityRules(v1.FirewallRules{IngressRules: []v1.FirewallRule{{
		FromPort: defaultSSHPort,
		ToPort:   defaultSSHPort,
		IPRanges: []string{"52.9.0.116/32"},
	}}})
	require.NoError(t, err)
	require.Len(t, rules, 1)
	assertSecurityRule(t, rules[0], "52.9.0.116/32", defaultSSHPort, defaultSSHPort)
}

func TestValidateCreateInstanceAttrs(t *testing.T) {
	valid := v1.CreateInstanceAttrs{
		RefID:        "ref",
		InstanceType: "n3-H100x1",
		PublicKey:    testSSHPublicKey,
	}
	require.NoError(t, validateCreateInstanceAttrs(valid, "CANADA-1"))

	spot := valid
	spot.InstanceType = "n3-H100x1-spot"
	require.Error(t, validateCreateInstanceAttrs(spot, "CANADA-1"))
	spot.UseSpot = true
	require.NoError(t, validateCreateInstanceAttrs(spot, "CANADA-1"))
}
