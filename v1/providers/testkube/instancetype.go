package v1

import (
	"context"
	"slices"
	"time"

	"github.com/alecthomas/units"
	"github.com/bojanz/currency"
	cloudv1 "github.com/brevdev/cloud/v1"
	corev1 "k8s.io/api/core/v1"
)

const (
	DefaultImageID = "testkube-ubuntu-vm"
	DefaultImage   = "ghcr.io/brevdev/cloud/testkube-ubuntu-vm:latest"

	DefaultPriceCentsPerHour = 1

	InstanceTypeOKCPU        = "test.ok.cpu"
	InstanceTypeFailCapacity = "test.fail.capacity"
	InstanceTypeFailQuota    = "test.fail.quota"
	InstanceTypeFailBuild    = "test.fail.build"
)

// instanceTypeSpec is used mainly as a tuple of instance type (from devplane) and service type (from k8s). When a request
// for instance provisioning is made, we need to determine the appropriate service type to use based on the incoming instance type.
type instanceTypeSpec struct {
	instanceType cloudv1.InstanceType
	imageID      string
	image        string
	serviceType  corev1.ServiceType
}

var allInstanceTypeSpecs = []instanceTypeSpec{
	makeInstanceTypeSpec(InstanceTypeOKCPU),
	makeInstanceTypeSpec(InstanceTypeFailCapacity),
	makeInstanceTypeSpec(InstanceTypeFailQuota),
	makeInstanceTypeSpec(InstanceTypeFailBuild),
}

func makeInstanceTypeSpec(instanceType string) instanceTypeSpec {
	estimatedDeployTime := 20 * time.Second
	return instanceTypeSpec{
		instanceType: makeCPUInstanceType(instanceType, true, &estimatedDeployTime),
		imageID:      DefaultImageID,
		image:        DefaultImage,
		serviceType:  corev1.ServiceTypeLoadBalancer,
	}
}

func (c *TestKubeClient) GetInstanceTypes(_ context.Context, args cloudv1.GetInstanceTypeArgs) ([]cloudv1.InstanceType, error) {
	// Instance types are statically defined, but in the future we should consider dynamic types, with capacity numbers based on
	// test input or devplane configuration.
	instanceTypes := c.allInstanceTypes()

	// Filter the instance types as any normal provider would do.
	instanceTypes = filterInstanceTypes(instanceTypes, args)
	return instanceTypes, nil
}

func (c *TestKubeClient) allInstanceTypes() []cloudv1.InstanceType {
	instanceTypes := make([]cloudv1.InstanceType, 0, len(allInstanceTypeSpecs))
	for _, spec := range allInstanceTypeSpecs {
		instanceTypes = append(instanceTypes, c.instanceTypeSpecToBrevInstanceType(spec))
	}
	return instanceTypes
}

func filterInstanceTypes(instanceTypes []cloudv1.InstanceType, args cloudv1.GetInstanceTypeArgs) []cloudv1.InstanceType {
	filtered := make([]cloudv1.InstanceType, 0, len(instanceTypes))
	for _, instanceType := range instanceTypes {
		if len(args.Locations) > 0 && !args.Locations.IsAll() && !args.Locations.IsAllowed(instanceType.Location) {
			continue
		}
		if len(args.InstanceTypes) > 0 && !slices.Contains(args.InstanceTypes, instanceType.Type) {
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
		return true // NB: always return CPU types
	}
	for _, gpu := range gpus {
		if filter.IsAllowed(gpu.Manufacturer) {
			return true
		}
	}
	return false
}

func (c *TestKubeClient) instanceTypeSpecToBrevInstanceType(spec instanceTypeSpec) cloudv1.InstanceType {
	instanceType := spec.instanceType
	instanceType.Location = c.location
	instanceType.ID = cloudv1.MakeGenericInstanceTypeID(instanceType)
	return instanceType
}

func makeCPUInstanceType(instanceType string, available bool, estimatedDeployTime *time.Duration) cloudv1.InstanceType {
	basePrice, _ := currency.NewAmountFromInt64(DefaultPriceCentsPerHour, "USD")
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
		IsContainer:         false,
		EstimatedDeployTime: estimatedDeployTime,
		Provider:            CloudProviderID,
		Cloud:               CloudProviderID,
	}
	return it
}

func getInstanceTypeSpec(instanceType string) (instanceTypeSpec, bool) {
	for _, spec := range allInstanceTypeSpecs {
		if spec.instanceType.Type == instanceType {
			return spec, true
		}
	}
	return instanceTypeSpec{}, false
}
