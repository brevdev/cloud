package verda

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/pkg/errors"

	v1 "github.com/brevdev/cloud/v1"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	verdago "github.com/verda-cloud/verdacloud-sdk-go/pkg/verda"
)

const (
	CloudProviderID = "verda"
	DefaultAPIURL   = verdago.DefaultBaseURL
)

type VerdaCredential struct {
	RefID        string
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	APIURL       string `json:"api_url"`
}

var _ v1.CloudCredential = &VerdaCredential{}

func NewVerdaCredential(refID, clientID, clientSecret string) *VerdaCredential {
	credential := &VerdaCredential{
		RefID:        refID,
		ClientID:     clientID,
		ClientSecret: clientSecret,
	}
	credential.SetDefaults()
	return credential
}

func (c *VerdaCredential) SetDefaults() {
	if c.APIURL == "" {
		c.APIURL = DefaultAPIURL
	}
	c.APIURL = strings.TrimRight(c.APIURL, "/")
}

func (c *VerdaCredential) Validate() error {
	c.SetDefaults()
	if err := validation.ValidateStruct(
		c,
		validation.Field(&c.ClientID, validation.Required),
		validation.Field(&c.ClientSecret, validation.Required),
		validation.Field(&c.APIURL, validation.Required),
	); err != nil {
		return errors.Wrap(err, "failed to validate verda credential")
	}
	return nil
}

func (c *VerdaCredential) GetReferenceID() string {
	return c.RefID
}

func (c *VerdaCredential) GetAPIType() v1.APIType {
	return v1.APITypeGlobal
}

func (c *VerdaCredential) GetCloudProviderID() v1.CloudProviderID {
	return CloudProviderID
}

func (c *VerdaCredential) GetTenantID() (string, error) {
	return makeTenantID(c.ClientID)
}

func makeTenantID(clientID string) (string, error) {
	hashedClientID, err := v1.HashSensitiveString(clientID)
	if err != nil {
		return "", errors.Wrap(err, "failed to hash verda client ID")
	}
	return fmt.Sprintf("%s-%s", CloudProviderID, hashedClientID), nil
}

func (c *VerdaCredential) MakeClient(ctx context.Context, location string) (v1.CloudClient, error) {
	return c.MakeClientWithOptions(ctx, location)
}

func (c *VerdaCredential) MakeClientWithOptions(_ context.Context, location string, opts ...VerdaClientOption) (v1.CloudClient, error) {
	return NewVerdaClient(*c, location, opts...)
}

func (c *VerdaCredential) GetCapabilities(_ context.Context) (v1.Capabilities, error) {
	return getCapabilities(), nil
}

type VerdaClient struct {
	v1.NotImplCloudClient

	refID      string
	clientID   string
	location   string
	client     *verdago.Client
	httpClient *http.Client
}

var _ v1.CloudClient = &VerdaClient{}

type VerdaClientOption func(*VerdaClient)

func WithHTTPClient(httpClient *http.Client) VerdaClientOption {
	return func(c *VerdaClient) {
		c.httpClient = httpClient
	}
}

func NewVerdaClient(credential VerdaCredential, location string, opts ...VerdaClientOption) (*VerdaClient, error) {
	if err := credential.Validate(); err != nil {
		return nil, err
	}

	client := &VerdaClient{
		refID:      credential.RefID,
		clientID:   credential.ClientID,
		location:   location,
		httpClient: http.DefaultClient,
	}
	for _, opt := range opts {
		opt(client)
	}

	sdkClient, err := verdago.NewClient(
		verdago.WithBaseURL(credential.APIURL),
		verdago.WithClientID(credential.ClientID),
		verdago.WithClientSecret(credential.ClientSecret),
		verdago.WithHTTPClient(client.httpClient),
		verdago.WithUserAgent("brev-cloud"),
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create verda SDK client")
	}
	client.client = sdkClient
	return client, nil
}

func (c *VerdaClient) GetReferenceID() string {
	return c.refID
}

func (c *VerdaClient) GetAPIType() v1.APIType {
	return v1.APITypeGlobal
}

func (c *VerdaClient) GetCloudProviderID() v1.CloudProviderID {
	return CloudProviderID
}

func (c *VerdaClient) GetTenantID() (string, error) {
	return makeTenantID(c.clientID)
}

func (c *VerdaClient) MakeClient(_ context.Context, location string) (v1.CloudClient, error) {
	clientCopy := *c
	clientCopy.location = location
	return &clientCopy, nil
}
