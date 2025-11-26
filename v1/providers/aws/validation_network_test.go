package v1

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/brevdev/cloud/internal/validation"
)

var (
	isValidationTest = os.Getenv("VALIDATION_TEST")
	accessKeyID      = os.Getenv("AWS_ACCESS_KEY_ID")
	secretAccessKey  = os.Getenv("AWS_SECRET_ACCESS_KEY")
)

func TestAWSNetworkValidation(t *testing.T) {
	t.Skip("Skipping AWS Network validation tests")

	if isValidationTest == "" {
		t.Skip("VALIDATION_TEST is not set, skipping AWS Network validation tests")
	}

	if accessKeyID == "" || secretAccessKey == "" {
		t.Fatalf("AWS_ACCESS_KEY_ID and AWS_SECRET_ACCESS_KEY must be set")
	}

	config := validation.ProviderConfig{
		Location:   "us-east-1",
		Credential: NewAWSCredential(fmt.Sprintf("validation-%s", t.Name()), accessKeyID, secretAccessKey),
	}

	// Use the test name as the name of the VPC
	name := fmt.Sprintf("cloud-sdk-%s-%s", t.Name(), time.Now().UTC().Format("20060102150405"))

	validation.RunNetworkValidation(t, config, validation.NetworkValidationOpts{
		Name:                  name,
		RefID:                 name,
		CidrBlock:             "172.16.0.0/16",
		PublicSubnetCidrBlock: "172.16.0.0/24",
		Tags: map[string]string{
			"test": "TestNetworkValidation",
		},
	})
}
