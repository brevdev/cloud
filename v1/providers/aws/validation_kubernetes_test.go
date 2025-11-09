package v1

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/brevdev/cloud/internal/validation"
	v1 "github.com/brevdev/cloud/v1"
)

func TestAWSKubernetesValidation(t *testing.T) {
	if isValidationTest == "" {
		t.Skip("VALIDATION_TEST is not set, skipping AWS Kubernetes validation tests")
	}

	testUserPrivateKeyPEMBase64 := os.Getenv("TEST_USER_PRIVATE_KEY_PEM_BASE64")

	if accessKeyID == "" || secretAccessKey == "" {
		t.Fatalf("AWS_ACCESS_KEY_ID and AWS_SECRET_ACCESS_KEY must be set")
	}

	config := validation.ProviderConfig{
		Location:   "us-east-1",
		Credential: NewAWSCredential(fmt.Sprintf("validation-%s", t.Name()), accessKeyID, secretAccessKey),
	}

	// Use the test name as the name of the cluster and node group
	name := fmt.Sprintf("testcloudsdk-%s", time.Now().UTC().Format("20060102150405"))

	// Network CIDR
	networkCidr := "10.0.0.0/16"

	// Network subnets
	pubSubnet1 := validation.KubernetesValidationSubnetOpts{Name: "pub-subnet-1", RefID: "pub-subnet-1", CidrBlock: "10.0.0.0/19", SubnetType: v1.SubnetTypePublic}
	prvSubnet1 := validation.KubernetesValidationSubnetOpts{Name: "prv-subnet-1", RefID: "prv-subnet-1", CidrBlock: "10.0.32.0/19", SubnetType: v1.SubnetTypePrivate}
	pubSubnet2 := validation.KubernetesValidationSubnetOpts{Name: "pub-subnet-2", RefID: "pub-subnet-2", CidrBlock: "10.0.64.0/19", SubnetType: v1.SubnetTypePublic}
	prvSubnet2 := validation.KubernetesValidationSubnetOpts{Name: "prv-subnet-2", RefID: "prv-subnet-2", CidrBlock: "10.0.96.0/19", SubnetType: v1.SubnetTypePrivate}

	validation.RunKubernetesValidation(t, config, validation.KubernetesValidationOpts{
		Name:              name,
		RefID:             name,
		KubernetesVersion: "1.34",
		// Associate the VPC with the private subnets
		Subnets: []validation.KubernetesValidationSubnetOpts{prvSubnet1, prvSubnet2},
		NetworkOpts: &validation.KubernetesValidationNetworkOpts{
			Name:      name,
			RefID:     name,
			CidrBlock: networkCidr,
			// Build the network with all subnets
			Subnets: []validation.KubernetesValidationSubnetOpts{pubSubnet1, prvSubnet1, pubSubnet2, prvSubnet2},
		},
		NodeGroupOpts: &validation.KubernetesValidationNodeGroupOpts{
			Name:         name,
			RefID:        name,
			InstanceType: "t3.medium",
			DiskSize:     v1.NewBytes(20, v1.Gibibyte),
			MinNodeCount: 1,
			MaxNodeCount: 1,
		},
		UserOpts: &validation.KubernetesValidationUserOpts{
			Username:     "test-user",
			Role:         "cluster-admin",
			RSAPEMBase64: testUserPrivateKeyPEMBase64,
		},
		Tags: map[string]string{
			"test": t.Name(),
		},
	})
}
