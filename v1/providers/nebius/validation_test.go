package v1

import (
	"os"
	"testing"

	"github.com/brevdev/cloud/internal/validation"
	v1 "github.com/brevdev/cloud/v1"
)

const (
	nebiusIsValidationTestEnvVar   = "VALIDATION_TEST"
	nebiusServiceAccountJSONEnvVar = "NEBIUS_SERVICE_ACCOUNT_JSON"
	nebiusTenantIDEnvVar           = "NEBIUS_TENANT_ID"
)

func TestValidationFunctions(t *testing.T) {
	t.Parallel()
	if os.Getenv(nebiusIsValidationTestEnvVar) == "" {
		t.Skipf("%s is not set, skipping Nebius validation tests", nebiusIsValidationTestEnvVar)
	}

	config := validation.ProviderConfig{
		Credential: newNebiusCredential(t),
		Location:   "eu-north1",
		StableIDs:  []v1.InstanceTypeID{"eu-north1-noSub-gpu-l40s-a.1gpu-8vcpu-32gb"},
	}

	validation.RunValidationSuite(t, config)
}

func TestInstanceLifecycleValidation(t *testing.T) {
	t.Parallel()
	if os.Getenv(nebiusIsValidationTestEnvVar) == "" {
		t.Skipf("%s is not set, skipping Nebius validation tests", nebiusIsValidationTestEnvVar)
	}

	config := validation.ProviderConfig{
		Credential: newNebiusCredential(t),
		Location:   "eu-north1",
		Tags:       ciRunTags(),
	}

	validation.RunInstanceLifecycleValidation(t, config)
}

// ciRunTags returns the CI run label when CI_RUN_ID is set (nil on local runs).
func ciRunTags() map[string]string {
	if id := os.Getenv("CI_RUN_ID"); id != "" {
		return map[string]string{CIRunIDLabel: id}
	}
	return nil
}

func newNebiusCredential(t *testing.T) *NebiusCredential {
	serviceAccountJSON := os.Getenv(nebiusServiceAccountJSONEnvVar)
	tenantID := os.Getenv(nebiusTenantIDEnvVar)

	var fail bool
	if serviceAccountJSON == "" {
		t.Logf("%s must be set", nebiusServiceAccountJSONEnvVar)
		fail = true
	}
	if tenantID == "" {
		t.Logf("%s must be set", nebiusTenantIDEnvVar)
		fail = true
	}

	if fail {
		t.Fatal("Missing environment variables")
	}
	return NewNebiusCredential("validation-test", serviceAccountJSON, tenantID)
}
