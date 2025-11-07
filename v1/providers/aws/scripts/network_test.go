//go:build scripts
// +build scripts

package scripts

import (
	"context"
	"fmt"
	"testing"

	"github.com/brevdev/cloud/internal/validation"
	v1 "github.com/brevdev/cloud/v1"
	aws "github.com/brevdev/cloud/v1/providers/aws"
)

func TestCreateVPC(t *testing.T) {
	awsClient, err := aws.NewAWSClient("test", accessKeyID, secretAccessKey, "us-east-1", aws.WithLogger(&validation.ValidationLogger{}))
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
			{CidrBlock: "10.0.64.0/19", Type: v1.SubnetTypePublic},
			{CidrBlock: "10.0.96.0/19", Type: v1.SubnetTypePrivate},
		},
	})
	if err != nil {
		t.Fatalf("failed to create VPC: %v", err)
	}

	vpc, err = awsClient.GetVPC(context.Background(), v1.GetVPCArgs{
		ID: vpc.GetID(),
	})
	if err != nil {
		t.Fatalf("failed to get VPC: %v", err)
	}
}

func TestDeleteVPC(t *testing.T) {
	awsClient, err := aws.NewAWSClient("test", accessKeyID, secretAccessKey, "us-east-1", aws.WithLogger(&validation.ValidationLogger{}))
	if err != nil {
		t.Fatalf("failed to create AWS client: %v", err)
	}

	err = awsClient.DeleteVPC(context.Background(), v1.DeleteVPCArgs{
		ID: v1.CloudProviderResourceID("vpc-01e30509323927f79"),
	})
	if err != nil {
		t.Fatalf("failed to delete VPC: %v", err)
	}

	fmt.Println("VPC deleted")
}
