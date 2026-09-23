package hyperstack

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/brevdev/cloud/internal/validation"
	v1 "github.com/brevdev/cloud/v1"
)

func TestValidationFunctions(t *testing.T) {
	checkValidationCredential(t)
	credential := validationCredential()

	validation.RunValidationSuite(t, validation.ProviderConfig{
		Credential: credential,
		StableIDs:  getStableInstanceTypeIDs(t, credential),
	})
}

func TestInstanceLifecycleValidation(t *testing.T) {
	checkValidationCredential(t)
	credential := validationCredential()

	validation.RunInstanceLifecycleValidation(t, validation.ProviderConfig{
		Credential: credential,
		StableIDs:  getStableInstanceTypeIDs(t, credential),
		CreateInstanceAttrs: v1.CreateInstanceAttrs{FirewallRules: v1.FirewallRules{
			IngressRules: []v1.FirewallRule{{
				FromPort: 22,
				ToPort:   22,
				IPRanges: []string{"0.0.0.0/0"},
			}},
		}},
	})
}

func TestGetLocations(t *testing.T) {
	checkValidationCredential(t)
	credential := validationCredential()

	client, err := credential.MakeClient(context.Background(), "")
	require.NoError(t, err)
	locations, err := client.GetLocations(context.Background(), v1.GetLocationsArgs{})
	require.NoError(t, err)
	require.NotEmpty(t, locations)
	for _, location := range locations {
		fmt.Println(location.Name)
	}
}

func TestGetInstanceTypes(t *testing.T) {
	checkValidationCredential(t)
	credential := validationCredential()

	client, err := credential.MakeClient(context.Background(), "")
	require.NoError(t, err)
	instanceTypes, err := client.GetInstanceTypes(context.Background(), v1.GetInstanceTypeArgs{})
	require.NoError(t, err)
	require.NotEmpty(t, instanceTypes)
}

func checkValidationCredential(t *testing.T) {
	t.Helper()
	if os.Getenv("HYPERSTACK_API_KEY") != "" {
		return
	}
	if os.Getenv("VALIDATION_TEST") != "" {
		t.Fatal("HYPERSTACK_API_KEY must be set when VALIDATION_TEST is set")
	}
	t.Skip("HYPERSTACK_API_KEY not set; skipping Hyperstack validation tests")
}

func validationCredential() *HyperstackCredential {
	credential := NewHyperstackCredential("validation-test", os.Getenv("HYPERSTACK_API_KEY"))
	if apiURL := os.Getenv("HYPERSTACK_API_URL"); apiURL != "" {
		credential.APIURL = apiURL
	}
	return credential
}

func getStableInstanceTypeIDs(t *testing.T, credential *HyperstackCredential) []v1.InstanceTypeID {
	t.Helper()
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()

	client, err := credential.MakeClient(ctx, "")
	require.NoError(t, err)
	instanceTypes, err := client.GetInstanceTypes(ctx, v1.GetInstanceTypeArgs{})
	require.NoError(t, err)
	require.NotEmpty(t, instanceTypes)

	stableIDs := make([]v1.InstanceTypeID, 0, len(instanceTypes))
	for _, instanceType := range instanceTypes {
		stableIDs = append(stableIDs, instanceType.ID)
	}
	return stableIDs
}
