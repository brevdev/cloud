package v1

import (
	"context"

	v1 "github.com/brevdev/cloud/v1"
)

func (c *NebiusClient) UpdateInstanceTags(_ context.Context, _ v1.UpdateInstanceTagsArgs) error {
	return v1.ErrNotImplemented
}
