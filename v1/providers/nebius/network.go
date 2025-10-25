package v1

import (
	"context"
	"fmt"
	"net"

	nebiuscommon "github.com/nebius/gosdk/proto/nebius/common/v1"
	nebiusvpc "github.com/nebius/gosdk/proto/nebius/vpc/v1"
	nebiusvpcv1 "github.com/nebius/gosdk/services/nebius/vpc/v1"
	grpccodes "google.golang.org/grpc/codes"
	grpcstatus "google.golang.org/grpc/status"

	"github.com/brevdev/cloud/internal/errors"
	v1 "github.com/brevdev/cloud/v1"
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

	err := validateCreateVPCArgs(args)
	if err != nil {
		return nil, errors.WrapAndTrace(err)
	}

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

func validateCreateVPCArgs(args v1.CreateVPCArgs) error {
	if args.Name == "" {
		return fmt.Errorf("VPC name is required")
	}
	if args.RefID == "" {
		return fmt.Errorf("VPC refID is required")
	}
	if args.Location == "" {
		return fmt.Errorf("VPC location is required")
	}
	if args.CidrBlock == "" {
		return fmt.Errorf("VPC CIDR block is required")
	}
	if len(args.Subnets) == 0 {
		return fmt.Errorf("VPC subnets are required")
	}
	for _, subnet := range args.Subnets {
		if subnet.CidrBlock == "" {
			return fmt.Errorf("VPC subnet CIDR block is required")
		}
		if subnet.Type == "" {
			return fmt.Errorf("VPC subnet type is required")
		}
	}

	// Subnet CIDR blocks must be grreater than /24
	for _, subnet := range args.Subnets {
		larger, err := cidrBlockLargerThanMask(subnet.CidrBlock, 24)
		if err != nil {
			return errors.WrapAndTrace(err)
		}
		if !larger {
			return fmt.Errorf("VPC subnet CIDR block must be greater than /24: %s", subnet.CidrBlock)
		}
	}
	return nil
}

func cidrBlockLargerThanMask(cidrBlock string, mask int) (bool, error) {
	_, ipnet, err := net.ParseCIDR(cidrBlock)
	if err != nil {
		return false, errors.WrapAndTrace(err)
	}
	ones, _ := ipnet.Mask.Size()
	return ones < mask, nil
}

func createNetwork(ctx context.Context, nebiusNetworkService nebiusvpcv1.NetworkService, nebiusPoolService nebiusvpcv1.PoolService, projectID string, args v1.CreateVPCArgs) (string, error) {
	// In Nebius, rather than creating a network with a CIDR, and subnets with slices of that CIDR, we instead first create a pool with
	// several specific CIDR blocks. These blocks will be intended to be used by subnets at the moment of their creation.
	// As we can add additional CIDR blocks to the pool later, we don't need to specify the entire network CIDR here.

	labels := make(map[string]string)
	for key, value := range args.Tags {
		labels[key] = value
	}

	// Add the required labels
	labels[labelBrevRefID] = args.RefID
	labels[labelCreatedBy] = labelBrevCloudSDK

	// Create the pool with the CIDR blocks for the subnets
	networkPoolCidrs := make([]*nebiusvpc.PoolCidr, 0)
	for _, subnet := range args.Subnets {
		networkPoolCidrs = append(networkPoolCidrs, &nebiusvpc.PoolCidr{Cidr: subnet.CidrBlock})
	}
	createPoolOperation, err := nebiusPoolService.Create(ctx, &nebiusvpc.CreatePoolRequest{
		Metadata: &nebiuscommon.ResourceMetadata{
			Name:     args.RefID,
			ParentId: projectID,
			Labels:   labels,
		},
		Spec: &nebiusvpc.PoolSpec{
			Version:    nebiusvpc.IpVersion_IPV4,
			Visibility: nebiusvpc.IpVisibility_PRIVATE,
			Cidrs:      networkPoolCidrs,
		},
	})
	if err != nil {
		return "", errors.WrapAndTrace(err)
	}

	// Here we must wait for the pool to be created, as otherwise we cannot proceed to create the network.
	createPoolOperation, err = createPoolOperation.Wait(ctx)
	if err != nil {
		return "", errors.WrapAndTrace(err)
	}
	poolID := createPoolOperation.ResourceID()

	// Create the network with the pool
	createNetworkOperation, err := nebiusNetworkService.Create(ctx, &nebiusvpc.CreateNetworkRequest{
		Metadata: &nebiuscommon.ResourceMetadata{
			Name:     args.Name,
			ParentId: projectID,
			Labels:   labels,
		},
		Spec: &nebiusvpc.NetworkSpec{
			Ipv4PrivatePools: &nebiusvpc.IPv4PrivateNetworkPools{
				Pools: []*nebiusvpc.NetworkPool{
					{Id: poolID},
				},
			},
		},
	})
	if err != nil {
		return "", errors.WrapAndTrace(err)
	}

	return createNetworkOperation.ResourceID(), nil
}

func createSubnet(ctx context.Context, nebiusSubnetService nebiusvpcv1.SubnetService, projectID string, networkID string, args v1.CreateSubnetArgs) (*v1.Subnet, error) {
	// In Nebius, the concept of "private" or "public" subnets is not a thing. Instead this concept is indirect -- subnets can be marked in such a
	// way as to allow for resources that are placed within them to allocate public IPs. This is controlled by the below "allowPublicIPAllocations" flag.

	var allowPublicIPAllocations bool
	if args.Type == v1.SubnetTypePublic {
		allowPublicIPAllocations = true
	} else {
		allowPublicIPAllocations = false
	}

	labels := make(map[string]string)
	for key, value := range args.Tags {
		labels[key] = value
	}

	// Add the required labels
	labels[labelBrevRefID] = fmt.Sprintf("%s-%s-%s", networkID, args.CidrBlock, args.Type)
	labels[labelCreatedBy] = labelBrevCloudSDK
	labels[labelBrevSubnetType] = string(args.Type)
	labels[labelBrevVPCID] = networkID
	labels[labelBrevCIDRBlock] = args.CidrBlock

	// Create the subnet, specifying the CIDR block (not the pool!) and the allowPublicIPAllocations flag.
	createSubnetOperation, err := nebiusSubnetService.Create(ctx, &nebiusvpc.CreateSubnetRequest{
		Metadata: &nebiuscommon.ResourceMetadata{
			Name:     fmt.Sprintf("%s-%s-%s", networkID, args.CidrBlock, args.Type),
			ParentId: projectID,
			Labels:   labels,
		},
		Spec: &nebiusvpc.SubnetSpec{
			NetworkId: networkID,
			Ipv4PrivatePools: &nebiusvpc.IPv4PrivateSubnetPools{
				UseNetworkPools: true,
			},
			Ipv4PublicPools: &nebiusvpc.IPv4PublicSubnetPools{
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
		Tags:      args.Tags,
	}, nil
}

func (c *NebiusClient) GetVPC(ctx context.Context, args v1.GetVPCArgs) (*v1.VPC, error) {
	nebiusNetworkService := c.sdk.Services().VPC().V1().Network()
	nebiusSubnetService := c.sdk.Services().VPC().V1().Subnet()

	network, err := nebiusNetworkService.Get(ctx, &nebiusvpc.GetNetworkRequest{
		Id: string(args.ID),
	})
	if err != nil {
		if grpcstatus.Code(err) == grpccodes.NotFound {
			return nil, v1.ErrResourceNotFound
		}
		return nil, errors.WrapAndTrace(err)
	}

	nebiusSubnets, err := nebiusSubnetService.ListByNetwork(ctx, &nebiusvpc.ListSubnetsByNetworkRequest{
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
			Tags:      v1.Tags(subnet.Metadata.Labels),
		})
	}

	return &v1.VPC{
		ID:       v1.CloudProviderResourceID(network.Metadata.Id),
		RefID:    network.Metadata.Labels[labelBrevRefID],
		Location: network.Metadata.ParentId,
		Status:   parseNebiusNetworkStatus(network.Status),
		Subnets:  subnets,
		Tags:     v1.Tags(network.Metadata.Labels),
	}, nil
}

func parseNebiusNetworkStatus(status *nebiusvpc.NetworkStatus) v1.VPCStatus {
	switch status.State {
	case nebiusvpc.NetworkStatus_CREATING:
		return v1.VPCStatusPending
	case nebiusvpc.NetworkStatus_READY:
		return v1.VPCStatusAvailable
	case nebiusvpc.NetworkStatus_DELETING:
		return v1.VPCStatusDeleting
	}
	return v1.VPCStatusUnknown
}

func (c *NebiusClient) DeleteVPC(ctx context.Context, args v1.DeleteVPCArgs) error {
	nebiusNetworkService := c.sdk.Services().VPC().V1().Network()
	nebiusPoolService := c.sdk.Services().VPC().V1().Pool()
	nebiusSubnetService := c.sdk.Services().VPC().V1().Subnet()

	// Find the network
	network, err := nebiusNetworkService.Get(ctx, &nebiusvpc.GetNetworkRequest{
		Id: string(args.ID),
	})
	if err != nil {
		return errors.WrapAndTrace(err)
	}

	// Find the network's subnets
	subnets, err := nebiusSubnetService.ListByNetwork(ctx, &nebiusvpc.ListSubnetsByNetworkRequest{
		NetworkId: network.Metadata.Id,
	})
	if err != nil {
		return errors.WrapAndTrace(err)
	}

	// Delete the subnets
	for _, subnet := range subnets.Items {
		_, err := nebiusSubnetService.Delete(ctx, &nebiusvpc.DeleteSubnetRequest{
			Id: subnet.Metadata.Id,
		})
		if err != nil {
			return errors.WrapAndTrace(err)
		}
	}

	pool, err := nebiusPoolService.GetByName(ctx, &nebiusvpc.GetPoolByNameRequest{
		ParentId: network.Metadata.ParentId,
		Name:     network.Metadata.Name,
	})
	if err != nil {
		if grpcstatus.Code(err) != grpccodes.NotFound {
			return errors.WrapAndTrace(err)
		}
		// Pool not found, continue
	}

	if pool != nil {
		// Remove pool from network
		updateNetworkOperation, err := nebiusNetworkService.Update(ctx, &nebiusvpc.UpdateNetworkRequest{
			Metadata: &nebiuscommon.ResourceMetadata{
				Name:     network.Metadata.Name,
				ParentId: network.Metadata.ParentId,
				Id:       network.Metadata.Id,
			},
			Spec: &nebiusvpc.NetworkSpec{
				Ipv4PrivatePools: &nebiusvpc.IPv4PrivateNetworkPools{
					Pools: []*nebiusvpc.NetworkPool{},
				},
			},
		})
		if err != nil {
			return errors.WrapAndTrace(err)
		}

		// Here we must wait for the network to be updated, as otherwise we cannot proceed to delete the pool.
		updateNetworkOperation, err = updateNetworkOperation.Wait(ctx)
		if err != nil {
			return errors.WrapAndTrace(err)
		}

		// Delete pool
		deletePoolOperation, err := nebiusPoolService.Delete(ctx, &nebiusvpc.DeletePoolRequest{
			Id: pool.Metadata.Id,
		})
		if err != nil {
			return errors.WrapAndTrace(err)
		}

		// Here we must wait for the pool to be deleted, as otherwise we cannot proceed to delete the network.
		deletePoolOperation, err = deletePoolOperation.Wait(ctx)
		if err != nil {
			return errors.WrapAndTrace(err)
		}
	}

	// Delete the network
	_, err = nebiusNetworkService.Delete(ctx, &nebiusvpc.DeleteNetworkRequest{
		Id: network.Metadata.Id,
	})
	if err != nil {
		return errors.WrapAndTrace(err)
	}

	return nil
}
