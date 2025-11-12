package v1

import (
	"context"
	"encoding/base64"
	"fmt"
	"strings"

	nebiuscommon "github.com/nebius/gosdk/proto/nebius/common/v1"
	nebiusmk8s "github.com/nebius/gosdk/proto/nebius/mk8s/v1"
	nebiusvpc "github.com/nebius/gosdk/proto/nebius/vpc/v1"
	grpccodes "google.golang.org/grpc/codes"
	grpcstatus "google.golang.org/grpc/status"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	k8scmd "k8s.io/client-go/tools/clientcmd/api"

	"github.com/brevdev/cloud/internal/errors"
	cloudk8s "github.com/brevdev/cloud/internal/kubernetes"
	"github.com/brevdev/cloud/internal/rsa"
	v1 "github.com/brevdev/cloud/v1"
)

var (
	maxDiskSize = v1.NewBytes(v1.BytesValue(64), v1.Gibibyte)

	errVPCHasNoPublicSubnets             = fmt.Errorf("VPC must have at least one public subnet with a CIDR block larger than /24")
	errVPCHasNoPrivateSubnets            = fmt.Errorf("VPC must have at least one private subnet with a CIDR block larger than /24")
	errNoSubnetIDsSpecifiedForVPC        = fmt.Errorf("no subnet IDs specified for VPC")
	errMultipleSubnetIDsNotAllowedForVPC = fmt.Errorf("multiple subnet IDs not allowed for VPC")

	errNodeGroupNameIsRequired                                     = fmt.Errorf("node group name is required")
	errNodeGroupRefIDIsRequired                                    = fmt.Errorf("node group refID is required")
	errNodeGroupMinNodeCountMustBeGreaterThan0                     = fmt.Errorf("node group minNodeCount must be greater than 0")
	errNodeGroupMaxNodeCountMustBeGreaterThan0                     = fmt.Errorf("node group maxNodeCount must be greater than 0")
	errNodeGroupMaxNodeCountMustBeGreaterThanOrEqualToMinNodeCount = fmt.Errorf("node group maxNodeCount must be greater than or equal to minNodeCount")
	errNodeGroupDiskSizeMustBeGreaterThanOrEqualToMax              = fmt.Errorf("node group diskSize must be greater than or equal to %v", maxDiskSize)
	errNodeGroupInstanceTypeIsRequired                             = fmt.Errorf("node group instanceType is required")

	errUsernameIsRequired     = fmt.Errorf("username is required")
	errRoleIsRequired         = fmt.Errorf("role is required")
	errClusterIDIsRequired    = fmt.Errorf("cluster ID is required")
	errRSAPEMBase64IsRequired = fmt.Errorf("RSA PEM base64 is required")
)

var _ v1.CloudMaintainKubernetes = &NebiusClient{}

