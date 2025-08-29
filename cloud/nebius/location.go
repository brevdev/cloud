package nebius

import (
	"context"

	"github.com/brevdev/sdk/cloud"
)

func (c *NebiusClient) GetLocations(_ context.Context, _ cloud.GetLocationsArgs) ([]cloud.Location, error) {
	return nil, cloud.ErrNotImplemented
}
