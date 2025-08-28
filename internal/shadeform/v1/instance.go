package v1

import (
	"context"
	"errors"
	"fmt"
	"github.com/alecthomas/units"
	openapi "github.com/brevdev/cloud/internal/shadeform/gen/shadeform"
	"github.com/brevdev/cloud/pkg/v1"
	"github.com/google/uuid"
	"strings"
)

const (
	hostname              = "shadecloud"
	refIDTagName          = "refID"
	cloudCredRefIDTagName = "cloudCredRefID"
)

func (c *ShadeformClient) CreateInstance(ctx context.Context, attrs v1.CreateInstanceAttrs) (*v1.Instance, error) {
	authCtx := c.makeAuthContext(ctx)

	// Check if the instance type is allowed by configuration
	if !c.isInstanceTypeAllowed(attrs.InstanceType) {
		return nil, fmt.Errorf("instance type: %v is not deployable", attrs.InstanceType)
	}

	sshKeyID := ""

	keyPairName := attrs.RefID
	if attrs.KeyPairName != nil {
		keyPairName = *attrs.KeyPairName
	}

	if keyPairName == "" {
		keyPairName = uuid.New().String()
	}

	if attrs.PublicKey != "" {
		var err error
		sshKeyID, err = c.addSshKey(ctx, keyPairName, attrs.PublicKey)
		if err != nil && !strings.Contains(err.Error(), "name must be unique") {
			return nil, fmt.Errorf("failed to add SSH key: %w", err)
		}
	}

	region := attrs.Location
	cloud, shadeInstanceType, err := c.getShadeformCloudAndInstanceType(attrs.InstanceType)
	if err != nil {
		return nil, err
	}

	cloudEnum, err := openapi.NewCloudFromValue(cloud)
	if err != nil {
		return nil, err
	}

	// Add refID tag
	refIDTag, err := c.createTag(refIDTagName, attrs.RefID)
	if err != nil {
		return nil, err
	}

	// Add cloudRefID tag
	cloudCredRefIDTag, err := c.createTag(cloudCredRefIDTagName, c.GetReferenceID())
	if err != nil {
		return nil, err
	}

	tags := []string{refIDTag, cloudCredRefIDTag}
	// Add all other tags
	for key, value := range attrs.Tags {
		createdTag, err := c.createTag(key, value)
		if err != nil {
			return nil, err
		}
		tags = append(tags, createdTag)
	}

	base64Script, err := c.GenerateFirewallScript(attrs.FirewallRules)
	if err != nil {
		return nil, err
	}

	req := openapi.CreateRequest{
		Cloud:             *cloudEnum,
		Region:            region,
		ShadeInstanceType: shadeInstanceType,
		Name:              attrs.Name,
		ShadeCloud:        true,
		Tags:              tags,
		SshKeyId:          &sshKeyID,
		LaunchConfiguration: &openapi.LaunchConfiguration{
			Type: "script",
			ScriptConfiguration: &openapi.ScriptConfiguration{
				Base64Script: base64Script,
			},
		},
	}

	resp, httpResp, err := c.client.DefaultAPI.InstancesCreate(authCtx).CreateRequest(req).Execute()
	if httpResp != nil && httpResp.Body != nil {
		defer func() { _ = httpResp.Body.Close() }()
	}
	if err != nil {
		return nil, fmt.Errorf("failed to create instance: %w", err)
	}

	if resp == nil {
		return nil, fmt.Errorf("no instance returned from create request")
	}

	// Since Shadeform doesn't return the full instance that's created, we need to make a second API call to get
	// the created instance after we call create
	createdInstance, err := c.GetInstance(authCtx, v1.CloudProviderInstanceID(resp.Id))
	if err != nil {
		return nil, err
	}

	return createdInstance, nil
}

func (c *ShadeformClient) addSshKey(ctx context.Context, keyPairName string, publicKey string) (string, error) {
	authCtx := c.makeAuthContext(ctx)

	request := openapi.AddSshKeyRequest{
		Name:      keyPairName,
		PublicKey: publicKey,
	}

	resp, httpResp, err := c.client.DefaultAPI.SshKeysAdd(authCtx).AddSshKeyRequest(request).Execute()
	if httpResp != nil && httpResp.Body != nil {
		defer func() { _ = httpResp.Body.Close() }()
	}
	if err != nil {
		return "", fmt.Errorf("failed to add SSH Key: %w", err)
	}

	if resp == nil {
		return "", fmt.Errorf("no instance returned from post request")
	}

	return resp.Id, nil
}

