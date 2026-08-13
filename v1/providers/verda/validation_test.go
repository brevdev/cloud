package verda

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/brevdev/cloud/internal/validation"
	v1 "github.com/brevdev/cloud/v1"
	"github.com/stretchr/testify/require"
)

func TestValidationFunctions(t *testing.T) {
	checkValidationCredentials(t)
	credential := validationCredential()

	validation.RunValidationSuite(t, validation.ProviderConfig{
		Credential: credential,
		StableIDs:  getStableInstanceTypeIDs(t, credential),
	})
}

func TestGetInstance(t *testing.T) {
	checkValidationCredentials(t)
	credential := validationCredential()

	client, err := credential.MakeClient(context.Background(), "")
	require.NoError(t, err)

	instance, err := client.GetInstance(context.Background(), v1.CloudProviderInstanceID("b395e9e7-9a21-4ff1-a4b5-fb06d9652942"))
	require.NoError(t, err)
	require.NotNil(t, instance)
}

func TestInstanceLifecycleValidation(t *testing.T) {
	checkValidationCredentials(t)

	validation.RunInstanceLifecycleValidation(t, validation.ProviderConfig{
		Credential: validationCredential(),
	})
}

func checkValidationCredentials(t *testing.T) {
	t.Helper()
	clientID := os.Getenv("VERDA_CLIENT_ID")
	clientSecret := os.Getenv("VERDA_CLIENT_SECRET")
	if clientID != "" && clientSecret != "" {
		return
	}
	if os.Getenv("VALIDATION_TEST") != "" {
		t.Fatal("VERDA_CLIENT_ID and VERDA_CLIENT_SECRET must be set when VALIDATION_TEST is set")
	}
	t.Skip("VERDA_CLIENT_ID or VERDA_CLIENT_SECRET not set; skipping Verda validation tests")
}

func validationCredential() *VerdaCredential {
	credential := NewVerdaCredential(
		"validation-test",
		os.Getenv("VERDA_CLIENT_ID"),
		os.Getenv("VERDA_CLIENT_SECRET"),
	)
	if apiURL := os.Getenv("VERDA_API_URL"); apiURL != "" {
		credential.APIURL = apiURL
	}
	return credential
}

func getStableInstanceTypeIDs(t *testing.T, credential *VerdaCredential) []v1.InstanceTypeID {
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
