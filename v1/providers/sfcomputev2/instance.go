package v2

import (
	"context"
	"encoding/base64"
	"fmt"
	"maps"
	"slices"
	"time"

	"github.com/alecthomas/units"
	"github.com/brevdev/cloud/internal/errors"
	v1 "github.com/brevdev/cloud/v1"
	"github.com/sfcompute/sfc-go/models/components"
	"github.com/sfcompute/sfc-go/models/operations"
	"github.com/sfcompute/sfc-go/optionalnullable"
)

func (c *SFCClientV2) CreateInstance(ctx context.Context, attrs v1.CreateInstanceAttrs) (*v1.Instance, error) {
	c.logger.Debug(ctx, "sfcv2: CreateInstance start",
		v1.LogField("name", attrs.Name),
		v1.LogField("location", attrs.Location),
	)

	tags := make(map[string]string, len(attrs.Tags)+2)
	maps.Copy(tags, attrs.Tags)
	tags[tagKeyCloudCredRefID] = c.refID
	tags[tagKeyRefID] = attrs.RefID

	cloudInit := sshKeyCloudInit(attrs.PublicKey)
	resp, err := c.client.Instances.Create(ctx, components.CreateInstanceRequest{
		Capacity:          GetDefaultCapacityID(c.workspace),
		Image:             GetDefaultImageID(),
		CloudInitUserData: &cloudInit,
		Tags:              optionalnullable.From(&tags),
		Name:              optionalnullable.From(&attrs.Name),
	})
	if err != nil {
		return nil, errors.WrapAndTrace(err)
	}
	if resp.InstanceResponse == nil {
		return nil, errors.WrapAndTrace(fmt.Errorf("no instance returned from create"))
	}

	instance, err := c.sfcInstanceToBrevInstance(resp.InstanceResponse, "")
	if err != nil {
		return nil, errors.WrapAndTrace(err)
	}

	c.logger.Debug(ctx, "sfcv2: CreateInstance end",
		v1.LogField("instanceID", resp.InstanceResponse.ID),
	)

	return instance, nil
}

func sshKeyCloudInit(sshKey string) string {
	script := fmt.Sprintf("#cloud-config\nssh_authorized_keys:\n  - %s", sshKey)
	return base64.StdEncoding.EncodeToString([]byte(script))
}

func (c *SFCClientV2) GetInstance(ctx context.Context, id v1.CloudProviderInstanceID) (*v1.Instance, error) {
	c.logger.Debug(ctx, "sfcv2: GetInstance start",
		v1.LogField("instanceID", id),
	)

	resp, err := c.client.Instances.Fetch(ctx, string(id), nil)
	if err != nil {
		return nil, errors.WrapAndTrace(err)
	}
	if resp.InstanceResponse == nil {
		return nil, errors.WrapAndTrace(fmt.Errorf("instance %s not found", id))
	}

	sshHostname, err := c.getSSHHostname(ctx, string(id), resp.InstanceResponse.Status)
	if err != nil {
		return nil, errors.WrapAndTrace(err)
	}

	instance, err := c.sfcInstanceToBrevInstance(resp.InstanceResponse, sshHostname)
	if err != nil {
		return nil, errors.WrapAndTrace(err)
	}

	c.logger.Debug(ctx, "sfcv2: GetInstance end",
		v1.LogField("instanceID", id),
		v1.LogField("status", resp.InstanceResponse.Status),
	)

	return instance, nil
}

func (c *SFCClientV2) ListInstances(ctx context.Context, args v1.ListInstancesArgs) ([]v1.Instance, error) {
	c.logger.Debug(ctx, "sfcv2: ListInstances start",
		v1.LogField("location", c.location),
	)

	capacityID := GetDefaultCapacityID(c.workspace)
	resp, err := c.client.Instances.List(ctx, operations.ListInstancesRequest{
		Capacity: &capacityID,
	})
	if err != nil {
		return nil, errors.WrapAndTrace(err)
	}
	if resp.ListInstancesResponse == nil {
		return []v1.Instance{}, nil
	}

	var instances []v1.Instance
	for _, inst := range resp.ListInstancesResponse.Data {
		// Filter by instance IDs if specified.
		if len(args.InstanceIDs) > 0 && !slices.Contains(args.InstanceIDs, v1.CloudProviderInstanceID(inst.ID)) {
			continue
		}

		sshHostname, err := c.getSSHHostname(ctx, inst.ID, inst.Status)
		if err != nil {
			c.logger.Error(ctx, err,
				v1.LogField("msg", "sfcv2: ListInstances skipping instance due to SSH error"),
				v1.LogField("instanceID", inst.ID),
			)
			continue
		}

		brevInst, err := c.sfcInstanceToBrevInstance(&inst, sshHostname)
		if err != nil {
			c.logger.Error(ctx, err,
				v1.LogField("msg", "sfcv2: ListInstances skipping instance due to conversion error"),
				v1.LogField("instanceID", inst.ID),
			)
			continue
		}
		instances = append(instances, *brevInst)
	}

	c.logger.Debug(ctx, "sfcv2: ListInstances end",
		v1.LogField("instance count", len(instances)),
	)

	return instances, nil
}

