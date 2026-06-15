package v1

import (
	"context"
	"testing"

	cloudv1 "github.com/brevdev/cloud/v1"
	"github.com/stretchr/testify/require"
	"k8s.io/client-go/kubernetes/fake"
)

func TestGetInstanceTypes(t *testing.T) {
	client := newTestClient(t)

	instanceTypes, err := client.GetInstanceTypes(context.Background(), cloudv1.GetInstanceTypeArgs{})
	require.NoError(t, err)
	require.Len(t, instanceTypes, 4)

	instanceTypeByName := map[string]cloudv1.InstanceType{}
	for _, instanceType := range instanceTypes {
		instanceTypeByName[instanceType.Type] = instanceType
	}

	for _, expected := range []string{
		InstanceTypeOKCPU,
		InstanceTypeFailCapacity,
		InstanceTypeFailQuota,
		InstanceTypeFailBuild,
	} {
		instanceType, ok := instanceTypeByName[expected]
		require.True(t, ok, "missing instance type %s", expected)
		require.True(t, instanceType.IsAvailable)
		require.Equal(t, CloudProviderID, instanceType.Provider)
		require.Equal(t, CloudProviderID, instanceType.Cloud)
	}
}

func TestCapabilitiesDoNotAdvertiseImages(t *testing.T) {
	client := newTestClient(t)

	capabilities, err := client.GetCapabilities(context.Background())
	require.NoError(t, err)
	require.True(t, capabilities.IsCapable(cloudv1.CapabilityCreateInstance))
	require.False(t, capabilities.IsCapable(cloudv1.CapabilityMachineImage))
}

func newTestClient(t *testing.T) *TestKubeClient {
	t.Helper()

	client, err := NewTestKubeClient("test-credential", nil,
		WithKubernetesClient(fake.NewSimpleClientset()),
		WithNamespace("testkube"),
		WithLocation("local"),
	)
	require.NoError(t, err)
	return client
}
