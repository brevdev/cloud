package hyperstack

import (
	"context"
	"fmt"
	"sort"
	"strings"

	"github.com/NexGenCloud/hyperstack-sdk-go/lib/flavor"
	"github.com/NexGenCloud/hyperstack-sdk-go/lib/region"

	v1 "github.com/brevdev/cloud/v1"
)

func (c *HyperstackClient) GetLocations(ctx context.Context, args v1.GetLocationsArgs) ([]v1.Location, error) {
	response, err := c.regions.ListRegionsWithResponse(ctx)
	if err != nil {
		return nil, wrapTransportError("list regions", err)
	}
	if response.StatusCode() != 200 {
		return nil, responseError("list regions", response.StatusCode(), response.Body, nil)
	}
	if response.JSON200 == nil || response.JSON200.Regions == nil {
		return nil, fmt.Errorf("hyperstack list regions response did not contain data")
	}
	regions := *response.JSON200.Regions

	allFlavors, err := c.listFlavors(ctx)
	if err != nil {
		return nil, err
	}
	regionsWithAvailability := availableFlavorLocations(allFlavors)

	locations := make([]v1.Location, 0, len(regions))
	for _, providerRegion := range regions {
		location, ok := hyperstackLocation(providerRegion, regionsWithAvailability)
		if !ok {
			continue
		}
		if !args.IncludeUnavailable && !location.Available {
			continue
		}
		locations = append(locations, location)
	}
	sort.Slice(locations, func(i, j int) bool {
		return locations[i].Name < locations[j].Name
	})
	return locations, nil
}

func availableFlavorLocations(groups []flavor.FlavorItemGetResponse) map[string]bool {
	available := make(map[string]bool)
	for _, group := range groups {
		if group.Flavors == nil {
			continue
		}
		for _, providerType := range *group.Flavors {
			location := stringValue(group.RegionName)
			if location == "" {
				location = stringValue(providerType.RegionName)
			}
			if providerType.StockAvailable == nil || *providerType.StockAvailable {
				available[location] = true
			}
		}
	}
	return available
}

func hyperstackLocation(providerRegion region.RegionFields, regionsWithAvailability map[string]bool) (v1.Location, bool) {
	regionName := strings.TrimSpace(stringValue(providerRegion.Name))
	if regionName == "" {
		return v1.Location{}, false
	}
	description := strings.TrimSpace(stringValue(providerRegion.Description))
	if description == "" {
		description = regionName
	}
	return v1.Location{
		Name:        regionName,
		Description: description,
		Available:   regionsWithAvailability[regionName] && supportsFloatingIP(providerRegion.Features),
		Country:     countryAlpha3(stringValue(providerRegion.Country)),
	}, true
}

// SupportsFloatingIP effectively checks if the region supports public IP addresses
// See: https://docs.hyperstack.cloud/docs/api-reference/floating-ip/
func supportsFloatingIP(features *map[string]interface{}) bool {
	if features == nil {
		return true
	}
	value, found := (*features)["floating_ip"]
	if !found {
		return true
	}
	supported, ok := value.(bool)
	return ok && supported
}

func countryAlpha3(alpha2 string) string {
	switch strings.ToUpper(strings.TrimSpace(alpha2)) {
	case "CA":
		return "CAN"
	case "NO":
		return "NOR"
	case "US":
		return "USA"
	default:
		return strings.ToUpper(strings.TrimSpace(alpha2))
	}
}
