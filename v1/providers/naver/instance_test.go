package v1

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	cloud "github.com/brevdev/cloud/v1"
)

func TestCreateInstanceImportsKeyAndCreatesServer(t *testing.T) {
	var sawImport bool
	var sawCreate bool
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch r.URL.Path {
		case "/vserver/v2/importLoginKey":
			sawImport = true
			if got := r.URL.Query().Get("publicKey"); got != "ssh-rsa AAAA test" {
				t.Fatalf("publicKey = %q", got)
			}
			_, _ = w.Write([]byte(`{"importLoginKeyResponse":{"returnCode":"0","returnMessage":"success","loginKeyList":[{"keyName":"brev-ref-1","fingerprint":"fp"}]}}`))
		case "/vserver/v2/createServerInstances":
			sawCreate = true
			query := r.URL.Query()
			assertQuery(t, query.Get("regionCode"), "KR")
			assertQuery(t, query.Get("vpcNo"), "123")
			assertQuery(t, query.Get("subnetNo"), "456")
			assertQuery(t, query.Get("serverImageProductCode"), "UBUNTU24")
			assertQuery(t, query.Get("serverProductCode"), "GPU-T4")
			assertQuery(t, query.Get("loginKeyName"), "brev-ref-1")
			assertQuery(t, query.Get("networkInterfaceList.1.networkInterfaceOrder"), "0")
			_, _ = w.Write([]byte(`{"createServerInstancesResponse":{"returnCode":"0","returnMessage":"success","serverInstanceList":[{"serverInstanceNo":"999","serverName":"brev-test","cpuCount":4,"memorySize":17179869184,"publicIp":"203.0.113.10","serverInstanceStatus":{"code":"INIT","codeName":"Server init state"},"serverInstanceStatusName":"init","createDate":"2025-06-11T17:00:14+0900","serverImageProductCode":"UBUNTU24","serverProductCode":"GPU-T4","zoneCode":"KR-1","regionCode":"KR","vpcNo":"123","subnetNo":"456","baseBlockStorageSize":53687091200,"baseBlockStorageDiskDetailType":{"code":"SSD","codeName":"SSD"}}]}}`))
		default:
			t.Fatalf("unexpected path %q", r.URL.Path)
		}
	}))
	defer server.Close()

	client := newTestNaverClient(t, server.URL)
	instance, err := client.CreateInstance(context.Background(), cloud.CreateInstanceAttrs{
		Location:     "KR",
		Name:         "brev-test",
		RefID:        "ref-1",
		VPCID:        "123",
		SubnetID:     "456",
		PublicKey:    "ssh-rsa AAAA test",
		ImageID:      "UBUNTU24",
		InstanceType: "GPU-T4",
	})
	if err != nil {
		t.Fatalf("CreateInstance() error = %v", err)
	}
	if !sawImport || !sawCreate {
		t.Fatalf("sawImport=%v sawCreate=%v", sawImport, sawCreate)
	}
	if instance.CloudID != "999" || instance.Status.LifecycleStatus != cloud.LifecycleStatusPending || instance.PublicIP != "203.0.113.10" {
		t.Fatalf("unexpected instance: %+v", instance)
	}
	if instance.SSHUser != defaultSSHUser || instance.SSHPort != defaultSSHPort {
		t.Fatalf("unexpected SSH settings: %+v", instance)
	}
}

func TestListGetAndTerminateInstances(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch r.URL.Path {
		case "/vserver/v2/getServerInstanceDetail":
			if got := r.URL.Query().Get("serverInstanceNo"); got != "999" {
				t.Fatalf("serverInstanceNo = %q", got)
			}
			_, _ = w.Write([]byte(`{"getServerInstanceDetailResponse":{"returnCode":"0","returnMessage":"success","totalRows":1,"serverInstanceList":[{"serverInstanceNo":"999","serverName":"brev-test","cpuCount":4,"memorySize":17179869184,"publicIp":"203.0.113.10","serverInstanceStatus":{"code":"RUN","codeName":"Server RUN status"},"serverInstanceStatusName":"running","createDate":"2025-06-11T17:00:14+0900","serverImageProductCode":"UBUNTU24","serverProductCode":"GPU-T4","zoneCode":"KR-1","regionCode":"KR","vpcNo":"123","subnetNo":"456","baseBlockStorageSize":53687091200}]}}`))
		case "/vserver/v2/getServerInstanceList":
			if got := r.URL.Query().Get("serverInstanceNoList.1"); got != "999" {
				t.Fatalf("serverInstanceNoList.1 = %q", got)
			}
			_, _ = w.Write([]byte(`{"getServerInstanceListResponse":{"returnCode":"0","returnMessage":"success","totalRows":1,"serverInstanceList":[{"serverInstanceNo":"999","serverName":"brev-test","cpuCount":4,"memorySize":17179869184,"publicIp":"203.0.113.10","serverInstanceStatus":{"code":"RUN","codeName":"Server RUN status"},"serverInstanceStatusName":"running","createDate":"2025-06-11T17:00:14+0900","serverImageProductCode":"UBUNTU24","serverProductCode":"GPU-T4","zoneCode":"KR-1","regionCode":"KR","vpcNo":"123","subnetNo":"456","baseBlockStorageSize":53687091200}]}}`))
		case "/vserver/v2/terminateServerInstances":
			if got := r.URL.Query().Get("serverInstanceNoList.1"); got != "999" {
				t.Fatalf("terminate serverInstanceNoList.1 = %q", got)
			}
			_, _ = w.Write([]byte(`{"terminateServerInstancesResponse":{"returnCode":"0","returnMessage":"success"}}`))
		default:
			t.Fatalf("unexpected path %q", r.URL.Path)
		}
	}))
	defer server.Close()

	client := newTestNaverClient(t, server.URL)
	got, err := client.GetInstance(context.Background(), "999")
	if err != nil {
		t.Fatalf("GetInstance() error = %v", err)
	}
	if got.CloudID != "999" || got.Status.LifecycleStatus != cloud.LifecycleStatusRunning {
		t.Fatalf("unexpected get instance: %+v", got)
	}

	list, err := client.ListInstances(context.Background(), cloud.ListInstancesArgs{
		InstanceIDs: []cloud.CloudProviderInstanceID{"999"},
	})
	if err != nil {
		t.Fatalf("ListInstances() error = %v", err)
	}
	if len(list) != 1 || list[0].CloudID != "999" {
		t.Fatalf("unexpected list: %+v", list)
	}

	if err := client.TerminateInstance(context.Background(), "999"); err != nil {
		t.Fatalf("TerminateInstance() error = %v", err)
	}
}

func assertQuery(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Fatalf("query = %q, want %q", got, want)
	}
}
