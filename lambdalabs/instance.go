package lambdalabs

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/alecthomas/units"
	"github.com/brevdev/cloud"
	"github.com/brevdev/cloud/internal/collections"
	openapi "github.com/brevdev/cloud/internal/gen/lambdalabs"
)

const lambdaLabsTimeNameFormat = "2006-01-02-15-04-05Z07-00"

// CreateInstance creates a new instance in Lambda Labs
// Supported via: POST /api/v1/instance-operations/launch
func (c *LambdaLabsClient) CreateInstance(ctx context.Context, attrs cloud.CreateInstanceAttrs) (*cloud.Instance, error) {
	keyPairName := attrs.RefID
	if attrs.KeyPairName != nil {
		keyPairName = *attrs.KeyPairName
	}

	if attrs.PublicKey != "" {
		request := openapi.AddSSHKeyRequest{
			Name:      keyPairName,
			PublicKey: &attrs.PublicKey,
		}
		keyPairResp, err := c.addSSHKey(ctx, request)
		if err != nil && !strings.Contains(err.Error(), "name must be unique") {
			return nil, fmt.Errorf("failed to add SSH key: %w", err)
		}
		keyPairName = keyPairResp.Data.Name
	}
	if keyPairName == "" {
		return nil, errors.New("keyPairName is required if public key not provided")
	}

	location := attrs.Location
	if location == "" {
		location = c.location
	}

	quantity := int32(1)
	request := openapi.LaunchInstanceRequest{
		RegionName:       location,
		InstanceTypeName: attrs.InstanceType,
		SshKeyNames:      []string{keyPairName},
		Quantity:         &quantity,
		FileSystemNames:  []string{},
	}

	name := fmt.Sprintf("%s--%s", c.GetReferenceID(), time.Now().UTC().Format(lambdaLabsTimeNameFormat))
	if len(name) > 64 {
		return nil, errors.New("name is too long")
	}

	request.Name = *openapi.NewNullableString(&name)

	resp, err := c.launchInstance(ctx, request)
	if err != nil {
		return nil, fmt.Errorf("failed to launch instance: %w", handleErrToCloudErr(err))
	}

	if len(resp.Data.InstanceIds) != 1 {
		return nil, fmt.Errorf("expected 1 instance ID, got %d", len(resp.Data.InstanceIds))
	}

	instanceID := cloud.CloudProviderInstanceID(resp.Data.InstanceIds[0])
	return c.GetInstance(ctx, instanceID)
}

// GetInstance retrieves an instance by ID
// Supported via: GET /api/v1/instances/{id}
func (c *LambdaLabsClient) GetInstance(ctx context.Context, instanceID cloud.CloudProviderInstanceID) (*cloud.Instance, error) {
	resp, err := c.getInstance(ctx, string(instanceID))
	if err != nil {
		return nil, fmt.Errorf("failed to get instance: %w", err)
	}

	return convertLambdaLabsInstanceToV1Instance(resp.Data), nil
}

// TerminateInstance terminates an instance
// Supported via: POST /api/v1/instance-operations/terminate
func (c *LambdaLabsClient) TerminateInstance(ctx context.Context, instanceID cloud.CloudProviderInstanceID) error {
	request := openapi.TerminateInstanceRequest{
		InstanceIds: []string{string(instanceID)},
	}

	_, err := c.terminateInstance(ctx, request)
	if err != nil {
		return fmt.Errorf("failed to terminate instance: %w", err)
	}

	return nil
}

// ListInstances lists all instances
// Supported via: GET /api/v1/instances
func (c *LambdaLabsClient) ListInstances(ctx context.Context, _ cloud.ListInstancesArgs) ([]cloud.Instance, error) {
	resp, err := c.listInstances(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to list instances: %w", err)
	}

	instances := make([]cloud.Instance, 0, len(resp.Data))
	for _, llInstance := range resp.Data {
		instance := convertLambdaLabsInstanceToV1Instance(llInstance)
		instances = append(instances, *instance)
	}

	return instances, nil
}

