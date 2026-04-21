package v1

import (
	"context"

	cloud "github.com/brevdev/cloud/v1"
)

type regionListResponse struct {
	Response regionList `json:"getRegionListResponse"`
}

func (r *regionListResponse) apiError() error {
	return r.Response.responseMeta.apiError()
}

type regionList struct {
	responseMeta
	TotalRows  int           `json:"totalRows"`
	RegionList []naverRegion `json:"regionList"`
}

type naverRegion struct {
	RegionNo   string `json:"regionNo"`
	RegionCode string `json:"regionCode"`
	RegionName string `json:"regionName"`
}

func (c *NaverClient) GetLocations(ctx context.Context, _ cloud.GetLocationsArgs) ([]cloud.Location, error) {
	var resp regionListResponse
	if err := c.do(ctx, "getRegionList", nil, &resp); err != nil {
		return nil, err
	}

	locations := make([]cloud.Location, 0, len(resp.Response.RegionList))
	for i, region := range resp.Response.RegionList {
		locations = append(locations, cloud.Location{
			Name:        region.RegionCode,
			Description: region.RegionName,
			Available:   true,
			Priority:    i,
			Country:     countryForRegion(region.RegionCode),
		})
	}
	return locations, nil
}

func countryForRegion(regionCode string) string {
	switch regionCode {
	case "KR":
		return "KOR"
	case "JPN":
		return "JPN"
	case "SGN":
		return "SGP"
	case "USWN", "USEN":
		return "USA"
	case "DEN":
		return "DEU"
	default:
		return ""
	}
}
