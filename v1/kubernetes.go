package v1

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/brevdev/cloud/internal/errors"
)

var (
	ErrRefIDRequired                                            = errors.New("refID is required")
	ErrNameRequired                                             = errors.New("name is required")
	ErrNodeGroupInvalidStatus                                   = errors.New("invalid node group status")
	ErrClusterInvalidStatus                                     = errors.New("invalid cluster status")
	ErrClusterUserClusterNameRequired                           = errors.New("clusterName is required")
	ErrClusterUserClusterCertificateAuthorityDataBase64Required = errors.New("clusterCertificateAuthorityDataBase64 is required")
	ErrClusterUserClusterServerURLRequired                      = errors.New("clusterServerURL is required")
	ErrClusterUserUsernameRequired                              = errors.New("username is required")
	ErrClusterUserUserClientCertificateDataBase64Required       = errors.New("userClientCertificateDataBase64 is required")
	ErrClusterUserUserClientKeyDataBase64Required               = errors.New("userClientKeyDataBase64 is required")
	ErrClusterUserKubeconfigBase64Required                      = errors.New("kubeconfigBase64 is required")
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

// clusterJSON is the JSON representation of a Cluster. This struct is maintained separately from the core Cluster
// struct to allow for unexported fields to be used in the MarshalJSON and UnmarshalJSON methods.
type clusterJSON struct {
	ID                         string            `json:"id"`
	Name                       string            `json:"name"`
	RefID                      string            `json:"refID"`
	Provider                   string            `json:"provider"`
	Cloud                      string            `json:"cloud"`
	Location                   string            `json:"location"`
	VPCID                      string            `json:"vpcID"`
	SubnetIDs                  []string          `json:"subnetIDs"`
	KubernetesVersion          string            `json:"kubernetesVersion"`
	Status                     string            `json:"status"`
	APIEndpoint                string            `json:"apiEndpoint"`
	ClusterCACertificateBase64 string            `json:"clusterCACertificateBase64"`
	NodeGroups                 []*NodeGroup      `json:"nodeGroups"`
	Tags                       map[string]string `json:"tags"`
}

// MarshalJSON implements the json.Marshaler interface
func (c *Cluster) MarshalJSON() ([]byte, error) {
	subnetIDs := make([]string, len(c.subnetIDs))
	for i, subnetID := range c.subnetIDs {
		subnetIDs[i] = string(subnetID)
	}

	return json.Marshal(clusterJSON{
		ID:                         string(c.id),
		Name:                       c.name,
		RefID:                      c.refID,
		Provider:                   c.provider,
		Cloud:                      c.cloud,
		Location:                   c.location,
		VPCID:                      string(c.vpcID),
		SubnetIDs:                  subnetIDs,
		KubernetesVersion:          c.kubernetesVersion,
		Status:                     c.status.value,
		APIEndpoint:                c.apiEndpoint,
		ClusterCACertificateBase64: c.clusterCACertificateBase64,
		NodeGroups:                 c.nodeGroups,
		Tags:                       c.tags,
	})
}

// UnmarshalJSON implements the json.Unmarshaler interface
func (c *Cluster) UnmarshalJSON(data []byte) error {
	var clusterJSON clusterJSON
	if err := json.Unmarshal(data, &clusterJSON); err != nil {
		return errors.WrapAndTrace(err)
	}

	subnetIDs := make([]CloudProviderResourceID, len(clusterJSON.SubnetIDs))
	for i, subnetID := range clusterJSON.SubnetIDs {
		subnetIDs[i] = CloudProviderResourceID(subnetID)
	}

	status, err := stringToClusterStatus(clusterJSON.Status)
	if err != nil {
		return errors.WrapAndTrace(err)
	}

	newCluster, err := NewCluster(ClusterSettings{
		ID:                         CloudProviderResourceID(clusterJSON.ID),
		Name:                       clusterJSON.Name,
		RefID:                      clusterJSON.RefID,
		Provider:                   clusterJSON.Provider,
		Cloud:                      clusterJSON.Cloud,
		Location:                   clusterJSON.Location,
		VPCID:                      CloudProviderResourceID(clusterJSON.VPCID),
		SubnetIDs:                  subnetIDs,
		KubernetesVersion:          clusterJSON.KubernetesVersion,
		Status:                     status,
		APIEndpoint:                clusterJSON.APIEndpoint,
		ClusterCACertificateBase64: clusterJSON.ClusterCACertificateBase64,
		NodeGroups:                 clusterJSON.NodeGroups,
		Tags:                       clusterJSON.Tags,
	})
	if err != nil {
		return errors.WrapAndTrace(err)
	}

	*c = *newCluster
	return nil
}

// ClusterStatus represents the status of a Kubernetes cluster. Note for maintainers: this is defined as a struct
// rather than a type alias to ensure stronger compile-time type checking and to avoid the need for a validation function.
type ClusterStatus struct {
	value string
}

var (
	ClusterStatusUnknown   = ClusterStatus{value: "unknown"}
	ClusterStatusPending   = ClusterStatus{value: "pending"}
	ClusterStatusAvailable = ClusterStatus{value: "available"}
	ClusterStatusDeleting  = ClusterStatus{value: "deleting"}
	ClusterStatusFailed    = ClusterStatus{value: "failed"}
)

func (s ClusterStatus) String() string {
	return s.value
}

func stringToClusterStatus(status string) (ClusterStatus, error) {
	switch status {
	case ClusterStatusUnknown.value:
		return ClusterStatusUnknown, nil
	case ClusterStatusPending.value:
		return ClusterStatusPending, nil
	case ClusterStatusAvailable.value:
		return ClusterStatusAvailable, nil
	case ClusterStatusDeleting.value:
		return ClusterStatusDeleting, nil
	case ClusterStatusFailed.value:
		return ClusterStatusFailed, nil
	}
	return ClusterStatusUnknown, errors.Join(ErrClusterInvalidStatus, fmt.Errorf("invalid status: %s", status))
}

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
		errs = append(errs, ErrRefIDRequired)
	}
	if s.Name == "" {
		errs = append(errs, ErrNameRequired)
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
	diskSize Bytes

	// The status of the node group.
	status NodeGroupStatus

	// The tags associated with the node group.
	tags Tags
}

