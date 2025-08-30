package v1

import (
	"context"

	v1 "github.com/brevdev/cloud/v1"
)

func (c *NebiusClient) GetInstanceTypeQuotas(_ context.Context, _ v1.GetInstanceTypeQuotasArgs) (v1.Quota, error) {
	return v1.Quota{}, v1.ErrNotImplemented
}
