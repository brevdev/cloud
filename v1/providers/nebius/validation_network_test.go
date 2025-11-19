package v1

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/brevdev/cloud/internal/validation"
)

var (
	isValidationTest   = os.Getenv("VALIDATION_TEST")
	serviceAccountJSON = os.Getenv("NEBIUS_SERVICE_ACCOUNT_JSON")
	tenantID           = os.Getenv("NEBIUS_TENANT_ID")
)

func TestNetworkValidation(t *testing.T) {
	if isValidationTest == "" {
		t.Skip("VALIDATION_TEST is not set, skipping Nebius Network validation tests")
	}

	if serviceAccountJSON == "" || tenantID == "" {
		t.Skip("NEBIUS_SERVICE_ACCOUNT_JSON and NEBIUS_TENANT_ID must be set")
	}

	config := validation.ProviderConfig{
		Credential: NewNebiusCredential(fmt.Sprintf("validation-%s", t.Name()), serviceAccountJSON, tenantID),
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
