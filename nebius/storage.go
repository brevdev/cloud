package nebius

import (
	"context"

	"github.com/brevdev/cloud"
)

func (c *NebiusClient) ResizeInstanceVolume(_ context.Context, _ cloud.ResizeInstanceVolumeArgs) error {
	return cloud.ErrNotImplemented
}
