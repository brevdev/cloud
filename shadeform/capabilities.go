package shadeform

import (
	"context"

	"github.com/brevdev/cloud"
)

func (c *ShadeformClient) GetCapabilities(_ context.Context) (cloud.Capabilities, error) {
	capabilities := cloud.Capabilities{
		cloud.CapabilityCreateInstance,
		cloud.CapabilityTerminateInstance,
		cloud.CapabilityTags,
		cloud.CapabilityRebootInstance,
		cloud.CapabilityMachineImage,
	}

	return capabilities, nil
}
