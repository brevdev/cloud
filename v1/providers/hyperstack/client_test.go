package hyperstack

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	v1 "github.com/brevdev/cloud/v1"
)

func TestHyperstackCredential(t *testing.T) {
	credential := NewHyperstackCredential("credential-ref", "api-key")

	assert.Equal(t, DefaultAPIURL, credential.APIURL)
	assert.Equal(t, v1.CloudProviderID(CloudProviderID), credential.GetCloudProviderID())
	assert.Equal(t, v1.APITypeGlobal, credential.GetAPIType())
	assert.Equal(t, "credential-ref", credential.GetReferenceID())
	require.NoError(t, credential.Validate())

	tenantID, err := credential.GetTenantID()
	require.NoError(t, err)
	assert.NotEmpty(t, tenantID)
	assert.NotContains(t, tenantID, credential.APIKey)

	capabilities, err := credential.GetCapabilities(context.Background())
	require.NoError(t, err)
	assert.Equal(t, getCapabilities(), capabilities)

	invalid := NewHyperstackCredential("credential-ref", "")
	require.Error(t, invalid.Validate())
}
