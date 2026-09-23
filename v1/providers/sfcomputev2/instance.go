package v2

import (
	"context"
	"encoding/base64"
	"fmt"
	"maps"
	"net/http"
	"regexp"
	"slices"
	"time"

	"github.com/alecthomas/units"
	"github.com/brevdev/cloud/internal/errors"
	v1 "github.com/brevdev/cloud/v1"
)

// SFC instance names must match `[a-zA-Z0-9][a-zA-Z0-9._-]{0,254}`: start with an
// alphanumeric character, then alphanumerics/dot/underscore/hyphen, max 255 chars.
var (
	sfcNamePattern    = regexp.MustCompile(`^[a-zA-Z0-9][a-zA-Z0-9._-]{0,254}$`)
	sfcNameDisallowed = regexp.MustCompile(`[^a-zA-Z0-9._-]`)
	sfcNameLeading    = regexp.MustCompile(`^[^a-zA-Z0-9]+`)
)

// sanitizeSFCName coerces a requested instance name into SFC's required format:
// disallowed characters are replaced with '-', leading non-alphanumeric characters are
// dropped (SFC requires an alphanumeric first character), and the result is truncated to
// the 255-char max. Returns "" if no usable characters remain.
func sanitizeSFCName(name string) string {
	name = sfcNameDisallowed.ReplaceAllString(name, "-")
	name = sfcNameLeading.ReplaceAllString(name, "")
	if len(name) > 255 {
		name = name[:255]
	}
	return name
}

func makeSFCName(refID string, tags v1.Tags) string {
	return sanitizeSFCName(refID + "-" + tags["dev-plane-x-environmentId"])
}