func (c *NebiusClient) CreateCluster(ctx context.Context, args v1.CreateClusterArgs) (*v1.Cluster, error) {
	nebiusClusterService := c.sdk.Services().MK8S().V1().Cluster()

	// Validate arguments
	if len(args.SubnetIDs) == 0 {
		return nil, errors.WrapAndTrace(errNoSubnetIDsSpecifiedForVPC)
	} else if len(args.SubnetIDs) > 1 {
		return nil, errors.WrapAndTrace(errMultipleSubnetIDsNotAllowedForVPC)
	}
	subnetID := string(args.SubnetIDs[0])

	// Fetch the target location
	location, err := c.GetLocation(ctx)
	if err != nil {
		return nil, errors.WrapAndTrace(err)
	}

	// Fetch the target VPC
	vpc, err := c.GetVPC(ctx, v1.GetVPCArgs{
		ID: args.VPCID,
	})
	if err != nil {
		return nil, errors.WrapAndTrace(err)
	}

	// Validate VPC
	err = validateVPC(vpc)
	if err != nil {
		return nil, errors.WrapAndTrace(err)
	}

	// Create a map of subnetID->subnet for this VPC so that we can find the target subnet
	subnetMap := make(map[string]*v1.Subnet)
	for _, subnet := range vpc.GetSubnets() {
		subnetMap[string(subnet.GetID())] = subnet
	}

	// Get the target subnet from the map
	var subnet *v1.Subnet
	if _, ok := subnetMap[subnetID]; !ok {
		return nil, errors.WrapAndTrace(fmt.Errorf("subnet ID %s does not match VPC %s", subnetID, vpc.GetID()))
	} else {
		subnet = subnetMap[subnetID]
	}

	labels := make(map[string]string)
	for key, value := range args.Tags {
		labels[key] = value
	}

	// Add the required labels
	labels[labelBrevRefID] = args.RefID
	labels[labelCreatedBy] = labelBrevCloudSDK

	// Create the cluster
	createClusterOperation, err := nebiusClusterService.Create(ctx, &nebiusmk8s.CreateClusterRequest{
		Metadata: &nebiuscommon.ResourceMetadata{
			Name:     args.Name,
			ParentId: c.projectID,
			Labels:   labels,
		},
		Spec: &nebiusmk8s.ClusterSpec{
			ControlPlane: &nebiusmk8s.ControlPlaneSpec{
				Version:         args.KubernetesVersion,
				SubnetId:        string(subnet.GetID()),
				EtcdClusterSize: 3,
				Endpoints: &nebiusmk8s.ControlPlaneEndpointsSpec{
					PublicEndpoint: &nebiusmk8s.PublicEndpointSpec{},
				},
			},
			KubeNetwork: &nebiusmk8s.KubeNetworkSpec{
				ServiceCidrs: []string{subnet.GetCidrBlock()},
			},
		},
	})
	if err != nil {
		return nil, errors.WrapAndTrace(err)
	}

	brevCluster, err := v1.NewCluster(v1.ClusterSettings{
		ID:                v1.CloudProviderResourceID(createClusterOperation.ResourceID()),
		Name:              args.Name,
		RefID:             args.RefID,
		Provider:          CloudProviderID,
		Cloud:             CloudProviderID,
		Location:          location,
		VPCID:             args.VPCID,
		SubnetIDs:         args.SubnetIDs,
		KubernetesVersion: args.KubernetesVersion,
		Status:            v1.ClusterStatusPending,
		Tags:              args.Tags,
	})
	if err != nil {
		return nil, errors.WrapAndTrace(err)
	}
	return brevCluster, nil
}

func validateVPC(vpc *v1.VPC) error {
	validPublicSubnetCount := 0
	validPrivateSubnetCount := 0

	for _, subnet := range vpc.GetSubnets() {
		larger, err := cidrBlockLargerThanMask(subnet.GetCidrBlock(), 24)
		if err != nil {
			return errors.WrapAndTrace(err)
		}
		if !larger {
			continue
		}

		if subnet.GetSubnetType() == v1.SubnetTypePublic {
			validPublicSubnetCount++
		} else {
			validPrivateSubnetCount++
		}
	}

	errs := []error{}
	if validPublicSubnetCount == 0 {
		errs = append(errs, errVPCHasNoPublicSubnets)
	}
	if validPrivateSubnetCount == 0 {
		errs = append(errs, errVPCHasNoPrivateSubnets)
	}

	return errors.WrapAndTrace(errors.Join(errs...))
}

