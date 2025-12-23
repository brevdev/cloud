package v1

import (
	"context"
	"fmt"

	"github.com/brevdev/cloud/v1"
	openapi "github.com/brevdev/cloud/v1/providers/fluidstack/gen/fluidstack"
)

func (c *FluidStackClient) CreateInstance(ctx context.Context, attrs v1.CreateInstanceAttrs) (*v1.Instance, error) {
	authCtx := c.makeAuthContext(ctx)
	projectCtx := c.makeProjectContext(authCtx)

	req := openapi.InstancesPostRequest{
		Name: attrs.Name,
		Type: attrs.InstanceType,
	}

	if attrs.UserDataBase64 != "" {
		req.SetUserData(attrs.UserDataBase64)
	}

	if len(attrs.Tags) > 0 {
		tags := make(map[string]string)
		for k, v := range attrs.Tags {
			tags[k] = v
		}
		req.SetTags(tags)
	}

	resp, httpResp, err := c.client.InstancesAPI.CreateInstance(projectCtx).XPROJECTID(c.projectID).InstancesPostRequest(req).Execute()
	if httpResp != nil && httpResp.Body != nil {
		defer func() { _ = httpResp.Body.Close() }()
	}
	if err != nil {
		return nil, fmt.Errorf("failed to create instance: %w", err)
	}

	if resp == nil {
		return nil, fmt.Errorf("no instance returned from create request")
	}

	return convertFluidStackInstanceToV1Instance(*resp), nil
}

func (c *FluidStackClient) GetInstance(ctx context.Context, instanceID v1.CloudProviderInstanceID) (*v1.Instance, error) {
	authCtx := c.makeAuthContext(ctx)
	projectCtx := c.makeProjectContext(authCtx)

	resp, httpResp, err := c.client.InstancesAPI.GetInstance(projectCtx, string(instanceID)).XPROJECTID(c.projectID).Execute()
	if httpResp != nil && httpResp.Body != nil {
		defer func() { _ = httpResp.Body.Close() }()
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get instance: %w", err)
	}

	if resp == nil {
		return nil, fmt.Errorf("no instance returned from get request")
	}

	return convertFluidStackInstanceToV1Instance(*resp), nil
}

func (c *FluidStackClient) TerminateInstance(ctx context.Context, instanceID v1.CloudProviderInstanceID) error {
	authCtx := c.makeAuthContext(ctx)
	projectCtx := c.makeProjectContext(authCtx)

	httpResp, err := c.client.InstancesAPI.DeleteInstance(projectCtx, string(instanceID)).XPROJECTID(c.projectID).Execute()
	if httpResp != nil && httpResp.Body != nil {
		defer func() { _ = httpResp.Body.Close() }()
	}
	if err != nil {
		return fmt.Errorf("failed to terminate instance: %w", err)
	}

	return nil
}

func (c *FluidStackClient) ListInstances(ctx context.Context, _ v1.ListInstancesArgs) ([]v1.Instance, error) {
	authCtx := c.makeAuthContext(ctx)
	projectCtx := c.makeProjectContext(authCtx)

	resp, httpResp, err := c.client.InstancesAPI.ListInstances(projectCtx).XPROJECTID(c.projectID).Execute()
	if httpResp != nil && httpResp.Body != nil {
		defer func() { _ = httpResp.Body.Close() }()
	}
	if err != nil {
		return nil, fmt.Errorf("failed to list instances: %w", err)
	}

	var instances []v1.Instance
	for _, fsInstance := range resp {
		instances = append(instances, *convertFluidStackInstanceToV1Instance(fsInstance))
	}

	return instances, nil
}

func (c *FluidStackClient) RebootInstance(_ context.Context, _ v1.CloudProviderInstanceID) error {
	return v1.ErrNotImplemented
}

func (c *FluidStackClient) MergeInstanceForUpdate(currInst v1.Instance, _ v1.Instance) v1.Instance {
	return currInst
}

func (c *FluidStackClient) MergeInstanceTypeForUpdate(currIt v1.InstanceType, _ v1.InstanceType) v1.InstanceType {
	return currIt
}

func convertFluidStackInstanceToV1Instance(fsInstance openapi.Instance) *v1.Instance {
	var privateIP string
	if fsInstance.Ip.IsSet() && fsInstance.Ip.Get() != nil {
		privateIP = *fsInstance.Ip.Get()
	}

	var lifecycleStatus v1.LifecycleStatus
	switch fsInstance.State {
	case openapi.INSTANCE_RUNNING:
		lifecycleStatus = v1.LifecycleStatusRunning
	case openapi.INSTANCE_STOPPED:
		lifecycleStatus = v1.LifecycleStatusStopped
	case openapi.INSTANCE_STOPPING:
		lifecycleStatus = v1.LifecycleStatusStopping
	case openapi.INSTANCE_STARTING, openapi.INSTANCE_CREATING:
		lifecycleStatus = v1.LifecycleStatusPending
	case openapi.INSTANCE_DELETING:
		lifecycleStatus = v1.LifecycleStatusTerminating
	case openapi.INSTANCE_ERROR:
		lifecycleStatus = v1.LifecycleStatusFailed
	default:
		lifecycleStatus = v1.LifecycleStatusPending
	}

	instance := &v1.Instance{
		Name:         fsInstance.Name,
		CloudID:      v1.CloudProviderInstanceID(fsInstance.Id),
		InstanceType: fsInstance.Type,
		PrivateIP:    privateIP,
		ImageID:      fsInstance.Image,
		Status: v1.Status{
			LifecycleStatus: lifecycleStatus,
		},
		Tags:        make(map[string]string),
		TunneledSSH: false,
	}

	for key, value := range fsInstance.Tags {
		instance.Tags[key] = value
	}

	return instance
}
