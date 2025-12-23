package v1

import (
	"context"
	"fmt"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/brevdev/cloud/internal/collections"
	v1 "github.com/brevdev/cloud/v1"
	openapi "github.com/brevdev/cloud/v1/providers/lambdalabs/gen/lambdalabs"
)

func TestLambdaLabsClient_CreateInstance_Success(t *testing.T) {
	client, cleanup := setupMockClient()
	defer cleanup()

	instanceID := testInstanceID
	publicKey := "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQ test@example.com"

	httpmock.RegisterResponder("POST", "https://cloud.lambda.ai/api/v1/ssh-keys",
		httpmock.NewJsonResponderOrPanic(200, openapi.AddSSHKey200Response{
			Data: openapi.SshKey{
				Id:        "ssh-key-id",
				Name:      "test-instance-id",
				PublicKey: publicKey,
			},
		}))

	httpmock.RegisterResponder("POST", "https://cloud.lambda.ai/api/v1/instance-operations/launch",
		httpmock.NewJsonResponderOrPanic(200, openapi.LaunchInstance200Response{
			Data: openapi.LaunchInstance200ResponseData{
				InstanceIds: []string{instanceID},
			},
		}))

	mockInstance := createMockInstance(instanceID)
	httpmock.RegisterResponder("GET", fmt.Sprintf("https://cloud.lambda.ai/api/v1/instances/%s", instanceID),
		httpmock.NewJsonResponderOrPanic(200, openapi.GetInstance200Response{
			Data: mockInstance,
		}))

	args := v1.CreateInstanceAttrs{
		InstanceType: "gpu_1x_a10",
		Location:     "us-west-1",
		PublicKey:    publicKey,
		Name:         "test-instance",
	}

	instance, err := client.CreateInstance(context.Background(), args)
	require.NoError(t, err)
	assert.Equal(t, instanceID, string(instance.CloudID))
	assert.Contains(t, instance.Name, "test-instance")
	assert.Equal(t, v1.LifecycleStatusRunning, instance.Status.LifecycleStatus)
	assert.False(t, instance.TunneledSSH)
}

func TestLambdaLabsClient_CreateInstance_WithoutPublicKey(t *testing.T) {
	client, cleanup := setupMockClient()
	defer cleanup()

	instanceID := testInstanceID

	httpmock.RegisterResponder("POST", "https://cloud.lambda.ai/api/v1/instance-operations/launch",
		httpmock.NewJsonResponderOrPanic(200, openapi.LaunchInstance200Response{
			Data: openapi.LaunchInstance200ResponseData{
				InstanceIds: []string{instanceID},
			},
		}))

	mockInstance := createMockInstance(instanceID)
	httpmock.RegisterResponder("GET", fmt.Sprintf("https://cloud.lambda.ai/api/v1/instances/%s", instanceID),
		httpmock.NewJsonResponderOrPanic(200, openapi.GetInstance200Response{
			Data: mockInstance,
		}))

	args := v1.CreateInstanceAttrs{
		InstanceType: "gpu_1x_a10",
		Location:     "us-west-1",
		Name:         "test-instance",
		KeyPairName:  collections.Ptr("test-key-pair"),
	}

	instance, err := client.CreateInstance(context.Background(), args)
	require.NoError(t, err)
	assert.Equal(t, instanceID, string(instance.CloudID))
}

func TestLambdaLabsClient_CreateInstance_SSHKeyError(t *testing.T) {
	client, cleanup := setupMockClient()
	defer cleanup()

	publicKey := "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQ test@example.com"

	httpmock.RegisterResponder("POST", "https://cloud.lambda.ai/api/v1/ssh-keys",
		httpmock.NewStringResponder(400, `{"error": {"code": "INVALID_REQUEST", "message": "SSH key already exists"}}`))

	args := v1.CreateInstanceAttrs{
		InstanceType: "gpu_1x_a10",
		Location:     "us-west-1",
		PublicKey:    publicKey,
		Name:         "test-instance",
	}

	_, err := client.CreateInstance(context.Background(), args)
	assert.Error(t, err)
}

func TestLambdaLabsClient_CreateInstance_LaunchError(t *testing.T) {
	client, cleanup := setupMockClient()
	defer cleanup()

	httpmock.RegisterResponder("POST", "https://cloud.lambda.ai/api/v1/instance-operations/launch",
		httpmock.NewStringResponder(400, `{"error": {"code": "INVALID_REQUEST", "message": "Instance type not available"}}`))

	args := v1.CreateInstanceAttrs{
		InstanceType: "gpu_1x_a10",
		Location:     "us-west-1",
		Name:         "test-instance",
	}

	_, err := client.CreateInstance(context.Background(), args)
	assert.Error(t, err)
}

func TestLambdaLabsClient_GetInstance_Success(t *testing.T) {
	client, cleanup := setupMockClient()
	defer cleanup()

	instanceID := testInstanceID
	mockInstance := createMockInstance(instanceID)

	httpmock.RegisterResponder("GET", fmt.Sprintf("https://cloud.lambda.ai/api/v1/instances/%s", instanceID),
		httpmock.NewJsonResponderOrPanic(200, openapi.GetInstance200Response{
			Data: mockInstance,
		}))

	instance, err := client.GetInstance(context.Background(), v1.CloudProviderInstanceID(instanceID))
	require.NoError(t, err)
	assert.Equal(t, instanceID, string(instance.CloudID))
	assert.Equal(t, "test-instance", instance.Name)
	assert.Equal(t, v1.LifecycleStatusRunning, instance.Status.LifecycleStatus)
	assert.False(t, instance.TunneledSSH)
}

