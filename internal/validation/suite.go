package validation

import (
	"context"
	"testing"
	"time"

	"github.com/brevdev/sdk/cloud"
	"github.com/brevdev/sdk/internal/ssh"
	"github.com/stretchr/testify/require"
)

type ProviderConfig struct {
	Location   string
	StableIDs  []cloud.InstanceTypeID
	Credential cloud.CloudCredential
}

func RunValidationSuite(t *testing.T, config ProviderConfig) {
	if testing.Short() {
		t.Skip("Skipping validation tests in short mode")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
	defer cancel()

	client, err := config.Credential.MakeClient(ctx, config.Location)
	if err != nil {
		t.Fatalf("Failed to create client for %s: %v", config.Credential.GetCloudProviderID(), err)
	}

	t.Run("ValidateGetLocations", func(t *testing.T) {
		err := cloud.ValidateGetLocations(ctx, client)
		require.NoError(t, err, "ValidateGetLocations should pass")
	})

	t.Run("ValidateGetInstanceTypes", func(t *testing.T) {
		err := cloud.ValidateGetInstanceTypes(ctx, client)
		require.NoError(t, err, "ValidateGetInstanceTypes should pass")
	})

	t.Run("ValidateRegionalInstanceTypes", func(t *testing.T) {
		err := cloud.ValidateLocationalInstanceTypes(ctx, client)
		require.NoError(t, err, "ValidateRegionalInstanceTypes should pass")
	})

	t.Run("ValidateStableInstanceTypeIDs", func(t *testing.T) {
		err = cloud.ValidateStableInstanceTypeIDs(ctx, client, config.StableIDs)
		require.NoError(t, err, "ValidateStableInstanceTypeIDs should pass")
	})
}

func RunInstanceLifecycleValidation(t *testing.T, config ProviderConfig) {
	if testing.Short() {
		t.Skip("Skipping validation tests in short mode")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Minute)
	defer cancel()

	client, err := config.Credential.MakeClient(ctx, config.Location)
	if err != nil {
		t.Fatalf("Failed to create client for %s: %v", config.Credential.GetCloudProviderID(), err)
	}
	capabilities, err := client.GetCapabilities(ctx)
	require.NoError(t, err)

	types, err := client.GetInstanceTypes(ctx, cloud.GetInstanceTypeArgs{})
	require.NoError(t, err)
	require.NotEmpty(t, types, "Should have instance types")

	locations, err := client.GetLocations(ctx, cloud.GetLocationsArgs{})
	require.NoError(t, err)
	require.NotEmpty(t, locations, "Should have locations")

	t.Run("ValidateCreateInstance", func(t *testing.T) {
		attrs := cloud.CreateInstanceAttrs{}
		for _, typ := range types {
			if typ.IsAvailable {
				attrs.InstanceType = typ.Type
				attrs.Location = typ.Location
				attrs.PublicKey = ssh.GetTestPublicKey()
				break
			}
		}
		instance, err := cloud.ValidateCreateInstance(ctx, client, attrs)
		if err != nil {
			t.Fatalf("ValidateCreateInstance failed: %v", err)
		}
		require.NotNil(t, instance)

		defer func() {
			if instance != nil {
				_ = client.TerminateInstance(ctx, instance.CloudID)
			}
		}()

		t.Run("ValidateListCreatedInstance", func(t *testing.T) {
			err := cloud.ValidateListCreatedInstance(ctx, client, instance)
			require.NoError(t, err, "ValidateListCreatedInstance should pass")
		})

		t.Run("ValidateSSHAccessible", func(t *testing.T) {
			err := cloud.ValidateInstanceSSHAccessible(ctx, client, instance, ssh.GetTestPrivateKey())
			require.NoError(t, err, "ValidateSSHAccessible should pass")
		})

		instance, err = client.GetInstance(ctx, instance.CloudID)
		require.NoError(t, err)

		t.Run("ValidateInstanceImage", func(t *testing.T) {
			err := cloud.ValidateInstanceImage(ctx, *instance, ssh.GetTestPrivateKey())
			require.NoError(t, err, "ValidateInstanceImage should pass")
		})

		if capabilities.IsCapable(cloud.CapabilityStopStartInstance) && instance.Stoppable {
			t.Run("ValidateStopStartInstance", func(t *testing.T) {
				err := cloud.ValidateStopStartInstance(ctx, client, instance)
				require.NoError(t, err, "ValidateStopStartInstance should pass")
			})
		}

		t.Run("ValidateTerminateInstance", func(t *testing.T) {
			err := cloud.ValidateTerminateInstance(ctx, client, instance)
			require.NoError(t, err, "ValidateTerminateInstance should pass")
		})
	})
}
