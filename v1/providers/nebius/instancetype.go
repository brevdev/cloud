package v1

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/alecthomas/units"
	"github.com/bojanz/currency"
	"github.com/brevdev/cloud/internal/errors"
	v1 "github.com/brevdev/cloud/v1"
	billing "github.com/nebius/gosdk/proto/nebius/billing/v1alpha1"
	common "github.com/nebius/gosdk/proto/nebius/common/v1"
	compute "github.com/nebius/gosdk/proto/nebius/compute/v1"
	quotas "github.com/nebius/gosdk/proto/nebius/quotas/v1"
)

func (c *NebiusClient) GetInstanceTypes(ctx context.Context, args v1.GetInstanceTypeArgs) ([]v1.InstanceType, error) {
	// Get platforms (instance types) from Nebius API
	platformsResp, err := c.sdk.Services().Compute().V1().Platform().List(ctx, &compute.ListPlatformsRequest{
		ParentId: c.projectID, // List platforms available in this project
	})
	if err != nil {
		return nil, errors.WrapAndTrace(err)
	}

	// Get all available locations for quota-aware enumeration
	// Default behavior: check ALL regions to show all available quota
	var locations []v1.Location

	if len(args.Locations) > 0 && !args.Locations.IsAll() {
		// User requested specific locations - filter to those
		allLocations, err := c.GetLocations(ctx, v1.GetLocationsArgs{})
		if err == nil {
			var filteredLocations []v1.Location
			for _, loc := range allLocations {
				for _, requestedLoc := range args.Locations {
					if loc.Name == requestedLoc {
						filteredLocations = append(filteredLocations, loc)
						break
					}
				}
			}
			locations = filteredLocations
		} else {
			// Fallback to client's configured location if we can't get all locations
			locations = []v1.Location{{Name: c.location}}
		}
	} else {
		// Default behavior: enumerate ALL regions for quota-aware discovery
		// This shows users all instance types they have quota for, regardless of region
		allLocations, err := c.GetLocations(ctx, v1.GetLocationsArgs{})
		if err == nil {
			locations = allLocations
		} else {
			// Fallback to client's configured location if we can't get all locations
			locations = []v1.Location{{Name: c.location}}
		}
	}

	// Get quota information for all regions
	quotaMap, err := c.getQuotaMap(ctx)
	if err != nil {
		// Log error but continue - we'll mark everything as unavailable
		quotaMap = make(map[string]*quotas.QuotaAllowance)
	}

	var instanceTypes []v1.InstanceType

	// For each location, get instance types with availability/quota info
	for _, location := range locations {
		locationInstanceTypes, err := c.getInstanceTypesForLocation(ctx, platformsResp, location, args, quotaMap)
		if err != nil {
			continue // Skip failed locations
		}
		instanceTypes = append(instanceTypes, locationInstanceTypes...)
	}

	// Apply filters
	instanceTypes = c.applyInstanceTypeFilters(instanceTypes, args)

	return instanceTypes, nil
}

func (c *NebiusClient) GetInstanceTypePollTime() time.Duration {
	return 5 * time.Minute
}

func (c *NebiusClient) MergeInstanceTypeForUpdate(currIt v1.InstanceType, newIt v1.InstanceType) v1.InstanceType {
	merged := newIt
	merged.ID = currIt.ID
	return merged
}

func (c *NebiusClient) GetInstanceTypeQuotas(ctx context.Context, args v1.GetInstanceTypeQuotasArgs) (v1.Quota, error) {
	// Query actual Nebius quotas from the compute service
	// For now, return a default quota structure
	quota := v1.Quota{
		ID:      "nebius-compute-quota",
		Name:    "Nebius Compute Quota",
		Maximum: 1000, // Default maximum instances - should be queried from API
		Current: 0,    // Would be calculated from actual usage
		Unit:    "instances",
	}

	return quota, nil
}

