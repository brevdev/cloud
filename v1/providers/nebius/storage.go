package v1

import (
	"context"

	v1 "github.com/brevdev/cloud/v1"
)

func (c *NebiusClient) ResizeInstanceVolume(_ context.Context, _ v1.ResizeInstanceVolumeArgs) error {
	return v1.ErrNotImplemented
}
