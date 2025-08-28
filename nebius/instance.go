package nebius

import (
	"context"

	"github.com/brevdev/cloud"
)

func (c *NebiusClient) CreateInstance(_ context.Context, _ cloud.CreateInstanceAttrs) (*cloud.Instance, error) {
	return nil, cloud.ErrNotImplemented
}

func (c *NebiusClient) GetInstance(_ context.Context, _ cloud.CloudProviderInstanceID) (*cloud.Instance, error) {
	return nil, cloud.ErrNotImplemented
}

func (c *NebiusClient) TerminateInstance(_ context.Context, _ cloud.CloudProviderInstanceID) error {
	return cloud.ErrNotImplemented
}

func (c *NebiusClient) ListInstances(_ context.Context, _ cloud.ListInstancesArgs) ([]cloud.Instance, error) {
	return nil, cloud.ErrNotImplemented
}

func (c *NebiusClient) StopInstance(_ context.Context, _ cloud.CloudProviderInstanceID) error {
	return cloud.ErrNotImplemented
}

func (c *NebiusClient) StartInstance(_ context.Context, _ cloud.CloudProviderInstanceID) error {
	return cloud.ErrNotImplemented
}

func (c *NebiusClient) RebootInstance(_ context.Context, _ cloud.CloudProviderInstanceID) error {
	return cloud.ErrNotImplemented
}

func (c *NebiusClient) MergeInstanceForUpdate(currInst cloud.Instance, newInst cloud.Instance) cloud.Instance {
	merged := newInst

	merged.Name = currInst.Name
	merged.RefID = currInst.RefID
	merged.CloudCredRefID = currInst.CloudCredRefID
	merged.CreatedAt = currInst.CreatedAt
	merged.CloudID = currInst.CloudID
	merged.Location = currInst.Location
	merged.SubLocation = currInst.SubLocation

	return merged
}
