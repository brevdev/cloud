package v1

import (
	"context"

	v1 "github.com/brevdev/cloud/v1"
)

func getSFCCapabilities() v1.Capabilities {
	return v1.Capabilities{
		v1.CapabilityCreateInstance,
		v1.CapabilityTerminateInstance,
		v1.CapabilityCreateTerminateInstance,
	}
}

func (c *SFCClient) GetCapabilities(_ context.Context) (v1.Capabilities, error) {
	return getSFCCapabilities(), nil
}

func (c *SFCCredential) GetCapabilities(_ context.Context) (v1.Capabilities, error) {
	return getSFCCapabilities(), nil
}
