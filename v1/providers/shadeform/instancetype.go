package v1

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/alecthomas/units"
	"github.com/bojanz/currency"

	v1 "github.com/brevdev/cloud/v1"
	openapi "github.com/brevdev/cloud/v1/providers/shadeform/gen/shadeform"
)

const (
	UsdCurrentCode = "USD"
	AllRegions     = "all"
)

// TODO: We need to apply a filter to specifically limit the integration and api to selected clouds and shade instance types

func (c *ShadeformClient) GetInstanceTypes(ctx context.Context, args v1.GetInstanceTypeArgs) ([]v1.InstanceType, error) {
	authCtx := c.makeAuthContext(ctx)

	request := c.client.DefaultAPI.InstancesTypes(authCtx)
	if len(args.Locations) > 0 && args.Locations[0] != AllRegions {
		regionFilter := args.Locations[0]
		request = request.Region(regionFilter)
	}

	resp, httpResp, err := request.Execute()
	if httpResp != nil && httpResp.Body != nil {
		defer func() { _ = httpResp.Body.Close() }()
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get instance types: %w", err)
	}

	var instanceTypes []v1.InstanceType
	for _, sfInstanceType := range resp.InstanceTypes {
		instanceTypesFromShadeformInstanceType, err := c.convertShadeformInstanceTypeToV1InstanceType(sfInstanceType)
		if err != nil {
			return nil, err
		}
		// Filter the list down to the instance types that are allowed by the configuration filter and the args
		for _, singleInstanceType := range instanceTypesFromShadeformInstanceType {
			if !isSelectedByArgs(singleInstanceType, args) {
				continue
			}
			if c.isInstanceTypeAllowed(singleInstanceType.Type) {
				instanceTypes = append(instanceTypes, singleInstanceType)
			}
		}
	}

	return instanceTypes, nil
}

func isSelectedByArgs(instanceType v1.InstanceType, args v1.GetInstanceTypeArgs) bool {
	if args.GPUManufactererFilter != nil {
		for _, supportedGPU := range instanceType.SupportedGPUs {
			if !args.GPUManufactererFilter.IsAllowed(supportedGPU.Manufacturer) {
				return false
			}
		}
	}

	if args.CloudFilter != nil {
		if !args.CloudFilter.IsAllowed(instanceType.Cloud) {
			return false
		}
	}

	if args.ArchitectureFilter != nil {
		for _, architecture := range instanceType.SupportedArchitectures {
			if !args.ArchitectureFilter.IsAllowed(architecture) {
				return false
			}
		}
	}

	return true
}

func (c *ShadeformClient) GetInstanceTypePollTime() time.Duration {
	return 5 * time.Minute
}

func (c *ShadeformClient) GetLocations(ctx context.Context, _ v1.GetLocationsArgs) ([]v1.Location, error) {
	authCtx := c.makeAuthContext(ctx)

	resp, httpResp, err := c.client.DefaultAPI.InstancesTypes(authCtx).Execute()
	if httpResp != nil && httpResp.Body != nil {
		defer func() { _ = httpResp.Body.Close() }()
	}

	if err != nil {
		return nil, fmt.Errorf("failed to get locations: %w", err)
	}

	// Shadeform doesn't have a dedicated locations API but we can get the same result from using the
	// instance types API and formatting the output

	dedupedLocations := map[string]v1.Location{}

	if resp != nil {
		for _, instanceType := range resp.InstanceTypes {
			for _, availability := range instanceType.Availability {
				_, ok := dedupedLocations[availability.Region]
				if !ok {
					dedupedLocations[availability.Region] = v1.Location{
						Name:        availability.Region,
						Description: availability.DisplayName,
						Available:   availability.Available,
					}
				}
			}
		}
	}

	locations := []v1.Location{}

	for _, location := range dedupedLocations {
		locations = append(locations, location)
	}

	return locations, nil
}

// isInstanceTypeAllowed - determines if an instance type is allowed based on configuration
func (c *ShadeformClient) isInstanceTypeAllowed(instanceType string) bool {
	// By default, everything is allowed
	if c.config == nil || c.config.AllowedInstanceTypes == nil {
		return true
	}

	// Convert to Cloud and Instance Type
	cloud, shadeInstanceType, err := c.getShadeformCloudAndInstanceType(instanceType)
	if err != nil {
		return false
	}

	// Convert to API Cloud Enum
	cloudEnum, err := openapi.NewCloudFromValue(cloud)
	if err != nil {
		return false
	}

	return c.config.isAllowed(*cloudEnum, shadeInstanceType)
}

// getInstanceType - gets the Brev instance type from the shadeform cloud and shade instance type
// TODO: determine if it would be better to include the shadeform cloud inside the region / location instead
func (c *ShadeformClient) getInstanceType(shadeformCloud string, shadeformInstanceType string) string {
	return fmt.Sprintf("%v_%v", shadeformCloud, shadeformInstanceType)
}

// getInstanceTypeID - unique identifier for the SKU
func (c *ShadeformClient) getInstanceTypeID(instanceType string, region string) string {
	return fmt.Sprintf("%v_%v", instanceType, region)
}

