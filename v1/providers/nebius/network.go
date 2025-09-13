package v1

import (
	"context"
	"fmt"

	v1 "github.com/brevdev/cloud/v1"

	common "github.com/nebius/gosdk/proto/nebius/common/v1"
	vpc "github.com/nebius/gosdk/proto/nebius/vpc/v1"
	nebiusVPC "github.com/nebius/gosdk/services/nebius/vpc/v1"
	grpcCodes "google.golang.org/grpc/codes"
	grpcStatus "google.golang.org/grpc/status"
)

func (c *NebiusClient) CreateVPC(ctx context.Context, args v1.CreateVPCArgs) (*v1.VPC, error) {
	nebiusNetworkService := c.sdk.Services().VPC().V1().Network()
	nebiusSubnetService := c.sdk.Services().VPC().V1().Subnet()
	nebiusPoolService := c.sdk.Services().VPC().V1().Pool()

	// Create the VPC
	vpcID, err := createVPC(ctx, nebiusNetworkService, nebiusPoolService, c.projectID, args)
	if err != nil {
		return nil, err
	}

	publicSubnetArgs := filterSubnetArgs(args.Subnets, v1.SubnetTypePublic)
	// Create the subnets
	for _, publicSubnet := range publicSubnetArgs {
		_, err := createPublicSubnet(ctx, nebiusSubnetService, c.projectID, vpcID, publicSubnet)
		if err != nil {
			return nil, err
		}
	}

	return &v1.VPC{
		CloudID: vpcID,
	}, nil
}

func createVPC(ctx context.Context, nebiusNetworkClient nebiusVPC.NetworkService, nebiusPoolClient nebiusVPC.PoolService, projectID string, args v1.CreateVPCArgs) (string, error) {
	// TODO: I think this is actually wrong -- might need to:
	//  1. Create a pool
	//  2. Create a pool for each subnet, referencing the first pool as the parent
	//  3. Associate the parent pool with the network
	//  4. Associate the subnet pools with the network
	// Then on delete, reverse the process ("update" each resource with a blank pool)
	createPoolOperation, err := nebiusPoolClient.Create(ctx, &vpc.CreatePoolRequest{
		Metadata: &common.ResourceMetadata{
			Name:     args.RefID,
			ParentId: projectID,
			Labels: map[string]string{
				"brev-ref-id": args.RefID,
				"CreatedBy":   "brev-cloud-sdk",
			},
		},
		Spec: &vpc.PoolSpec{
			Version:    vpc.IpVersion_IPV4,
			Visibility: vpc.IpVisibility_PRIVATE,
			Cidrs: []*vpc.PoolCidr{
				{Cidr: args.CidrBlock},
			},
		},
	})
	if err != nil {
		return "", err
	}
	createPoolOperation, err = createPoolOperation.Wait(ctx)
	if err != nil {
		return "", err
	}

	createNetworkOperation, err := nebiusNetworkClient.Create(ctx, &vpc.CreateNetworkRequest{
		Metadata: &common.ResourceMetadata{
			Name:     args.Name,
			ParentId: projectID,
			Labels: map[string]string{
				"brev-ref-id": args.RefID,
				"CreatedBy":   "brev-cloud-sdk",
			},
		},
		Spec: &vpc.NetworkSpec{
			Ipv4PrivatePools: &vpc.IPv4PrivateNetworkPools{
				Pools: []*vpc.NetworkPool{
					{Id: createPoolOperation.ResourceID()},
				},
			},
		},
	})
	if err != nil {
		return "", err
	}

	createNetworkOperation, err = createNetworkOperation.Wait(ctx)
	if err != nil {
		return "", err
	}

	return createNetworkOperation.ResourceID(), nil
}

