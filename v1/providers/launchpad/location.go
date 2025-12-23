package v1

import (
	"context"

	"github.com/brevdev/cloud/internal/clouderrors"
	openapi "github.com/brevdev/cloud/v1/providers/launchpad/gen/launchpad"

	v1 "github.com/brevdev/cloud/v1"
)

func (c *LaunchpadClient) GetLocations(ctx context.Context, args v1.GetLocationsArgs) ([]v1.Location, error) {
	launchpadInstanceTypes, err := c.paginateInstanceTypes(ctx, 100)
	if err != nil {
		return nil, clouderrors.WrapAndTrace(err)
	}

	locations := launchpadInstanceTypesToLocations(launchpadInstanceTypes, args)

	return locations, nil
}

func launchpadInstanceTypesToLocations(launchpadInstanceTypes []openapi.InstanceType, args v1.GetLocationsArgs) []v1.Location {
	locationsSet := make(map[string]v1.Location)
	for _, launchpadInstanceType := range launchpadInstanceTypes {
		for locationName, capacity := range launchpadInstanceType.Capacity {
			if capacity <= 0 && !args.IncludeUnavailable {
				continue
			}
			if _, ok := locationsSet[locationName]; ok {
				continue
			}
			locationsSet[locationName] = v1.Location{
				Name:        locationName,
				Description: locationName,
				Available:   capacity > 0,
			}
		}
	}

	locations := []v1.Location{}
	for _, location := range locationsSet {
		locations = append(locations, location)
	}

	return locations
}
