package v1

import (
	"context"
	"fmt"
	"github.com/brevdev/cloud/internal/collections"
	"slices"
	"time"

	v1 "github.com/brevdev/cloud/v1"
)

func (c *SFCClient) getInstanceTypeID(region string) string {
	return fmt.Sprintf("h100v_%v", region)
}

func (c *SFCClient) GetInstanceTypes(ctx context.Context, args v1.GetInstanceTypeArgs) ([]v1.InstanceType, error) {
	resp, err := c.client.Zones.List(ctx)
	if err != nil {
		return nil, err
	}
	types := make([]v1.InstanceType, 0)
	for _, zone := range resp.Data {
		var available = false
		if len(zone.AvailableCapacity) > 0 && zone.DeliveryType == "VM" {
			available = true
		}

		types = append(types, v1.InstanceType{
			ID:                  v1.InstanceTypeID(c.getInstanceTypeID(zone.Name)),
			IsAvailable:         available,
			Type:                "h100v",
			Location:            zone.Name,
			Stoppable:           false,
			Rebootable:          false,
			IsContainer:         false,
			EstimatedDeployTime: collections.Ptr(time.Duration(15 * time.Minute)),
			SupportedGPUs: []v1.GPU{{
				Count:        8,
				Type:         "h100v",
				Manufacturer: "nvidia",
				Name:         "h100v",
				MemoryBytes:  v1.NewBytes(80, v1.Gibibyte),
			}},
		})

	}
	return types, nil
}

func (c *SFCClient) GetLocations(ctx context.Context, _ v1.GetLocationsArgs) ([]v1.Location, error) {
	resp, err := c.client.Zones.List(ctx)
	if err != nil {
		return nil, err
	}
	locations := make(map[string]v1.Location)
	allowedZones := []string{"hayesvalley"}
	if resp != nil {
		for _, zone := range resp.Data {
			var available = false
			if len(zone.AvailableCapacity) > 0 && zone.DeliveryType == "VM" && slices.Contains(allowedZones, zone.Name) == true {
				available = true
				locations[zone.Name] = v1.Location{
					Name:        zone.Name,
					Description: fmt.Sprintf("sfc_%s_%s", zone.Name, string(zone.HardwareType)),
					Available:   available}
			} else {
				available = false
				locations[zone.Name] = v1.Location{
					Name:        zone.Name,
					Description: fmt.Sprintf("sfc_%s_%s", zone.Name, string(zone.HardwareType)),
					Available:   false}
			}
		}
	}
	availableLocations := []v1.Location{}
	for _, location := range locations {
		availableLocations = append(availableLocations, location)
	}
	return availableLocations, nil
}
