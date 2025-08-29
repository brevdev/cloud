package v1

import (
	"context"

	v1 "github.com/brevdev/cloud/v1"
)

func (c *NebiusClient) CreateInstance(_ context.Context, _ v1.CreateInstanceAttrs) (*v1.Instance, error) {
	return nil, v1.ErrNotImplemented
}

func (c *NebiusClient) GetInstance(_ context.Context, _ v1.CloudProviderInstanceID) (*v1.Instance, error) {
	return nil, v1.ErrNotImplemented
}

func (c *NebiusClient) TerminateInstance(_ context.Context, _ v1.CloudProviderInstanceID) error {
	return v1.ErrNotImplemented
}

func (c *NebiusClient) ListInstances(_ context.Context, _ v1.ListInstancesArgs) ([]v1.Instance, error) {
	return nil, v1.ErrNotImplemented
}

func (c *NebiusClient) StopInstance(_ context.Context, _ v1.CloudProviderInstanceID) error {
	return v1.ErrNotImplemented
}

func (c *NebiusClient) StartInstance(_ context.Context, _ v1.CloudProviderInstanceID) error {
	return v1.ErrNotImplemented
}

func (c *NebiusClient) RebootInstance(_ context.Context, _ v1.CloudProviderInstanceID) error {
	return v1.ErrNotImplemented
}

func (c *NebiusClient) MergeInstanceForUpdate(currInst v1.Instance, newInst v1.Instance) v1.Instance {
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