func (c *NebiusClient) GetCluster(ctx context.Context, args v1.GetClusterArgs) (*v1.Cluster, error) {
	nebiusClusterService := c.sdk.Services().MK8S().V1().Cluster()
	nebiusSubnetService := c.sdk.Services().VPC().V1().Subnet()

	cluster, err := nebiusClusterService.Get(ctx, &nebiusmk8s.GetClusterRequest{
		Id: string(args.ID),
	})
	if err != nil {
		if grpcstatus.Code(err) == grpccodes.NotFound {
			return nil, v1.ErrResourceNotFound
		}
		return nil, errors.WrapAndTrace(err)
	}

	nebiusSubnet, err := nebiusSubnetService.Get(ctx, &nebiusvpc.GetSubnetRequest{
		Id: cluster.Spec.ControlPlane.SubnetId,
	})
	if err != nil {
		return nil, errors.WrapAndTrace(err)
	}

	nodeGroups, err := c.getClusterNodeGroups(ctx, cluster)
	if err != nil {
		return nil, errors.WrapAndTrace(err)
	}

	brevCluster, err := v1.NewCluster(v1.ClusterSettings{
		RefID:                      cluster.Metadata.Labels[labelBrevRefID],
		ID:                         v1.CloudProviderResourceID(cluster.Metadata.Id),
		Name:                       cluster.Metadata.Name,
		APIEndpoint:                getClusterAPIEndpoint(cluster),
		KubernetesVersion:          cluster.Spec.ControlPlane.Version,
		Status:                     parseNebiusClusterStatus(cluster.Status),
		VPCID:                      v1.CloudProviderResourceID(nebiusSubnet.Spec.NetworkId),
		SubnetIDs:                  []v1.CloudProviderResourceID{v1.CloudProviderResourceID(nebiusSubnet.Metadata.Id)},
		NodeGroups:                 nodeGroups,
		ClusterCACertificateBase64: getClusterCACertificateBase64(cluster),
		Tags:                       v1.Tags(cluster.Metadata.Labels),
	})
	if err != nil {
		return nil, errors.WrapAndTrace(err)
	}
	return brevCluster, err
}

func getClusterCACertificateBase64(cluster *nebiusmk8s.Cluster) string {
	if cluster == nil {
		return ""
	}
	if cluster.Status == nil || cluster.Status.ControlPlane == nil || cluster.Status.ControlPlane.Auth == nil {
		return ""
	}
	return base64.StdEncoding.EncodeToString([]byte(cluster.Status.ControlPlane.Auth.ClusterCaCertificate))
}

func getClusterAPIEndpoint(cluster *nebiusmk8s.Cluster) string {
	if cluster == nil {
		return ""
	}
	if cluster.Status == nil || cluster.Status.ControlPlane == nil || cluster.Status.ControlPlane.Endpoints == nil {
		return ""
	}
	return cluster.Status.ControlPlane.Endpoints.PublicEndpoint
}

func (c *NebiusClient) getClusterNodeGroups(ctx context.Context, cluster *nebiusmk8s.Cluster) ([]*v1.NodeGroup, error) {
	nebiusNodeGroupService := c.sdk.Services().MK8S().V1().NodeGroup()

	nebiusNodeGroups, err := nebiusNodeGroupService.List(ctx, &nebiusmk8s.ListNodeGroupsRequest{
		ParentId: cluster.Metadata.Id,
	})
	if err != nil {
		return nil, errors.WrapAndTrace(err)
	}

	nodeGroups := make([]*v1.NodeGroup, len(nebiusNodeGroups.Items))
	for i, nebiusNodeGroup := range nebiusNodeGroups.Items {
		brevNodeGroup, err := parseNebiusNodeGroup(nebiusNodeGroup)
		if err != nil {
			return nil, errors.WrapAndTrace(err)
		}
		nodeGroups[i] = brevNodeGroup
	}
	return nodeGroups, nil
}

func parseNebiusClusterStatus(status *nebiusmk8s.ClusterStatus) v1.ClusterStatus {
	if status == nil {
		return v1.ClusterStatusUnknown
	}
	switch status.State {
	case nebiusmk8s.ClusterStatus_PROVISIONING:
		return v1.ClusterStatusPending
	case nebiusmk8s.ClusterStatus_RUNNING:
		return v1.ClusterStatusAvailable
	case nebiusmk8s.ClusterStatus_DELETING:
		return v1.ClusterStatusDeleting
	}
	return v1.ClusterStatusUnknown
}

