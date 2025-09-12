//go:build scripts
// +build scripts

package scripts

import (
	"context"
	"fmt"
	"testing"

	v1 "github.com/brevdev/cloud/v1"
	nebius "github.com/brevdev/cloud/v1/providers/nebius"
)

const (
	privateKeyPEMBase64 = "test"
	publicKeyID         = "test"
	serviceAccountID    = "test"
	projectID           = "test"
)

func TestCreateVPC(t *testing.T) {
	if privateKeyPEMBase64 == "" || publicKeyID == "" || serviceAccountID == "" || projectID == "" {
		t.Fatalf("NEBIUS_PRIVATE_KEY_PEM_BASE64, NEBIUS_PUBLIC_KEY_ID, NEBIUS_SERVICE_ACCOUNT_ID, and NEBIUS_PROJECT_ID must be set")
	}

	nebiusClient, err := nebius.NewNebiusClient(context.Background(), "test", publicKeyID, privateKeyPEMBase64, serviceAccountID, projectID)
	if err != nil {
		t.Fatalf("failed to create Nebius client: %v", err)
	}

	vpc, err := nebiusClient.CreateVPC(context.Background(), v1.CreateVPCArgs{
		Name:      "cloud-sdk-test",
		RefID:     "cloud-sdk-test",
		CidrBlock: "172.16.0.0/16",
		Subnets: []v1.CreateSubnetArgs{
			{CidrBlock: "172.16.0.0/24", Type: v1.SubnetTypePublic},
			{CidrBlock: "172.16.1.0/24", Type: v1.SubnetTypePrivate},
			{CidrBlock: "172.16.2.0/24", Type: v1.SubnetTypePublic},
			{CidrBlock: "172.16.3.0/24", Type: v1.SubnetTypePrivate},
		},
	})
	if err != nil {
		t.Fatalf("failed to get VPC: %v", err)
	}

	fmt.Println(vpc)
}

func TestGetVPC(t *testing.T) {
	if privateKeyPEMBase64 == "" || publicKeyID == "" || serviceAccountID == "" || projectID == "" {
		t.Fatalf("NEBIUS_PRIVATE_KEY_PEM_BASE64, NEBIUS_PUBLIC_KEY_ID, NEBIUS_SERVICE_ACCOUNT_ID, and NEBIUS_PROJECT_ID must be set")
	}

	nebiusClient, err := nebius.NewNebiusClient(context.Background(), "test", publicKeyID, privateKeyPEMBase64, serviceAccountID, projectID)
	if err != nil {
		t.Fatalf("failed to create Nebius client: %v", err)
	}

	vpc, err := nebiusClient.GetVPC(context.Background(), v1.GetVPCArgs{
		ID: v1.CloudProviderResourceID("cloud-sdk-test"),
	})
	if err != nil {
		t.Fatalf("failed to get VPC: %v", err)
	}

	fmt.Println(vpc)
}

func TestDeleteVPC(t *testing.T) {
	if privateKeyPEMBase64 == "" || publicKeyID == "" || serviceAccountID == "" || projectID == "" {
		t.Fatalf("NEBIUS_PRIVATE_KEY_PEM_BASE64, NEBIUS_PUBLIC_KEY_ID, NEBIUS_SERVICE_ACCOUNT_ID, and NEBIUS_PROJECT_ID must be set")
	}

	nebiusClient, err := nebius.NewNebiusClient(context.Background(), "test", publicKeyID, privateKeyPEMBase64, serviceAccountID, projectID)
	if err != nil {
		t.Fatalf("failed to create Nebius client: %v", err)
	}

	err = nebiusClient.DeleteVPC(context.Background(), v1.DeleteVPCArgs{
		ID: v1.CloudProviderResourceID("vpcnetwork-u00r9ya5rc8wntbffr"),
	})
	if err != nil {
		t.Fatalf("failed to delete VPC: %v", err)
	}

	fmt.Println("VPC deleted")
}
