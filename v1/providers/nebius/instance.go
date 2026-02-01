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

const (
	platformTypeCPU = "cpu"
)

//nolint:gocyclo,funlen // Complex instance creation with resource management
func (c *NebiusClient) CreateInstance(ctx context.Context, attrs v1.CreateInstanceAttrs) (*v1.Instance, error) {
	// Track created resources for automatic cleanup on failure
	var networkID, subnetID, bootDiskID, instanceID string
	cleanupOnError := true
	defer func() {
		if cleanupOnError {
			c.logger.Info(ctx, "cleaning up resources after instance creation failure",
				v1.LogField("refID", attrs.RefID),
				v1.LogField("instanceID", instanceID),
				v1.LogField("networkID", networkID),
				v1.LogField("subnetID", subnetID),
				v1.LogField("bootDiskID", bootDiskID))

			// Clean up instance if it was created
			if instanceID != "" {
				if err := c.deleteInstanceIfExists(ctx, v1.CloudProviderInstanceID(instanceID)); err != nil {
					c.logger.Error(ctx, err, v1.LogField("instanceID", instanceID))
				}
			}

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
	// Use RefID for naming consistency with VPC, subnet, and boot disk
	createReq := &compute.CreateInstanceRequest{
		Metadata: &common.ResourceMetadata{
			ParentId: c.projectID,
			Name:     attrs.RefID,
		},
		Spec: instanceSpec,
	}

	// Add labels/tags to metadata (always create labels for resource tracking)
	createReq.Metadata.Labels = make(map[string]string)
	c.logger.Info(ctx, "Setting instance tags during CreateInstance",
		v1.LogField("providedTagsCount", len(attrs.Tags)),
		v1.LogField("providedTags", fmt.Sprintf("%+v", attrs.Tags)),
		v1.LogField("refID", attrs.RefID))
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
	// Store full instance type ID for later retrieval (dot format: "gpu-h100-sxm.8gpu-128vcpu-1600gb")
	createReq.Metadata.Labels["instance-type-id"] = attrs.InstanceType

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
	// Assign to the outer variable for cleanup tracking
	instanceID = finalOp.ResourceID()
	if instanceID == "" {
		return nil, fmt.Errorf("failed to get instance ID from operation")
	}

	// Wait for instance to reach a stable state (RUNNING or terminal failure)
	// This prevents leaving orphaned resources if the instance fails after creation
	c.logger.Info(ctx, "waiting for instance to reach RUNNING state",
		v1.LogField("instanceID", instanceID),
		v1.LogField("refID", attrs.RefID))

	createdInstance, err := c.waitForInstanceRunning(ctx, v1.CloudProviderInstanceID(instanceID), attrs.RefID, 5*time.Minute)
	if err != nil {
		// Instance failed to reach RUNNING state - cleanup will be triggered by defer
		c.logger.Error(ctx, fmt.Errorf("instance failed to reach RUNNING state: %w", err),
			v1.LogField("instanceID", instanceID))
		return nil, fmt.Errorf("instance failed to reach RUNNING state: %w", err)
	}

	// Return the full instance details with IP addresses and SSH info
	createdInstance.RefID = attrs.RefID
	createdInstance.CloudCredRefID = c.refID
	createdInstance.Tags = attrs.Tags

	// Success - instance reached RUNNING state
	// Disable cleanup and return
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

	return c.convertNebiusInstanceToV1(ctx, instance, nil)
}

// convertNebiusInstanceToV1 converts a Nebius instance to v1.Instance
// This is used by both GetInstance and ListInstances for consistent conversion
// projectToRegion is an optional map of project ID to region for determining instance location
//
//nolint:gocognit,gocyclo,funlen // Complex function converting Nebius instance to v1.Instance with many field mappings
func (c *NebiusClient) convertNebiusInstanceToV1(ctx context.Context, instance *compute.Instance, projectToRegion map[string]string) (*v1.Instance, error) {
	if instance.Metadata == nil || instance.Spec == nil {
		return nil, fmt.Errorf("invalid instance response from Nebius API")
	}

	instanceID := v1.CloudProviderInstanceID(instance.Metadata.Id)

	// Determine location from instance's parent project
	// This ensures instances are correctly attributed to their actual region
	location := c.location // Default to client's location
	if instance.Metadata.ParentId != "" && projectToRegion != nil {
		if region, exists := projectToRegion[instance.Metadata.ParentId]; exists && region != "" {
			location = region
		}
	}

	c.logger.Debug(ctx, "determined instance location",
		v1.LogField("instanceID", instance.Metadata.Id),
		v1.LogField("parentProjectID", instance.Metadata.ParentId),
		v1.LogField("determinedLocation", location),
		v1.LogField("clientLocation", c.location))

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

	// Extract disk size from boot disk by querying the disk
	var diskSize int64 // in bytes
	if instance.Metadata != nil && instance.Metadata.Labels != nil {
		bootDiskID := instance.Metadata.Labels["boot-disk-id"]
		if bootDiskID != "" {
			diskSizeBytes, err := c.getBootDiskSize(ctx, bootDiskID)
			if err != nil {
				c.logger.Error(ctx, fmt.Errorf("failed to get boot disk size: %w", err),
					v1.LogField("bootDiskID", bootDiskID))
				// Don't fail, just use 0 as fallback
			} else {
				diskSize = diskSizeBytes
			}
		}
	}

	// Extract creation time
	createdAt := time.Now()
	if instance.Metadata.CreatedAt != nil {
		createdAt = instance.Metadata.CreatedAt.AsTime()
	}

	// Extract labels from metadata
	var tags map[string]string
	var refID string
	var instanceTypeID string
	if instance.Metadata != nil && len(instance.Metadata.Labels) > 0 {
		tags = instance.Metadata.Labels
		refID = instance.Metadata.Labels["brev-user"]                 // Extract from labels if available
		instanceTypeID = instance.Metadata.Labels["instance-type-id"] // Full instance type ID (dot format)
	}

	// If instance type ID is not in labels (older instances), reconstruct it from platform + preset
	// This is a fallback for backwards compatibility
	if instanceTypeID == "" && instance.Spec.Resources != nil {
		platform := instance.Spec.Resources.Platform
		var preset string
		if instance.Spec.Resources.Size != nil {
			if presetSpec, ok := instance.Spec.Resources.Size.(*compute.ResourcesSpec_Preset); ok {
				preset = presetSpec.Preset
			}
		}
		if platform != "" && preset != "" {
			instanceTypeID = fmt.Sprintf("%s.%s", platform, preset)
		} else {
			// Last resort: just use platform name (less accurate but prevents total failure)
			instanceTypeID = platform
		}
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

	inst := &v1.Instance{
		RefID:          refID,
		CloudCredRefID: c.refID,
		Name:           instance.Metadata.Name,
		CloudID:        instanceID,
		Location:       location,
		CreatedAt:      createdAt,
		InstanceType:   instanceTypeID, // Full instance type ID (e.g., "gpu-h100-sxm.8gpu-128vcpu-1600gb")
		ImageID:        imageFamily,
		DiskSize:       units.Base2Bytes(diskSize),
		DiskSizeBytes:  v1.NewBytes(v1.BytesValue(diskSize), v1.Byte), // diskSize is already in bytes from getBootDiskSize
		Tags:           tags,
		Status:         v1.Status{LifecycleStatus: lifecycleStatus},
		// SSH connectivity details
		PublicIP:  publicIP,
		PrivateIP: privateIP,
		PublicDNS: publicIP, // Nebius doesn't provide separate DNS, use public IP
		Hostname:  hostname,
		SSHUser:   sshUser,
		SSHPort:   22, // Standard SSH port
	}
	inst.InstanceTypeID = v1.MakeGenericInstanceTypeIDFromInstance(*inst)
	return inst, nil
}

// waitForInstanceRunning polls the instance until it reaches RUNNING state or fails
// This prevents orphaned resources when instances fail after the create API call succeeds
func (c *NebiusClient) waitForInstanceRunning(ctx context.Context, instanceID v1.CloudProviderInstanceID, refID string, timeout time.Duration) (*v1.Instance, error) {
	deadline := time.Now().Add(timeout)
	pollInterval := 10 * time.Second

	c.logger.Info(ctx, "polling instance state until RUNNING or terminal failure",
		v1.LogField("instanceID", instanceID),
		v1.LogField("refID", refID),
		v1.LogField("timeout", timeout.String()))

	for {
		// Check if we've exceeded the timeout
		if time.Now().After(deadline) {
			return nil, fmt.Errorf("timeout waiting for instance to reach RUNNING state after %v", timeout)
		}

		// Check if context is canceled
		if ctx.Err() != nil {
			return nil, fmt.Errorf("context canceled while waiting for instance: %w", ctx.Err())
		}

		// Get current instance state
		instance, err := c.GetInstance(ctx, instanceID)
		if err != nil {
			c.logger.Error(ctx, fmt.Errorf("failed to query instance state: %w", err),
				v1.LogField("instanceID", instanceID))
			// Don't fail immediately on transient errors, keep polling
			time.Sleep(pollInterval)
			continue
		}

		c.logger.Info(ctx, "instance state check",
			v1.LogField("instanceID", instanceID),
			v1.LogField("status", instance.Status.LifecycleStatus))

		// Check for success: RUNNING state
		if instance.Status.LifecycleStatus == v1.LifecycleStatusRunning {
			c.logger.Info(ctx, "instance reached RUNNING state",
				v1.LogField("instanceID", instanceID),
				v1.LogField("refID", refID))
			return instance, nil
		}

		// Check for terminal failure states
		if instance.Status.LifecycleStatus == v1.LifecycleStatusFailed ||
			instance.Status.LifecycleStatus == v1.LifecycleStatusTerminated {
			return nil, fmt.Errorf("instance entered terminal failure state: %s", instance.Status.LifecycleStatus)
		}

		// Instance is still in transitional state (PENDING, STARTING, etc.)
		// Wait and poll again
		c.logger.Info(ctx, "instance still transitioning, waiting...",
			v1.LogField("instanceID", instanceID),
			v1.LogField("currentStatus", instance.Status.LifecycleStatus),
			v1.LogField("pollInterval", pollInterval.String()))
		time.Sleep(pollInterval)
	}
}

// waitForInstanceState is a generic helper that waits for an instance to reach a specific lifecycle state
// Used by StopInstance (wait for STOPPED), StartInstance (wait for RUNNING), etc.
func (c *NebiusClient) waitForInstanceState(ctx context.Context, instanceID v1.CloudProviderInstanceID, targetState v1.LifecycleStatus, timeout time.Duration) error {
	deadline := time.Now().Add(timeout)
	pollInterval := 5 * time.Second

	c.logger.Info(ctx, "waiting for instance to reach target state",
		v1.LogField("instanceID", instanceID),
		v1.LogField("targetState", targetState),
		v1.LogField("timeout", timeout.String()))

	for {
		// Check if we've exceeded the timeout
		if time.Now().After(deadline) {
			return fmt.Errorf("timeout waiting for instance to reach %s state after %v", targetState, timeout)
		}

		// Check if context is canceled
		if ctx.Err() != nil {
			return fmt.Errorf("context canceled while waiting for instance: %w", ctx.Err())
		}

		// Get current instance state
		instance, err := c.GetInstance(ctx, instanceID)
		if err != nil {
			c.logger.Error(ctx, fmt.Errorf("failed to query instance state: %w", err),
				v1.LogField("instanceID", instanceID))
			// Don't fail immediately on transient errors, keep polling
			time.Sleep(pollInterval)
			continue
		}

		c.logger.Info(ctx, "instance state check",
			v1.LogField("instanceID", instanceID),
			v1.LogField("currentState", instance.Status.LifecycleStatus),
			v1.LogField("targetState", targetState))

		// Check if we've reached the target state
		if instance.Status.LifecycleStatus == targetState {
			c.logger.Info(ctx, "instance reached target state",
				v1.LogField("instanceID", instanceID),
				v1.LogField("state", targetState))
			return nil
		}

		// Check for terminal failure states (unless we're specifically waiting for a failed state)
		if targetState != v1.LifecycleStatusFailed && targetState != v1.LifecycleStatusTerminated {
			if instance.Status.LifecycleStatus == v1.LifecycleStatusFailed ||
				instance.Status.LifecycleStatus == v1.LifecycleStatusTerminated {
				return fmt.Errorf("instance entered terminal failure state: %s while waiting for %s",
					instance.Status.LifecycleStatus, targetState)
			}
		}

		// Instance is still transitioning, wait and poll again
		c.logger.Info(ctx, "instance still transitioning, waiting...",
			v1.LogField("instanceID", instanceID),
			v1.LogField("currentState", instance.Status.LifecycleStatus),
			v1.LogField("targetState", targetState),
			v1.LogField("pollInterval", pollInterval.String()))
		time.Sleep(pollInterval)
	}
}

// waitForInstanceDeleted polls until the instance is fully deleted (NotFound)
// This is different from waitForInstanceState because deletion results in the instance disappearing
func (c *NebiusClient) waitForInstanceDeleted(ctx context.Context, instanceID v1.CloudProviderInstanceID, timeout time.Duration) error {
	deadline := time.Now().Add(timeout)
	pollInterval := 5 * time.Second

	c.logger.Info(ctx, "waiting for instance to be fully deleted",
		v1.LogField("instanceID", instanceID),
		v1.LogField("timeout", timeout.String()))

	for {
		// Check if we've exceeded the timeout
		if time.Now().After(deadline) {
			return fmt.Errorf("timeout waiting for instance to be deleted after %v", timeout)
		}

		// Check if context is canceled
		if ctx.Err() != nil {
			return fmt.Errorf("context canceled while waiting for instance deletion: %w", ctx.Err())
		}

		// Try to get the instance
		instance, err := c.GetInstance(ctx, instanceID)
		if err != nil {
			// Check if it's a NotFound error - that means the instance is fully deleted
			if isNotFoundError(err) {
				c.logger.Info(ctx, "instance successfully deleted (NotFound)",
					v1.LogField("instanceID", instanceID))
				return nil
			}
			// Other errors - log but keep polling
			c.logger.Error(ctx, fmt.Errorf("error querying instance during deletion wait: %w", err),
				v1.LogField("instanceID", instanceID))
			time.Sleep(pollInterval)
			continue
		}

		// Instance still exists - check its state
		c.logger.Info(ctx, "instance still exists, checking state",
			v1.LogField("instanceID", instanceID),
			v1.LogField("state", instance.Status.LifecycleStatus))

		// If instance is in TERMINATED state, consider it deleted
		if instance.Status.LifecycleStatus == v1.LifecycleStatusTerminated {
			c.logger.Info(ctx, "instance reached TERMINATED state",
				v1.LogField("instanceID", instanceID))
			return nil
		}

		// Instance still in DELETING or other transitional state, wait and poll again
		c.logger.Info(ctx, "instance still deleting, waiting...",
			v1.LogField("instanceID", instanceID),
			v1.LogField("currentState", instance.Status.LifecycleStatus),
			v1.LogField("pollInterval", pollInterval.String()))
		time.Sleep(pollInterval)
	}
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
//
//nolint:unparam // Reserved for future image metadata extraction
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
	c.logger.Info(ctx, "initiating instance termination",
		v1.LogField("instanceID", instanceID))

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

	// Wait for the deletion operation to complete
	finalOp, err := operation.Wait(ctx)
	if err != nil {
		return fmt.Errorf("failed to wait for instance termination: %w", err)
	}

	if !finalOp.Successful() {
		return fmt.Errorf("instance termination failed: %v", finalOp.Status())
	}

	c.logger.Info(ctx, "delete operation completed, waiting for instance to be fully deleted",
		v1.LogField("instanceID", instanceID))

	// Step 2: Wait for instance to be actually deleted (not just "DELETING")
	// We MUST wait because we need to clean up boot disk, subnet, and VPC
	// These resources cannot be deleted while still attached to the instance
	if err := c.waitForInstanceDeleted(ctx, instanceID, 5*time.Minute); err != nil {
		return fmt.Errorf("instance failed to complete deletion: %w", err)
	}

	c.logger.Info(ctx, "instance fully deleted, proceeding with resource cleanup",
		v1.LogField("instanceID", instanceID))

	// Step 3: Delete boot disk if it exists and wasn't auto-deleted
	if bootDiskID != "" {
		if err := c.deleteBootDiskIfExists(ctx, bootDiskID); err != nil {
			// Log but don't fail - disk may have been auto-deleted with instance
			c.logger.Error(ctx, fmt.Errorf("failed to delete boot disk: %w", err),
				v1.LogField("bootDiskID", bootDiskID))
		}
	}

	// Step 4: Delete network resources (subnet, then VPC)
	if err := c.cleanupNetworkResources(ctx, networkID, subnetID); err != nil {
		// Log but don't fail - cleanup is best-effort
		c.logger.Error(ctx, fmt.Errorf("failed to cleanup network resources: %w", err),
			v1.LogField("networkID", networkID),
			v1.LogField("subnetID", subnetID))
	}

	c.logger.Info(ctx, "instance successfully terminated and cleaned up",
		v1.LogField("instanceID", instanceID))

	return nil
}

// deleteInstanceIfExists deletes an instance and ignores NotFound errors
// Used during cleanup to handle cases where the instance may have already been deleted
func (c *NebiusClient) deleteInstanceIfExists(ctx context.Context, instanceID v1.CloudProviderInstanceID) error {
	if instanceID == "" {
		return nil
	}

	// Try to delete the instance - TerminateInstance handles all cleanup
	err := c.TerminateInstance(ctx, instanceID)
	if err != nil {
		// Ignore NotFound errors - instance may have already been deleted
		if isNotFoundError(err) {
			c.logger.Info(ctx, "instance already deleted or not found",
				v1.LogField("instanceID", instanceID))
			return nil
		}
		return fmt.Errorf("failed to delete instance: %w", err)
	}

	c.logger.Info(ctx, "successfully deleted instance",
		v1.LogField("instanceID", instanceID))
	return nil
}

//nolint:gocognit,gocyclo,funlen // Complex function listing instances across multiple projects with filtering
func (c *NebiusClient) ListInstances(ctx context.Context, args v1.ListInstancesArgs) ([]v1.Instance, error) {
	c.logger.Info(ctx, "listing nebius instances",
		v1.LogField("primaryProjectID", c.projectID),
		v1.LogField("location", c.location),
		v1.LogField("tagFilters", fmt.Sprintf("%+v", args.TagFilters)),
		v1.LogField("instanceIDFilter", fmt.Sprintf("%+v", args.InstanceIDs)),
		v1.LogField("locationFilter", fmt.Sprintf("%+v", args.Locations)))

	// Query ALL projects in the tenant to find all instances
	// Projects are region-specific, so we need to check all projects to find all instances
	// Build project-to-region mapping to correctly set Location field on instances
	projectToRegion, err := c.discoverAllProjectsWithRegions(ctx)
	if err != nil {
		c.logger.Error(ctx, fmt.Errorf("failed to discover projects with regions: %w", err))
		// Fallback: just use primary project with client's location
		projectToRegion = map[string]string{c.projectID: c.location}
	}

	c.logger.Info(ctx, "querying instances across all projects",
		v1.LogField("projectCount", len(projectToRegion)),
		v1.LogField("projects", fmt.Sprintf("%v", projectToRegion)))

	// Collect instances from all projects
	allNebiusInstances := make([]*compute.Instance, 0)
	for projectID := range projectToRegion {
		var pageToken string
		for {
			response, err := c.sdk.Services().Compute().V1().Instance().List(ctx, &compute.ListInstancesRequest{
				ParentId:  projectID,
				PageSize:  100,
				PageToken: pageToken,
			})
			if err != nil {
				c.logger.Error(ctx, fmt.Errorf("failed to list instances in project %s: %w", projectID, err),
					v1.LogField("projectID", projectID))
				// Continue to next project instead of failing completely
				break
			}

			// If the response is nil, we've reached the end of the list
			if response == nil {
				break
			}

			if len(response.Items) > 0 {
				c.logger.Info(ctx, "found instances in project",
					v1.LogField("projectID", projectID),
					v1.LogField("region", projectToRegion[projectID]),
					v1.LogField("count", len(response.Items)),
					v1.LogField("page", pageToken))
				allNebiusInstances = append(allNebiusInstances, response.Items...)
			}

			pageToken = response.GetNextPageToken()
			if pageToken == "" {
				break
			}
		}
	}

	if len(allNebiusInstances) == 0 {
		c.logger.Info(ctx, "no instances found across all projects")
		return []v1.Instance{}, nil
	}

	c.logger.Info(ctx, "found raw instances from Nebius API across all projects",
		v1.LogField("totalCount", len(allNebiusInstances)))

	// Convert and filter each Nebius instance to v1.Instance
	instances := make([]v1.Instance, 0, len(allNebiusInstances))
	for _, nebiusInstance := range allNebiusInstances {
		if nebiusInstance.Metadata == nil {
			c.logger.Error(ctx, fmt.Errorf("instance has no metadata"),
				v1.LogField("instanceID", "unknown"))
			continue
		}

		c.logger.Info(ctx, "Processing instance from Nebius API",
			v1.LogField("instanceID", nebiusInstance.Metadata.Id),
			v1.LogField("instanceName", nebiusInstance.Metadata.Name),
			v1.LogField("rawLabelsCount", len(nebiusInstance.Metadata.Labels)),
			v1.LogField("rawLabels", fmt.Sprintf("%+v", nebiusInstance.Metadata.Labels)))

		// Convert to v1.Instance using convertNebiusInstanceToV1 for consistent conversion
		// Pass projectToRegion mapping so instances get correct location from their parent project
		instance, err := c.convertNebiusInstanceToV1(ctx, nebiusInstance, projectToRegion)
		if err != nil {
			c.logger.Error(ctx, fmt.Errorf("failed to convert instance: %w", err),
				v1.LogField("instanceID", nebiusInstance.Metadata.Id))
			continue
		}

		c.logger.Info(ctx, "Instance after conversion",
			v1.LogField("instanceID", instance.CloudID),
			v1.LogField("convertedTagsCount", len(instance.Tags)),
			v1.LogField("convertedTags", fmt.Sprintf("%+v", instance.Tags)))

		// Apply tag filtering if TagFilters are provided
		if len(args.TagFilters) > 0 {
			c.logger.Info(ctx, "🔎 Checking tag filters",
				v1.LogField("instanceID", instance.CloudID),
				v1.LogField("requiredFilters", fmt.Sprintf("%+v", args.TagFilters)),
				v1.LogField("instanceTags", fmt.Sprintf("%+v", instance.Tags)))

			if !matchesTagFilters(instance.Tags, args.TagFilters) {
				c.logger.Warn(ctx, "❌ Instance FILTERED OUT by tag filters",
					v1.LogField("instanceID", instance.CloudID),
					v1.LogField("instanceTags", fmt.Sprintf("%+v", instance.Tags)),
					v1.LogField("requiredFilters", fmt.Sprintf("%+v", args.TagFilters)))
				continue
			}

			c.logger.Info(ctx, "✅ Instance PASSED tag filters",
				v1.LogField("instanceID", instance.CloudID))
		}

		// Apply instance ID filtering if provided
		if len(args.InstanceIDs) > 0 {
			found := false
			for _, id := range args.InstanceIDs {
				if instance.CloudID == id {
					found = true
					break
				}
			}
			if !found {
				c.logger.Debug(ctx, "instance filtered out by instance ID filter",
					v1.LogField("instanceID", instance.CloudID))
				continue
			}
		}

		// Apply location filtering if provided
		if len(args.Locations) > 0 && !args.Locations.IsAllowed(instance.Location) {
			c.logger.Debug(ctx, "instance filtered out by location filter",
				v1.LogField("instanceID", instance.CloudID),
				v1.LogField("instanceLocation", instance.Location))
			continue
		}

		c.logger.Debug(ctx, "instance passed all filters",
			v1.LogField("instanceID", instance.CloudID),
			v1.LogField("instanceTags", fmt.Sprintf("%+v", instance.Tags)))

		instances = append(instances, *instance)
	}

	c.logger.Info(ctx, "successfully listed and filtered instances",
		v1.LogField("totalFromAPI", len(allNebiusInstances)),
		v1.LogField("afterFiltering", len(instances)))

	return instances, nil
}

// matchesTagFilters checks if the instance tags match the required tag filters.
// TagFilters is a map where the key is the tag name and the value is a list of acceptable values.
// An instance matches if for every filter key, the instance has that tag and its value is in the list.
func matchesTagFilters(instanceTags map[string]string, tagFilters map[string][]string) bool {
	for filterKey, acceptableValues := range tagFilters {
		instanceValue, hasTag := instanceTags[filterKey]
		if !hasTag {
			// Instance doesn't have this required tag
			return false
		}

		// Check if the instance's tag value is in the list of acceptable values
		valueMatches := false
		for _, acceptableValue := range acceptableValues {
			if instanceValue == acceptableValue {
				valueMatches = true
				break
			}
		}

		if !valueMatches {
			// Instance has the tag but the value doesn't match any acceptable value
			return false
		}
	}

	// All filters passed
	return true
}

//nolint:dupl // StopInstance and StartInstance have similar structure but different operations
func (c *NebiusClient) StopInstance(ctx context.Context, instanceID v1.CloudProviderInstanceID) error {
	c.logger.Info(ctx, "initiating instance stop operation",
		v1.LogField("instanceID", instanceID))

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

	c.logger.Info(ctx, "stop operation completed, waiting for instance to reach STOPPED state",
		v1.LogField("instanceID", instanceID))

	// Wait for instance to actually reach STOPPED state
	// The operation completing doesn't mean the instance is fully stopped yet
	if err := c.waitForInstanceState(ctx, instanceID, v1.LifecycleStatusStopped, 3*time.Minute); err != nil {
		return fmt.Errorf("instance failed to reach STOPPED state: %w", err)
	}

	c.logger.Info(ctx, "instance successfully stopped",
		v1.LogField("instanceID", instanceID))

	return nil
}

//nolint:dupl // StartInstance and StopInstance have similar structure but different operations
func (c *NebiusClient) StartInstance(ctx context.Context, instanceID v1.CloudProviderInstanceID) error {
	c.logger.Info(ctx, "initiating instance start operation",
		v1.LogField("instanceID", instanceID))

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

	c.logger.Info(ctx, "start operation completed, waiting for instance to reach RUNNING state",
		v1.LogField("instanceID", instanceID))

	// Wait for instance to actually reach RUNNING state
	// The operation completing doesn't mean the instance is fully running yet
	if err := c.waitForInstanceState(ctx, instanceID, v1.LifecycleStatusRunning, 5*time.Minute); err != nil {
		return fmt.Errorf("instance failed to reach RUNNING state: %w", err)
	}

	c.logger.Info(ctx, "instance successfully started",
		v1.LogField("instanceID", instanceID))

	return nil
}

func (c *NebiusClient) RebootInstance(_ context.Context, _ v1.CloudProviderInstanceID) error {
	return fmt.Errorf("nebius reboot instance implementation pending: %w", v1.ErrNotImplemented)
}

func (c *NebiusClient) ChangeInstanceType(_ context.Context, _ v1.CloudProviderInstanceID, _ string) error {
	return fmt.Errorf("nebius change instance type implementation pending: %w", v1.ErrNotImplemented)
}

func (c *NebiusClient) UpdateInstanceTags(_ context.Context, _ v1.UpdateInstanceTagsArgs) error {
	return fmt.Errorf("nebius update instance tags implementation pending: %w", v1.ErrNotImplemented)
}

func (c *NebiusClient) ResizeInstanceVolume(_ context.Context, _ v1.ResizeInstanceVolumeArgs) error {
	return fmt.Errorf("nebius resize instance volume implementation pending: %w", v1.ErrNotImplemented)
}

func (c *NebiusClient) AddFirewallRulesToInstance(_ context.Context, _ v1.AddFirewallRulesToInstanceArgs) error {
	return fmt.Errorf("nebius firewall rules management not yet implemented: %w", v1.ErrNotImplemented)
}

func (c *NebiusClient) RevokeSecurityGroupRules(_ context.Context, _ v1.RevokeSecurityGroupRuleArgs) error {
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
	if attrs.DiskSize == 0 {
		attrs.DiskSize = 1280 * units.Gibibyte // Defaulted by the Nebius Console
	}

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
//
//nolint:gocognit,gocyclo // Complex function trying multiple image resolution strategies
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
			//nolint:gocritic // if-else chain is clearer than switch for version matching logic
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
//
//nolint:gocognit,gocyclo,funlen // Complex function with multiple fallback strategies for parsing instance types
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
		// Region could be multi-part (eu-north1) so we need to find the GPU type or platformTypeCPU
		var gpuType string
		var presetStartIdx int

		// Look for GPU type indicators or platformTypeCPU
		for i := 1; i < len(parts); i++ {
			partLower := strings.ToLower(parts[i])
			// Check if this part is a known GPU type or platformTypeCPU
			if partLower == platformTypeCPU || partLower == "l40s" || partLower == "h100" ||
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
				if (gpuType == platformTypeCPU && strings.Contains(platformNameLower, platformTypeCPU)) ||
					(gpuType != platformTypeCPU && strings.Contains(platformNameLower, gpuType)) {
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
//
//nolint:gocyclo,unparam // Complex image family resolution with fallback logic
func (c *NebiusClient) resolveImageFamily(ctx context.Context, imageID string) (string, error) {
	// Common Nebius image families - if ImageID matches one of these, use it directly
	commonFamilies := []string{
		"ubuntu22.04-cuda12",
		"mk8s-worker-node-v-1-32-ubuntu24.04",
		"mk8s-worker-node-v-1-32-ubuntu24.04-cuda12.8",
		"mk8s-worker-node-v-1-31-ubuntu24.04-cuda12",
		"ubuntu22.04",
		"ubuntu20.04",
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

// getBootDiskSize queries a boot disk and returns its size in bytes
func (c *NebiusClient) getBootDiskSize(ctx context.Context, diskID string) (int64, error) {
	disk, err := c.sdk.Services().Compute().V1().Disk().Get(ctx, &compute.GetDiskRequest{
		Id: diskID,
	})
	if err != nil {
		return 0, fmt.Errorf("failed to get disk details: %w", err)
	}

	if disk.Spec == nil {
		return 0, fmt.Errorf("disk spec is nil")
	}

	// Extract size from the Size oneof field
	if sizeGiB, ok := disk.Spec.Size.(*compute.DiskSpec_SizeGibibytes); ok {
		// Convert GiB to bytes
		return sizeGiB.SizeGibibytes * int64(units.Gibibyte), nil
	}

	return 0, fmt.Errorf("disk size not available")
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
				// Continue on error - don't fail the entire cleanup
				continue
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

	var commands []string
	// Generate UFW firewall commands (similar to Shadeform's approach)
	// UFW (Uncomplicated Firewall) is available on Ubuntu/Debian instances
	commands = append(commands, generateUFWCommands(firewallRules)...)

	// Generate IPTables firewall commands to ensure docker ports are not made immediately
	// accessible from the internet by default.
	commands = append(commands, generateIPTablesCommands()...)

	if len(commands) > 0 {
		// Use runcmd to execute firewall setup commands
		script += "\nruncmd:\n"
		for _, cmd := range commands {
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

// generateIPTablesCommands generates IPTables firewall commands to ensure docker ports are not made immediately
// accessible from the internet by default.
func generateIPTablesCommands() []string {
	commands := []string{
		"iptables -I DOCKER-USER -i lo -j ACCEPT",
		"iptables -I DOCKER-USER -m conntrack --ctstate ESTABLISHED,RELATED -j ACCEPT",
		"iptables -A DOCKER-USER -j DROP",
	}
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
