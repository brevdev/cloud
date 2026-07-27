package v1

import (
	"testing"

	"github.com/stretchr/testify/require"

	cloudv1 "github.com/brevdev/cloud/v1"
)

func TestNebiusPlatformArchitecture(t *testing.T) {
	t.Parallel()

	tests := []struct {
		platform string
		want     cloudv1.Architecture
	}{
		{platform: "gpu-b300-sxm", want: cloudv1.ArchitectureX86_64},
		{platform: "gpu-b200-sxm", want: cloudv1.ArchitectureX86_64},
		{platform: "gpu-b200-sxm-a", want: cloudv1.ArchitectureX86_64},
		{platform: "gpu-h200-sxm", want: cloudv1.ArchitectureX86_64},
		{platform: "gpu-h100-sxm", want: cloudv1.ArchitectureX86_64},
		{platform: "gpu-rtx6000", want: cloudv1.ArchitectureX86_64},
		{platform: "gpu-l40s-a", want: cloudv1.ArchitectureX86_64},
		{platform: "gpu-l40s-d", want: cloudv1.ArchitectureX86_64},
		{platform: "cpu-d3", want: cloudv1.ArchitectureX86_64},
		{platform: "cpu-e2", want: cloudv1.ArchitectureX86_64},
		{platform: "future-platform", want: cloudv1.ArchitectureUnknown},
	}

	for _, tt := range tests {
		t.Run(tt.platform, func(t *testing.T) {
			t.Parallel()
			require.Equal(t, tt.want, nebiusPlatformArchitecture(tt.platform))
		})
	}
}

func TestApplyInstanceTypeFiltersUsesArchitectureMetadata(t *testing.T) {
	t.Parallel()

	x86Type := cloudv1.InstanceType{
		Type:                   "x86-type",
		SupportedArchitectures: []cloudv1.Architecture{cloudv1.ArchitectureX86_64},
	}
	armType := cloudv1.InstanceType{
		Type:                   "arm-type",
		SupportedArchitectures: []cloudv1.Architecture{cloudv1.ArchitectureARM64},
	}
	unknownType := cloudv1.InstanceType{
		Type:                   "unknown-type",
		SupportedArchitectures: []cloudv1.Architecture{cloudv1.ArchitectureUnknown},
	}
	all := []cloudv1.InstanceType{x86Type, armType, unknownType}

	tests := []struct {
		name string
		args cloudv1.GetInstanceTypeArgs
		want []cloudv1.InstanceType
	}{
		{name: "unfiltered", want: all},
		{
			name: "include x86",
			args: cloudv1.GetInstanceTypeArgs{
				ArchitectureFilter: &cloudv1.ArchitectureFilter{
					IncludeArchitectures: []cloudv1.Architecture{
						cloudv1.ArchitectureX86_64,
					},
				},
			},
			want: []cloudv1.InstanceType{x86Type},
		},
		{
			name: "include ARM",
			args: cloudv1.GetInstanceTypeArgs{
				ArchitectureFilter: &cloudv1.ArchitectureFilter{
					IncludeArchitectures: []cloudv1.Architecture{
						cloudv1.ArchitectureARM64,
					},
				},
			},
			want: []cloudv1.InstanceType{armType},
		},
		{
			name: "exclude ARM",
			args: cloudv1.GetInstanceTypeArgs{
				ArchitectureFilter: &cloudv1.ArchitectureFilter{
					ExcludeArchitectures: []cloudv1.Architecture{
						cloudv1.ArchitectureARM64,
					},
				},
			},
			want: []cloudv1.InstanceType{x86Type, unknownType},
		},
		{
			name: "exclude x86",
			args: cloudv1.GetInstanceTypeArgs{
				ArchitectureFilter: &cloudv1.ArchitectureFilter{
					ExcludeArchitectures: []cloudv1.Architecture{
						cloudv1.ArchitectureX86_64,
					},
				},
			},
			want: []cloudv1.InstanceType{armType, unknownType},
		},
	}

	client := &NebiusClient{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			require.Equal(t, tt.want, client.applyInstanceTypeFilters(all, tt.args))
		})
	}
}
