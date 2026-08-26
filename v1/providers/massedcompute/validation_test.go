package massedcompute

import (
	"context"
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
	})
}

func checkValidationCredential(t *testing.T) {
	t.Helper()
	if os.Getenv("MASSED_COMPUTE_API_TOKEN") != "" {
		return
	}
	if os.Getenv("VALIDATION_TEST") != "" {
		t.Fatal("MASSED_COMPUTE_API_TOKEN must be set when VALIDATION_TEST is set")
	}
	t.Skip("MASSED_COMPUTE_API_TOKEN not set; skipping Massed Compute validation tests")
}

func validationCredential() *MassedComputeCredential {
	credential := NewMassedComputeCredential("validation-test", os.Getenv("MASSED_COMPUTE_API_TOKEN"))
	if apiURL := os.Getenv("MASSED_COMPUTE_API_URL"); apiURL != "" {
		credential.APIURL = apiURL
	}
	return credential
}

func getStableInstanceTypeIDs(t *testing.T, credential *MassedComputeCredential) []v1.InstanceTypeID {
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
