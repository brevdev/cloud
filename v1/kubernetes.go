package v1

import (
	"context"
	"fmt"

	"github.com/brevdev/cloud/internal/errors"
)

// Cluster represents the complete specification of a Brev Kubernetes cluster.
type Cluster struct {
	// The ID assigned by the cloud provider to the cluster.
	id CloudProviderResourceID

	// The name of the cluster, displayed on clients.
	name string

	// The unique ID used to associate with this cluster.
	refID string

	// The cloud provider that manages the cluster.
	provider string

	// The cloud that hosts the cluster.
	cloud string

	// The location of the cluster.
	location string

	// The ID of the VPC that the cluster is associated with.
	vpcID CloudProviderResourceID

	// The subnet IDs that the cluster's nodes are deployed into.
	subnetIDs []CloudProviderResourceID

	// The version of Kubernetes that the cluster is running.
	kubernetesVersion string

	// The status of the cluster.
	status ClusterStatus

	// The API endpoint of the cluster.
	apiEndpoint string

	// The CA certificate of the cluster, in base64.
	clusterCACertificateBase64 string

	// The node groups associated with the cluster.
	nodeGroups []*NodeGroup

	// The tags associated with the cluster.
	tags Tags
}

type ClusterStatus string

const (
	ClusterStatusUnknown   ClusterStatus = "unknown"
	ClusterStatusPending   ClusterStatus = "pending"
	ClusterStatusAvailable ClusterStatus = "available"
	ClusterStatusDeleting  ClusterStatus = "deleting"
	ClusterStatusFailed    ClusterStatus = "failed"
)

func (c *Cluster) GetID() CloudProviderResourceID {
	return c.id
}

func (c *Cluster) GetName() string {
	return c.name
}

func (c *Cluster) GetRefID() string {
	return c.refID
}

func (c *Cluster) GetProvider() string {
	return c.provider
}

func (c *Cluster) GetCloud() string {
	return c.cloud
}

func (c *Cluster) GetLocation() string {
	return c.location
}

func (c *Cluster) GetVPCID() CloudProviderResourceID {
	return c.vpcID
}

func (c *Cluster) GetSubnetIDs() []CloudProviderResourceID {
	return c.subnetIDs
}

func (c *Cluster) GetKubernetesVersion() string {
	return c.kubernetesVersion
}

func (c *Cluster) GetStatus() ClusterStatus {
	return c.status
}

func (c *Cluster) GetAPIEndpoint() string {
	return c.apiEndpoint
}

func (c *Cluster) GetClusterCACertificateBase64() string {
	return c.clusterCACertificateBase64
}

func (c *Cluster) GetNodeGroups() []*NodeGroup {
	return c.nodeGroups
}

func (c *Cluster) GetTags() Tags {
	return c.tags
}

// ClusterSettings represents the settings for a Kubernetes cluster. This is the input to the NewCluster function.
type ClusterSettings struct {
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

	// The tags associated with the cluster.
	Tags Tags
}

func (s *ClusterSettings) setDefaults() {
}

func (s *ClusterSettings) validate() error {
	var errs []error
	if s.RefID == "" {
		errs = append(errs, fmt.Errorf("refID is required"))
	}
	if s.Name == "" {
		errs = append(errs, fmt.Errorf("name is required"))
	}
	if s.Status == "" {
		errs = append(errs, fmt.Errorf("status is required"))
	}
	return errors.WrapAndTrace(errors.Join(errs...))
}

// NewCluster creates a new Cluster from the provided settings.
func NewCluster(settings ClusterSettings) (*Cluster, error) {
	settings.setDefaults()
	err := settings.validate()
	if err != nil {
		return nil, errors.WrapAndTrace(err)
	}
	return &Cluster{
		id:                         settings.ID,
		name:                       settings.Name,
		refID:                      settings.RefID,
		provider:                   settings.Provider,
		cloud:                      settings.Cloud,
		location:                   settings.Location,
		vpcID:                      settings.VPCID,
		subnetIDs:                  settings.SubnetIDs,
		kubernetesVersion:          settings.KubernetesVersion,
		status:                     settings.Status,
		apiEndpoint:                settings.APIEndpoint,
		clusterCACertificateBase64: settings.ClusterCACertificateBase64,
		nodeGroups:                 settings.NodeGroups,
		tags:                       settings.Tags,
	}, nil
}

