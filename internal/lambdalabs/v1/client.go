package v1

import (
	"context"
	"net/http"
	"time"

	openapi "github.com/brevdev/cloud/internal/lambdalabs/gen/lambdalabs"
	v1 "github.com/brevdev/cloud/pkg/v1"
	"github.com/cenkalti/backoff/v4"
)

// LambdaLabsClient implements the CloudClient interface for Lambda Labs
// It embeds NotImplCloudClient to handle unsupported features
type LambdaLabsClient struct {
	v1.NotImplCloudClient
	refID    string
	apiKey   string
	baseURL  string
	client   *openapi.APIClient
	location string
	backoff  backoff.BackOff
}

var _ v1.CloudClient = &LambdaLabsClient{}

type options struct {
	baseURL  string
	client   *openapi.APIClient
	location string
	backoff  backoff.BackOff
}

type Option func(options *options) error

func WithBaseURL(baseURL string) Option {
	return func(options *options) error {
		options.baseURL = baseURL
		return nil
	}
}

func WithClient(client *openapi.APIClient) Option {
	return func(options *options) error {
		options.client = client
		return nil
	}
}

func WithLocation(location string) Option {
	return func(options *options) error {
		options.location = location
		return nil
	}
}

func WithBackoff(backoff backoff.BackOff) Option {
	return func(options *options) error {
		options.backoff = backoff
		return nil
	}
}

// NewLambdaLabsClient creates a new Lambda Labs client
func NewLambdaLabsClient(refID, apiKey string, opts ...Option) (*LambdaLabsClient, error) {
	var options options
	for _, opt := range opts {
		if err := opt(&options); err != nil {
			return nil, err
		}
	}

	if options.baseURL == "" {
		options.baseURL = "https://cloud.lambda.ai/api/v1"
	}

	if options.client == nil {
		config := openapi.NewConfiguration()
		config.HTTPClient = http.DefaultClient
		options.client = openapi.NewAPIClient(config)
	}

	if options.backoff == nil {
		bo := backoff.NewExponentialBackOff()
		bo.InitialInterval = 1000 * time.Millisecond
		bo.MaxElapsedTime = 120 * time.Second
		options.backoff = bo
	}

	return &LambdaLabsClient{
		refID:    refID,
		apiKey:   apiKey,
		baseURL:  options.baseURL,
		client:   options.client,
		location: options.location,
		backoff:  options.backoff,
	}, nil
}

// GetAPIType returns the API type for Lambda Labs
func (c *LambdaLabsClient) GetAPIType() v1.APIType {
	return v1.APITypeGlobal // Lambda Labs uses a global API
}

// GetCloudProviderID returns the cloud provider ID for Lambda Labs
func (c *LambdaLabsClient) GetCloudProviderID() v1.CloudProviderID {
	return CloudProviderID
}

// MakeClient creates a new client instance
func (c *LambdaLabsClient) MakeClient(_ context.Context, location string) (v1.CloudClient, error) {
	if location == "" {
		location = DefaultRegion
	}
	c.location = location
	return c, nil
}

// GetTenantID returns the tenant ID for Lambda Labs
func (c *LambdaLabsClient) GetTenantID() (string, error) {
	// TODO: Implement tenant ID retrieval for Lambda Labs
	// This could be derived from the API key or account information
	return "", nil
}

// GetReferenceID returns the reference ID for this client
func (c *LambdaLabsClient) GetReferenceID() string {
	return c.refID
}

func (c *LambdaLabsClient) makeAuthContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, openapi.ContextBasicAuth, openapi.BasicAuth{
		UserName: c.apiKey,
	})
}
