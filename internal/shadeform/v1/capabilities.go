package v1

import (
	"context"
	v1 "github.com/brevdev/cloud/pkg/v1"
)

func (c *ShadeformClient) GetCapabilities(_ context.Context) (v1.Capabilities, error) {
	capabilities := v1.Capabilities{
		v1.CapabilityCreateInstance,
		v1.CapabilityTerminateInstance,
		v1.CapabilityTags,
		v1.CapabilityRebootInstance,
		v1.CapabilityMachineImage,
	}

	return capabilities, nil
}
