package v1

import (
	"context"

	v1 "github.com/brevdev/cloud/v1"
)

func (c *AWSClient) GetCapabilities(_ context.Context) (v1.Capabilities, error) {
	capabilities := v1.Capabilities{
		v1.CapabilityVPC,
	}

	return capabilities, nil
}
