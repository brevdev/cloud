package v1

import (
	"context"

	v1 "github.com/brevdev/cloud/v1"
)

func (c *NebiusClient) GetLocations(_ context.Context, _ v1.GetLocationsArgs) ([]v1.Location, error) {
	return nil, v1.ErrNotImplemented
}
