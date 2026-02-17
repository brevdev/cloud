package v1

import (
	"os"
	"testing"

	"github.com/brevdev/cloud/internal/validation"
	v1 "github.com/brevdev/cloud/v1"
)

func TestValidationFunctions(t *testing.T) {
	checkSkip(t)
	apiKey := getAPIKey()

	config := validation.ProviderConfig{
		Credential: NewSFCCredential("validation-test", apiKey),
		StableIDs: []v1.InstanceTypeID{
			"h100v_hayesvalley",
			"h100v_yerba",
			"h100v_richmond",
		},
	}

	validation.RunValidationSuite(t, config)
}

func checkSkip(t *testing.T) {
	apiKey := getAPIKey()
	isValidation := os.Getenv("VALIDATION_TEST")
	if apiKey == "" && isValidation != "true" {
		t.Fatal("SFCOMPUTE_API_KEY not set, but VALIDATION_TEST is set")
	} else if apiKey == "" && isValidation == "false" {
		t.Skip("SFCOMPUTE_API_KEY not set, skipping sfcompute validation tests")
	}
}

func getAPIKey() string {
	return os.Getenv("SFCOMPUTE_API_KEY")
}
