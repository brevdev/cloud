package v1

import (
	"os"
	"testing"

	"github.com/brevdev/cloud/internal/validation"
)

func TestInstanceLifecycleValidation(t *testing.T) {
	t.Parallel()
	checkSkip(t)
	apiKey := getAPIKey()

	config := validation.ProviderConfig{
		Credential: NewLaunchpadCredential("validation-test", apiKey, "https://launchpad.api.nvidia.com"),
	}

	validation.RunInstanceLifecycleValidation(t, config)
}

func checkSkip(t *testing.T) {
	apiKey := getAPIKey()
	isValidationTest := os.Getenv("VALIDATION_TEST")
	if apiKey == "" && isValidationTest != "" {
		t.Fatal("LAUNCHPAD_API_TOKEN not set, but VALIDATION_TEST is set")
	} else if apiKey == "" && isValidationTest == "" {
		t.Skip("LAUNCHPAD_API_TOKEN not set, skipping launchpad validation tests")
	}
}

func getAPIKey() string {
	return os.Getenv("LAUNCHPAD_API_TOKEN")
}
