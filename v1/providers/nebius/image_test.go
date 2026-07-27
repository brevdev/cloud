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

	tests := []struct {
		name  string
		image *compute.Image
		want  string
	}{
		{
			name: "authoritative AMD64 enum",
			image: &compute.Image{
				Spec: &compute.ImageSpec{
					CpuArchitecture: compute.ImageSpec_AMD64,
				},
			},
			want: string(cloudv1.ArchitectureX86_64),
		},
		{
			name: "authoritative ARM64 enum",
			image: &compute.Image{
				Spec: &compute.ImageSpec{
					CpuArchitecture: compute.ImageSpec_ARM64,
				},
			},
			want: string(cloudv1.ArchitectureARM64),
		},
		{
			name: "enum wins over legacy label",
			image: &compute.Image{
				Metadata: &common.ResourceMetadata{
					Labels: map[string]string{"architecture": "amd64"},
				},
				Spec: &compute.ImageSpec{
					CpuArchitecture: compute.ImageSpec_ARM64,
				},
			},
			want: string(cloudv1.ArchitectureARM64),
		},
		{
			name: "legacy label is canonicalized",
			image: &compute.Image{
				Metadata: &common.ResourceMetadata{
					Labels: map[string]string{"arch": "aarch64"},
				},
			},
			want: string(cloudv1.ArchitectureARM64),
		},
		{
			name: "legacy name fallback",
			image: &compute.Image{
				Metadata: &common.ResourceMetadata{Name: "ubuntu-24.04-amd64"},
			},
			want: string(cloudv1.ArchitectureX86_64),
		},
		{
			name:  "unknown stays unknown",
			image: &compute.Image{},
			want:  string(cloudv1.ArchitectureUnknown),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			require.Equal(t, tt.want, extractArchitecture(tt.image))
		})
	}
}

func TestApplyImageFiltersWithoutArchitectureReturnsAll(t *testing.T) {
	t.Parallel()

	images := []cloudv1.Image{
		{ID: "amd64", Architecture: string(cloudv1.ArchitectureX86_64)},
		{ID: "arm64", Architecture: string(cloudv1.ArchitectureARM64)},
		{ID: "unknown", Architecture: string(cloudv1.ArchitectureUnknown)},
	}

	require.Equal(
		t,
		images,
		applyImageFilters(images, cloudv1.GetImageArgs{}),
	)
}
