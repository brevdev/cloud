package v1

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"k8s.io/client-go/rest"
)

func TestCredentialRestConfigRequiresKubeconfigByDefault(t *testing.T) {
	t.Parallel()

	credential := NewTestKubeCredential("test", "", "testkube")

	_, err := credential.restConfig()
	require.EqualError(t, err, "kubeconfigBase64 is required")
}

func TestCredentialRestConfigUsesInClusterAuthMode(t *testing.T) {
	original := restInClusterConfig
	t.Cleanup(func() {
		restInClusterConfig = original
	})

	called := false
	restInClusterConfig = func() (*rest.Config, error) {
		called = true
		return &rest.Config{Host: "https://kubernetes.default.svc"}, nil
	}

	credential := NewInClusterTestKubeCredential("test", "testkube")

	config, err := credential.restConfig()
	require.NoError(t, err)
	require.True(t, called)
	require.Equal(t, "https://kubernetes.default.svc", config.Host)
}

func TestCredentialRestConfigRejectsMixedAuthConfig(t *testing.T) {
	t.Parallel()

	credential := &TestKubeCredential{
		RefID:            "test",
		AuthMode:         TestKubeAuthModeInCluster,
		KubeconfigBase64: "not-allowed",
		Namespace:        "testkube",
	}

	_, err := credential.restConfig()
	require.EqualError(t, err, `kubeconfigBase64 must be empty when authMode is "in-cluster"`)
}

func TestCredentialRestConfigRejectsUnknownAuthMode(t *testing.T) {
	t.Parallel()

	credential := &TestKubeCredential{
		RefID:     "test",
		AuthMode:  "token",
		Namespace: "testkube",
	}

	_, err := credential.restConfig()
	require.EqualError(t, err, "unknown testkube auth mode: token")
}

func TestCredentialMakeClientUsesInClusterConfig(t *testing.T) {
	original := restInClusterConfig
	t.Cleanup(func() {
		restInClusterConfig = original
	})

	restInClusterConfig = func() (*rest.Config, error) {
		return &rest.Config{Host: "https://kubernetes.default.svc"}, nil
	}

	credential := NewInClusterTestKubeCredential("test", "testkube")

	client, err := credential.MakeClient(context.Background(), "staging")
	require.NoError(t, err)

	testKubeClient, ok := client.(*TestKubeClient)
	require.True(t, ok)
	require.Equal(t, "testkube", testKubeClient.namespace)
	require.Equal(t, "staging", testKubeClient.location)
}
