package v1

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	cloud "github.com/brevdev/cloud/v1"
)

func TestNaverClientSignsRequests(t *testing.T) {
	const (
		accessKey = "test-access"
		secretKey = "test-secret"
	)
	fixedNow := time.UnixMilli(1700000000123)
	var sawRequest bool

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sawRequest = true
		if got := r.Header.Get("x-ncp-apigw-timestamp"); got != "1700000000123" {
			t.Fatalf("timestamp header = %q, want fixed timestamp", got)
		}
		if got := r.Header.Get("x-ncp-iam-access-key"); got != accessKey {
			t.Fatalf("access key header = %q, want %q", got, accessKey)
		}

		expectedMessage := fmt.Sprintf("%s\n%s\n%s\n%s", r.Method, r.URL.RequestURI(), "1700000000123", accessKey)
		mac := hmac.New(sha256.New, []byte(secretKey))
		_, _ = mac.Write([]byte(expectedMessage))
		expectedSignature := base64.StdEncoding.EncodeToString(mac.Sum(nil))
		if got := r.Header.Get("x-ncp-apigw-signature-v2"); got != expectedSignature {
			t.Fatalf("signature header = %q, want %q", got, expectedSignature)
		}

		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"getRegionListResponse":{"returnCode":"0","returnMessage":"success","totalRows":1,"regionList":[{"regionCode":"KR","regionName":"Korea"}]}}`))
	}))
	defer server.Close()

	client, err := NewNaverClient("ref-1", accessKey, secretKey, WithBaseURL(server.URL), WithClock(func() time.Time {
		return fixedNow
	}))
	if err != nil {
		t.Fatalf("NewNaverClient() error = %v", err)
	}

	_, err = client.GetLocations(context.Background(), cloud.GetLocationsArgs{})
	if err != nil {
		t.Fatalf("GetLocations() error = %v", err)
	}
	if !sawRequest {
		t.Fatal("server did not receive request")
	}
}

func TestNaverCredentialCreatesClient(t *testing.T) {
	cred := NewNaverCredential("ref-1", "access", "secret")
	if got := cred.GetReferenceID(); got != "ref-1" {
		t.Fatalf("reference ID = %q", got)
	}
	if got := cred.GetCloudProviderID(); got != CloudProviderID {
		t.Fatalf("cloud provider ID = %q", got)
	}
	if got := cred.GetAPIType(); got != cloud.APITypeGlobal {
		t.Fatalf("API type = %q", got)
	}

	client, err := cred.MakeClient(context.Background(), "KR")
	if err != nil {
		t.Fatalf("MakeClient() error = %v", err)
	}
	if got := client.GetReferenceID(); got != "ref-1" {
		t.Fatalf("client reference ID = %q", got)
	}
}
