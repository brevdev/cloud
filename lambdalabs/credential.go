package lambdalabs

import (
	"context"
	"crypto/sha256"
	"fmt"

	"github.com/brevdev/cloud"
)

const CloudProviderID = "lambda-labs"

const DefaultRegion string = "us-west-1"

// LambdaLabsCredential implements the CloudCredential interface for Lambda Labs
type LambdaLabsCredential struct {
	RefID  string
	APIKey string
}

var _ cloud.CloudCredential = &LambdaLabsCredential{}

// NewLambdaLabsCredential creates a new Lambda Labs credential
func NewLambdaLabsCredential(refID, apiKey string) *LambdaLabsCredential {
	return &LambdaLabsCredential{
		RefID:  refID,
		APIKey: apiKey,
	}
}

// GetReferenceID returns the reference ID for this credential
func (c *LambdaLabsCredential) GetReferenceID() string {
	return c.RefID
}

// GetAPIType returns the API type for Lambda Labs
func (c *LambdaLabsCredential) GetAPIType() cloud.APIType {
	return cloud.APITypeGlobal
}

// GetCloudProviderID returns the cloud provider ID for Lambda Labs
func (c *LambdaLabsCredential) GetCloudProviderID() cloud.CloudProviderID {
	return CloudProviderID
}

// GetTenantID returns the tenant ID for Lambda Labs
func (c *LambdaLabsCredential) GetTenantID() (string, error) {
	return fmt.Sprintf("lambdalabs-%x", sha256.Sum256([]byte(c.APIKey))), nil
}

// MakeClient creates a new Lambda Labs client from this credential
func (c *LambdaLabsCredential) MakeClient(_ context.Context, _ string) (cloud.CloudClient, error) {
	return NewLambdaLabsClient(c.RefID, c.APIKey)
}
