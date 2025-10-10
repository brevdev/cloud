package v1

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/alecthomas/units"
	v1 "github.com/brevdev/cloud/v1"
	common "github.com/nebius/gosdk/proto/nebius/common/v1"
	compute "github.com/nebius/gosdk/proto/nebius/compute/v1"
	vpc "github.com/nebius/gosdk/proto/nebius/vpc/v1"
)

func (c *NebiusClient) CreateInstance(ctx context.Context, attrs v1.CreateInstanceAttrs) (*v1.Instance, error) {
	// Ensure networking infrastructure exists
	subnetID, err := c.ensureNetworkInfrastructure(ctx, attrs.Name)
	if err != nil {
		return nil, fmt.Errorf("failed to ensure network infrastructure: %w", err)
	}

	// Create boot disk first using image family
	bootDiskID, err := c.createBootDisk(ctx, attrs)
	if err != nil {
		return nil, fmt.Errorf("failed to create boot disk: %w", err)
	}

	// Parse platform and preset from instance type
	platform, preset, err := c.parseInstanceType(ctx, attrs.InstanceType)
	if err != nil {
		return nil, fmt.Errorf("failed to parse instance type %s: %w", attrs.InstanceType, err)
	}

	// Create instance specification
	instanceSpec := &compute.InstanceSpec{
		Resources: &compute.ResourcesSpec{
			Platform: platform,
			Size: &compute.ResourcesSpec_Preset{
				Preset: preset,
			},
		},
		NetworkInterfaces: []*compute.NetworkInterfaceSpec{
			{
				Name:      "eth0",
				SubnetId:  subnetID,
				IpAddress: &compute.IPAddress{}, // Auto-assign IP
			},
		},
		BootDisk: &compute.AttachedDiskSpec{
			AttachMode: compute.AttachedDiskSpec_READ_WRITE,
			Type: &compute.AttachedDiskSpec_ExistingDisk{
				ExistingDisk: &compute.ExistingDisk{
					Id: bootDiskID,
				},
			},
			DeviceId: "boot-disk", // User-defined device identifier
		},
	}

	// Create the instance - labels should be in metadata
	createReq := &compute.CreateInstanceRequest{
		Metadata: &common.ResourceMetadata{
			ParentId: c.projectID,
			Name:     attrs.Name,
		},
		Spec: instanceSpec,
	}

	// Add labels/tags to metadata if provided
	if len(attrs.Tags) > 0 {
		createReq.Metadata.Labels = make(map[string]string)
		for k, v := range attrs.Tags {
			createReq.Metadata.Labels[k] = v
		}
		// Add Brev-specific labels
		createReq.Metadata.Labels["created-by"] = "brev-cloud-sdk"
		createReq.Metadata.Labels["brev-user"] = attrs.RefID
	}

	operation, err := c.sdk.Services().Compute().V1().Instance().Create(ctx, createReq)
	if err != nil {
		return nil, fmt.Errorf("failed to create Nebius instance: %w", err)
	}

	// Wait for the operation to complete and get the actual instance ID
	finalOp, err := operation.Wait(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to wait for instance creation: %w", err)
	}

	if !finalOp.Successful() {
		return nil, fmt.Errorf("instance creation failed: %v", finalOp.Status())
	}

	// Get the actual instance ID from the completed operation
	instanceID := finalOp.ResourceID()
	if instanceID == "" {
		return nil, fmt.Errorf("failed to get instance ID from operation")
	}

	instance := &v1.Instance{
		RefID:          attrs.RefID,
		CloudCredRefID: c.refID,
		Name:           attrs.Name,
		Location:       c.location,
		CreatedAt:      time.Now(),
		InstanceType:   attrs.InstanceType,
		ImageID:        attrs.ImageID,
		DiskSize:       attrs.DiskSize,
		Tags:           attrs.Tags,
		CloudID:        v1.CloudProviderInstanceID(instanceID),                // Use actual instance ID
		Status:         v1.Status{LifecycleStatus: v1.LifecycleStatusRunning}, // Instance should be running after successful operation
	}

	return instance, nil
}

