package v2

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	v1 "github.com/brevdev/cloud/v1"
	"github.com/stretchr/testify/require"
)

const sshDetailsNotAvailableBody = `{"error":{"type":"not_found","message":"SSH details not currently available"}}`

func TestNormalizeTerminateInstanceErrorTreatsNotFoundAsAlreadyGone(t *testing.T) {
	t.Parallel()

	providerErr := fmt.Errorf("release instance: %w", &apiError{
		statusCode: http.StatusNotFound,
		body:       `{"error":"not found"}`,
	})

	err := normalizeTerminateInstanceError(providerErr)

	require.ErrorIs(t, err, v1.ErrInstanceNotFound)
	var responseErr *apiError
	require.False(t, errors.As(err, &responseErr))
}

func TestNormalizeTerminateInstanceErrorPreservesRetryableFailure(t *testing.T) {
	t.Parallel()

	providerErr := errors.New("provider rate limited")

	err := normalizeTerminateInstanceError(providerErr)

	require.ErrorIs(t, err, providerErr)
}

func TestMakeSFCNameIsDeterministic(t *testing.T) {
	t.Parallel()

	tags := v1.Tags{"dev-plane-x-environmentId": "p82qfn5qs"}

	first := makeSFCName("inst-2toqsvHXfalevkjPXY2QNJZL9HF", tags)
	second := makeSFCName("inst-2toqsvHXfalevkjPXY2QNJZL9HF", tags)

	require.Equal(t, first, second)
}

func TestMakeSFCNameIsUniquePerInstance(t *testing.T) {
	t.Parallel()

	// Same environment, two instances: re-provisioning must not reuse the previous name.
	tags := v1.Tags{"dev-plane-x-environmentId": "p82qfn5qs"}

	first := makeSFCName("inst-2toqsvHXfalevkjPXY2QNJZL9HF", tags)
	second := makeSFCName("inst-2gAgPPd3QC1nmSjaWg8kc7UfyDx", tags)

	require.NotEqual(t, first, second)
}

func TestIsSSHDetailsPending(t *testing.T) {
	t.Parallel()

	now := time.Unix(1_800_000_000, 0)
	notFound := fmt.Errorf("get ssh info: %w", &apiError{statusCode: http.StatusNotFound, body: sshDetailsNotAvailableBody})
	serverError := &apiError{statusCode: http.StatusInternalServerError, body: "internal error"}

	tests := []struct {
		name      string
		err       error
		createdAt time.Time
		want      bool
	}{
		{name: "404 for a new instance", err: notFound, createdAt: now.Add(-time.Minute), want: true},
		{name: "404 just before the window ends", err: notFound, createdAt: now.Add(-sshDetailsPendingWindow + time.Second), want: true},
		{name: "404 once the window has ended", err: notFound, createdAt: now.Add(-sshDetailsPendingWindow), want: false},
		{name: "server error for a new instance", err: serverError, createdAt: now.Add(-time.Minute), want: false},
		{name: "transport error for a new instance", err: errors.New("connection reset"), createdAt: now.Add(-time.Minute), want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			require.Equal(t, tt.want, isSSHDetailsPending(tt.err, tt.createdAt.Unix(), now))
		})
	}
}

func TestListInstancesKeepsNewInstanceWithoutSSHDetailsAsPending(t *testing.T) {
	t.Parallel()

	recently := time.Now().Add(-time.Minute).Unix()
	listed := []instanceResponse{
		{ID: "inst_starting", Status: instanceStatusRunning, CreatedAt: recently},
		{ID: "inst_ready", Status: instanceStatusRunning, CreatedAt: recently},
		{ID: "inst_old", Status: instanceStatusRunning, CreatedAt: time.Now().Add(-2 * sshDetailsPendingWindow).Unix()},
	}
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		switch request.URL.Path {
		case "/integrations/brev/v1/instances":
			writeJSON(t, writer, listInstancesResponse{Data: listed})
		case "/integrations/brev/v1/instances/inst_ready/ssh":
			writeJSON(t, writer, instanceSSHInfo{Hostname: "192.0.2.1", Port: 2222})
		default:
			http.Error(writer, sshDetailsNotAvailableBody, http.StatusNotFound)
		}
	}))
	t.Cleanup(server.Close)

	instances, err := newTestSFCClient(server.URL).ListInstances(context.Background(), v1.ListInstancesArgs{})
	require.NoError(t, err)

	byID := map[v1.CloudProviderInstanceID]v1.Instance{}
	for _, instance := range instances {
		byID[instance.CloudID] = instance
	}
	require.Equal(t, v1.LifecycleStatusPending, byID["inst_starting"].Status.LifecycleStatus)
	require.Empty(t, byID["inst_starting"].PublicIP)
	require.Equal(t, v1.LifecycleStatusRunning, byID["inst_ready"].Status.LifecycleStatus)
	require.Equal(t, "192.0.2.1", byID["inst_ready"].PublicIP)
	require.Equal(t, 2222, byID["inst_ready"].SSHPort)
	require.NotContains(t, byID, v1.CloudProviderInstanceID("inst_old"))
}

func TestGetInstanceReportsNewInstanceWithoutSSHDetailsAsPending(t *testing.T) {
	t.Parallel()

	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if request.URL.Path == "/integrations/brev/v1/instances/inst_starting" {
			writeJSON(t, writer, instanceResponse{ID: "inst_starting", Status: instanceStatusRunning, CreatedAt: time.Now().Add(-time.Minute).Unix()})
			return
		}
		http.Error(writer, sshDetailsNotAvailableBody, http.StatusNotFound)
	}))
	t.Cleanup(server.Close)

	instance, err := newTestSFCClient(server.URL).GetInstance(context.Background(), "inst_starting")
	require.NoError(t, err)
	require.Equal(t, v1.LifecycleStatusPending, instance.Status.LifecycleStatus)
	require.Empty(t, instance.PublicIP)
}

func newTestSFCClient(baseURL string) *SFCClientV2 {
	client := newAPIClient("api-key")
	client.baseURL = baseURL
	return &SFCClientV2{
		refID:        "cloud-cred",
		organization: "account",
		workspace:    "workspace",
		client:       client,
		logger:       &v1.NoopLogger{},
	}
}
