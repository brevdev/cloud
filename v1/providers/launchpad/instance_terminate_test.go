package v1

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/brevdev/cloud/internal/validation"
	v1 "github.com/brevdev/cloud/v1"
)

func TestLaunchpadClient_TerminateInstance(t *testing.T) {
	t.Parallel()
	checkSkip(t)
	apiKey := getAPIKey()

	config := validation.ProviderConfig{
		Credential: NewLaunchpadCredential("brev-cloud-sdk-test", apiKey, "https://stage.launchpad.api.nvidia.com"),
	}

	client, err := config.Credential.MakeClient(context.Background(), config.Location)
	if err != nil {
		t.Fatalf("failed to make client: %v", err)
	}

	err = client.TerminateInstance(context.Background(), v1.CloudProviderInstanceID("1f507fb0-a90e-4d9e-89da-1ccd6abe4ce1"))
	require.NoError(t, err)
}
