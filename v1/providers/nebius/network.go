package v1

import (
	"context"
	"fmt"

	"github.com/brevdev/cloud/internal/errors"
	v1 "github.com/brevdev/cloud/v1"

	common "github.com/nebius/gosdk/proto/nebius/common/v1"
	vpc "github.com/nebius/gosdk/proto/nebius/vpc/v1"
	nebiusVPC "github.com/nebius/gosdk/services/nebius/vpc/v1"
	grpcCodes "google.golang.org/grpc/codes"
	grpcStatus "google.golang.org/grpc/status"
)

const (
	labelBrevRefID      = "brev-ref-id"
	labelBrevVPCID      = "brev-vpc-id"
	labelBrevSubnetType = "brev-subnet-type"
	labelBrevCIDRBlock  = "brev-cidr-block"
	labelCreatedBy      = "CreatedBy"
	labelBrevCloudSDK   = "brev-cloud-sdk"
)

func (c *NebiusClient) CreateVPC(ctx context.Context, args v1.CreateVPCArgs) (*v1.VPC, error) {
	nebiusNetworkService := c.sdk.Services().VPC().V1().Network()
	nebiusSubnetService := c.sdk.Services().VPC().V1().Subnet()
	nebiusPoolService := c.sdk.Services().VPC().V1().Pool()

	// Create the network
	vpcID, err := createNetwork(ctx, nebiusNetworkService, nebiusPoolService, c.projectID, args)
	if err != nil {
		return nil, errors.WrapAndTrace(err)
	}

	// Create the subnets
	subnets := make([]*v1.Subnet, 0)
	for _, subnetArgs := range args.Subnets {
		subnet, err := createSubnet(ctx, nebiusSubnetService, c.projectID, vpcID, subnetArgs)
		if err != nil {
			return nil, errors.WrapAndTrace(err)
		}
		subnets = append(subnets, subnet)
	}

	return &v1.VPC{
		RefID:     args.RefID,
		Name:      args.Name,
		Location:  args.Location,
		CidrBlock: args.CidrBlock,
		Status:    v1.VPCStatusPending,
		ID:        v1.CloudProviderResourceID(vpcID),
		Subnets:   subnets,
	}, nil
}

func createNetwork(ctx context.Context, nebiusNetworkService nebiusVPC.NetworkService, nebiusPoolService nebiusVPC.PoolService, projectID string, args v1.CreateVPCArgs) (string, error) {
	// In Nebius, rather than creating a network with a CIDR, and subnets with slices of that CIDR, we instead first create a pool with
	// several specific CIDR blocks. These blocks will be intended to be used by subnets at the moment of their creation.
	// As we can add additional CIDR blocks to the pool later, we don't need to specify the entire network CIDR here.

	// Create the pool with the CIDR blocks for the subnets
	networkPoolCidrs := make([]*vpc.PoolCidr, 0)
	for _, subnet := range args.Subnets {
		networkPoolCidrs = append(networkPoolCidrs, &vpc.PoolCidr{Cidr: subnet.CidrBlock})
	}
	createPoolOperation, err := nebiusPoolService.Create(ctx, &vpc.CreatePoolRequest{
		Metadata: &common.ResourceMetadata{
			Name:     args.RefID,
			ParentId: projectID,
			Labels: map[string]string{
				labelBrevRefID: args.RefID,
				labelCreatedBy: labelBrevCloudSDK,
			},
		},
		Spec: &vpc.PoolSpec{
			Version:    vpc.IpVersion_IPV4,
			Visibility: vpc.IpVisibility_PRIVATE,
			Cidrs:      networkPoolCidrs,
		},
	})
	if err != nil {
		return "", errors.WrapAndTrace(err)
	}
	createPoolOperation, err = createPoolOperation.Wait(ctx)
	if err != nil {
		return "", errors.WrapAndTrace(err)
	}
	poolID := createPoolOperation.ResourceID()

	// Create the network with the pool
	createNetworkOperation, err := nebiusNetworkService.Create(ctx, &vpc.CreateNetworkRequest{
		Metadata: &common.ResourceMetadata{
			Name:     args.Name,
			ParentId: projectID,
			Labels: map[string]string{
				labelBrevRefID: args.RefID,
				labelCreatedBy: labelBrevCloudSDK,
			},
		},
		Spec: &vpc.NetworkSpec{
			Ipv4PrivatePools: &vpc.IPv4PrivateNetworkPools{
				Pools: []*vpc.NetworkPool{
					{Id: poolID},
				},
			},
		},
	})
	if err != nil {
		return "", errors.WrapAndTrace(err)
	}
	createNetworkOperation, err = createNetworkOperation.Wait(ctx)
	if err != nil {
		return "", errors.WrapAndTrace(err)
	}

	return createNetworkOperation.ResourceID(), nil
}

