package v1

import (
	"context"
)

type APIType string

const (
	APITypeLocational APIType = "locational"
	APITypeGlobal     APIType = "global"
)

type CloudProviderID string // aws, gcp, azure, etc.

type CloudClient interface {
	CloudCredential
	ImageManager
	InstanceManager
	LocationManager
	NetworkManager
}

type CloudCredential interface {
	MakeClient(ctx context.Context, location string) (CloudClient, error)
	GetTenantID() (string, error)
	GetReferenceID() string
	CloudProvider
}

type CloudProvider interface {
	GetAPIType() APIType
	GetCapabilities(ctx context.Context) (Capabilities, error)
	GetCloudProviderID() CloudProviderID
}

type ImageManager interface {
	ImageGetter
}
type InstanceManager interface {
	InstanceCreator
	InstanceTerminator
	InstanceGetter
	InstanceLister
	InstanceRebooter
	InstanceStopStarter
	InstanceTypeChanger
	InstanceTagsUpdater
	InstanceUpdateHandler
	InstanceTypeGetter
	InstanceTypePollTimeGetter
	InstanceTypeQuotaGetter
	InstanceVolumeResizer
}

type LocationManager interface {
	LocationGetter
}

type NetworkManager interface {
	NetworkFirewallModifier
	NetworkSecurityGroupModifier
}
