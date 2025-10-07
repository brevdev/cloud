package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strings"

	v1 "github.com/brevdev/cloud/v1"
	nebius "github.com/brevdev/cloud/v1/providers/nebius"
	"github.com/nebius/gosdk"
	"github.com/nebius/gosdk/auth"
	billing "github.com/nebius/gosdk/proto/nebius/billing/v1alpha1"
	common "github.com/nebius/gosdk/proto/nebius/common/v1"
	compute "github.com/nebius/gosdk/proto/nebius/compute/v1"
)

// AggregatedInstanceType matches the LaunchPad format with regional capacity
type AggregatedInstanceType struct {
	// Semantic identifier for this instance type configuration
	// Format: {platform}-{preset} (e.g., "gpu-h200-sxm-8gpu-128vcpu-1600gb")
	ID string `json:"id"`

	// Cloud provider
	Cloud string `json:"cloud"`

	// Platform name (e.g., "gpu-l40s-d", "cpu-d3")
	Platform string `json:"platform"`

	// Preset name (e.g., "1gpu-16vcpu-96gb")
	Preset string `json:"preset"`

	// Nebius internal platform ID (includes routing code like e00)
	// Kept for reference but not used as primary ID
	NebiusPlatformID string `json:"nebius_platform_id,omitempty"`

	// Key/value pairs of region name and availability (0 or 1 for Nebius quota-based)
	Capacity map[string]int `json:"capacity"`

	// List of regions where this instance type is available
	Regions []string `json:"regions"`

	// Resources
	CPU      int32 `json:"cpu"`
	MemoryGB int   `json:"memory_gb"`

	// GPU information (if applicable)
	GPU *GPUInfo `json:"gpu,omitempty"`

	// Storage
	Storage []StorageInfo `json:"storage"`

	// Architecture
	SystemArch string `json:"system_arch"`

	// Pricing (from Nebius billing API if available)
	Price PriceInfo `json:"price"`
}

type GPUInfo struct {
	Count               int    `json:"count"`
	Family              string `json:"family"`                         // e.g., "l40s", "h100"
	Model               string `json:"model"`                          // e.g., "L40S-48GB", "H100-80GB"
	Manufacturer        string `json:"manufacturer"`                   // "NVIDIA"
	MemoryGB            int    `json:"memory_gb,omitempty"`            // GPU memory
	InterconnectionType string `json:"interconnection_type,omitempty"` // "nvlink", "pcie"
}

type StorageInfo struct {
	Type      string `json:"type"`        // "network-ssd"
	SizeMinGB int    `json:"size_min_gb"` // Minimum size
	SizeMaxGB int    `json:"size_max_gb"` // Maximum size
	IsElastic bool   `json:"is_elastic"`  // Can be resized
}

type PriceInfo struct {
	Currency         string  `json:"currency"`
	OnDemandPerHour  float64 `json:"on_demand_per_hour"`
	EstimatedMonthly float64 `json:"estimated_monthly,omitempty"`
}

func main() {
	ctx := context.Background()

	// Read credentials from environment
	saJSON := os.Getenv("NEBIUS_SERVICE_ACCOUNT_JSON")
	tenantID := os.Getenv("NEBIUS_TENANT_ID")
	location := os.Getenv("NEBIUS_LOCATION")
	fetchPricing := os.Getenv("FETCH_PRICING") == "true"

	if saJSON == "" || tenantID == "" {
		fmt.Fprintln(os.Stderr, "Error: Set NEBIUS_SERVICE_ACCOUNT_JSON and NEBIUS_TENANT_ID")
		os.Exit(1)
	}

	if location == "" {
		location = "eu-north1" // Default location
	}

	// Read service account JSON
	saKey, err := os.ReadFile(saJSON)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading service account: %v\n", err)
		os.Exit(1)
	}

	// Create client (it will create/find a project automatically)
	cred := nebius.NewNebiusCredential("integration-test", string(saKey), tenantID)
	client, err := cred.MakeClient(ctx, location)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating client: %v\n", err)
		os.Exit(1)
	}

	// Get all instance types (across all regions)
	instanceTypes, err := client.GetInstanceTypes(ctx, v1.GetInstanceTypeArgs{})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error getting instance types: %v\n", err)
		os.Exit(1)
	}

	// Aggregate by preset configuration
	aggregated := aggregateInstanceTypes(instanceTypes)

	// Optionally fetch pricing (can be slow, so make it opt-in via FETCH_PRICING=true)
	if fetchPricing {
		fmt.Fprintln(os.Stderr, "Fetching pricing information from Nebius Billing API...")
		fmt.Fprintf(os.Stderr, "This may take 30-60 seconds for %d instance types...\n", len(aggregated))

		// Get project ID from client for billing API (just needs any valid project for pricing catalog)
		projectID, err := client.GetTenantID()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Warning: Could not get project ID for pricing: %v\n", err)
			fmt.Fprintln(os.Stderr, "Continuing with placeholder pricing...")
		} else {
			// We need to recreate the SDK since the client doesn't expose it
			if err := enrichWithRealPricing(ctx, string(saKey), projectID, aggregated); err != nil {
				fmt.Fprintf(os.Stderr, "Warning: Could not fetch pricing: %v\n", err)
				fmt.Fprintln(os.Stderr, "Continuing with placeholder pricing...")
			} else {
				fmt.Fprintln(os.Stderr, "✅ Pricing data successfully retrieved from Nebius Billing API")
			}
		}
	} else {
		fmt.Fprintln(os.Stderr, "Note: Using placeholder pricing. Set FETCH_PRICING=true to query real pricing from Nebius Billing API")
	}

	// Sort by ID for consistent output
	sort.Slice(aggregated, func(i, j int) bool {
		return aggregated[i].ID < aggregated[j].ID
	})

	// Output as JSON
	output, err := json.MarshalIndent(aggregated, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error marshaling JSON: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(string(output))
}

