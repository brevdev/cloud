package v1

import (
	"context"
	"testing"

	"github.com/cenkalti/backoff"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	openapi "github.com/brevdev/cloud/internal/lambdalabs/gen/lambdalabs"
	v1 "github.com/brevdev/cloud/pkg/v1"
)

func TestLambdaLabsClient_GetAPIType(t *testing.T) {
	client := &LambdaLabsClient{}
	assert.Equal(t, v1.APITypeGlobal, client.GetAPIType())
}

func TestLambdaLabsClient_GetCloudProviderID(t *testing.T) {
	client := &LambdaLabsClient{}
	assert.Equal(t, v1.CloudProviderID(CloudProviderID), client.GetCloudProviderID())
}

func TestLambdaLabsClient_MakeClient(t *testing.T) {
	client := &LambdaLabsClient{
		refID:  "test-ref-id",
		apiKey: "test-api-key",
	}

	newClient, err := client.MakeClient(context.Background(), "test-tenant")
	require.NoError(t, err)
	lambdaClient, ok := newClient.(*LambdaLabsClient)
	require.True(t, ok)
	assert.Equal(t, client, lambdaClient)
}

func TestLambdaLabsClient_GetReferenceID(t *testing.T) {
	client := &LambdaLabsClient{refID: "test-ref-id"}
	assert.Equal(t, "test-ref-id", client.GetReferenceID())
}

func TestLambdaLabsClient_makeAuthContext(t *testing.T) {
	client := &LambdaLabsClient{apiKey: "test-api-key"}
	ctx := client.makeAuthContext(context.Background())

	auth := ctx.Value(openapi.ContextBasicAuth)
	require.NotNil(t, auth)

	basicAuth, ok := auth.(openapi.BasicAuth)
	require.True(t, ok)
	assert.Equal(t, "test-api-key", basicAuth.UserName)
	assert.Equal(t, "", basicAuth.Password)
}

func TestLambdaLabsClient_NewLambdaLabsClientRequiredFields(t *testing.T) {
	_, err := NewLambdaLabsClient("", "")
	require.Error(t, err)
	assert.Equal(t, "refID and apiKey are required", err.Error())
}

func TestLambdaLabsClient_NewLambdaLabsClientWithBaseURL(t *testing.T) {
	baseURL := "https://test.lambda.ai/api/v1"

	client, err := NewLambdaLabsClient("test-ref-id", "test-api-key", WithBaseURL(baseURL))
	require.NoError(t, err)
	assert.Equal(t, baseURL, client.baseURL)
}

func TestLambdaLabsClient_NewLambdaLabsClientWithClient(t *testing.T) {
	apiClient := openapi.NewAPIClient(openapi.NewConfiguration())

	client, err := NewLambdaLabsClient("test-ref-id", "test-api-key", WithClient(apiClient))
	require.NoError(t, err)
	assert.Equal(t, apiClient, client.client)
}

func TestLambdaLabsClient_NewLambdaLabsClientWithLocation(t *testing.T) {
	location := "us-west-1"

	client, err := NewLambdaLabsClient("test-ref-id", "test-api-key", WithLocation(location))
	require.NoError(t, err)
	assert.Equal(t, location, client.location)
}

func TestLambdaLabsClient_NewLambdaLabsClientWithBackoff(t *testing.T) {
	backoff := &backoff.ZeroBackOff{}

	client, err := NewLambdaLabsClient("test-ref-id", "test-api-key", WithBackoff(backoff))
	require.NoError(t, err)
	assert.Equal(t, backoff, client.backoff)
}
