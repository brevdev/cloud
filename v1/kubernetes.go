package v1

import "context"

type Cluster struct {
	// The name of the cluster, displayed on clients.
	Name string

	// The unique ID used to associate with this cluster.
	RefID string

	// The ID assigned by the cloud provider to the cluster.
	CloudID string

	// The cloud provider that manages the cluster.
	Provider string

	// The cloud that hosts the cluster.
	Cloud string

	// The location of the cluster.
	Location string

	// The ID of the VPC that the cluster is associated with.
	VPCID string

	// The subnet IDs that the cluster's nodes are deployed into.
	SubnetIDs []string

	// The version of Kubernetes that the cluster is running.
	KubernetesVersion string

	// The status of the cluster.
	Status ClusterStatus

	// The node groups associated with the cluster.
	NodeGroups []NodeGroup
}

type NodeGroup struct {
	// The name of the node group, displayed on clients.
	Name string

	// The unique ID used to associate with this node group.
	RefID string
	// The minimum number of nodes in the node group.
	MinNodeCount int

	// The maximum number of nodes in the node group.
	MaxNodeCount int

	// The instance type of the nodes in the node group.
	InstanceType string
}

type ClusterStatus string

const (
	ClusterStatusPending   ClusterStatus = "pending"
	ClusterStatusAvailable ClusterStatus = "available"
)

type CreateClusterArgs struct {
	Name              string
	RefID             string
	VPCID             string
	SubnetIDs         []string
	KubernetesVersion string
	Location          string
	NodeGroups        []CreateNodeGroupArgs
}

type CreateNodeGroupArgs struct {
	Name         string
	RefID        string
	MinNodeCount int
	MaxNodeCount int
	InstanceType string
}

type GetClusterArgs struct {
	RefID    string
	CloudID  string
	Location string
}

type DeleteClusterArgs struct {
	Cluster *Cluster
}

type CloudMaintainKubernetes interface {
	CreateCluster(ctx context.Context, args CreateClusterArgs) (*Cluster, error)
	GetCluster(ctx context.Context, args GetClusterArgs) (*Cluster, error)
	DeleteCluster(ctx context.Context, args DeleteClusterArgs) error
}
