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

var _ v1.CloudMaintainKubernetes = &NebiusClient{}

func (c *NebiusClient) CreateCluster(ctx context.Context, args v1.CreateClusterArgs) (*v1.Cluster, error) {
	nebiusClusterService := c.sdk.Services().MK8S().V1().Cluster()

	vpc, err := c.GetVPC(ctx, v1.GetVPCArgs{
		ID: args.VPCID,
	})
	if err != nil {
		return nil, errors.WrapAndTrace(err)
	}

	// validate
	if len(args.SubnetIDs) == 0 {
		return nil, errors.WrapAndTrace(fmt.Errorf("no subnet IDs specified for VPC %s", vpc.ID))
	} else if len(args.SubnetIDs) > 1 {
		return nil, errors.WrapAndTrace(fmt.Errorf("multiple subnet IDs not allowed for VPC %s", vpc.ID))
	}
	subnetID := string(args.SubnetIDs[0])

	// make a map of ID to subnet for this VPC
	subnetMap := make(map[string]*v1.Subnet)
	for _, subnet := range vpc.Subnets {
		subnetMap[string(subnet.ID)] = subnet
	}

	// get the specified subnet
	var subnet *v1.Subnet
	if _, ok := subnetMap[subnetID]; !ok {
		return nil, errors.WrapAndTrace(fmt.Errorf("subnet ID %s does not match VPC %s", subnetID, vpc.ID))
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

	createClusterOperation, err := nebiusClusterService.Create(ctx, &nebiusmk8s.CreateClusterRequest{
		Metadata: &nebiuscommon.ResourceMetadata{
			Name:     args.Name,
			ParentId: c.projectID,
			Labels:   labels,
		},
		Spec: &nebiusmk8s.ClusterSpec{
			ControlPlane: &nebiusmk8s.ControlPlaneSpec{
				Version:         args.KubernetesVersion,
				SubnetId:        string(subnet.ID),
				EtcdClusterSize: 3,
				Endpoints: &nebiusmk8s.ControlPlaneEndpointsSpec{
					PublicEndpoint: &nebiusmk8s.PublicEndpointSpec{},
				},
			},
			KubeNetwork: &nebiusmk8s.KubeNetworkSpec{
				ServiceCidrs: []string{subnet.CidrBlock},
			},
		},
	})
	if err != nil {
		return nil, errors.WrapAndTrace(err)
	}

	return &v1.Cluster{
		ID:                v1.CloudProviderResourceID(createClusterOperation.ResourceID()),
		Name:              args.Name,
		RefID:             args.RefID,
		Provider:          "nebius",
		Cloud:             "nebius",
		Location:          args.Location,
		VPCID:             args.VPCID,
		SubnetIDs:         args.SubnetIDs,
		KubernetesVersion: args.KubernetesVersion,
		Status:            v1.ClusterStatusPending,
		Tags:              args.Tags,
	}, nil
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

	return &v1.Cluster{
		RefID:                      cluster.Metadata.Labels[labelBrevRefID],
		ID:                         v1.CloudProviderResourceID(cluster.Metadata.Id),
		Name:                       cluster.Metadata.Name,
		APIEndpoint:                getClusterAPIEndpoint(cluster),
		KubernetesVersion:          cluster.Spec.ControlPlane.Version,
		Status:                     parseNebiusClusterStatus(cluster.Status),
		VPCID:                      v1.CloudProviderResourceID(nebiusSubnet.Spec.NetworkId),
		SubnetIDs:                  []v1.CloudProviderResourceID{v1.CloudProviderResourceID(nebiusSubnet.Metadata.Id)},
		ClusterCACertificateBase64: getClusterCACertificateBase64(cluster),
		NodeGroups:                 nodeGroups,
		Tags:                       v1.Tags(cluster.Metadata.Labels),
	}, nil
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

	nodeGroups := make([]*v1.NodeGroup, 0)
	for _, nebiusNodeGroup := range nebiusNodeGroups.Items {
		nodeGroups = append(nodeGroups, parseNebiusNodeGroup(nebiusNodeGroup))
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

// PutUser implements v1.CloudMaintainKubernetes.
func (c *NebiusClient) PutUser(ctx context.Context, args v1.PutUserArgs) (*v1.PutUserResponse, error) {
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
	certificateAuthorityData, err := base64.StdEncoding.DecodeString(cluster.ClusterCACertificateBase64)
	if err != nil {
		return nil, errors.WrapAndTrace(fmt.Errorf("failed to decode certificate authority data: %w", err))
	}

	// Generate the complete kubeconfig
	kubeconfigBytes, err := clientcmd.Write(k8scmd.Config{
		Kind:       "Config",
		APIVersion: "v1",
		Clusters: map[string]*k8scmd.Cluster{
			cluster.RefID: {
				Server:                   cluster.APIEndpoint,
				CertificateAuthorityData: certificateAuthorityData,
			},
		},
		AuthInfos: map[string]*k8scmd.AuthInfo{
			cluster.RefID: {
				ClientCertificateData: signedCertificate,
				ClientKeyData:         privateKeyBytes,
			},
		},
	})
	if err != nil {
		return nil, errors.WrapAndTrace(fmt.Errorf("failed to write kubeconfig: %w", err))
	}

	return &v1.PutUserResponse{
		ClusterName:                           cluster.RefID,
		ClusterCertificateAuthorityDataBase64: cluster.ClusterCACertificateBase64,
		ClusterServerURL:                      cluster.APIEndpoint,
		Username:                              args.Username,
		UserClientCertificateDataBase64:       base64.StdEncoding.EncodeToString(signedCertificate),
		UserClientKeyDataBase64:               base64.StdEncoding.EncodeToString(privateKeyBytes),
		KubeconfigBase64:                      base64.StdEncoding.EncodeToString(kubeconfigBytes),
	}, nil
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

	// create the node groups
	createNodeGroupOperation, err := nebiusNodeGroupService.Create(ctx, &nebiusmk8s.CreateNodeGroupRequest{
		Metadata: &nebiuscommon.ResourceMetadata{
			Name:     args.Name,
			ParentId: string(cluster.ID),
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
						SizeGibibytes: int64(args.DiskSizeGiB),
					},
				},
			},
		},
	})
	if err != nil {
		return nil, errors.WrapAndTrace(err)
	}

	return &v1.NodeGroup{
		ID:           v1.CloudProviderResourceID(createNodeGroupOperation.ResourceID()),
		Name:         args.Name,
		RefID:        args.RefID,
		MinNodeCount: args.MinNodeCount,
		MaxNodeCount: args.MaxNodeCount,
		InstanceType: args.InstanceType,
		DiskSizeGiB:  args.DiskSizeGiB,
		Status:       v1.NodeGroupStatusPending,
		Tags:         args.Tags,
	}, nil
}

