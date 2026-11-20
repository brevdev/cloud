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
		w.Header().Set("Content-Type", "application/json")
		switch r.URL.Path {
		case "/vserver/v2/getServerProductList":
			writeServerProductList(t, w, r)
		case "/billing/v1/product/getProductPriceList":
			writeProductPriceList(t, w, r)
		default:
			t.Fatalf("path = %q", r.URL.Path)
		}
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
	assertPricedGPUInstanceType(t, types[0])
}

func assertPricedGPUInstanceType(t *testing.T, it cloud.InstanceType) {
	t.Helper()
	if it.Provider != CloudProviderID || it.Cloud != CloudProviderID || it.Location != "KR" {
		t.Fatalf("unexpected provider fields: %+v", it)
	}
	if it.Type != "SVR.VSVR.GPU.T4.C004.M016.NET.SSD.B050.G002" || it.VCPU != 4 {
		t.Fatalf("unexpected instance type: %+v", it)
	}
	if len(it.SupportedGPUs) != 1 || it.SupportedGPUs[0].Name != "T4" || it.SupportedGPUs[0].Count != 1 {
		t.Fatalf("unexpected GPU conversion: %+v", it.SupportedGPUs)
	}
	if it.BasePrice == nil || it.BasePrice.CurrencyCode() != "USD" || it.BasePrice.Number() != "0.5" {
		t.Fatalf("unexpected base price: %v", it.BasePrice)
	}
	if it.ID == "" {
		t.Fatal("instance type ID is empty")
	}
}

func writeServerProductList(t *testing.T, w http.ResponseWriter, r *http.Request) {
	t.Helper()
	if got := r.URL.Query().Get("regionCode"); got != "KR" {
		t.Fatalf("regionCode = %q", got)
	}
	if got := r.URL.Query().Get("serverImageProductCode"); got != defaultServerImageProductCode {
		t.Fatalf("serverImageProductCode = %q", got)
	}
	_, _ = w.Write([]byte(`{"getServerProductListResponse":{"returnCode":"0","returnMessage":"success","totalRows":2,"productList":[{"productCode":"SVR.VSVR.GPU.T4.C004.M016.NET.SSD.B050.G002","productName":"vCPU 4EA, Memory 16GB, NVIDIA T4 1EA, [SSD]Disk 50GB","productType":{"code":"GPU","codeName":"GPU"},"productDescription":"vCPU 4EA, Memory 16GB, NVIDIA T4 1EA, [SSD]Disk 50GB","cpuCount":4,"memorySize":17179869184,"baseBlockStorageSize":53687091200,"diskType":{"code":"NET","codeName":"Network storage"},"generationCode":"G2"},{"productCode":"SVR.VSVR.STAND.C002.M008.NET.SSD.B050.G002","productName":"vCPU 2EA, Memory 8GB, [SSD]Disk 50GB","productType":{"code":"STAND","codeName":"Standard"},"productDescription":"vCPU 2EA, Memory 8GB, [SSD]Disk 50GB","cpuCount":2,"memorySize":8589934592,"baseBlockStorageSize":53687091200,"diskType":{"code":"NET","codeName":"Network storage"},"generationCode":"G2"}]}}`))
}

func writeProductPriceList(t *testing.T, w http.ResponseWriter, r *http.Request) {
	t.Helper()
	if got := r.URL.Query().Get("regionCode"); got != "KR" {
		t.Fatalf("price regionCode = %q", got)
	}
	if got := r.URL.Query().Get("productItemKindCode"); got != ncloudServerProductKindVPC {
		t.Fatalf("productItemKindCode = %q", got)
	}
	if got := r.URL.Query().Get("payCurrencyCode"); got != defaultPayCurrencyCode {
		t.Fatalf("payCurrencyCode = %q", got)
	}
	_, _ = w.Write([]byte(`{"getProductPriceListResponse":{"returnCode":"0","returnMessage":"success","totalRows":1,"productPriceList":[{"productCode":"SVR.VSVR.GPU.T4.C004.M016.NET.SSD.B050.G002","priceList":[{"priceType":{"code":"FXSUM","codeName":"Monthly flat rate"},"unit":{"code":"USAGE_TIME","codeName":"Usage time"},"price":1500,"payCurrency":{"code":"USD","codeName":"US Dollar"}},{"priceType":{"code":"MTRAT","codeName":"Meter rate"},"unit":{"code":"USAGE_HH","codeName":"Usage time (per hour)"},"price":0.5,"payCurrency":{"code":"USD","codeName":"US Dollar"}}]}]}}`))
}

func TestGetInstanceTypesIgnoresPriceLookupFailure(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch r.URL.Path {
		case "/vserver/v2/getServerProductList":
			_, _ = w.Write([]byte(`{"getServerProductListResponse":{"returnCode":"0","returnMessage":"success","totalRows":1,"productList":[{"productCode":"SVR.VSVR.STAND.C002.M008.NET.SSD.B050.G002","productName":"vCPU 2EA, Memory 8GB, [SSD]Disk 50GB","productType":{"code":"STAND","codeName":"Standard"},"productDescription":"vCPU 2EA, Memory 8GB, [SSD]Disk 50GB","cpuCount":2,"memorySize":8589934592,"baseBlockStorageSize":53687091200,"diskType":{"code":"NET","codeName":"Network storage"},"generationCode":"G2"}]}}`))
		case "/billing/v1/product/getProductPriceList":
			http.Error(w, "billing unavailable", http.StatusInternalServerError)
		default:
			t.Fatalf("path = %q", r.URL.Path)
		}
	}))
	defer server.Close()

	client := newTestNaverClient(t, server.URL)
	types, err := client.GetInstanceTypes(context.Background(), cloud.GetInstanceTypeArgs{
		Locations: cloud.LocationsFilter{"KR"},
	})
	if err != nil {
		t.Fatalf("GetInstanceTypes() error = %v", err)
	}
	if len(types) != 1 {
		t.Fatalf("instance types len = %d, want 1", len(types))
	}
	if types[0].BasePrice != nil {
		t.Fatalf("BasePrice = %v, want nil after price lookup failure", types[0].BasePrice)
	}
}
