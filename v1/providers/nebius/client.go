package v1

import (
	"context"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"

	"github.com/brevdev/cloud/internal/errors"
	v1 "github.com/brevdev/cloud/v1"
	"github.com/nebius/gosdk"
	"github.com/nebius/gosdk/auth"
	nebiusiamv1 "github.com/nebius/gosdk/proto/nebius/iam/v1"
)

const CloudProviderID string = "nebius"

type NebiusCredential struct {
	RefID               string
	PublicKeyID         string
	PrivateKeyPEMBase64 string
	ServiceAccountID    string
	ProjectID           string
}

var _ v1.CloudCredential = &NebiusCredential{}

func NewNebiusCredential(refID string, publicKeyID string, privateKeyPEMBase64 string, serviceAccountID string, projectID string) *NebiusCredential {
	return &NebiusCredential{
		RefID:               refID,
		PublicKeyID:         publicKeyID,
		PrivateKeyPEMBase64: privateKeyPEMBase64,
		ServiceAccountID:    serviceAccountID,
		ProjectID:           projectID,
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
	return v1.CloudProviderID(CloudProviderID)
}

// GetTenantID returns the tenant ID for Nebius (project ID)
func (c *NebiusCredential) GetTenantID() (string, error) {
	if c.ServiceAccountID == "" {
		return "", fmt.Errorf("service account ID is required for Nebius")
	}
	return c.ServiceAccountID, nil
}

func (c *NebiusCredential) MakeClient(ctx context.Context, _ string) (v1.CloudClient, error) {
	return NewNebiusClient(ctx, c.RefID, c.PublicKeyID, c.PrivateKeyPEMBase64, c.ServiceAccountID, c.ProjectID)
}

// It embeds NotImplCloudClient to handle unsupported features
type NebiusClient struct {
	v1.NotImplCloudClient
	refID     string
	projectID string
	sdk       *gosdk.SDK
	logger    v1.Logger
}

var _ v1.CloudClient = &NebiusClient{}

type NebiusClientOption func(c *NebiusClient)

func WithLogger(logger v1.Logger) NebiusClientOption {
	return func(c *NebiusClient) {
		c.logger = logger
	}
}

func NewNebiusClient(ctx context.Context, refID string, publicKeyID string, privateKeyPEMBase64 string, serviceAccountID string, projectID string, opts ...NebiusClientOption) (*NebiusClient, error) {
	// Decode base64 into raw PEM bytes
	pemBytes, err := base64.StdEncoding.DecodeString(privateKeyPEMBase64)
	if err != nil {
		return nil, fmt.Errorf("failed to base64 decode: %w", err)
	}

	// Decode the PEM block
	block, _ := pem.Decode(pemBytes)
	if block == nil {
		return nil, fmt.Errorf("failed to parse PEM block")
	}

	parsedKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse PKCS8 private key: %w", err)
	}
	var ok bool
	privateKey, ok := parsedKey.(*rsa.PrivateKey)
	if !ok {
		return nil, fmt.Errorf("not an RSA private key")
	}

	sdk, err := gosdk.New(ctx, gosdk.WithCredentials(
		gosdk.ServiceAccount(auth.ServiceAccount{
			PrivateKey:       privateKey,
			PublicKeyID:      publicKeyID,
			ServiceAccountID: serviceAccountID,
		}),
	))
	if err != nil {
		return nil, fmt.Errorf("failed to initialize Nebius SDK: %w", err)
	}

	nebiusClient := &NebiusClient{
		refID:     refID,
		projectID: projectID,
		sdk:       sdk,
		logger:    &v1.NoopLogger{},
	}

	for _, opt := range opts {
		opt(nebiusClient)
	}

	return nebiusClient, nil
}

// GetAPIType returns the API type for Nebius
func (c *NebiusClient) GetAPIType() v1.APIType {
	return v1.APITypeLocational
}

// GetCloudProviderID returns the cloud provider ID for Nebius
func (c *NebiusClient) GetCloudProviderID() v1.CloudProviderID {
	return "nebius"
}

// GetTenantID returns the tenant ID for Nebius
func (c *NebiusClient) GetTenantID() (string, error) {
	return c.projectID, nil
}

// GetReferenceID returns the reference ID for this client
func (c *NebiusClient) GetReferenceID() string {
	return c.refID
}

func (c *NebiusClient) GetLocation(ctx context.Context) (string, error) {
	nebiusProjectService := c.sdk.Services().IAM().V1().Project()

	// The target region is the same as the client's project region
	project, err := nebiusProjectService.Get(ctx, &nebiusiamv1.GetProjectRequest{
		Id: c.projectID,
	})
	if err != nil {
		return "", errors.WrapAndTrace(err)
	}
	return project.GetSpec().GetRegion(), nil
}