func (c *ShadeformClient) GetInstance(ctx context.Context, instanceID v1.CloudProviderInstanceID) (*v1.Instance, error) {
	authCtx := c.makeAuthContext(ctx)

	resp, httpResp, err := c.client.DefaultAPI.InstancesInfo(authCtx, string(instanceID)).Execute()
	if httpResp != nil && httpResp.Body != nil {
		defer func() { _ = httpResp.Body.Close() }()
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get instance: %w", err)
	}

	if resp == nil {
		return nil, fmt.Errorf("no instance returned from get request")
	}

	instance, err := c.convertInstanceInfoResponseToV1Instance(*resp)
	if err != nil {
		return nil, err
	}

	return instance, nil
}

func (c *ShadeformClient) TerminateInstance(ctx context.Context, instanceID v1.CloudProviderInstanceID) error {
	authCtx := c.makeAuthContext(ctx)

	httpResp, err := c.client.DefaultAPI.InstancesDelete(authCtx, string(instanceID)).Execute()
	if httpResp != nil && httpResp.Body != nil {
		defer func() { _ = httpResp.Body.Close() }()
	}
	if err != nil {
		return fmt.Errorf("failed to terminate instance: %w", err)
	}

	return nil
}

func (c *ShadeformClient) ListInstances(ctx context.Context, _ v1.ListInstancesArgs) ([]v1.Instance, error) {
	authCtx := c.makeAuthContext(ctx)

	resp, httpResp, err := c.client.DefaultAPI.Instances(authCtx).Execute()
	if httpResp != nil && httpResp.Body != nil {
		defer func() { _ = httpResp.Body.Close() }()
	}
	if err != nil {
		return nil, fmt.Errorf("failed to list instances: %w", err)
	}

	var instances []v1.Instance
	for _, instance := range resp.Instances {
		singleInstance, err := c.convertShadeformInstanceToV1Instance(instance)
		if err != nil {
			return nil, err
		}
		instances = append(instances, *singleInstance)
	}

	return instances, nil
}

func (c *ShadeformClient) RebootInstance(_ context.Context, _ v1.CloudProviderInstanceID) error {
	return v1.ErrNotImplemented
}

func (c *ShadeformClient) MergeInstanceForUpdate(currInst v1.Instance, _ v1.Instance) v1.Instance {
	return currInst
}

func (c *ShadeformClient) MergeInstanceTypeForUpdate(currIt v1.InstanceType, _ v1.InstanceType) v1.InstanceType {
	return currIt
}

func (c *ShadeformClient) getLifecycleStatus(status string) v1.LifecycleStatus {
	var lifecycleStatus v1.LifecycleStatus
	switch status {
	case "creating":
		lifecycleStatus = v1.LifecycleStatusPending
	case "pending_provider":
		lifecycleStatus = v1.LifecycleStatusPending
	case "pending":
		lifecycleStatus = v1.LifecycleStatusPending
	case "active":
		lifecycleStatus = v1.LifecycleStatusRunning
	case "error":
		lifecycleStatus = v1.LifecycleStatusFailed
	default:
		lifecycleStatus = v1.LifecycleStatusPending
	}
	return lifecycleStatus
}