func (c *NebiusClient) GetInstance(ctx context.Context, instanceID v1.CloudProviderInstanceID) (*v1.Instance, error) {
	// Query actual Nebius instance
	instance, err := c.sdk.Services().Compute().V1().Instance().Get(ctx, &compute.GetInstanceRequest{
		Id: string(instanceID),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get Nebius instance: %w", err)
	}

	if instance.Metadata == nil || instance.Spec == nil {
		return nil, fmt.Errorf("invalid instance response from Nebius API")
	}

	// Convert Nebius instance status to our status
	var lifecycleStatus v1.LifecycleStatus
	if instance.Status != nil {
		switch instance.Status.State {
		case compute.InstanceStatus_RUNNING:
			lifecycleStatus = v1.LifecycleStatusRunning
		case compute.InstanceStatus_STARTING:
			lifecycleStatus = v1.LifecycleStatusPending
		case compute.InstanceStatus_STOPPING:
			lifecycleStatus = v1.LifecycleStatusStopping
		case compute.InstanceStatus_STOPPED:
			lifecycleStatus = v1.LifecycleStatusStopped
		case compute.InstanceStatus_CREATING:
			lifecycleStatus = v1.LifecycleStatusPending
		case compute.InstanceStatus_DELETING:
			lifecycleStatus = v1.LifecycleStatusTerminating
		case compute.InstanceStatus_ERROR:
			lifecycleStatus = v1.LifecycleStatusFailed
		default:
			lifecycleStatus = v1.LifecycleStatusFailed
		}
	} else {
		lifecycleStatus = v1.LifecycleStatusFailed
	}

	// Extract disk size from boot disk spec
	// Note: For existing disks, we'd need to query the disk separately to get size
	// This is a limitation of the current structure
	var diskSize int
	// TODO: Query the actual disk to get its size if needed

	// Extract creation time
	createdAt := time.Now()
	if instance.Metadata.CreatedAt != nil {
		createdAt = instance.Metadata.CreatedAt.AsTime()
	}

	// Extract labels from metadata
	var tags map[string]string
	var refID string
	if instance.Metadata != nil && len(instance.Metadata.Labels) > 0 {
		tags = instance.Metadata.Labels
		refID = instance.Metadata.Labels["brev-user"] // Extract from labels if available
	}

	return &v1.Instance{
		RefID:          refID,
		CloudCredRefID: c.refID,
		Name:           instance.Metadata.Name,
		CloudID:        instanceID,
		Location:       c.location,
		CreatedAt:      createdAt,
		InstanceType:   instance.Spec.Resources.Platform,
		ImageID:        extractImageFamily(instance.Spec.BootDisk),
		DiskSize:       units.Base2Bytes(diskSize) * units.Gibibyte,
		Tags:           tags,
		Status:         v1.Status{LifecycleStatus: lifecycleStatus},
	}, nil
}

// extractImageFamily extracts the image family from attached disk spec
func extractImageFamily(bootDisk *compute.AttachedDiskSpec) string {
	if bootDisk == nil {
		return ""
	}

	// For existing disks, we'd need to query the disk separately to get its image family
	// This is a limitation when querying existing instances
	// TODO: Query the actual disk to get its source image family if needed
	return ""
}

func (c *NebiusClient) TerminateInstance(ctx context.Context, instanceID v1.CloudProviderInstanceID) error {
	// Delete the instance
	operation, err := c.sdk.Services().Compute().V1().Instance().Delete(ctx, &compute.DeleteInstanceRequest{
		Id: string(instanceID),
	})
	if err != nil {
		return fmt.Errorf("failed to initiate instance termination: %w", err)
	}

	// Wait for the deletion to complete
	finalOp, err := operation.Wait(ctx)
	if err != nil {
		return fmt.Errorf("failed to wait for instance termination: %w", err)
	}

	if !finalOp.Successful() {
		return fmt.Errorf("instance termination failed: %v", finalOp.Status())
	}

	return nil
}

func (c *NebiusClient) ListInstances(ctx context.Context, args v1.ListInstancesArgs) ([]v1.Instance, error) {
	// Simplified implementation - would list actual instances
	return []v1.Instance{}, fmt.Errorf("nebius list instances implementation pending: %w", v1.ErrNotImplemented)
}

func (c *NebiusClient) StopInstance(ctx context.Context, instanceID v1.CloudProviderInstanceID) error {
	return fmt.Errorf("nebius stop instance implementation pending: %w", v1.ErrNotImplemented)
}

func (c *NebiusClient) StartInstance(ctx context.Context, instanceID v1.CloudProviderInstanceID) error {
	return fmt.Errorf("nebius start instance implementation pending: %w", v1.ErrNotImplemented)
}

func (c *NebiusClient) RebootInstance(ctx context.Context, instanceID v1.CloudProviderInstanceID) error {
	return fmt.Errorf("nebius reboot instance implementation pending: %w", v1.ErrNotImplemented)
}

func (c *NebiusClient) ChangeInstanceType(ctx context.Context, instanceID v1.CloudProviderInstanceID, newInstanceType string) error {
	return fmt.Errorf("nebius change instance type implementation pending: %w", v1.ErrNotImplemented)
}

func (c *NebiusClient) UpdateInstanceTags(ctx context.Context, args v1.UpdateInstanceTagsArgs) error {
	return fmt.Errorf("nebius update instance tags implementation pending: %w", v1.ErrNotImplemented)
}

func (c *NebiusClient) ResizeInstanceVolume(ctx context.Context, args v1.ResizeInstanceVolumeArgs) error {
	return fmt.Errorf("nebius resize instance volume implementation pending: %w", v1.ErrNotImplemented)
}

func (c *NebiusClient) AddFirewallRulesToInstance(ctx context.Context, args v1.AddFirewallRulesToInstanceArgs) error {
	return fmt.Errorf("nebius firewall rules management not yet implemented: %w", v1.ErrNotImplemented)
}

func (c *NebiusClient) RevokeSecurityGroupRules(ctx context.Context, args v1.RevokeSecurityGroupRuleArgs) error {
	return fmt.Errorf("nebius security group rules management not yet implemented: %w", v1.ErrNotImplemented)
}

func (c *NebiusClient) GetMaxCreateRequestsPerMinute() int {
	return 10
}

func (c *NebiusClient) MergeInstanceForUpdate(currInst v1.Instance, newInst v1.Instance) v1.Instance {
	merged := newInst
	merged.Name = currInst.Name
	merged.RefID = currInst.RefID
	merged.CloudCredRefID = currInst.CloudCredRefID
	merged.CreatedAt = currInst.CreatedAt
	merged.CloudID = currInst.CloudID
	merged.Location = currInst.Location
	return merged
}

// ensureNetworkInfrastructure creates VPC network and subnet for instance if needed
func (c *NebiusClient) ensureNetworkInfrastructure(ctx context.Context, instanceName string) (string, error) {
	// Create or get VPC network
	networkID, err := c.ensureVPCNetwork(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to ensure VPC network: %w", err)
	}

	// Create or get subnet
	subnetID, err := c.ensureSubnet(ctx, networkID, instanceName)
	if err != nil {
		return "", fmt.Errorf("failed to ensure subnet: %w", err)
	}

	return subnetID, nil
}

// ensureVPCNetwork creates a VPC network for the project if it doesn't exist
func (c *NebiusClient) ensureVPCNetwork(ctx context.Context) (string, error) {
	networkName := fmt.Sprintf("%s-network", c.projectID)

	// Try to find existing network
	networksResp, err := c.sdk.Services().VPC().V1().Network().List(ctx, &vpc.ListNetworksRequest{
		ParentId: c.projectID,
	})
	if err == nil {
		for _, network := range networksResp.GetItems() {
			if network.Metadata != nil && network.Metadata.Name == networkName {
				return network.Metadata.Id, nil
			}
		}
	}

	// Create new VPC network
	createReq := &vpc.CreateNetworkRequest{
		Metadata: &common.ResourceMetadata{
			ParentId: c.projectID,
			Name:     networkName,
			Labels: map[string]string{
				"created-by": "brev-cloud-sdk",
				"brev-user":  c.refID,
			},
		},
		Spec: &vpc.NetworkSpec{
			// Use default network pools
		},
	}

	operation, err := c.sdk.Services().VPC().V1().Network().Create(ctx, createReq)
	if err != nil {
		return "", fmt.Errorf("failed to create VPC network: %w", err)
	}

	// Wait for network creation to complete
	finalOp, err := operation.Wait(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to wait for VPC network creation: %w", err)
	}

	if !finalOp.Successful() {
		return "", fmt.Errorf("VPC network creation failed: %v", finalOp.Status())
	}

	// Get the resource ID directly
	networkID := finalOp.ResourceID()
	if networkID == "" {
		return "", fmt.Errorf("failed to get network ID from operation")
	}

	return networkID, nil
}

// ensureSubnet creates a subnet within the VPC network if it doesn't exist
func (c *NebiusClient) ensureSubnet(ctx context.Context, networkID, instanceName string) (string, error) {
	subnetName := fmt.Sprintf("%s-subnet", strings.ReplaceAll(instanceName, "_", "-"))

	// Try to find existing subnet
	subnetsResp, err := c.sdk.Services().VPC().V1().Subnet().List(ctx, &vpc.ListSubnetsRequest{
		ParentId: c.projectID,
	})
	if err == nil {
		for _, subnet := range subnetsResp.GetItems() {
			if subnet.Metadata != nil && subnet.Metadata.Name == subnetName {
				return subnet.Metadata.Id, nil
			}
		}
	}

	// Create new subnet
	createReq := &vpc.CreateSubnetRequest{
		Metadata: &common.ResourceMetadata{
			ParentId: c.projectID,
			Name:     subnetName,
			Labels: map[string]string{
				"created-by": "brev-cloud-sdk",
				"brev-user":  c.refID,
			},
		},
		Spec: &vpc.SubnetSpec{
			NetworkId: networkID,
			// Use default network pools without explicit CIDR specification
		},
	}

	operation, err := c.sdk.Services().VPC().V1().Subnet().Create(ctx, createReq)
	if err != nil {
		return "", fmt.Errorf("failed to create subnet: %w", err)
	}

	// Wait for subnet creation to complete
	finalOp, err := operation.Wait(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to wait for subnet creation: %w", err)
	}

	if !finalOp.Successful() {
		return "", fmt.Errorf("subnet creation failed: %v", finalOp.Status())
	}

	// Get the resource ID directly
	subnetID := finalOp.ResourceID()
	if subnetID == "" {
		return "", fmt.Errorf("failed to get subnet ID from operation")
	}

	return subnetID, nil
}

// createBootDisk creates a boot disk for the instance using image family or specific image ID
func (c *NebiusClient) createBootDisk(ctx context.Context, attrs v1.CreateInstanceAttrs) (string, error) {
	diskName := fmt.Sprintf("%s-boot-disk", attrs.Name)

	// Try to use image family first, then fallback to specific image ID
	createReq, err := c.buildDiskCreateRequest(ctx, diskName, attrs)
	if err != nil {
		return "", fmt.Errorf("failed to build disk create request: %w", err)
	}

	operation, err := c.sdk.Services().Compute().V1().Disk().Create(ctx, createReq)
	if err != nil {
		return "", fmt.Errorf("failed to create boot disk: %w", err)
	}

	// Wait for disk creation to complete
	finalOp, err := operation.Wait(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to wait for boot disk creation: %w", err)
	}

	if !finalOp.Successful() {
		return "", fmt.Errorf("boot disk creation failed: %v", finalOp.Status())
	}

	// Get the resource ID directly
	diskID := finalOp.ResourceID()
	if diskID == "" {
		return "", fmt.Errorf("failed to get disk ID from operation")
	}

	return diskID, nil
}

// buildDiskCreateRequest builds a disk creation request, trying image family first, then image ID
func (c *NebiusClient) buildDiskCreateRequest(ctx context.Context, diskName string, attrs v1.CreateInstanceAttrs) (*compute.CreateDiskRequest, error) {
	baseReq := &compute.CreateDiskRequest{
		Metadata: &common.ResourceMetadata{
			ParentId: c.projectID,
			Name:     diskName,
			Labels: map[string]string{
				"created-by": "brev-cloud-sdk",
				"brev-user":  c.refID,
			},
		},
		Spec: &compute.DiskSpec{
			Size: &compute.DiskSpec_SizeGibibytes{
				SizeGibibytes: int64(attrs.DiskSize / units.Gibibyte),
			},
			Type: compute.DiskSpec_NETWORK_SSD,
		},
	}

	// First, try to resolve and use image family
	if imageFamily, err := c.resolveImageFamily(ctx, attrs.ImageID); err == nil {
		publicImagesParent := c.getPublicImagesParent()

		// Skip validation for known-good common families to speed up instance start
		knownFamilies := []string{"ubuntu22.04-cuda12", "mk8s-worker-node-v-1-32-ubuntu24.04", "mk8s-worker-node-v-1-32-ubuntu24.04-cuda12.8"}
		isKnownFamily := false
		for _, known := range knownFamilies {
			if imageFamily == known {
				isKnownFamily = true
				break
			}
		}

		if isKnownFamily {
			// Use known family without validation
			baseReq.Spec.Source = &compute.DiskSpec_SourceImageFamily{
				SourceImageFamily: &compute.SourceImageFamily{
					ImageFamily: imageFamily,
					ParentId:    publicImagesParent,
				},
			}
			baseReq.Metadata.Labels["image-family"] = imageFamily
			return baseReq, nil
		}

		// For unknown families, validate first
		_, err := c.sdk.Services().Compute().V1().Image().GetLatestByFamily(ctx, &compute.GetImageLatestByFamilyRequest{
			ParentId:    publicImagesParent,
			ImageFamily: imageFamily,
		})
		if err == nil {
			// Family works, use it
			baseReq.Spec.Source = &compute.DiskSpec_SourceImageFamily{
				SourceImageFamily: &compute.SourceImageFamily{
					ImageFamily: imageFamily,
					ParentId:    publicImagesParent,
				},
			}
			baseReq.Metadata.Labels["image-family"] = imageFamily
			return baseReq, nil
		}
	}

	// Family approach failed, try to use a known working public image ID
	publicImageID, err := c.getWorkingPublicImageID(ctx, attrs.ImageID)
	if err == nil {
		baseReq.Spec.Source = &compute.DiskSpec_SourceImageId{
			SourceImageId: publicImageID,
		}
		baseReq.Metadata.Labels["source-image-id"] = publicImageID
		return baseReq, nil
	}

	// Both approaches failed
	return nil, fmt.Errorf("could not resolve image %s to either a working family or image ID: %w", attrs.ImageID, err)
}

// getWorkingPublicImageID gets a working public image ID based on the requested image type
func (c *NebiusClient) getWorkingPublicImageID(ctx context.Context, requestedImage string) (string, error) {
	// Get available public images from the correct region
	publicImagesParent := c.getPublicImagesParent()
	imagesResp, err := c.sdk.Services().Compute().V1().Image().List(ctx, &compute.ListImagesRequest{
		ParentId: publicImagesParent,
	})
	if err != nil {
		return "", fmt.Errorf("failed to list public images: %w", err)
	}

	if len(imagesResp.GetItems()) == 0 {
		return "", fmt.Errorf("no public images available")
	}

	// Try to find the best match based on the requested image
	requestedLower := strings.ToLower(requestedImage)

	var bestMatch *compute.Image
	var fallbackImage *compute.Image

	for _, image := range imagesResp.GetItems() {
		if image.Metadata == nil {
			continue
		}

		imageName := strings.ToLower(image.Metadata.Name)

		// Set fallback to first available image
		if fallbackImage == nil {
			fallbackImage = image
		}

		// Look for Ubuntu matches
		if strings.Contains(requestedLower, "ubuntu") && strings.Contains(imageName, "ubuntu") {
			// Prefer specific version matches
			if strings.Contains(requestedLower, "24.04") || strings.Contains(requestedLower, "24") {
				if strings.Contains(imageName, "ubuntu24.04") {
					bestMatch = image
					break
				}
			} else if strings.Contains(requestedLower, "22.04") || strings.Contains(requestedLower, "22") {
				if strings.Contains(imageName, "ubuntu22.04") {
					bestMatch = image
					break
				}
			} else if strings.Contains(requestedLower, "20.04") || strings.Contains(requestedLower, "20") {
				if strings.Contains(imageName, "ubuntu20.04") {
					bestMatch = image
					break
				}
			}

			// Any Ubuntu image is better than non-Ubuntu
			if bestMatch == nil {
				bestMatch = image
			}
		}
	}

	// Use best match if found, otherwise fallback
	selectedImage := bestMatch
	if selectedImage == nil {
		selectedImage = fallbackImage
	}

	if selectedImage == nil {
		return "", fmt.Errorf("no suitable public image found")
	}

	return selectedImage.Metadata.Id, nil
}

// getPublicImagesParent determines the correct public images parent ID based on project routing code
func (c *NebiusClient) getPublicImagesParent() string {
	// Extract routing code from project ID
	// Project ID format: project-{routing-code}{identifier}
	// Examples: project-e00a2zkhpr004gvq7e9e07 -> e00
	//          project-u00public-images -> u00

	if len(c.projectID) >= 11 && strings.HasPrefix(c.projectID, "project-") {
		// Extract the 3-character routing code after "project-"
		routingCode := c.projectID[8:11] // e.g., "e00", "u00"
		return fmt.Sprintf("project-%spublic-images", routingCode)
	}

	// Fallback to default if we can't parse the routing code
	return "project-e00public-images" // Default to e00 region
}

// parseInstanceType parses an instance type ID to extract platform and preset
// NEW Format: nebius-{region}-{gpu-type}-{preset} or nebius-{region}-cpu-{preset}
// Examples:
//
//	nebius-eu-north1-l40s-4gpu-96vcpu-768gb
//	nebius-eu-north1-cpu-4vcpu-16gb
func (c *NebiusClient) parseInstanceType(ctx context.Context, instanceTypeID string) (platform string, preset string, err error) {
	// Get the compute platforms to find the correct platform and preset
	platformsResp, err := c.sdk.Services().Compute().V1().Platform().List(ctx, &compute.ListPlatformsRequest{
		ParentId: c.projectID,
	})
	if err != nil {
		return "", "", fmt.Errorf("failed to list platforms: %w", err)
	}

	// Parse the NEW instance type ID format: nebius-{region}-{gpu-type}-{preset}
	// Split by "-" and extract components
	parts := strings.Split(instanceTypeID, "-")
	if len(parts) >= 4 && parts[0] == "nebius" {
		// Format: nebius-{region}-{gpu-type}-{preset-parts...}
		// Example: nebius-eu-north1-l40s-4gpu-96vcpu-768gb
		//          parts[0]=nebius, parts[1]=eu, parts[2]=north1, parts[3]=l40s, parts[4+]=preset

		// Find where the preset starts (after region and gpu-type)
		// Region could be multi-part (eu-north1) so we need to find the GPU type or "cpu"
		var gpuType string
		var presetStartIdx int

		// Look for GPU type indicators or "cpu"
		for i := 1; i < len(parts); i++ {
			partLower := strings.ToLower(parts[i])
			// Check if this part is a known GPU type or "cpu"
			if partLower == "cpu" || partLower == "l40s" || partLower == "h100" ||
				partLower == "h200" || partLower == "a100" || partLower == "v100" ||
				partLower == "b200" || partLower == "a10" || partLower == "t4" || partLower == "l4" {
				gpuType = partLower
				presetStartIdx = i + 1
				break
			}
		}

		if presetStartIdx > 0 && presetStartIdx < len(parts) {
			// Reconstruct the preset name from remaining parts
			presetName := strings.Join(parts[presetStartIdx:], "-")

			// Now find the matching platform based on GPU type
			for _, p := range platformsResp.GetItems() {
				if p.Metadata == nil || p.Spec == nil {
					continue
				}

				platformNameLower := strings.ToLower(p.Metadata.Name)

				// Match platform by GPU type
				if (gpuType == "cpu" && strings.Contains(platformNameLower, "cpu")) ||
					(gpuType != "cpu" && strings.Contains(platformNameLower, gpuType)) {

					// Verify the preset exists in this platform
					for _, preset := range p.Spec.Presets {
						if preset != nil && preset.Name == presetName {
							return p.Metadata.Name, preset.Name, nil
						}
					}

					// If preset not found, use first preset as fallback
					if len(p.Spec.Presets) > 0 && p.Spec.Presets[0] != nil {
						return p.Metadata.Name, p.Spec.Presets[0].Name, nil
					}
				}
			}
		}
	}

	// OLD Format fallback: {platform-id}-{preset}
	// This handles any legacy instance type IDs that might still exist
	for _, platform := range platformsResp.GetItems() {
		if platform.Metadata == nil || platform.Spec == nil {
			continue
		}

		platformID := platform.Metadata.Id

		// Check if the instance type starts with this platform ID
		if strings.HasPrefix(instanceTypeID, platformID+"-") {
			// Extract the preset part (everything after platform ID + "-")
			presetPart := instanceTypeID[len(platformID)+1:] // +1 for the "-"

			// Find the matching preset in this platform
			for _, preset := range platform.Spec.Presets {
				if preset != nil && preset.Name == presetPart {
					// Return platform NAME (not ID) for ResourcesSpec
					return platform.Metadata.Name, preset.Name, nil
				}
			}

			// If preset not found but platform matches, use the first preset as fallback
			if len(platform.Spec.Presets) > 0 && platform.Spec.Presets[0] != nil {
				return platform.Metadata.Name, platform.Spec.Presets[0].Name, nil
			}
		}
	}

	// Fallback: try to find any platform that contains parts of the instance type
	legacyParts := strings.Split(instanceTypeID, "-")
	if len(legacyParts) >= 3 { // computeplatform-xxx-preset
		for _, platform := range platformsResp.GetItems() {
			if platform.Metadata == nil || platform.Spec == nil {
				continue
			}

			// Check if any part of the instance type matches this platform
			platformID := platform.Metadata.Id
			for _, part := range legacyParts {
				if strings.Contains(platformID, part) {
					// Use first available preset
					if len(platform.Spec.Presets) > 0 && platform.Spec.Presets[0] != nil {
						return platform.Metadata.Name, platform.Spec.Presets[0].Name, nil
					}
				}
			}
		}
	}

	// Final fallback: use first available platform and preset
	if len(platformsResp.GetItems()) > 0 {
		platform := platformsResp.GetItems()[0]
		if platform.Metadata != nil && platform.Spec != nil && len(platform.Spec.Presets) > 0 {
			firstPreset := platform.Spec.Presets[0]
			if firstPreset != nil {
				return platform.Metadata.Id, firstPreset.Name, nil
			}
		}
	}

	return "", "", fmt.Errorf("could not parse instance type %s or find suitable platform/preset", instanceTypeID)
}

// resolveImageFamily resolves an ImageID to an image family name
// If ImageID is already a family name, use it directly
// Otherwise, try to get the image and extract its family
func (c *NebiusClient) resolveImageFamily(ctx context.Context, imageID string) (string, error) {
	// Common Nebius image families - if ImageID matches one of these, use it directly
	commonFamilies := []string{
		"ubuntu22.04-cuda12",
		"mk8s-worker-node-v-1-32-ubuntu24.04",
		"mk8s-worker-node-v-1-32-ubuntu24.04-cuda12.8",
		"mk8s-worker-node-v-1-31-ubuntu24.04-cuda12",
		"ubuntu22.04",
		"ubuntu20.04",
		"ubuntu18.04",
	}

	// Check if ImageID is already a known family name
	for _, family := range commonFamilies {
		if imageID == family {
			return family, nil
		}
	}

	// If ImageID looks like a family name pattern (contains dots, dashes, no UUIDs)
	// and doesn't look like a UUID, assume it's a family name
	if !strings.Contains(imageID, "-") || len(imageID) < 32 {
		// Likely a family name, use it directly
		return imageID, nil
	}

	// If it looks like a UUID/ID, try to get the image and extract its family
	image, err := c.sdk.Services().Compute().V1().Image().Get(ctx, &compute.GetImageRequest{
		Id: imageID,
	})
	if err != nil {
		// If we can't get the image, try using the ID as a family name anyway
		// This allows for custom family names that don't match our patterns
		return imageID, nil
	}

	// Extract family from image metadata/labels if available
	if image.Metadata != nil && image.Metadata.Labels != nil {
		if family, exists := image.Metadata.Labels["family"]; exists && family != "" {
			return family, nil
		}
		if family, exists := image.Metadata.Labels["image-family"]; exists && family != "" {
			return family, nil
		}
	}

	// Extract family from image name as fallback
	if image.Metadata != nil && image.Metadata.Name != "" {
		// Try to extract a reasonable family name from the image name
		name := strings.ToLower(image.Metadata.Name)
		if strings.Contains(name, "ubuntu22") || strings.Contains(name, "ubuntu-22") {
			return "ubuntu22.04", nil
		}
		if strings.Contains(name, "ubuntu20") || strings.Contains(name, "ubuntu-20") {
			return "ubuntu20.04", nil
		}
		if strings.Contains(name, "ubuntu18") || strings.Contains(name, "ubuntu-18") {
			return "ubuntu18.04", nil
		}
	}

	// Default fallback - use the original ImageID as family
	// This handles cases where users provide custom family names
	return imageID, nil
}

// deleteBootDisk deletes a boot disk by ID
func (c *NebiusClient) deleteBootDisk(ctx context.Context, diskID string) error {
	operation, err := c.sdk.Services().Compute().V1().Disk().Delete(ctx, &compute.DeleteDiskRequest{
		Id: diskID,
	})
	if err != nil {
		return fmt.Errorf("failed to delete boot disk: %w", err)
	}

	// Wait for disk deletion to complete
	finalOp, err := operation.Wait(ctx)
	if err != nil {
		return fmt.Errorf("failed to wait for boot disk deletion: %w", err)
	}

	if !finalOp.Successful() {
		return fmt.Errorf("boot disk deletion failed: %v", finalOp.Status())
	}

	return nil
}

// cleanupOrphanedBootDisks finds and cleans up boot disks created by smoke tests
func (c *NebiusClient) cleanupOrphanedBootDisks(ctx context.Context, testID string) error {
	// List all disks in the project
	disksResp, err := c.sdk.Services().Compute().V1().Disk().List(ctx, &compute.ListDisksRequest{
		ParentId: c.projectID,
	})
	if err != nil {
		return fmt.Errorf("failed to list disks: %w", err)
	}

	// Find disks that match our test pattern
	for _, disk := range disksResp.GetItems() {
		if disk.Metadata == nil {
			continue
		}

		// Check if this disk belongs to our smoke test
		if strings.Contains(disk.Metadata.Name, testID) ||
			(disk.Metadata.Labels != nil &&
				(disk.Metadata.Labels["test-id"] == testID ||
					disk.Metadata.Labels["created-by"] == "brev-cloud-sdk")) {

			// Delete this orphaned disk
			err := c.deleteBootDisk(ctx, disk.Metadata.Id)
			if err != nil {
				// Log but continue - don't fail the entire cleanup
				fmt.Printf("Failed to delete orphaned disk %s: %v\n", disk.Metadata.Id, err)
			}
		}
	}

	return nil
}
