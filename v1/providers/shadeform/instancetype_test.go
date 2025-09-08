package v1

import (
	"testing"

	v1 "github.com/brevdev/cloud/v1"
	"github.com/stretchr/testify/assert"
)

func TestIsSelectedByArgs(t *testing.T) {
	t.Parallel()

	x8664nvidiaaws := v1.InstanceType{SupportedArchitectures: []v1.Architecture{v1.ArchitectureX86_64}, SupportedGPUs: []v1.GPU{{Manufacturer: v1.ManufacturerNVIDIA}}, Cloud: "aws"}
	x8664nvidiagcp := v1.InstanceType{SupportedArchitectures: []v1.Architecture{v1.ArchitectureX86_64}, SupportedGPUs: []v1.GPU{{Manufacturer: v1.ManufacturerNVIDIA}}, Cloud: "gcp"}
	x8664intelaws := v1.InstanceType{SupportedArchitectures: []v1.Architecture{v1.ArchitectureX86_64}, SupportedGPUs: []v1.GPU{{Manufacturer: v1.ManufacturerIntel}}, Cloud: "aws"}
	x8664intelgcp := v1.InstanceType{SupportedArchitectures: []v1.Architecture{v1.ArchitectureX86_64}, SupportedGPUs: []v1.GPU{{Manufacturer: v1.ManufacturerIntel}}, Cloud: "gcp"}
	arm64nvidiaaws := v1.InstanceType{SupportedArchitectures: []v1.Architecture{v1.ArchitectureARM64}, SupportedGPUs: []v1.GPU{{Manufacturer: v1.ManufacturerNVIDIA}}, Cloud: "aws"}
	arm64nvidiagcp := v1.InstanceType{SupportedArchitectures: []v1.Architecture{v1.ArchitectureARM64}, SupportedGPUs: []v1.GPU{{Manufacturer: v1.ManufacturerNVIDIA}}, Cloud: "gcp"}
	arm64intelaws := v1.InstanceType{SupportedArchitectures: []v1.Architecture{v1.ArchitectureARM64}, SupportedGPUs: []v1.GPU{{Manufacturer: v1.ManufacturerIntel}}, Cloud: "aws"}
	arm64intelgcp := v1.InstanceType{SupportedArchitectures: []v1.Architecture{v1.ArchitectureARM64}, SupportedGPUs: []v1.GPU{{Manufacturer: v1.ManufacturerIntel}}, Cloud: "gcp"}

	all := []v1.InstanceType{x8664nvidiaaws, x8664intelaws, arm64nvidiaaws, arm64intelaws, x8664nvidiagcp, arm64nvidiagcp, x8664intelgcp, arm64intelgcp}

	cases := []struct {
		name          string
		instanceTypes []v1.InstanceType
		args          v1.GetInstanceTypeArgs
		want          []v1.InstanceType
	}{
		{
			name:          "no filters",
			instanceTypes: all,
			args:          v1.GetInstanceTypeArgs{},
			want:          all,
		},
		{
			name:          "include only x86_64 architecture",
			instanceTypes: all,
			args:          v1.GetInstanceTypeArgs{ArchitectureFilter: &v1.ArchitectureFilter{IncludeArchitectures: []v1.Architecture{v1.ArchitectureX86_64}}},
			want:          []v1.InstanceType{x8664nvidiaaws, x8664intelaws, x8664nvidiagcp, x8664intelgcp},
		},
		{
			name:          "exclude x86_64 architecture",
			instanceTypes: all,
			args:          v1.GetInstanceTypeArgs{ArchitectureFilter: &v1.ArchitectureFilter{ExcludeArchitectures: []v1.Architecture{v1.ArchitectureX86_64}}},
			want:          []v1.InstanceType{arm64nvidiaaws, arm64intelaws, arm64nvidiagcp, arm64intelgcp},
		},
		{
			name:          "include only nvidia manufacturer",
			instanceTypes: all,
			args:          v1.GetInstanceTypeArgs{GPUManufactererFilter: &v1.GPUManufacturerFilter{IncludeGPUManufacturers: []v1.Manufacturer{v1.ManufacturerNVIDIA}}},
			want:          []v1.InstanceType{x8664nvidiaaws, x8664nvidiagcp, arm64nvidiaaws, arm64nvidiagcp},
		},
		{
			name:          "exclude nvidia manufacturer",
			instanceTypes: all,
			args:          v1.GetInstanceTypeArgs{GPUManufactererFilter: &v1.GPUManufacturerFilter{ExcludeGPUManufacturers: []v1.Manufacturer{v1.ManufacturerNVIDIA}}},
			want:          []v1.InstanceType{x8664intelaws, x8664intelgcp, arm64intelaws, arm64intelgcp},
		},
		{
			name:          "include only aws cloud",
			instanceTypes: all,
			args:          v1.GetInstanceTypeArgs{CloudFilter: &v1.CloudFilter{IncludeClouds: []string{"aws"}}},
			want:          []v1.InstanceType{x8664nvidiaaws, x8664intelaws, arm64nvidiaaws, arm64intelaws},
		},
		{
			name:          "exclude aws cloud",
			instanceTypes: all,
			args:          v1.GetInstanceTypeArgs{CloudFilter: &v1.CloudFilter{ExcludeClouds: []string{"aws"}}},
			want:          []v1.InstanceType{x8664nvidiagcp, x8664intelgcp, arm64nvidiagcp, arm64intelgcp},
		},
		{
			name:          "include only aws cloud, exclude arm64 architecture, include nvidia manufacturer",
			instanceTypes: all,
			args: v1.GetInstanceTypeArgs{
				CloudFilter:           &v1.CloudFilter{IncludeClouds: []string{"aws"}},
				ArchitectureFilter:    &v1.ArchitectureFilter{ExcludeArchitectures: []v1.Architecture{v1.ArchitectureARM64}},
				GPUManufactererFilter: &v1.GPUManufacturerFilter{IncludeGPUManufacturers: []v1.Manufacturer{v1.ManufacturerNVIDIA}},
			},
			want: []v1.InstanceType{x8664nvidiaaws},
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			selectedInstanceTypes := []v1.InstanceType{}
			for _, instanceType := range tt.instanceTypes {
				if isSelectedByArgs(instanceType, tt.args) {
					selectedInstanceTypes = append(selectedInstanceTypes, instanceType)
				}
			}
			assert.ElementsMatch(t, tt.want, selectedInstanceTypes)
		})
	}
}
