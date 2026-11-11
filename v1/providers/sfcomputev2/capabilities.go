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
		v1.CapabilityTags,
	}
}

func (c *SFCClientV2) GetCapabilities(ctx context.Context) (v1.Capabilities, error) {
	capabilities := getSFCCapabilitiesV2()
	if !c.enableConfigurableFirewall {
		return capabilities, nil
	}
	pool, err := c.client.getPool(ctx, c.GetDefaultPoolResourcePath())
	if err != nil {
		return nil, err
	}
	if pool.PublicIPv4SKUs != nil && len(*pool.PublicIPv4SKUs) > 0 {
		capabilities = append(capabilities, v1.CapabilityModifyFirewall)
	}
	return capabilities, nil
}

func (c *SFCCredentialV2) GetCapabilities(ctx context.Context) (v1.Capabilities, error) {
	if !c.EnableConfigurableFirewall {
		return getSFCCapabilitiesV2(), nil
	}
	client, err := c.MakeClient(ctx, "")
	if err != nil {
		return nil, err
	}
	return client.GetCapabilities(ctx)
}
