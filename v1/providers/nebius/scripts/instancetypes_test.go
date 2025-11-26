//go:build scripts
// +build scripts

package scripts

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strings"
	"testing"

	v1 "github.com/brevdev/cloud/v1"
	nebius "github.com/brevdev/cloud/v1/providers/nebius"
)

// Test_EnumerateInstanceTypes enumerates all instance types across all Nebius regions
// Usage:
//
//	export NEBIUS_SERVICE_ACCOUNT_JSON='/path/to/service-account.json'
//	export NEBIUS_TENANT_ID='tenant-e00xxx'
//	go test -tags scripts -v -run Test_EnumerateInstanceTypes
func Test_EnumerateInstanceTypes(t *testing.T) {
	serviceAccountJSON := os.Getenv("NEBIUS_SERVICE_ACCOUNT_JSON")
	tenantID := os.Getenv("NEBIUS_TENANT_ID")

	if serviceAccountJSON == "" || tenantID == "" {
		t.Skip("NEBIUS_SERVICE_ACCOUNT_JSON and NEBIUS_TENANT_ID must be set")
	}

	ctx := context.Background()

	// List of regions to enumerate
	regions := []string{
		"eu-north1",
		"eu-west1",
		"eu-west2",
		"us-central1",
		"us-east1",
		"asia-east1",
	}

	t.Logf("Enumerating instance types across %d regions...", len(regions))

	allInstanceTypes := make([]v1.InstanceType, 0)
	regionStats := make(map[string]int)
	gpuFamilies := make(map[string]map[string]int) // region -> gpu_family -> count

	for _, region := range regions {
		t.Logf("Querying region: %s...", region)

		// Create client for this region
		client, err := nebius.NewNebiusClient(ctx, "enum-script", serviceAccountJSON, tenantID, "", region)
		if err != nil {
			t.Logf("  Warning: Failed to create client for %s: %v", region, err)
			continue
		}

		// Get instance types for this region
		instanceTypes, err := client.GetInstanceTypes(ctx, v1.GetInstanceTypeArgs{})
		if err != nil {
			t.Logf("  Warning: Failed to get instance types for %s: %v", region, err)
			continue
		}

		regionStats[region] = len(instanceTypes)
		allInstanceTypes = append(allInstanceTypes, instanceTypes...)

		// Count GPUs by family for this region
		gpuCount := 0
		regionGPUs := make(map[string]int)
		for _, it := range instanceTypes {
			if len(it.SupportedGPUs) > 0 {
				gpuCount++
				family := strings.ToLower(it.SupportedGPUs[0].Type)
				regionGPUs[family]++
			}
		}
		gpuFamilies[region] = regionGPUs

		t.Logf("  Found %d instance types (%d with GPUs)", len(instanceTypes), gpuCount)
	}

	// Sort by ID
	sort.Slice(allInstanceTypes, func(i, j int) bool {
		return allInstanceTypes[i].ID < allInstanceTypes[j].ID
	})

	// Output statistics
	t.Logf("\n=== Summary ===")
	t.Logf("Total instance types: %d", len(allInstanceTypes))
	t.Logf("\nBy region:")
	for _, region := range regions {
		if count, ok := regionStats[region]; ok {
			t.Logf("  %s: %d", region, count)
		}
	}

	// GPU families summary
	t.Logf("\nGPU types by region:")
	for _, region := range regions {
		if gpus, ok := gpuFamilies[region]; ok && len(gpus) > 0 {
			t.Logf("  %s:", region)
			families := make([]string, 0, len(gpus))
			for family := range gpus {
				families = append(families, family)
			}
			sort.Strings(families)
			for _, family := range families {
				t.Logf("    %s: %d instance types", strings.ToUpper(family), gpus[family])
			}
		}
	}

	// Write detailed JSON to file
	outputFile := "instance_types_all_regions.json"
	output, err := json.MarshalIndent(allInstanceTypes, "", "  ")
	if err != nil {
		t.Fatalf("Error marshaling JSON: %v", err)
	}

	err = os.WriteFile(outputFile, output, 0o644)
	if err != nil {
		t.Fatalf("Error writing to file: %v", err)
	}

	t.Logf("\nDetailed results written to: %s", outputFile)
}

