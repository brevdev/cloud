package fluidstack

import (
	"context"
	"crypto/sha256"
	"fmt"
	"net/http"

	"github.com/brevdev/sdk/cloud"
	openapi "github.com/brevdev/sdk/cloud/fluidstack/gen/fluidstack"
)

const CloudProviderID = "fluidstack"

// FluidStackCredential implements the CloudCredential interface for FluidStack
type FluidStackCredential struct {
	RefID  string
	APIKey string
}

var _ cloud.CloudCredential = &FluidStackCredential{}

func NewFluidStackCredential(refID, apiKey string) *FluidStackCredential {
	return &FluidStackCredential{
		RefID:  refID,
		APIKey: apiKey,
	}
}

// GetReferenceID returns the reference ID for this credential
func (c *FluidStackCredential) GetReferenceID() string {
	return c.RefID
}

// GetAPIType returns the API type for FluidStack
func (c *FluidStackCredential) GetAPIType() cloud.APIType {
	return cloud.APITypeGlobal
}

// GetCloudProviderID returns the cloud provider ID for FluidStack
func (c *FluidStackCredential) GetCloudProviderID() cloud.CloudProviderID {
	return CloudProviderID
}

// GetTenantID returns the tenant ID for FluidStack
func (c *FluidStackCredential) GetTenantID() (string, error) {
	return fmt.Sprintf("%s-%x", CloudProviderID, sha256.Sum256([]byte(c.APIKey))), nil
}

// GetCapabilities returns the capabilities for FluidStack
func (c *FluidStackCredential) GetCapabilities(ctx context.Context) (cloud.Capabilities, error) {
	client, err := c.MakeClient(ctx, "")
	if err != nil {
		return nil, err
	}
	return client.GetCapabilities(ctx)
}

// MakeClient creates a new FluidStack client from this credential
func (c *FluidStackCredential) MakeClient(_ context.Context, _ string) (cloud.CloudClient, error) {
	return NewFluidStackClient(c.RefID, c.APIKey), nil
}

// FluidStackClient implements the CloudClient interface for FluidStack
// It embeds NotImplCloudClient to handle unsupported features
type FluidStackClient struct {
	cloud.NotImplCloudClient
	refID     string
	apiKey    string
	baseURL   string
	projectID string
	client    *openapi.APIClient
}

var _ cloud.CloudClient = &FluidStackClient{}

func NewFluidStackClient(refID, apiKey string) *FluidStackClient {
	config := openapi.NewConfiguration()
	config.HTTPClient = http.DefaultClient
	client := openapi.NewAPIClient(config)

	return &FluidStackClient{
		refID:   refID,
		apiKey:  apiKey,
		baseURL: "https://api.fluidstack.io/v1alpha1",
		client:  client,
	}
}

// GetAPIType returns the API type for FluidStack
func (c *FluidStackClient) GetAPIType() cloud.APIType {
	return cloud.APITypeGlobal
}

// GetCloudProviderID returns the cloud provider ID for FluidStack
func (c *FluidStackClient) GetCloudProviderID() cloud.CloudProviderID {
	return CloudProviderID
}

// MakeClient creates a new client instance
func (c *FluidStackClient) MakeClient(_ context.Context, _ string) (cloud.CloudClient, error) {
	return c, nil
}

// GetReferenceID returns the reference ID for this client
func (c *FluidStackClient) GetReferenceID() string {
	return c.refID
}

func (c *FluidStackClient) makeAuthContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, openapi.ContextAccessToken, c.apiKey)
}

func (c *FluidStackClient) makeProjectContext(ctx context.Context) context.Context {
	// FluidStack requires project ID to be passed, but we'll use a default for now
	if c.projectID == "" {
		c.projectID = "default-project-id"
	}
	return ctx
}
