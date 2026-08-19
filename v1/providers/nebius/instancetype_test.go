package v1

import (
	"testing"

	"github.com/stretchr/testify/require"

	cloudv1 "github.com/brevdev/cloud/v1"
)

func TestNebiusPlatformArchitecture(t *testing.T) {
	t.Parallel()

	require.Equal(t, cloudv1.ArchitectureX86_64, nebiusPlatformArchitecture(" GPU-H100-SXM "))
	require.Equal(t, cloudv1.ArchitectureX86_64, nebiusPlatformArchitecture("cpu-d3"))
	require.Equal(t, cloudv1.ArchitectureARM64, nebiusPlatformArchitecture("gpu-gb200"))
	require.Equal(t, cloudv1.ArchitectureARM64, nebiusPlatformArchitecture("gpu-gb300"))
	require.Equal(t, cloudv1.ArchitectureUnknown, nebiusPlatformArchitecture("future-platform"))
}

func TestNebiusGPUMemoryUsesNominalGB(t *testing.T) {
	t.Parallel()

	memory := getGPUMemory("L40S")

	require.Equal(t, cloudv1.NewBytes(48, cloudv1.Gigabyte), gpuMemoryBytes(memory))
}

func TestApplyInstanceTypeFiltersUsesArchitectureMetadata(t *testing.T) {
	t.Parallel()

	instanceType := func(name string, architecture cloudv1.Architecture) cloudv1.InstanceType {
		return cloudv1.InstanceType{
			Type:                   name,
			SupportedArchitectures: []cloudv1.Architecture{architecture},
		}
	}
	x86 := instanceType("arm-in-name", cloudv1.ArchitectureX86_64)
	arm := instanceType("x86-in-name", cloudv1.ArchitectureARM64)
	unknown := instanceType("unknown", cloudv1.ArchitectureUnknown)
	instanceTypes := []cloudv1.InstanceType{x86, arm, unknown}
	client := &NebiusClient{}
	apply := func(filter *cloudv1.ArchitectureFilter) []cloudv1.InstanceType {
		return client.applyInstanceTypeFilters(
			instanceTypes,
			cloudv1.GetInstanceTypeArgs{ArchitectureFilter: filter},
		)
	}

	require.Equal(t, []cloudv1.InstanceType{arm}, apply(
		&cloudv1.ArchitectureFilter{
			IncludeArchitectures: []cloudv1.Architecture{cloudv1.ArchitectureARM64},
		},
	))
	require.Equal(t, []cloudv1.InstanceType{x86, unknown}, apply(
		&cloudv1.ArchitectureFilter{
			ExcludeArchitectures: []cloudv1.Architecture{cloudv1.ArchitectureARM64},
		},
	))
}
