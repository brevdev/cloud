package v2

import (
	"os"
	"testing"

	"github.com/brevdev/cloud/internal/validation"
	v1 "github.com/brevdev/cloud/v1"
)

func TestValidationFunctions(t *testing.T) {
	checkSkip(t)

	config := validation.ProviderConfig{
		Credential: NewSFCCredentialV2("validation-test", getAPIKey(), getOrganization(), getWorkspace()),
		StableIDs: []v1.InstanceTypeID{
			h100InstanceTypeMetadata.instanceTypeID,
		},
	}

	validation.RunValidationSuite(t, config)
}

func TestInstanceLifecycleValidation(t *testing.T) {
	checkSkip(t)

	config := validation.ProviderConfig{
		Credential: NewSFCCredentialV2("validation-test", getAPIKey(), getOrganization(), getWorkspace()),
		Location:   sfcLocation,
	}

	validation.RunInstanceLifecycleValidation(t, config)
}

func checkSkip(t *testing.T) {
	t.Helper()
	apiKey := getAPIKey()
	isValidationTest := os.Getenv("VALIDATION_TEST")
	if apiKey == "" && isValidationTest != "" {
		t.Fatal("SFCOMPUTE_API_KEY not set, but VALIDATION_TEST is set")
	} else if apiKey == "" {
		t.Skip("SFCOMPUTE_API_KEY not set, skipping sfcomputev2 validation tests")
	}
}

func getAPIKey() string {
	return os.Getenv("SFCOMPUTE_API_KEY")
}

func getOrganization() string {
	return os.Getenv("SFCOMPUTE_ORGANIZATION")
}

func getWorkspace() string {
	return os.Getenv("SFCOMPUTE_WORKSPACE")
}
