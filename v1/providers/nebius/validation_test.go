package v1

import (
	"os"
	"testing"

	"github.com/brevdev/cloud/internal/validation"
	v1 "github.com/brevdev/cloud/v1"
)

var (
	nebiusIsValidationTest   = os.Getenv("VALIDATION_TEST")
	nebiusServiceAccountJSON = os.Getenv("NEBIUS_SERVICE_ACCOUNT_JSON")
	nebiusTenantID           = os.Getenv("NEBIUS_TENANT_ID")
)

func TestValidationFunctions(t *testing.T) {
	t.Parallel()
	checkSkip(t)

	config := validation.ProviderConfig{
		Credential: NewNebiusCredential("validation-test", nebiusServiceAccountJSON, nebiusTenantID),
		StableIDs:  []v1.InstanceTypeID{"gpu-l40s-a.1gpu-8vcpu-32gb"},
	}

	validation.RunValidationSuite(t, config)
}

func TestInstanceLifecycleValidation(t *testing.T) {
	t.Parallel()
	checkSkip(t)

	config := validation.ProviderConfig{
		Credential: NewNebiusCredential("validation-test", nebiusServiceAccountJSON, nebiusTenantID),
	}

	validation.RunInstanceLifecycleValidation(t, config)
}

func checkSkip(t *testing.T) {
	if nebiusIsValidationTest == "" {
		t.Skip("VALIDATION_TEST is not set, skipping Nebius validation tests")
	}

	if nebiusServiceAccountJSON == "" || nebiusTenantID == "" {
		t.Fatal("NEBIUS_SERVICE_ACCOUNT_JSON and NEBIUS_TENANT_ID must be set")
	}
}