// NodeGroup represents the complete specification of a Brev Kubernetes node group.
type NodeGroup struct {
	// The name of the node group, displayed on clients.
	name string

	// The unique ID used to associate with this node group.
	refID string

	// The ID assigned by the cloud provider to the node group.
	id CloudProviderResourceID

	// The minimum number of nodes in the node group.
	minNodeCount int

	// The maximum number of nodes in the node group.
	maxNodeCount int

	// The instance type of the nodes in the node group.
	instanceType string

	// The disk size of the nodes in the node group.
	diskSizeGiB int

	// The status of the node group.
	status NodeGroupStatus

	// The tags associated with the node group.
	tags Tags
}

type NodeGroupStatus string

const (
	NodeGroupStatusUnknown   NodeGroupStatus = "unknown"
	NodeGroupStatusPending   NodeGroupStatus = "pending"
	NodeGroupStatusAvailable NodeGroupStatus = "available"
	NodeGroupStatusDeleting  NodeGroupStatus = "deleting"
	NodeGroupStatusFailed    NodeGroupStatus = "failed"
)

func (n *NodeGroup) GetName() string {
	return n.name
}

func (n *NodeGroup) GetRefID() string {
	return n.refID
}

func (n *NodeGroup) GetID() CloudProviderResourceID {
	return n.id
}

func (n *NodeGroup) GetMinNodeCount() int {
	return n.minNodeCount
}

func (n *NodeGroup) GetMaxNodeCount() int {
	return n.maxNodeCount
}

func (n *NodeGroup) GetInstanceType() string {
	return n.instanceType
}

func (n *NodeGroup) GetDiskSizeGiB() int {
	return n.diskSizeGiB
}

func (n *NodeGroup) GetStatus() NodeGroupStatus {
	return n.status
}

func (n *NodeGroup) GetTags() Tags {
	return n.tags
}

// NodeGroupSettings represents the settings for a Kubernetes node group. This is the input to the NewNodeGroup function.
type NodeGroupSettings struct {
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

	// The status of the node group.
	Status NodeGroupStatus

	// The tags associated with the node group.
	Tags Tags
}

func (s *NodeGroupSettings) setDefaults() {
}

func (s *NodeGroupSettings) validate() error {
	var errs []error
	if s.RefID == "" {
		errs = append(errs, fmt.Errorf("refID is required"))
	}
	if s.Name == "" {
		errs = append(errs, fmt.Errorf("name is required"))
	}
	if s.Status == "" {
		errs = append(errs, fmt.Errorf("status is required"))
	}
	return errors.WrapAndTrace(errors.Join(errs...))
}

// NewNodeGroup creates a new NodeGroup from the provided settings.
func NewNodeGroup(settings NodeGroupSettings) (*NodeGroup, error) {
	settings.setDefaults()
	err := settings.validate()
	if err != nil {
		return nil, errors.WrapAndTrace(err)
	}
	return &NodeGroup{
		name:         settings.Name,
		refID:        settings.RefID,
		id:           settings.ID,
		minNodeCount: settings.MinNodeCount,
		maxNodeCount: settings.MaxNodeCount,
		instanceType: settings.InstanceType,
		diskSizeGiB:  settings.DiskSizeGiB,
		status:       settings.Status,
		tags:         settings.Tags,
	}, nil
}

// ClusterUser represents the complete specification of a Brev Kubernetes cluster user.
type ClusterUser struct {
	// The name of the cluster that the user is associated with.
	clusterName string

	// The CA certificate of the cluster, in base64.
	clusterCertificateAuthorityDataBase64 string

	// The API endpoint of the cluster.
	clusterServerURL string

	// The username of the user.
	username string

	// The client certificate of the user, in base64.
	userClientCertificateDataBase64 string

	// The client key of the user, in base64.
	userClientKeyDataBase64 string

	// The kubeconfig of the user, in base64.
	kubeconfigBase64 string
}

func (c *ClusterUser) GetClusterName() string {
	return c.clusterName
}

func (c *ClusterUser) GetClusterCertificateAuthorityDataBase64() string {
	return c.clusterCertificateAuthorityDataBase64
}

func (c *ClusterUser) GetClusterServerURL() string {
	return c.clusterServerURL
}

func (c *ClusterUser) GetUsername() string {
	return c.username
}

func (c *ClusterUser) GetUserClientCertificateDataBase64() string {
	return c.userClientCertificateDataBase64
}

func (c *ClusterUser) GetUserClientKeyDataBase64() string {
	return c.userClientKeyDataBase64
}

func (c *ClusterUser) GetKubeconfigBase64() string {
	return c.kubeconfigBase64
}

