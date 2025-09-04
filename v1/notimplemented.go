package v1

import (
	"context"
	"errors"
	"time"
)

var ErrNotImplemented = errors.New("not implemented")

type NotImplCloudClient struct {
	InnerNotImplCloudClient
}

var _ CloudClient = NotImplCloudClient{}

type InnerNotImplCloudClient struct {
	notImplCloudClient
}

type notImplCloudClient struct{}

func (c notImplCloudClient) GetReferenceID() string {
	return "not-implemented-reference-id"
}

func (c notImplCloudClient) GetCloudProviderID() CloudProviderID {
	return "not-implemented-cloud-provider-id"
}

var _ CloudClient = notImplCloudClient{}

func (c notImplCloudClient) GetAPIType() APIType {
	return APITypeGlobal
}

func (c notImplCloudClient) GetTenantID() (string, error) {
	return "", ErrNotImplemented
}

func (c notImplCloudClient) GetInstanceTypePollTime() time.Duration {
	pollTime := time.Minute
	return pollTime
}

func (c notImplCloudClient) CreateInstance(_ context.Context, _ CreateInstanceAttrs) (*Instance, error) {
	return nil, ErrNotImplemented
}

func (c notImplCloudClient) GetInstance(_ context.Context, _ CloudProviderInstanceID) (*Instance, error) {
	return nil, ErrNotImplemented
}

func (c notImplCloudClient) ChangeInstanceType(_ context.Context, _ CloudProviderInstanceID, _ string) error {
	return ErrNotImplemented
}

func (c notImplCloudClient) StopInstance(_ context.Context, _ CloudProviderInstanceID) error {
	return ErrNotImplemented
}

func (c notImplCloudClient) RebootInstance(_ context.Context, _ CloudProviderInstanceID) error {
	return ErrNotImplemented
}

func (c notImplCloudClient) StartInstance(_ context.Context, _ CloudProviderInstanceID) error {
	return ErrNotImplemented
}

func (c notImplCloudClient) TerminateInstance(_ context.Context, _ CloudProviderInstanceID) error {
	return ErrNotImplemented
}

func (c notImplCloudClient) GetInstanceTypes(_ context.Context, _ GetInstanceTypeArgs) ([]InstanceType, error) {
	return nil, ErrNotImplemented
}

func (c notImplCloudClient) GetImages(_ context.Context, _ GetImageArgs) ([]Image, error) {
	return nil, ErrNotImplemented
}

func (c notImplCloudClient) GetInstanceTypeQuotas(_ context.Context, _ GetInstanceTypeQuotasArgs) (Quota, error) {
	return Quota{}, ErrNotImplemented
}

func (c notImplCloudClient) AddFirewallRulesToInstance(_ context.Context, _ AddFirewallRulesToInstanceArgs) error {
	return ErrNotImplemented
}

func (c notImplCloudClient) RevokeSecurityGroupRules(_ context.Context, _ RevokeSecurityGroupRuleArgs) error {
	return ErrNotImplemented
}

func (c notImplCloudClient) ListInstances(_ context.Context, _ ListInstancesArgs) ([]Instance, error) {
	return nil, ErrNotImplemented
}

func (c notImplCloudClient) MakeClient(_ context.Context, _ string) (CloudClient, error) {
	return nil, ErrNotImplemented
}

func (c notImplCloudClient) GetLocations(_ context.Context, _ GetLocationsArgs) ([]Location, error) {
	return nil, ErrNotImplemented
}

func (c notImplCloudClient) ResizeInstanceVolume(_ context.Context, _ ResizeInstanceVolumeArgs) error {
	return ErrNotImplemented
}

func (c notImplCloudClient) GetCapabilities(_ context.Context) (Capabilities, error) {
	return []Capability{}, ErrNotImplemented
}

func (c notImplCloudClient) UpdateInstanceTags(_ context.Context, _ UpdateInstanceTagsArgs) error {
	return ErrNotImplemented
}

func (c notImplCloudClient) MergeInstanceForUpdate(_, i Instance) Instance {
	return i
}

func (c notImplCloudClient) MergeInstanceTypeForUpdate(_, i InstanceType) InstanceType {
	return i
}

func (c notImplCloudClient) GetMaxCreateRequestsPerMinute() int {
	return 10
}
