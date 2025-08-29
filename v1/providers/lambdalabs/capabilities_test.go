package v1

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	v1 "github.com/brevdev/cloud/v1"
)

func TestLambdaLabsClient_GetCapabilities(t *testing.T) {
	client := &LambdaLabsClient{}
	capabilities, err := client.GetCapabilities(context.Background())
	require.NoError(t, err)

	assert.Contains(t, capabilities, v1.CapabilityCreateInstance)
	assert.Contains(t, capabilities, v1.CapabilityTerminateInstance)
	assert.Contains(t, capabilities, v1.CapabilityRebootInstance)
	assert.NotContains(t, capabilities, v1.CapabilityStopStartInstance)
}

func TestLambdaLabsCredential_GetCapabilities(t *testing.T) {
	cred := &LambdaLabsCredential{}
	capabilities, err := cred.GetCapabilities(context.Background())
	require.NoError(t, err)

	assert.Contains(t, capabilities, v1.CapabilityCreateInstance)
	assert.Contains(t, capabilities, v1.CapabilityTerminateInstance)
	assert.Contains(t, capabilities, v1.CapabilityRebootInstance)
	assert.NotContains(t, capabilities, v1.CapabilityStopStartInstance)
}
