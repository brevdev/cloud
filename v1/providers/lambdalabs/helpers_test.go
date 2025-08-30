package v1

import (
	"testing"

	"github.com/stretchr/testify/assert"

	v1 "github.com/brevdev/cloud/v1"
)

func TestConvertLambdaLabsInstanceToV1Instance(t *testing.T) {
	lambdaInstance := createMockInstance("test-instance-id")

	v1Instance := convertLambdaLabsInstanceToV1Instance(lambdaInstance)

	assert.Equal(t, "test-instance-id", string(v1Instance.CloudID))
	assert.Equal(t, "test-instance", v1Instance.Name)
	assert.Equal(t, v1.LifecycleStatusRunning, v1Instance.Status.LifecycleStatus)
	assert.Equal(t, "192.168.1.100", v1Instance.PublicIP)
	assert.Equal(t, "10.0.1.100", v1Instance.PrivateIP)
	assert.Equal(t, "us-west-1", v1Instance.Location)
	assert.Equal(t, "gpu_1x_a10", v1Instance.InstanceType)
}

func TestConvertLambdaLabsStatusToV1Status(t *testing.T) {
	tests := []struct {
		lambdaStatus string
		expected     v1.LifecycleStatus
	}{
		{"active", v1.LifecycleStatusRunning},
		{"booting", v1.LifecycleStatusPending},
		{"unhealthy", v1.LifecycleStatusRunning},
		{"terminating", v1.LifecycleStatusTerminating},
		{"terminated", v1.LifecycleStatusTerminated},
		{"error", v1.LifecycleStatusFailed},
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
	original := v1.Instance{
		CloudID: "test-id",
		Name:    "original-name",
		Status:  v1.Status{LifecycleStatus: v1.LifecycleStatusRunning},
	}

	update := v1.Instance{
		Name:   "updated-name",
		Status: v1.Status{LifecycleStatus: v1.LifecycleStatusTerminated},
	}

	merged := client.MergeInstanceForUpdate(original, update)

	assert.Equal(t, "updated-name", merged.Name)
	assert.Equal(t, v1.LifecycleStatusTerminated, merged.Status.LifecycleStatus)
}

func TestMergeInstanceTypeForUpdate(t *testing.T) {
	client := &LambdaLabsClient{}
	original := v1.InstanceType{
		ID:   "test-id",
		Type: "original-type",
	}

	update := v1.InstanceType{
		Type: "updated-type",
	}

	merged := client.MergeInstanceTypeForUpdate(original, update)

	assert.Equal(t, "updated-type", merged.Type)
}
