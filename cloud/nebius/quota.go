package nebius

import (
	"context"

	"github.com/brevdev/sdk/cloud"
)

func (c *NebiusClient) GetInstanceTypeQuotas(_ context.Context, _ cloud.GetInstanceTypeQuotasArgs) (cloud.Quota, error) {
	return cloud.Quota{}, cloud.ErrNotImplemented
}
