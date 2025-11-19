package v1

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/brevdev/cloud/internal/validation"
	v1 "github.com/brevdev/cloud/v1"
)

func TestKubernetesValidation(t *testing.T) {
	isValidationTest := os.Getenv("VALIDATION_TEST")
	if isValidationTest == "" {
		t.Skip("VALIDATION_TEST is not set, skipping Nebius Kubernetes validation tests")
	}

	testUserPrivateKeyPEMBase64 := os.Getenv("TEST_USER_PRIVATE_KEY_PEM_BASE64")
	serviceAccountJSON := os.Getenv("NEBIUS_SERVICE_ACCOUNT_JSON")
	tenantID := os.Getenv("NEBIUS_TENANT_ID")

	if serviceAccountJSON == "" || tenantID == "" {
		t.Fatalf("NEBIUS_SERVICE_ACCOUNT_JSON and NEBIUS_TENANT_ID must be set")
	}

	config := validation.ProviderConfig{
		Credential: NewNebiusCredential(fmt.Sprintf("validation-%s", t.Name()), serviceAccountJSON, tenantID),
	}

	// Use the test name as the name of the cluster and node group
	name := fmt.Sprintf("cloud-sdk-%s-%s", t.Name(), time.Now().UTC().Format("20060102150405"))

	// Network CIDR
	networkCidr := "10.0.0.0/16"

	// Network subnets
	pubSubnet1 := validation.KubernetesValidationSubnetOpts{Name: "pub-subnet-1", RefID: "pub-subnet-1", CidrBlock: "10.0.0.0/19", SubnetType: v1.SubnetTypePublic}
	prvSubnet1 := validation.KubernetesValidationSubnetOpts{Name: "prv-subnet-1", RefID: "prv-subnet-1", CidrBlock: "10.0.32.0/19", SubnetType: v1.SubnetTypePrivate}

	validation.RunKubernetesValidation(t, config, validation.KubernetesValidationOpts{
		Name:              name,
		RefID:             name,
		KubernetesVersion: "1.31",
		// Associate the VPC with the private subnets
		Subnets: []validation.KubernetesValidationSubnetOpts{prvSubnet1},
		NetworkOpts: &validation.KubernetesValidationNetworkOpts{
			Name:      name,
			RefID:     name,
			CidrBlock: networkCidr,
			// Build the network with all subnets
			Subnets: []validation.KubernetesValidationSubnetOpts{pubSubnet1, prvSubnet1},
		},
		NodeGroupOpts: &validation.KubernetesValidationNodeGroupOpts{
			Name:         name,
			RefID:        name,
			InstanceType: "cpu-d3.4vcpu-16gb",
			DiskSize:     v1.NewBytes(64, v1.Gibibyte),
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
