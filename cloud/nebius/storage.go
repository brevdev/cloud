package nebius

import (
	"context"

	"github.com/brevdev/sdk/cloud"
)

func (c *NebiusClient) ResizeInstanceVolume(_ context.Context, _ cloud.ResizeInstanceVolumeArgs) error {
	return cloud.ErrNotImplemented
}
