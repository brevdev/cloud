package v1

import (
	"context"
	"crypto/sha256"
	"fmt"

	v1 "github.com/brevdev/cloud/v1"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
)

const CloudProviderID = "aws"

type AWSCredential struct {
	RefID           string
	AccessKeyID     string
	SecretAccessKey string
}

var _ v1.CloudCredential = &AWSCredential{}

func NewAWSCredential(refID string, accessKeyID string, secretAccessKey string) *AWSCredential {
	return &AWSCredential{
		RefID:           refID,
		AccessKeyID:     accessKeyID,
		SecretAccessKey: secretAccessKey,
	}
}

func (c *AWSCredential) GetReferenceID() string {
	return c.RefID
}

func (c *AWSCredential) GetAPIType() v1.APIType {
	return v1.APITypeGlobal
}

func (c *AWSCredential) GetCloudProviderID() v1.CloudProviderID {
	return CloudProviderID
}

func (c *AWSCredential) GetTenantID() (string, error) {
	return fmt.Sprintf("%s-%x", CloudProviderID, sha256.Sum256([]byte(c.AccessKeyID))), nil
}

func (c *AWSCredential) MakeClient(_ context.Context, region string) (v1.CloudClient, error) {
	return NewAWSClient(c.RefID, c.AccessKeyID, c.SecretAccessKey, region)
}

type AWSClient struct {
	v1.NotImplCloudClient
	refID     string
	awsConfig aws.Config
	region    string
}

var _ v1.CloudClient = &AWSClient{}

func NewAWSClient(refID string, accessKeyID string, secretAccessKey string, region string) (*AWSClient, error) {
	ctx := context.Background()

	awsCredentials := credentials.NewStaticCredentialsProvider(accessKeyID, secretAccessKey, "")

	awsConfig, err := config.LoadDefaultConfig(ctx,
		config.WithCredentialsProvider(awsCredentials),
		config.WithRegion(region),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to load AWS config: %w", err)
	}

	return &AWSClient{
		refID:     refID,
		awsConfig: awsConfig,
		region:    region,
	}, nil
}

func (c *AWSClient) GetAPIType() v1.APIType {
	return v1.APITypeGlobal
}

func (c *AWSClient) GetCloudProviderID() v1.CloudProviderID {
	return CloudProviderID
}

func (c *AWSClient) GetReferenceID() string {
	return c.refID
}
