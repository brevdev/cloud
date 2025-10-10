package v1

import (
	"context"
	"encoding/base64"
	"fmt"
	"strings"

	cloudk8s "github.com/brevdev/cloud/internal/kubernetes"
	"github.com/brevdev/cloud/internal/rsa"
	v1 "github.com/brevdev/cloud/v1"

	common "github.com/nebius/gosdk/proto/nebius/common/v1"
	mk8s "github.com/nebius/gosdk/proto/nebius/mk8s/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
)

var _ v1.CloudMaintainKubernetes = &NebiusClient{}

func (c *NebiusClient) CreateCluster(ctx context.Context, args v1.CreateClusterArgs) (*v1.Cluster, error) {
	nebiusClusterService := c.sdk.Services().MK8S().V1().Cluster()

	vpc, err := c.GetVPC(ctx, v1.GetVPCArgs{
		CloudID: args.VPCID,
	})
	if err != nil {
		return nil, err
	}

	// validate
	if len(args.SubnetIDs) == 0 {
		return nil, fmt.Errorf("no subnet IDs specified for VPC %s", vpc.CloudID)
	} else if len(args.SubnetIDs) > 1 {
		return nil, fmt.Errorf("multiple subnet IDs not allowed for VPC %s", vpc.CloudID)
	}
	subnetID := args.SubnetIDs[0]

	// make a map of ID to subnet for this VPC
	subnetMap := make(map[string]v1.Subnet)
	for _, subnet := range vpc.Subnets {
		subnetMap[subnet.CloudID] = subnet
	}

	// get the specified subnet
	var subnet v1.Subnet
	if _, ok := subnetMap[subnetID]; !ok {
		return nil, fmt.Errorf("subnet ID %s does not match VPC %s", subnetID, vpc.CloudID)
	} else {
		subnet = subnetMap[subnetID]
	}

	createClusterOperation, err := nebiusClusterService.Create(ctx, &mk8s.CreateClusterRequest{
		Metadata: &common.ResourceMetadata{
			Name:     args.Name,
			ParentId: c.projectID,
			Labels: map[string]string{
				labelBrevRefID: args.RefID,
				labelCreatedBy: labelBrevCloudSDK,
			},
		},
		Spec: &mk8s.ClusterSpec{
			ControlPlane: &mk8s.ControlPlaneSpec{
				Version:         args.KubernetesVersion,
				SubnetId:        subnet.CloudID,
				EtcdClusterSize: 3,
				Endpoints: &mk8s.ControlPlaneEndpointsSpec{
					PublicEndpoint: &mk8s.PublicEndpointSpec{},
				},
			},
			KubeNetwork: &mk8s.KubeNetworkSpec{
				ServiceCidrs: []string{subnet.CidrBlock},
			},
		},
	})
	if err != nil {
		return nil, err
	}

	createClusterOperation, err = createClusterOperation.Wait(ctx)
	if err != nil {
		return nil, err
	}

	return &v1.Cluster{
		CloudID: createClusterOperation.ResourceID(),
	}, nil
}

func (c *NebiusClient) GetCluster(ctx context.Context, args v1.GetClusterArgs) (*v1.Cluster, error) {
	nebiusClusterService := c.sdk.Services().MK8S().V1().Cluster()

	var cluster *mk8s.Cluster
	var err error
	if args.CloudID == "" {
		cluster, err = nebiusClusterService.GetByName(ctx, &common.GetByNameRequest{
			ParentId: c.projectID,
			Name:     args.RefID,
		})
		if err != nil {
			return nil, err
		}
	} else {
		cluster, err = nebiusClusterService.Get(ctx, &mk8s.GetClusterRequest{
			Id: args.CloudID,
		})
		if err != nil {
			return nil, err
		}
	}

	clusterCACertificate := cluster.Status.ControlPlane.Auth.ClusterCaCertificate
	clusterCACertificateBase64 := base64.StdEncoding.EncodeToString([]byte(clusterCACertificate))

	return &v1.Cluster{
		RefID:                      cluster.Metadata.Labels[labelBrevRefID],
		CloudID:                    cluster.Metadata.Id,
		APIEndpoint:                cluster.Status.ControlPlane.Endpoints.PublicEndpoint,
		ClusterCACertificateBase64: clusterCACertificateBase64,
	}, nil
}

