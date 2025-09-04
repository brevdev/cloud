package v1

import (
	"context"

	"github.com/brevdev/cloud/v1"
)

func (c *FluidStackClient) GetCapabilities(_ context.Context) (v1.Capabilities, error) {
	capabilities := v1.Capabilities{
		v1.CapabilityCreateInstance,
		v1.CapabilityTerminateInstance,
		v1.CapabilityStopStartInstance,
		v1.CapabilityTags,
		v1.CapabilityInstanceUserData,
	}

	return capabilities, nil
}
