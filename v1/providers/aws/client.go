package v1

import (
	"context"
	"crypto/sha256"
	"fmt"

	v1 "github.com/brevdev/cloud/v1"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	awslogging "github.com/aws/smithy-go/logging"
)

const CloudProviderID string = "aws"

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
	return v1.CloudProviderID(CloudProviderID)
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
	logger    v1.Logger
}

var _ v1.CloudClient = &AWSClient{}

type AWSClientOption func(c *AWSClient)

func WithLogger(logger v1.Logger) AWSClientOption {
	return func(c *AWSClient) {
		c.logger = logger
	}
}

func NewAWSClient(refID string, accessKeyID string, secretAccessKey string, region string, opts ...AWSClientOption) (*AWSClient, error) {
	ctx := context.Background()

	awsCredentials := credentials.NewStaticCredentialsProvider(accessKeyID, secretAccessKey, "")

	awsClient := &AWSClient{
		refID:  refID,
		region: region,
		logger: &v1.NoopLogger{},
	}

	for _, opt := range opts {
		opt(awsClient)
	}

	awsConfig, err := config.LoadDefaultConfig(ctx,
		config.WithCredentialsProvider(awsCredentials),
		config.WithRegion(region),
		config.WithLogger(&AWSLoggerAdapter{
			logger: awsClient.logger,
		}),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to load AWS config: %w", err)
	}

	awsClient.awsConfig = awsConfig

	return awsClient, nil
}

func (c *AWSClient) GetAPIType() v1.APIType {
	return v1.APITypeGlobal
}

func (c *AWSClient) GetCloudProviderID() v1.CloudProviderID {
	return v1.CloudProviderID(CloudProviderID)
}

func (c *AWSClient) GetReferenceID() string {
	return c.refID
}

type AWSLoggerAdapter struct {
	logger v1.Logger
}

func (l *AWSLoggerAdapter) Logf(classification awslogging.Classification, format string, v ...interface{}) {
	ctx := context.Background()
	switch classification {
	case awslogging.Debug:
		l.logger.Debug(ctx, fmt.Sprintf(format, v...))
	case awslogging.Warn:
		l.logger.Warn(ctx, fmt.Sprintf(format, v...))
	}
}
