package v1

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/brevdev/cloud/internal/validation"
)

func TestKubernetesValidation(t *testing.T) {
	testUserPrivateKeyPEMBase64 := os.Getenv("TEST_USER_PRIVATE_KEY_PEM_BASE64")

	if privateKeyPEMBase64 == "" || publicKeyID == "" || serviceAccountID == "" || projectID == "" {
		t.Fatalf("NEBIUS_PRIVATE_KEY_PEM_BASE64, NEBIUS_PUBLIC_KEY_ID, NEBIUS_SERVICE_ACCOUNT_ID, and NEBIUS_PROJECT_ID must be set")
	}

	config := validation.ProviderConfig{
		Credential: NewNebiusCredential(fmt.Sprintf("validation-%s", t.Name()), publicKeyID, privateKeyPEMBase64, serviceAccountID, projectID),
	}

	// Use the test name as the name of the cluster and node group
	name := fmt.Sprintf("cloud-sdk-%s-%s", t.Name(), time.Now().UTC().Format("20060102150405"))

	validation.RunKubernetesValidation(t, config, validation.KubernetesValidationOpts{
		Name:              name,
		RefID:             name,
		KubernetesVersion: "1.31",
		NetworkOpts: &validation.KubernetesValidationNetworkOpts{
			Name:                   name,
			RefID:                  name,
			CidrBlock:              "172.16.0.0/16",
			PublicSubnetCidrBlock:  "172.16.0.0/19",
			PrivateSubnetCidrBlock: "172.16.32.0/19",
		},
		NodeGroupOpts: &validation.KubernetesValidationNodeGroupOpts{
			Name:         name,
			RefID:        name,
			InstanceType: "cpu-d3.4vcpu-16gb",
			DiskSizeGiB:  64,
			MinNodeCount: 1,
			MaxNodeCount: 1,
		},
		UserOpts: &validation.KubernetesValidationUserOpts{
			Username:     "test-user",
			Role:         "cluster-admin",
			RSAPEMBase64: testUserPrivateKeyPEMBase64,
		},
		Tags: map[string]string{
			"test": "TestKubernetesValidation",
		},
	})
}
