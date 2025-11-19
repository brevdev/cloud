package v1

import (
	"context"
	"fmt"

	v1 "github.com/brevdev/cloud/v1"
	quotas "github.com/nebius/gosdk/proto/nebius/quotas/v1"
)

// GetLocations returns all Nebius regions where the tenant has quota allocated
// This queries the actual Quotas API to discover regions with active quota
func (c *NebiusClient) GetLocations(ctx context.Context, args v1.GetLocationsArgs) ([]v1.Location, error) {
	// Query quota allocations to discover available regions
	quotaResp, err := c.sdk.Services().Quotas().V1().QuotaAllowance().List(ctx, &quotas.ListQuotaAllowancesRequest{
		ParentId: c.tenantID,
		PageSize: 1000, // Get all quotas
	})
	if err != nil {
		// Fallback to returning just the configured location if quota query fails
		return []v1.Location{{
			Name:        c.location,
			Description: getRegionDescription(c.location),
			Available:   true,
			Country:     getRegionCountry(c.location),
		}}, nil
	}

	// Extract unique regions from quota allocations
	regionMap := make(map[string]bool)
	for _, quota := range quotaResp.GetItems() {
		if quota.Spec == nil || quota.Status == nil {
			continue
		}

		// Only include regions with active quotas
		if quota.Status.State == quotas.QuotaAllowanceStatus_STATE_ACTIVE {
			region := quota.Spec.Region
			if region != "" {
				regionMap[region] = true
			}
		}
	}

	// Convert to location list
	var locations []v1.Location
	for region := range regionMap {
		// Only include available regions unless explicitly requested
		if !args.IncludeUnavailable && len(regionMap) == 0 {
			continue
		}

		locations = append(locations, v1.Location{
			Name:        region,
			Description: getRegionDescription(region),
			Available:   true, // If we have quota here, it's available
			Country:     getRegionCountry(region),
		})
	}

	// If no regions found from quota (shouldn't happen), return configured location
	if len(locations) == 0 {
		locations = []v1.Location{{
			Name:        c.location,
			Description: getRegionDescription(c.location),
			Available:   true,
			Country:     getRegionCountry(c.location),
		}}
	}

	return locations, nil
}

// getRegionDescription returns a human-readable description for a Nebius region
func getRegionDescription(region string) string {
	descriptions := map[string]string{
		"eu-north1":   "Europe North 1 (Finland)",
		"eu-west1":    "Europe West 1 (Netherlands)",
		"eu-west2":    "Europe West 2 (Belgium)",
		"eu-west3":    "Europe West 3 (Germany)",
		"eu-west4":    "Europe West 4 (France)",
		"us-central1": "US Central 1 (Iowa)",
		"us-east1":    "US East 1 (Virginia)",
		"us-west1":    "US West 1 (California)",
		"asia-east1":  "Asia East 1 (Taiwan)",
	}

	if desc, ok := descriptions[region]; ok {
		return desc
	}
	return fmt.Sprintf("Nebius Region %s", region)
}

// getRegionCountry returns the ISO 3166-1 alpha-3 country code for a region
func getRegionCountry(region string) string {
	countries := map[string]string{
		"eu-north1":   "FIN",
		"eu-west1":    "NLD",
		"eu-west2":    "BEL",
		"eu-west3":    "DEU",
		"eu-west4":    "FRA",
		"us-central1": "USA",
		"us-east1":    "USA",
		"us-west1":    "USA",
		"asia-east1":  "TWN",
	}

	if country, ok := countries[region]; ok {
		return country
	}
	return ""
}
