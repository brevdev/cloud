package v1

import (
	"context"
	"errors"
	"fmt"
	"github.com/alecthomas/units"
	"github.com/bojanz/currency"
	"strings"
	"time"

	openapi "github.com/brevdev/cloud/internal/shadeform/gen/shadeform"

	"github.com/brevdev/cloud/pkg/v1"
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
		// Filter the list down to the instance types that are allowed by the configuration filter
		for _, singleInstanceType := range instanceTypesFromShadeformInstanceType {
			if c.isInstanceTypeAllowed(singleInstanceType.Type) {
				instanceTypes = append(instanceTypes, singleInstanceType)
			}
		}
	}

	return instanceTypes, nil
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

// convertShadeformInstanceTypeToV1InstanceTypes - converts a shadeform returned instance type to a specific instance type and region of availability
func (c *ShadeformClient) convertShadeformInstanceTypeToV1InstanceType(shadeformInstanceType openapi.InstanceType) ([]v1.InstanceType, error) {
	instanceType := c.getInstanceType(string(shadeformInstanceType.Cloud), shadeformInstanceType.ShadeInstanceType)

	instanceTypes := []v1.InstanceType{}

	basePrice, err := convertHourlyPriceToAmount(shadeformInstanceType.HourlyPrice)
	if err != nil {
		return nil, err
	}

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
					NetworkDetails: "",
					Manufacturer:   "",
					Name:           shadeformInstanceType.Configuration.GpuType,
					Type:           shadeformInstanceType.Configuration.GpuType,
				},
			},
			BasePrice:   basePrice,
			IsAvailable: region.Available,
			Location:    region.Region,
			Provider:    CloudProviderID,
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
