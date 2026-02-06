package v1

import (
	"context"
	"encoding/base64"
	"fmt"
	"slices"
	"strings"
	"time"

	"github.com/alecthomas/units"
	"github.com/brevdev/cloud/internal/errors"
	v1 "github.com/brevdev/cloud/v1"
	sfcnodes "github.com/sfcompute/nodes-go"
	"github.com/sfcompute/nodes-go/packages/param"
)

const (
	maxPricePerNodeHour = 1600
	defaultPort         = 2222
	defaultSSHUsername  = "ubuntu"
)

func (c *SFCClient) CreateInstance(ctx context.Context, attrs v1.CreateInstanceAttrs) (*v1.Instance, error) {
	// Get the zone for the location (do not include unavailable zones)
	zone, err := c.getZone(ctx, attrs.Location, false)
	if err != nil {
		return nil, errors.WrapAndTrace(err)
	}

	// Create a name for the node
	name := brevDataToSFCName(attrs.RefID, attrs.Name)

	// Create the node
	resp, err := c.client.Nodes.New(ctx, sfcnodes.NodeNewParams{
		CreateNodesRequest: sfcnodes.CreateNodesRequestParam{
			DesiredCount:        1,
			MaxPricePerNodeHour: maxPricePerNodeHour,
			Zone:                zone.Name,
			Names:               []string{name},
			CloudInitUserData:   param.Opt[string]{Value: sshKeyCloudInit(attrs.PublicKey)},
		},
	})
	if err != nil {
		return nil, errors.WrapAndTrace(err)
	}
	if len(resp.Data) == 0 {
		return nil, errors.WrapAndTrace(fmt.Errorf("no nodes returned"))
	}
	node := resp.Data[0]

	// Get the instance
	instance, err := c.GetInstance(ctx, v1.CloudProviderInstanceID(node.ID))
	if err != nil {
		return nil, errors.WrapAndTrace(err)
	}

	return instance, nil
}

func sshKeyCloudInit(sshKey string) string {
	script := fmt.Sprintf("#cloud-config\nssh_authorized_keys:\n  - %s", sshKey)
	return base64.StdEncoding.EncodeToString([]byte(script))
}

func (c *SFCClient) GetInstance(ctx context.Context, id v1.CloudProviderInstanceID) (*v1.Instance, error) {
	c.logger.Debug(ctx, "sfc: GetInstance start",
		v1.LogField("instanceID", id),
		v1.LogField("location", c.location),
	)

	// Get the node from the API
	node, err := c.client.Nodes.Get(ctx, string(id))
	if err != nil {
		return nil, errors.WrapAndTrace(err)
	}

	// Get the zone for the location (include unavailable zones, in case the zone is not available but the node is still running)
	zone, err := c.getZone(ctx, node.Zone, true)
	if err != nil {
		return nil, errors.WrapAndTrace(err)
	}

	nodeInfo, err := c.sfcNodeInfoFromNode(ctx, node, zone)
	if err != nil {
		return nil, errors.WrapAndTrace(err)
	}

	instance, err := c.sfcNodeToBrevInstance(*nodeInfo)
	if err != nil {
		return nil, errors.WrapAndTrace(err)
	}

	c.logger.Debug(ctx, "sfc: GetInstance end",
		v1.LogField("instanceID", id),
		v1.LogField("instance", instance),
	)

	return instance, nil
}

func (c *SFCClient) getZone(ctx context.Context, location string, includeUnavailable bool) (*sfcnodes.ZoneListResponseData, error) {
	// Fetch the zones to ensure the location is valid
	zones, err := c.getZones(ctx, includeUnavailable)
	if err != nil {
		return nil, errors.WrapAndTrace(err)
	}
	if len(zones) == 0 {
		return nil, errors.WrapAndTrace(fmt.Errorf("no zones available"))
	}

	// Find the zone that matches the location
	var zone *sfcnodes.ZoneListResponseData
	for _, z := range zones {
		if z.Name == location {
			zone = &z
			break
		}
	}
	if zone == nil {
		return nil, errors.WrapAndTrace(fmt.Errorf("zone not found in location %s", location))
	}

	return zone, nil
}

