package hyperstack

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	virtualmachine "github.com/NexGenCloud/hyperstack-sdk-go/lib/virtual_machine"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	v1 "github.com/brevdev/cloud/v1"
)

const (
	testSSHPublicKey  = "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIDBdptDTzJ2cOmdyryG1B7yb1YssiCQs6SWu4HlbZXGE"
	statusActive      = "ACTIVE"
	statusShuttingOff = "SHUTOFF"
	statusAttached    = "ATTACHED"
	statusAttaching   = "ATTACHING"
)

func TestCreateInstance(t *testing.T) {
	var payload virtualmachine.CreateInstancesPayload
	id := 42
	keyPairID := 7
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		switch {
		case request.URL.Path == "/v1/core/environments" && request.Method == http.MethodGet:
			writeJSON(t, writer, map[string]any{"status": true, "environments": []map[string]any{{
				"id": 5, "name": "default-CANADA-1", "region": "CANADA-1",
			}}})
		case request.URL.Path == "/v1/core/keypairs" && request.Method == http.MethodGet:
			writeJSON(t, writer, map[string]any{"status": true, "keypairs": []any{}})
		case request.URL.Path == "/v1/core/keypairs" && request.Method == http.MethodPost:
			writeJSON(t, writer, map[string]any{"status": true, "keypair": map[string]any{"id": keyPairID, "name": "ref-123"}})
		case request.URL.Path == "/v1/core/virtual-machines" && request.Method == http.MethodPost:
			require.NoError(t, json.NewDecoder(request.Body).Decode(&payload))
			writeJSON(t, writer, map[string]any{"status": true, "instances": []map[string]any{{"id": 42}}})
		case request.URL.Path == fmt.Sprintf("/v1/core/virtual-machines/%d", id) && request.Method == http.MethodGet:
			labels := append(*payload.Labels, readinessLabel)
			writeJSON(t, writer, map[string]any{
				"status": true,
				"instance": map[string]any{
					"id":          id,
					"name":        "ref-123",
					"status":      statusActive,
					"floating_ip": "203.0.113.42",
					"environment": map[string]any{"name": "default-CANADA-1", "region": "CANADA-1"},
					"image":       map[string]any{"name": defaultImageName},
					"flavor":      map[string]any{"name": "n3-H100x1"},
					"labels":      labels,
				},
			})
		default:
			http.NotFound(writer, request)
		}
	}))
	defer server.Close()

	client := newTestClient(t, server.URL+"/v1")
	instance, err := client.CreateInstance(context.Background(), v1.CreateInstanceAttrs{
		Location:     "CANADA-1",
		RefID:        "ref-123",
		PublicKey:    testSSHPublicKey,
		InstanceType: "n3-H100x1",
		Tags:         v1.Tags{"dev-plane-x-instanceId": "instance-123"},
	})
	require.NoError(t, err)

	assert.Equal(t, "ref-123", payload.Name)
	assert.Equal(t, "default-CANADA-1", payload.EnvironmentName)
	assert.Equal(t, "n3-H100x1", payload.FlavorName)
	require.NotNil(t, payload.AssignFloatingIp)
	assert.True(t, *payload.AssignFloatingIp)
	require.NotNil(t, payload.UserData)
	assert.Contains(t, *payload.UserData, readinessMarker)

	assert.Equal(t, v1.CloudProviderInstanceID("42"), instance.CloudID)
	assert.Equal(t, "ref-123", instance.RefID)
	assert.Equal(t, "instance-123", instance.Tags["dev-plane-x-instanceId"])
	assert.Equal(t, v1.LifecycleStatusRunning, instance.Status.LifecycleStatus)
	assert.Equal(t, "203.0.113.42", instance.PublicIP)
}

func TestStartAndStopInstance(t *testing.T) {
	status := statusActive
	startCalls := 0
	stopCalls := 0
	id := 42
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		switch request.URL.Path {
		case fmt.Sprintf("/v1/core/virtual-machines/%d", id):
			writeJSON(t, writer, map[string]any{"status": true, "instance": map[string]any{"id": id, "status": status}})
		case fmt.Sprintf("/v1/core/virtual-machines/%d/stop", id):
			stopCalls++
			status = statusShuttingOff
			writeJSON(t, writer, map[string]any{"status": true})
		case fmt.Sprintf("/v1/core/virtual-machines/%d/start", id):
			startCalls++
			status = statusActive
			writeJSON(t, writer, map[string]any{"status": true})
		default:
			http.NotFound(writer, request)
		}
	}))
	defer server.Close()

	client := newTestClient(t, server.URL+"/v1")
	require.NoError(t, client.StopInstance(context.Background(), v1.CloudProviderInstanceID(strconv.Itoa(id))))
	require.NoError(t, client.StartInstance(context.Background(), v1.CloudProviderInstanceID(strconv.Itoa(id))))
	require.NoError(t, client.StartInstance(context.Background(), v1.CloudProviderInstanceID(strconv.Itoa(id))))
	assert.Equal(t, 1, stopCalls)
	assert.Equal(t, 1, startCalls)
}

