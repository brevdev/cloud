package validation

import (
	"context"
	"testing"
	"time"

	"github.com/brevdev/cloud/internal/ssh"
	v1 "github.com/brevdev/cloud/v1"
	"github.com/stretchr/testify/require"
)

type ProviderConfig struct {
	Location   string
	StableIDs  []v1.InstanceTypeID
	Credential v1.CloudCredential
}

func RunValidationSuite(t *testing.T, config ProviderConfig) {
	if testing.Short() {
		t.Skip("Skipping validation tests in short mode")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
	defer cancel()

	client, err := config.Credential.MakeClient(ctx, config.Location)
	if err != nil {
		t.Fatalf("Failed to create client for %s: %v", config.Credential.GetCloudProviderID(), err)
	}

	t.Run("ValidateGetLocations", func(t *testing.T) {
		err := v1.ValidateGetLocations(ctx, client)
		require.NoError(t, err, "ValidateGetLocations should pass")
	})

	t.Run("ValidateGetInstanceTypes", func(t *testing.T) {
		err := v1.ValidateGetInstanceTypes(ctx, client)
		require.NoError(t, err, "ValidateGetInstanceTypes should pass")
	})

	t.Run("ValidateRegionalInstanceTypes", func(t *testing.T) {
		err := v1.ValidateLocationalInstanceTypes(ctx, client)
		require.NoError(t, err, "ValidateRegionalInstanceTypes should pass")
	})

	t.Run("ValidateStableInstanceTypeIDs", func(t *testing.T) {
		err = v1.ValidateStableInstanceTypeIDs(ctx, client, config.StableIDs)
		require.NoError(t, err, "ValidateStableInstanceTypeIDs should pass")
	})
}

func RunInstanceLifecycleValidation(t *testing.T, config ProviderConfig) {
	if testing.Short() {
		t.Skip("Skipping validation tests in short mode")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Minute)
	defer cancel()

	client, err := config.Credential.MakeClient(ctx, config.Location)
	if err != nil {
		t.Fatalf("Failed to create client for %s: %v", config.Credential.GetCloudProviderID(), err)
	}
	capabilities, err := client.GetCapabilities(ctx)
	require.NoError(t, err)

	types, err := client.GetInstanceTypes(ctx, v1.GetInstanceTypeArgs{
		ArchitectureFilter: &v1.ArchitectureFilter{
			IncludeArchitectures: []v1.Architecture{v1.ArchitectureX86_64},
		},
	})
	require.NoError(t, err)
	require.NotEmpty(t, types, "Should have instance types")

	locations, err := client.GetLocations(ctx, v1.GetLocationsArgs{})
	require.NoError(t, err)
	require.NotEmpty(t, locations, "Should have locations")

	t.Run("ValidateCreateInstance", func(t *testing.T) {
		attrs := v1.CreateInstanceAttrs{}
		selectedType := v1.InstanceType{}
		for _, typ := range types {
			if typ.IsAvailable {
				attrs.InstanceType = typ.Type
				attrs.Location = typ.Location
				attrs.PublicKey = ssh.GetTestPublicKey()
				selectedType = typ
				break
			}
		}
		instance, err := v1.ValidateCreateInstance(ctx, client, attrs, selectedType)
		if err != nil {
			t.Fatalf("ValidateCreateInstance failed: %v", err)
		}
		require.NotNil(t, instance)

		defer func() {
			if instance != nil {
				_ = client.TerminateInstance(ctx, instance.CloudID)
			}
		}()

		t.Run("ValidateListCreatedInstance", func(t *testing.T) {
			err := v1.ValidateListCreatedInstance(ctx, client, instance)
			require.NoError(t, err, "ValidateListCreatedInstance should pass")
		})

		t.Run("ValidateSSHAccessible", func(t *testing.T) {
			err := v1.ValidateInstanceSSHAccessible(ctx, client, instance, ssh.GetTestPrivateKey())
			require.NoError(t, err, "ValidateSSHAccessible should pass")
		})

		instance, err = client.GetInstance(ctx, instance.CloudID)
		require.NoError(t, err)

		t.Run("ValidateInstanceImage", func(t *testing.T) {
			err := v1.ValidateInstanceImage(ctx, *instance, ssh.GetTestPrivateKey())
			require.NoError(t, err, "ValidateInstanceImage should pass")
		})

		t.Run("ValidateFirewallBlocksPort", func(t *testing.T) {
			err := v1.ValidateFirewallBlocksPort(ctx, client, instance, ssh.GetTestPrivateKey(), v1.DefaultFirewallTestPort)
			require.NoError(t, err, "ValidateFirewallBlocksPort should pass - non-allowed port should be blocked")
		})

		t.Run("ValidateDockerFirewallBlocksPort", func(t *testing.T) {
			err := v1.ValidateDockerFirewallBlocksPort(ctx, client, instance, ssh.GetTestPrivateKey(), v1.DefaultFirewallTestPort)
			require.NoError(t, err, "ValidateDockerFirewallBlocksPort should pass - docker port should be blocked by iptables")
		})

		t.Run("ValidateDockerFirewallAllowsEgress", func(t *testing.T) {
			err := v1.ValidateDockerFirewallAllowsEgress(ctx, client, instance, ssh.GetTestPrivateKey())
			require.NoError(t, err, "ValidateDockerFirewallAllowsEgress should pass - egress should be allowed")
		})

		t.Run("ValidateDockerFirewallAllowsContainerToContainerCommunication", func(t *testing.T) {
			err := v1.ValidateDockerFirewallAllowsContainerToContainerCommunication(ctx, client, instance, ssh.GetTestPrivateKey())
			require.NoError(t, err, "ValidateDockerFirewallAllowsContainerToContainerCommunication should pass - container to container communication should be allowed")
		})

		t.Run("ValidateMicroK8sFirewallAllowsEgress", func(t *testing.T) {
			err := v1.ValidateMicroK8sFirewallAllowsEgress(ctx, client, instance, ssh.GetTestPrivateKey())
			require.NoError(t, err, "ValidateMicroK8sFirewallAllowsEgress should pass - microk8s pod egress should be allowed")
		})
	
		t.Run("ValidateMicroK8sFirewallAllowsPodToPodCommunication", func(t *testing.T) {
			err := v1.ValidateMicroK8sFirewallAllowsPodToPodCommunication(ctx, client, instance, ssh.GetTestPrivateKey())
			require.NoError(t, err, "ValidateMicroK8sFirewallAllowsPodToPodCommunication should pass - microk8s pod to pod communication should be allowed")
		})

		if capabilities.IsCapable(v1.CapabilityStopStartInstance) && instance.Stoppable {
			t.Run("ValidateStopStartInstance", func(t *testing.T) {
				err := v1.ValidateStopStartInstance(ctx, client, instance)
				require.NoError(t, err, "ValidateStopStartInstance should pass")
			})
		}

		t.Run("ValidateTerminateInstance", func(t *testing.T) {
			err := v1.ValidateTerminateInstance(ctx, client, instance)
			require.NoError(t, err, "ValidateTerminateInstance should pass")
		})
	})
}

type NetworkValidationOpts struct {
	Name                  string
	RefID                 string
	CidrBlock             string
	PublicSubnetCidrBlock string
	Tags                  map[string]string
}

func RunNetworkValidation(t *testing.T, config ProviderConfig, opts NetworkValidationOpts) {
	if testing.Short() {
		t.Skip("Skipping validation tests in short mode")
	}

	// Set a default timeout of 15 minutes for the validation suite
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Minute)
	defer cancel()

	client, err := config.Credential.MakeClient(ctx, config.Location)
	if err != nil {
		t.Fatalf("Failed to create client for %s: %v", config.Credential.GetCloudProviderID(), err)
	}

	// Test #1: ValidateCreateVPC
	var vpcID v1.CloudProviderResourceID
	t.Run("ValidateCreateVPC", func(t *testing.T) {
		vpc, err := v1.ValidateCreateVPC(ctx, client, v1.CreateVPCArgs{
			Name:      opts.Name,
			RefID:     opts.RefID,
			CidrBlock: opts.CidrBlock,
			Subnets: []v1.CreateSubnetArgs{
				{CidrBlock: opts.PublicSubnetCidrBlock, Type: v1.SubnetTypePublic},
			},
			Tags: opts.Tags,
		})
		require.NoError(t, err, "ValidateCreateVPC should pass")
		vpcID = vpc.GetID()
	})

	// The VPC was created successfully -- create a defer function to delete the VPC if the tests fail
	deletionSucceeded := false
	defer func() {
		if !deletionSucceeded && vpcID != "" {
			t.Logf("Cleaning up VPC after failed tests: %s", vpcID)
			err = v1.ValidateDeleteVPC(ctx, client, v1.DeleteVPCArgs{
				ID: vpcID,
			})
			if err != nil {
				t.Fatalf("Failed to cleanup after validation of VPC: %v", err)
			}
		}
	}()

	// Test #2: ValidateGetVPC
	t.Run("ValidateGetVPC", func(t *testing.T) {
		vpc, err := v1.ValidateGetVPC(ctx, client, v1.GetVPCArgs{
			ID: vpcID,
		})
		require.NoError(t, err, "ValidateGetVPC should pass")
		require.NotNil(t, vpc)
	})

	// Test #3: WaitForVPCToBeAvailable
	t.Run("WaitForVPCToBeAvailable", func(t *testing.T) {
		err := WaitForResourcePredicate(ctx, WaitForResourcePredicateOpts[*v1.VPC]{
			GetResource: func() (*v1.VPC, error) {
				return client.GetVPC(ctx, v1.GetVPCArgs{ID: vpcID})
			},
			Predicate: func(vpc *v1.VPC) bool {
				return vpc.GetStatus() == v1.VPCStatusAvailable
			},
			Timeout:  5 * time.Minute,
			Interval: 5 * time.Second,
		})
		require.NoError(t, err, "WaitForVPCToBeAvailable should pass")
	})

	// Test #4: ValidateDeleteVPC
	t.Run("ValidateDeleteVPC", func(t *testing.T) {
		err := v1.ValidateDeleteVPC(ctx, client, v1.DeleteVPCArgs{
			ID: vpcID,
		})
		require.NoError(t, err, "ValidateDeleteVPC should pass")
		deletionSucceeded = true
	})

	// Test #5: WaitForVPCToBeDeleted
	t.Run("WaitForVPCToBeDeleted", func(t *testing.T) {
		err := WaitForResourcePredicate(ctx, WaitForResourcePredicateOpts[*v1.VPC]{
			GetResource: func() (*v1.VPC, error) {
				return client.GetVPC(ctx, v1.GetVPCArgs{ID: vpcID})
			},
			Predicate: func(_ *v1.VPC) bool {
				return false // continue until failure
			},
			Timeout:  5 * time.Minute,
			Interval: 5 * time.Second,
		})
		require.ErrorIs(t, err, v1.ErrResourceNotFound)
		deletionSucceeded = true
	})
}

