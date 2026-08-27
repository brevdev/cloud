package massedcompute

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/pkg/errors"

	v1 "github.com/brevdev/cloud/v1"
	openapi "github.com/brevdev/cloud/v1/providers/massedcompute/gen/massedcompute"
)

const (
	CloudProviderID       = "massedcompute"
	DefaultAPIURL         = "https://vm.massedcompute.com/api/v1"
	massedComputeLocation = "massedcompute" // Massed Compute does not support location selection, so report all locations as "massedcompute"
	massedComputeRegion   = "any"           // The instance creation API expects 'any' as the region
)

type MassedComputeCredential struct {
	RefID    string
	APIToken string `json:"api_token"`
	APIURL   string `json:"api_url"`
}

var _ v1.CloudCredential = &MassedComputeCredential{}

func NewMassedComputeCredential(refID, apiToken string) *MassedComputeCredential {
	credential := &MassedComputeCredential{
		RefID:    refID,
		APIToken: apiToken,
	}
	credential.SetDefaults()
	return credential
}

func (c *MassedComputeCredential) SetDefaults() {
	if c.APIURL == "" {
		c.APIURL = DefaultAPIURL
	}
	c.APIURL = strings.TrimRight(c.APIURL, "/")
}

func (c *MassedComputeCredential) Validate() error {
	c.SetDefaults()
	if err := validation.ValidateStruct(
		c,
		validation.Field(&c.APIToken, validation.Required),
		validation.Field(&c.APIURL, validation.Required),
	); err != nil {
		return errors.Wrap(err, "failed to validate massed compute credential")
	}
	return nil
}

func (c *MassedComputeCredential) GetReferenceID() string {
	return c.RefID
}

func (c *MassedComputeCredential) GetAPIType() v1.APIType {
	return v1.APITypeGlobal
}

func (c *MassedComputeCredential) GetCloudProviderID() v1.CloudProviderID {
	return CloudProviderID
}

func (c *MassedComputeCredential) GetTenantID() (string, error) {
	return makeTenantID(c.APIToken)
}

func makeTenantID(apiToken string) (string, error) {
	hashedToken, err := v1.HashSensitiveString(apiToken)
	if err != nil {
		return "", errors.Wrap(err, "failed to hash massed compute API token")
	}
	return fmt.Sprintf("%s-%s", CloudProviderID, hashedToken), nil
}

func (c *MassedComputeCredential) MakeClient(ctx context.Context, location string) (v1.CloudClient, error) {
	return c.MakeClientWithOptions(ctx, location)
}

func (c *MassedComputeCredential) MakeClientWithOptions(_ context.Context, location string, opts ...MassedComputeClientOption) (v1.CloudClient, error) {
	return NewMassedComputeClient(*c, location, opts...)
}

func (c *MassedComputeCredential) GetCapabilities(_ context.Context) (v1.Capabilities, error) {
	return getCapabilities(), nil
}

type MassedComputeClient struct {
	v1.NotImplCloudClient

	refID      string
	apiToken   string
	client     *openapi.APIClient
	httpClient *http.Client
}

var _ v1.CloudClient = &MassedComputeClient{}

type MassedComputeClientOption func(*MassedComputeClient)

func WithHTTPClient(httpClient *http.Client) MassedComputeClientOption {
	return func(c *MassedComputeClient) {
		c.httpClient = httpClient
	}
}

func NewMassedComputeClient(credential MassedComputeCredential, _ string, opts ...MassedComputeClientOption) (*MassedComputeClient, error) {
	if err := credential.Validate(); err != nil {
		return nil, err
	}

	client := &MassedComputeClient{
		refID:      credential.RefID,
		apiToken:   credential.APIToken,
		httpClient: http.DefaultClient,
	}
	for _, opt := range opts {
		opt(client)
	}

	configuration := openapi.NewConfiguration()
	configuration.HTTPClient = client.httpClient
	configuration.UserAgent = "brev-cloud"
	configuration.Servers = []openapi.ServerConfiguration{{URL: credential.APIURL}}
	configuration.AddDefaultHeader("Authorization", "Bearer "+credential.APIToken)
	client.client = openapi.NewAPIClient(configuration)

	return client, nil
}

func (c *MassedComputeClient) GetReferenceID() string {
	return c.refID
}

func (c *MassedComputeClient) GetAPIType() v1.APIType {
	return v1.APITypeGlobal
}

func (c *MassedComputeClient) GetCloudProviderID() v1.CloudProviderID {
	return CloudProviderID
}

func (c *MassedComputeClient) GetTenantID() (string, error) {
	return makeTenantID(c.apiToken)
}

func (c *MassedComputeClient) MakeClient(_ context.Context, _ string) (v1.CloudClient, error) {
	return c, nil
}
