package massedcompute

import (
	"context"

	v1 "github.com/brevdev/cloud/v1"
)

func (c *MassedComputeClient) GetLocations(_ context.Context, _ v1.GetLocationsArgs) ([]v1.Location, error) {
	return []v1.Location{{
		Name:        massedComputeLocation,
		Description: "Massed Compute",
		Available:   true,
	}}, nil
}
