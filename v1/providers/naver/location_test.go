package v1

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	cloud "github.com/brevdev/cloud/v1"
)

func TestGetLocationsConvertsRegions(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/vserver/v2/getRegionList" {
			t.Fatalf("path = %q", r.URL.Path)
		}
		if got := r.URL.Query().Get("responseFormatType"); got != "json" {
			t.Fatalf("responseFormatType = %q", got)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"getRegionListResponse":{"returnCode":"0","returnMessage":"success","totalRows":2,"regionList":[{"regionCode":"KR","regionName":"Korea"},{"regionCode":"JPN","regionName":"Japan"}]}}`))
	}))
	defer server.Close()

	client := newTestNaverClient(t, server.URL)
	locations, err := client.GetLocations(context.Background(), cloud.GetLocationsArgs{})
	if err != nil {
		t.Fatalf("GetLocations() error = %v", err)
	}
	if len(locations) != 2 {
		t.Fatalf("locations len = %d, want 2", len(locations))
	}
	if locations[0].Name != "KR" || locations[0].Description != "Korea" || !locations[0].Available || locations[0].Country != "KOR" {
		t.Fatalf("unexpected KR location: %+v", locations[0])
	}
	if locations[1].Name != "JPN" || locations[1].Country != "JPN" {
		t.Fatalf("unexpected JPN location: %+v", locations[1])
	}
}

func newTestNaverClient(t *testing.T, baseURL string) *NaverClient {
	t.Helper()
	client, err := NewNaverClient("ref-1", "access", "secret", WithBaseURL(baseURL), WithClock(func() time.Time {
		return time.UnixMilli(1700000000123)
	}))
	if err != nil {
		t.Fatalf("NewNaverClient() error = %v", err)
	}
	return client
}
