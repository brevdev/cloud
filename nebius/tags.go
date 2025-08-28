package nebius

import (
	"context"

	"github.com/brevdev/cloud"
)

func (c *NebiusClient) UpdateInstanceTags(_ context.Context, _ cloud.UpdateInstanceTagsArgs) error {
	return cloud.ErrNotImplemented
}