func (c *SFCClientV2) TerminateInstance(ctx context.Context, id v1.CloudProviderInstanceID) error {
	c.logger.Debug(ctx, "sfcv2: TerminateInstance start",
		v1.LogField("instanceID", id),
	)

	_, err := c.client.Instances.TerminateInstance(ctx, string(id))
	if err != nil {
		return errors.WrapAndTrace(err)
	}

	c.logger.Debug(ctx, "sfcv2: TerminateInstance end",
		v1.LogField("instanceID", id),
	)

	return nil
}

func (c *SFCClientV2) getSSHHostname(ctx context.Context, id string, status components.InstanceStatus) (string, error) {
	if status != components.InstanceStatusRunning {
		return "", nil
	}

	resp, err := c.client.Instances.GetSSHInfoForInstance(ctx, id)
	if err != nil {
		return "", errors.WrapAndTrace(err)
	}
	if resp.InstanceSSHInfo == nil {
		return "", nil
	}

	return resp.InstanceSSHInfo.Hostname, nil
}

func (c *SFCClientV2) sfcInstanceToBrevInstance(inst *components.InstanceResponse, sshHostname string) (*v1.Instance, error) {
	tags, _ := inst.GetTags().GetOrZero()

	cloudCredRefID := tags[tagKeyCloudCredRefID]
	if cloudCredRefID == "" {
		cloudCredRefID = c.refID
	}

	userTags := make(v1.Tags)
	for k, v := range tags {
		switch k {
		case tagKeyCloudCredRefID, tagKeyRefID:
		default:
			userTags[k] = v
		}
	}

	status := sfcStatusToLifecycleStatus(inst.Status)

	diskInt64, err := h100InstanceTypeMetadata.diskBytes.ByteCountInUnitInt64(v1.Gibibyte)
	if err != nil {
		return nil, err
	}
	diskSize := units.Base2Bytes(diskInt64 * int64(units.Gibibyte))

	return &v1.Instance{
		Name:          inst.Name,
		CloudID:       v1.CloudProviderInstanceID(inst.ID),
		RefID:         tags[tagKeyRefID],
		PublicDNS:     sshHostname,
		PublicIP:      sshHostname,
		SSHUser:       defaultSSHUsername,
		SSHPort:       defaultPort,
		CreatedAt:     time.Unix(inst.CreatedAt, 0),
		DiskSize:      diskSize,
		DiskSizeBytes: h100InstanceTypeMetadata.diskBytes,
		Status: v1.Status{
			LifecycleStatus: status,
		},
		InstanceTypeID: h100InstanceTypeMetadata.instanceTypeID,
		InstanceType:   h100InstanceType,
		Location:       sfcLocation,
		Spot:           false,
		Stoppable:      false,
		Rebootable:     false,
		CloudCredRefID: cloudCredRefID,
		Tags:           userTags,
	}, nil
}

func sfcStatusToLifecycleStatus(status components.InstanceStatus) v1.LifecycleStatus {
	switch status {
	case components.InstanceStatusAwaitingAllocation:
		return v1.LifecycleStatusPending
	case components.InstanceStatusRunning:
		return v1.LifecycleStatusRunning
	case components.InstanceStatusTerminated:
		return v1.LifecycleStatusTerminated
	case components.InstanceStatusFailed:
		return v1.LifecycleStatusFailed
	default:
		return v1.LifecycleStatusPending
	}
}

func (c *SFCClientV2) RebootInstance(_ context.Context, _ v1.CloudProviderInstanceID) error {
	return v1.ErrNotImplemented
}

func (c *SFCClientV2) StopInstance(_ context.Context, _ v1.CloudProviderInstanceID) error {
	return v1.ErrNotImplemented
}

func (c *SFCClientV2) StartInstance(_ context.Context, _ v1.CloudProviderInstanceID) error {
	return v1.ErrNotImplemented
}

func (c *SFCClientV2) MergeInstanceForUpdate(_ v1.Instance, newInst v1.Instance) v1.Instance {
	return newInst
}

func (c *SFCClientV2) MergeInstanceTypeForUpdate(_ v1.InstanceType, newIt v1.InstanceType) v1.InstanceType {
	return newIt
}
