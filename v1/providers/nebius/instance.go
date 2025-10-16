package v1

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/alecthomas/units"
	"github.com/brevdev/cloud/internal/errors"
	v1 "github.com/brevdev/cloud/v1"
	common "github.com/nebius/gosdk/proto/nebius/common/v1"
	compute "github.com/nebius/gosdk/proto/nebius/compute/v1"
	vpc "github.com/nebius/gosdk/proto/nebius/vpc/v1"
)

func (c *NebiusClient) CreateInstance(ctx context.Context, attrs v1.CreateInstanceAttrs) (*v1.Instance, error) {
	// Track created resources for automatic cleanup on failure
	var networkID, subnetID, bootDiskID string
	cleanupOnError := true
	defer func() {
		if cleanupOnError {
			c.logger.Info(ctx, "cleaning up resources after instance creation failure",
				v1.LogField("refID", attrs.RefID),
				v1.LogField("networkID", networkID),
				v1.LogField("subnetID", subnetID),
				v1.LogField("bootDiskID", bootDiskID))

			// Clean up boot disk
			if bootDiskID != "" {
				if err := c.deleteBootDiskIfExists(ctx, bootDiskID); err != nil {
					c.logger.Error(ctx, err, v1.LogField("bootDiskID", bootDiskID))
				}
			}

			// Clean up network resources
			if err := c.cleanupNetworkResources(ctx, networkID, subnetID); err != nil {
				c.logger.Error(ctx, err, v1.LogField("networkID", networkID), v1.LogField("subnetID", subnetID))
			}
		}
	}()

	// Create isolated networking infrastructure for this instance
	// Use RefID (environmentId) for resource correlation
	var err error
	networkID, subnetID, err = c.createIsolatedNetwork(ctx, attrs.RefID)
	if err != nil {
		return nil, fmt.Errorf("failed to create isolated network: %w", err)
	}

	// Create boot disk first using image family
	bootDiskID, err = c.createBootDisk(ctx, attrs)
	if err != nil {
		return nil, fmt.Errorf("failed to create boot disk: %w", err)
	}

	// Parse platform and preset from instance type
	platform, preset, err := c.parseInstanceType(ctx, attrs.InstanceType)
	if err != nil {
		return nil, fmt.Errorf("failed to parse instance type %s: %w", attrs.InstanceType, err)
	}

	// Generate cloud-init user-data for SSH key injection and firewall configuration
	// This is similar to Shadeform's LaunchConfiguration approach but uses cloud-init
	cloudInitUserData := generateCloudInitUserData(attrs.PublicKey, attrs.FirewallRules)

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
				Name:     "eth0",
				SubnetId: subnetID,
				// Auto-assign private IP
				IpAddress: &compute.IPAddress{},
				// Request public IP for SSH connectivity
				// Static=false means ephemeral IP (allocated with instance, freed on deletion)
				PublicIpAddress: &compute.PublicIPAddress{
					Static: false,
				},
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
		CloudInitUserData: cloudInitUserData, // Inject SSH keys and configure instance via cloud-init
	}

	// Create the instance - labels should be in metadata
	createReq := &compute.CreateInstanceRequest{
		Metadata: &common.ResourceMetadata{
			ParentId: c.projectID,
			Name:     attrs.Name,
		},
		Spec: instanceSpec,
	}

	// Add labels/tags to metadata (always create labels for resource tracking)
	createReq.Metadata.Labels = make(map[string]string)
	for k, v := range attrs.Tags {
		createReq.Metadata.Labels[k] = v
	}
	// Add Brev-specific labels and resource tracking
	createReq.Metadata.Labels["created-by"] = "brev-cloud-sdk"
	createReq.Metadata.Labels["brev-user"] = attrs.RefID
	createReq.Metadata.Labels["environment-id"] = attrs.RefID
	// Track associated resources for cleanup
	createReq.Metadata.Labels["network-id"] = networkID
	createReq.Metadata.Labels["subnet-id"] = subnetID
	createReq.Metadata.Labels["boot-disk-id"] = bootDiskID

	operation, err := c.sdk.Services().Compute().V1().Instance().Create(ctx, createReq)
	if err != nil {
		return nil, errors.WrapAndTrace(err)
	}

	// Wait for the operation to complete and get the actual instance ID
	finalOp, err := operation.Wait(ctx)
	if err != nil {
		return nil, errors.WrapAndTrace(err)
	}

	if !finalOp.Successful() {
		return nil, fmt.Errorf("instance creation failed: %v", finalOp.Status())
	}

	// Get the actual instance ID from the completed operation
	instanceID := finalOp.ResourceID()
	if instanceID == "" {
		return nil, fmt.Errorf("failed to get instance ID from operation")
	}

	// Query the created instance to get IP addresses and full details
	createdInstance, err := c.GetInstance(ctx, v1.CloudProviderInstanceID(instanceID))
	if err != nil {
		// If we can't get instance details, return basic info
		return &v1.Instance{
			RefID:          attrs.RefID,
			CloudCredRefID: c.refID,
			Name:           attrs.Name,
			Location:       c.location,
			CreatedAt:      time.Now(),
			InstanceType:   attrs.InstanceType,
			ImageID:        attrs.ImageID,
			DiskSize:       attrs.DiskSize,
			Tags:           attrs.Tags,
			CloudID:        v1.CloudProviderInstanceID(instanceID),
			Status:         v1.Status{LifecycleStatus: v1.LifecycleStatusPending},
		}, nil
	}

	// Return the full instance details with IP addresses and SSH info
	createdInstance.RefID = attrs.RefID
	createdInstance.CloudCredRefID = c.refID
	createdInstance.Tags = attrs.Tags

	// Success - disable cleanup
	cleanupOnError = false
	return createdInstance, nil
}

