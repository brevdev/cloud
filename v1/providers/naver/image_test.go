package v1

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	cloud "github.com/brevdev/cloud/v1"
)

func TestGetImagesConvertsAndFiltersImageProducts(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/vserver/v2/getServerImageProductList" {
			t.Fatalf("path = %q", r.URL.Path)
		}
		if got := r.URL.Query().Get("platformTypeCodeList.1"); got != "LNX64" {
			t.Fatalf("platform filter = %q", got)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"getServerImageProductListResponse":{"returnCode":"0","returnMessage":"success","totalRows":2,"productList":[{"productCode":"UBUNTU24","productName":"ubuntu-24.04","productDescription":"Ubuntu Server 24.04","platformType":{"code":"LNX64","codeName":"Linux 64 Bit"},"osInformation":"Ubuntu 24.04 (64-bit)"},{"productCode":"WIN","productName":"windows","productDescription":"Windows","platformType":{"code":"WND64","codeName":"Windows 64 Bit"},"osInformation":"Windows"}]}}`))
	}))
	defer server.Close()

	client := newTestNaverClient(t, server.URL)
	images, err := client.GetImages(context.Background(), cloud.GetImageArgs{
		Architectures: []string{"x86_64"},
		ImageIDs:      []string{"UBUNTU24"},
	})
	if err != nil {
		t.Fatalf("GetImages() error = %v", err)
	}
	if len(images) != 1 {
		t.Fatalf("images len = %d, want 1", len(images))
	}
	if images[0].ID != "UBUNTU24" || images[0].Architecture != "x86_64" || images[0].Name != "ubuntu-24.04" {
		t.Fatalf("unexpected image: %+v", images[0])
	}
}
