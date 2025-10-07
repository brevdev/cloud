package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/nebius/gosdk"
	"github.com/nebius/gosdk/auth"
	billing "github.com/nebius/gosdk/proto/nebius/billing/v1alpha1"
	compute "github.com/nebius/gosdk/proto/nebius/compute/v1"
	common "github.com/nebius/gosdk/proto/nebius/common/v1"
)

// PricingEstimate represents the cost estimate for an instance type
type PricingEstimate struct {
	PlatformID       string  `json:"platform_id"`
	PlatformName     string  `json:"platform_name"`
	PresetName       string  `json:"preset_name"`
	Region           string  `json:"region"`
	Currency         string  `json:"currency"`
	HourlyRate       float64 `json:"hourly_rate"`
	DailyRate        float64 `json:"daily_rate"`
	MonthlyRate      float64 `json:"monthly_rate"`
	AnnualRate       float64 `json:"annual_rate"`
}

func main() {
	ctx := context.Background()
	
	saJSON := os.Getenv("NEBIUS_SERVICE_ACCOUNT_JSON")
	tenantID := os.Getenv("NEBIUS_TENANT_ID")
	projectID := os.Getenv("NEBIUS_PROJECT_ID")
	
	if saJSON == "" || tenantID == "" {
		fmt.Fprintln(os.Stderr, "Error: Set NEBIUS_SERVICE_ACCOUNT_JSON, NEBIUS_TENANT_ID, and optionally NEBIUS_PROJECT_ID")
		os.Exit(1)
	}
	
	// Read service account
	saKey, err := os.ReadFile(saJSON)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading service account: %v\n", err)
		os.Exit(1)
	}
	
	// Initialize SDK
	var credFile auth.ServiceAccountCredentials
	if err := json.Unmarshal(saKey, &credFile); err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing service account: %v\n", err)
		os.Exit(1)
	}
	
	parser := auth.NewPrivateKeyParser(
		[]byte(credFile.SubjectCredentials.PrivateKey),
		credFile.SubjectCredentials.KeyID,
		credFile.SubjectCredentials.Subject,
	)
	creds := gosdk.ServiceAccountReader(parser)
	
	sdk, err := gosdk.New(ctx, gosdk.WithCredentials(creds))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error initializing SDK: %v\n", err)
		os.Exit(1)
	}
	
	// Default project ID if not provided
	if projectID == "" {
		projectID = fmt.Sprintf("project-integration-test")
	}
	
	// List all platforms to get pricing for each
	platformsResp, err := sdk.Services().Compute().V1().Platform().List(ctx, &compute.ListPlatformsRequest{
		ParentId: projectID,
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error listing platforms: %v\n", err)
		os.Exit(1)
	}
	
	var estimates []PricingEstimate
	
	// For each platform, estimate pricing for each preset
	for _, platform := range platformsResp.GetItems() {
		if platform.Metadata == nil || platform.Spec == nil {
			continue
		}
		
		for _, preset := range platform.Spec.Presets {
			if preset == nil {
				continue
			}
			
			// Estimate for first available region (eu-north1 as default)
			region := "eu-north1"
			
			estimate, err := estimatePlatformPresetPricing(ctx, sdk, projectID, platform.Metadata.Id, platform.Metadata.Name, preset.Name, region)
			if err != nil {
				// Skip on error, just log
				fmt.Fprintf(os.Stderr, "Warning: Could not estimate pricing for %s/%s: %v\n", platform.Metadata.Name, preset.Name, err)
				continue
			}
			
			estimates = append(estimates, *estimate)
		}
	}
	
	// Output as JSON
	output, err := json.MarshalIndent(estimates, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error marshaling JSON: %v\n", err)
		os.Exit(1)
	}
	
	fmt.Println(string(output))
}

func estimatePlatformPresetPricing(
	ctx context.Context,
	sdk *gosdk.SDK,
	projectID string,
	platformID string,
	platformName string,
	presetName string,
	region string,
) (*PricingEstimate, error) {
	// Build a minimal instance spec for pricing estimation
	// Only the platform and preset are required for pricing calculation
	req := &billing.EstimateRequest{
		ResourceSpec: &billing.ResourceSpec{
			ResourceSpec: &billing.ResourceSpec_ComputeInstanceSpec{
				ComputeInstanceSpec: &compute.CreateInstanceRequest{
					Metadata: &common.ResourceMetadata{
						ParentId: projectID,
						Name:     "pricing-estimate",
					},
					Spec: &compute.InstanceSpec{
						Resources: &compute.ResourcesSpec{
							Platform: platformName,
							Size: &compute.ResourcesSpec_Preset{
								Preset: presetName,
							},
						},
					},
				},
			},
		},
		// Use unspecified to get default/on-demand pricing
		OfferTypes: []billing.OfferType{
			billing.OfferType_OFFER_TYPE_UNSPECIFIED,
		},
	}
	
	resp, err := sdk.Services().Billing().V1Alpha1().Calculator().Estimate(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to estimate pricing: %w", err)
	}
	
	// Extract costs from nested structure
	var hourlyRate, monthlyRate float64
	
	if resp.HourlyCost != nil && resp.HourlyCost.GetGeneral() != nil && resp.HourlyCost.GetGeneral().Total != nil {
		hourlyRate = parseDecimalCost(resp.HourlyCost.GetGeneral().Total.Cost)
	}
	
	if resp.MonthlyCost != nil && resp.MonthlyCost.GetGeneral() != nil && resp.MonthlyCost.GetGeneral().Total != nil {
		monthlyRate = parseDecimalCost(resp.MonthlyCost.GetGeneral().Total.Cost)
	}
	
	// Calculate daily and annual from hourly and monthly
	dailyRate := hourlyRate * 24
	annualRate := monthlyRate * 12
	
	estimate := &PricingEstimate{
		PlatformID:   platformID,
		PlatformName: platformName,
		PresetName:   presetName,
		Region:       region,
		Currency:     "USD", // Nebius pricing currency
		HourlyRate:   hourlyRate,
		DailyRate:    dailyRate,
		MonthlyRate:  monthlyRate,
		AnnualRate:   annualRate,
	}
	
	return estimate, nil
}

// parseDecimalCost converts the decimal string cost to float64
func parseDecimalCost(costStr string) float64 {
	if costStr == "" {
		return 0.0
	}
	
	var cost float64
	fmt.Sscanf(costStr, "%f", &cost)
	return cost
}