func TestLambdaLabsClient_GetInstance_NotFound(t *testing.T) {
	client, cleanup := setupMockClient()
	defer cleanup()

	instanceID := nonexistentInstance

	httpmock.RegisterResponder("GET", fmt.Sprintf("https://cloud.lambda.ai/api/v1/instances/%s", instanceID),
		httpmock.NewStringResponder(404, `{"error": {"code": "NOT_FOUND", "message": "Instance not found"}}`))

	_, err := client.GetInstance(context.Background(), v1.CloudProviderInstanceID(instanceID))
	assert.Error(t, err)
}

func TestLambdaLabsClient_ListInstances_Success(t *testing.T) {
	client, cleanup := setupMockClient()
	defer cleanup()

	mockInstances := []openapi.Instance{
		createMockInstance("instance-1"),
		createMockInstance("instance-2"),
	}

	httpmock.RegisterResponder("GET", "https://cloud.lambda.ai/api/v1/instances",
		httpmock.NewJsonResponderOrPanic(200, openapi.ListInstances200Response{
			Data: mockInstances,
		}))

	instances, err := client.ListInstances(context.Background(), v1.ListInstancesArgs{})
	require.NoError(t, err)
	assert.Len(t, instances, 2)
	assert.Equal(t, "instance-1", string(instances[0].CloudID))
	assert.False(t, instances[0].TunneledSSH)
	assert.Equal(t, "instance-2", string(instances[1].CloudID))
	assert.False(t, instances[1].TunneledSSH)
}

func TestLambdaLabsClient_ListInstances_Empty(t *testing.T) {
	client, cleanup := setupMockClient()
	defer cleanup()

	httpmock.RegisterResponder("GET", "https://cloud.lambda.ai/api/v1/instances",
		httpmock.NewJsonResponderOrPanic(200, openapi.ListInstances200Response{
			Data: []openapi.Instance{},
		}))

	instances, err := client.ListInstances(context.Background(), v1.ListInstancesArgs{})
	require.NoError(t, err)
	assert.Len(t, instances, 0)
}

func TestLambdaLabsClient_ListInstances_Error(t *testing.T) {
	client, cleanup := setupMockClient()
	defer cleanup()

	httpmock.RegisterResponder("GET", "https://cloud.lambda.ai/api/v1/instances",
		httpmock.NewStringResponder(500, `{"error": {"code": "INTERNAL_ERROR", "message": "Internal server error"}}`))

	_, err := client.ListInstances(context.Background(), v1.ListInstancesArgs{})
	assert.Error(t, err)
}

func TestLambdaLabsClient_TerminateInstance_Success(t *testing.T) {
	client, cleanup := setupMockClient()
	defer cleanup()

	instanceID := testInstanceID

	httpmock.RegisterResponder("POST", "https://cloud.lambda.ai/api/v1/instance-operations/terminate",
		httpmock.NewJsonResponderOrPanic(200, openapi.TerminateInstance200Response{
			Data: openapi.TerminateInstance200ResponseData{
				TerminatedInstances: []openapi.Instance{
					createMockInstance(instanceID),
				},
			},
		}))

	err := client.TerminateInstance(context.Background(), v1.CloudProviderInstanceID(instanceID))
	assert.NoError(t, err)
}

func TestLambdaLabsClient_TerminateInstance_Error(t *testing.T) {
	client, cleanup := setupMockClient()
	defer cleanup()

	instanceID := nonexistentInstance

	httpmock.RegisterResponder("POST", "https://cloud.lambda.ai/api/v1/instance-operations/terminate",
		httpmock.NewStringResponder(404, `{"error": {"code": "NOT_FOUND", "message": "Instance not found"}}`))

	err := client.TerminateInstance(context.Background(), v1.CloudProviderInstanceID(instanceID))
	assert.Error(t, err)
}

func TestLambdaLabsClient_RebootInstance_Success(t *testing.T) {
	client, cleanup := setupMockClient()
	defer cleanup()

	instanceID := testInstanceID

	httpmock.RegisterResponder("POST", "https://cloud.lambda.ai/api/v1/instance-operations/restart",
		httpmock.NewJsonResponderOrPanic(200, openapi.RestartInstance200Response{
			Data: openapi.RestartInstance200ResponseData{
				RestartedInstances: []openapi.Instance{
					createMockInstance(instanceID),
				},
			},
		}))

	err := client.RebootInstance(context.Background(), v1.CloudProviderInstanceID(instanceID))
	assert.NoError(t, err)
}

func TestLambdaLabsClient_RebootInstance_Error(t *testing.T) {
	client, cleanup := setupMockClient()
	defer cleanup()

	instanceID := nonexistentInstance

	httpmock.RegisterResponder("POST", "https://cloud.lambda.ai/api/v1/instance-operations/restart",
		httpmock.NewStringResponder(404, `{"error": {"code": "NOT_FOUND", "message": "Instance not found"}}`))

	err := client.RebootInstance(context.Background(), v1.CloudProviderInstanceID(instanceID))
	assert.Error(t, err)
}
