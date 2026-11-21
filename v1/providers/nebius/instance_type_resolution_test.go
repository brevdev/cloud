package v1

import (
	"testing"

	v1 "github.com/brevdev/cloud/v1"
	common "github.com/nebius/gosdk/proto/nebius/common/v1"
	compute "github.com/nebius/gosdk/proto/nebius/compute/v1"
	"github.com/stretchr/testify/require"
)

//nolint:funlen // Table covers exact CPU/GPU matching and rejects substitution.
func TestResolveInstanceType(t *testing.T) {
	t.Parallel()
	const cpuType = "cpu-e2.4vcpu-16gb"
	gpu := &compute.Platform{
		Metadata: &common.ResourceMetadata{
			Id:   "computeplatform-gpu",
			Name: "gpu-h100-sxm",
		},
		Spec: &compute.PlatformSpec{
			Presets: []*compute.Preset{
				{
					Name: "8gpu-128vcpu-1600gb",
					Resources: &compute.PresetResources{
						GpuCount: 8,
					},
				},
			},
		},
	}
	cpu := &compute.Platform{
		Metadata: &common.ResourceMetadata{
			Id:   "computeplatform-cpu",
			Name: "cpu-e2",
		},
		Spec: &compute.PlatformSpec{
			Presets: []*compute.Preset{
				{
					Name:      "4vcpu-16gb",
					Resources: &compute.PresetResources{},
				},
			},
		},
	}
	tests := []struct {
		name          string
		instanceType  string
		platforms     []*compute.Platform
		wantPlatform  string
		wantPreset    string
		errorContains string
	}{
		{
			name:          "absent CPU platform never falls back to GPU",
			instanceType:  cpuType,
			platforms:     []*compute.Platform{gpu},
			errorContains: "no exact platform/preset match",
		},
		{
			name:         "valid CPU platform is selected even when GPU is first",
			instanceType: cpuType,
			platforms:    []*compute.Platform{gpu, cpu},
			wantPlatform: "cpu-e2",
			wantPreset:   "4vcpu-16gb",
		},
		{
			name:         "exact GPU request is preserved",
			instanceType: "gpu-h100-sxm.8gpu-128vcpu-1600gb",
			platforms:    []*compute.Platform{cpu, gpu},
			wantPlatform: "gpu-h100-sxm",
			wantPreset:   "8gpu-128vcpu-1600gb",
		},
		{
			name:          "missing CPU preset never substitutes a different size",
			instanceType:  "cpu-e2.8vcpu-32gb",
			platforms:     []*compute.Platform{gpu, cpu},
			errorContains: "no exact platform/preset match",
		},
		{
			name:          "missing GPU platform never substitutes CPU",
			instanceType:  "gpu-h200-sxm.8gpu-128vcpu-1600gb",
			platforms:     []*compute.Platform{cpu},
			errorContains: "no exact platform/preset match",
		},
		{
			name:          "missing GPU preset never substitutes more GPUs",
			instanceType:  "gpu-h100-sxm.1gpu-16vcpu-200gb",
			platforms:     []*compute.Platform{gpu},
			errorContains: "no exact platform/preset match",
		},
		{
			name:          "family-only request cannot select a platform",
			instanceType:  "nebius-eu-north1-cpu-4vcpu-16gb",
			platforms:     []*compute.Platform{gpu, cpu},
			errorContains: "no exact platform/preset match",
		},
		{
			name:          "platform ID cannot replace platform name",
			instanceType:  "computeplatform-cpu-4vcpu-16gb",
			platforms:     []*compute.Platform{gpu, cpu},
			errorContains: "no exact platform/preset match",
		},
		{
			name:          "unknown type cannot match any provider default",
			instanceType:  "not-a-real-instance-type",
			platforms:     []*compute.Platform{gpu, cpu},
			errorContains: "no exact platform/preset match",
		},
		{
			name:          "empty type cannot match provider default",
			platforms:     []*compute.Platform{gpu, cpu},
			errorContains: "no exact platform/preset match",
		},
		{
			name:          "nil platform metadata cannot authorize provisioning",
			instanceType:  cpuType,
			platforms:     []*compute.Platform{nil, {}, {Metadata: &common.ResourceMetadata{Name: "cpu-e2"}}},
			errorContains: "no exact platform/preset match",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			platform, preset, err := resolveInstanceType(tt.instanceType, tt.platforms)
			if tt.errorContains != "" {
				require.ErrorContains(t, err, tt.errorContains)
				if tt.errorContains == "no exact platform/preset match" {
					require.ErrorIs(t, err, v1.ErrInsufficientResources)
				}
				require.Empty(t, platform)
				require.Empty(t, preset)
				return
			}
			require.NoError(t, err)
			require.Equal(t, tt.wantPlatform, platform)
			require.Equal(t, tt.wantPreset, preset)
		})
	}
}
