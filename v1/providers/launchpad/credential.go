package v1

import (
	"context"
	"fmt"
	"time"

	v1 "github.com/brevdev/cloud/v1"
	validation "github.com/go-ozzo/ozzo-validation/v4"

	"github.com/pkg/errors"
)

const CloudProviderID = "launchpad"

type LaunchpadCredential struct {
	RefID    string
	APIToken string `json:"api_token"`
	APIURL   string `json:"api_url"`
}

var _ v1.CloudCredential = &LaunchpadCredential{}

func NewLaunchpadCredential(refID, apiToken, apiURL string) *LaunchpadCredential {
	return &LaunchpadCredential{
		RefID:    refID,
		APIToken: apiToken,
		APIURL:   apiURL,
	}
}

func (c *LaunchpadCredential) SetDefaults() {
	if c.APIURL == "" {
		c.APIURL = "https://launchpad.api.nvidia.com"
	}
}

func (c *LaunchpadCredential) Validate() error {
	err := validation.ValidateStruct(
		c,
		validation.Field(&c.APIToken, validation.Required),
		validation.Field(&c.APIURL, validation.Required),
	)
	if err != nil {
		return errors.Wrap(err, "failed to validate launchpad credential")
	}
	return nil
}

// GetAPIType implements v1.CloudCredential.
func (c *LaunchpadCredential) GetAPIType() v1.APIType {
	return v1.APITypeGlobal
}

// GetCloudProviderID implements v1.CloudCredential.
func (c *LaunchpadCredential) GetCloudProviderID() v1.CloudProviderID {
	return CloudProviderID
}

// GetReferenceID implements v1.CloudCredential.
func (c *LaunchpadCredential) GetReferenceID() string {
	return c.RefID
}

// GetTenantID implements v1.CloudCredential.
func (c *LaunchpadCredential) GetTenantID() (string, error) {
	hAPIKey, err := v1.HashSensitiveString(c.APIToken)
	if err != nil {
		return "", errors.Wrap(err, "failed to hash api key")
	}
	return fmt.Sprintf("launchpad-%s", hAPIKey), nil
}

func (c *LaunchpadCredential) MakeClientWithOptions(_ context.Context, _ string, opts ...LaunchpadClientOption) (v1.CloudClient, error) {
	return NewLaunchpadClient(*c, opts...)
}

// MakeClient implements v1.CloudCredential.
func (c *LaunchpadCredential) MakeClient(ctx context.Context, location string) (v1.CloudClient, error) {
	return c.MakeClientWithOptions(ctx, location)
}

func (c *LaunchpadCredential) GetInstancePollTime() time.Duration {
	return 1 * time.Minute
}

// GetCapabilities returns the capabilities for Launchpad
func (c *LaunchpadCredential) GetCapabilities(ctx context.Context) (v1.Capabilities, error) {
	client, err := c.MakeClient(ctx, "")
	if err != nil {
		return nil, err
	}
	return client.GetCapabilities(ctx)
}
