package v1

import (
	openapi "github.com/brevdev/cloud/internal/lambdalabs/gen/lambdalabs"
	"github.com/cenkalti/backoff/v4"
	"github.com/jarcoal/httpmock"
)

const (
	testInstanceID      = "test-instance-id"
	nonexistentInstance = "nonexistent-instance"
)

func setupMockClient() (*LambdaLabsClient, func()) {
	httpmock.Activate()
	client, err := NewLambdaLabsClient("test-ref-id", "test-api-key", WithBackoff(&backoff.StopBackOff{}))
	if err != nil {
		panic(err)
	}
	return client, httpmock.DeactivateAndReset
}

func createMockInstance(instanceID string) openapi.Instance {
	name := "test-instance"
	ip := "192.168.1.100"
	privateIP := "10.0.1.100"
	hostname := "test-instance.lambda.ai"

	return openapi.Instance{
		Id:              instanceID,
		Name:            *openapi.NewNullableString(&name),
		Ip:              *openapi.NewNullableString(&ip),
		PrivateIp:       *openapi.NewNullableString(&privateIP),
		Status:          "active",
		SshKeyNames:     []string{"test-key"},
		FileSystemNames: []string{},
		Region: &openapi.Region{
			Name:        "us-west-1",
			Description: "US West 1",
		},
		InstanceType: &openapi.InstanceType{
			Name:              "gpu_1x_a10",
			Description:       "1x NVIDIA A10 GPU",
			GpuDescription:    "NVIDIA A10",
			PriceCentsPerHour: 100,
			Specs: openapi.InstanceTypeSpecs{
				MemoryGib:  32,
				StorageGib: 512,
			},
		},
		Hostname: *openapi.NewNullableString(&hostname),
	}
}
