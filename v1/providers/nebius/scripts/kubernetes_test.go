//go:build scripts
// +build scripts

package scripts

import (
	"context"
	"fmt"
	"os"
	"testing"

	v1 "github.com/brevdev/cloud/v1"
	nebius "github.com/brevdev/cloud/v1/providers/nebius"
)

var platformPresetMap = map[string][]string{
	"cpu-d3":       {"4vcpu-16gb", "8vcpu-32gb", "16vcpu-64gb", "32vcpu-128gb", "48vcpu-192gb", "64vcpu-256gb", "96vcpu-384gb", "128vcpu-512gb"},
	"cpu-e2":       {"2vcpu-8gb", "4vcpu-16gb", "8vcpu-32gb", "16vcpu-64gb", "32vcpu-128gb", "48vcpu-192gb", "64vcpu-256gb", "80vcpu-320gb"},
	"gpu-h200-sxm": {"1gpu-16vcpu-200gb", "8gpu-128vcpu-1600gb"},
	"gpu-h100-sxm": {"1gpu-16vcpu-200gb", "8gpu-128vcpu-1600gb"},
	"gpu-l40s-a":   {"1gpu-8vcpu-32gb", "1gpu-16vcpu-64gb", "1gpu-24vcpu-96gb", "1gpu-32vcpu-128gb", "1gpu-40vcpu-160gb"},
	"gpu-l40s-d":   {"1gpu-16vcpu-96gb", "1gpu-32vcpu-192gb", "1gpu-48vcpu-288gb", "2gpu-64vcpu-384gb", "2gpu-96vcpu-576gb", "4gpu-128vcpu-768gb", "4gpu-192vcpu-1152gb"},
}

func Test_CreateVPCAndCluster(t *testing.T) {
	privateKeyPEMBase64 := os.Getenv("NEBIUS_PRIVATE_KEY_PEM_BASE64")
	publicKeyID := os.Getenv("NEBIUS_PUBLIC_KEY_ID")
	serviceAccountID := os.Getenv("NEBIUS_SERVICE_ACCOUNT_ID")

	projectID := "project-e00nrhefpr009ynkkzcgba" // eu-north1

	nebiusClient, err := nebius.NewNebiusClient(context.Background(), "test", publicKeyID, privateKeyPEMBase64, serviceAccountID, projectID)
	if err != nil {
		t.Fatalf("failed to create Nebius client: %v", err)
	}

	vpc, err := nebiusClient.CreateVPC(context.Background(), v1.CreateVPCArgs{
		Name:      "cloud-sdk-test",
		RefID:     "cloud-sdk-test",
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
		VPCID:             v1.CloudProviderResourceID(vpc.GetID()),
		SubnetIDs:         []v1.CloudProviderResourceID{v1.CloudProviderResourceID(vpc.GetSubnets()[0].GetID())},
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

	nebiusClient, err := nebius.NewNebiusClient(context.Background(), "test", publicKeyID, privateKeyPEMBase64, serviceAccountID, projectID)
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

	nebiusClient, err := nebius.NewNebiusClient(context.Background(), "test", publicKeyID, privateKeyPEMBase64, serviceAccountID, projectID)
	if err != nil {
		t.Fatalf("failed to create Nebius client: %v", err)
	}

	putUserResponse, err := nebiusClient.SetClusterUser(context.Background(), v1.SetClusterUserArgs{
		Username:     "test-user",
		Role:         "cluster-admin",
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

	nebiusClient, err := nebius.NewNebiusClient(context.Background(), "test", publicKeyID, privateKeyPEMBase64, serviceAccountID, projectID)
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
	projectID := os.Getenv("NEBIUS_PROJECT_ID")

	nebiusClient, err := nebius.NewNebiusClient(context.Background(), "test", publicKeyID, privateKeyPEMBase64, serviceAccountID, projectID)
	if err != nil {
		t.Fatalf("failed to create Nebius client: %v", err)
	}

	err = nebiusClient.DeleteCluster(context.Background(), v1.DeleteClusterArgs{
		ID: v1.CloudProviderResourceID("mk8scluster-u00vgffpfgh3ze60vr"),
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

	nebiusClient, err := nebius.NewNebiusClient(context.Background(), "test", publicKeyID, privateKeyPEMBase64, serviceAccountID, projectID)
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