// getInstanceTypesForLocation gets instance types for a specific location with quota/availability checking
//
//nolint:gocognit // Complex function iterating platforms, presets, and quota checks
func (c *NebiusClient) getInstanceTypesForLocation(ctx context.Context, platformsResp *compute.ListPlatformsResponse, location v1.Location, args v1.GetInstanceTypeArgs, quotaMap map[string]*quotas.QuotaAllowance) ([]v1.InstanceType, error) {
	var instanceTypes []v1.InstanceType

	for _, platform := range platformsResp.GetItems() {
		if platform.Metadata == nil || platform.Spec == nil {
			continue
		}

		// Filter platforms to only supported ones
		if !c.isPlatformSupported(platform.Metadata.Name) {
			continue
		}

		// Check if this is a CPU-only platform
		isCPUOnly := c.isCPUOnlyPlatform(platform.Metadata.Name)

		// For CPU platforms, limit the number of presets to avoid pollution
		maxCPUPresets := 3
		cpuPresetCount := 0

		// For each preset, create an instance type
		for _, preset := range platform.Spec.Presets {
			if preset == nil || preset.Resources == nil {
				continue
			}

			// For CPU platforms, limit to first N presets
			if isCPUOnly {
				if cpuPresetCount >= maxCPUPresets {
					continue
				}
			}

			// Determine GPU type and details from platform name
			gpuType, gpuName := extractGPUTypeAndName(platform.Metadata.Name)

			// Check quota/availability for this instance type in this location
			isAvailable := c.checkPresetQuotaAvailability(preset.Resources, location.Name, platform.Metadata.Name, quotaMap)

			// Skip instance types with no quota at all
			if !isAvailable {
				continue
			}

			// Increment CPU preset counter if this is a CPU platform
			if isCPUOnly {
				cpuPresetCount++
			}

			// Build instance type ID in dot-separated format: {platform}.{preset}
			// Examples:
			//   gpu-l40s.4gpu-96vcpu-768gb
			//   gpu-h100-sxm.8gpu-128vcpu-1600gb
			//   cpu-e2.4vcpu-16gb
			// ID and Type are the same - no region/provider prefix
			instanceTypeID := fmt.Sprintf("%s.%s", platform.Metadata.Name, preset.Name)

			c.logger.Debug(ctx, "building instance type",
				v1.LogField("instanceTypeID", instanceTypeID),
				v1.LogField("platformName", platform.Metadata.Name),
				v1.LogField("presetName", preset.Name),
				v1.LogField("location", location.Name),
				v1.LogField("gpuType", gpuType))

			// Convert Nebius platform preset to our InstanceType format
			instanceType := v1.InstanceType{
				ID:                 v1.InstanceTypeID(instanceTypeID), // Dot-separated format (e.g., "gpu-h100-sxm.8gpu-128vcpu-1600gb")
				Location:           location.Name,
				Type:               instanceTypeID, // Same as ID - both use dot-separated format
				VCPU:               preset.Resources.VcpuCount,
				MemoryBytes:        v1.NewBytes(v1.BytesValue(preset.Resources.MemoryGibibytes), v1.Gibibyte), // Memory in GiB
				NetworkPerformance: "standard", // Default network performance
				IsAvailable:        isAvailable,
				Stoppable:          true, // All Nebius instances support stop/start operations
				ElasticRootVolume:  true, // Nebius supports dynamic disk allocation
				SupportedStorage:   c.buildSupportedStorage(),
				Provider:           CloudProviderID, // Nebius is the provider
			}

			// Add GPU information if available
			if preset.Resources.GpuCount > 0 && !isCPUOnly {
				gpu := v1.GPU{
					Count:        preset.Resources.GpuCount,
					Type:         gpuType,
					Name:         gpuName,
					Manufacturer: v1.ManufacturerNVIDIA, // Nebius currently only supports NVIDIA GPUs
					Memory:       getGPUMemory(gpuType), // Populate VRAM based on GPU type
				}
				instanceType.SupportedGPUs = []v1.GPU{gpu}
			}

			// Enrich with pricing information from Nebius Billing API
			pricing := c.getPricingForInstanceType(ctx, platform.Metadata.Name, preset.Name, location.Name)
			if pricing != nil {
				instanceType.BasePrice = pricing
			}

			instanceTypes = append(instanceTypes, instanceType)
		}
	}

	return instanceTypes, nil
}

