package v1

import (
	"context"
	"fmt"
	"slices"

	v1 "github.com/brevdev/cloud/v1"
)

func (c *SFCClient) getInstanceTypeID(region string) string {
	return fmt.Sprintf("h100v_%v", region)
}

func (c *SFCClient) GetLocations(ctx context.Context, _ v1.GetLocationsArgs) ([]v1.Location, error) {
	resp, err := c.client.Zones.List(ctx)
	if err != nil {
		return nil, err
	}
	var locations map[string]v1.Location
	allowedZones := []string{"hayesvalley"}
	if resp != nil {
		for _, zone := range resp.Data {
			var available = false
			if len(zone.AvailableCapacity) != 0 && zone.DeliveryType == "VM" && slices.Contains(allowedZones, zone.Name) == true {
				available = true
				locations[zone.Name] = v1.Location{
					Name:        zone.Name,
					Description: string(zone.HardwareType),
					Available:   available}
			}
		}
	}
	availableLocations := []v1.Location{}
	for _, location := range locations {
		availableLocations = append(availableLocations, location)
	}
	return availableLocations, nil
}
