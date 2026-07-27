package v2

import (
	"testing"

	"github.com/stretchr/testify/require"

	cloudv1 "github.com/brevdev/cloud/v1"
)

func TestSFComputeV2InstanceTypeArchitecture(t *testing.T) {
	t.Parallel()

	instanceType := buildInstanceType(h100InstanceTypeMetadata, true)
	require.Equal(
		t,
		[]cloudv1.Architecture{cloudv1.ArchitectureX86_64},
		instanceType.SupportedArchitectures,
	)
	require.False(
		t,
		cloudv1.IsSelectedByArgs(
			instanceType,
			cloudv1.GetInstanceTypeArgs{
				ArchitectureFilter: &cloudv1.ArchitectureFilter{
					IncludeArchitectures: []cloudv1.Architecture{
						cloudv1.ArchitectureARM64,
					},
				},
			},
		),
	)
}
