package v1

import (
	"context"

	v1 "github.com/brevdev/cloud/v1"
	"github.com/sfcompute/nodes-go/option"

	sfcnodes "github.com/sfcompute/nodes-go"
)

const CloudProviderID = "sfcompute"

type SFCCredential struct {
	RefID  string
	APIKey string `json:"api_key"`
}

var _ v1.CloudCredential = &SFCCredential{}

func NewSFCCredential(refID string, apiKey string) *SFCCredential {
	return &SFCCredential{
		RefID:  refID,
		APIKey: apiKey,
	}
}

func (c *SFCCredential) GetReferenceID() string {
	return c.RefID
}

func (c *SFCCredential) GetAPIType() v1.APIType {
	return v1.APITypeGlobal
}

func (c *SFCCredential) GetCloudProviderID() v1.CloudProviderID {
	return CloudProviderID
}

func (c *SFCCredential) GetTenantID() (string, error) {
	// sfc does not have a tenant system, return empty string
	return "", nil
}

type SFCClient struct {
	v1.NotImplCloudClient
	refID    string
	location string
	apiKey   string
	client   sfcnodes.Client
	logger   v1.Logger
}

var _ v1.CloudClient = &SFCClient{}

type SFCClientOption func(c *SFCClient)

func WithLogger(logger v1.Logger) SFCClientOption {
	return func(c *SFCClient) {
		c.logger = logger
	}
}

func (c *SFCCredential) MakeClientWithOptions(ctx context.Context, location string, opts ...SFCClientOption) (v1.CloudClient, error) {
	sfcClient := &SFCClient{
		refID:    c.RefID,
		apiKey:   c.APIKey,
		client:   sfcnodes.NewClient(option.WithBearerToken(c.APIKey)),
		location: location,
		logger:   &v1.NoopLogger{},
	}

	for _, opt := range opts {
		opt(sfcClient)
	}

	return sfcClient, nil
}

func (c *SFCCredential) MakeClient(ctx context.Context, location string) (v1.CloudClient, error) {
	return c.MakeClientWithOptions(ctx, location)
}

func (c *SFCClient) GetAPIType() v1.APIType {
	return v1.APITypeGlobal
}

func (c *SFCClient) GetCloudProviderID() v1.CloudProviderID {
	return CloudProviderID
}

func (c *SFCClient) GetReferenceID() string {
	return c.refID
}

func (c *SFCClient) GetTenantID() (string, error) {
	// sfc does not have a tenant system, return empty string
	return "", nil
}

func (c *SFCClient) MakeClient(_ context.Context, location string) (v1.CloudClient, error) {
	c.location = location
	return c, nil
}
