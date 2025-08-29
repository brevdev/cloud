package v1

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	v1 "github.com/brevdev/cloud/v1"
)

func TestLambdaLabsCredential_GetReferenceID(t *testing.T) {
	cred := &LambdaLabsCredential{
		RefID:  "test-ref-id",
		APIKey: "test-api-key",
	}

	assert.Equal(t, "test-ref-id", cred.GetReferenceID())
}

func TestLambdaLabsCredential_GetAPIType(t *testing.T) {
	cred := &LambdaLabsCredential{}
	assert.Equal(t, v1.APITypeGlobal, cred.GetAPIType())
}

func TestLambdaLabsCredential_GetCloudProviderID(t *testing.T) {
	cred := &LambdaLabsCredential{}
	assert.Equal(t, v1.CloudProviderID("lambda-labs"), cred.GetCloudProviderID())
}

func TestLambdaLabsCredential_GetTenantID(t *testing.T) {
	cred := &LambdaLabsCredential{APIKey: "test-key"}
	tenantID, err := cred.GetTenantID()
	assert.NoError(t, err)
	assert.Contains(t, tenantID, "lambdalabs-")
}

func TestLambdaLabsCredential_MakeClient(t *testing.T) {
	cred := &LambdaLabsCredential{
		RefID:  "test-ref-id",
		APIKey: "test-api-key",
	}

	client, err := cred.MakeClient(context.Background(), "test-tenant")
	require.NoError(t, err)

	lambdaClient, ok := client.(*LambdaLabsClient)
	require.True(t, ok)
	assert.Equal(t, "test-ref-id", lambdaClient.refID)
	assert.Equal(t, "test-api-key", lambdaClient.apiKey)
}
