package v1

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	cloud "github.com/brevdev/cloud/v1"
)

func TestGetInstanceTypesConvertsProducts(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/vserver/v2/getServerProductList" {
			t.Fatalf("path = %q", r.URL.Path)
		}
		if got := r.URL.Query().Get("regionCode"); got != "KR" {
			t.Fatalf("regionCode = %q", got)
		}
		if got := r.URL.Query().Get("serverImageProductCode"); got != defaultServerImageProductCode {
			t.Fatalf("serverImageProductCode = %q", got)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"getServerProductListResponse":{"returnCode":"0","returnMessage":"success","totalRows":2,"productList":[{"productCode":"SVR.VSVR.GPU.T4.C004.M016.NET.SSD.B050.G002","productName":"vCPU 4EA, Memory 16GB, NVIDIA T4 1EA, [SSD]Disk 50GB","productType":{"code":"GPU","codeName":"GPU"},"productDescription":"vCPU 4EA, Memory 16GB, NVIDIA T4 1EA, [SSD]Disk 50GB","cpuCount":4,"memorySize":17179869184,"baseBlockStorageSize":53687091200,"diskType":{"code":"NET","codeName":"Network storage"},"generationCode":"G2"},{"productCode":"SVR.VSVR.STAND.C002.M008.NET.SSD.B050.G002","productName":"vCPU 2EA, Memory 8GB, [SSD]Disk 50GB","productType":{"code":"STAND","codeName":"Standard"},"productDescription":"vCPU 2EA, Memory 8GB, [SSD]Disk 50GB","cpuCount":2,"memorySize":8589934592,"baseBlockStorageSize":53687091200,"diskType":{"code":"NET","codeName":"Network storage"},"generationCode":"G2"}]}}`))
	}))
	defer server.Close()

	client := newTestNaverClient(t, server.URL)
	types, err := client.GetInstanceTypes(context.Background(), cloud.GetInstanceTypeArgs{
		Locations:     cloud.LocationsFilter{"KR"},
		InstanceTypes: []string{"SVR.VSVR.GPU.T4.C004.M016.NET.SSD.B050.G002"},
	})
	if err != nil {
		t.Fatalf("GetInstanceTypes() error = %v", err)
	}
	if len(types) != 1 {
		t.Fatalf("instance types len = %d, want 1", len(types))
	}
	it := types[0]
	if it.Provider != CloudProviderID || it.Cloud != CloudProviderID || it.Location != "KR" {
		t.Fatalf("unexpected provider fields: %+v", it)
	}
	if it.Type != "SVR.VSVR.GPU.T4.C004.M016.NET.SSD.B050.G002" || it.VCPU != 4 {
		t.Fatalf("unexpected instance type: %+v", it)
	}
	if len(it.SupportedGPUs) != 1 || it.SupportedGPUs[0].Name != "T4" || it.SupportedGPUs[0].Count != 1 {
		t.Fatalf("unexpected GPU conversion: %+v", it.SupportedGPUs)
	}
	if it.ID == "" {
		t.Fatal("instance type ID is empty")
	}
}
