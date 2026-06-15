package v1

import (
	"context"

	cloudv1 "github.com/brevdev/cloud/v1"
)

func getTestKubeCapabilities() cloudv1.Capabilities {
	return cloudv1.Capabilities{
		cloudv1.CapabilityCreateInstance,
		cloudv1.CapabilityTerminateInstance,
		cloudv1.CapabilityStopStartInstance,
		cloudv1.CapabilityRebootInstance,
		cloudv1.CapabilityTags,
	}
}

func (c *TestKubeCredential) GetCapabilities(_ context.Context) (cloudv1.Capabilities, error) {
	return getTestKubeCapabilities(), nil
}

func (c *TestKubeClient) GetCapabilities(_ context.Context) (cloudv1.Capabilities, error) {
	return getTestKubeCapabilities(), nil
}