func createSubnet(ctx context.Context, nebiusSubnetService nebiusVPC.SubnetService, projectID string, networkID string, args v1.CreateSubnetArgs) (*v1.Subnet, error) {
	// In Nebius, the concept of "private" or "public" subnets is not a thing. Instead this concept is indirect -- subnets can be marked in such a
	// way as to allow for resources that are placed within them to allocate public IPs. This is controlled by the below "allowPublicIPAllocations" flag.
	var allowPublicIPAllocations bool
	if args.Type == v1.SubnetTypePublic {
		allowPublicIPAllocations = true
	} else {
		allowPublicIPAllocations = false
	}

	// Create the subnet, specifying the CIDR block (not the pool!) and the allowPublicIPAllocations flag.
	createSubnetOperation, err := nebiusSubnetService.Create(ctx, &vpc.CreateSubnetRequest{
		Metadata: &common.ResourceMetadata{
			Name:     fmt.Sprintf("%s-%s-%s", networkID, args.CidrBlock, args.Type),
			ParentId: projectID,
			Labels: map[string]string{
				labelBrevRefID:      fmt.Sprintf("%s-%s-%s", networkID, args.CidrBlock, args.Type),
				labelCreatedBy:      labelBrevCloudSDK,
				labelBrevSubnetType: string(args.Type),
				labelBrevVPCID:      networkID,
				labelBrevCIDRBlock:  args.CidrBlock,
			},
		},
		Spec: &vpc.SubnetSpec{
			NetworkId: networkID,
			Ipv4PrivatePools: &vpc.IPv4PrivateSubnetPools{
				UseNetworkPools: true,
			},
			Ipv4PublicPools: &vpc.IPv4PublicSubnetPools{
				UseNetworkPools: allowPublicIPAllocations,
			},
		},
	})
	if err != nil {
		return nil, errors.WrapAndTrace(err)
	}
	createSubnetOperation, err = createSubnetOperation.Wait(ctx)
	if err != nil {
		return nil, errors.WrapAndTrace(err)
	}

	return &v1.Subnet{
		ID:        v1.CloudProviderResourceID(createSubnetOperation.ResourceID()),
		CidrBlock: args.CidrBlock,
		Type:      args.Type,
		VPCID:     v1.CloudProviderResourceID(networkID),
		Name:      fmt.Sprintf("%s-%s-%s", networkID, args.CidrBlock, args.Type),
	}, nil
}

func (c *NebiusClient) GetVPC(ctx context.Context, args v1.GetVPCArgs) (*v1.VPC, error) {
	nebiusNetworkService := c.sdk.Services().VPC().V1().Network()
	nebiusSubnetService := c.sdk.Services().VPC().V1().Subnet()

	network, err := nebiusNetworkService.Get(ctx, &vpc.GetNetworkRequest{
		Id: string(args.ID),
	})
	if err != nil {
		if grpcStatus.Code(err) == grpcCodes.NotFound {
			return nil, v1.ErrResourceNotFound
		}
		return nil, errors.WrapAndTrace(err)
	}

	nebiusSubnets, err := nebiusSubnetService.ListByNetwork(ctx, &vpc.ListSubnetsByNetworkRequest{
		NetworkId: network.Metadata.Id,
	})
	if err != nil {
		return nil, errors.WrapAndTrace(err)
	}

	subnets := make([]*v1.Subnet, 0)
	for _, subnet := range nebiusSubnets.Items {
		subnets = append(subnets, &v1.Subnet{
			ID:        v1.CloudProviderResourceID(subnet.Metadata.Id),
			RefID:     subnet.Metadata.Labels[labelBrevRefID],
			Location:  subnet.Metadata.ParentId,
			CidrBlock: subnet.Metadata.Labels[labelBrevCIDRBlock],
			Type:      v1.SubnetType(subnet.Metadata.Labels[labelBrevSubnetType]),
			VPCID:     v1.CloudProviderResourceID(network.Metadata.Id),
			Name:      subnet.Metadata.Name,
		})
	}

	return &v1.VPC{
		ID:       v1.CloudProviderResourceID(network.Metadata.Id),
		RefID:    network.Metadata.Labels[labelBrevRefID],
		Location: network.Metadata.ParentId,
		Status:   parseNebiusNetworkStatus(network.Status),
		Subnets:  subnets,
	}, nil
}

