package v1

import (
	"context"
	"fmt"
	"os"
	"testing"

	v1 "github.com/brevdev/cloud/v1"
)

func Test_CreateVPCAndCluster(t *testing.T) {
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
		CidrBlock: "172.16.0.0/16",
		Subnets: []v1.CreateSubnetArgs{
			{CidrBlock: "172.16.0.0/19", Type: v1.SubnetTypePublic}, // note, /24 (IP count: 256) is too small for a subnet
			{CidrBlock: "172.16.32.0/19", Type: v1.SubnetTypePrivate},
		},
	})
	if err != nil {
		t.Fatalf("failed to create VPC: %v", err)
	}

	cluster, err := nebiusClient.CreateCluster(context.Background(), v1.CreateClusterArgs{
		Name:              "cloud-sdk-test",
		RefID:             "cloud-sdk-test",
		VPCID:             string(vpc.ID),
		SubnetIDs:         []string{string(vpc.Subnets[0].ID)},
		KubernetesVersion: "1.31",
	})
	if err != nil {
		t.Fatalf("failed to create cluster: %v", err)
	}

	fmt.Println(cluster)
}

func Test_GetCluster(t *testing.T) {
	privateKeyPEMBase64 := os.Getenv("NEBIUS_PRIVATE_KEY_PEM_BASE64")
	publicKeyID := os.Getenv("NEBIUS_PUBLIC_KEY_ID")
	serviceAccountID := os.Getenv("NEBIUS_SERVICE_ACCOUNT_ID")

	projectID := "project-e00nrhefpr009ynkkzcgba" // eu-north1

	nebiusClient, err := NewNebiusClient(context.Background(), "test", publicKeyID, privateKeyPEMBase64, serviceAccountID, projectID)
	if err != nil {
		t.Fatalf("failed to create Nebius client: %v", err)
	}

	cluster, err := nebiusClient.GetCluster(context.Background(), v1.GetClusterArgs{
		ID: v1.CloudProviderResourceID("cloud-sdk-test"),
	})
	if err != nil {
		t.Fatalf("failed to get cluster: %v", err)
	}

	fmt.Println(cluster)
}

func Test_PutUser(t *testing.T) {
	testUserPrivateKeyPEMBase64 := os.Getenv("TEST_USER_PRIVATE_KEY_PEM_BASE64")
	privateKeyPEMBase64 := os.Getenv("NEBIUS_PRIVATE_KEY_PEM_BASE64")
	publicKeyID := os.Getenv("NEBIUS_PUBLIC_KEY_ID")
	serviceAccountID := os.Getenv("NEBIUS_SERVICE_ACCOUNT_ID")

	projectID := "project-e00nrhefpr009ynkkzcgba" // eu-north1

	nebiusClient, err := NewNebiusClient(context.Background(), "test", publicKeyID, privateKeyPEMBase64, serviceAccountID, projectID)
	if err != nil {
		t.Fatalf("failed to create Nebius client: %v", err)
	}

	putUserResponse, err := nebiusClient.PutUser(context.Background(), v1.PutUserArgs{
		Username:     "test-user",
		ClusterID:    v1.CloudProviderResourceID("cloud-sdk-test"),
		RSAPEMBase64: testUserPrivateKeyPEMBase64,
	})
	if err != nil {
		t.Fatalf("failed to put user: %v", err)
	}

	fmt.Println(putUserResponse)
}

func Test_CreateNodeGroup(t *testing.T) {
	privateKeyPEMBase64 := os.Getenv("NEBIUS_PRIVATE_KEY_PEM_BASE64")
	publicKeyID := os.Getenv("NEBIUS_PUBLIC_KEY_ID")
	serviceAccountID := os.Getenv("NEBIUS_SERVICE_ACCOUNT_ID")

	projectID := "project-e00nrhefpr009ynkkzcgba" // eu-north1

	nebiusClient, err := NewNebiusClient(context.Background(), "test", publicKeyID, privateKeyPEMBase64, serviceAccountID, projectID)
	if err != nil {
		t.Fatalf("failed to create Nebius client: %v", err)
	}

	platform := "gpu-h100-sxm"
	preset := platformPresetMap[platform][0]

	createNodeGroupResponse, err := nebiusClient.CreateNodeGroup(context.Background(), v1.CreateNodeGroupArgs{
		ClusterID:    v1.CloudProviderResourceID("cloud-sdk-test"),
		Name:         "test-node-group3",
		RefID:        "test-node-group3",
		MinNodeCount: 1,
		MaxNodeCount: 2,
		InstanceType: fmt.Sprintf("%s.%s", platform, preset),
		DiskSizeGiB:  96,
	})
	if err != nil {
		t.Fatalf("failed to create node group: %v", err)
	}

	fmt.Println(createNodeGroupResponse)
}

func Test_DeleteCluster(t *testing.T) {
	privateKeyPEMBase64 := os.Getenv("NEBIUS_PRIVATE_KEY_PEM_BASE64")
	publicKeyID := os.Getenv("NEBIUS_PUBLIC_KEY_ID")
	serviceAccountID := os.Getenv("NEBIUS_SERVICE_ACCOUNT_ID")

	projectID := "project-e00nrhefpr009ynkkzcgba" // eu-north1

	nebiusClient, err := NewNebiusClient(context.Background(), "test", publicKeyID, privateKeyPEMBase64, serviceAccountID, projectID)
	if err != nil {
		t.Fatalf("failed to create Nebius client: %v", err)
	}

	err = nebiusClient.DeleteCluster(context.Background(), v1.DeleteClusterArgs{
		ID: v1.CloudProviderResourceID("cloud-sdk-test"),
	})
	if err != nil {
		t.Fatalf("failed to delete cluster: %v", err)
	}

	fmt.Println("Cluster deleted")
}

func Test_DeleteVPC(t *testing.T) {
	privateKeyPEMBase64 := os.Getenv("NEBIUS_PRIVATE_KEY_PEM_BASE64")
	publicKeyID := os.Getenv("NEBIUS_PUBLIC_KEY_ID")
	serviceAccountID := os.Getenv("NEBIUS_SERVICE_ACCOUNT_ID")

	projectID := "project-e00nrhefpr009ynkkzcgba" // eu-north1

	nebiusClient, err := NewNebiusClient(context.Background(), "test", publicKeyID, privateKeyPEMBase64, serviceAccountID, projectID)
	if err != nil {
		t.Fatalf("failed to create Nebius client: %v", err)
	}

	err = nebiusClient.DeleteVPC(context.Background(), v1.DeleteVPCArgs{
		ID: v1.CloudProviderResourceID("cloud-sdk-test"),
	})
	if err != nil {
		t.Fatalf("failed to delete VPC: %v", err)
	}

	fmt.Println("VPC deleted")
}
