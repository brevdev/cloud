package v1

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/alecthomas/units"
	"github.com/brevdev/cloud/internal/collections"
	v1 "github.com/brevdev/cloud/v1"
	openapi "github.com/brevdev/cloud/v1/providers/lambdalabs/gen/lambdalabs"
)

const lambdaLabsTimeNameFormat = "2006-01-02-15-04-05Z07-00"

// CreateInstance creates a new instance in Lambda Labs
// Supported via: POST /api/v1/instance-operations/launch
func (c *LambdaLabsClient) CreateInstance(ctx context.Context, attrs v1.CreateInstanceAttrs) (*v1.Instance, error) {
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

	instanceID := v1.CloudProviderInstanceID(resp.Data.InstanceIds[0])
	return c.GetInstance(ctx, instanceID)
}

// GetInstance retrieves an instance by ID
// Supported via: GET /api/v1/instances/{id}
func (c *LambdaLabsClient) GetInstance(ctx context.Context, instanceID v1.CloudProviderInstanceID) (*v1.Instance, error) {
	resp, err := c.getInstance(ctx, string(instanceID))
	if err != nil {
		return nil, fmt.Errorf("failed to get instance: %w", err)
	}

	return convertLambdaLabsInstanceToV1Instance(resp.Data), nil
}

// TerminateInstance terminates an instance
// Supported via: POST /api/v1/instance-operations/terminate
func (c *LambdaLabsClient) TerminateInstance(ctx context.Context, instanceID v1.CloudProviderInstanceID) error {
	request := openapi.TerminateInstanceRequest{
		InstanceIds: []string{string(instanceID)},
	}

	// TODO: since "createInstance" also creates an SSH key, we should consider cleaning up the key pair here
	_, err := c.terminateInstance(ctx, request)
	if err != nil {
		return fmt.Errorf("failed to terminate instance: %w", err)
	}

	return nil
}

// ListInstances lists all instances
// Supported via: GET /api/v1/instances
func (c *LambdaLabsClient) ListInstances(ctx context.Context, _ v1.ListInstancesArgs) ([]v1.Instance, error) {
	resp, err := c.listInstances(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to list instances: %w", err)
	}

	instances := make([]v1.Instance, 0, len(resp.Data))
	for _, llInstance := range resp.Data {
		instance := convertLambdaLabsInstanceToV1Instance(llInstance)
		instances = append(instances, *instance)
	}

	return instances, nil
}

// RebootInstance reboots an instance
// Supported via: POST /api/v1/instance-operations/restart
func (c *LambdaLabsClient) RebootInstance(ctx context.Context, instanceID v1.CloudProviderInstanceID) error {
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
func convertLambdaLabsInstanceToV1Instance(instance openapi.Instance) *v1.Instance {
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

	inst := v1.Instance{
		RefID:          instance.SshKeyNames[0],
		CloudCredRefID: cloudCredRefID,
		CreatedAt:      createTime,
		CloudID:        v1.CloudProviderInstanceID(instance.Id),
		Name:           instanceName,
		PublicIP:       instanceIP,
		PublicDNS:      instanceIP,
		PrivateIP:      instancePrivateIP,
		Hostname:       instanceHostname,
		Status: v1.Status{
			LifecycleStatus: convertLambdaLabsStatusToV1Status(instance.Status),
		},
		InstanceType:      instance.InstanceType.Name,
		VolumeType:        "ssd",
		DiskSize:          units.GiB * units.Base2Bytes(instance.InstanceType.Specs.StorageGib),
		DiskSizeByteValue: v1.NewByteValue(instance.InstanceType.Specs.StorageGib, v1.Gibibyte),
		FirewallRules: v1.FirewallRules{
			IngressRules: []v1.FirewallRule{generateFirewallRouteFromPort(22), generateFirewallRouteFromPort(2222)}, // TODO pull from api
			EgressRules:  []v1.FirewallRule{generateFirewallRouteFromPort(22), generateFirewallRouteFromPort(2222)}, // TODO pull from api
		},
		SSHUser:    "ubuntu",
		SSHPort:    22,
		Stoppable:  false,
		Rebootable: true,
		Location:   instance.Region.Name,
	}
	inst.InstanceTypeID = v1.MakeGenericInstanceTypeIDFromInstance(inst)
	return &inst
}

func generateFirewallRouteFromPort(port int32) v1.FirewallRule {
	return v1.FirewallRule{
		FromPort: port,
		ToPort:   port,
		IPRanges: []string{"0.0.0.0/0"},
	}
}

func convertLambdaLabsStatusToV1Status(status string) v1.LifecycleStatus {
	switch status {
	case "booting":
		return v1.LifecycleStatusPending
	case "active":
		return v1.LifecycleStatusRunning
	case "terminating":
		return v1.LifecycleStatusTerminating
	case "terminated":
		return v1.LifecycleStatusTerminated
	case "error":
		return v1.LifecycleStatusFailed
	case "unhealthy":
		return v1.LifecycleStatusRunning
	default:
		return v1.LifecycleStatusPending
	}
}

func (c *LambdaLabsClient) MergeInstanceForUpdate(_ v1.Instance, newInst v1.Instance) v1.Instance {
	return newInst
}

func (c *LambdaLabsClient) MergeInstanceTypeForUpdate(_ v1.InstanceType, newIt v1.InstanceType) v1.InstanceType {
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
