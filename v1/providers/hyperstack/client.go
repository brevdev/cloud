package hyperstack

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	pricebook "github.com/NexGenCloud/hyperstack-sdk-go/lib/Pricebook"
	"github.com/NexGenCloud/hyperstack-sdk-go/lib/environment"
	"github.com/NexGenCloud/hyperstack-sdk-go/lib/flavor"
	"github.com/NexGenCloud/hyperstack-sdk-go/lib/keypair"
	"github.com/NexGenCloud/hyperstack-sdk-go/lib/region"
	virtualmachine "github.com/NexGenCloud/hyperstack-sdk-go/lib/virtual_machine"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/pkg/errors"

	v1 "github.com/brevdev/cloud/v1"
)

const (
	CloudProviderID = "hyperstack"
	DefaultAPIURL   = "https://infrahub-api.nexgencloud.com/v1"
)

type HyperstackCredential struct {
	RefID  string
	APIKey string `json:"api_key"`
	APIURL string `json:"api_url"`
}

var _ v1.CloudCredential = &HyperstackCredential{}

func NewHyperstackCredential(refID, apiKey string) *HyperstackCredential {
	credential := &HyperstackCredential{
		RefID:  refID,
		APIKey: apiKey,
	}
	credential.SetDefaults()
	return credential
}

func (c *HyperstackCredential) SetDefaults() {
	if c.APIURL == "" {
		c.APIURL = DefaultAPIURL
	}
	c.APIURL = strings.TrimRight(c.APIURL, "/")
}

func (c *HyperstackCredential) Validate() error {
	c.SetDefaults()
	if err := validation.ValidateStruct(
		c,
		validation.Field(&c.APIKey, validation.Required),
		validation.Field(&c.APIURL, validation.Required),
	); err != nil {
		return errors.Wrap(err, "failed to validate hyperstack credential")
	}
	return nil
}

func (c *HyperstackCredential) GetReferenceID() string {
	return c.RefID
}

func (c *HyperstackCredential) GetAPIType() v1.APIType {
	return v1.APITypeGlobal
}

func (c *HyperstackCredential) GetCloudProviderID() v1.CloudProviderID {
	return CloudProviderID
}

func (c *HyperstackCredential) GetTenantID() (string, error) {
	return makeTenantID(c.APIKey)
}

func makeTenantID(apiKey string) (string, error) {
	hashedAPIKey, err := v1.HashSensitiveString(apiKey)
	if err != nil {
		return "", errors.Wrap(err, "failed to hash hyperstack API key")
	}
	return fmt.Sprintf("%s-%s", CloudProviderID, hashedAPIKey), nil
}

func (c *HyperstackCredential) MakeClient(ctx context.Context, location string) (v1.CloudClient, error) {
	return c.MakeClientWithOptions(ctx, location)
}

func (c *HyperstackCredential) MakeClientWithOptions(_ context.Context, location string, opts ...HyperstackClientOption) (v1.CloudClient, error) {
	return NewHyperstackClient(*c, location, opts...)
}

type HyperstackClient struct {
	v1.NotImplCloudClient

	refID      string
	apiKey     string
	location   string
	httpClient *http.Client
	logger     v1.Logger

	virtualMachines *virtualmachine.ClientWithResponses
	environments    *environment.ClientWithResponses
	flavors         *flavor.ClientWithResponses
	keypairs        *keypair.ClientWithResponses
	regions         *region.ClientWithResponses
	pricebook       *pricebook.Client
}

var _ v1.CloudClient = &HyperstackClient{}

type HyperstackClientOption func(*HyperstackClient)

func WithHTTPClient(httpClient *http.Client) HyperstackClientOption {
	return func(c *HyperstackClient) {
		c.httpClient = httpClient
	}
}

func WithLogger(logger v1.Logger) HyperstackClientOption {
	return func(c *HyperstackClient) {
		c.logger = logger
	}
}

func NewHyperstackClient(credential HyperstackCredential, location string, opts ...HyperstackClientOption) (*HyperstackClient, error) {
	if err := credential.Validate(); err != nil {
		return nil, err
	}

	client := &HyperstackClient{
		refID:      credential.RefID,
		apiKey:     credential.APIKey,
		location:   location,
		httpClient: http.DefaultClient,
		logger:     &v1.NoopLogger{},
	}
	for _, opt := range opts {
		opt(client)
	}
	if client.httpClient == nil {
		return nil, errors.New("hyperstack HTTP client is required")
	}

	doer := &authenticatedDoer{apiKey: credential.APIKey, client: client.httpClient}
	var err error
	client.virtualMachines, err = virtualmachine.NewClientWithResponses(
		credential.APIURL,
		virtualmachine.WithHTTPClient(doer),
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create hyperstack virtual-machine client")
	}
	client.environments, err = environment.NewClientWithResponses(credential.APIURL, environment.WithHTTPClient(doer))
	if err != nil {
		return nil, errors.Wrap(err, "failed to create hyperstack environment client")
	}
	client.flavors, err = flavor.NewClientWithResponses(credential.APIURL, flavor.WithHTTPClient(doer))
	if err != nil {
		return nil, errors.Wrap(err, "failed to create hyperstack flavor client")
	}
	client.keypairs, err = keypair.NewClientWithResponses(credential.APIURL, keypair.WithHTTPClient(doer))
	if err != nil {
		return nil, errors.Wrap(err, "failed to create hyperstack keypair client")
	}
	client.regions, err = region.NewClientWithResponses(credential.APIURL, region.WithHTTPClient(doer))
	if err != nil {
		return nil, errors.Wrap(err, "failed to create hyperstack region client")
	}
	client.pricebook, err = pricebook.NewClient(credential.APIURL, pricebook.WithHTTPClient(doer))
	if err != nil {
		return nil, errors.Wrap(err, "failed to create hyperstack pricebook client")
	}

	return client, nil
}

type authenticatedDoer struct {
	apiKey string
	client *http.Client
}

func (d *authenticatedDoer) Do(request *http.Request) (*http.Response, error) {
	request = request.Clone(request.Context())
	request.Header.Set("api_key", d.apiKey)
	request.Header.Set("User-Agent", "brev-cloud")
	request.Header.Set("Accept", "application/json")
	return d.client.Do(request)
}

func (c *HyperstackClient) GetReferenceID() string {
	return c.refID
}

func (c *HyperstackClient) GetAPIType() v1.APIType {
	return v1.APITypeGlobal
}

func (c *HyperstackClient) GetCloudProviderID() v1.CloudProviderID {
	return CloudProviderID
}

func (c *HyperstackClient) GetTenantID() (string, error) {
	return makeTenantID(c.apiKey)
}

func (c *HyperstackClient) MakeClient(_ context.Context, location string) (v1.CloudClient, error) {
	clientCopy := *c
	clientCopy.location = location
	return &clientCopy, nil
}
