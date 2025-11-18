package v1

import (
	"context"
	"fmt"
	"strings"

	v1 "github.com/brevdev/cloud/v1"
	compute "github.com/nebius/gosdk/proto/nebius/compute/v1"
)

func (c *NebiusClient) GetImages(ctx context.Context, args v1.GetImageArgs) ([]v1.Image, error) {
	var images []v1.Image

	// First, try to get project-specific images
	projectImages, err := c.getProjectImages(ctx)
	if err == nil && len(projectImages) > 0 {
		images = append(images, projectImages...)
	}

	// Then, get region-specific public images (always include these for broader selection)
	publicImages, err := c.getRegionalPublicImages(ctx, c.location)
	if err == nil {
		images = append(images, publicImages...)
	}

	// If still no images, try cross-region public images as fallback
	if len(images) == 0 {
		fallbackImages, err := c.getCrossRegionPublicImages(ctx)
		if err == nil {
			images = append(images, fallbackImages...)
		}
	}

	// Apply architecture filters - default to x86_64 if no architecture specified
	architectures := args.Architectures
	if len(architectures) == 0 {
		architectures = []string{"x86_64"} // Default to x86_64
	}
	images = filterImagesByArchitectures(images, architectures)

	// Apply name filter if specified
	if len(args.NameFilters) > 0 {
		images = filterImagesByNameFilters(images, args.NameFilters)
	}

	return images, nil
}

