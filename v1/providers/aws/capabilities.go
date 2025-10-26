package v1

import (
	"context"

	v1 "github.com/brevdev/cloud/v1"
)

func getAWSCapabilities() v1.Capabilities {
	return v1.Capabilities{
		v1.CapabilityVPC,
		v1.CapabilityManagedKubernetes,
	}
}

func (c *AWSClient) GetCapabilities(_ context.Context) (v1.Capabilities, error) {
	return getAWSCapabilities(), nil
}

func (c *AWSCredential) GetCapabilities(_ context.Context) (v1.Capabilities, error) {
	return getAWSCapabilities(), nil
}
