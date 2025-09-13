package v1

import (
	"context"
	"fmt"
	"os"
	"testing"

	v1 "github.com/brevdev/cloud/v1"
)

func TestCreateVPC(t *testing.T) {
	privateKeyPEMBase64 := os.Getenv("NEBIUS_PRIVATE_KEY_PEM_BASE64")
	publicKeyID := os.Getenv("NEBIUS_PUBLIC_KEY_ID")
	serviceAccountID := os.Getenv("NEBIUS_SERVICE_ACCOUNT_ID")

	projectID := "project-e00nrhefpr009ynkkzcgba" // eu-north1

	nebiusClient, err := NewNebiusClient(context.Background(), "test", publicKeyID, privateKeyPEMBase64, serviceAccountID, projectID)
	if err != nil {
		t.Fatalf("failed to create Nebius client: %v", err)
	}

	vpc, err := nebiusClient.CreateVPC(context.Background(), v1.CreateVPCArgs{
		Name:      "cloud-sdk-test",
		RefID:     "cloud-sdk-test",
		Location:  "eu-north1",
		CidrBlock: "10.0.0.0/16",
		Subnets: []v1.CreateSubnetArgs{
			{CidrBlock: "10.0.0.0/24", Type: v1.SubnetTypePublic},
			{CidrBlock: "10.0.1.0/24", Type: v1.SubnetTypePrivate},
			{CidrBlock: "10.0.2.0/24", Type: v1.SubnetTypePublic},
			{CidrBlock: "10.0.3.0/24", Type: v1.SubnetTypePrivate},
		},
	})
	if err != nil {
		t.Fatalf("failed to get VPC: %v", err)
	}

	fmt.Println(vpc)
}

func TestGetVPC(t *testing.T) {
	privateKeyPEMBase64 := os.Getenv("NEBIUS_PRIVATE_KEY_PEM_BASE64")
	publicKeyID := os.Getenv("NEBIUS_PUBLIC_KEY_ID")
	serviceAccountID := os.Getenv("NEBIUS_SERVICE_ACCOUNT_ID")

	projectID := "project-e00nrhefpr009ynkkzcgba" // eu-north1

	nebiusClient, err := NewNebiusClient(context.Background(), "test", publicKeyID, privateKeyPEMBase64, serviceAccountID, projectID)
	if err != nil {
		t.Fatalf("failed to create Nebius client: %v", err)
	}

	vpc, err := nebiusClient.GetVPC(context.Background(), v1.GetVPCArgs{
		CloudID: "vpcnetwork-e00g39sp5rk783qf2q",
	})
	if err != nil {
		t.Fatalf("failed to get VPC: %v", err)
	}

	fmt.Println(vpc)
}

func TestDeleteVPC(t *testing.T) {
	privateKeyPEMBase64 := os.Getenv("NEBIUS_PRIVATE_KEY_PEM_BASE64")
	publicKeyID := os.Getenv("NEBIUS_PUBLIC_KEY_ID")
	serviceAccountID := os.Getenv("NEBIUS_SERVICE_ACCOUNT_ID")

	projectID := "project-e00nrhefpr009ynkkzcgba" // eu-north1

	nebiusClient, err := NewNebiusClient(context.Background(), "test", publicKeyID, privateKeyPEMBase64, serviceAccountID, projectID)
	if err != nil {
		t.Fatalf("failed to create Nebius client: %v", err)
	}

	err = nebiusClient.DeleteVPC(context.Background(), v1.DeleteVPCArgs{
		VPC: &v1.VPC{
			RefID: "cloud-sdk-test",
		},
	})
	if err != nil {
		t.Fatalf("failed to delete VPC: %v", err)
	}

	fmt.Println("VPC deleted")
}