func (c *SFCClient) ListInstances(ctx context.Context, args v1.ListInstancesArgs) ([]v1.Instance, error) {
	c.logger.Debug(ctx, "sfc: ListInstances start",
		v1.LogField("location", c.location),
		v1.LogField("args", fmt.Sprintf("%+v", args)),
	)

	resp, err := c.client.Nodes.List(ctx, sfcnodes.NodeListParams{})
	if err != nil {
		return nil, errors.WrapAndTrace(err)
	}

	c.logger.Debug(ctx, "sfc: ListInstances nodes list",
		v1.LogField("node count", len(resp.Data)),
	)

	zoneCache := make(map[string]*sfcnodes.ZoneListResponseData)

	var instances []v1.Instance
	for _, node := range resp.Data {
		// Get the zone for the node, checking the cache first
		zone, ok := zoneCache[node.Zone]
		if !ok {
			z, err := c.getZone(ctx, node.Zone, true)
			if err != nil {
				return nil, errors.WrapAndTrace(err)
			}
			zoneCache[node.Zone] = z
			zone = z
		}

		// Filter by locations
		if args.Locations != nil && !args.Locations.IsAllowed(zone.Name) {
			c.logger.Debug(ctx, "sfc: ListInstances node filtered out by location",
				v1.LogField("nodeID", node.ID),
				v1.LogField("location", zone.Name),
			)
			continue
		}

		// Filter by instance IDs
		if len(args.InstanceIDs) > 0 && !slices.Contains(args.InstanceIDs, v1.CloudProviderInstanceID(node.ID)) {
			c.logger.Debug(ctx, "sfc: ListInstances node filtered out by instance ID",
				v1.LogField("nodeID", node.ID),
				v1.LogField("instanceID", v1.CloudProviderInstanceID(node.ID)),
			)
			continue
		}

		nodeInfo, err := c.sfcNodeInfoFromNodeListResponseData(ctx, &node, zone)
		if err != nil {
			return nil, errors.WrapAndTrace(err)
		}

		inst, err := c.sfcNodeToBrevInstance(*nodeInfo)
		if err != nil {
			return nil, errors.WrapAndTrace(err)
		}
		instances = append(instances, *inst)
	}

	c.logger.Debug(ctx, "sfc: ListInstances end",
		v1.LogField("instance count", len(instances)),
	)

	return instances, nil
}

func (c *SFCClient) TerminateInstance(ctx context.Context, id v1.CloudProviderInstanceID) error {
	c.logger.Debug(ctx, "sfc: TerminateInstance start",
		v1.LogField("instanceID", id),
	)

	_, err := c.client.Nodes.Release(ctx, string(id))
	if err != nil {
		return errors.WrapAndTrace(err)
	}

	c.logger.Debug(ctx, "sfc: TerminateInstance end",
		v1.LogField("instanceID", id),
	)

	return nil
}

type sfcNodeInfo struct {
	id          string
	name        string
	createdAt   time.Time
	status      v1.LifecycleStatus
	gpuType     string
	sshUsername string
	sshHostname string
	zone        *sfcnodes.ZoneListResponseData
}

func (c *SFCClient) sfcNodeToBrevInstance(node sfcNodeInfo) (*v1.Instance, error) {
	// Get the refID and name from the node name
	refID, name, err := sfcNameToBrevData(node.name)
	if err != nil {
		return nil, errors.WrapAndTrace(err)
	}

	// Get the instance type for the zone
	instanceType, err := getInstanceTypeForZone(*node.zone)
	if err != nil {
		return nil, errors.WrapAndTrace(err)
	}

	diskSizeInt64, err := instanceType.SupportedStorage[0].SizeBytes.ByteCountInUnitInt64(v1.Gibibyte)
	if err != nil {
		return nil, err
	}
	diskSize := units.Base2Bytes(diskSizeInt64 * int64(units.Gibibyte))

	// Create the instance
	inst := &v1.Instance{
		Name:          name,
		CloudID:       v1.CloudProviderInstanceID(node.id),
		RefID:         refID,
		PublicDNS:     node.sshHostname,
		PublicIP:      node.sshHostname,
		SSHUser:       node.sshUsername,
		SSHPort:       defaultPort,
		CreatedAt:     node.createdAt,
		DiskSize:      diskSize,
		DiskSizeBytes: instanceType.SupportedStorage[0].SizeBytes, // TODO: this should be pulled from the node itself
		Status: v1.Status{
			LifecycleStatus: node.status,
		},
		InstanceTypeID: instanceType.ID,
		InstanceType:   instanceType.Type,
		Location:       node.zone.Name,
		Spot:           false,
		Stoppable:      false,
		Rebootable:     false,
		CloudCredRefID: c.refID, // TODO: this should be pulled from the node itself
	}
	return inst, nil
}

func (c *SFCClient) sfcNodeInfoFromNode(ctx context.Context, node *sfcnodes.Node, zone *sfcnodes.ZoneListResponseData) (*sfcNodeInfo, error) {
	var sshHostname string
	if len(node.VMs.Data) == 1 { //nolint:gocritic // ok
		hostname, err := c.getSSHHostnameFromVM(ctx, node.VMs.Data[0].ID, node.VMs.Data[0].Status)
		if err != nil {
			return nil, errors.WrapAndTrace(err)
		}
		sshHostname = hostname
	} else if len(node.VMs.Data) == 0 {
		sshHostname = ""
	} else {
		return nil, errors.WrapAndTrace(fmt.Errorf("multiple VMs found for node %s", node.ID))
	}

	return &sfcNodeInfo{
		id:          node.ID,
		name:        node.Name,
		createdAt:   time.Unix(node.CreatedAt, 0),
		status:      sfcStatusToLifecycleStatus(fmt.Sprint(node.Status)),
		gpuType:     string(node.GPUType),
		sshUsername: defaultSSHUsername,
		sshHostname: sshHostname,
		zone:        zone,
	}, nil
}

