package v1

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/brevdev/cloud/internal/validation"
	v1 "github.com/brevdev/cloud/v1"
)

func TestLaunchpadClient_ListInstances(t *testing.T) {
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

	instances, err := client.ListInstances(context.Background(), v1.ListInstancesArgs{})
	require.NoError(t, err)
	require.Len(t, instances, 0)
}