func createPublicSubnet(ctx context.Context, nebiusSubnetService nebiusVPC.SubnetService, projectID string, networkID string, args v1.CreateSubnetArgs) (*v1.Subnet, error) {
	createSubnetOperation, err := nebiusSubnetService.Create(ctx, &vpc.CreateSubnetRequest{
		Metadata: &common.ResourceMetadata{
			Name:     fmt.Sprintf("%s-%s-public", networkID, args.CidrBlock),
			ParentId: projectID,
			Labels: map[string]string{
				"CreatedBy": "brev-cloud-sdk",
			},
		},
		Spec: &vpc.SubnetSpec{
			NetworkId: networkID,
			Ipv4PrivatePools: &vpc.IPv4PrivateSubnetPools{
				Pools: []*vpc.SubnetPool{
					{Cidrs: []*vpc.SubnetCidr{
						{Cidr: args.CidrBlock},
					}},
				},
			},
		},
	})
	if err != nil {
		return nil, err
	}

	createSubnetOperation, err = createSubnetOperation.Wait(ctx)
	if err != nil {
		return nil, err
	}

	return &v1.Subnet{
		CloudID:   createSubnetOperation.ResourceID(),
		CidrBlock: args.CidrBlock,
		Type:      v1.SubnetTypePublic,
		VPCID:     networkID,
		Name:      fmt.Sprintf("%s-%s-public", networkID, args.CidrBlock),
	}, nil
}

func (c *NebiusClient) GetVPC(ctx context.Context, args v1.GetVPCArgs) (*v1.VPC, error) {
	nebiusClient := c.sdk.Services().VPC().V1().Network()

	vpc, err := nebiusClient.Get(ctx, &vpc.GetNetworkRequest{
		Id: args.CloudID,
	})
	if err != nil {
		return nil, err
	}

	return &v1.VPC{
		CloudID: vpc.Metadata.Id,
	}, nil
}

func (c *NebiusClient) DeleteVPC(ctx context.Context, args v1.DeleteVPCArgs) error {
	nebiusNetworkService := c.sdk.Services().VPC().V1().Network()
	nebiusPoolService := c.sdk.Services().VPC().V1().Pool()
	nebiusSubnetService := c.sdk.Services().VPC().V1().Subnet()

	network, err := nebiusNetworkService.GetByName(ctx, &vpc.GetNetworkByNameRequest{
		ParentId: c.projectID,
		Name:     args.VPC.RefID,
	})
	if err != nil {
		return err
	}

	pool, err := nebiusPoolService.GetByName(ctx, &vpc.GetPoolByNameRequest{
		ParentId: c.projectID,
		Name:     args.VPC.RefID,
	})
	if err != nil {
		if grpcStatus.Code(err) != grpcCodes.NotFound {
			return err
		}
		// Pool not found, continue
	}

	if pool != nil {
		// Remove pool from network
		updateNetworkOperation, err := nebiusNetworkService.Update(ctx, &vpc.UpdateNetworkRequest{
			Metadata: network.Metadata,
			Spec: &vpc.NetworkSpec{
				Ipv4PrivatePools: &vpc.IPv4PrivateNetworkPools{
					Pools: []*vpc.NetworkPool{},
				},
			},
		})
		if err != nil {
			return err
		}
		updateNetworkOperation, err = updateNetworkOperation.Wait(ctx)
		if err != nil {
			return err
		}

		// Delete pool
		deletePoolOperation, err := nebiusPoolService.Delete(ctx, &vpc.DeletePoolRequest{
			Id: pool.Metadata.Id,
		})
		if err != nil {
			return err
		}
		deletePoolOperation, err = deletePoolOperation.Wait(ctx)
		if err != nil {
			return err
		}
	}

	subnets, err := nebiusSubnetService.ListByNetwork(ctx, &vpc.ListSubnetsByNetworkRequest{
		NetworkId: network.Metadata.Id,
	})
	if err != nil {
		return err
	}

	for _, subnet := range subnets.Items {
		deleteSubnetOperation, err := nebiusSubnetService.Delete(ctx, &vpc.DeleteSubnetRequest{
			Id: subnet.Metadata.Id,
		})
		if err != nil {
			return err
		}
		deleteSubnetOperation, err = deleteSubnetOperation.Wait(ctx)
		if err != nil {
			return err
		}
	}

	deleteNetworkOperation, err := nebiusNetworkService.Delete(ctx, &vpc.DeleteNetworkRequest{
		Id: network.Metadata.Id,
	})
	if err != nil {
		return err
	}

	deleteNetworkOperation, err = deleteNetworkOperation.Wait(ctx)
	if err != nil {
		return err
	}

	return nil
}

// Helper function to filter subnet arguments by type
func filterSubnetArgs(subnets []v1.CreateSubnetArgs, subnetType v1.SubnetType) []v1.CreateSubnetArgs {
	filteredSubnets := make([]v1.CreateSubnetArgs, 0)
	for _, subnet := range subnets {
		if subnet.Type == subnetType {
			filteredSubnets = append(filteredSubnets, subnet)
		}
	}
	return filteredSubnets
}
