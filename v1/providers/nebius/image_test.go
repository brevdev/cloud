package v1

import (
	"testing"

	common "github.com/nebius/gosdk/proto/nebius/common/v1"
	compute "github.com/nebius/gosdk/proto/nebius/compute/v1"
	"github.com/stretchr/testify/require"

	cloudv1 "github.com/brevdev/cloud/v1"
)

func TestExtractArchitecture(t *testing.T) {
	t.Parallel()

	image := func(
		architecture compute.ImageSpec_CPUArchitecture,
		labels map[string]string,
	) *compute.Image {
		return &compute.Image{
			Metadata: &common.ResourceMetadata{Labels: labels},
			Spec:     &compute.ImageSpec{CpuArchitecture: architecture},
		}
	}

	tests := []struct {
		name  string
		image *compute.Image
		want  cloudv1.Architecture
	}{
		{name: "AMD64 enum", image: image(compute.ImageSpec_AMD64, nil), want: cloudv1.ArchitectureX86_64},
		{
			name: "ARM64 enum wins over legacy metadata",
			image: image(
				compute.ImageSpec_ARM64,
				map[string]string{"architecture": "amd64"},
			),
			want: cloudv1.ArchitectureARM64,
		},
		{name: "legacy label", image: image(compute.ImageSpec_UNSPECIFIED, map[string]string{"arch": "aarch64"}), want: cloudv1.ArchitectureARM64},
		{name: "unknown enum", image: image(compute.ImageSpec_CPUArchitecture(99), map[string]string{"architecture": "amd64"}), want: cloudv1.ArchitectureUnknown},
		{name: "unknown", image: &compute.Image{}, want: cloudv1.ArchitectureUnknown},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			require.Equal(t, string(tt.want), extractArchitecture(tt.image))
		})
	}
}

func TestApplyImageFiltersDoesNotDefaultToX86(t *testing.T) {
	t.Parallel()

	x86 := cloudv1.Image{ID: "x86", Architecture: string(cloudv1.ArchitectureX86_64)}
	arm := cloudv1.Image{ID: "arm", Architecture: string(cloudv1.ArchitectureARM64)}
	images := []cloudv1.Image{x86, arm}

	require.Equal(t, images, applyImageFilters(images, cloudv1.GetImageArgs{}))
}