// RebootInstance reboots an instance
// Supported via: POST /api/v1/instance-operations/restart
func (c *LambdaLabsClient) RebootInstance(ctx context.Context, instanceID cloud.CloudProviderInstanceID) error {
	request := openapi.RestartInstanceRequest{
		InstanceIds: []string{string(instanceID)},
	}

	_, err := c.restartInstance(ctx, request)
	if err != nil {
		return fmt.Errorf("failed to reboot instance: %w", err)
	}

	return nil
}

// MergeInstanceForUpdate merges instance data for updates
func convertLambdaLabsInstanceToV1Instance(instance openapi.Instance) *cloud.Instance {
	var instanceIP string
	if instance.Ip.IsSet() {
		instanceIP = *instance.Ip.Get()
	}

	var instanceName string
	if instance.Name.IsSet() {
		instanceName = *instance.Name.Get()
	}

	var instanceHostname string
	if instance.Hostname.IsSet() {
		instanceHostname = *instance.Hostname.Get()
	}

	nameSplit := strings.Split(instanceName, "--")
	var cloudCredRefID string
	if len(nameSplit) > 0 {
		cloudCredRefID = nameSplit[0]
	}
	var createTime time.Time
	if len(nameSplit) > 1 {
		createTimeStr := nameSplit[1]
		createTime, _ = time.Parse(lambdaLabsTimeNameFormat, createTimeStr)
	}

	var instancePrivateIP string
	if instance.PrivateIp.IsSet() {
		instancePrivateIP = *instance.PrivateIp.Get()
	}

	inst := cloud.Instance{
		RefID:          instance.SshKeyNames[0],
		CloudCredRefID: cloudCredRefID,
		CreatedAt:      createTime,
		CloudID:        cloud.CloudProviderInstanceID(instance.Id),
		Name:           instanceName,
		PublicIP:       instanceIP,
		PublicDNS:      instanceIP,
		PrivateIP:      instancePrivateIP,
		Hostname:       instanceHostname,
		Status: cloud.Status{
			LifecycleStatus: convertLambdaLabsStatusToV1Status(instance.Status),
		},
		InstanceType: instance.InstanceType.Name,
		VolumeType:   "ssd",
		DiskSize:     units.GiB * units.Base2Bytes(instance.InstanceType.Specs.StorageGib),
		FirewallRules: cloud.FirewallRules{
			IngressRules: []cloud.FirewallRule{generateFirewallRouteFromPort(22), generateFirewallRouteFromPort(2222)}, // TODO pull from api
			EgressRules:  []cloud.FirewallRule{generateFirewallRouteFromPort(22), generateFirewallRouteFromPort(2222)}, // TODO pull from api
		},
		SSHUser:    "ubuntu",
		SSHPort:    22,
		Stoppable:  false,
		Rebootable: true,
		Location:   instance.Region.Name,
	}
	inst.InstanceTypeID = cloud.MakeGenericInstanceTypeIDFromInstance(inst)
	return &inst
}

func generateFirewallRouteFromPort(port int32) cloud.FirewallRule {
	return cloud.FirewallRule{
		FromPort: port,
		ToPort:   port,
		IPRanges: []string{"0.0.0.0/0"},
	}
}

func convertLambdaLabsStatusToV1Status(status string) cloud.LifecycleStatus {
	switch status {
	case "booting":
		return cloud.LifecycleStatusPending
	case "active":
		return cloud.LifecycleStatusRunning
	case "terminating":
		return cloud.LifecycleStatusTerminating
	case "terminated":
		return cloud.LifecycleStatusTerminated
	case "error":
		return cloud.LifecycleStatusFailed
	case "unhealthy":
		return cloud.LifecycleStatusRunning
	default:
		return cloud.LifecycleStatusPending
	}
}

func (c *LambdaLabsClient) MergeInstanceForUpdate(_ cloud.Instance, newInst cloud.Instance) cloud.Instance {
	return newInst
}

func (c *LambdaLabsClient) MergeInstanceTypeForUpdate(_ cloud.InstanceType, newIt cloud.InstanceType) cloud.InstanceType {
	return newIt
}

