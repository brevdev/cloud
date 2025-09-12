package v1

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/brevdev/cloud/internal/validation"
)

var (
	privateKeyPEMBase64 = os.Getenv("NEBIUS_PRIVATE_KEY_PEM_BASE64")
	publicKeyID         = os.Getenv("NEBIUS_PUBLIC_KEY_ID")
	serviceAccountID    = os.Getenv("NEBIUS_SERVICE_ACCOUNT_ID")
	projectID           = os.Getenv("NEBIUS_PROJECT_ID")
)

func TestNetworkValidation(t *testing.T) {
	if privateKeyPEMBase64 == "" || publicKeyID == "" || serviceAccountID == "" || projectID == "" {
		t.Fatalf("NEBIUS_PRIVATE_KEY_PEM_BASE64, NEBIUS_PUBLIC_KEY_ID, NEBIUS_SERVICE_ACCOUNT_ID, and NEBIUS_PROJECT_ID must be set")
	}

	config := validation.ProviderConfig{
		Credential: NewNebiusCredential(fmt.Sprintf("validation-%s", t.Name()), publicKeyID, privateKeyPEMBase64, serviceAccountID, projectID),
	}

	// Use the test name as the name of the VPC
	name := fmt.Sprintf("cloud-sdk-%s-%s", t.Name(), time.Now().UTC().Format("20060102150405"))

	validation.RunNetworkValidation(t, config, validation.NetworkValidationOpts{
		Name:                  name,
		RefID:                 name,
		CidrBlock:             "172.16.0.0/16",
		PublicSubnetCidrBlock: "172.16.0.0/23",
		Tags: map[string]string{
			"test": "TestNetworkValidation",
		},
	})
}
