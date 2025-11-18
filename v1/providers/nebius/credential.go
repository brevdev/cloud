package v1

import (
	"context"
	"fmt"

	"github.com/brevdev/cloud/internal/errors"
	v1 "github.com/brevdev/cloud/v1"
)

const CloudProviderID = "nebius"

const defaultNebiusLocation = "eu-north1"

// NebiusCredential implements the CloudCredential interface for Nebius AI Cloud
type NebiusCredential struct {
	RefID             string
	ServiceAccountKey string `json:"sa_json"`   // JSON service account key
	TenantID          string `json:"tenant_id"` // Nebius tenant ID (top-level organization)
}

var _ v1.CloudCredential = &NebiusCredential{}

// NewNebiusCredential creates a new Nebius credential
func NewNebiusCredential(refID, serviceAccountKey, tenantID string) *NebiusCredential {
	return &NebiusCredential{
		RefID:             refID,
		ServiceAccountKey: serviceAccountKey,
		TenantID:          tenantID,
	}
}

// NewNebiusCredentialWithOrg creates a new Nebius credential with organization ID
func NewNebiusCredentialWithOrg(refID, serviceAccountKey, tenantID, organizationID string) *NebiusCredential {
	return &NebiusCredential{
		RefID:             refID,
		ServiceAccountKey: serviceAccountKey,
		TenantID:          tenantID,
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
	return CloudProviderID
}

// GetTenantID returns the tenant ID
// Note: Project IDs are now determined per-region as default-project-{region}
func (c *NebiusCredential) GetTenantID() (string, error) {
	if c.TenantID == "" {
		return "", fmt.Errorf("tenant ID is required")
	}
	return c.TenantID, nil
}

// MakeClient creates a new Nebius client from this credential
func (c *NebiusCredential) MakeClient(ctx context.Context, location string) (v1.CloudClient, error) {
	return c.MakeClientWithOptions(ctx, location)
}

// MakeClientWithOptions creates a new Nebius client with options (e.g., logger)
func (c *NebiusCredential) MakeClientWithOptions(ctx context.Context, location string, opts ...NebiusClientOption) (v1.CloudClient, error) {
	// If no location is provided, use the default locaiton
	if location == "" {
		location = defaultNebiusLocation
	}

	// ProjectID is now determined in NewNebiusClient as default-project-{location}
	// Pass empty string and let the client constructor set it
	client, err := NewNebiusClientWithOrg(ctx, c.RefID, c.ServiceAccountKey, c.TenantID, "", "", location, opts...)
	if err != nil {
		return nil, errors.WrapAndTrace(err)
	}
	return client, nil
}