func TestTerminateInstanceDeletesManagedKeyPair(t *testing.T) {
	vmDeletes := 0
	keyPairDeletes := 0
	id := 42
	keyPairID := 7
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		switch {
		case request.URL.Path == fmt.Sprintf("/v1/core/virtual-machines/%d", id) && request.Method == http.MethodGet:
			writeJSON(t, writer, map[string]any{"status": true, "instance": map[string]any{
				"id": id, "labels": []string{managedKeyIDLabelPrefix + strconv.Itoa(keyPairID)},
			}})
		case request.URL.Path == fmt.Sprintf("/v1/core/virtual-machines/%d", id) && request.Method == http.MethodDelete:
			vmDeletes++
			writeJSON(t, writer, map[string]any{"status": true})
		case request.URL.Path == fmt.Sprintf("/v1/core/keypair/%d", keyPairID) && request.Method == http.MethodDelete:
			keyPairDeletes++
			writeJSON(t, writer, map[string]any{"status": true})
		default:
			http.NotFound(writer, request)
		}
	}))
	defer server.Close()

	client := newTestClient(t, server.URL+"/v1")
	require.NoError(t, client.TerminateInstance(context.Background(), "42"))
	assert.Equal(t, 1, vmDeletes)
	assert.Equal(t, 1, keyPairDeletes)
}

func TestResolveKeyPairSearchesEveryPage(t *testing.T) {
	listCalls := 0
	importCalls := 0
	keyPairID := 7
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
				"id": keyPairID, "name": "ref", "public_key": testSSHPublicKey,
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
	assert.Equal(t, keyPairID, keyPair.managedID)
	assert.Equal(t, 2, listCalls)
	assert.Zero(t, importCalls)
}

func TestLabelsRoundTrip(t *testing.T) {
	const refID = "environment-123"
	keyPairID := 7
	tags := v1.Tags{
		"dev-plane-managedBy":    "dev-plane",
		"dev-plane-x-instanceId": "instance-123",
	}
	labels := makeLabels(refID, tags)
	labels = append(labels, managedKeyIDLabelPrefix+strconv.Itoa(keyPairID))

	parsedRefID, parsedTags := parseLabels(&labels)
	assert.Equal(t, refID, parsedRefID)
	assert.Equal(t, tags, parsedTags)
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
	status := statusActive
	client := &HyperstackClient{}

	withoutIP := client.convertInstance(virtualmachine.InstanceFields{Status: &status}, false)
	assert.Equal(t, v1.LifecycleStatusPending, withoutIP.Status.LifecycleStatus)

	publicIP := "203.0.113.42"
	attaching := statusAttaching
	withAttachingIP := client.convertInstance(virtualmachine.InstanceFields{
		Status: &status, FloatingIp: &publicIP, FloatingIpStatus: &attaching,
	}, false)
	assert.Equal(t, v1.LifecycleStatusPending, withAttachingIP.Status.LifecycleStatus)
}

func TestActiveInstanceWaitsForGuestBoot(t *testing.T) {
	client := &HyperstackClient{}
	status := statusActive
	publicIP := "203.0.113.42"
	attached := statusAttached
	running := statusActive
	providerInstance := virtualmachine.InstanceFields{
		Status: &status, FloatingIp: &publicIP, FloatingIpStatus: &attached,
		VmState: &status, PowerState: &running,
	}

	booting := client.convertInstance(providerInstance, false)
	assert.Equal(t, v1.LifecycleStatusPending, booting.Status.LifecycleStatus)

	ready := client.convertInstance(providerInstance, true)
	assert.Equal(t, v1.LifecycleStatusRunning, ready.Status.LifecycleStatus)
}

func TestConsoleReadyPollsUntilLogsAreAvailable(t *testing.T) {
	requestCalls := 0
	getCalls := 0
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		switch request.Method {
		case http.MethodPost:
			requestCalls++
			writer.WriteHeader(http.StatusAccepted)
			require.NoError(t, json.NewEncoder(writer).Encode(map[string]any{"request_id": 17}))
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
	id := 42
	ready, err := client.vmOperatingSystemReportsReady(context.Background(), virtualmachine.InstanceFields{Id: &id})
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
