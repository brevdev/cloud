package v1

import (
	"context"

	v1 "github.com/brevdev/cloud/v1"
)

func (c *NebiusClient) GetImages(_ context.Context, _ v1.GetImageArgs) ([]v1.Image, error) {
	return nil, v1.ErrNotImplemented
}
