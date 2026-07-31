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
	require.Len(t, instanceTypes, 5)

	instanceTypeByName := map[string]cloudv1.InstanceType{}
	for _, instanceType := range instanceTypes {
		instanceTypeByName[instanceType.Type] = instanceType
	}

	for _, expected := range []string{
		InstanceTypeOKCPU,
		InstanceTypeFailCapacity,
		InstanceTypeFailQuota,
		InstanceTypeFailBuild,
		InstanceTypeOKCPUARM64,
	} {
		instanceType, ok := instanceTypeByName[expected]
		require.True(t, ok, "missing instance type %s", expected)
		require.True(t, instanceType.IsAvailable)
		require.Equal(t, CloudProviderID, instanceType.Provider)
		require.Equal(t, CloudProviderID, instanceType.Cloud)
		require.NotNil(t, instanceType.BasePrice)

		basePrice, err := instanceType.BasePrice.Int64()
		require.NoError(t, err)
		require.Greater(t, basePrice, int64(0))
	}
}

func TestGetInstanceTypesFiltersByArchitecture(t *testing.T) {
	client := newTestClient(t)

	tests := []struct {
		name         string
		architecture cloudv1.Architecture
		expected     []string
	}{
		{
			name:         "x86_64",
			architecture: cloudv1.ArchitectureX86_64,
			expected: []string{
				InstanceTypeOKCPU,
				InstanceTypeFailCapacity,
				InstanceTypeFailQuota,
				InstanceTypeFailBuild,
			},
		},
		{
			name:         "arm64",
			architecture: cloudv1.ArchitectureARM64,
			expected:     []string{InstanceTypeOKCPUARM64},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			instanceTypes, err := client.GetInstanceTypes(context.Background(), cloudv1.GetInstanceTypeArgs{
				ArchitectureFilter: &cloudv1.ArchitectureFilter{
					IncludeArchitectures: []cloudv1.Architecture{tt.architecture},
				},
			})
			require.NoError(t, err)

			actual := make([]string, 0, len(instanceTypes))
			for _, instanceType := range instanceTypes {
				actual = append(actual, instanceType.Type)
			}
			require.ElementsMatch(t, tt.expected, actual)
		})
	}
}

func TestGetInstanceTypesWithGPUManufacturerFilterIncludesCPU(t *testing.T) {
	client := newTestClient(t)

	instanceTypes, err := client.GetInstanceTypes(context.Background(), cloudv1.GetInstanceTypeArgs{
		GPUManufactererFilter: &cloudv1.GPUManufacturerFilter{
			IncludeGPUManufacturers: []cloudv1.Manufacturer{cloudv1.ManufacturerNVIDIA},
		},
	})
	require.NoError(t, err)
	require.Len(t, instanceTypes, 5)
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