// enrichWithRealPricing fetches real pricing from Nebius Billing Calculator API
func enrichWithRealPricing(ctx context.Context, serviceAccountKey string, projectID string, aggregated []AggregatedInstanceType) error {
	// Initialize SDK for billing API access
	var credFile auth.ServiceAccountCredentials
	if err := json.Unmarshal([]byte(serviceAccountKey), &credFile); err != nil {
		return fmt.Errorf("failed to parse service account: %w", err)
	}

	parser := auth.NewPrivateKeyParser(
		[]byte(credFile.SubjectCredentials.PrivateKey),
		credFile.SubjectCredentials.KeyID,
		credFile.SubjectCredentials.Subject,
	)
	creds := gosdk.ServiceAccountReader(parser)

	sdk, err := gosdk.New(ctx, gosdk.WithCredentials(creds))
	if err != nil {
		return fmt.Errorf("failed to initialize SDK: %w", err)
	}

	// Fetch pricing for each instance type
	for i := range aggregated {
		if len(aggregated[i].Regions) == 0 {
			continue
		}

		// Build estimate request with minimal spec
		// Pricing is catalog-level and doesn't vary by region for the same preset
		req := &billing.EstimateRequest{
			ResourceSpec: &billing.ResourceSpec{
				ResourceSpec: &billing.ResourceSpec_ComputeInstanceSpec{
					ComputeInstanceSpec: &compute.CreateInstanceRequest{
						Metadata: &common.ResourceMetadata{
							ParentId: projectID,
							Name:     fmt.Sprintf("pricing-%s", aggregated[i].Platform),
						},
						Spec: &compute.InstanceSpec{
							Resources: &compute.ResourcesSpec{
								Platform: aggregated[i].Platform, // Use semantic platform name
								Size: &compute.ResourcesSpec_Preset{
									Preset: aggregated[i].Preset,
								},
							},
						},
					},
				},
			},
			OfferTypes: []billing.OfferType{
				billing.OfferType_OFFER_TYPE_UNSPECIFIED, // On-demand pricing
			},
		}

		resp, err := sdk.Services().Billing().V1Alpha1().Calculator().Estimate(ctx, req)
		if err != nil {
			// Log warning but continue
			fmt.Fprintf(os.Stderr, "  Warning: Could not get pricing for %s/%s: %v\n", aggregated[i].Platform, aggregated[i].Preset, err)
			continue
		}

		// Extract hourly and monthly costs
		var hourlyRate, monthlyRate float64

		if resp.HourlyCost != nil && resp.HourlyCost.GetGeneral() != nil && resp.HourlyCost.GetGeneral().Total != nil {
			hourlyRate = parseDecimalCost(resp.HourlyCost.GetGeneral().Total.Cost)
		}

		if resp.MonthlyCost != nil && resp.MonthlyCost.GetGeneral() != nil && resp.MonthlyCost.GetGeneral().Total != nil {
			monthlyRate = parseDecimalCost(resp.MonthlyCost.GetGeneral().Total.Cost)
		}

		// Update the aggregated entry with real pricing
		aggregated[i].Price.OnDemandPerHour = hourlyRate
		aggregated[i].Price.EstimatedMonthly = monthlyRate
	}

	return nil
}