// SetClusterUser implements v1.CloudMaintainKubernetes.
func (c *NebiusClient) SetClusterUser(ctx context.Context, args v1.SetClusterUserArgs) (*v1.ClusterUser, error) {
	err := validatePutUserArgs(args)
	if err != nil {
		return nil, errors.WrapAndTrace(err)
	}

	// Fetch the cluster the user key will be added to
	cluster, err := c.GetCluster(ctx, v1.GetClusterArgs{
		ID: args.ClusterID,
	})
	if err != nil {
		return nil, errors.WrapAndTrace(fmt.Errorf("failed to get cluster: %w", err))
	}

	// Create a clientset to interact with the cluster using the bearer token and CA certificate
	clientset, err := c.newK8sClient(ctx, cluster)
	if err != nil {
		return nil, errors.WrapAndTrace(fmt.Errorf("failed to create clientset: %w", err))
	}

	// Prepare the private key for the CSR
	privateKeyBytes, err := base64.StdEncoding.DecodeString(args.RSAPEMBase64)
	if err != nil {
		return nil, errors.WrapAndTrace(fmt.Errorf("failed to decode base64 string: %w", err))
	}

	// Parse the private key
	privateKey, err := rsa.BytesToRSAKey(privateKeyBytes)
	if err != nil {
		return nil, errors.WrapAndTrace(fmt.Errorf("failed to parse private key: %w", err))
	}

	// Create the client certificate to allow for external access to the cluster for the holders of this private key
	signedCertificate, err := cloudk8s.ClientCertificateData(ctx, clientset, args.Username, privateKey)
	if err != nil {
		return nil, errors.WrapAndTrace(fmt.Errorf("failed to get signed certificate: %w", err))
	}

	// Make the user a cluster admin
	err = cloudk8s.SetUserRole(ctx, clientset, args.Username, args.Role)
	if err != nil {
		return nil, errors.WrapAndTrace(fmt.Errorf("failed to set user role: %w", err))
	}

	// Get the certificate authority data
	certificateAuthorityData, err := base64.StdEncoding.DecodeString(cluster.GetClusterCACertificateBase64())
	if err != nil {
		return nil, errors.WrapAndTrace(fmt.Errorf("failed to decode certificate authority data: %w", err))
	}

	// Generate the complete kubeconfig
	kubeconfigBytes, err := clientcmd.Write(k8scmd.Config{
		Kind:       "Config",
		APIVersion: "v1",
		Clusters: map[string]*k8scmd.Cluster{
			cluster.GetRefID(): {
				Server:                   cluster.GetAPIEndpoint(),
				CertificateAuthorityData: certificateAuthorityData,
			},
		},
		AuthInfos: map[string]*k8scmd.AuthInfo{
			cluster.GetRefID(): {
				ClientCertificateData: signedCertificate,
				ClientKeyData:         privateKeyBytes,
			},
		},
	})
	if err != nil {
		return nil, errors.WrapAndTrace(fmt.Errorf("failed to write kubeconfig: %w", err))
	}

	brevClusterUser, err := v1.NewClusterUser(v1.ClusterUserSettings{
		ClusterName:                           cluster.GetRefID(),
		ClusterCertificateAuthorityDataBase64: cluster.GetClusterCACertificateBase64(),
		ClusterServerURL:                      cluster.GetAPIEndpoint(),
		Username:                              args.Username,
		UserClientCertificateDataBase64:       base64.StdEncoding.EncodeToString(signedCertificate),
		UserClientKeyDataBase64:               base64.StdEncoding.EncodeToString(privateKeyBytes),
		KubeconfigBase64:                      base64.StdEncoding.EncodeToString(kubeconfigBytes),
	})
	if err != nil {
		return nil, errors.WrapAndTrace(fmt.Errorf("failed to create cluster user: %w", err))
	}
	return brevClusterUser, nil
}

func validatePutUserArgs(args v1.SetClusterUserArgs) error {
	errs := []error{}
	if args.Username == "" {
		errs = append(errs, errUsernameIsRequired)
	}
	if args.Role == "" {
		errs = append(errs, errRoleIsRequired)
	}
	if args.ClusterID == "" {
		errs = append(errs, errClusterIDIsRequired)
	}
	if args.RSAPEMBase64 == "" {
		errs = append(errs, errRSAPEMBase64IsRequired)
	}
	return errors.WrapAndTrace(errors.Join(errs...))
}

