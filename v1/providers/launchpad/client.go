package v1

import (
	"context"
	"fmt"
	"net/http"

	v1 "github.com/brevdev/cloud/v1"
	openapi "github.com/brevdev/cloud/v1/providers/launchpad/gen/launchpad"
	"github.com/pkg/errors"
)

type LaunchpadClient struct {
	v1.NotImplCloudClient
	refID  string
	apiKey string
	client *openapi.APIClient
	logger v1.Logger
}

var _ v1.CloudClient = &LaunchpadClient{}

type LaunchpadClientOption func(c *LaunchpadClient)

func WithLogger(logger v1.Logger) LaunchpadClientOption {
	return func(c *LaunchpadClient) {
		c.logger = logger
	}
}

func NewLaunchpadClient(launchpadCredential LaunchpadCredential, opts ...LaunchpadClientOption) (*LaunchpadClient, error) {
	launchpadCredential.SetDefaults()
	err := launchpadCredential.Validate()
	if err != nil {
		return nil, errors.Wrap(err, "failed to validate launchpad credential")
	}

	conf := openapi.NewConfiguration()
	conf.HTTPClient = http.DefaultClient
	conf.Servers = []openapi.ServerConfiguration{
		{
			URL:         launchpadCredential.APIURL,
			Description: "Production server",
		},
	}
	client := openapi.NewAPIClient(conf)

	launchpadClient := &LaunchpadClient{
		refID:  launchpadCredential.RefID,
		apiKey: launchpadCredential.APIToken,
		client: client,
		logger: &v1.NoopLogger{},
	}

	for _, opt := range opts {
		opt(launchpadClient)
	}

	return launchpadClient, nil
}

func (c *LaunchpadClient) makeAuthContext(ctx context.Context) context.Context {
	authCtx := context.WithValue(ctx, openapi.ContextAPIKeys, map[string]openapi.APIKey{
		"TokenAuthentication": {
			Key: fmt.Sprintf("Token %s", c.apiKey),
		},
	})
	return authCtx
}

func (c *LaunchpadClient) GetAPIType() v1.APIType {
	return v1.APITypeGlobal
}

func (c *LaunchpadClient) GetCloudProviderID() v1.CloudProviderID {
	return CloudProviderID
}

func (c *LaunchpadClient) GetReferenceID() string {
	return c.refID
}

// GetCapabilities implements v1.CloudCredential.
func (c LaunchpadClient) GetCapabilities(_ context.Context) (v1.Capabilities, error) {
	return v1.Capabilities{
		v1.CapabilityCreateTerminateInstance,
		v1.CapabilityCreateInstance,
		v1.CapabilityTerminateInstance,
		v1.CapabilityCreateIdempotentInstance,
	}, nil
}
