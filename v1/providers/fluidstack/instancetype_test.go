package v1

import (
	"testing"

	"github.com/stretchr/testify/require"

	cloudv1 "github.com/brevdev/cloud/v1"
	openapi "github.com/brevdev/cloud/v1/providers/fluidstack/gen/fluidstack"
)

func TestConvertFluidStackInstanceTypeArchitecture(t *testing.T) {
	t.Parallel()

	gh200 := "gh200"
	b200 := "B200"
	tests := []struct {
		name     string
		gpuModel *string
		want     cloudv1.Architecture
	}{
		{name: "GH200", gpuModel: &gh200, want: cloudv1.ArchitectureARM64},
		{name: "B200", gpuModel: &b200, want: cloudv1.ArchitectureX86_64},
		{name: "CPU", gpuModel: nil, want: cloudv1.ArchitectureX86_64},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			gpuCount := int32(1)
			if tt.gpuModel == nil {
				gpuCount = 0
			}
			got := convertFluidStackInstanceTypeToV1InstanceType(
				"",
				openapi.InstanceType{
					Name:     "test-" + tt.name,
					Cpu:      64,
					Memory:   "432GB",
					GpuModel: tt.gpuModel,
					GpuCount: &gpuCount,
				},
				true,
			)
			require.Equal(t, []cloudv1.Architecture{tt.want}, got.SupportedArchitectures)
		})
	}
}

func TestFilterFluidStackInstanceTypesByArchitecture(t *testing.T) {
	t.Parallel()

	x86Type := cloudv1.InstanceType{
		Type:                   "x86",
		SupportedArchitectures: []cloudv1.Architecture{cloudv1.ArchitectureX86_64},
	}
	armType := cloudv1.InstanceType{
		Type:                   "arm",
		SupportedArchitectures: []cloudv1.Architecture{cloudv1.ArchitectureARM64},
	}
	instanceTypes := []cloudv1.InstanceType{x86Type, armType}

	require.Equal(
		t,
		instanceTypes,
		filterInstanceTypesByArchitecture(instanceTypes, nil),
	)
	require.Equal(
		t,
		[]cloudv1.InstanceType{x86Type},
		filterInstanceTypesByArchitecture(
			instanceTypes,
			&cloudv1.ArchitectureFilter{
				ExcludeArchitectures: []cloudv1.Architecture{cloudv1.ArchitectureARM64},
			},
		),
	)
}
