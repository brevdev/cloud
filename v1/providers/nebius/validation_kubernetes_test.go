package v1

import (
	"os"
	"testing"

	"github.com/brevdev/cloud/internal/validation"
)

func TestKubernetesValidation(t *testing.T) {
	privateKeyPEMBase64 := os.Getenv("NEBIUS_PRIVATE_KEY_PEM_BASE64")
	publicKeyID := os.Getenv("NEBIUS_PUBLIC_KEY_ID")
	serviceAccountID := os.Getenv("NEBIUS_SERVICE_ACCOUNT_ID")
	projectID := os.Getenv("NEBIUS_PROJECT_ID")

	if privateKeyPEMBase64 == "" || publicKeyID == "" || serviceAccountID == "" || projectID == "" {
		t.Fatalf("NEBIUS_PRIVATE_KEY_PEM_BASE64, NEBIUS_PUBLIC_KEY_ID, NEBIUS_SERVICE_ACCOUNT_ID, and NEBIUS_PROJECT_ID must be set")
	}

	config := validation.ProviderConfig{
		Credential: NewNebiusCredential("validation-test", publicKeyID, privateKeyPEMBase64, serviceAccountID, projectID),
	}

	validation.RunKubernetesValidation(t, config, validation.KubernetesValidationOpts{
		Name:              "cloud-sdk-test",
		RefID:             "cloud-sdk-test",
		KubernetesVersion: "1.31",
		NetworkOpts: &validation.KubernetesValidationNetworkOpts{
			Name:                   "cloud-sdk-test",
			RefID:                  "cloud-sdk-test",
			Location:               "us-central1",
			CidrBlock:              "172.16.0.0/16",
			PublicSubnetCidrBlock:  "172.16.0.0/19", // todo: validate greater than /24 (IP count: 256)
			PrivateSubnetCidrBlock: "172.16.32.0/19",
		},
	})
}