// parseDecimalCost converts Nebius decimal string cost to float64
func parseDecimalCost(costStr string) float64 {
	if costStr == "" {
		return 0.0
	}

	var cost float64
	fmt.Sscanf(costStr, "%f", &cost)
	return cost
}

// aggregateInstanceTypes aggregates v1.InstanceType entries by preset configuration
// Returns one entry per preset with regional capacity information
func aggregateInstanceTypes(instanceTypes []v1.InstanceType) []AggregatedInstanceType {
	// Group by semantic ID (platform + preset, not Nebius internal ID)
	groups := make(map[string]*AggregatedInstanceType)

	for _, it := range instanceTypes {
		// Extract platform and preset from the Type field
		platform, preset := extractPlatformAndPreset(it.Type)

		// Generate semantic ID: {platform}-{preset}
		// This is stable across regions and routing codes
		semanticID := fmt.Sprintf("%s-%s", platform, preset)

		// Extract the Nebius internal platform ID (for reference)
		nebiusPlatformID := extractNebiusPlatformID(string(it.ID))

		if existing, ok := groups[semanticID]; ok {
			// Add this region to the existing entry
			existing.Regions = append(existing.Regions, it.Location)
			if it.IsAvailable {
				existing.Capacity[it.Location] = 1
			} else {
				existing.Capacity[it.Location] = 0
			}
		} else {
			// Create new aggregated entry
			agg := &AggregatedInstanceType{
				ID:               semanticID,
				Cloud:            "nebius",
				Platform:         platform,
				Preset:           preset,
				NebiusPlatformID: nebiusPlatformID,
				Capacity:         make(map[string]int),
				Regions:          []string{it.Location},
				CPU:              it.VCPU,
				MemoryGB:         int(it.Memory / (1024 * 1024 * 1024)),
				SystemArch:       determineArch(it),
				Storage:          convertStorage(it.SupportedStorage),
				Price: PriceInfo{
					Currency:        "USD",
					OnDemandPerHour: 0.0, // Will be populated if FETCH_PRICING=true
				},
			}

			if it.IsAvailable {
				agg.Capacity[it.Location] = 1
			} else {
				agg.Capacity[it.Location] = 0
			}

			// Add GPU info if present
			if len(it.SupportedGPUs) > 0 {
				gpu := it.SupportedGPUs[0]
				agg.GPU = &GPUInfo{
					Count:               int(gpu.Count),
					Family:              strings.ToLower(gpu.Type),
					Model:               gpu.Name,
					Manufacturer:        string(gpu.Manufacturer),
					MemoryGB:            int(gpu.Memory / (1024 * 1024 * 1024)), // Convert bytes to GB
					InterconnectionType: gpu.NetworkDetails,
				}
			}

			groups[semanticID] = agg
		}
	}

	// Convert map to slice
	result := make([]AggregatedInstanceType, 0, len(groups))
	for _, agg := range groups {
		// Sort regions for consistent output
		sort.Strings(agg.Regions)
		result = append(result, *agg)
	}

	return result
}

func extractPlatformAndPreset(typeStr string) (platform, preset string) {
	// Type format: "gpu-l40s-d (1gpu-16vcpu-96gb)" or "cpu-d3 (4vcpu-16gb)"
	parts := strings.Split(typeStr, " (")
	if len(parts) == 2 {
		platform = parts[0]
		preset = strings.TrimSuffix(parts[1], ")")
		return
	}
	return typeStr, ""
}

func extractNebiusPlatformID(fullID string) string {
	// Full ID format: "computeplatform-e00xxx-preset-name"
	// Extract just the platform part: "computeplatform-e00xxx"
	parts := strings.SplitN(fullID, "-", 3) // Split into max 3 parts
	if len(parts) >= 2 {
		// Return "computeplatform-e00xxx"
		return strings.Join(parts[0:2], "-")
	}
	return fullID
}

func determineArch(it v1.InstanceType) string {
	if len(it.SupportedArchitectures) > 0 {
		return string(it.SupportedArchitectures[0])
	}
	return "amd64" // Default
}

func convertStorage(storage []v1.Storage) []StorageInfo {
	result := make([]StorageInfo, 0, len(storage))
	for _, s := range storage {
		info := StorageInfo{
			Type:      s.Type,
			IsElastic: s.IsElastic,
		}
		if s.MinSize != nil {
			info.SizeMinGB = int(*s.MinSize / (1024 * 1024 * 1024))
		}
		if s.MaxSize != nil {
			info.SizeMaxGB = int(*s.MaxSize / (1024 * 1024 * 1024))
		}
		result = append(result, info)
	}
	return result
}
