package v1

import (
	"context"
	"time"

	"github.com/alecthomas/units"
	"github.com/bojanz/currency"
	cloudv1 "github.com/brevdev/cloud/v1"
	corev1 "k8s.io/api/core/v1"
)

const (
	DefaultImageID = "testkube-ubuntu-vm"
	DefaultImage   = "ghcr.io/brevdev/cloud/testkube-ubuntu-vm:latest"

	InstanceTypeOKCPU        = "test.ok.cpu"
	InstanceTypeFailCapacity = "test.fail.capacity"
	InstanceTypeFailQuota    = "test.fail.quota"
	InstanceTypeFailBuild    = "test.fail.build"
)

type testInstanceTypeSpec struct {
	instanceType cloudv1.InstanceType
	imageID      string
	image        string
	serviceType  corev1.ServiceType
}

var testInstanceTypeSpecs = []testInstanceTypeSpec{
	makeTestInstanceTypeSpec(InstanceTypeOKCPU),
	makeTestInstanceTypeSpec(InstanceTypeFailCapacity),
	makeTestInstanceTypeSpec(InstanceTypeFailQuota),
	makeTestInstanceTypeSpec(InstanceTypeFailBuild),
}

func makeTestInstanceTypeSpec(instanceType string) testInstanceTypeSpec {
	estimatedDeployTime := 20 * time.Second
	return testInstanceTypeSpec{
		instanceType: makeCPUInstanceType(instanceType, true, &estimatedDeployTime),
		imageID:      DefaultImageID,
		image:        DefaultImage,
		serviceType:  corev1.ServiceTypeLoadBalancer,
	}
}

func (c *TestKubeClient) GetInstanceTypes(_ context.Context, args cloudv1.GetInstanceTypeArgs) ([]cloudv1.InstanceType, error) {
	instanceTypes := c.testInstanceTypes()
	instanceTypes = filterInstanceTypes(instanceTypes, args)
	return instanceTypes, nil
}

func (c *TestKubeClient) testInstanceTypes() []cloudv1.InstanceType {
	instanceTypes := make([]cloudv1.InstanceType, 0, len(testInstanceTypeSpecs))
	for _, spec := range testInstanceTypeSpecs {
		instanceTypes = append(instanceTypes, c.instanceTypeFromSpec(spec))
	}
	return instanceTypes
}

func (c *TestKubeClient) instanceTypeFromSpec(spec testInstanceTypeSpec) cloudv1.InstanceType {
	instanceType := spec.instanceType
	instanceType.Location = c.location
	instanceType.ID = cloudv1.MakeGenericInstanceTypeID(instanceType)
	return instanceType
}

func makeCPUInstanceType(instanceType string, available bool, estimatedDeployTime *time.Duration) cloudv1.InstanceType {
	basePrice, _ := currency.NewAmountFromInt64(0, "USD")
	it := cloudv1.InstanceType{
		Type: instanceType,
		SupportedStorage: []cloudv1.Storage{
			{
				Type:        "ephemeral",
				Count:       1,
				Size:        units.GiB * 20,
				SizeBytes:   cloudv1.NewBytes(20, cloudv1.Gibibyte),
				IsEphemeral: true,
			},
		},
		ElasticRootVolume:     false,
		SupportedUsageClasses: []string{"on-demand"},
		Memory:                units.GiB * 4,
		MemoryBytes:           cloudv1.NewBytes(4, cloudv1.Gibibyte),
		SupportedNumCores:     []int32{2},
		DefaultCores:          2,
		VCPU:                  2,
		SupportedArchitectures: []cloudv1.Architecture{
			cloudv1.ArchitectureX86_64,
		},
		Stoppable:           true,
		Rebootable:          true,
		IsAvailable:         available,
		BasePrice:           &basePrice,
		IsContainer:         true,
		EstimatedDeployTime: estimatedDeployTime,
		Provider:            CloudProviderID,
		Cloud:               CloudProviderID,
	}
	return it
}

func filterInstanceTypes(instanceTypes []cloudv1.InstanceType, args cloudv1.GetInstanceTypeArgs) []cloudv1.InstanceType {
	filtered := make([]cloudv1.InstanceType, 0, len(instanceTypes))
	for _, instanceType := range instanceTypes {
		if len(args.Locations) > 0 && !args.Locations.IsAll() && !args.Locations.IsAllowed(instanceType.Location) {
			continue
		}
		if len(args.InstanceTypes) > 0 && !containsString(args.InstanceTypes, instanceType.Type) {
			continue
		}
		if args.CloudFilter != nil && !args.CloudFilter.IsAllowed(instanceType.Cloud) {
			continue
		}
		if args.ArchitectureFilter != nil && !isArchitectureAllowed(args.ArchitectureFilter, instanceType.SupportedArchitectures) {
			continue
		}
		if args.GPUManufactererFilter != nil && !isGPUManufacturerAllowed(args.GPUManufactererFilter, instanceType.SupportedGPUs) {
			continue
		}
		filtered = append(filtered, instanceType)
	}
	return filtered
}

func isArchitectureAllowed(filter *cloudv1.ArchitectureFilter, architectures []cloudv1.Architecture) bool {
	for _, architecture := range architectures {
		if filter.IsAllowed(architecture) {
			return true
		}
	}
	return false
}

func isGPUManufacturerAllowed(filter *cloudv1.GPUManufacturerFilter, gpus []cloudv1.GPU) bool {
	if len(gpus) == 0 {
		return filter.IsAllowed(cloudv1.ManufacturerUnknown)
	}
	for _, gpu := range gpus {
		if filter.IsAllowed(gpu.Manufacturer) {
			return true
		}
	}
	return false
}

func getInstanceTypeSpec(instanceType string) (testInstanceTypeSpec, bool) {
	for _, spec := range testInstanceTypeSpecs {
		if spec.instanceType.Type == instanceType {
			return spec, true
		}
	}
	return testInstanceTypeSpec{}, false
}

func containsString(values []string, value string) bool {
	for _, v := range values {
		if v == value {
			return true
		}
	}
	return false
}
