package v1

import (
	"context"
	"fmt"

	v1 "github.com/brevdev/cloud/v1"
	"github.com/nebius/gosdk"
)

type NebiusCredential struct {
	RefID             string
	ServiceAccountKey string // JSON service account key
	ProjectID         string
}

var _ v1.CloudCredential = &NebiusCredential{}

func NewNebiusCredential(refID, serviceAccountKey, projectID string) *NebiusCredential {
	return &NebiusCredential{
		RefID:             refID,
		ServiceAccountKey: serviceAccountKey,
		ProjectID:         projectID,
	}
}

// GetReferenceID returns the reference ID for this credential
func (c *NebiusCredential) GetReferenceID() string {
	return c.RefID
}

// GetAPIType returns the API type for Nebius
func (c *NebiusCredential) GetAPIType() v1.APIType {
	return v1.APITypeLocational // Nebius uses location-specific endpoints
}

// GetCloudProviderID returns the cloud provider ID for Nebius
func (c *NebiusCredential) GetCloudProviderID() v1.CloudProviderID {
	return "nebius"
}

// GetTenantID returns the tenant ID for Nebius (project ID)
func (c *NebiusCredential) GetTenantID() (string, error) {
	if c.ProjectID == "" {
		return "", fmt.Errorf("project ID is required for Nebius")
	}
	return c.ProjectID, nil
}

func (c *NebiusCredential) MakeClient(ctx context.Context, location string) (v1.CloudClient, error) {
	return NewNebiusClient(ctx, c.RefID, c.ServiceAccountKey, c.ProjectID, location)
}

// It embeds NotImplCloudClient to handle unsupported features
type NebiusClient struct {
	v1.NotImplCloudClient
	refID             string
	serviceAccountKey string
	projectID         string
	location          string
	sdk               *gosdk.SDK
}

var _ v1.CloudClient = &NebiusClient{}

func NewNebiusClient(ctx context.Context, refID, serviceAccountKey, projectID, location string) (*NebiusClient, error) {
	sdk, err := gosdk.New(ctx, gosdk.WithCredentials(
		gosdk.IAMToken(serviceAccountKey), // For now, treat as IAM token - will need proper service account handling later
	))
	if err != nil {
		return nil, fmt.Errorf("failed to initialize Nebius SDK: %w", err)
	}

	return &NebiusClient{
		refID:             refID,
		serviceAccountKey: serviceAccountKey,
		projectID:         projectID,
		location:          location,
		sdk:               sdk,
	}, nil
}

// GetAPIType returns the API type for Nebius
func (c *NebiusClient) GetAPIType() v1.APIType {
	return v1.APITypeLocational
}

// GetCloudProviderID returns the cloud provider ID for Nebius
func (c *NebiusClient) GetCloudProviderID() v1.CloudProviderID {
	return "nebius"
}

// MakeClient creates a new client instance for a different location
func (c *NebiusClient) MakeClient(ctx context.Context, location string) (v1.CloudClient, error) {
	return NewNebiusClient(ctx, c.refID, c.serviceAccountKey, c.projectID, location)
}

// GetTenantID returns the tenant ID for Nebius
func (c *NebiusClient) GetTenantID() (string, error) {
	return c.projectID, nil
}

// GetReferenceID returns the reference ID for this client
func (c *NebiusClient) GetReferenceID() string {
	return c.refID
}