func (c *SFCClientV2) CreateInstance(ctx context.Context, attrs v1.CreateInstanceAttrs) (*v1.Instance, error) {
	c.logger.Debug(ctx, "sfcv2: CreateInstance start",
		v1.LogField("name", attrs.Name),
		v1.LogField("location", attrs.Location),
	)

	tags := make(map[string]string, len(attrs.Tags)+3)
	maps.Copy(tags, attrs.Tags)
	tags[tagKeyCloudCredRefID] = c.refID
	tags[tagKeyRefID] = attrs.RefID

	// Spread instances across every SKU in the capacity rather than piling onto one.
	sku, err := c.selectAvailableSku(ctx)
	if err != nil {
		return nil, errors.WrapAndTrace(err)
	}

	cloudInit := sshKeyCloudInit(attrs.PublicKey)
	req := createInstanceRequest{
		Pool:              c.GetDefaultPoolResourcePath(),
		Image:             c.GetDefaultImageResourcePath(),
		InstanceSKU:       sku,
		CloudInitUserData: &cloudInit,
		Tags:              tags,
	}
	if name := makeSFCName(attrs.RefID, attrs.Tags); sfcNamePattern.MatchString(name) {
		req.Name = &name
	}
	if c.enableConfigurableFirewall {
		firewall, err := c.createInstanceFirewall(ctx, attrs.FirewallRules)
		if err != nil {
			return nil, errors.WrapAndTrace(err)
		}
		req.EnablePublicIPv4 = true
		req.Firewall = firewall.ID
		req.Tags[tagKeyFirewallID] = firewall.ID
	}
	resp, err := c.client.createInstance(ctx, req)
	if err != nil {
		if req.Firewall != "" {
			err = c.cleanupFirewall(ctx, req.Firewall, err)
		}
		return nil, errors.WrapAndTrace(err)
	}
	if resp == nil {
		return nil, errors.WrapAndTrace(fmt.Errorf("no instance returned from create"))
	}
	if req.EnablePublicIPv4 && (!resp.EnablePublicIPv4 || resp.Firewall != req.Firewall) {
		err := fmt.Errorf("SFCompute API did not accept the requested public networking")
		cleanupCtx, cancel := context.WithTimeout(context.WithoutCancel(ctx), 30*time.Second)
		defer cancel()
		if terminateErr := c.client.terminateInstance(cleanupCtx, resp.ID); terminateErr != nil && !isAPIStatus(terminateErr, http.StatusNotFound) {
			return nil, errors.WrapAndTrace(fmt.Errorf("%w; instance cleanup failed: %w", err, terminateErr))
		}
		return nil, errors.WrapAndTrace(c.cleanupFirewall(ctx, req.Firewall, err))
	}

	instance, err := c.sfcInstanceToBrevInstance(resp, nil)
	if err != nil {
		return nil, errors.WrapAndTrace(err)
	}

	c.logger.Debug(ctx, "sfcv2: CreateInstance end",
		v1.LogField("instanceID", resp.ID),
		v1.LogField("instanceSku", sku),
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

	resp, err := c.client.getInstance(ctx, string(id))
	if err != nil {
		return nil, errors.WrapAndTrace(err)
	}
	if resp == nil {
		return nil, errors.WrapAndTrace(fmt.Errorf("instance %s not found", id))
	}

	var sshInfo *instanceSSHInfo
	if !resp.EnablePublicIPv4 {
		sshInfo, err = c.getSSHInfo(ctx, string(id), resp.Status)
		if err != nil {
			return nil, errors.WrapAndTrace(err)
		}
	}

	instance, err := c.sfcInstanceToBrevInstance(resp, sshInfo)
	if err != nil {
		return nil, errors.WrapAndTrace(err)
	}
	if err := c.loadInstanceFirewall(ctx, resp, instance); err != nil {
		return nil, errors.WrapAndTrace(err)
	}

	c.logger.Debug(ctx, "sfcv2: GetInstance end",
		v1.LogField("instanceID", id),
		v1.LogField("status", resp.Status),
	)

	return instance, nil
}

func (c *SFCClientV2) ListInstances(ctx context.Context, args v1.ListInstancesArgs) ([]v1.Instance, error) {
	c.logger.Debug(ctx, "sfcv2: ListInstances start",
		v1.LogField("location", c.location),
	)

	poolID := c.GetDefaultPoolResourcePath()
	resp, err := c.client.listInstances(ctx, c.GetWorkspaceResourcePath(), poolID)
	if err != nil {
		return nil, errors.WrapAndTrace(err)
	}
	if resp == nil {
		return []v1.Instance{}, nil
	}

	var instances []v1.Instance
	for _, inst := range resp.Data {
		// Filter by instance IDs if specified.
		if len(args.InstanceIDs) > 0 && !slices.Contains(args.InstanceIDs, v1.CloudProviderInstanceID(inst.ID)) {
			continue
		}

		var sshInfo *instanceSSHInfo
		var err error
		if !inst.EnablePublicIPv4 {
			sshInfo, err = c.getSSHInfo(ctx, inst.ID, inst.Status)
		}
		if err != nil {
			c.logger.Error(ctx, err,
				v1.LogField("msg", "sfcv2: ListInstances skipping instance due to SSH error"),
				v1.LogField("instanceID", inst.ID),
			)
			continue
		}

		brevInst, err := c.sfcInstanceToBrevInstance(&inst, sshInfo)
		if err != nil {
			c.logger.Error(ctx, err,
				v1.LogField("msg", "sfcv2: ListInstances skipping instance due to conversion error"),
				v1.LogField("instanceID", inst.ID),
			)
			continue
		}
		if err := c.loadInstanceFirewall(ctx, &inst, brevInst); err != nil {
			return nil, errors.WrapAndTrace(err)
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

	instance, err := c.client.terminateInstanceWithResponse(ctx, string(id))
	if err != nil {
		return normalizeTerminateInstanceError(err)
	}
	if instance.EnablePublicIPv4 && instance.Firewall != "" && instance.Tags[tagKeyFirewallID] == instance.Firewall {
		return c.cleanupFirewall(ctx, instance.Firewall, nil)
	}

	c.logger.Debug(ctx, "sfcv2: TerminateInstance end",
		v1.LogField("instanceID", id),
	)

	return nil
}

func normalizeTerminateInstanceError(err error) error {
	var responseErr *apiError
	if errors.As(err, &responseErr) && responseErr.statusCode == http.StatusNotFound {
		// Termination is idempotent: a missing instance is already in the
		// requested terminal state. Do not retain the provider response body.
		return errors.WrapAndTrace(v1.ErrInstanceNotFound)
	}
	return errors.WrapAndTrace(err)
}

func (c *SFCClientV2) getSSHInfo(ctx context.Context, id string, status instanceStatus) (*instanceSSHInfo, error) {
	if status != instanceStatusRunning {
		return nil, nil
	}

	resp, err := c.client.getSSHInfo(ctx, id)
	if err != nil {
		return nil, errors.WrapAndTrace(err)
	}
	if resp == nil {
		return nil, nil
	}

	return resp, nil
}

func (c *SFCClientV2) sfcInstanceToBrevInstance(inst *instanceResponse, sshInfo *instanceSSHInfo) (*v1.Instance, error) {
	tags := inst.Tags

	cloudCredRefID := tags[tagKeyCloudCredRefID]
	if cloudCredRefID == "" {
		cloudCredRefID = c.refID
	}

	userTags := make(v1.Tags)
	for k, v := range tags {
		switch k {
		case tagKeyCloudCredRefID, tagKeyRefID, tagKeyFirewallID:
		default:
			userTags[k] = v
		}
	}

	status := sfcStatusToLifecycleStatus(inst.Status)
	hostname, sshPort := sshInfo.GetHostname(), int(sshInfo.GetPort())
	if inst.EnablePublicIPv4 {
		hostname, sshPort = inst.PublicIP, 22
		if hostname == "" && status == v1.LifecycleStatusRunning {
			status = v1.LifecycleStatusPending
		}
	}

	diskInt64, err := h100InstanceTypeMetadata.diskBytes.ByteCountInUnitInt64(v1.Gibibyte)
	if err != nil {
		return nil, err
	}
	diskSize := units.Base2Bytes(diskInt64 * int64(units.Gibibyte))

	return &v1.Instance{
		Name:          inst.Name,
		CloudID:       v1.CloudProviderInstanceID(inst.ID),
		RefID:         tags[tagKeyRefID],
		PublicDNS:     hostname,
		PublicIP:      hostname,
		SSHUser:       defaultSSHUsername,
		SSHPort:       sshPort,
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

func sfcStatusToLifecycleStatus(status instanceStatus) v1.LifecycleStatus {
	switch status {
	case instanceStatusAwaitingAllocation:
		return v1.LifecycleStatusPending
	case instanceStatusRunning:
		return v1.LifecycleStatusRunning
	case instanceStatusTerminated:
		return v1.LifecycleStatusTerminated
	case instanceStatusFailed:
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
