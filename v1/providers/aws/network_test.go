package v1

import (
	"context"
	"fmt"
	"testing"

	v1 "github.com/brevdev/cloud/v1"
)

func TestCreateVPC(t *testing.T) {
	awsClient, err := NewAWSClient("test", accessKeyID, secretAccessKey, "us-east-1")
	if err != nil {
		t.Fatalf("failed to create AWS client: %v", err)
	}

	vpc, err := awsClient.CreateVPC(context.Background(), v1.CreateVPCArgs{
		Name:      "cloud-sdk-test",
		RefID:     "cloud-sdk-test",
		CidrBlock: "10.0.0.0/16",
		Subnets: []v1.CreateSubnetArgs{
			{CidrBlock: "10.0.0.0/19", Type: v1.SubnetTypePublic},
			{CidrBlock: "10.0.32.0/19", Type: v1.SubnetTypePrivate},
			// {CidrBlock: "10.0.2.0/24", Type: v1.SubnetTypePublic},
			// {CidrBlock: "10.0.3.0/24", Type: v1.SubnetTypePrivate},
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

	// err = awsClient.DeleteVPC(context.Background(), v1.DeleteVPCArgs{
	// 	ID: vpc.ID,
	// })
	// if err != nil {
	// 	t.Fatalf("failed to delete VPC: %v", err)
	// }
}

func TestDeleteVPC(t *testing.T) {
	awsClient, err := NewAWSClient("test", accessKeyID, secretAccessKey, "us-east-1")
	if err != nil {
		t.Fatalf("failed to create AWS client: %v", err)
	}

	err = awsClient.DeleteVPC(context.Background(), v1.DeleteVPCArgs{
		ID: v1.CloudProviderResourceID("vpc-0e7fe7887a7908c41"),
	})
	if err != nil {
		t.Fatalf("failed to delete VPC: %v", err)
	}

	fmt.Println("VPC deleted")
}