func (c *ShadeformClient) getShadeformCloudAndInstanceType(instanceType string) (string, string, error) {
	shadeformCloud, shadeformInstanceType, found := strings.Cut(instanceType, "_")
	if !found {
		return "", "", errors.New("could not determine shadeform cloud and instance type from instance type")
	}
	return shadeformCloud, shadeformInstanceType, nil
}

func (c *ShadeformClient) getEstimatedDeployTime(shadeformInstanceType openapi.InstanceType) *time.Duration {
	bootTime := shadeformInstanceType.BootTime
	if bootTime == nil {
		return nil
	}

	minSec := bootTime.MinBootInSec
	maxSec := bootTime.MaxBootInSec

	var estimatedDeployTime *time.Duration
	if minSec != nil && maxSec != nil { //nolint:gocritic // if else fine
		avg := (*minSec + *maxSec) / 2
		avgDuration := time.Duration(avg) * time.Second
		estimatedDeployTime = &avgDuration
	} else if minSec != nil {
		d := time.Duration(*minSec) * time.Second
		estimatedDeployTime = &d
	} else if maxSec != nil {
		d := time.Duration(*maxSec) * time.Second
		estimatedDeployTime = &d
	}
	return estimatedDeployTime
}

// convertShadeformInstanceTypeToV1InstanceTypes - converts a shadeform returned instance type to a specific instance type and region of availability
func (c *ShadeformClient) convertShadeformInstanceTypeToV1InstanceType(shadeformInstanceType openapi.InstanceType) ([]v1.InstanceType, error) {
	instanceType := c.getInstanceType(string(shadeformInstanceType.Cloud), shadeformInstanceType.ShadeInstanceType)

	instanceTypes := []v1.InstanceType{}

	basePrice, err := convertHourlyPriceToAmount(shadeformInstanceType.HourlyPrice)
	if err != nil {
		return nil, err
	}

	gpuName := shadeformGPUTypeToBrevGPUName(shadeformInstanceType.Configuration.GpuType)
	gpuManufacturer := v1.GetManufacturer(shadeformInstanceType.Configuration.GpuManufacturer)
	cloud := shadeformCloud(shadeformInstanceType.Cloud)
	architecture := shadeformArchitecture(gpuName)

	estimatedDeployTime := c.getEstimatedDeployTime(shadeformInstanceType)

	for _, region := range shadeformInstanceType.Availability {
		instanceTypes = append(instanceTypes, v1.InstanceType{
			ID:     v1.InstanceTypeID(c.getInstanceTypeID(instanceType, region.Region)),
			Type:   instanceType,
			VCPU:   shadeformInstanceType.Configuration.Vcpus,
			Memory: units.Base2Bytes(shadeformInstanceType.Configuration.MemoryInGb) * units.GiB,
			SupportedGPUs: []v1.GPU{
				{
					Count:          shadeformInstanceType.Configuration.NumGpus,
					Memory:         units.Base2Bytes(shadeformInstanceType.Configuration.VramPerGpuInGb) * units.GiB,
					MemoryDetails:  "",
					NetworkDetails: shadeformInstanceType.Configuration.Interconnect,
					Manufacturer:   gpuManufacturer,
					Name:           gpuName,
					Type:           shadeformInstanceType.Configuration.GpuType,
				},
			},
			SupportedStorage: []v1.Storage{ // TODO: add storage (look in configuration)
				{
					Type:  "ssd",
					Count: 1,
					Size:  units.Base2Bytes(shadeformInstanceType.Configuration.StorageInGb) * units.GiB,
				},
			},
			SupportedArchitectures: []v1.Architecture{architecture},
			BasePrice:              basePrice,
			IsAvailable:            region.Available,
			Location:               region.Region,
			Provider:               CloudProviderID,
			Cloud:                  cloud,
			EstimatedDeployTime:    estimatedDeployTime,
		})
	}

	return instanceTypes, nil
}

func convertHourlyPriceToAmount(hourlyPrice int32) (*currency.Amount, error) {
	number := fmt.Sprintf("%.2f", float64(hourlyPrice)/100)

	amount, err := currency.NewAmount(number, UsdCurrentCode)
	if err != nil {
		return nil, err
	}
	return &amount, nil
}

func shadeformGPUTypeToBrevGPUName(gpuType string) string {
	// Shadeform may include a memory size as a suffix. This must be cleaned up before
	// being used as a name.
	// e.g. A100_80GB -> A100, H100_40GB -> H100
	gpuType = strings.Split(gpuType, "_")[0]
	return gpuType
}

func shadeformCloud(cloud openapi.Cloud) string {
	// Shadeform will return the cloud as "excesssupply" if the instance type is retrieved
	// from cloud partners and not a direct cloud provider. In this case, we should just return
	// the Shadeform Cloud Provider ID.
	if cloud == openapi.EXCESSSUPPLY {
		return CloudProviderID
	}

	return string(cloud)
}

func shadeformArchitecture(gpuName string) v1.Architecture {
	// Shadeform currently does not specify the architecture, so we need to infer it from the GPU name.
	if strings.HasPrefix(gpuName, "GH") || strings.HasPrefix(gpuName, "GB") {
		return v1.ArchitectureARM64
	}
	return v1.ArchitectureX86_64
}
