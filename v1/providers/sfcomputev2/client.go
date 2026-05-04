package v2

import (
	"context"

	v1 "github.com/brevdev/cloud/v1"
	sfc "github.com/sfcompute/sfc-go"
)

const CloudProviderID = "sfcompute"

// SFCCredentialV2 holds only authentication details. Operational config (capacity, image)
// is set on SFCClientV2 at MakeClient time via the Brev credential config.
type SFCCredentialV2 struct {
	RefID              string
	APIKey             string `json:"api_key"`
	CapacityID         string `json:"capacity_id"`
	ImageID            string `json:"image_id"`
}

var _ v1.CloudCredential = &SFCCredentialV2{}

func NewSFCCredentialV2(refID, apiKey, capacityID, imageID string) *SFCCredentialV2 {
	return &SFCCredentialV2{
		RefID:      refID,
		APIKey:     apiKey,
		CapacityID: capacityID,
		ImageID:    imageID,
	}
}

func (c *SFCCredentialV2) GetReferenceID() string {
	return c.RefID
}

func (c *SFCCredentialV2) GetAPIType() v1.APIType {
	return v1.APITypeGlobal
}

func (c *SFCCredentialV2) GetCloudProviderID() v1.CloudProviderID {
	return CloudProviderID
}

func (c *SFCCredentialV2) GetTenantID() (string, error) {
	return "", nil
}

type SFCClientV2 struct {
	v1.NotImplCloudClient
	refID      string
	location   string
	capacityID string
	imageID    string
	client     *sfc.SDK
	logger     v1.Logger
}

var _ v1.CloudClient = &SFCClientV2{}

type SFCClientV2Option func(c *SFCClientV2)

func WithLogger(logger v1.Logger) SFCClientV2Option {
	return func(c *SFCClientV2) {
		c.logger = logger
	}
}

func (c *SFCCredentialV2) MakeClientWithOptions(_ context.Context, location string, opts ...SFCClientV2Option) (v1.CloudClient, error) {
	sfcClient := &SFCClientV2{
		refID:      c.RefID,
		location:   location,
		capacityID: c.CapacityID,
		imageID:    c.ImageID,
		client:     sfc.New(sfc.WithSecurity(c.APIKey)),
		logger:     &v1.NoopLogger{},
	}

	for _, opt := range opts {
		opt(sfcClient)
	}

	return sfcClient, nil
}

func (c *SFCCredentialV2) MakeClient(ctx context.Context, location string) (v1.CloudClient, error) {
	return c.MakeClientWithOptions(ctx, location)
}

func (c *SFCClientV2) GetAPIType() v1.APIType {
	return v1.APITypeGlobal
}

func (c *SFCClientV2) GetCloudProviderID() v1.CloudProviderID {
	return CloudProviderID
}

func (c *SFCClientV2) GetReferenceID() string {
	return c.refID
}

func (c *SFCClientV2) GetTenantID() (string, error) {
	return "", nil
}

func (c *SFCClientV2) MakeClient(_ context.Context, location string) (v1.CloudClient, error) {
	c.location = location
	return c, nil
}