// nodeGroupJSON is the JSON representation of a NodeGroup. This struct is maintained separately from the core NodeGroup
// struct to allow for unexported fields to be used in the MarshalJSON and UnmarshalJSON methods.
type nodeGroupJSON struct {
	Name         string            `json:"name"`
	RefID        string            `json:"refID"`
	ID           string            `json:"id"`
	MinNodeCount int               `json:"minNodeCount"`
	MaxNodeCount int               `json:"maxNodeCount"`
	InstanceType string            `json:"instanceType"`
	DiskSize     Bytes             `json:"diskSize"`
	Status       string            `json:"status"`
	Tags         map[string]string `json:"tags"`
}

// MarshalJSON implements the json.Marshaler interface
func (n *NodeGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(nodeGroupJSON{
		Name:         n.name,
		RefID:        n.refID,
		ID:           string(n.id),
		MinNodeCount: n.minNodeCount,
		MaxNodeCount: n.maxNodeCount,
		InstanceType: n.instanceType,
		DiskSize:     n.diskSize,
		Status:       n.status.value,
		Tags:         n.tags,
	})
}

// UnmarshalJSON implements the json.Unmarshaler interface
func (n *NodeGroup) UnmarshalJSON(data []byte) error {
	var nodeGroupJSON nodeGroupJSON
	if err := json.Unmarshal(data, &nodeGroupJSON); err != nil {
		return errors.WrapAndTrace(err)
	}

	status, err := stringToNodeGroupStatus(nodeGroupJSON.Status)
	if err != nil {
		return errors.WrapAndTrace(err)
	}

	newNodeGroup, err := NewNodeGroup(NodeGroupSettings{
		Name:         nodeGroupJSON.Name,
		RefID:        nodeGroupJSON.RefID,
		ID:           CloudProviderResourceID(nodeGroupJSON.ID),
		MinNodeCount: nodeGroupJSON.MinNodeCount,
		MaxNodeCount: nodeGroupJSON.MaxNodeCount,
		InstanceType: nodeGroupJSON.InstanceType,
		DiskSize:     nodeGroupJSON.DiskSize,
		Status:       status,
		Tags:         nodeGroupJSON.Tags,
	})
	if err != nil {
		return errors.WrapAndTrace(err)
	}

	*n = *newNodeGroup
	return nil
}

// NodeGroupStatus represents the status of a Kubernetes node group. Note for maintainers: this is defined as a struct
// rather than a type alias to ensure stronger compile-time type checking and to avoid the need for a validation function.
type NodeGroupStatus struct {
	value string
}

func (s NodeGroupStatus) String() string {
	return s.value
}

var (
	NodeGroupStatusUnknown   = NodeGroupStatus{value: "unknown"}
	NodeGroupStatusPending   = NodeGroupStatus{value: "pending"}
	NodeGroupStatusAvailable = NodeGroupStatus{value: "available"}
	NodeGroupStatusDeleting  = NodeGroupStatus{value: "deleting"}
	NodeGroupStatusFailed    = NodeGroupStatus{value: "failed"}
)

func stringToNodeGroupStatus(status string) (NodeGroupStatus, error) {
	switch status {
	case NodeGroupStatusUnknown.value:
		return NodeGroupStatusUnknown, nil
	case NodeGroupStatusPending.value:
		return NodeGroupStatusPending, nil
	case NodeGroupStatusAvailable.value:
		return NodeGroupStatusAvailable, nil
	case NodeGroupStatusDeleting.value:
		return NodeGroupStatusDeleting, nil
	case NodeGroupStatusFailed.value:
		return NodeGroupStatusFailed, nil
	}
	return NodeGroupStatusUnknown, errors.Join(ErrNodeGroupInvalidStatus, fmt.Errorf("invalid status: %s", status))
}

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

func (n *NodeGroup) GetDiskSize() Bytes {
	return n.diskSize
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
	DiskSize Bytes

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
		errs = append(errs, ErrRefIDRequired)
	}
	if s.Name == "" {
		errs = append(errs, ErrNameRequired)
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
		diskSize:     settings.DiskSize,
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
		errs = append(errs, ErrClusterUserClusterNameRequired)
	}
	if s.ClusterCertificateAuthorityDataBase64 == "" {
		errs = append(errs, ErrClusterUserClusterCertificateAuthorityDataBase64Required)
	}
	if s.ClusterServerURL == "" {
		errs = append(errs, ErrClusterUserClusterServerURLRequired)
	}
	if s.Username == "" {
		errs = append(errs, ErrClusterUserUsernameRequired)
	}
	if s.UserClientCertificateDataBase64 == "" {
		errs = append(errs, ErrClusterUserUserClientCertificateDataBase64Required)
	}
	if s.UserClientKeyDataBase64 == "" {
		errs = append(errs, ErrClusterUserUserClientKeyDataBase64Required)
	}
	if s.KubeconfigBase64 == "" {
		errs = append(errs, ErrClusterUserKubeconfigBase64Required)
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
	DiskSize     Bytes
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