// ClusterUserSettings represents the settings for a Kubernetes cluster user. This is the input to the NewClusterUser function.
type ClusterUserSettings struct {
	// The name of the cluster that the user is associated with.
	ClusterName string

	// The CA certificate of the cluster, in base64.
	ClusterCertificateAuthorityDataBase64 string

	// The API endpoint of the cluster.
	ClusterServerURL string

	// The username of the user.
	Username string

	// The client certificate of the user, in base64.
	UserClientCertificateDataBase64 string

	// The client key of the user, in base64.
	UserClientKeyDataBase64 string

	// The kubeconfig of the user, in base64.
	KubeconfigBase64 string
}

func (s *ClusterUserSettings) setDefaults() {
}

func (s *ClusterUserSettings) validate() error {
	var errs []error
	if s.ClusterName == "" {
		errs = append(errs, fmt.Errorf("clusterName is required"))
	}
	if s.ClusterCertificateAuthorityDataBase64 == "" {
		errs = append(errs, fmt.Errorf("clusterCertificateAuthorityDataBase64 is required"))
	}
	if s.ClusterServerURL == "" {
		errs = append(errs, fmt.Errorf("clusterServerURL is required"))
	}
	if s.Username == "" {
		errs = append(errs, fmt.Errorf("username is required"))
	}
	if s.UserClientCertificateDataBase64 == "" {
		errs = append(errs, fmt.Errorf("userClientCertificateDataBase64 is required"))
	}
	if s.UserClientKeyDataBase64 == "" {
		errs = append(errs, fmt.Errorf("userClientKeyDataBase64 is required"))
	}
	if s.KubeconfigBase64 == "" {
		errs = append(errs, fmt.Errorf("kubeconfigBase64 is required"))
	}
	return errors.WrapAndTrace(errors.Join(errs...))
}

// NewClusterUser creates a new ClusterUser from the provided settings.
func NewClusterUser(settings ClusterUserSettings) (*ClusterUser, error) {
	settings.setDefaults()
	err := settings.validate()
	if err != nil {
		return nil, errors.WrapAndTrace(err)
	}
	return &ClusterUser{
		clusterName:                           settings.ClusterName,
		clusterCertificateAuthorityDataBase64: settings.ClusterCertificateAuthorityDataBase64,
		clusterServerURL:                      settings.ClusterServerURL,
		username:                              settings.Username,
		userClientCertificateDataBase64:       settings.UserClientCertificateDataBase64,
		userClientKeyDataBase64:               settings.UserClientKeyDataBase64,
		kubeconfigBase64:                      settings.KubeconfigBase64,
	}, nil
}

type CloudMaintainKubernetes interface {
	// Create a new Kubernetes cluster.
	CreateCluster(ctx context.Context, args CreateClusterArgs) (*Cluster, error)

	// Get a Kubernetes cluster identified by the provided args.
	GetCluster(ctx context.Context, args GetClusterArgs) (*Cluster, error)

	// Idempotently set a user into a Kubernetes cluster.
	SetClusterUser(ctx context.Context, args SetClusterUserArgs) (*ClusterUser, error)

	// Create a new Kubernetes node group.
	CreateNodeGroup(ctx context.Context, args CreateNodeGroupArgs) (*NodeGroup, error)

	// Get a Kubernetes node group identified by the provided args.
	GetNodeGroup(ctx context.Context, args GetNodeGroupArgs) (*NodeGroup, error)

	// Modify a Kubernetes node group.
	ModifyNodeGroup(ctx context.Context, args ModifyNodeGroupArgs) error

	// Delete a Kubernetes node group identified by the provided args.
	DeleteNodeGroup(ctx context.Context, args DeleteNodeGroupArgs) error

	// Delete a Kubernetes cluster identified by the provided args.
	DeleteCluster(ctx context.Context, args DeleteClusterArgs) error
}

type CreateClusterArgs struct {
	Name              string
	RefID             string
	VPCID             CloudProviderResourceID
	SubnetIDs         []CloudProviderResourceID
	KubernetesVersion string
	Tags              Tags
}

type SetClusterUserArgs struct {
	ClusterID    CloudProviderResourceID
	Username     string
	RSAPEMBase64 string
	Role         string
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
	Tags         Tags
}

type GetNodeGroupArgs struct {
	ClusterID CloudProviderResourceID
	ID        CloudProviderResourceID
}

type ModifyNodeGroupArgs struct {
	ClusterID    CloudProviderResourceID
	ID           CloudProviderResourceID
	MinNodeCount int
	MaxNodeCount int
}

type DeleteNodeGroupArgs struct {
	ClusterID CloudProviderResourceID
	ID        CloudProviderResourceID
}

type DeleteClusterArgs struct {
	ID CloudProviderResourceID
}
