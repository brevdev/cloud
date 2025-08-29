package v1

import (
	"context"
	"time"

	v1 "github.com/brevdev/cloud/v1"
)

func (c *NebiusClient) GetInstanceTypes(_ context.Context, _ v1.GetInstanceTypeArgs) ([]v1.InstanceType, error) {
	return nil, v1.ErrNotImplemented
}

func (c *NebiusClient) GetInstanceTypePollTime() time.Duration {
	return 5 * time.Minute
}

func (c *NebiusClient) MergeInstanceTypeForUpdate(currIt v1.InstanceType, newIt v1.InstanceType) v1.InstanceType {
	merged := newIt

	merged.ID = currIt.ID

	return merged
}