// convertInstanceInfoResponseToV1Instance - converts Instance Info to v1 instance
func (c *ShadeformClient) convertInstanceInfoResponseToV1Instance(instanceInfo openapi.InstanceInfoResponse) (*v1.Instance, error) {
	instanceType := c.getInstanceType(string(instanceInfo.Cloud), instanceInfo.ShadeInstanceType)
	lifeCycleStatus := c.getLifecycleStatus(string(instanceInfo.Status))

	tags, err := c.convertShadeformTagToV1Tag(instanceInfo.Tags)
	if err != nil {
		return nil, err
	}

	refID, found := tags[refIDTagName]
	if !found {
		return nil, errors.New("could not find refID tag")
	}
	delete(tags, refIDTagName)

	cloudCredRefID := tags[cloudCredRefIDTagName]
	if err != nil {
		return nil, errors.New("could not find cloudCredRefID tag")
	}
	delete(tags, cloudCredRefIDTagName)

	instance := &v1.Instance{
		Name:         instanceInfo.Name,
		CreatedAt:    instanceInfo.CreatedAt,
		CloudID:      v1.CloudProviderInstanceID(instanceInfo.Id),
		PublicIP:     instanceInfo.Ip,
		Hostname:     hostname,
		ImageID:      instanceInfo.Configuration.Os,
		InstanceType: instanceType,
		DiskSize:     units.Base2Bytes(instanceInfo.Configuration.StorageInGb) * units.GiB,
		SSHUser:      instanceInfo.SshUser,
		SSHPort:      int(instanceInfo.SshPort),
		Status: v1.Status{
			LifecycleStatus: lifeCycleStatus,
		},
		Spot:           false,
		Location:       instanceInfo.Region,
		Stoppable:      false,
		Rebootable:     true,
		RefID:          refID,
		CloudCredRefID: cloudCredRefID,
	}

	return instance, nil
}

// convertInstanceInfoResponseToV1Instance - converts /instances response to v1 instance; the api struct is slightly
// different from instance info response and expected to diverge so keeping it as a separate function for now
func (c *ShadeformClient) convertShadeformInstanceToV1Instance(shadeformInstance openapi.Instance) (*v1.Instance, error) {
	instanceType := c.getInstanceType(string(shadeformInstance.Cloud), shadeformInstance.ShadeInstanceType)
	lifeCycleStatus := c.getLifecycleStatus(string(shadeformInstance.Status))

	tags, err := c.convertShadeformTagToV1Tag(shadeformInstance.Tags)
	if err != nil {
		return nil, err
	}

	refID, found := tags[refIDTagName]
	if !found {
		return nil, errors.New("could not find refID tag")
	}
	delete(tags, refIDTagName)

	cloudCredRefID := tags[cloudCredRefIDTagName]
	if err != nil {
		return nil, errors.New("could not find cloudCredRefID tag")
	}
	delete(tags, cloudCredRefIDTagName)

	instance := &v1.Instance{
		Name:         shadeformInstance.Name,
		CreatedAt:    shadeformInstance.CreatedAt,
		CloudID:      v1.CloudProviderInstanceID(shadeformInstance.Id),
		PublicIP:     shadeformInstance.Ip,
		Hostname:     hostname,
		ImageID:      shadeformInstance.Configuration.Os,
		InstanceType: instanceType,
		DiskSize:     units.Base2Bytes(shadeformInstance.Configuration.StorageInGb) * units.GiB,
		SSHUser:      shadeformInstance.SshUser,
		SSHPort:      int(shadeformInstance.SshPort),
		Status: v1.Status{
			LifecycleStatus: lifeCycleStatus,
		},
		Spot:           false,
		Location:       shadeformInstance.Region,
		Stoppable:      false,
		Rebootable:     true,
		RefID:          refID,
		Tags:           tags,
		CloudCredRefID: cloudCredRefID,
	}

	return instance, nil
}

func (c *ShadeformClient) convertShadeformTagToV1Tag(shadeformTags []string) (v1.Tags, error) {
	tags := v1.Tags{}
	for _, tag := range shadeformTags {
		key, value, err := c.getTag(tag)
		if err != nil {
			return nil, err
		}
		tags[key] = value
	}
	return tags, nil
}

func (c *ShadeformClient) createTag(key string, value string) (string, error) {
	if strings.Contains(key, "=") {
		return "", errors.New("tags cannot contain the '=' character")
	}

	return fmt.Sprintf("%v=%v", key, value), nil
}

func (c *ShadeformClient) getTag(shadeformTag string) (string, string, error) {
	key, value, found := strings.Cut(shadeformTag, "=")
	if !found {
		return "", "", errors.New(fmt.Sprintf("tag %v does not conform to the key value tag format", shadeformTag))
	}
	return key, value, nil
}