// getQuotaMap retrieves all quota allowances for the tenant and creates a lookup map
func (c *NebiusClient) getQuotaMap(ctx context.Context) (map[string]*quotas.QuotaAllowance, error) {
	quotaMap := make(map[string]*quotas.QuotaAllowance)

	// List all quota allowances for the tenant
	resp, err := c.sdk.Services().Quotas().V1().QuotaAllowance().List(ctx, &quotas.ListQuotaAllowancesRequest{
		ParentId: c.tenantID, // Use tenant ID to list all quotas
		PageSize: 1000,       // Get all quotas in one request
	})
	if err != nil {
		return nil, errors.WrapAndTrace(err)
	}

	// Build a map of quota name + region -> quota allowance
	for _, quota := range resp.GetItems() {
		if quota.Metadata == nil || quota.Spec == nil || quota.Status == nil {
			continue
		}

		// Only include active quotas with available capacity
		if quota.Status.State != quotas.QuotaAllowanceStatus_STATE_ACTIVE {
			continue
		}

		// Key format: "quota-name:region" (e.g., "compute.instance.gpu.h100:eu-north1")
		key := fmt.Sprintf("%s:%s", quota.Metadata.Name, quota.Spec.Region)
		quotaMap[key] = quota
	}

	return quotaMap, nil
}

// checkPresetQuotaAvailability checks if a preset has available quota in the specified region
func (c *NebiusClient) checkPresetQuotaAvailability(resources *compute.PresetResources, region string, platformName string, quotaMap map[string]*quotas.QuotaAllowance) bool {
	// Check GPU quota if GPUs are requested
	if resources.GpuCount > 0 {
		// Determine GPU type from platform name
		gpuQuotaName := c.getGPUQuotaName(platformName)
		if gpuQuotaName == "" {
			return false // Unknown GPU type
		}

		key := fmt.Sprintf("%s:%s", gpuQuotaName, region)
		quota, exists := quotaMap[key]
		if !exists {
			return false // No quota for this GPU in this region
		}

		// Check if quota has available capacity
		if quota.Status == nil || quota.Spec == nil || quota.Spec.Limit == nil {
			return false
		}

		available := int64(*quota.Spec.Limit) - int64(quota.Status.Usage)
		if available < int64(resources.GpuCount) {
			return false // Not enough GPU quota
		}

		return true
	}

	// For CPU-only instances, check CPU and memory quotas
	// Nebius uses "compute.instance.non-gpu.vcpu" for CPU quota (not "compute.cpu")
	cpuQuotaKey := fmt.Sprintf("compute.instance.non-gpu.vcpu:%s", region)
	if cpuQuota, exists := quotaMap[cpuQuotaKey]; exists {
		if cpuQuota.Status != nil && cpuQuota.Spec != nil && cpuQuota.Spec.Limit != nil {
			cpuAvailable := int64(*cpuQuota.Spec.Limit) - int64(cpuQuota.Status.Usage)
			if cpuAvailable < int64(resources.VcpuCount) {
				return false
			}
		}
	}

	// Check memory quota - Nebius uses "compute.instance.non-gpu.memory"
	memoryQuotaKey := fmt.Sprintf("compute.instance.non-gpu.memory:%s", region)
	if memQuota, exists := quotaMap[memoryQuotaKey]; exists {
		if memQuota.Status != nil && memQuota.Spec != nil && memQuota.Spec.Limit != nil {
			memoryRequired := int64(resources.MemoryGibibytes) * 1024 * 1024 * 1024 // Convert GiB to bytes
			memAvailable := int64(*memQuota.Spec.Limit) - int64(memQuota.Status.Usage)
			if memAvailable < memoryRequired {
				return false
			}
		}
	}

	return true // CPU-only instances are available if we get here
}

