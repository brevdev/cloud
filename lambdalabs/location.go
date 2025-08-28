package lambdalabs

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/brevdev/cloud"
)

const lambdaLocationsData = `[
    {"location_name": "us-west-1", "description": "California, USA", "country": "USA"},
    {"location_name": "us-west-2", "description": "Arizona, USA", "country": "USA"},
    {"location_name": "us-west-3", "description": "Utah, USA", "country": "USA"},
    {"location_name": "us-south-1", "description": "Texas, USA", "country": "USA"},
    {"location_name": "us-east-1", "description": "Virginia, USA", "country": "USA"},
    {"location_name": "us-midwest-1", "description": "Illinois, USA", "country": "USA"},
    {"location_name": "australia-southeast-1", "description": "Australia", "country": "AUS"},
    {"location_name": "europe-central-1", "description": "Germany", "country": "DEU"},
    {"location_name": "asia-south-1", "description": "India", "country": "IND"},
    {"location_name": "me-west-1", "description": "Israel", "country": "ISR"},
    {"location_name": "europe-south-1", "description": "Italy", "country": "ITA"},
    {"location_name": "asia-northeast-1", "description": "Osaka, Japan", "country": "JPN"},
    {"location_name": "asia-northeast-2", "description": "Tokyo, Japan", "country": "JPN"},
    {"location_name": "us-east-3", "description": "Washington D.C, USA", "country": "USA"},
    {"location_name": "us-east-2", "description": "Washington D.C, USA", "country": "USA"},
    {"location_name": "australia-east-1", "description": "Sydney, Australia", "country": "AUS"},
    {"location_name": "us-south-3", "description": "Central Texas, USA", "country": "USA"},
    {"location_name": "us-south-2", "description": "North Texas, USA", "country": "USA"}
]`

type LambdaLocation struct {
	LocationName string `json:"location_name"`
	Description  string `json:"description"`
	Country      string `json:"country"`
}

func getLambdaLabsLocations() ([]LambdaLocation, error) {
	var locationData []LambdaLocation
	if err := json.Unmarshal([]byte(lambdaLocationsData), &locationData); err != nil {
		return nil, fmt.Errorf("failed to parse location data: %w", err)
	}
	return locationData, nil
}

func (c *LambdaLabsClient) GetLocations(_ context.Context, _ cloud.GetLocationsArgs) ([]cloud.Location, error) {
	locationData, err := getLambdaLabsLocations()
	if err != nil {
		return nil, err
	}

	locations := make([]cloud.Location, 0, len(locationData))
	for _, location := range locationData {
		locations = append(locations, cloud.Location{
			Name:        location.LocationName,
			Description: location.Description,
			Available:   true,
			Country:     location.Country,
		})
	}

	return locations, nil
}
