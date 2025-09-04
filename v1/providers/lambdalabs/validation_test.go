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
		Credential: NewLambdaLabsCredential("validation-test", apiKey),
		StableIDs:  []v1.InstanceTypeID{"us-west-1-noSub-gpu_8x_a100_80gb_sxm4", "us-east-1-noSub-gpu_8x_a100_80gb_sxm4"},
	}

	validation.RunValidationSuite(t, config)
}

func TestInstanceLifecycleValidation(t *testing.T) {
	checkSkip(t)
	apiKey := getAPIKey()

	config := validation.ProviderConfig{
		Credential: NewLambdaLabsCredential("validation-test", apiKey),
	}

	validation.RunInstanceLifecycleValidation(t, config)
}

func checkSkip(t *testing.T) {
	apiKey := getAPIKey()
	isValidationTest := os.Getenv("VALIDATION_TEST")
	if apiKey == "" && isValidationTest != "" {
		t.Fatal("LAMBDALABS_API_KEY not set, but VALIDATION_TEST is set")
	} else if apiKey == "" && isValidationTest == "" {
		t.Skip("LAMBDALABS_API_KEY not set, skipping LambdaLabs validation tests")
	}
}

func getAPIKey() string {
	return os.Getenv("LAMBDALABS_API_KEY")
}
