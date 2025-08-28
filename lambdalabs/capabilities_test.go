package lambdalabs

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/brevdev/cloud"
)

func TestLambdaLabsClient_GetCapabilities(t *testing.T) {
	client := &LambdaLabsClient{}
	capabilities, err := client.GetCapabilities(context.Background())
	require.NoError(t, err)

	assert.Contains(t, capabilities, cloud.CapabilityCreateInstance)
	assert.Contains(t, capabilities, cloud.CapabilityTerminateInstance)
	assert.Contains(t, capabilities, cloud.CapabilityRebootInstance)
	assert.NotContains(t, capabilities, cloud.CapabilityStopStartInstance)
}

func TestLambdaLabsCredential_GetCapabilities(t *testing.T) {
	cred := &LambdaLabsCredential{}
	capabilities, err := cred.GetCapabilities(context.Background())
	require.NoError(t, err)

	assert.Contains(t, capabilities, cloud.CapabilityCreateInstance)
	assert.Contains(t, capabilities, cloud.CapabilityTerminateInstance)
	assert.Contains(t, capabilities, cloud.CapabilityRebootInstance)
	assert.NotContains(t, capabilities, cloud.CapabilityStopStartInstance)
}
