package fluidstack

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/alecthomas/units"
	"github.com/bojanz/currency"
	"github.com/brevdev/sdk/cloud"
	openapi "github.com/brevdev/sdk/cloud/fluidstack/gen/fluidstack"
)

func (c *FluidStackClient) GetInstanceTypes(ctx context.Context, _ cloud.GetInstanceTypeArgs) ([]cloud.InstanceType, error) {
	authCtx := c.makeAuthContext(ctx)
	projectCtx := c.makeProjectContext(authCtx)

	resp, httpResp, err := c.client.InstanceTypesAPI.ListInstanceTypes(projectCtx).Execute()
	if httpResp != nil && httpResp.Body != nil {
		defer func() { _ = httpResp.Body.Close() }()
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get instance types: %w", err)
	}

	var instanceTypes []cloud.InstanceType
	for _, fsInstanceType := range resp {
		instanceType := convertFluidStackInstanceTypeToV1InstanceType("", fsInstanceType, true)
		instanceTypes = append(instanceTypes, instanceType)
	}

	return instanceTypes, nil
}

func (c *FluidStackClient) GetInstanceTypePollTime() time.Duration {
	return 5 * time.Minute
}

func (c *FluidStackClient) GetLocations(ctx context.Context, _ cloud.GetLocationsArgs) ([]cloud.Location, error) {
	authCtx := c.makeAuthContext(ctx)
	projectCtx := c.makeProjectContext(authCtx)

	resp, httpResp, err := c.client.CapacityAPI.ListCapacity(projectCtx).Execute()
	if httpResp != nil && httpResp.Body != nil {
		defer func() { _ = httpResp.Body.Close() }()
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get locations: %w", err)
	}

	var locations []cloud.Location
	for _, capacity := range resp {
		location := cloud.Location{
			Name:        capacity.Name,
			Description: capacity.Name,
			Available:   capacity.Capacity > 0,
		}
		locations = append(locations, location)
	}

	return locations, nil
}

func convertFluidStackInstanceTypeToV1InstanceType(location string, fsInstanceType openapi.InstanceType, isAvailable bool) cloud.InstanceType {
	var gpus []cloud.GPU

	if fsInstanceType.GpuCount != nil && *fsInstanceType.GpuCount > 0 {
		count := int(*fsInstanceType.GpuCount)
		model := "GPU"
		if fsInstanceType.GpuModel != nil {
			model = *fsInstanceType.GpuModel
		}

		for i := 0; i < count; i++ {
			gpus = append(gpus, cloud.GPU{
				Name: model,
			})
		}
	}

	var ram units.Base2Bytes
	if fsInstanceType.Memory != "" {
		memoryStr := strings.TrimSuffix(fsInstanceType.Memory, "GB")
		memoryStr = strings.TrimSpace(memoryStr)
		if memoryGB, err := strconv.ParseFloat(memoryStr, 64); err == nil {
			ram = units.Base2Bytes(memoryGB) * units.Gibibyte
		}
	}

	vcpus := fsInstanceType.Cpu
	if vcpus < 0 {
		vcpus = 0
	}

	price, _ := currency.NewAmount("0", "USD")

	return cloud.InstanceType{
		Type:          fsInstanceType.Name,
		VCPU:          vcpus,
		Memory:        ram,
		SupportedGPUs: gpus,
		BasePrice:     &price,
		IsAvailable:   isAvailable,
		Location:      location,
		Provider:      CloudProviderID,
	}
}