func (c *NebiusClient) CreateNodeGroup(ctx context.Context, args v1.CreateNodeGroupArgs) (*v1.NodeGroup, error) {
	nebiusNodeGroupService := c.sdk.Services().MK8S().V1().NodeGroup()

	err := validateCreateNodeGroupArgs(args)
	if err != nil {
		return nil, errors.WrapAndTrace(err)
	}

	// Fetch the cluster the user key will be added to
	cluster, err := c.GetCluster(ctx, v1.GetClusterArgs{
		ID: args.ClusterID,
	})
	if err != nil {
		return nil, errors.WrapAndTrace(fmt.Errorf("failed to get cluster: %w", err))
	}

	// Placeholder for parsing instance type
	parts := strings.Split(args.InstanceType, ".")
	platform := parts[0]
	preset := parts[1]

	labels := make(map[string]string)
	for key, value := range args.Tags {
		labels[key] = value
	}

	// Add the required labels
	labels[labelBrevRefID] = args.RefID
	labels[labelCreatedBy] = labelBrevCloudSDK

	// Nebius expects the disk size in GiB, so we need to convert the disk size to GiB
	diskSizeGiB, err := args.DiskSize.ByteCountInUnitInt64(v1.Gibibyte)
	if err != nil {
		return nil, errors.WrapAndTrace(err)
	}

	// create the node groups
	createNodeGroupOperation, err := nebiusNodeGroupService.Create(ctx, &nebiusmk8s.CreateNodeGroupRequest{
		Metadata: &nebiuscommon.ResourceMetadata{
			Name:     args.Name,
			ParentId: string(cluster.GetID()),
			Labels:   labels,
		},
		Spec: &nebiusmk8s.NodeGroupSpec{
			Size: &nebiusmk8s.NodeGroupSpec_Autoscaling{
				Autoscaling: &nebiusmk8s.NodeGroupAutoscalingSpec{
					MinNodeCount: int64(args.MinNodeCount),
					MaxNodeCount: int64(args.MaxNodeCount),
				},
			},
			Template: &nebiusmk8s.NodeTemplate{
				Resources: &nebiusmk8s.ResourcesSpec{
					Platform: platform,
					Size: &nebiusmk8s.ResourcesSpec_Preset{
						Preset: preset,
					},
				},
				GpuSettings: &nebiusmk8s.GpuSettings{
					DriversPreset: "cuda12",
				},
				BootDisk: &nebiusmk8s.DiskSpec{
					Type: nebiusmk8s.DiskSpec_NETWORK_SSD,
					Size: &nebiusmk8s.DiskSpec_SizeGibibytes{
						SizeGibibytes: diskSizeGiB,
					},
				},
			},
		},
	})
	if err != nil {
		return nil, errors.WrapAndTrace(err)
	}

	brevNodeGroup, err := v1.NewNodeGroup(v1.NodeGroupSettings{
		ID:           v1.CloudProviderResourceID(createNodeGroupOperation.ResourceID()),
		Name:         args.Name,
		RefID:        args.RefID,
		MinNodeCount: args.MinNodeCount,
		MaxNodeCount: args.MaxNodeCount,
		InstanceType: args.InstanceType,
		DiskSize:     args.DiskSize,
		Status:       v1.NodeGroupStatusPending,
		Tags:         args.Tags,
	})
	if err != nil {
		return nil, errors.WrapAndTrace(err)
	}
	return brevNodeGroup, nil
}

func validateCreateNodeGroupArgs(args v1.CreateNodeGroupArgs) error {
	if args.Name == "" {
		return errNodeGroupNameIsRequired
	}
	if args.RefID == "" {
		return errNodeGroupRefIDIsRequired
	}
	if args.MinNodeCount < 1 {
		return errNodeGroupMinNodeCountMustBeGreaterThan0
	}
	if args.MaxNodeCount < 1 {
		return errNodeGroupMaxNodeCountMustBeGreaterThan0
	}
	if args.MaxNodeCount < args.MinNodeCount {
		return errNodeGroupMaxNodeCountMustBeGreaterThanOrEqualToMinNodeCount
	}
	if args.DiskSize.LessThan(maxDiskSize) {
		return errNodeGroupDiskSizeMustBeGreaterThanOrEqualToMax
	}
	if args.InstanceType == "" {
		return errNodeGroupInstanceTypeIsRequired
	}
	return nil
}

