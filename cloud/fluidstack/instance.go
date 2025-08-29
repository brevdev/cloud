package fluidstack

import (
	"context"
	"fmt"

	"github.com/brevdev/sdk/cloud"
	openapi "github.com/brevdev/sdk/cloud/fluidstack/gen/fluidstack"
)

func (c *FluidStackClient) CreateInstance(ctx context.Context, attrs cloud.CreateInstanceAttrs) (*cloud.Instance, error) {
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

func (c *FluidStackClient) GetInstance(ctx context.Context, instanceID cloud.CloudProviderInstanceID) (*cloud.Instance, error) {
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

func (c *FluidStackClient) TerminateInstance(ctx context.Context, instanceID cloud.CloudProviderInstanceID) error {
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

func (c *FluidStackClient) ListInstances(ctx context.Context, _ cloud.ListInstancesArgs) ([]cloud.Instance, error) {
	authCtx := c.makeAuthContext(ctx)
	projectCtx := c.makeProjectContext(authCtx)

	resp, httpResp, err := c.client.InstancesAPI.ListInstances(projectCtx).XPROJECTID(c.projectID).Execute()
	if httpResp != nil && httpResp.Body != nil {
		defer func() { _ = httpResp.Body.Close() }()
	}
	if err != nil {
		return nil, fmt.Errorf("failed to list instances: %w", err)
	}

	var instances []cloud.Instance
	for _, fsInstance := range resp {
		instances = append(instances, *convertFluidStackInstanceToV1Instance(fsInstance))
	}

	return instances, nil
}

func (c *FluidStackClient) RebootInstance(_ context.Context, _ cloud.CloudProviderInstanceID) error {
	return cloud.ErrNotImplemented
}

func (c *FluidStackClient) MergeInstanceForUpdate(currInst cloud.Instance, _ cloud.Instance) cloud.Instance {
	return currInst
}

func (c *FluidStackClient) MergeInstanceTypeForUpdate(currIt cloud.InstanceType, _ cloud.InstanceType) cloud.InstanceType {
	return currIt
}

func convertFluidStackInstanceToV1Instance(fsInstance openapi.Instance) *cloud.Instance {
	var privateIP string
	if fsInstance.Ip.IsSet() && fsInstance.Ip.Get() != nil {
		privateIP = *fsInstance.Ip.Get()
	}

	var lifecycleStatus cloud.LifecycleStatus
	switch fsInstance.State {
	case openapi.INSTANCE_RUNNING:
		lifecycleStatus = cloud.LifecycleStatusRunning
	case openapi.INSTANCE_STOPPED:
		lifecycleStatus = cloud.LifecycleStatusStopped
	case openapi.INSTANCE_STOPPING:
		lifecycleStatus = cloud.LifecycleStatusStopping
	case openapi.INSTANCE_STARTING, openapi.INSTANCE_CREATING:
		lifecycleStatus = cloud.LifecycleStatusPending
	case openapi.INSTANCE_DELETING:
		lifecycleStatus = cloud.LifecycleStatusTerminating
	case openapi.INSTANCE_ERROR:
		lifecycleStatus = cloud.LifecycleStatusFailed
	default:
		lifecycleStatus = cloud.LifecycleStatusPending
	}

	instance := &cloud.Instance{
		Name:         fsInstance.Name,
		CloudID:      cloud.CloudProviderInstanceID(fsInstance.Id),
		InstanceType: fsInstance.Type,
		PrivateIP:    privateIP,
		ImageID:      fsInstance.Image,
		Status: cloud.Status{
			LifecycleStatus: lifecycleStatus,
		},
		Tags: make(map[string]string),
	}

	for key, value := range fsInstance.Tags {
		instance.Tags[key] = value
	}

	return instance
}
