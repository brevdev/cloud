package v1

import (
	"context"
	"fmt"
	"time"
)

type Cluster struct {
	// The ID assigned by the cloud provider to the cluster.
	ID CloudProviderResourceID

	// The name of the cluster, displayed on clients.
	Name string

	// The unique ID used to associate with this cluster.
	RefID string

	// The cloud provider that manages the cluster.
	Provider string

	// The cloud that hosts the cluster.
	Cloud string

	// The location of the cluster.
	Location string

	// The ID of the VPC that the cluster is associated with.
	VPCID CloudProviderResourceID

	// The subnet IDs that the cluster's nodes are deployed into.
	SubnetIDs []CloudProviderResourceID

	// The version of Kubernetes that the cluster is running.
	KubernetesVersion string

	// The status of the cluster.
	Status ClusterStatus

	// The API endpoint of the cluster.
	APIEndpoint string

	// The CA certificate of the cluster, in base64.
	ClusterCACertificateBase64 string

	// The node groups associated with the cluster.
	NodeGroups []*NodeGroup
}

type NodeGroup struct {
	// The name of the node group, displayed on clients.
	Name string

	// The unique ID used to associate with this node group.
	RefID string

	// The ID assigned by the cloud provider to the node group.
	ID CloudProviderResourceID

	// The minimum number of nodes in the node group.
	MinNodeCount int

	// The maximum number of nodes in the node group.
	MaxNodeCount int

	// The instance type of the nodes in the node group.
	InstanceType string

	// The disk size of the nodes in the node group.
	DiskSizeGiB int
}

type ClusterStatus string

const (
	ClusterStatusUnknown   ClusterStatus = "unknown"
	ClusterStatusPending   ClusterStatus = "pending"
	ClusterStatusAvailable ClusterStatus = "available"
	ClusterStatusDeleting  ClusterStatus = "deleting"
)

type CloudProviderResourceID string

type CreateClusterArgs struct {
	Name              string
	RefID             string
	VPCID             CloudProviderResourceID
	SubnetIDs         []CloudProviderResourceID
	KubernetesVersion string
	Location          string
}

type PutUserArgs struct {
	ClusterID    CloudProviderResourceID
	Username     string
	RSAPEMBase64 string
}

type PutUserResponse struct {
	ClusterName                           string
	ClusterCertificateAuthorityDataBase64 string
	ClusterServerURL                      string
	Username                              string
	UserClientCertificateDataBase64       string
	UserClientKeyDataBase64               string
	KubeconfigBase64                      string
}

type GetClusterArgs struct {
	ID CloudProviderResourceID
}

type CreateNodeGroupArgs struct {
	ClusterID    CloudProviderResourceID
	Name         string
	RefID        string
	MinNodeCount int
	MaxNodeCount int
	InstanceType string
	DiskSizeGiB  int
}

type GetNodeGroupArgs struct {
	ID CloudProviderResourceID
}

type ModifyNodeGroupArgs struct {
	ID           CloudProviderResourceID
	MinNodeCount int
	MaxNodeCount int
}

type DeleteNodeGroupArgs struct {
	ID CloudProviderResourceID
}

type CreateNodeGroupResponse struct {
	ID           CloudProviderResourceID
	Name         string
	RefID        string
	MinNodeCount int
	MaxNodeCount int
	InstanceType string
	DiskSizeGiB  int
}

type DeleteClusterArgs struct {
	ID CloudProviderResourceID
}

type CloudMaintainKubernetes interface {
	CreateCluster(ctx context.Context, args CreateClusterArgs) (*Cluster, error)
	GetCluster(ctx context.Context, args GetClusterArgs) (*Cluster, error)
	PutUser(ctx context.Context, args PutUserArgs) (*PutUserResponse, error)
	CreateNodeGroup(ctx context.Context, args CreateNodeGroupArgs) (*CreateNodeGroupResponse, error)
	GetNodeGroup(ctx context.Context, args GetNodeGroupArgs) (*NodeGroup, error)
	ModifyNodeGroup(ctx context.Context, args ModifyNodeGroupArgs) error
	DeleteNodeGroup(ctx context.Context, args DeleteNodeGroupArgs) error
	DeleteCluster(ctx context.Context, args DeleteClusterArgs) error
}

