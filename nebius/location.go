package nebius

import (
	"context"

	"github.com/brevdev/cloud"
)

func (c *NebiusClient) GetLocations(_ context.Context, _ cloud.GetLocationsArgs) ([]cloud.Location, error) {
	return nil, cloud.ErrNotImplemented
}
