package nebius

import (
	"context"

	"github.com/brevdev/cloud"
)

func (c *NebiusClient) GetImages(_ context.Context, _ cloud.GetImageArgs) ([]cloud.Image, error) {
	return nil, cloud.ErrNotImplemented
}
