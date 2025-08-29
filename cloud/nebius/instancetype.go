package nebius

import (
	"context"
	"time"

	"github.com/brevdev/sdk/cloud"
)

func (c *NebiusClient) GetInstanceTypes(_ context.Context, _ cloud.GetInstanceTypeArgs) ([]cloud.InstanceType, error) {
	return nil, cloud.ErrNotImplemented
}

func (c *NebiusClient) GetInstanceTypePollTime() time.Duration {
	return 5 * time.Minute
}

func (c *NebiusClient) MergeInstanceTypeForUpdate(currIt cloud.InstanceType, newIt cloud.InstanceType) cloud.InstanceType {
	merged := newIt

	merged.ID = currIt.ID

	return merged
}
