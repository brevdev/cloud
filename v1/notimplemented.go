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

func (c notImplCloudClient) GetInstancePollTime() time.Duration {
	pollTime := 5 * time.Second
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

func (c notImplCloudClient) CreateVPC(_ context.Context, _ CreateVPCArgs) (*VPC, error) {
	return nil, ErrNotImplemented
}

func (c notImplCloudClient) GetVPC(_ context.Context, _ GetVPCArgs) (*VPC, error) {
	return nil, ErrNotImplemented
}

func (c notImplCloudClient) DeleteVPC(_ context.Context, _ DeleteVPCArgs) error {
	return ErrNotImplemented
}

func (c notImplCloudClient) CreateCluster(_ context.Context, _ CreateClusterArgs) (*Cluster, error) {
	return nil, ErrNotImplemented
}

func (c notImplCloudClient) GetCluster(_ context.Context, _ GetClusterArgs) (*Cluster, error) {
	return nil, ErrNotImplemented
}

func (c notImplCloudClient) SetClusterUser(_ context.Context, _ SetClusterUserArgs) (*ClusterUser, error) {
	return nil, ErrNotImplemented
}

func (c notImplCloudClient) CreateNodeGroup(_ context.Context, _ CreateNodeGroupArgs) (*NodeGroup, error) {
	return nil, ErrNotImplemented
}

func (c notImplCloudClient) GetNodeGroup(_ context.Context, _ GetNodeGroupArgs) (*NodeGroup, error) {
	return nil, ErrNotImplemented
}

func (c notImplCloudClient) ModifyNodeGroup(_ context.Context, _ ModifyNodeGroupArgs) error {
	return ErrNotImplemented
}

func (c notImplCloudClient) DeleteNodeGroup(_ context.Context, _ DeleteNodeGroupArgs) error {
	return ErrNotImplemented
}

func (c notImplCloudClient) DeleteCluster(_ context.Context, _ DeleteClusterArgs) error {
	return ErrNotImplemented
}
