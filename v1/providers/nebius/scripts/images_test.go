//go:build scripts
// +build scripts

package scripts

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"testing"

	v1 "github.com/brevdev/cloud/v1"
	nebius "github.com/brevdev/cloud/v1/providers/nebius"
)

// Test_EnumerateImages enumerates all available images in Nebius
// Usage:
//
//	export NEBIUS_SERVICE_ACCOUNT_JSON='/path/to/service-account.json'
//	export NEBIUS_TENANT_ID='tenant-e00xxx'
//	export NEBIUS_LOCATION='eu-north1'
//	go test -tags scripts -v -run Test_EnumerateImages
func Test_EnumerateImages(t *testing.T) {
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

	t.Logf("Enumerating images in region: %s", location)

	// Create client
	client, err := nebius.NewNebiusClient(ctx, "enum-script", serviceAccountJSON, tenantID, "", location)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	// Get images
	images, err := client.GetImages(ctx, v1.GetImagesArgs{})
	if err != nil {
		t.Fatalf("Failed to get images: %v", err)
	}

	t.Logf("Found %d images", len(images))

	// Assert we got at least one image
	if len(images) == 0 {
		t.Fatal("Expected to receive at least one image, but got zero")
	}

	// Categorize by OS
	imagesByOS := make(map[string][]v1.Image)
	for _, img := range images {
		imagesByOS[img.OS] = append(imagesByOS[img.OS], img)
	}

	// Print summary
	t.Logf("\nImages by OS:")
	osList := make([]string, 0, len(imagesByOS))
	for os := range imagesByOS {
		osList = append(osList, os)
	}
	sort.Strings(osList)

	for _, os := range osList {
		imgs := imagesByOS[os]
		t.Logf("\n  %s (%d images):", os, len(imgs))

		// Sort by version
		sort.Slice(imgs, func(i, j int) bool {
			return imgs[i].Version < imgs[j].Version
		})

		for _, img := range imgs {
			t.Logf("    - %s: %s (Arch: %s, Version: %s)",
				img.ID, img.Name, img.Architecture, img.Version)
		}
	}

	// Write to JSON
	outputFile := fmt.Sprintf("images_%s.json", location)
	output, err := json.MarshalIndent(images, "", "  ")
	if err != nil {
		t.Fatalf("Error marshaling JSON: %v", err)
	}

	err = os.WriteFile(outputFile, output, 0o644)
	if err != nil {
		t.Fatalf("Error writing to file: %v", err)
	}

	t.Logf("\nDetailed results written to: %s", outputFile)
}

// Test_EnumerateImagesAllRegions enumerates images across all Nebius regions
// Usage:
//
//	export NEBIUS_SERVICE_ACCOUNT_JSON='/path/to/service-account.json'
//	export NEBIUS_TENANT_ID='tenant-e00xxx'
//	go test -tags scripts -v -run Test_EnumerateImagesAllRegions
func Test_EnumerateImagesAllRegions(t *testing.T) {
	serviceAccountJSON := os.Getenv("NEBIUS_SERVICE_ACCOUNT_JSON")
	tenantID := os.Getenv("NEBIUS_TENANT_ID")

	if serviceAccountJSON == "" || tenantID == "" {
		t.Skip("NEBIUS_SERVICE_ACCOUNT_JSON and NEBIUS_TENANT_ID must be set")
	}

	ctx := context.Background()

	regions := []string{
		"eu-north1",
		"eu-west1",
		"eu-west2",
		"us-central1",
		"us-east1",
		"asia-east1",
	}

	t.Logf("Enumerating images across %d regions...", len(regions))

	allImages := make(map[string][]v1.Image) // region -> images
	imageIDsByRegion := make(map[string]map[string]bool)

	for _, region := range regions {
		t.Logf("Querying region: %s...", region)

		client, err := nebius.NewNebiusClient(ctx, "enum-script", serviceAccountJSON, tenantID, "", region)
		if err != nil {
			t.Logf("  Warning: Failed to create client for %s: %v", region, err)
			continue
		}

		images, err := client.GetImages(ctx, v1.GetImagesArgs{})
		if err != nil {
			t.Logf("  Warning: Failed to get images for %s: %v", region, err)
			continue
		}

		allImages[region] = images
		t.Logf("  Found %d images", len(images))

		// Track unique image IDs per region
		if imageIDsByRegion[region] == nil {
			imageIDsByRegion[region] = make(map[string]bool)
		}
		for _, img := range images {
			imageIDsByRegion[region][img.ID] = true
		}
	}

	// Summary
	t.Logf("\n=== Summary ===")
	t.Logf("Images by region:")
	for _, region := range regions {
		if imgs, ok := allImages[region]; ok {
			t.Logf("  %s: %d images", region, len(imgs))
		}
	}

	// Write to JSON
	outputFile := "images_all_regions.json"
	output, err := json.MarshalIndent(allImages, "", "  ")
	if err != nil {
		t.Fatalf("Error marshaling JSON: %v", err)
	}

	err = os.WriteFile(outputFile, output, 0o644)
	if err != nil {
		t.Fatalf("Error writing to file: %v", err)
	}

	t.Logf("\nDetailed results written to: %s", outputFile)
}

// Test_FilterGPUImages filters images suitable for GPU instances
// Usage:
//
//	export NEBIUS_SERVICE_ACCOUNT_JSON='/path/to/service-account.json'
//	export NEBIUS_TENANT_ID='tenant-e00xxx'
//	export NEBIUS_LOCATION='eu-north1'
//	go test -tags scripts -v -run Test_FilterGPUImages
func Test_FilterGPUImages(t *testing.T) {
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

	images, err := client.GetImages(ctx, v1.GetImagesArgs{})
	if err != nil {
		t.Fatalf("Failed to get images: %v", err)
	}

	t.Logf("GPU-optimized Images in %s:", location)
	t.Logf("%-50s %-20s %-15s %-20s", "ID", "Name", "OS", "Version")
	t.Logf(strings.Repeat("-", 110))

	gpuImageCount := 0
	for _, img := range images {
		// Look for GPU-related keywords in name or description
		name := strings.ToLower(img.Name)
		if strings.Contains(name, "gpu") ||
			strings.Contains(name, "cuda") ||
			strings.Contains(name, "nvidia") ||
			strings.Contains(name, "ml") ||
			strings.Contains(name, "deep learning") {

			gpuImageCount++
			t.Logf("%-50s %-20s %-15s %-20s",
				img.ID, img.Name, img.OS, img.Version)
		}
	}

	if gpuImageCount == 0 {
		t.Logf("No GPU-specific images found. Showing Ubuntu images (typically GPU-compatible):\n")
		for _, img := range images {
			if strings.Contains(strings.ToLower(img.OS), "ubuntu") {
				t.Logf("%-50s %-20s %-15s %-20s",
					img.ID, img.Name, img.OS, img.Version)
			}
		}
	}

	t.Logf("\nTotal GPU-optimized images: %d", gpuImageCount)
}
