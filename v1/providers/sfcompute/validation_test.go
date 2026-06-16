package v1

import (
	"os"
	"testing"

	"github.com/brevdev/cloud/internal/validation"
	v1 "github.com/brevdev/cloud/v1"
)

func TestValidationFunctions(t *testing.T) {
	t.Parallel()
	checkSkip(t)
	apiKey := getAPIKey()

	config := validation.ProviderConfig{
		Credential: NewSFCCredential("validation-test", apiKey),
		StableIDs: []v1.InstanceTypeID{
			"richmond-noSub-h100",
		},
	}

	validation.RunValidationSuite(t, config)
}

func TestInstanceLifecycleValidation(t *testing.T) {
	t.Parallel()
	checkSkip(t)
	apiKey := getAPIKey()

	config := validation.ProviderConfig{
		Credential: NewSFCCredential("validation-test", apiKey),
		Location:   "yerba",
	}

	validation.RunInstanceLifecycleValidation(t, config)
}

func checkSkip(t *testing.T) {
	isValidationTest := os.Getenv("VALIDATION_TEST")
	if isValidationTest != "true" {
		t.Skip("VALIDATION_TEST is not set to true, skipping sfcompute validation tests")
	}

	apiKey := getAPIKey()
	if apiKey == "" {
		t.Fatal("SFCOMPUTE_API_KEY not set, but VALIDATION_TEST is set")
	}
}

func getAPIKey() string {
	return os.Getenv("SFCOMPUTE_API_KEY")
}