// PutUser implements v1.CloudMaintainKubernetes.
func (c *NebiusClient) PutUser(ctx context.Context, args v1.PutUserArgs) (*v1.PutUserResponse, error) {
	// Fetch the cluster the user key will be added to
	cluster, err := c.GetCluster(ctx, v1.GetClusterArgs{
		RefID: args.ClusterRefID,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get cluster: %w", err)
	}

	// Create a clientset to interact with the cluster using the bearer token and CA certificate
	clientset, err := c.newClusterClient(ctx, cluster)
	if err != nil {
		return nil, fmt.Errorf("failed to create clientset: %w", err)
	}

	// Prepare the private key for the CSR
	privateKeyBytes, err := base64.StdEncoding.DecodeString(args.RSAPEMBase64)
	if err != nil {
		return nil, fmt.Errorf("failed to decode base64 string: %w", err)
	}

	// Parse the private key
	privateKey, err := rsa.BytesToRSAKey(privateKeyBytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse private key: %w", err)
	}

	// Create the client certificate to allow for external access to the cluster for the holders of this private key
	signedCertificate, err := cloudk8s.ClientCertificateData(ctx, clientset, args.Username, privateKey)
	if err != nil {
		return nil, fmt.Errorf("failed to get signed certificate: %w", err)
	}

	// Make the user a cluster admin
	err = cloudk8s.SetUserRole(ctx, clientset, args.Username, "cluster-admin")
	if err != nil {
		return nil, fmt.Errorf("failed to set user role: %w", err)
	}

	// Get the certificate authority data
	certificateAuthorityData, err := base64.StdEncoding.DecodeString(cluster.ClusterCACertificateBase64)
	if err != nil {
		return nil, fmt.Errorf("failed to decode certificate authority data: %w", err)
	}

	// Generate the complete kubeconfig
	kubeconfigBytes, err := clientcmd.Write(clientcmdapi.Config{
		Kind:       "Config",
		APIVersion: "v1",
		Clusters: map[string]*clientcmdapi.Cluster{
			cluster.RefID: {
				Server:                   cluster.APIEndpoint,
				CertificateAuthorityData: certificateAuthorityData,
			},
		},
		AuthInfos: map[string]*clientcmdapi.AuthInfo{
			cluster.RefID: {
				ClientCertificateData: signedCertificate,
				ClientKeyData:         privateKeyBytes,
			},
		},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to write kubeconfig: %w", err)
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

var platformPresetMap = map[string][]string{
	"cpu-d3":       {"4vcpu-16gb", "8vcpu-32gb", "16vcpu-64gb", "32vcpu-128gb", "48vcpu-192gb", "64vcpu-256gb", "96vcpu-384gb", "128vcpu-512gb"},
	"cpu-e2":       {"2vcpu-8gb", "4vcpu-16gb", "8vcpu-32gb", "16vcpu-64gb", "32vcpu-128gb", "48vcpu-192gb", "64vcpu-256gb", "80vcpu-320gb"},
	"gpu-h200-sxm": {"1gpu-16vcpu-200gb", "8gpu-128vcpu-1600gb"},
	"gpu-h100-sxm": {"1gpu-16vcpu-200gb", "8gpu-128vcpu-1600gb"},
	"gpu-l40s-a":   {"1gpu-8vcpu-32gb", "1gpu-16vcpu-64gb", "1gpu-24vcpu-96gb", "1gpu-32vcpu-128gb", "1gpu-40vcpu-160gb"},
	"gpu-l40s-d":   {"1gpu-16vcpu-96gb", "1gpu-32vcpu-192gb", "1gpu-48vcpu-288gb", "2gpu-64vcpu-384gb", "2gpu-96vcpu-576gb", "4gpu-128vcpu-768gb", "4gpu-192vcpu-1152gb"},
}

func (c *NebiusClient) CreateNodeGroup(ctx context.Context, args v1.CreateNodeGroupArgs) (*v1.CreateNodeGroupResponse, error) {
	nebiusNodeGroupService := c.sdk.Services().MK8S().V1().NodeGroup()

	// Fetch the cluster the user key will be added to
	cluster, err := c.GetCluster(ctx, v1.GetClusterArgs{
		RefID: args.ClusterRefID,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get cluster: %w", err)
	}

	// Placeholder for parsing instance type
	parts := strings.Split(args.InstanceType, ".")
	platform := parts[0]
	preset := parts[1]

	// create the node groups
	createNodeGroupOperation, err := nebiusNodeGroupService.Create(ctx, &mk8s.CreateNodeGroupRequest{
		Metadata: &common.ResourceMetadata{
			Name:     args.Name,
			ParentId: cluster.CloudID,
			Labels: map[string]string{
				labelBrevRefID: args.RefID,
				labelCreatedBy: labelBrevCloudSDK,
			},
		},
		Spec: &mk8s.NodeGroupSpec{
			Size: &mk8s.NodeGroupSpec_Autoscaling{
				Autoscaling: &mk8s.NodeGroupAutoscalingSpec{
					MinNodeCount: int64(args.MinNodeCount),
					MaxNodeCount: int64(args.MaxNodeCount),
				},
			},
			Template: &mk8s.NodeTemplate{
				Resources: &mk8s.ResourcesSpec{
					Platform: platform,
					Size: &mk8s.ResourcesSpec_Preset{
						Preset: preset,
					},
				},
				GpuSettings: &mk8s.GpuSettings{
					DriversPreset: "cuda12",
				},
				BootDisk: &mk8s.DiskSpec{
					Type: mk8s.DiskSpec_NETWORK_SSD,
					Size: &mk8s.DiskSpec_SizeGibibytes{
						SizeGibibytes: int64(args.DiskSizeGiB),
					},
				},
			},
		},
	})
	if err != nil {
		return nil, err
	}

	_, err = createNodeGroupOperation.Wait(ctx)
	if err != nil {
		return nil, err
	}

	return &v1.CreateNodeGroupResponse{
		ClusterRefID: args.ClusterRefID,
		Name:         args.Name,
		RefID:        args.RefID,
	}, nil
}

func (c *NebiusClient) DeleteCluster(ctx context.Context, args v1.DeleteClusterArgs) error {
	// Fetch the cluster the user key will be added to
	cluster, err := c.GetCluster(ctx, v1.GetClusterArgs{
		RefID: args.ClusterRefID,
	})
	if err != nil {
		return fmt.Errorf("failed to get cluster: %w", err)
	}

	deleteClusterOperation, err := c.sdk.Services().MK8S().V1().Cluster().Delete(ctx, &mk8s.DeleteClusterRequest{
		Id: cluster.CloudID,
	})
	if err != nil {
		return fmt.Errorf("failed to delete cluster: %w", err)
	}

	_, err = deleteClusterOperation.Wait(ctx)
	if err != nil {
		return fmt.Errorf("failed to delete cluster: %w", err)
	}

	return nil
}

func (c *NebiusClient) newClusterClient(ctx context.Context, cluster *v1.Cluster) (*kubernetes.Clientset, error) {
	// Decode the cluster CA certificate
	clusterCACertificate, err := base64.StdEncoding.DecodeString(cluster.ClusterCACertificateBase64)
	if err != nil {
		return nil, fmt.Errorf("failed to decode cluster CA certificate: %w", err)
	}

	// Get a bearer token to authenticate to the cluster
	bearerToken, err := c.sdk.BearerToken(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get bearer token: %w", err)
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
		return nil, fmt.Errorf("failed to create clientset: %w", err)
	}

	return clientset, nil
}
