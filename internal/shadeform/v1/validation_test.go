package v1

import (
	"context"
	openapi "github.com/brevdev/cloud/internal/shadeform/gen/shadeform"
	"github.com/brevdev/cloud/internal/validation"
	"github.com/brevdev/cloud/pkg/ssh"
	v1 "github.com/brevdev/cloud/pkg/v1"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
	"time"
)

func TestValidationFunctions(t *testing.T) {
	t.Parallel()
	checkSkip(t)
	apiKey := getAPIKey()

	config := validation.ProviderConfig{
		Credential: NewShadeformCredential("validation-test", apiKey),
		StableIDs:  []v1.InstanceTypeID{"datacrunch_B200_helsinki-finland-5", "massedcompute_L40_desmoines-usa-1"},
	}

	validation.RunValidationSuite(t, config)
}

func TestInstanceLifecycleValidation(t *testing.T) {
	t.Parallel()
	checkSkip(t)
	apiKey := getAPIKey()

	config := validation.ProviderConfig{
		Credential: NewShadeformCredential("validation-test", apiKey),
	}

	validation.RunInstanceLifecycleValidation(t, config)
}

func TestInstanceTypeFilter(t *testing.T) {
	t.Parallel()
	checkSkip(t)
	apiKey := getAPIKey()

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Minute)
	defer cancel()

	client := NewShadeformClient("validation-test", apiKey)
	client.WithConfiguration(Configuration{
		AllowedInstanceTypes: map[openapi.Cloud]map[string]bool{
			openapi.DATACRUNCH: {
				"B200": true,
			},
		},
	})

	types, err := client.GetInstanceTypes(ctx, v1.GetInstanceTypeArgs{})
	require.NoError(t, err)
	require.NotEmpty(t, types, "Should have instance types")
	require.True(t, len(types) == 1, "Instance types should return only one entry")
	require.True(t, types[0].Type == "datacrunch_B200", "returned instance type does not match expectations")

	if !types[0].IsAvailable {
		return
	}

	instance, err := client.CreateInstance(ctx, v1.CreateInstanceAttrs{
		RefID:        uuid.New().String(),
		InstanceType: types[0].Type,
		Location:     types[0].Location,
		PublicKey:    ssh.GetTestPublicKey(),
		Name:         "test_name",
		FirewallRules: v1.FirewallRules{
			EgressRules: []v1.FirewallRule{
				{
					ID:       "test-rule1",
					FromPort: 80,
					ToPort:   8080,
					IPRanges: []string{"127.0.0.1", "10.0.0.0/24"},
				},
				{
					ID:       "test-rule2",
					FromPort: 5432,
					ToPort:   5432,
					IPRanges: []string{"127.0.0.1", "10.0.0.0/24"},
				},
			},
			IngressRules: []v1.FirewallRule{
				{
					ID:       "test-rule3",
					FromPort: 80,
					ToPort:   8080,
					IPRanges: []string{"127.0.0.1", "10.0.0.0/24"},
				},
				{
					ID:       "test-rule4",
					FromPort: 5432,
					ToPort:   5432,
					IPRanges: []string{"127.0.0.1", "10.0.0.0/24"},
				},
			},
		},
	})

	if err != nil {
		t.Fatalf("ValidateCreateInstance failed: %v", err)
	}
	require.NotNil(t, instance)

	t.Run("ValidateSSHAccessible", func(t *testing.T) {
		err := v1.ValidateInstanceSSHAccessible(ctx, client, instance, ssh.GetTestPrivateKey())
		require.NoError(t, err, "ValidateSSHAccessible should pass")
	})

	t.Run("ValidateTerminateInstance", func(t *testing.T) {
		err := v1.ValidateTerminateInstance(ctx, client, instance)
		require.NoError(t, err, "ValidateTerminateInstance should pass")
	})
}

func checkSkip(t *testing.T) {
	apiKey := getAPIKey()
	isValidationTest := os.Getenv("VALIDATION_TEST")
	if apiKey == "" && isValidationTest != "" {
		t.Fatal("SHADEFORM_API_KEY not set, but VALIDATION_TEST is set")
	} else if apiKey == "" && isValidationTest == "" {
		t.Skip("SHADEFORM_API_KEY not set, skipping shadeform validation tests")
	}
}

func getAPIKey() string {
	return os.Getenv("SHADEFORM_API_KEY")
}