// getGPUQuotaName determines the quota name for a GPU based on the platform name
func (c *NebiusClient) getGPUQuotaName(platformName string) string {
	// Nebius GPU quota names follow pattern: "compute.instance.gpu.{type}"
	// Examples: "compute.instance.gpu.h100", "compute.instance.gpu.h200", "compute.instance.gpu.l40s"

	platformLower := strings.ToLower(platformName)

	if strings.Contains(platformLower, "h100") {
		return "compute.instance.gpu.h100"
	}
	if strings.Contains(platformLower, "h200") {
		return "compute.instance.gpu.h200"
	}
	if strings.Contains(platformLower, "l40s") {
		return "compute.instance.gpu.l40s"
	}
	if strings.Contains(platformLower, "a100") {
		return "compute.instance.gpu.a100"
	}
	if strings.Contains(platformLower, "v100") {
		return "compute.instance.gpu.v100"
	}
	if strings.Contains(platformLower, "b200") {
		return "compute.instance.gpu.b200"
	}

	return ""
}

// isPlatformSupported checks if a platform should be included in instance types
func (c *NebiusClient) isPlatformSupported(platformName string) bool {
	platformLower := strings.ToLower(platformName)

	// For GPU platforms: accept any GPU platform (filtered by quota availability)
	// Look for common GPU indicators in platform names
	gpuIndicators := []string{"gpu", "h100", "h200", "l40s", "a100", "v100", "a10", "t4", "l4", "b200"}
	for _, indicator := range gpuIndicators {
		if strings.Contains(platformLower, indicator) {
			return true
		}
	}

	// For CPU platforms: only accept specific types to avoid polluting the list
	if strings.Contains(platformLower, "cpu-d3") || strings.Contains(platformLower, "cpu-e2") {
		return true
	}

	return false
}

// isCPUOnlyPlatform checks if a platform is CPU-only (no GPUs)
func (c *NebiusClient) isCPUOnlyPlatform(platformName string) bool {
	platformLower := strings.ToLower(platformName)
	return strings.Contains(platformLower, "cpu-d3") || strings.Contains(platformLower, "cpu-e2")
}

// buildSupportedStorage creates storage configuration for Nebius instances
func (c *NebiusClient) buildSupportedStorage() []v1.Storage {
	// Nebius supports dynamically allocatable network SSD disks
	// Minimum: 50GB, Maximum: 2560GB
	minSize := units.Base2Bytes(50 * units.GiB)
	maxSize := units.Base2Bytes(2560 * units.GiB)

	// Pricing is roughly $0.10 per GB-month, which is ~$0.00014 per GB-hour
	pricePerGBHr, _ := currency.NewAmount("0.00014", "USD")

	return []v1.Storage{
		{
			Type:         "network-ssd",
			Count:        1,
			MinSize:      &minSize,
			MaxSize:      &maxSize,
			IsElastic:    true,
			PricePerGBHr: &pricePerGBHr,
		},
	}
}

// applyInstanceTypeFilters applies various filters to the instance type list
func (c *NebiusClient) applyInstanceTypeFilters(instanceTypes []v1.InstanceType, args v1.GetInstanceTypeArgs) []v1.InstanceType {
	var filtered []v1.InstanceType

	for _, instanceType := range instanceTypes {
		// Apply specific instance type filters
		if len(args.InstanceTypes) > 0 {
			found := false
			for _, requestedType := range args.InstanceTypes {
				if string(instanceType.ID) == requestedType {
					found = true
					break
				}
			}
			if !found {
				continue
			}
		}

		// Apply architecture filter
		if args.ArchitectureFilter != nil {
			arch := determineInstanceTypeArchitecture(instanceType)
			// Check if architecture matches the filter requirements
			if len(args.ArchitectureFilter.IncludeArchitectures) > 0 {
				found := false
				for _, allowedArch := range args.ArchitectureFilter.IncludeArchitectures {
					if arch == string(allowedArch) {
						found = true
						break
					}
				}
				if !found {
					continue
				}
			}
		}

		filtered = append(filtered, instanceType)
	}

	return filtered
}

