package v1

import (
	"os"
	"testing"

	"github.com/brevdev/cloud/internal/validation"
)

func TestNetworkValidation(t *testing.T) {
	privateKeyPEMBase64 := os.Getenv("NEBIUS_PRIVATE_KEY_PEM_BASE64")
	publicKeyID := os.Getenv("NEBIUS_PUBLIC_KEY_ID")
	serviceAccountID := os.Getenv("NEBIUS_SERVICE_ACCOUNT_ID")
	projectID := "project-e00nrhefpr009ynkkzcgba" // eu-north1

	config := validation.ProviderConfig{
		Credential: NewNebiusCredential("validation-test", publicKeyID, privateKeyPEMBase64, serviceAccountID, projectID),
	}

	validation.RunNetworkValidation(t, config)
}
