package v1

import (
	"context"
	"fmt"

	v1 "github.com/brevdev/cloud/v1"
)

func (c *SFCClient) CreateInstance(ctx context.Context, attrs v1.CreateInstanceAttrs) (*v1.Instance, error) {
	// 1) ensure SSH key present (or inject via API) per ../docs/SECURITY.md
	// 2) map attrs to provider request (location, instance type, image, tags, firewall rules if supported)
	// 3) launch and return instance converted to v1.Instance
	return nil, fmt.Errorf("not implemented")
}

func (c *SFCClient) GetInstance(ctx context.Context, id v1.CloudProviderInstanceID) (*v1.Instance, error) {
	return nil, fmt.Errorf("not implemented")
}

func (c *SFCClient) ListInstances(ctx context.Context, args v1.ListInstancesArgs) ([]v1.Instance, error) {
	return nil, fmt.Errorf("not implemented")
}

func (c *SFCClient) TerminateInstance(ctx context.Context, id v1.CloudProviderInstanceID) error {
	return fmt.Errorf("not implemented")
}

// Optional if supported:
func (c *SFCClient) RebootInstance(ctx context.Context, id v1.CloudProviderInstanceID) error {
	return fmt.Errorf("not implemented")
}
func (c *SFCClient) StopInstance(ctx context.Context, id v1.CloudProviderInstanceID) error {
	return fmt.Errorf("not implemented")
}
func (c *SFCClient) StartInstance(ctx context.Context, id v1.CloudProviderInstanceID) error {
	return fmt.Errorf("not implemented")
}

// Merge strategies (pass-through is acceptable baseline).
func (c *SFCClient) MergeInstanceForUpdate(_ v1.Instance, newInst v1.Instance) v1.Instance {
	return newInst
}
func (c *SFCClient) MergeInstanceTypeForUpdate(_ v1.InstanceType, newIt v1.InstanceType) v1.InstanceType {
	return newIt
}