func ValidateCreateKubernetesCluster(ctx context.Context, client CloudMaintainKubernetes, attrs CreateClusterArgs) (*Cluster, error) {
	cluster, err := client.CreateCluster(ctx, attrs)
	if err != nil {
		return nil, err
	}

	return cluster, nil
}

func ValidateGetKubernetesCluster(ctx context.Context, client CloudMaintainKubernetes, attrs GetClusterArgs) (*Cluster, error) {
	cluster, err := client.GetCluster(ctx, attrs)
	if err != nil {
		return nil, err
	}
	return cluster, nil
}

// WaitForKubernetesClusterPredicate waits for the Kubernetes cluster to satisfy the predicate function. If the predicate returns true, the loop breaks.
type WaitForKubernetesClusterPredicateOpts struct {
	Predicate func(cluster *Cluster) bool
	Timeout   time.Duration
	Interval  time.Duration
}

func WaitForKubernetesClusterPredicate(ctx context.Context, client CloudMaintainKubernetes, attrs GetClusterArgs, opts WaitForKubernetesClusterPredicateOpts) error {
	ctx, cancel := context.WithTimeout(ctx, opts.Timeout)
	defer cancel()

	ticker := time.NewTicker(opts.Interval)
	defer ticker.Stop()

	fmt.Printf("Entering WaitForKubernetesClusterPredicate, timeout: %s, interval: %s\n", opts.Timeout.String(), opts.Interval.String())
	for {
		cluster, err := client.GetCluster(ctx, attrs)
		if err != nil {
			return err
		}

		if opts.Predicate(cluster) {
			break
		}
		fmt.Printf("Waiting %s for cluster to satisfy predicate\n", opts.Interval.String())
		select {
		case <-ctx.Done():
			return fmt.Errorf("timeout waiting for cluster to satisfy predicate")
		case <-ticker.C:
			continue
		}
	}
	return nil
}

func ValidateGetKubernetesClusterCredentials(ctx context.Context, client CloudMaintainKubernetes, attrs GetClusterArgs) (*PutUserResponse, error) {
	putUserResponse, err := client.PutUser(ctx, PutUserArgs{
		ClusterID:    attrs.ID,
		Username:     "admin",
		RSAPEMBase64: "test",
	})
	if err != nil {
		return nil, err
	}
	return putUserResponse, nil
}

func ValidateCreateKubernetesNodeGroup(ctx context.Context, client CloudMaintainKubernetes, attrs CreateNodeGroupArgs) error {
	_, err := client.CreateNodeGroup(ctx, attrs)
	if err != nil {
		return err
	}
	return nil
}

func ValidateGetKubernetesNodeGroup(ctx context.Context, client CloudMaintainKubernetes, attrs GetNodeGroupArgs) error {
	_, err := client.GetNodeGroup(ctx, attrs)
	if err != nil {
		return err
	}
	return nil
}

func ValidateModifyKubernetesNodeGroup(ctx context.Context, client CloudMaintainKubernetes, attrs ModifyNodeGroupArgs) error {
	err := client.ModifyNodeGroup(ctx, attrs)
	if err != nil {
		return err
	}
	return nil
}

func ValidateDeleteKubernetesNodeGroup(ctx context.Context, client CloudMaintainKubernetes, attrs DeleteNodeGroupArgs) error {
	err := client.DeleteNodeGroup(ctx, attrs)
	if err != nil {
		return err
	}
	return nil
}

func ValidateDeleteKubernetesCluster(ctx context.Context, client CloudMaintainKubernetes, attrs DeleteClusterArgs) error {
	err := client.DeleteCluster(ctx, attrs)
	if err != nil {
		return err
	}
	return nil
}