// Test_EnumerateInstanceTypesSingleRegion enumerates instance types for a specific region
// Usage:
//
//	export NEBIUS_SERVICE_ACCOUNT_JSON='/path/to/service-account.json'
//	export NEBIUS_TENANT_ID='tenant-e00xxx'
//	export NEBIUS_LOCATION='eu-north1'
//	go test -tags scripts -v -run Test_EnumerateInstanceTypesSingleRegion
func Test_EnumerateInstanceTypesSingleRegion(t *testing.T) {
	serviceAccountJSON := os.Getenv("NEBIUS_SERVICE_ACCOUNT_JSON")
	tenantID := os.Getenv("NEBIUS_TENANT_ID")
	location := os.Getenv("NEBIUS_LOCATION")

	if serviceAccountJSON == "" || tenantID == "" {
		t.Skip("NEBIUS_SERVICE_ACCOUNT_JSON and NEBIUS_TENANT_ID must be set")
	}

	if location == "" {
		location = "eu-north1" // default
	}

	ctx := context.Background()

	t.Logf("Enumerating instance types for region: %s", location)

	// Create client
	client, err := nebius.NewNebiusClient(ctx, "enum-script", serviceAccountJSON, tenantID, "", location)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	// Get instance types
	instanceTypes, err := client.GetInstanceTypes(ctx, v1.GetInstanceTypeArgs{})
	if err != nil {
		t.Fatalf("Failed to get instance types: %v", err)
	}

	t.Logf("Found %d instance types", len(instanceTypes))

	// Assert we got at least one instance type
	if len(instanceTypes) == 0 {
		t.Fatal("Expected to receive at least one instance type, but got zero")
	}

	// Categorize by GPU
	cpuTypes := make([]v1.InstanceType, 0)
	gpuTypesByFamily := make(map[string][]v1.InstanceType)

	for _, it := range instanceTypes {
		if len(it.SupportedGPUs) > 0 {
			family := strings.ToUpper(it.SupportedGPUs[0].Type)
			gpuTypesByFamily[family] = append(gpuTypesByFamily[family], it)
		} else {
			cpuTypes = append(cpuTypes, it)
		}
	}

	// Print summary
	t.Logf("\nCPU-only instance types: %d", len(cpuTypes))
	for _, it := range cpuTypes {
		t.Logf("  - %s: %d vCPUs, %d GB RAM", it.ID, it.CPU, it.MemoryGB)
	}

	t.Logf("\nGPU instance types:")
	gpuFamilies := make([]string, 0, len(gpuTypesByFamily))
	for family := range gpuTypesByFamily {
		gpuFamilies = append(gpuFamilies, family)
	}
	sort.Strings(gpuFamilies)

	for _, family := range gpuFamilies {
		types := gpuTypesByFamily[family]
		t.Logf("\n  %s (%d types):", family, len(types))
		for _, it := range types {
			gpu := it.SupportedGPUs[0]
			vramGB := int64(gpu.Memory) / (1024 * 1024 * 1024)
			t.Logf("    - %s: %dx %s (%d GB VRAM each), %d vCPUs, %d GB RAM",
				it.ID, gpu.Count, gpu.Name, vramGB, it.CPU, it.MemoryGB)
			if it.Price != nil {
				t.Logf("      Price: $%.4f/hr", float64(it.Price.Amount)/float64(it.Price.Precision))
			}
		}
	}

	// Write to JSON
	outputFile := fmt.Sprintf("instance_types_%s.json", location)
	output, err := json.MarshalIndent(instanceTypes, "", "  ")
	if err != nil {
		t.Fatalf("Error marshaling JSON: %v", err)
	}

	err = os.WriteFile(outputFile, output, 0o644)
	if err != nil {
		t.Fatalf("Error writing to file: %v", err)
	}

	t.Logf("\nDetailed results written to: %s", outputFile)
}

// Test_EnumerateGPUTypes filters and displays only GPU instance types with detailed specs
// Usage:
//
//	export NEBIUS_SERVICE_ACCOUNT_JSON='/path/to/service-account.json'
//	export NEBIUS_TENANT_ID='tenant-e00xxx'
//	export NEBIUS_LOCATION='eu-north1'
//	go test -tags scripts -v -run Test_EnumerateGPUTypes
func Test_EnumerateGPUTypes(t *testing.T) {
	serviceAccountJSON := os.Getenv("NEBIUS_SERVICE_ACCOUNT_JSON")
	tenantID := os.Getenv("NEBIUS_TENANT_ID")
	location := os.Getenv("NEBIUS_LOCATION")

	if serviceAccountJSON == "" || tenantID == "" {
		t.Skip("NEBIUS_SERVICE_ACCOUNT_JSON and NEBIUS_TENANT_ID must be set")
	}

	if location == "" {
		location = "eu-north1"
	}

	ctx := context.Background()
	client, err := nebius.NewNebiusClient(ctx, "enum-script", serviceAccountJSON, tenantID, "", location)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	instanceTypes, err := client.GetInstanceTypes(ctx, v1.GetInstanceTypeArgs{})
	if err != nil {
		t.Fatalf("Failed to get instance types: %v", err)
	}

	// Assert we got at least one instance type to search through
	if len(instanceTypes) == 0 {
		t.Fatal("Expected to receive at least one instance type, but got zero")
	}

	t.Logf("GPU Instance Types in %s:\n", location)
	t.Logf("%-50s %-15s %-8s %-10s %-10s %-15s", "ID", "GPU Type", "Count", "vCPUs", "RAM (GB)", "VRAM/GPU (GB)")
	t.Logf(strings.Repeat("-", 120))

	gpuCount := 0
	for _, it := range instanceTypes {
		if len(it.SupportedGPUs) > 0 {
			gpuCount++
			gpu := it.SupportedGPUs[0]
			vramGB := int64(gpu.Memory) / (1024 * 1024 * 1024)
			t.Logf("%-50s %-15s %-8d %-10d %-10d %-15d",
				it.ID, gpu.Type, gpu.Count, it.CPU, it.MemoryGB, vramGB)
		}
	}

	t.Logf("\nTotal GPU instance types: %d", gpuCount)
}
