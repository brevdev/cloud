package lambdalabs

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/brevdev/sdk/cloud"
)

func TestConvertLambdaLabsInstanceToV1Instance(t *testing.T) {
	lambdaInstance := createMockInstance("test-instance-id")

	v1Instance := convertLambdaLabsInstanceToV1Instance(lambdaInstance)

	assert.Equal(t, "test-instance-id", string(v1Instance.CloudID))
	assert.Equal(t, "test-instance", v1Instance.Name)
	assert.Equal(t, cloud.LifecycleStatusRunning, v1Instance.Status.LifecycleStatus)
	assert.Equal(t, "192.168.1.100", v1Instance.PublicIP)
	assert.Equal(t, "10.0.1.100", v1Instance.PrivateIP)
	assert.Equal(t, "us-west-1", v1Instance.Location)
	assert.Equal(t, "gpu_1x_a10", v1Instance.InstanceType)
}

func TestConvertLambdaLabsStatusToV1Status(t *testing.T) {
	tests := []struct {
		lambdaStatus string
		expected     cloud.LifecycleStatus
	}{
		{"active", cloud.LifecycleStatusRunning},
		{"booting", cloud.LifecycleStatusPending},
		{"unhealthy", cloud.LifecycleStatusRunning},
		{"terminating", cloud.LifecycleStatusTerminating},
		{"terminated", cloud.LifecycleStatusTerminated},
		{"error", cloud.LifecycleStatusFailed},
	}

	for _, test := range tests {
		t.Run(test.lambdaStatus, func(t *testing.T) {
			result := convertLambdaLabsStatusToV1Status(test.lambdaStatus)
			assert.Equal(t, test.expected, result)
		})
	}
}

func TestMergeInstanceForUpdate(t *testing.T) {
	client := &LambdaLabsClient{}
	original := cloud.Instance{
		CloudID: "test-id",
		Name:    "original-name",
		Status:  cloud.Status{LifecycleStatus: cloud.LifecycleStatusRunning},
	}

	update := cloud.Instance{
		Name:   "updated-name",
		Status: cloud.Status{LifecycleStatus: cloud.LifecycleStatusTerminated},
	}

	merged := client.MergeInstanceForUpdate(original, update)

	assert.Equal(t, "updated-name", merged.Name)
	assert.Equal(t, cloud.LifecycleStatusTerminated, merged.Status.LifecycleStatus)
}

func TestMergeInstanceTypeForUpdate(t *testing.T) {
	client := &LambdaLabsClient{}
	original := cloud.InstanceType{
		ID:   "test-id",
		Type: "original-type",
	}

	update := cloud.InstanceType{
		Type: "updated-type",
	}

	merged := client.MergeInstanceTypeForUpdate(original, update)

	assert.Equal(t, "updated-type", merged.Type)
}