// extractGPUTypeAndName extracts GPU type and name from platform name
// Note: Returns model name only (e.g., "H100"), not full name with manufacturer
// Manufacturer info is stored separately in GPU.Manufacturer field
func extractGPUTypeAndName(platformName string) (string, string) {
	platformLower := strings.ToLower(platformName)

	if strings.Contains(platformLower, "h100") {
		return "H100", "H100"
	}
	if strings.Contains(platformLower, "h200") {
		return "H200", "H200"
	}
	if strings.Contains(platformLower, "l40s") {
		return "L40S", "L40S"
	}
	if strings.Contains(platformLower, "a100") {
		return "A100", "A100"
	}
	if strings.Contains(platformLower, "v100") {
		return "V100", "V100"
	}
	if strings.Contains(platformLower, "b200") {
		return "B200", "B200"
	}

	return "GPU", "GPU" // Generic fallback
}

// getGPUMemory returns the VRAM for a given GPU type in GiB
func getGPUMemory(gpuType string) units.Base2Bytes {
	// Static mapping of GPU types to their VRAM capacities
	vramMap := map[string]int64{
		"L40S": 48,  // 48 GiB VRAM
		"H100": 80,  // 80 GiB VRAM
		"H200": 141, // 141 GiB VRAM
		"A100": 80,  // 80 GiB VRAM (most common variant)
		"V100": 32,  // 32 GiB VRAM (most common variant)
		"A10":  24,  // 24 GiB VRAM
		"T4":   16,  // 16 GiB VRAM
		"L4":   24,  // 24 GiB VRAM
		"B200": 192, // 192 GiB VRAM
	}

	if vramGiB, exists := vramMap[gpuType]; exists {
		return units.Base2Bytes(vramGiB * int64(units.Gibibyte))
	}

	// Default fallback for unknown GPU types
	return units.Base2Bytes(0)
}

// determineInstanceTypeArchitecture determines architecture from instance type
func determineInstanceTypeArchitecture(instanceType v1.InstanceType) string {
	// Check if ARM architecture is indicated in the type or name
	typeLower := strings.ToLower(instanceType.Type)
	if strings.Contains(typeLower, "arm") || strings.Contains(typeLower, "aarch64") {
		return "arm64"
	}

	return "x86_64" // Default assumption
}

// getPricingForInstanceType fetches real pricing from Nebius Billing Calculator API
// Returns nil if pricing cannot be fetched (non-critical failure)
func (c *NebiusClient) getPricingForInstanceType(ctx context.Context, platformName, presetName, region string) *currency.Amount {
	// Build minimal instance spec for pricing estimation
	req := &billing.EstimateRequest{
		ResourceSpec: &billing.ResourceSpec{
			ResourceSpec: &billing.ResourceSpec_ComputeInstanceSpec{
				ComputeInstanceSpec: &compute.CreateInstanceRequest{
					Metadata: &common.ResourceMetadata{
						ParentId: c.projectID,
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
		OfferTypes: []billing.OfferType{
			billing.OfferType_OFFER_TYPE_UNSPECIFIED, // On-demand pricing
		},
	}

	// Query Nebius Billing Calculator API
	resp, err := c.sdk.Services().Billing().V1Alpha1().Calculator().Estimate(ctx, req)
	if err != nil {
		// Non-critical failure - pricing is optional enrichment
		// Log error but don't fail the entire GetInstanceTypes call
		return nil
	}

	// Extract hourly cost
	if resp.HourlyCost == nil || resp.HourlyCost.GetGeneral() == nil || resp.HourlyCost.GetGeneral().Total == nil {
		return nil
	}

	costStr := resp.HourlyCost.GetGeneral().Total.Cost
	if costStr == "" {
		return nil
	}

	// Parse cost string to currency.Amount
	amount, err := currency.NewAmount(costStr, "USD")
	if err != nil {
		return nil
	}

	return &amount
}