func (c *NebiusClient) GetInstance(ctx context.Context, instanceID v1.CloudProviderInstanceID) (*v1.Instance, error) {
	// Query actual Nebius instance
	instance, err := c.sdk.Services().Compute().V1().Instance().Get(ctx, &compute.GetInstanceRequest{
		Id: string(instanceID),
	})
	if err != nil {
		return nil, errors.WrapAndTrace(err)
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

	// Extract IP addresses from network interfaces
	var publicIP, privateIP, hostname string
	if instance.Status != nil && len(instance.Status.NetworkInterfaces) > 0 {
		// Get the first network interface (usually eth0)
		netInterface := instance.Status.NetworkInterfaces[0]

		// Extract private IP (strip CIDR notation if present)
		if netInterface.IpAddress != nil {
			privateIP = stripCIDR(netInterface.IpAddress.Address)
		}

		// Extract public IP (strip CIDR notation if present)
		if netInterface.PublicIpAddress != nil {
			publicIP = stripCIDR(netInterface.PublicIpAddress.Address)
		}

		// Use public IP as hostname if available, otherwise use private IP
		if publicIP != "" {
			hostname = publicIP
		} else {
			hostname = privateIP
		}
	}

	// Determine SSH user based on image
	sshUser := "ubuntu" // Default SSH user for Nebius instances
	imageFamily := extractImageFamily(instance.Spec.BootDisk)
	if strings.Contains(strings.ToLower(imageFamily), "centos") {
		sshUser = "centos"
	} else if strings.Contains(strings.ToLower(imageFamily), "debian") {
		sshUser = "admin"
	}

	return &v1.Instance{
		RefID:          refID,
		CloudCredRefID: c.refID,
		Name:           instance.Metadata.Name,
		CloudID:        instanceID,
		Location:       c.location,
		CreatedAt:      createdAt,
		InstanceType:   instance.Spec.Resources.Platform,
		ImageID:        imageFamily,
		DiskSize:       units.Base2Bytes(diskSize) * units.Gibibyte,
		Tags:           tags,
		Status:         v1.Status{LifecycleStatus: lifecycleStatus},
		// SSH connectivity details
		PublicIP:  publicIP,
		PrivateIP: privateIP,
		PublicDNS: publicIP, // Nebius doesn't provide separate DNS, use public IP
		Hostname:  hostname,
		SSHUser:   sshUser,
		SSHPort:   22, // Standard SSH port
	}, nil
}

// stripCIDR removes CIDR notation from an IP address string
// Nebius API returns IPs in CIDR format (e.g., "192.168.1.1/32")
// We need just the IP address for SSH connectivity
func stripCIDR(ipWithCIDR string) string {
	if ipWithCIDR == "" {
		return ""
	}
	// Check if CIDR notation is present
	if idx := strings.Index(ipWithCIDR, "/"); idx != -1 {
		return ipWithCIDR[:idx]
	}
	return ipWithCIDR
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
	// Get instance details to retrieve associated resource IDs
	instance, err := c.sdk.Services().Compute().V1().Instance().Get(ctx, &compute.GetInstanceRequest{
		Id: string(instanceID),
	})
	if err != nil {
		return fmt.Errorf("failed to get instance details: %w", err)
	}

	// Extract resource IDs from labels
	var networkID, subnetID, bootDiskID string
	if instance.Metadata != nil && instance.Metadata.Labels != nil {
		networkID = instance.Metadata.Labels["network-id"]
		subnetID = instance.Metadata.Labels["subnet-id"]
		bootDiskID = instance.Metadata.Labels["boot-disk-id"]
	}

	// Step 1: Delete the instance
	operation, err := c.sdk.Services().Compute().V1().Instance().Delete(ctx, &compute.DeleteInstanceRequest{
		Id: string(instanceID),
	})
	if err != nil {
		return fmt.Errorf("failed to initiate instance termination: %w", err)
	}

	// Wait for the instance deletion to complete
	finalOp, err := operation.Wait(ctx)
	if err != nil {
		return fmt.Errorf("failed to wait for instance termination: %w", err)
	}

	if !finalOp.Successful() {
		return fmt.Errorf("instance termination failed: %v", finalOp.Status())
	}

	// Step 2: Delete boot disk if it exists and wasn't auto-deleted
	if bootDiskID != "" {
		if err := c.deleteBootDiskIfExists(ctx, bootDiskID); err != nil {
			// Log but don't fail - disk may have been auto-deleted with instance
			fmt.Printf("Warning: failed to delete boot disk %s: %v\n", bootDiskID, err)
		}
	}

	// Step 3: Delete network resources (subnet, then VPC)
	if err := c.cleanupNetworkResources(ctx, networkID, subnetID); err != nil {
		// Log but don't fail - cleanup is best-effort
		fmt.Printf("Warning: failed to cleanup network resources: %v\n", err)
	}

	return nil
}

func (c *NebiusClient) ListInstances(ctx context.Context, args v1.ListInstancesArgs) ([]v1.Instance, error) {
	// Simplified implementation - would list actual instances
	return []v1.Instance{}, fmt.Errorf("nebius list instances implementation pending: %w", v1.ErrNotImplemented)
}

func (c *NebiusClient) StopInstance(ctx context.Context, instanceID v1.CloudProviderInstanceID) error {
	// Initiate instance stop operation
	operation, err := c.sdk.Services().Compute().V1().Instance().Stop(ctx, &compute.StopInstanceRequest{
		Id: string(instanceID),
	})
	if err != nil {
		return fmt.Errorf("failed to initiate instance stop: %w", err)
	}

	// Wait for the stop operation to complete
	finalOp, err := operation.Wait(ctx)
	if err != nil {
		return fmt.Errorf("failed to wait for instance stop: %w", err)
	}

	if !finalOp.Successful() {
		return fmt.Errorf("instance stop failed: %v", finalOp.Status())
	}

	return nil
}

func (c *NebiusClient) StartInstance(ctx context.Context, instanceID v1.CloudProviderInstanceID) error {
	// Initiate instance start operation
	operation, err := c.sdk.Services().Compute().V1().Instance().Start(ctx, &compute.StartInstanceRequest{
		Id: string(instanceID),
	})
	if err != nil {
		return fmt.Errorf("failed to initiate instance start: %w", err)
	}

	// Wait for the start operation to complete
	finalOp, err := operation.Wait(ctx)
	if err != nil {
		return fmt.Errorf("failed to wait for instance start: %w", err)
	}

	if !finalOp.Successful() {
		return fmt.Errorf("instance start failed: %v", finalOp.Status())
	}

	return nil
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

// createIsolatedNetwork creates a dedicated VPC and subnet for a single instance
// This ensures complete network isolation between instances
// Uses refID (environmentId) for resource correlation
func (c *NebiusClient) createIsolatedNetwork(ctx context.Context, refID string) (networkID, subnetID string, err error) {
	// Create VPC network (unique per instance, named with refID for correlation)
	networkName := fmt.Sprintf("%s-vpc", refID)

	createNetworkReq := &vpc.CreateNetworkRequest{
		Metadata: &common.ResourceMetadata{
			ParentId: c.projectID,
			Name:     networkName,
			Labels: map[string]string{
				"created-by":     "brev-cloud-sdk",
				"brev-user":      c.refID,
				"environment-id": refID,
			},
		},
		Spec: &vpc.NetworkSpec{
			// Use default network pools
		},
	}

	networkOp, err := c.sdk.Services().VPC().V1().Network().Create(ctx, createNetworkReq)
	if err != nil {
		return "", "", fmt.Errorf("failed to create isolated VPC network: %w", err)
	}

	// Wait for network creation
	finalNetworkOp, err := networkOp.Wait(ctx)
	if err != nil {
		return "", "", fmt.Errorf("failed to wait for VPC network creation: %w", err)
	}

	if !finalNetworkOp.Successful() {
		return "", "", fmt.Errorf("VPC network creation failed: %v", finalNetworkOp.Status())
	}

	networkID = finalNetworkOp.ResourceID()
	if networkID == "" {
		return "", "", fmt.Errorf("failed to get network ID from operation")
	}

	// Create subnet within the VPC
	subnetName := fmt.Sprintf("%s-subnet", refID)

	createSubnetReq := &vpc.CreateSubnetRequest{
		Metadata: &common.ResourceMetadata{
			ParentId: c.projectID,
			Name:     subnetName,
			Labels: map[string]string{
				"created-by":     "brev-cloud-sdk",
				"brev-user":      c.refID,
				"environment-id": refID,
				"network-id":     networkID,
			},
		},
		Spec: &vpc.SubnetSpec{
			NetworkId: networkID,
			// Use default network pools without explicit CIDR specification
		},
	}

	subnetOp, err := c.sdk.Services().VPC().V1().Subnet().Create(ctx, createSubnetReq)
	if err != nil {
		// Cleanup network if subnet creation fails
		_ = c.deleteNetworkIfExists(ctx, networkID)
		return "", "", fmt.Errorf("failed to create subnet: %w", err)
	}

	// Wait for subnet creation
	finalSubnetOp, err := subnetOp.Wait(ctx)
	if err != nil {
		// Cleanup network if subnet wait fails
		_ = c.deleteNetworkIfExists(ctx, networkID)
		return "", "", fmt.Errorf("failed to wait for subnet creation: %w", err)
	}

	if !finalSubnetOp.Successful() {
		// Cleanup network if subnet creation fails
		_ = c.deleteNetworkIfExists(ctx, networkID)
		return "", "", fmt.Errorf("subnet creation failed: %v", finalSubnetOp.Status())
	}

	subnetID = finalSubnetOp.ResourceID()
	if subnetID == "" {
		// Cleanup network if we can't get subnet ID
		_ = c.deleteNetworkIfExists(ctx, networkID)
		return "", "", fmt.Errorf("failed to get subnet ID from operation")
	}

	return networkID, subnetID, nil
}

// cleanupNetworkResources deletes subnet and VPC network
func (c *NebiusClient) cleanupNetworkResources(ctx context.Context, networkID, subnetID string) error {
	// Delete subnet first (must be deleted before VPC)
	if subnetID != "" {
		if err := c.deleteSubnetIfExists(ctx, subnetID); err != nil {
			return fmt.Errorf("failed to delete subnet: %w", err)
		}
	}

	// Then delete VPC network
	if networkID != "" {
		if err := c.deleteNetworkIfExists(ctx, networkID); err != nil {
			return fmt.Errorf("failed to delete network: %w", err)
		}
	}

	return nil
}

// deleteSubnetIfExists deletes a subnet if it exists
func (c *NebiusClient) deleteSubnetIfExists(ctx context.Context, subnetID string) error {
	operation, err := c.sdk.Services().VPC().V1().Subnet().Delete(ctx, &vpc.DeleteSubnetRequest{
		Id: subnetID,
	})
	if err != nil {
		// Ignore NotFound errors
		if isNotFoundError(err) {
			return nil
		}
		return fmt.Errorf("failed to delete subnet: %w", err)
	}

	// Wait for deletion to complete
	finalOp, err := operation.Wait(ctx)
	if err != nil {
		return fmt.Errorf("failed to wait for subnet deletion: %w", err)
	}

	if !finalOp.Successful() {
		return fmt.Errorf("subnet deletion failed: %v", finalOp.Status())
	}

	return nil
}

// deleteNetworkIfExists deletes a VPC network if it exists
func (c *NebiusClient) deleteNetworkIfExists(ctx context.Context, networkID string) error {
	operation, err := c.sdk.Services().VPC().V1().Network().Delete(ctx, &vpc.DeleteNetworkRequest{
		Id: networkID,
	})
	if err != nil {
		// Ignore NotFound errors
		if isNotFoundError(err) {
			return nil
		}
		return fmt.Errorf("failed to delete network: %w", err)
	}

	// Wait for deletion to complete
	finalOp, err := operation.Wait(ctx)
	if err != nil {
		return fmt.Errorf("failed to wait for network deletion: %w", err)
	}

	if !finalOp.Successful() {
		return fmt.Errorf("network deletion failed: %v", finalOp.Status())
	}

	return nil
}

// createBootDisk creates a boot disk for the instance using image family or specific image ID
// Uses refID (environmentId) for resource correlation
func (c *NebiusClient) createBootDisk(ctx context.Context, attrs v1.CreateInstanceAttrs) (string, error) {
	diskName := fmt.Sprintf("%s-boot-disk", attrs.RefID)

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
				"created-by":     "brev-cloud-sdk",
				"brev-user":      c.refID,
				"environment-id": attrs.RefID,
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
	c.logger.Info(ctx, "parsing instance type",
		v1.LogField("instanceTypeID", instanceTypeID),
		v1.LogField("projectID", c.projectID))

	// Get the compute platforms to find the correct platform and preset
	platformsResp, err := c.sdk.Services().Compute().V1().Platform().List(ctx, &compute.ListPlatformsRequest{
		ParentId: c.projectID,
	})
	if err != nil {
		return "", "", errors.WrapAndTrace(err)
	}

	c.logger.Info(ctx, "listed platforms",
		v1.LogField("platformCount", len(platformsResp.GetItems())))

	// DOT Format: {platform-name}.{preset-name}
	// Example: "gpu-h100-sxm.8gpu-128vcpu-1600gb"
	if strings.Contains(instanceTypeID, ".") {
		dotParts := strings.SplitN(instanceTypeID, ".", 2)
		if len(dotParts) == 2 {
			platformName := dotParts[0]
			presetName := dotParts[1]

			c.logger.Info(ctx, "parsed DOT format instance type",
				v1.LogField("platformName", platformName),
				v1.LogField("presetName", presetName))

			// Find matching platform by name
			for _, p := range platformsResp.GetItems() {
				if p.Metadata == nil || p.Spec == nil {
					continue
				}

				if p.Metadata.Name == platformName {
					// Verify the preset exists
					for _, preset := range p.Spec.Presets {
						if preset != nil && preset.Name == presetName {
							c.logger.Info(ctx, "✓ DOT format EXACT MATCH",
								v1.LogField("platformName", p.Metadata.Name),
								v1.LogField("presetName", preset.Name))
							return p.Metadata.Name, preset.Name, nil
						}
					}

					// If preset not found but platform matches, use first preset
					if len(p.Spec.Presets) > 0 && p.Spec.Presets[0] != nil {
						c.logger.Warn(ctx, "✗ DOT format - preset not found, using first preset",
							v1.LogField("requestedPreset", presetName),
							v1.LogField("fallbackPreset", p.Spec.Presets[0].Name))
						return p.Metadata.Name, p.Spec.Presets[0].Name, nil
					}
				}
			}
		}
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

			c.logger.Info(ctx, "parsed NEW format instance type",
				v1.LogField("gpuType", gpuType),
				v1.LogField("presetName", presetName),
				v1.LogField("presetStartIdx", presetStartIdx))

			// Now find the matching platform based on GPU type
			for _, p := range platformsResp.GetItems() {
				if p.Metadata == nil || p.Spec == nil {
					continue
				}

				platformNameLower := strings.ToLower(p.Metadata.Name)

				// Match platform by GPU type
				if (gpuType == "cpu" && strings.Contains(platformNameLower, "cpu")) ||
					(gpuType != "cpu" && strings.Contains(platformNameLower, gpuType)) {

					// Log ALL available presets for this platform for debugging
					availablePresets := make([]string, 0, len(p.Spec.Presets))
					for _, preset := range p.Spec.Presets {
						if preset != nil {
							availablePresets = append(availablePresets, preset.Name)
						}
					}

					c.logger.Info(ctx, "found matching platform",
						v1.LogField("platformName", p.Metadata.Name),
						v1.LogField("platformID", p.Metadata.Id),
						v1.LogField("presetCount", len(p.Spec.Presets)),
						v1.LogField("requestedPreset", presetName),
						v1.LogField("availablePresets", strings.Join(availablePresets, ", ")))

					// Verify the preset exists in this platform
					for _, preset := range p.Spec.Presets {
						if preset != nil && preset.Name == presetName {
							c.logger.Info(ctx, "✓ EXACT MATCH - using requested preset",
								v1.LogField("platformName", p.Metadata.Name),
								v1.LogField("presetName", preset.Name))
							return p.Metadata.Name, preset.Name, nil
						}
					}

					// If preset not found, use first preset as fallback
					if len(p.Spec.Presets) > 0 && p.Spec.Presets[0] != nil {
						c.logger.Warn(ctx, "✗ MISMATCH - preset not found, using FIRST preset as fallback",
							v1.LogField("requestedPreset", presetName),
							v1.LogField("fallbackPreset", p.Spec.Presets[0].Name),
							v1.LogField("platformName", p.Metadata.Name),
							v1.LogField("availablePresets", strings.Join(availablePresets, ", ")))
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
				c.logger.Warn(ctx, "using final fallback - first available platform/preset",
					v1.LogField("requestedInstanceType", instanceTypeID),
					v1.LogField("fallbackPlatform", platform.Metadata.Name),
					v1.LogField("fallbackPreset", firstPreset.Name))
				return platform.Metadata.Name, firstPreset.Name, nil
			}
		}
	}

	c.logger.Error(ctx, fmt.Errorf("no platforms available"),
		v1.LogField("instanceTypeID", instanceTypeID))
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

// deleteBootDiskIfExists deletes a boot disk if it exists (ignores NotFound errors)
func (c *NebiusClient) deleteBootDiskIfExists(ctx context.Context, diskID string) error {
	operation, err := c.sdk.Services().Compute().V1().Disk().Delete(ctx, &compute.DeleteDiskRequest{
		Id: diskID,
	})
	if err != nil {
		// Ignore NotFound errors - disk may have been auto-deleted with instance
		if isNotFoundError(err) {
			return nil
		}
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

// generateCloudInitUserData generates a cloud-init user-data script for SSH key injection and firewall configuration
// This is inspired by Shadeform's LaunchConfiguration approach but uses cloud-init instead of base64 scripts
func generateCloudInitUserData(publicKey string, firewallRules v1.FirewallRules) string {
	// Start with cloud-init header
	script := "#cloud-config\n"

	// Add SSH key configuration if provided
	if publicKey != "" {
		script += fmt.Sprintf(`ssh_authorized_keys:
  - %s
`, publicKey)
	}

	// Generate UFW firewall commands (similar to Shadeform's approach)
	// UFW (Uncomplicated Firewall) is available on Ubuntu/Debian instances
	ufwCommands := generateUFWCommands(firewallRules)

	if len(ufwCommands) > 0 {
		// Use runcmd to execute firewall setup commands
		script += "\nruncmd:\n"
		for _, cmd := range ufwCommands {
			script += fmt.Sprintf("  - %s\n", cmd)
		}
	}

	return script
}

// generateUFWCommands generates UFW firewall commands similar to Shadeform
// This follows the same pattern as Shadeform's GenerateFirewallScript
func generateUFWCommands(firewallRules v1.FirewallRules) []string {
	commands := []string{
		"ufw --force reset",          // Reset to clean state
		"ufw default deny incoming",  // Default deny incoming
		"ufw default allow outgoing", // Default allow outgoing
		"ufw allow 22/tcp",           // Always allow SSH on port 22
		"ufw allow 2222/tcp",         // Also allow alternate SSH port
	}

	// Add ingress rules
	for _, rule := range firewallRules.IngressRules {
		commands = append(commands, convertIngressRuleToUFW(rule)...)
	}

	// Add egress rules
	for _, rule := range firewallRules.EgressRules {
		commands = append(commands, convertEgressRuleToUFW(rule)...)
	}

	// Enable the firewall
	commands = append(commands, "ufw --force enable")

	return commands
}

// convertIngressRuleToUFW converts an ingress firewall rule to UFW command(s)
func convertIngressRuleToUFW(rule v1.FirewallRule) []string {
	cmds := []string{}
	portSpecs := []string{}

	if rule.FromPort == rule.ToPort {
		portSpecs = append(portSpecs, fmt.Sprintf("port %d", rule.FromPort))
	} else {
		// Port ranges require two separate rules for tcp and udp
		portSpecs = append(portSpecs, fmt.Sprintf("port %d:%d proto tcp", rule.FromPort, rule.ToPort))
		portSpecs = append(portSpecs, fmt.Sprintf("port %d:%d proto udp", rule.FromPort, rule.ToPort))
	}

	if len(rule.IPRanges) == 0 {
		for _, portSpec := range portSpecs {
			cmds = append(cmds, fmt.Sprintf("ufw allow in from any to any %s", portSpec))
		}
	} else {
		for _, ipRange := range rule.IPRanges {
			for _, portSpec := range portSpecs {
				cmds = append(cmds, fmt.Sprintf("ufw allow in from %s to any %s", ipRange, portSpec))
			}
		}
	}

	return cmds
}

// convertEgressRuleToUFW converts an egress firewall rule to UFW command(s)
func convertEgressRuleToUFW(rule v1.FirewallRule) []string {
	cmds := []string{}
	portSpecs := []string{}

	if rule.FromPort == rule.ToPort {
		portSpecs = append(portSpecs, fmt.Sprintf("port %d", rule.FromPort))
	} else {
		// Port ranges require two separate rules for tcp and udp
		portSpecs = append(portSpecs, fmt.Sprintf("port %d:%d proto tcp", rule.FromPort, rule.ToPort))
		portSpecs = append(portSpecs, fmt.Sprintf("port %d:%d proto udp", rule.FromPort, rule.ToPort))
	}

	if len(rule.IPRanges) == 0 {
		for _, portSpec := range portSpecs {
			cmds = append(cmds, fmt.Sprintf("ufw allow out to any %s", portSpec))
		}
	} else {
		for _, ipRange := range rule.IPRanges {
			for _, portSpec := range portSpecs {
				cmds = append(cmds, fmt.Sprintf("ufw allow out to %s %s", ipRange, portSpec))
			}
		}
	}

	return cmds
}