func (c *SFCClient) sfcNodeInfoFromNodeListResponseData(ctx context.Context, node *sfcnodes.ListResponseNodeData, zone *sfcnodes.ZoneListResponseData) (*sfcNodeInfo, error) {
	sfcNode := sfcListResponseNodeDataToNode(node)
	return c.sfcNodeInfoFromNode(ctx, sfcNode, zone)
}

// Convert the sfcnodes.ListResponseNodeData into a node *sfcnodes.Node -- these are fundamentally the same object, but they
// lack a common interface. One type is returned from a single "get" call, the other is the type of each object returned by
// a "list" call. This conversion function allows the rest of our business logic to treat these as the same type.
func sfcListResponseNodeDataToNode(node *sfcnodes.ListResponseNodeData) *sfcnodes.Node {
	vms := make([]sfcnodes.NodeVMsData, len(node.VMs.Data))
	for i, vm := range node.VMs.Data {
		vms[i] = sfcnodes.NodeVMsData{ //nolint:staticcheck // ok
			ID:        vm.ID,
			CreatedAt: vm.CreatedAt,
			EndAt:     vm.EndAt,
			Object:    vm.Object,
			StartAt:   vm.StartAt,
			Status:    vm.Status,
			UpdatedAt: vm.UpdatedAt,
			ImageID:   vm.ImageID,
			JSON:      vm.JSON,
		}
	}

	return &sfcnodes.Node{
		ID:                  node.ID,
		GPUType:             node.GPUType,
		Name:                node.Name,
		NodeType:            node.NodeType,
		Object:              node.Object,
		Owner:               node.Owner,
		Status:              node.Status,
		CreatedAt:           node.CreatedAt,
		DeletedAt:           node.DeletedAt,
		EndAt:               node.EndAt,
		MaxPricePerNodeHour: node.MaxPricePerNodeHour,
		ProcurementID:       node.ProcurementID,
		StartAt:             node.StartAt,
		UpdatedAt:           node.UpdatedAt,
		Zone:                node.Zone,
		JSON:                node.JSON,
		VMs: sfcnodes.NodeVMs{
			Data:   vms,
			Object: node.VMs.Object,
			JSON:   node.VMs.JSON,
		},
	}
}

func sfcStatusToLifecycleStatus(status string) v1.LifecycleStatus {
	switch strings.ToLower(status) {
	case "pending", "unspecified", "awaitingcapacity", "unknown":
		return v1.LifecycleStatusPending
	case "running":
		return v1.LifecycleStatusRunning
	case "stopped":
		return v1.LifecycleStatusStopped
	case "terminating":
		return v1.LifecycleStatusTerminating
	case "released", "destroyed", "deleted":
		return v1.LifecycleStatusTerminated
	case "nodefailure", "failed":
		return v1.LifecycleStatusFailed
	default:
		return v1.LifecycleStatusPending
	}
}

func (c *SFCClient) getSSHHostnameFromVM(ctx context.Context, vmID string, vmStatus string) (string, error) {
	// If the VM is not running, set the SSH username and hostname to empty strings
	if strings.ToLower(vmStatus) != "running" {
		return "", nil
	}

	// If the VM is running, get the SSH username and hostname
	sshResponse, err := c.client.VMs.SSH(ctx, sfcnodes.VMSSHParams{VMID: vmID})
	if err != nil {
		return "", errors.WrapAndTrace(err)
	}

	return sshResponse.SSHHostname, nil
}

func brevDataToSFCName(refID string, name string) string {
	return fmt.Sprintf("%s_%s", refID, name)
}

func sfcNameToBrevData(name string) (string, string, error) {
	parts := strings.SplitN(name, "_", 2)
	if len(parts) != 2 {
		return "", "", errors.WrapAndTrace(fmt.Errorf("invalid node name %s", name))
	}
	return parts[0], parts[1], nil
}

// Optional if supported:
func (c *SFCClient) RebootInstance(_ context.Context, _ v1.CloudProviderInstanceID) error {
	return v1.ErrNotImplemented
}

func (c *SFCClient) StopInstance(_ context.Context, _ v1.CloudProviderInstanceID) error {
	return v1.ErrNotImplemented
}

func (c *SFCClient) StartInstance(_ context.Context, _ v1.CloudProviderInstanceID) error {
	return v1.ErrNotImplemented
}

// Merge strategies (pass-through is acceptable baseline).
func (c *SFCClient) MergeInstanceForUpdate(_ v1.Instance, newInst v1.Instance) v1.Instance {
	return newInst
}

func (c *SFCClient) MergeInstanceTypeForUpdate(_ v1.InstanceType, newIt v1.InstanceType) v1.InstanceType {
	return newIt
}
