package v1

import (
	"context"

	v1 "github.com/brevdev/cloud/v1"
	"github.com/sfcompute/nodes-go/option"

	sfcnodes "github.com/sfcompute/nodes-go"
)

type SFCCredential struct {
	RefID  string
	apiKey string `json:"api_key"`
}

var _ v1.CloudCredential = &SFCCredential{}

func NewSFCCredential(refID string, apiKey string /* auth fields */) *SFCCredential {
	return &SFCCredential{
		RefID:  refID,
		apiKey: apiKey,
		// ...
	}
}

func (c *SFCCredential) GetReferenceID() string { return c.RefID }
func (c *SFCCredential) GetAPIType() v1.APIType { return v1.APITypeLocational /* or v1.APITypeGlobal */ }
func (c *SFCCredential) GetCloudProviderID() v1.CloudProviderID {
	return "sfcompute" // e.g., "lambdalabs"
}
func (c *SFCCredential) GetTenantID() (string, error) {
	// sfc does not have a tenant system, return empty string
	return "", nil
}

func (c *SFCCredential) MakeClient(ctx context.Context, location string) (v1.CloudClient, error) {
	// Create a client configured for a given location if locational API
	return NewSFCClient(c.RefID, c.apiKey /* auth fields */).MakeClient(ctx, location)
}

// ---------------- Client ----------------

type SFCClient struct {
	v1.NotImplCloudClient
	refID    string
	location string
	apiKey   string
	client   sfcnodes.Client // Add this field
	// add http/sdk client fields, base URLs, etc.
}

var _ v1.CloudClient = &SFCClient{}

func NewSFCClient(refID string, apiKey string /* auth fields */) *SFCClient {
	return &SFCClient{
		refID:  refID,
		apiKey: apiKey,
		client: sfcnodes.NewClient(
			option.WithBearerToken(apiKey)),
		// init http/sdk clients here
	}
}

func (c *SFCClient) GetAPIType() v1.APIType                 { return v1.APITypeGlobal /* or Global */ }
func (c *SFCClient) GetCloudProviderID() v1.CloudProviderID { return "sfcompute" }
func (c *SFCClient) GetReferenceID() string                 { return c.refID }
func (c *SFCClient) GetTenantID() (string, error)           { return "", nil }

func (c *SFCClient) MakeClient(_ context.Context, location string) (v1.CloudClient, error) {
	c.location = location
	return c, nil
}
