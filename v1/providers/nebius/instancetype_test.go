package v1

import (
	"context"
	"testing"

	"github.com/alecthomas/units"
	"github.com/bojanz/currency"
	common "github.com/nebius/gosdk/proto/nebius/common/v1"
	compute "github.com/nebius/gosdk/proto/nebius/compute/v1"
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

func TestGetInstanceTypesForLocationConstructsMappedPlatforms(t *testing.T) {
	t.Parallel()

	client := &NebiusClient{
		projectID: "test-project",
		logger:    &cloudv1.NoopLogger{},
		getInstanceTypePrice: func(
			context.Context,
			string,
			string,
			string,
		) *currency.Amount {
			return nil
		},
	}
	platforms := &compute.ListPlatformsResponse{
		Items: []*compute.Platform{
			{
				Metadata: &common.ResourceMetadata{Name: "gpu-b300-sxm"},
				Spec: &compute.PlatformSpec{
					GpuMemoryGigabytes: 270,
					Presets: []*compute.Preset{
						{
							Name: "1gpu-24vcpu-346gb",
							Resources: &compute.PresetResources{
								GpuCount:        1,
								VcpuCount:       24,
								MemoryGibibytes: 346,
							},
						},
					},
				},
			},
			{
				Metadata: &common.ResourceMetadata{Name: "gpu-rtx6000"},
				Spec: &compute.PlatformSpec{
					GpuMemoryGigabytes: 96,
					Presets: []*compute.Preset{
						{
							Name: "1gpu-24vcpu-218gb",
							Resources: &compute.PresetResources{
								GpuCount:        1,
								VcpuCount:       24,
								MemoryGibibytes: 218,
							},
						},
					},
				},
			},
		},
	}

	instanceTypes, err := client.getInstanceTypesForLocation(
		context.Background(),
		platforms,
		cloudv1.Location{Name: "test-region"},
		cloudv1.GetInstanceTypeArgs{},
		nil,
		nil,
	)
	require.NoError(t, err)
	require.Len(t, instanceTypes, 2)

	tests := []struct {
		instanceType string
		gpuType      string
		memoryGB     int64
	}{
		{
			instanceType: "gpu-b300-sxm.1gpu-24vcpu-346gb",
			gpuType:      "B300",
			memoryGB:     270,
		},
		{
			instanceType: "gpu-rtx6000.1gpu-24vcpu-218gb",
			gpuType:      "RTX6000",
			memoryGB:     96,
		},
	}
	for i, tt := range tests {
		instanceType := instanceTypes[i]
		require.Equal(t, tt.instanceType, instanceType.Type)
		require.Equal(
			t,
			[]cloudv1.Architecture{cloudv1.ArchitectureX86_64},
			instanceType.SupportedArchitectures,
		)
		require.False(t, instanceType.IsAvailable)
		require.Len(t, instanceType.SupportedGPUs, 1)

		gpu := instanceType.SupportedGPUs[0]
		require.Equal(t, int32(1), gpu.Count)
		require.Equal(t, tt.gpuType, gpu.Type)
		require.Equal(t, tt.gpuType, gpu.Name)
		require.Equal(t, cloudv1.ManufacturerNVIDIA, gpu.Manufacturer)
		require.Equal(
			t,
			units.Base2Bytes(tt.memoryGB*int64(units.Gibibyte)),
			gpu.Memory,
		)
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
