package v1

import (
	"context"

	cloud "github.com/brevdev/cloud/v1"
)

func getNaverCapabilities() cloud.Capabilities {
	return cloud.Capabilities{
		cloud.CapabilityCreateInstance,
		cloud.CapabilityTerminateInstance,
		cloud.CapabilityCreateTerminateInstance,
		cloud.CapabilityRebootInstance,
		cloud.CapabilityStopStartInstance,
		cloud.CapabilityMachineImage,
	}
}

func (c *NaverClient) GetCapabilities(_ context.Context) (cloud.Capabilities, error) {
	return getNaverCapabilities(), nil
}
