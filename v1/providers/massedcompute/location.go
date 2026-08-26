package massedcompute

import (
	"context"

	v1 "github.com/brevdev/cloud/v1"
)

func (c *MassedComputeClient) GetLocations(ctx context.Context, args v1.GetLocationsArgs) ([]v1.Location, error) {
	inventory, err := c.getInventory(ctx)
	if err != nil {
		return nil, err
	}

	byName := make(map[string]v1.Location)
	for _, item := range inventory {
		if item.InstanceType == nil || isSpotDescription(stringValue(item.InstanceType.Description)) {
			continue
		}
		available := item.CapacityAvailable == nil || *item.CapacityAvailable > 0
		for _, region := range item.RegionsWithCapacityAvailable {
			name := stringValue(region.Name)
			if name == "" || (!available && !args.IncludeUnavailable) {
				continue
			}
			location := byName[name]
			location.Name = name
			if location.Description == "" {
				location.Description = stringValue(region.Description)
			}
			location.Available = location.Available || available
			byName[name] = location
		}
	}

	locations := make([]v1.Location, 0, len(byName))
	for _, location := range byName {
		locations = append(locations, location)
	}
	return locations, nil
}
