package v1

import (
	"context"
	"crypto/sha256"
	"fmt"
	"net/http"

	openapi "github.com/brevdev/cloud/internal/shadeform/gen/shadeform"
	"github.com/brevdev/cloud/pkg/v1"
)

const CloudProviderID = "shadeform"

// ShadeformCredential implements the CloudCredential interface for Shadeform
type ShadeformCredential struct {
	RefID  string
	APIKey string
}

var _ v1.CloudCredential = &ShadeformCredential{}

func NewShadeformCredential(refID, apiKey string) *ShadeformCredential {
	return &ShadeformCredential{
		RefID:  refID,
		APIKey: apiKey,
	}
}

// GetReferenceID returns the reference ID for this credential
func (c *ShadeformCredential) GetReferenceID() string {
	return c.RefID
}

// GetAPIType returns the API type for Shadeform
func (c *ShadeformCredential) GetAPIType() v1.APIType {
	return v1.APITypeGlobal
}

// GetCloudProviderID returns the cloud provider ID for Shadeform
func (c *ShadeformCredential) GetCloudProviderID() v1.CloudProviderID {
	return CloudProviderID
}

// GetTenantID returns the account ID for Shadeform
func (c *ShadeformCredential) GetTenantID() (string, error) {
	// TODO: Return the account ID for the API key
	return fmt.Sprintf("%s-%x", CloudProviderID, sha256.Sum256([]byte(c.APIKey))), nil
}

// GetCapabilities returns the capabilities for Shadeform
func (c *ShadeformCredential) GetCapabilities(ctx context.Context) (v1.Capabilities, error) {
	client, err := c.MakeClient(ctx, "")
	if err != nil {
		return nil, err
	}
	return client.GetCapabilities(ctx)
}

// MakeClient creates a new Shadeform client from this credential
func (c *ShadeformCredential) MakeClient(_ context.Context, _ string) (v1.CloudClient, error) {
	return NewShadeformClient(c.RefID, c.APIKey), nil
}

// Shadeform implements the CloudClient interface for Shadeform
// It embeds NotImplCloudClient to handle unsupported features
type ShadeformClient struct {
	v1.NotImplCloudClient
	refID     string
	apiKey    string
	baseURL   string
	projectID string
	client    *openapi.APIClient
	config    *Configuration
}

var _ v1.CloudClient = &ShadeformClient{}

func NewShadeformClient(refID, apiKey string) *ShadeformClient {
	config := openapi.NewConfiguration()
	config.HTTPClient = http.DefaultClient
	client := openapi.NewAPIClient(config)

	return &ShadeformClient{
		refID:   refID,
		apiKey:  apiKey,
		baseURL: "https://api.shadeform.ai/v1",
		client:  client,
	}
}

func (c *ShadeformClient) WithConfiguration(config Configuration) *ShadeformClient {
	c.config = &config
	return c
}

// GetAPIType returns the API type for Shadeform
func (c *ShadeformClient) GetAPIType() v1.APIType {
	return v1.APITypeGlobal
}

// GetCloudProviderID returns the cloud provider ID for Shadeform
func (c *ShadeformClient) GetCloudProviderID() v1.CloudProviderID {
	return CloudProviderID
}

// MakeClient creates a new client instance
func (c *ShadeformClient) MakeClient(_ context.Context, _ string) (v1.CloudClient, error) {
	return c, nil
}

// GetReferenceID returns the reference ID for this client
func (c *ShadeformClient) GetReferenceID() string {
	return c.refID
}

func (c *ShadeformClient) makeAuthContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, openapi.ContextAPIKeys, map[string]openapi.APIKey{
		"ApiKeyAuth": {
			Key:    c.apiKey,
			Prefix: "", // or "" if no prefix is needed
		},
	})
}
