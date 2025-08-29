package fluidstack

import (
	"context"

	"github.com/brevdev/sdk/cloud"
)

func (c *FluidStackClient) GetCapabilities(_ context.Context) (cloud.Capabilities, error) {
	capabilities := cloud.Capabilities{
		cloud.CapabilityCreateInstance,
		cloud.CapabilityTerminateInstance,
		cloud.CapabilityStopStartInstance,
		cloud.CapabilityTags,
		cloud.CapabilityInstanceUserData,
	}

	return capabilities, nil
}