func (c *NebiusClient) GetNodeGroup(ctx context.Context, args v1.GetNodeGroupArgs) (*v1.NodeGroup, error) {
	nebiusNodeGroupService := c.sdk.Services().MK8S().V1().NodeGroup()

	nodeGroup, err := nebiusNodeGroupService.Get(ctx, &nebiusmk8s.GetNodeGroupRequest{
		Id: string(args.ID),
	})
	if err != nil {
		if grpcstatus.Code(err) == grpccodes.NotFound {
			return nil, v1.ErrResourceNotFound
		}
		return nil, errors.WrapAndTrace(err)
	}

	brevNodeGroup, err := parseNebiusNodeGroup(nodeGroup)
	if err != nil {
		return nil, errors.WrapAndTrace(err)
	}
	return brevNodeGroup, nil
}

func parseNebiusNodeGroup(nodeGroup *nebiusmk8s.NodeGroup) (*v1.NodeGroup, error) {
	brevNodeGroup, err := v1.NewNodeGroup(v1.NodeGroupSettings{
		ID:           v1.CloudProviderResourceID(nodeGroup.Metadata.Id),
		RefID:        nodeGroup.Metadata.Labels[labelBrevRefID],
		Name:         nodeGroup.Metadata.Name,
		MinNodeCount: int(nodeGroup.Spec.GetAutoscaling().MinNodeCount),
		MaxNodeCount: int(nodeGroup.Spec.GetAutoscaling().MaxNodeCount),
		InstanceType: nodeGroup.Spec.Template.Resources.Platform + "." + nodeGroup.Spec.Template.Resources.GetPreset(),
		DiskSize:     v1.NewBytes(v1.BytesValue(nodeGroup.Spec.Template.BootDisk.GetSizeGibibytes()), v1.Gibibyte),
		Status:       parseNebiusNodeGroupStatus(nodeGroup.Status),
		Tags:         v1.Tags(nodeGroup.Metadata.Labels),
	})
	if err != nil {
		return nil, errors.WrapAndTrace(err)
	}
	return brevNodeGroup, nil
}

func parseNebiusNodeGroupStatus(status *nebiusmk8s.NodeGroupStatus) v1.NodeGroupStatus {
	if status == nil {
		return v1.NodeGroupStatusUnknown
	}
	switch status.State {
	case nebiusmk8s.NodeGroupStatus_PROVISIONING:
		return v1.NodeGroupStatusPending
	case nebiusmk8s.NodeGroupStatus_RUNNING:
		return v1.NodeGroupStatusAvailable
	case nebiusmk8s.NodeGroupStatus_DELETING:
		return v1.NodeGroupStatusDeleting
	}
	return v1.NodeGroupStatusUnknown
}

func (c *NebiusClient) ModifyNodeGroup(ctx context.Context, args v1.ModifyNodeGroupArgs) error {
	nebiusNodeGroupService := c.sdk.Services().MK8S().V1().NodeGroup()

	err := validateModifyNodeGroupArgs(args)
	if err != nil {
		return errors.WrapAndTrace(err)
	}

	nodeGroup, err := nebiusNodeGroupService.Get(ctx, &nebiusmk8s.GetNodeGroupRequest{
		Id: string(args.ID),
	})
	if err != nil {
		return errors.WrapAndTrace(err)
	}

	_, err = nebiusNodeGroupService.Update(ctx, &nebiusmk8s.UpdateNodeGroupRequest{
		Metadata: &nebiuscommon.ResourceMetadata{
			Id:       nodeGroup.Metadata.Id,
			Name:     nodeGroup.Metadata.Name,
			ParentId: nodeGroup.Metadata.ParentId,
			Labels:   nodeGroup.Metadata.Labels,
		},
		Spec: &nebiusmk8s.NodeGroupSpec{
			Size: &nebiusmk8s.NodeGroupSpec_Autoscaling{
				Autoscaling: &nebiusmk8s.NodeGroupAutoscalingSpec{
					MinNodeCount: int64(args.MinNodeCount),
					MaxNodeCount: int64(args.MaxNodeCount),
				},
			},
			Template: &nebiusmk8s.NodeTemplate{
				Resources: &nebiusmk8s.ResourcesSpec{
					Platform: nodeGroup.Spec.GetTemplate().GetResources().GetPlatform(),
					Size: &nebiusmk8s.ResourcesSpec_Preset{
						Preset: nodeGroup.Spec.GetTemplate().GetResources().GetPreset(),
					},
				},
				GpuSettings: &nebiusmk8s.GpuSettings{
					DriversPreset: nodeGroup.Spec.GetTemplate().GetGpuSettings().GetDriversPreset(),
				},
				BootDisk: &nebiusmk8s.DiskSpec{
					Type: nodeGroup.Spec.GetTemplate().GetBootDisk().GetType(),
					Size: &nebiusmk8s.DiskSpec_SizeGibibytes{
						SizeGibibytes: nodeGroup.Spec.GetTemplate().GetBootDisk().GetSizeGibibytes(),
					},
				},
			},
		},
	})
	if err != nil {
		return errors.WrapAndTrace(fmt.Errorf("failed to modify node group: %w", err))
	}

	return nil
}