type FirewallValidationOpts struct {
	// TestPort is the port to test firewall blocking on (should NOT be in allowed ingress)
	TestPort int
}

func RunFirewallValidation(t *testing.T, config ProviderConfig, opts FirewallValidationOpts) {
	if testing.Short() {
		t.Skip("Skipping validation tests in short mode")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Minute)
	defer cancel()

	client, err := config.Credential.MakeClient(ctx, config.Location)
	if err != nil {
		t.Fatalf("Failed to create client for %s: %v", config.Credential.GetCloudProviderID(), err)
	}

	types, err := client.GetInstanceTypes(ctx, v1.GetInstanceTypeArgs{
		ArchitectureFilter: &v1.ArchitectureFilter{
			IncludeArchitectures: []v1.Architecture{v1.ArchitectureX86_64},
		},
	})
	require.NoError(t, err)
	require.NotEmpty(t, types, "Should have instance types")

	// Find an available instance type
	attrs := v1.CreateInstanceAttrs{}
	selectedType := v1.InstanceType{}
	for _, typ := range types {
		if typ.IsAvailable {
			attrs.InstanceType = typ.Type
			attrs.Location = typ.Location
			attrs.PublicKey = ssh.GetTestPublicKey()
			selectedType = typ
			break
		}
	}
	require.NotEmpty(t, attrs.InstanceType, "Should find available instance type")

	// Create instance for firewall testing
	instance, err := v1.ValidateCreateInstance(ctx, client, attrs, selectedType)
	require.NoError(t, err, "ValidateCreateInstance should pass")
	require.NotNil(t, instance)

	defer func() {
		if instance != nil {
			_ = client.TerminateInstance(ctx, instance.CloudID)
		}
	}()

	// Wait for instance to be running and SSH accessible
	t.Run("ValidateSSHAccessible", func(t *testing.T) {
		err := v1.ValidateInstanceSSHAccessible(ctx, client, instance, ssh.GetTestPrivateKey())
		require.NoError(t, err, "ValidateSSHAccessible should pass")
	})

	// Refresh instance data
	instance, err = client.GetInstance(ctx, instance.CloudID)
	require.NoError(t, err)

	testPort := opts.TestPort
	if testPort == 0 {
		testPort = v1.DefaultFirewallTestPort
	}

	// Test that regular server on 0.0.0.0 is blocked
	t.Run("ValidateFirewallBlocksPort", func(t *testing.T) {
		err := v1.ValidateFirewallBlocksPort(ctx, client, instance, ssh.GetTestPrivateKey(), testPort)
		require.NoError(t, err, "ValidateFirewallBlocksPort should pass - port should be blocked")
	})

	t.Run("ValidateDockerFirewallBlocksPort", func(t *testing.T) {
		err := v1.ValidateDockerFirewallBlocksPort(ctx, client, instance, ssh.GetTestPrivateKey(), testPort)
		require.NoError(t, err, "ValidateDockerFirewallBlocksPort should pass - docker port should be blocked")
	})

	t.Run("ValidateDockerFirewallAllowsEgress", func(t *testing.T) {
		err := v1.ValidateDockerFirewallAllowsEgress(ctx, client, instance, ssh.GetTestPrivateKey())
		require.NoError(t, err, "ValidateDockerFirewallAllowsEgress should pass - egress should be allowed")
	})

	t.Run("ValidateDockerFirewallAllowsContainerToContainerCommunication", func(t *testing.T) {
		err := v1.ValidateDockerFirewallAllowsContainerToContainerCommunication(ctx, client, instance, ssh.GetTestPrivateKey())
		require.NoError(t, err, "ValidateDockerFirewallAllowsContainerToContainerCommunication should pass - container to container communication should be allowed")
	})

	t.Run("ValidateMicroK8sFirewallAllowsEgress", func(t *testing.T) {
		err := v1.ValidateMicroK8sFirewallAllowsEgress(ctx, client, instance, ssh.GetTestPrivateKey())
		require.NoError(t, err, "ValidateMicroK8sFirewallAllowsEgress should pass - microk8s pod egress should be allowed")
	})

	t.Run("ValidateMicroK8sFirewallAllowsPodToPodCommunication", func(t *testing.T) {
		err := v1.ValidateMicroK8sFirewallAllowsPodToPodCommunication(ctx, client, instance, ssh.GetTestPrivateKey())
		require.NoError(t, err, "ValidateMicroK8sFirewallAllowsPodToPodCommunication should pass - microk8s pod to pod communication should be allowed")
	})

	// Test that SSH port is accessible (sanity check)
	t.Run("ValidateSSHPortAccessible", func(t *testing.T) {
		err := v1.ValidateFirewallAllowsPort(ctx, client, instance, ssh.GetTestPrivateKey(), instance.SSHPort)
		require.NoError(t, err, "ValidateFirewallAllowsPort should pass for SSH port")
	})

	// Terminate instance
	t.Run("ValidateTerminateInstance", func(t *testing.T) {
		err := v1.ValidateTerminateInstance(ctx, client, instance)
		require.NoError(t, err, "ValidateTerminateInstance should pass")
	})
}

