package v1

import (
	"context"
	"fmt"
	"os"
	"testing"

	v1 "github.com/brevdev/cloud/v1"
)

func TestCreateVPC(t *testing.T) {
	// get env var
	accessKeyID := os.Getenv("AWS_ACCESS_KEY_ID")
	secretAccessKey := os.Getenv("AWS_SECRET_ACCESS_KEY")

	awsClient, err := NewAWSClient("test", accessKeyID, secretAccessKey)
	if err != nil {
		t.Fatalf("failed to create AWS client: %v", err)
	}

	vpc, err := awsClient.CreateVPC(context.Background(), v1.CreateVPCArgs{
		Name:      "cloud-sdk-test",
		RefID:     "cloud-sdk-test",
		Location:  "us-east-1",
		CidrBlock: "10.0.0.0/16",
	})
	if err != nil {
		t.Fatalf("failed to create VPC: %v", err)
	}

	vpc, err = awsClient.GetVPC(context.Background(), v1.GetVPCArgs{
		CloudID:  vpc.CloudID,
		Location: vpc.Location,
	})
	if err != nil {
		t.Fatalf("failed to get VPC: %v", err)
	}
	fmt.Println("VPC retrieved")

	err = awsClient.DeleteVPC(context.Background(), v1.DeleteVPCArgs{
		VPC: &v1.VPC{
			CloudID:  vpc.CloudID,
			Location: "us-east-1",
		},
	})
	if err != nil {
		t.Fatalf("failed to delete VPC: %v", err)
	}

	fmt.Println("VPC deleted")
}
