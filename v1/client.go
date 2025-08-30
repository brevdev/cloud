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

type CloudAPI interface {
	GetAPIType() APIType
	GetCapabilities(ctx context.Context) (Capabilities, error)
	GetCloudProviderID() CloudProviderID
}

type CloudCredential interface {
	MakeClient(ctx context.Context, location string) (CloudClient, error)
	GetTenantID() (string, error)
	GetReferenceID() string
	CloudAPI
}

type CloudBase interface {
	CloudCreateTerminateInstance
}

type CloudClient interface {
	CloudCredential
	CloudBase
	CloudQuota
	CloudRebootInstance
	CloudStopStartInstance
	CloudResizeInstanceVolume
	CloudMachineImage
	CloudChangeInstanceType
	CloudModifyFirewall
	CloudInstanceTags
	UpdateHandler
}
