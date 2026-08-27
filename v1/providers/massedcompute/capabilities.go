package massedcompute

import (
	"context"

	v1 "github.com/brevdev/cloud/v1"
)

func getCapabilities() v1.Capabilities {
	return v1.Capabilities{
		v1.CapabilityCreateInstance,
		v1.CapabilityTerminateInstance,
		v1.CapabilityCreateTerminateInstance,
	}
}

func (c *MassedComputeClient) GetCapabilities(_ context.Context) (v1.Capabilities, error) {
	return getCapabilities(), nil
}