func validateModifyNodeGroupArgs(args v1.ModifyNodeGroupArgs) error {
	if args.MinNodeCount < 1 {
		return errNodeGroupMinNodeCountMustBeGreaterThan0
	}
	if args.MaxNodeCount < 1 {
		return errNodeGroupMaxNodeCountMustBeGreaterThan0
	}
	if args.MaxNodeCount < args.MinNodeCount {
		return errNodeGroupMaxNodeCountMustBeGreaterThanOrEqualToMinNodeCount
	}
	return nil
}

func (c *NebiusClient) DeleteNodeGroup(ctx context.Context, args v1.DeleteNodeGroupArgs) error {
	nebiusNodeGroupService := c.sdk.Services().MK8S().V1().NodeGroup()

	nodeGroup, err := c.GetNodeGroup(ctx, v1.GetNodeGroupArgs{
		ID: args.ID,
	})
	if err != nil {
		return errors.WrapAndTrace(fmt.Errorf("failed to get node group: %w", err))
	}

	_, err = nebiusNodeGroupService.Delete(ctx, &nebiusmk8s.DeleteNodeGroupRequest{
		Id: string(nodeGroup.GetID()),
	})
	if err != nil {
		return errors.WrapAndTrace(fmt.Errorf("failed to delete node group: %w", err))
	}

	return nil
}

func (c *NebiusClient) DeleteCluster(ctx context.Context, args v1.DeleteClusterArgs) error {
	nebiusClusterService := c.sdk.Services().MK8S().V1().Cluster()

	// Fetch the cluster the user key will be added to
	cluster, err := c.GetCluster(ctx, v1.GetClusterArgs{ //nolint:staticcheck // prefer explicit struct literal
		ID: args.ID,
	})
	if err != nil {
		return errors.WrapAndTrace(fmt.Errorf("failed to get cluster: %w", err))
	}

	_, err = nebiusClusterService.Delete(ctx, &nebiusmk8s.DeleteClusterRequest{
		Id: string(cluster.GetID()),
	})
	if err != nil {
		return errors.WrapAndTrace(fmt.Errorf("failed to delete cluster: %w", err))
	}

	return nil
}

func (c *NebiusClient) newK8sClient(ctx context.Context, cluster *v1.Cluster) (*kubernetes.Clientset, error) {
	// Decode the cluster CA certificate
	clusterCACertificate, err := base64.StdEncoding.DecodeString(cluster.GetClusterCACertificateBase64())
	if err != nil {
		return nil, errors.WrapAndTrace(fmt.Errorf("failed to decode cluster CA certificate: %w", err))
	}

	// Get a bearer token to authenticate to the cluster
	bearerToken, err := c.sdk.BearerToken(ctx)
	if err != nil {
		return nil, errors.WrapAndTrace(fmt.Errorf("failed to get bearer token: %w", err))
	}

	// Create a clientset to interact with the cluster using the bearer token and CA certificate
	clientset, err := kubernetes.NewForConfig(&rest.Config{
		Host:        cluster.GetAPIEndpoint(),
		BearerToken: bearerToken.Token,
		TLSClientConfig: rest.TLSClientConfig{
			CAData: clusterCACertificate,
		},
	})
	if err != nil {
		return nil, errors.WrapAndTrace(fmt.Errorf("failed to create clientset: %w", err))
	}

	return clientset, nil
}
