package v1

import (
	"context"
	"fmt"
	"os"
	"testing"

	v1 "github.com/brevdev/cloud/v1"
)

func TestDeleteVPC(t *testing.T) {
	awsClient, err := NewAWSClient("test",
		os.Getenv("AWS_ACCESS_KEY_ID"),
		os.Getenv("AWS_SECRET_ACCESS_KEY"),
		"us-east-1",
	)
	if err != nil {
		t.Fatalf("failed to create AWS client: %v", err)
	}

	err = awsClient.DeleteVPC(context.Background(), v1.DeleteVPCArgs{
		ID: v1.CloudProviderResourceID("vpc-0cea13ba205f6f523"),
	})
	if err != nil {
		t.Fatalf("failed to delete VPC: %v", err)
	}

	fmt.Println("VPC deleted")
}

func TestCreateVPC(t *testing.T) {
	awsClient, err := NewAWSClient("test",
		os.Getenv("AWS_ACCESS_KEY_ID"),
		os.Getenv("AWS_SECRET_ACCESS_KEY"),
		"us-east-1",
	)
	if err != nil {
		t.Fatalf("failed to create AWS client: %v", err)
	}

	vpc, err := awsClient.CreateVPC(context.Background(), v1.CreateVPCArgs{
		Name:      "cloud-sdk-test",
		RefID:     "cloud-sdk-test",
		CidrBlock: "10.0.0.0/16",
		Subnets: []v1.CreateSubnetArgs{
			{CidrBlock: "10.0.0.0/24", Type: v1.SubnetTypePublic},
			{CidrBlock: "10.0.1.0/24", Type: v1.SubnetTypePrivate},
			{CidrBlock: "10.0.2.0/24", Type: v1.SubnetTypePublic},
			{CidrBlock: "10.0.3.0/24", Type: v1.SubnetTypePrivate},
		},
	})
	if err != nil {
		t.Fatalf("failed to create VPC: %v", err)
	}

	vpc, err = awsClient.GetVPC(context.Background(), v1.GetVPCArgs{
		ID: vpc.ID,
	})
	if err != nil {
		t.Fatalf("failed to get VPC: %v", err)
	}

	err = awsClient.DeleteVPC(context.Background(), v1.DeleteVPCArgs{
		ID: vpc.ID,
	})
	if err != nil {
		t.Fatalf("failed to delete VPC: %v", err)
	}
}
