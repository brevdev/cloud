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

	types, err := client.GetInstanceTypes(ctx, v1.GetInstanceTypeArgs{})
	require.NoError(t, err)
	require.NotEmpty(t, types, "Should have instance types")

	locations, err := client.GetLocations(ctx, v1.GetLocationsArgs{})
	require.NoError(t, err)
	require.NotEmpty(t, locations, "Should have locations")

	t.Run("ValidateCreateInstance", func(t *testing.T) {
		attrs := v1.CreateInstanceAttrs{}
		for _, typ := range types {
			if typ.IsAvailable {
				attrs.InstanceType = typ.Type
				attrs.Location = typ.Location
				attrs.PublicKey = ssh.GetTestPublicKey()
				break
			}
		}
		instance, err := v1.ValidateCreateInstance(ctx, client, attrs)
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
	Location              string
	CidrBlock             string
	PublicSubnetCidrBlock string
}

func RunNetworkValidation(t *testing.T, config ProviderConfig, opts NetworkValidationOpts) {
	if testing.Short() {
		t.Skip("Skipping validation tests in short mode")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Minute)
	defer cancel()

	client, err := config.Credential.MakeClient(ctx, config.Location)
	if err != nil {
		t.Fatalf("Failed to create client for %s: %v", config.Credential.GetCloudProviderID(), err)
	}

	var vpcID v1.CloudProviderResourceID
	t.Run("ValidateCreateVPC", func(t *testing.T) {
		vpc, err := v1.ValidateCreateVPC(ctx, client, v1.CreateVPCArgs{
			Name:      opts.Name,
			RefID:     opts.RefID,
			Location:  opts.Location,
			CidrBlock: opts.CidrBlock,
			Subnets: []v1.CreateSubnetArgs{
				{CidrBlock: opts.PublicSubnetCidrBlock, Type: v1.SubnetTypePublic},
			},
		})
		require.NoError(t, err, "ValidateCreateVPC should pass")
		vpcID = vpc.ID
	})

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

	t.Run("ValidateGetVPC", func(t *testing.T) {
		vpc, err := v1.ValidateGetVPC(ctx, client, v1.GetVPCArgs{
			ID: vpcID,
		})
		require.NoError(t, err, "ValidateGetVPC should pass")
		require.NotNil(t, vpc)
	})

	t.Run("WaitForVPCToBeAvailable", func(t *testing.T) {
		err := v1.WaitForVPCPredicate(ctx, client, v1.GetVPCArgs{ID: vpcID},
			v1.WaitForVPCPredicateOpts{
				Predicate: func(vpc *v1.VPC) bool {
					return vpc.Status == v1.VPCStatusAvailable
				},
				Timeout:  5 * time.Minute,
				Interval: 5 * time.Second,
			},
		)
		require.NoError(t, err, "WaitForVPCToBeAvailable should pass")
	})

	t.Run("ValidateDeleteVPC", func(t *testing.T) {
		err := v1.ValidateDeleteVPC(ctx, client, v1.DeleteVPCArgs{
			ID: vpcID,
		})
		require.NoError(t, err, "ValidateDeleteVPC should pass")
		deletionSucceeded = true
	})

	t.Run("ValidateVPCNotFound", func(t *testing.T) {
		vpc, err := v1.ValidateGetVPC(ctx, client, v1.GetVPCArgs{
			ID: vpcID,
		})
		require.Nil(t, vpc)
		require.ErrorIs(t, err, v1.ErrResourceNotFound)
	})
}

type KubernetesValidationOpts struct {
	Name              string
	RefID             string
	KubernetesVersion string
	NetworkOpts       *KubernetesValidationNetworkOpts
}

type KubernetesValidationNetworkOpts struct {
	Name                   string
	RefID                  string
	Location               string
	CidrBlock              string
	PublicSubnetCidrBlock  string
	PrivateSubnetCidrBlock string
}

func RunKubernetesValidation(t *testing.T, config ProviderConfig, opts KubernetesValidationOpts) {
	if testing.Short() {
		t.Skip("Skipping validation tests in short mode")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Minute)
	defer cancel()

	client, err := config.Credential.MakeClient(ctx, config.Location)
	if err != nil {
		t.Fatalf("Failed to create client for %s: %v", config.Credential.GetCloudProviderID(), err)
	}

	if opts.NetworkOpts == nil {
		t.Fatalf("KubernetesValidationOpts.NetworkOpts is required")
	}

	// Create the initial VPC
	vpc, err := v1.ValidateCreateVPC(ctx, client, v1.CreateVPCArgs{
		Name:      opts.NetworkOpts.Name,
		RefID:     opts.NetworkOpts.RefID,
		Location:  opts.NetworkOpts.Location,
		CidrBlock: opts.NetworkOpts.CidrBlock,
		Subnets: []v1.CreateSubnetArgs{
			{CidrBlock: opts.NetworkOpts.PublicSubnetCidrBlock, Type: v1.SubnetTypePublic},
			{CidrBlock: opts.NetworkOpts.PrivateSubnetCidrBlock, Type: v1.SubnetTypePrivate},
		},
	})
	require.NoError(t, err, "ValidateCreateVPC should pass")

	// Wait for the VPC to be available
	err = v1.WaitForVPCPredicate(ctx, client, v1.GetVPCArgs{ID: vpc.ID},
		v1.WaitForVPCPredicateOpts{
			Predicate: func(vpc *v1.VPC) bool {
				return vpc.Status == v1.VPCStatusAvailable
			},
			Timeout:  5 * time.Minute,
			Interval: 5 * time.Second,
		},
	)
	require.NoError(t, err, "WaitForVPCToBeAvailable should pass")

	defer func() {
		// Clean up the VPC if it was created
		if vpc != nil {
			err = v1.ValidateDeleteVPC(ctx, client, v1.DeleteVPCArgs{
				ID: vpc.ID,
			})
			if err != nil {
				t.Fatalf("Failed to cleanup after validation of VPC: %v", err)
			}
		}
	}()

	var clusterID v1.CloudProviderResourceID
	t.Run("ValidateCreateKubernetesCluster", func(t *testing.T) {
		cluster, err := v1.ValidateCreateKubernetesCluster(ctx, client, v1.CreateClusterArgs{
			Name:              opts.Name,
			RefID:             opts.RefID,
			VPCID:             vpc.ID,
			SubnetIDs:         []v1.CloudProviderResourceID{vpc.Subnets[0].ID},
			KubernetesVersion: opts.KubernetesVersion,
			Location:          opts.NetworkOpts.Location,
		})
		require.NoError(t, err, "ValidateCreateKubernetesCluster should pass")
		require.NotNil(t, cluster)
		clusterID = cluster.ID
	})

	t.Run("ValidateGetKubernetesCluster", func(t *testing.T) {
		cluster, err := v1.ValidateGetKubernetesCluster(ctx, client, v1.GetClusterArgs{
			ID: clusterID,
		})
		require.NoError(t, err, "ValidateGetKubernetesCluster should pass")
		require.NotNil(t, cluster)
	})

	t.Run("WaitForKubernetesClusterToBeAvailable", func(t *testing.T) {
		err := v1.WaitForKubernetesClusterPredicate(ctx, client, v1.GetClusterArgs{ID: clusterID},
			v1.WaitForKubernetesClusterPredicateOpts{
				Predicate: func(cluster *v1.Cluster) bool {
					return cluster.Status == v1.ClusterStatusAvailable
				},
				Timeout:  20 * time.Minute,
				Interval: 15 * time.Second,
			},
		)
		require.NoError(t, err, "WaitForKubernetesClusterToBeAvailable should pass")
	})

	// t.Run("ValidateGetKubernetesClusterCredentials", func(t *testing.T) {
	// 	err := v1.ValidateGetKubernetesClusterCredentials(ctx, client, v1.GetClusterArgs{
	// 		ID: v1.CloudProviderResourceID("test-cluster"),
	// 	})
	// 	require.NoError(t, err, "ValidateGetKubernetesClusterCredentials should pass")
	// })

	// t.Run("ValidateCreateKubernetesNodeGroup", func(t *testing.T) {
	// 	err := v1.ValidateCreateKubernetesNodeGroup(ctx, client, v1.CreateNodeGroupArgs{
	// 		ClusterID:    v1.CloudProviderResourceID("test-cluster"),
	// 		Name:         "test-node-group",
	// 		RefID:        "test-node-group",
	// 		MinNodeCount: 1,
	// 		MaxNodeCount: 1,
	// 		InstanceType: "test-instance-type",
	// 		DiskSizeGiB:  100,
	// 	})
	// 	require.NoError(t, err, "ValidateCreateKubernetesNodeGroup should pass")
	// })

	// t.Run("ValidateDeleteKubernetesNodeGroup", func(t *testing.T) {
	// 	err := v1.ValidateDeleteKubernetesNodeGroup(ctx, client, v1.DeleteNodeGroupArgs{
	// 		ID: v1.CloudProviderResourceID("test-node-group"),
	// 	})
	// 	require.NoError(t, err, "ValidateDeleteKubernetesNodeGroup should pass")
	// })

	t.Run("ValidateDeleteKubernetesCluster", func(t *testing.T) {
		err := v1.ValidateDeleteKubernetesCluster(ctx, client, v1.DeleteClusterArgs{
			ID: clusterID,
		})
		require.NoError(t, err, "ValidateDeleteKubernetesCluster should pass")
	})

	t.Run("WaitForKubernetesClusterToBeDeleted", func(t *testing.T) {
		err := v1.WaitForKubernetesClusterPredicate(ctx, client, v1.GetClusterArgs{ID: clusterID},
			v1.WaitForKubernetesClusterPredicateOpts{
				Predicate: func(_ *v1.Cluster) bool {
					return false // continue until failure
				},
				Timeout:  5 * time.Minute,
				Interval: 5 * time.Second,
			},
		)
		require.ErrorIs(t, err, v1.ErrResourceNotFound)
	})
}
