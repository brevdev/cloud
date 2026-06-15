package v1

import (
	"context"

	cloudv1 "github.com/brevdev/cloud/v1"
)

func (c *TestKubeClient) GetLocations(_ context.Context, _ cloudv1.GetLocationsArgs) ([]cloudv1.Location, error) {
	return []cloudv1.Location{
		{
			Name:        c.location,
			Description: "Developer test Kubernetes cluster",
			Available:   true,
			Country:     "USA",
		},
	}, nil
}
