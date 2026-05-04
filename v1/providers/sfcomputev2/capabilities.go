package v2

import (
	"context"

	v1 "github.com/brevdev/cloud/v1"
)

func getSFCCapabilitiesV2() v1.Capabilities {
	return v1.Capabilities{
		v1.CapabilityCreateInstance,
		v1.CapabilityTerminateInstance,
		v1.CapabilityCreateTerminateInstance,
	}
}

func (c *SFCClientV2) GetCapabilities(_ context.Context) (v1.Capabilities, error) {
	return getSFCCapabilitiesV2(), nil
}

func (c *SFCCredentialV2) GetCapabilities(_ context.Context) (v1.Capabilities, error) {
	return getSFCCapabilitiesV2(), nil
}