// getProjectImages retrieves images specific to the current project
func (c *NebiusClient) getProjectImages(ctx context.Context) ([]v1.Image, error) {
	imagesResp, err := c.sdk.Services().Compute().V1().Image().List(ctx, &compute.ListImagesRequest{
		ParentId: c.projectID,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to list project images: %w", err)
	}

	var images []v1.Image
	for _, image := range imagesResp.GetItems() {
		if image.Metadata == nil || image.Spec == nil {
			continue
		}

		img := v1.Image{
			ID:           image.Metadata.Id,
			Name:         image.Metadata.Name,
			Description:  getImageDescription(image),
			Architecture: extractArchitecture(image),
		}

		if image.Metadata.CreatedAt != nil {
			img.CreatedAt = image.Metadata.CreatedAt.AsTime()
		}

		images = append(images, img)
	}

	return images, nil
}

// getRegionalPublicImages retrieves public images for the specified region
func (c *NebiusClient) getRegionalPublicImages(ctx context.Context, region string) ([]v1.Image, error) {
	// Determine the correct public images parent for this region
	publicParent := c.getPublicImagesParentForRegion(region)

	imagesResp, err := c.sdk.Services().Compute().V1().Image().List(ctx, &compute.ListImagesRequest{
		ParentId: publicParent,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to list public images for region %s: %w", region, err)
	}

	var images []v1.Image
	for _, image := range imagesResp.GetItems() {
		if image.Metadata == nil {
			continue
		}

		img := v1.Image{
			ID:           image.Metadata.Id,
			Name:         image.Metadata.Name,
			Description:  getImageDescription(image),
			Architecture: extractArchitecture(image),
		}

		if image.Metadata.CreatedAt != nil {
			img.CreatedAt = image.Metadata.CreatedAt.AsTime()
		}

		images = append(images, img)
	}

	return images, nil
}

// getCrossRegionPublicImages tries to get public images from other regions as fallback
func (c *NebiusClient) getCrossRegionPublicImages(ctx context.Context) ([]v1.Image, error) {
	// Common region patterns to try
	regions := []string{"eu-north1", "eu-west1", "us-central1"}

	for _, region := range regions {
		if region == c.location {
			continue // Skip current region since we already tried it
		}

		images, err := c.getRegionalPublicImages(ctx, region)
		if err == nil && len(images) > 0 {
			return images, nil // Return first successful region
		}
	}

	return c.getDefaultImages(ctx) // Final fallback
}

// getPublicImagesParentForRegion determines the correct public images parent ID for a region
func (c *NebiusClient) getPublicImagesParentForRegion(region string) string {
	// Map region to routing code patterns
	regionToRoutingCode := map[string]string{
		"eu-north1":  "e00",
		"eu-west1":   "e00",
		"us-central1": "u00",
		"us-west1":   "u00",
		"asia-southeast1": "a00",
	}

	if routingCode, exists := regionToRoutingCode[region]; exists {
		return fmt.Sprintf("project-%spublic-images", routingCode)
	}

	// Fallback: try to extract from current project ID
	return c.getPublicImagesParent()
}

// getDefaultImages returns common public images when no project-specific images are found
func (c *NebiusClient) getDefaultImages(ctx context.Context) ([]v1.Image, error) {
	// Common Nebius public image families
	defaultFamilies := []string{
		"ubuntu22.04-cuda12",
		"ubuntu20.04",
		"ubuntu18.04",
	}

	var images []v1.Image
	for _, family := range defaultFamilies {
		// Try to get latest image from family (use tenant ID for public images)
		image, err := c.sdk.Services().Compute().V1().Image().GetLatestByFamily(ctx, &compute.GetImageLatestByFamilyRequest{
			ParentId:    c.tenantID,
			ImageFamily: family,
		})
		if err != nil {
			continue // Skip if family not available
		}

		if image.Metadata == nil {
			continue
		}

		img := v1.Image{
			ID:          image.Metadata.Id,
			Name:        image.Metadata.Name,
			Description: getImageDescription(image),
			Architecture: "x86_64",
		}

		// Set creation time if available
		if image.Metadata.CreatedAt != nil {
			img.CreatedAt = image.Metadata.CreatedAt.AsTime()
		}

		images = append(images, img)
	}

	return images, nil
}

// getImageDescription extracts description from ImageSpec if available
func getImageDescription(image *compute.Image) string {
	if image.Spec != nil && image.Spec.Description != nil {
		return *image.Spec.Description
	}
	return ""
}

// extractArchitecture extracts architecture information from image metadata
func extractArchitecture(image *compute.Image) string {
	// Check labels for architecture info
	if image.Metadata != nil && image.Metadata.Labels != nil {
		if arch, exists := image.Metadata.Labels["architecture"]; exists {
			return arch
		}
		if arch, exists := image.Metadata.Labels["arch"]; exists {
			return arch
		}
	}

	// Infer from image name
	if image.Metadata != nil {
		name := strings.ToLower(image.Metadata.Name)
		if strings.Contains(name, "arm64") || strings.Contains(name, "aarch64") {
			return "arm64"
		}
		if strings.Contains(name, "x86_64") || strings.Contains(name, "amd64") {
			return "x86_64"
		}
	}

	return "x86_64" // Default assumption
}

// filterImagesByArchitecture filters images by architecture
func filterImagesByArchitecture(images []v1.Image, architecture string) []v1.Image {
	var filtered []v1.Image
	for _, img := range images {
		if img.Architecture == architecture {
			filtered = append(filtered, img)
		}
	}
	return filtered
}

// filterImagesByArchitectures filters images by multiple architectures
func filterImagesByArchitectures(images []v1.Image, architectures []string) []v1.Image {
	if len(architectures) == 0 {
		return images
	}

	var filtered []v1.Image
	for _, img := range images {
		for _, arch := range architectures {
			if img.Architecture == arch {
				filtered = append(filtered, img)
				break
			}
		}
	}
	return filtered
}

// filterImagesByNameFilters filters images by name patterns
func filterImagesByNameFilters(images []v1.Image, nameFilters []string) []v1.Image {
	if len(nameFilters) == 0 {
		return images
	}

	var filtered []v1.Image
	for _, img := range images {
		for _, filter := range nameFilters {
			if strings.Contains(strings.ToLower(img.Name), strings.ToLower(filter)) {
				filtered = append(filtered, img)
				break
			}
		}
	}
	return filtered
}