func parseNebiusNetworkStatus(status *vpc.NetworkStatus) v1.VPCStatus {
	switch status.State {
	case vpc.NetworkStatus_CREATING:
		return v1.VPCStatusPending
	case vpc.NetworkStatus_READY:
		return v1.VPCStatusAvailable
	case vpc.NetworkStatus_DELETING:
		return v1.VPCStatusDeleting
	}
	return v1.VPCStatusUnknown
}

func (c *NebiusClient) DeleteVPC(ctx context.Context, args v1.DeleteVPCArgs) error {
	nebiusNetworkService := c.sdk.Services().VPC().V1().Network()
	nebiusPoolService := c.sdk.Services().VPC().V1().Pool()
	nebiusSubnetService := c.sdk.Services().VPC().V1().Subnet()

	// Find the network
	network, err := nebiusNetworkService.Get(ctx, &vpc.GetNetworkRequest{
		Id: string(args.ID),
	})
	if err != nil {
		return errors.WrapAndTrace(err)
	}

	// Find the network's subnets
	subnets, err := nebiusSubnetService.ListByNetwork(ctx, &vpc.ListSubnetsByNetworkRequest{
		NetworkId: network.Metadata.Id,
	})
	if err != nil {
		return errors.WrapAndTrace(err)
	}

	// Delete the subnets
	for _, subnet := range subnets.Items {
		deleteSubnetOperation, err := nebiusSubnetService.Delete(ctx, &vpc.DeleteSubnetRequest{
			Id: subnet.Metadata.Id,
		})
		if err != nil {
			return errors.WrapAndTrace(err)
		}
		deleteSubnetOperation, err = deleteSubnetOperation.Wait(ctx)
		if err != nil {
			return errors.WrapAndTrace(err)
		}
	}

	pool, err := nebiusPoolService.GetByName(ctx, &vpc.GetPoolByNameRequest{
		ParentId: network.Metadata.ParentId,
		Name:     network.Metadata.Name,
	})
	if err != nil {
		if grpcStatus.Code(err) != grpcCodes.NotFound {
			return errors.WrapAndTrace(err)
		}
		// Pool not found, continue
	}

	if pool != nil {
		// Remove pool from network
		updateNetworkOperation, err := nebiusNetworkService.Update(ctx, &vpc.UpdateNetworkRequest{
			Metadata: &common.ResourceMetadata{
				Name:     network.Metadata.Name,
				ParentId: network.Metadata.ParentId,
				Id:       network.Metadata.Id,
			},
			Spec: &vpc.NetworkSpec{
				Ipv4PrivatePools: &vpc.IPv4PrivateNetworkPools{
					Pools: []*vpc.NetworkPool{},
				},
			},
		})
		if err != nil {
			return errors.WrapAndTrace(err)
		}
		updateNetworkOperation, err = updateNetworkOperation.Wait(ctx)
		if err != nil {
			return errors.WrapAndTrace(err)
		}

		// Delete pool
		deletePoolOperation, err := nebiusPoolService.Delete(ctx, &vpc.DeletePoolRequest{
			Id: pool.Metadata.Id,
		})
		if err != nil {
			return errors.WrapAndTrace(err)
		}
		deletePoolOperation, err = deletePoolOperation.Wait(ctx)
		if err != nil {
			return errors.WrapAndTrace(err)
		}
	}

	// Delete the network
	deleteNetworkOperation, err := nebiusNetworkService.Delete(ctx, &vpc.DeleteNetworkRequest{
		Id: network.Metadata.Id,
	})
	if err != nil {
		return errors.WrapAndTrace(err)
	}

	deleteNetworkOperation, err = deleteNetworkOperation.Wait(ctx)
	if err != nil {
		return errors.WrapAndTrace(err)
	}

	return nil
}