func validateCreateNodeGroupArgs(args v1.CreateNodeGroupArgs) error {
	if args.Name == "" {
		return fmt.Errorf("node group name is required")
	}
	if args.RefID == "" {
		return fmt.Errorf("node group refID is required")
	}
	if args.MinNodeCount < 1 {
		return fmt.Errorf("node group minNodeCount must be greater than 0")
	}
	if args.MaxNodeCount < 1 {
		return fmt.Errorf("node group maxNodeCount must be greater than 0")
	}
	if args.MaxNodeCount < args.MinNodeCount {
		return fmt.Errorf("node group maxNodeCount must be greater than or equal to minNodeCount")
	}
	if args.DiskSizeGiB < 64 {
		return fmt.Errorf("node group diskSizeGiB must be greater than or equal to 64")
	}
	if args.InstanceType == "" {
		return fmt.Errorf("node group instanceType is required")
	}
	return nil
}

func (c *NebiusClient) GetNodeGroup(ctx context.Context, args v1.GetNodeGroupArgs) (*v1.NodeGroup, error) {
	nebiusNodeGroupService := c.sdk.Services().MK8S().V1().NodeGroup()

	nodeGroup, err := nebiusNodeGroupService.Get(ctx, &nebiusmk8s.GetNodeGroupRequest{
		Id: string(args.ID),
	})
	if err != nil {
		return nil, errors.WrapAndTrace(err)
	}

	return parseNebiusNodeGroup(nodeGroup), nil
}

func parseNebiusNodeGroup(nodeGroup *nebiusmk8s.NodeGroup) *v1.NodeGroup {
	return &v1.NodeGroup{
		ID:           v1.CloudProviderResourceID(nodeGroup.Metadata.Id),
		RefID:        nodeGroup.Metadata.Labels[labelBrevRefID],
		Name:         nodeGroup.Metadata.Name,
		MinNodeCount: int(nodeGroup.Spec.GetAutoscaling().MinNodeCount),
		MaxNodeCount: int(nodeGroup.Spec.GetAutoscaling().MaxNodeCount),
		InstanceType: nodeGroup.Spec.Template.Resources.Platform + "." + nodeGroup.Spec.Template.Resources.GetPreset(),
		DiskSizeGiB:  int(nodeGroup.Spec.Template.BootDisk.GetSizeGibibytes()),
		Status:       parseNebiusNodeGroupStatus(nodeGroup.Status),
		Tags:         v1.Tags(nodeGroup.Metadata.Labels),
	}
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
		return fmt.Errorf("node group minNodeCount must be greater than 0")
	}
	if args.MaxNodeCount < 1 {
		return fmt.Errorf("node group maxNodeCount must be greater than 0")
	}
	if args.MaxNodeCount < args.MinNodeCount {
		return fmt.Errorf("node group maxNodeCount must be greater than or equal to minNodeCount")
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
		Id: string(nodeGroup.ID),
	})
	if err != nil {
		return errors.WrapAndTrace(fmt.Errorf("failed to delete node group: %w", err))
	}

	return nil
}

func (c *NebiusClient) DeleteCluster(ctx context.Context, args v1.DeleteClusterArgs) error {
	nebiusClusterService := c.sdk.Services().MK8S().V1().Cluster()

	// Fetch the cluster the user key will be added to
	cluster, err := c.GetCluster(ctx, v1.GetClusterArgs{
		ID: args.ID,
	})
	if err != nil {
		return errors.WrapAndTrace(fmt.Errorf("failed to get cluster: %w", err))
	}

	_, err = nebiusClusterService.Delete(ctx, &nebiusmk8s.DeleteClusterRequest{
		Id: string(cluster.ID),
	})
	if err != nil {
		return errors.WrapAndTrace(fmt.Errorf("failed to delete cluster: %w", err))
	}

	return nil
}

func (c *NebiusClient) newK8sClient(ctx context.Context, cluster *v1.Cluster) (*kubernetes.Clientset, error) {
	// Decode the cluster CA certificate
	clusterCACertificate, err := base64.StdEncoding.DecodeString(cluster.ClusterCACertificateBase64)
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
		Host:        cluster.APIEndpoint,
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
