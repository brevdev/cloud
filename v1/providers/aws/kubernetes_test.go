package v1

import (
	"context"
	"fmt"
	"testing"

	v1 "github.com/brevdev/cloud/v1"
)

func TestCreateKubernetesCluster(t *testing.T) {
	awsClient, err := NewAWSClient("test", accessKeyID, secretAccessKey, "us-east-1")
	if err != nil {
		t.Fatalf("failed to create AWS client: %v", err)
	}

	cluster, err := awsClient.CreateCluster(context.Background(), v1.CreateClusterArgs{
		Name:  "cloud-sdk-test",
		RefID: "cloud-sdk-test",
		VPCID: v1.CloudProviderResourceID("vpc-0e7fe7887a7908c41"),
		SubnetIDs: []v1.CloudProviderResourceID{
			v1.CloudProviderResourceID("subnet-09a0a21834bedd85c"),
			v1.CloudProviderResourceID("subnet-01b65e00514f075df"),
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
		ID: cluster.ID,
	})
	if err != nil {
		t.Fatalf("failed to get cluster: %v", err)
	}

	fmt.Println(cluster)
}