type KubernetesValidationOpts struct {
	Name              string
	RefID             string
	KubernetesVersion string
	Subnets           []KubernetesValidationSubnetOpts
	NodeGroupOpts     *KubernetesValidationNodeGroupOpts
	NetworkOpts       *KubernetesValidationNetworkOpts
	UserOpts          *KubernetesValidationUserOpts
	Tags              map[string]string
}

type KubernetesValidationNodeGroupOpts struct {
	Name         string
	RefID        string
	MinNodeCount int
	MaxNodeCount int
	InstanceType string
	DiskSize     v1.Bytes
}

type KubernetesValidationNetworkOpts struct {
	Name      string
	RefID     string
	CidrBlock string
	Subnets   []KubernetesValidationSubnetOpts
}

type KubernetesValidationSubnetOpts struct {
	Name       string
	RefID      string
	CidrBlock  string
	SubnetType v1.SubnetType
}

type KubernetesValidationUserOpts struct {
	Username     string
	Role         string
	RSAPEMBase64 string
}

func RunKubernetesValidation(t *testing.T, config ProviderConfig, opts KubernetesValidationOpts) { //nolint:funlen,gocyclo // This function is long but it is a validation suite
	if testing.Short() {
		t.Skip("Skipping validation tests in short mode")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Minute)
	defer cancel()

	client, err := config.Credential.MakeClient(ctx, config.Location)
	if err != nil {
		t.Fatalf("Failed to create client for %s: %v", config.Credential.GetCloudProviderID(), err)
	}

	if opts.NetworkOpts == nil {
		t.Fatalf("KubernetesValidationOpts.NetworkOpts is required")
	}

	subnets := []v1.CreateSubnetArgs{}
	for _, subnet := range opts.NetworkOpts.Subnets {
		subnets = append(subnets, v1.CreateSubnetArgs{
			RefID:     subnet.RefID,
			CidrBlock: subnet.CidrBlock,
			Type:      subnet.SubnetType,
		})
	}

	// Create the initial VPC
	vpc, err := v1.ValidateCreateVPC(ctx, client, v1.CreateVPCArgs{
		Name:      opts.NetworkOpts.Name,
		RefID:     opts.NetworkOpts.RefID,
		CidrBlock: opts.NetworkOpts.CidrBlock,
		Subnets:   subnets,
		Tags:      opts.Tags,
	})
	require.NoError(t, err, "ValidateCreateVPC should pass")

	// Wait for the VPC to be available
	err = WaitForResourcePredicate(ctx, WaitForResourcePredicateOpts[*v1.VPC]{
		GetResource: func() (*v1.VPC, error) {
			return client.GetVPC(ctx, v1.GetVPCArgs{ID: vpc.GetID()})
		},
		Predicate: func(vpc *v1.VPC) bool {
			return vpc.GetStatus() == v1.VPCStatusAvailable
		},
		Timeout:  5 * time.Minute,
		Interval: 5 * time.Second,
	})
	require.NoError(t, err, "WaitForVPCToBeAvailable should pass")
	t.Logf("VPC created: %s", vpc.GetID())

	// The VPC was created successfully -- create a defer function to delete the VPC if the tests fail
	defer func() {
		if vpc != nil {
			err = v1.ValidateDeleteVPC(ctx, client, v1.DeleteVPCArgs{
				ID: vpc.GetID(),
			})
			if err != nil {
				t.Fatalf("Failed to cleanup after validation of VPC: %v", err)
			}
			t.Logf("VPC deleted: %s", vpc.GetID())
		}
	}()

	// Map the input subnet ref IDs to their real subnets
	subnetRefIDs := make(map[string]*v1.Subnet)
	for _, subnet := range vpc.GetSubnets() {
		subnetRefIDs[subnet.GetRefID()] = subnet
	}

	// Convert the input subnet ref IDs to their real subnet IDs
	subnetIDs := []v1.CloudProviderResourceID{}
	for _, subnet := range opts.Subnets {
		subnetIDs = append(subnetIDs, subnetRefIDs[subnet.RefID].GetID())
	}

	// Test: Create Kubernetes Cluster
	var clusterID v1.CloudProviderResourceID
	t.Run("ValidateCreateKubernetesCluster", func(t *testing.T) {
		cluster, err := v1.ValidateCreateKubernetesCluster(ctx, client, v1.CreateClusterArgs{
			Name:              opts.Name,
			RefID:             opts.RefID,
			VPCID:             vpc.GetID(),
			SubnetIDs:         subnetIDs,
			KubernetesVersion: opts.KubernetesVersion,
			Tags:              opts.Tags,
		})
		require.NoError(t, err, "ValidateCreateKubernetesCluster should pass")
		require.NotNil(t, cluster)
		clusterID = cluster.GetID()
	})

	// The Kubernetes cluster was created successfully -- create a defer function to delete the Kubernetes cluster if the tests fail
	clusterDeletionSucceeded := false
	defer func() {
		if !clusterDeletionSucceeded && clusterID != "" {
			t.Logf("Cleaning up Kubernetes cluster after failed tests: %s", clusterID)
			err = v1.ValidateDeleteKubernetesCluster(ctx, client, v1.DeleteClusterArgs{
				ID: clusterID,
			})
			if err != nil {
				t.Fatalf("Failed to cleanup after validation of Kubernetes cluster: %v", err)
			}
		}
	}()

	// Test: Get Kubernetes Cluster
	t.Run("ValidateGetKubernetesCluster", func(t *testing.T) {
		cluster, err := v1.ValidateGetKubernetesCluster(ctx, client, v1.GetClusterArgs{
			ID: clusterID,
		})
		require.NoError(t, err, "ValidateGetKubernetesCluster should pass")
		require.NotNil(t, cluster)
	})

	// Test: WaitFor Kubernetes Cluster to Be Available
	t.Run("WaitForKubernetesClusterToBeAvailable", func(t *testing.T) {
		err := WaitForResourcePredicate(ctx, WaitForResourcePredicateOpts[*v1.Cluster]{
			GetResource: func() (*v1.Cluster, error) {
				return client.GetCluster(ctx, v1.GetClusterArgs{ID: clusterID})
			},
			Predicate: func(cluster *v1.Cluster) bool {
				return cluster.GetStatus() == v1.ClusterStatusAvailable
			},
			Timeout:  20 * time.Minute,
			Interval: 15 * time.Second,
		})
		require.NoError(t, err, "WaitForKubernetesClusterToBeAvailable should pass")
	})

	// Test: Get Kubernetes Cluster Credentials
	t.Run("ValidateGetKubernetesClusterCredentials", func(t *testing.T) {
		t.Skip("Under development")
		_, err := v1.ValidateSetKubernetesClusterUser(ctx, client, v1.SetClusterUserArgs{
			ClusterID:    clusterID,
			Username:     opts.UserOpts.Username,
			Role:         opts.UserOpts.Role,
			RSAPEMBase64: opts.UserOpts.RSAPEMBase64,
		})
		require.NoError(t, err, "ValidateGetKubernetesClusterCredentials should pass")
	})

	// Test: Create Kubernetes Node Group
	var nodeGroup v1.NodeGroup
	t.Run("ValidateCreateKubernetesNodeGroup", func(t *testing.T) {
		ng, err := v1.ValidateCreateKubernetesNodeGroup(ctx, client, v1.CreateNodeGroupArgs{
			ClusterID:    clusterID,
			Name:         opts.NodeGroupOpts.Name,
			RefID:        opts.NodeGroupOpts.RefID,
			MinNodeCount: opts.NodeGroupOpts.MinNodeCount,
			MaxNodeCount: opts.NodeGroupOpts.MaxNodeCount,
			InstanceType: opts.NodeGroupOpts.InstanceType,
			DiskSize:     opts.NodeGroupOpts.DiskSize,
			Tags:         opts.Tags,
		})
		require.NoError(t, err, "ValidateCreateKubernetesNodeGroup should pass")
		require.NotNil(t, ng)
		nodeGroup = *ng
	})

	// The node group was created successfully -- create a defer function to delete the node group if the tests fail
	nodeGroupDeletionSucceeded := false
	defer func() {
		if !nodeGroupDeletionSucceeded && nodeGroup.GetID() != "" {
			t.Logf("Cleaning up Kubernetes node group after failed tests: %s", nodeGroup.GetID())
			err = v1.ValidateDeleteKubernetesNodeGroup(ctx, client, v1.DeleteNodeGroupArgs{
				ID: nodeGroup.GetID(),
			})
			if err != nil {
				t.Fatalf("Failed to cleanup after validation of Kubernetes node group: %v", err)
			}
		}
	}()

	// Test: WaitFor Kubernetes Node Group to Be Available
	t.Run("WaitForKubernetesNodeGroupToBeAvailable", func(t *testing.T) {
		err := WaitForResourcePredicate(ctx, WaitForResourcePredicateOpts[*v1.NodeGroup]{
			GetResource: func() (*v1.NodeGroup, error) {
				return client.GetNodeGroup(ctx, v1.GetNodeGroupArgs{
					ClusterID: clusterID,
					ID:        nodeGroup.GetID(),
				})
			},
			Predicate: func(nodeGroup *v1.NodeGroup) bool {
				return nodeGroup.GetStatus() == v1.NodeGroupStatusAvailable
			},
			Timeout:  20 * time.Minute,
			Interval: 15 * time.Second,
		})
		require.NoError(t, err, "WaitForKubernetesNodeGroupToBeAvailable should pass")
	})

	// Test: Validate Cluster Node Groups matches the created node group
	t.Run("ValidateClusterNodeGroups", func(t *testing.T) {
		err := v1.ValidateClusterNodeGroups(ctx, client, v1.GetClusterArgs{ID: clusterID}, nodeGroup)
		require.NoError(t, err, "ValidateClusterNodeGroups should pass")
	})

	// Test: Modify Kubernetes Node Group
	t.Run("ValidateModifyKubernetesNodeGroup", func(t *testing.T) {
		err := v1.ValidateModifyKubernetesNodeGroup(ctx, client, v1.ModifyNodeGroupArgs{
			ClusterID:    clusterID,
			ID:           nodeGroup.GetID(),
			MinNodeCount: opts.NodeGroupOpts.MinNodeCount + 1,
			MaxNodeCount: opts.NodeGroupOpts.MaxNodeCount + 1,
		})
		require.NoError(t, err, "ValidateModifyKubernetesNodeGroup should pass")
	})

	// Test: WaitFor Kubernetes Node Group to Be Available
	t.Run("WaitForKubernetesNodeGroupToBeAvailable", func(t *testing.T) {
		err := WaitForResourcePredicate(ctx, WaitForResourcePredicateOpts[*v1.NodeGroup]{
			GetResource: func() (*v1.NodeGroup, error) {
				return client.GetNodeGroup(ctx, v1.GetNodeGroupArgs{
					ClusterID: clusterID,
					ID:        nodeGroup.GetID(),
				})
			},
			Predicate: func(nodeGroup *v1.NodeGroup) bool {
				return nodeGroup.GetStatus() == v1.NodeGroupStatusAvailable &&
					nodeGroup.GetMinNodeCount() == opts.NodeGroupOpts.MinNodeCount+1 &&
					nodeGroup.GetMaxNodeCount() == opts.NodeGroupOpts.MaxNodeCount+1
			},
			Timeout:  20 * time.Minute,
			Interval: 15 * time.Second,
		})
		require.NoError(t, err, "WaitForKubernetesNodeGroupToBeAvailable should pass")
	})

	// Test: Delete Kubernetes Node Group
	t.Run("ValidateDeleteKubernetesNodeGroup", func(t *testing.T) {
		err := v1.ValidateDeleteKubernetesNodeGroup(ctx, client, v1.DeleteNodeGroupArgs{
			ClusterID: clusterID,
			ID:        nodeGroup.GetID(),
		})
		require.NoError(t, err, "ValidateDeleteKubernetesNodeGroup should pass")
		nodeGroupDeletionSucceeded = true
	})

	// Test: WaitFor Kubernetes Node Group to Be Deleted
	t.Run("WaitForKubernetesNodeGroupToBeDeleted", func(t *testing.T) {
		err := WaitForResourcePredicate(ctx, WaitForResourcePredicateOpts[*v1.NodeGroup]{
			GetResource: func() (*v1.NodeGroup, error) {
				return client.GetNodeGroup(ctx, v1.GetNodeGroupArgs{
					ClusterID: clusterID,
					ID:        nodeGroup.GetID(),
				})
			},
			Predicate: func(_ *v1.NodeGroup) bool {
				return false // continue until failure
			},
			Timeout:  20 * time.Minute,
			Interval: 15 * time.Second,
		})
		require.ErrorIs(t, err, v1.ErrResourceNotFound)
	})

	// Test: Delete Kubernetes Cluster
	t.Run("ValidateDeleteKubernetesCluster", func(t *testing.T) {
		err := v1.ValidateDeleteKubernetesCluster(ctx, client, v1.DeleteClusterArgs{
			ID: clusterID,
		})
		require.NoError(t, err, "ValidateDeleteKubernetesCluster should pass")
	})

	// Test: WaitFor Kubernetes Cluster to Be Deleted
	t.Run("WaitForKubernetesClusterToBeDeleted", func(t *testing.T) {
		err := WaitForResourcePredicate(ctx, WaitForResourcePredicateOpts[*v1.Cluster]{
			GetResource: func() (*v1.Cluster, error) {
				return client.GetCluster(ctx, v1.GetClusterArgs{ID: clusterID})
			},
			Predicate: func(_ *v1.Cluster) bool {
				return false // continue until failure
			},
			Timeout:  20 * time.Minute,
			Interval: 15 * time.Second,
		})
		require.ErrorIs(t, err, v1.ErrResourceNotFound)
		clusterDeletionSucceeded = true
	})
}