func (c *LambdaLabsClient) addSSHKey(ctx context.Context, request openapi.AddSSHKeyRequest) (*openapi.AddSSHKey200Response, error) {
	result, err := collections.RetryWithDataAndAttemptCount(func() (*openapi.AddSSHKey200Response, error) {
		res, resp, err := c.client.DefaultAPI.AddSSHKey(c.makeAuthContext(ctx)).AddSSHKeyRequest(request).Execute()
		if resp != nil {
			defer resp.Body.Close() //nolint:errcheck // ignore because using defer (for some reason HandleErrDefer)
		}
		if err != nil {
			return &openapi.AddSSHKey200Response{}, handleAPIError(ctx, resp, err)
		}
		return res, nil
	}, c.backoff)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *LambdaLabsClient) launchInstance(ctx context.Context, request openapi.LaunchInstanceRequest) (*openapi.LaunchInstance200Response, error) {
	result, err := collections.RetryWithDataAndAttemptCount(func() (*openapi.LaunchInstance200Response, error) {
		res, resp, err := c.client.DefaultAPI.LaunchInstance(c.makeAuthContext(ctx)).LaunchInstanceRequest(request).Execute()
		if resp != nil {
			defer resp.Body.Close() //nolint:errcheck // ignore because using defer (for some reason HandleErrDefer)
		}
		if err != nil {
			return &openapi.LaunchInstance200Response{}, handleAPIError(ctx, resp, err)
		}
		return res, nil
	}, c.backoff)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *LambdaLabsClient) getInstance(ctx context.Context, instanceID string) (*openapi.GetInstance200Response, error) {
	result, err := collections.RetryWithDataAndAttemptCount(func() (*openapi.GetInstance200Response, error) {
		res, resp, err := c.client.DefaultAPI.GetInstance(c.makeAuthContext(ctx), instanceID).Execute()
		if resp != nil {
			defer resp.Body.Close() //nolint:errcheck // ignore because using defer (for some reason HandleErrDefer)
		}
		if err != nil {
			return &openapi.GetInstance200Response{}, handleAPIError(ctx, resp, err)
		}
		return res, nil
	}, c.backoff)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *LambdaLabsClient) terminateInstance(ctx context.Context, request openapi.TerminateInstanceRequest) (*openapi.TerminateInstance200Response, error) {
	result, err := collections.RetryWithDataAndAttemptCount(func() (*openapi.TerminateInstance200Response, error) {
		res, resp, err := c.client.DefaultAPI.TerminateInstance(c.makeAuthContext(ctx)).TerminateInstanceRequest(request).Execute()
		if resp != nil {
			defer resp.Body.Close() //nolint:errcheck // ignore because using defer (for some reason HandleErrDefer)
		}
		if err != nil {
			return &openapi.TerminateInstance200Response{}, handleAPIError(ctx, resp, err)
		}
		return res, nil
	}, c.backoff)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *LambdaLabsClient) listInstances(ctx context.Context) (*openapi.ListInstances200Response, error) {
	result, err := collections.RetryWithDataAndAttemptCount(func() (*openapi.ListInstances200Response, error) {
		res, resp, err := c.client.DefaultAPI.ListInstances(c.makeAuthContext(ctx)).Execute()
		if resp != nil {
			defer resp.Body.Close() //nolint:errcheck // ignore because using defer (for some reason HandleErrDefer)
		}
		if err != nil {
			return &openapi.ListInstances200Response{}, handleAPIError(ctx, resp, err)
		}
		return res, nil
	}, c.backoff)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *LambdaLabsClient) restartInstance(ctx context.Context, request openapi.RestartInstanceRequest) (*openapi.RestartInstance200Response, error) {
	result, err := collections.RetryWithDataAndAttemptCount(func() (*openapi.RestartInstance200Response, error) {
		res, resp, err := c.client.DefaultAPI.RestartInstance(c.makeAuthContext(ctx)).RestartInstanceRequest(request).Execute()
		if resp != nil {
			defer resp.Body.Close() //nolint:errcheck // ignore because using defer (for some reason HandleErrDefer)
		}
		if err != nil {
			return &openapi.RestartInstance200Response{}, handleAPIError(ctx, resp, err)
		}
		return res, nil
	}, c.backoff)
	if err != nil {
		return nil, err
	}
	return result, nil
}
