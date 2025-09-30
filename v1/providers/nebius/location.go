package v1

import (
	"context"

	v1 "github.com/brevdev/cloud/v1"
)

// Common Nebius regions based on the projects we observed
const nebiusLocationsData = `[
    {"location_name": "eu-north1", "description": "Europe North 1 (Finland)", "country": "FIN"},
    {"location_name": "eu-west1", "description": "Europe West 1 (Netherlands)", "country": "NLD"},
    {"location_name": "us-central1", "description": "US Central 1 (Iowa)", "country": "USA"}
]`

// For now, support the current location pattern
func (c *NebiusClient) GetLocations(ctx context.Context, args v1.GetLocationsArgs) ([]v1.Location, error) {
	// Return the current configured location
	// In a full implementation, this would query the Nebius API for available regions
	location := v1.Location{
		Name:      c.location,
		Available: true,
	}

	// Add description based on known regions
	switch c.location {
	case "eu-north1":
		location.Description = "Europe North 1 (Finland)"
		location.Country = "FIN"
	case "eu-west1":
		location.Description = "Europe West 1 (Netherlands)"
		location.Country = "NLD"
	case "us-central1":
		location.Description = "US Central 1 (Iowa)"
		location.Country = "USA"
	default:
		location.Description = c.location
		location.Country = ""
	}

	return []v1.Location{location}, nil
}