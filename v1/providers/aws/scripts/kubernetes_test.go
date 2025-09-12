//go:build scripts
// +build scripts

package scripts

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/brevdev/cloud/internal/validation"
	v1 "github.com/brevdev/cloud/v1"
	aws "github.com/brevdev/cloud/v1/providers/aws"
)

const (
	accessKeyID     = "test"
	secretAccessKey = "test"
)

func TestCreateKubernetesCluster(t *testing.T) {
	awsClient, err := aws.NewAWSClient("test", accessKeyID, secretAccessKey, "us-east-1", aws.WithLogger(&validation.ValidationLogger{}))
	if err != nil {
		t.Fatalf("failed to create AWS client: %v", err)
	}

	cluster, err := awsClient.CreateCluster(context.Background(), v1.CreateClusterArgs{
		Name:  "cloud-sdk-test",
		RefID: "cloud-sdk-test",
		VPCID: v1.CloudProviderResourceID("vpc-09035a20d5b393eff"),
		SubnetIDs: []v1.CloudProviderResourceID{
			v1.CloudProviderResourceID("subnet-0ba8c98b636237a2d"),
			v1.CloudProviderResourceID("subnet-07fadc6ba1992285b"),
		},
		KubernetesVersion: "1.31",
		Tags: v1.Tags{
			"test": "test",
		},
	})
	if err != nil {
		t.Fatalf("failed to create cluster: %v", err)
	}

	cluster, err = awsClient.GetCluster(context.Background(), v1.GetClusterArgs{
		ID: cluster.GetID(),
	})
	if err != nil {
		t.Fatalf("failed to get cluster: %v", err)
	}

	fmt.Println(cluster)
}

func TestGetKubernetesCluster(t *testing.T) {
	awsClient, err := aws.NewAWSClient("test", accessKeyID, secretAccessKey, "us-east-1", aws.WithLogger(&validation.ValidationLogger{}))
	if err != nil {
		t.Fatalf("failed to create AWS client: %v", err)
	}

	cluster, err := awsClient.GetCluster(context.Background(), v1.GetClusterArgs{
		ID: v1.CloudProviderResourceID("cloud-sdk-test2"),
	})
	if err != nil {
		t.Fatalf("failed to get cluster: %v", err)
	}

	fmt.Println(cluster)
}

func TestAddNodeGroupToKubernetesCluster(t *testing.T) {
	awsClient, err := aws.NewAWSClient("test", accessKeyID, secretAccessKey, "us-east-1", aws.WithLogger(&validation.ValidationLogger{}))
	if err != nil {
		t.Fatalf("failed to create AWS client: %v", err)
	}

	nodeGroup, err := awsClient.CreateNodeGroup(context.Background(), v1.CreateNodeGroupArgs{
		ClusterID:    v1.CloudProviderResourceID("cloud-sdk-test2"),
		Name:         "cloud-sdk-test-node-group",
		RefID:        "cloud-sdk-test-node-group",
		MinNodeCount: 1,
		MaxNodeCount: 1,
		InstanceType: "t3.medium",
		DiskSizeGiB:  20,
		Tags: v1.Tags{
			"test": "test",
		},
	})
	if err != nil {
		t.Fatalf("failed to create node group: %v", err)
	}

	fmt.Println(nodeGroup)
}

func TestGetNodeGroup(t *testing.T) {
	awsClient, err := aws.NewAWSClient("test", accessKeyID, secretAccessKey, "us-east-1", aws.WithLogger(&validation.ValidationLogger{}))
	if err != nil {
		t.Fatalf("failed to create AWS client: %v", err)
	}

	nodeGroup, err := awsClient.GetNodeGroup(context.Background(), v1.GetNodeGroupArgs{
		ID:        v1.CloudProviderResourceID("cloud-sdk-test-node-group"),
		ClusterID: v1.CloudProviderResourceID("cloud-sdk-test"),
	})
	if err != nil {
		t.Fatalf("failed to get node group: %v", err)
	}

	fmt.Println(nodeGroup)
}

func TestModifyNodeGroup(t *testing.T) {
	awsClient, err := aws.NewAWSClient("test", accessKeyID, secretAccessKey, "us-east-1", aws.WithLogger(&validation.ValidationLogger{}))
	if err != nil {
		t.Fatalf("failed to create AWS client: %v", err)
	}

	err = awsClient.ModifyNodeGroup(context.Background(), v1.ModifyNodeGroupArgs{
		ClusterID:    v1.CloudProviderResourceID("cloud-sdk-test"),
		ID:           v1.CloudProviderResourceID("cloud-sdk-test-node-group"),
		MinNodeCount: 2,
		MaxNodeCount: 2,
	})
	if err != nil {
		t.Fatalf("failed to modify node group: %v", err)
	}
}

func TestDeleteNodeGroup(t *testing.T) {
	awsClient, err := aws.NewAWSClient("test", accessKeyID, secretAccessKey, "us-east-1", aws.WithLogger(&validation.ValidationLogger{}))
	if err != nil {
		t.Fatalf("failed to create AWS client: %v", err)
	}

	err = awsClient.DeleteNodeGroup(context.Background(), v1.DeleteNodeGroupArgs{
		ClusterID: v1.CloudProviderResourceID("testcloudsdk-20251103191744"),
		ID:        v1.CloudProviderResourceID("testcloudsdk-20251103191744"),
	})
	if err != nil {
		t.Fatalf("failed to delete node group: %v", err)
	}
}

func TestPutUser(t *testing.T) {
	testUserPrivateKeyPEMBase64 := os.Getenv("TEST_USER_PRIVATE_KEY_PEM_BASE64")

	awsClient, err := aws.NewAWSClient("test", accessKeyID, secretAccessKey, "us-east-1", aws.WithLogger(&validation.ValidationLogger{}))
	if err != nil {
		t.Fatalf("failed to create AWS client: %v", err)
	}

	config, err := awsClient.SetClusterUser(context.Background(), v1.SetClusterUserArgs{
		ClusterID:    v1.CloudProviderResourceID("cloud-sdk-test2"),
		Username:     "test-user",
		Role:         "cluster-admin",
		RSAPEMBase64: testUserPrivateKeyPEMBase64,
	})
	if err != nil {
		t.Fatalf("failed to put user: %v", err)
	}

	fmt.Println(config)
}

func TestDeleteKubernetesCluster(t *testing.T) {
	awsClient, err := aws.NewAWSClient("test", accessKeyID, secretAccessKey, "us-east-1", aws.WithLogger(&validation.ValidationLogger{}))
	if err != nil {
		t.Fatalf("failed to create AWS client: %v", err)
	}

	err = awsClient.DeleteCluster(context.Background(), v1.DeleteClusterArgs{
		ID: v1.CloudProviderResourceID("testcloudsdk-20251103200615"),
	})
	if err != nil {
		t.Fatalf("failed to delete cluster: %v", err)
	}
}
