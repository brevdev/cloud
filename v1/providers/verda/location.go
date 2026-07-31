package verda

import (
	"context"
	"sort"
	"strings"

	v1 "github.com/brevdev/cloud/v1"
)

func (c *VerdaClient) GetLocations(ctx context.Context, args v1.GetLocationsArgs) ([]v1.Location, error) {
	verdaLocations, err := c.client.Locations.Get(ctx)
	if err != nil {
		return nil, wrapVerdaError(err)
	}
	availabilities, err := c.client.InstanceAvailability.GetAllAvailabilities(ctx, false, "")
	if err != nil {
		return nil, wrapVerdaError(err)
	}

	available := make(map[string]bool, len(availabilities))
	for _, availability := range availabilities {
		available[availability.LocationCode] = len(availability.Availabilities) > 0
	}

	locations := make([]v1.Location, 0, len(verdaLocations))
	for _, verdaLocation := range verdaLocations {
		isAvailable := available[verdaLocation.Code]
		if !args.IncludeUnavailable && !isAvailable {
			continue
		}
		locations = append(locations, v1.Location{
			Name:        verdaLocation.Code,
			Description: verdaLocation.Name,
			Available:   isAvailable,
			Country:     countryAlpha3(verdaLocation.CountryCode),
		})
	}
	sort.Slice(locations, func(i, j int) bool {
		return locations[i].Name < locations[j].Name
	})
	return locations, nil
}

func countryAlpha3(alpha2 string) string {
	code := strings.ToUpper(alpha2)

	// The known current Verda locations
	switch code {
	case "FI":
		return "FIN"
	case "IS":
		return "ISL"
	case "NO":
		return "NOR"
	default:
		return code
	}
}
