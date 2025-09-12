package v1

import (
	"context"
	"fmt"
	"time"

	v1 "github.com/brevdev/cloud/v1"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

func (c *AWSClient) CreateVPC(ctx context.Context, args v1.CreateVPCArgs) (*v1.VPC, error) {
	awsVPC, err := c.createVPC(ctx, args)
	if err != nil {
		return nil, err
	}

	return &v1.VPC{
		RefID:          args.RefID,
		Name:           args.Name,
		Location:       args.Location,
		CloudCredRefID: c.GetReferenceID(),
		Provider:       CloudProviderID,
		Cloud:          CloudProviderID,
		CloudID:        *awsVPC.VpcId,
		CidrBlock:      *awsVPC.CidrBlock,
	}, nil
}

func (c *AWSClient) GetVPC(ctx context.Context, args v1.GetVPCArgs) (*v1.VPC, error) {
	awsVPC, err := c.getVPC(ctx, args)
	if err != nil {
		return nil, err
	}

	nameTag := ""
	refIDTag := ""
	for _, tag := range awsVPC.Tags {
		if *tag.Key == "Name" {
			nameTag = *tag.Value
		}
		if *tag.Key == "brev-ref-id" {
			refIDTag = *tag.Value
		}
	}

	return &v1.VPC{
		RefID:          refIDTag,
		Name:           nameTag,
		Location:       args.Location,
		CloudID:        *awsVPC.VpcId,
		CloudCredRefID: c.GetReferenceID(),
		Provider:       CloudProviderID,
		Cloud:          CloudProviderID,
		CidrBlock:      *awsVPC.CidrBlock,
	}, nil
}

func (c *AWSClient) DeleteVPC(ctx context.Context, args v1.DeleteVPCArgs) error {
	err := c.deleteVPC(ctx, args)
	if err != nil {
		return err
	}
	return nil
}

func (c *AWSClient) CreatePublicSubnet(ctx context.Context, args v1.CreatePublicSubnetArgs) (*v1.Subnet, error) {
	awsSubnet, err := c.createPublicSubnet(ctx, args)
	if err != nil {
		return nil, err
	}

	return &v1.Subnet{
		RefID:     *awsSubnet.SubnetId,
		Location:  args.Location,
		Name:      *awsSubnet.Tags[0].Value,
		VPCID:     *awsSubnet.VpcId,
		CloudID:   *awsSubnet.SubnetId,
		CidrBlock: *awsSubnet.CidrBlock,
	}, nil
}

func (c *AWSClient) CreatePrivateSubnet(ctx context.Context, args v1.CreatePrivateSubnetArgs) (*v1.Subnet, error) {
	awsSubnet, err := c.createPrivateSubnet(ctx, args)
	if err != nil {
		return nil, err
	}

	return &v1.Subnet{
		RefID:     *awsSubnet.SubnetId,
		Location:  args.Location,
		Name:      *awsSubnet.Tags[0].Value,
		VPCID:     *awsSubnet.VpcId,
		CloudID:   *awsSubnet.SubnetId,
		CidrBlock: *awsSubnet.CidrBlock,
	}, nil
}

func (c *AWSClient) DeleteSubnet(ctx context.Context, args v1.DeleteSubnetArgs) error {
	err := c.deleteSubnet(ctx, args)
	if err != nil {
		return err
	}
	return nil
}

func (c *AWSClient) createVPC(ctx context.Context, args v1.CreateVPCArgs) (*types.Vpc, error) {
	// Create the AWS client in the specified region
	awsClient := ec2.NewFromConfig(c.awsConfig, func(o *ec2.Options) {
		o.Region = args.Location
	})

	// Create the VPC and associate relevant tag data
	createVPCOutput, err := awsClient.CreateVpc(ctx, &ec2.CreateVpcInput{
		CidrBlock: &args.CidrBlock,
		TagSpecifications: []types.TagSpecification{
			{
				ResourceType: types.ResourceTypeVpc,
				Tags: makeTags(map[string]string{
					"Name":        args.Name,
					"brev-ref-id": args.RefID,
				}),
			},
		},
	})
	if err != nil {
		return nil, err
	}

	// Enable DNS support for the VPC
	_, err = awsClient.ModifyVpcAttribute(ctx, &ec2.ModifyVpcAttributeInput{
		VpcId: createVPCOutput.Vpc.VpcId,
		EnableDnsSupport: &types.AttributeBooleanValue{
			Value: aws.Bool(true),
		},
	})
	if err != nil {
		return nil, err
	}

	// Enable DNS hostnames for the VPC
	_, err = awsClient.ModifyVpcAttribute(ctx, &ec2.ModifyVpcAttributeInput{
		VpcId: createVPCOutput.Vpc.VpcId,
		EnableDnsHostnames: &types.AttributeBooleanValue{
			Value: aws.Bool(true),
		},
	})
	if err != nil {
		return nil, err
	}

	// Create an Internet Gateway for the VPC
	createInternetGatewayOutput, err := awsClient.CreateInternetGateway(ctx, &ec2.CreateInternetGatewayInput{
		TagSpecifications: []types.TagSpecification{
			{
				ResourceType: types.ResourceTypeInternetGateway,
				Tags: makeTags(map[string]string{
					"Name":        args.Name,
					"brev-ref-id": args.RefID,
				}),
			},
		},
	})
	if err != nil {
		return nil, err
	}

	// Attach the Internet Gateway to the VPC
	_, err = awsClient.AttachInternetGateway(ctx, &ec2.AttachInternetGatewayInput{
		InternetGatewayId: createInternetGatewayOutput.InternetGateway.InternetGatewayId,
		VpcId:             createVPCOutput.Vpc.VpcId,
	})
	if err != nil {
		return nil, err
	}

	return createVPCOutput.Vpc, nil
}

func (c *AWSClient) getVPC(ctx context.Context, args v1.GetVPCArgs) (*types.Vpc, error) {
	awsClient := ec2.NewFromConfig(c.awsConfig, func(o *ec2.Options) {
		o.Region = args.Location
	})

	var describeVPCsOutput *ec2.DescribeVpcsOutput

	if args.CloudID != "" {
		var err error
		describeVPCsOutput, err = awsClient.DescribeVpcs(ctx, &ec2.DescribeVpcsInput{
			VpcIds: []string{args.CloudID},
		})
		if err != nil {
			return nil, err
		}
	} else if args.RefID != "" {
		var err error
		describeVPCsOutput, err = awsClient.DescribeVpcs(ctx, &ec2.DescribeVpcsInput{
			Filters: []types.Filter{
				{
					Name:   aws.String("tag:brev-ref-id"),
					Values: []string{args.RefID},
				},
			},
		})
		if err != nil {
			return nil, err
		}
	}

	if len(describeVPCsOutput.Vpcs) == 0 {
		return nil, nil
	}

	return &describeVPCsOutput.Vpcs[0], nil
}

func (c *AWSClient) deleteVPC(ctx context.Context, args v1.DeleteVPCArgs) error {
	awsClient := ec2.NewFromConfig(c.awsConfig, func(o *ec2.Options) {
		o.Region = args.VPC.Location
	})

	// Find all Internet Gateways
	describeInternetGatewaysOutput, err := awsClient.DescribeInternetGateways(ctx, &ec2.DescribeInternetGatewaysInput{
		Filters: []types.Filter{
			{
				Name:   aws.String("attachment.vpc-id"),
				Values: []string{args.VPC.CloudID},
			},
		},
	})
	if err != nil {
		return err
	}

	for _, internetGateway := range describeInternetGatewaysOutput.InternetGateways {
		// Detach the Internet Gateway from the VPC
		_, err = awsClient.DetachInternetGateway(ctx, &ec2.DetachInternetGatewayInput{
			InternetGatewayId: internetGateway.InternetGatewayId,
			VpcId:             aws.String(args.VPC.CloudID),
		})
		if err != nil {
			return err
		}

		// Delete the Internet Gateway
		_, err = awsClient.DeleteInternetGateway(ctx, &ec2.DeleteInternetGatewayInput{
			InternetGatewayId: internetGateway.InternetGatewayId,
		})
		if err != nil {
			return err
		}
	}

	// Delete the VPC
	_, err = awsClient.DeleteVpc(ctx, &ec2.DeleteVpcInput{
		VpcId: aws.String(args.VPC.CloudID),
	})
	if err != nil {
		return err
	}

	return nil
}

func (c *AWSClient) createPublicSubnet(ctx context.Context, args v1.CreatePublicSubnetArgs) (*types.Subnet, error) {
	awsClient := ec2.NewFromConfig(c.awsConfig, func(o *ec2.Options) {
		o.Region = args.VPC.Location
	})

	// Create the Public Subnet
	createSubnetOutput, err := awsClient.CreateSubnet(ctx, &ec2.CreateSubnetInput{
		VpcId:            aws.String(args.VPC.CloudID),
		CidrBlock:        aws.String(args.CidrBlock),
		AvailabilityZone: aws.String(args.Location),
		TagSpecifications: []types.TagSpecification{
			{
				ResourceType: types.ResourceTypeSubnet,
				Tags: makeTags(map[string]string{
					"brev-vpc-id":      args.VPC.CloudID,
					"brev-subnet-type": "public",
				}),
			},
		},
	})
	if err != nil {
		return nil, err
	}

	// Get or create the Public Route Table
	awsPublicRouteTable, err := c.getOrCreatePublicRouteTable(ctx, args)
	if err != nil {
		return nil, err
	}

	// Associate the Public Subnet with the Public Route Table
	_, err = awsClient.AssociateRouteTable(ctx, &ec2.AssociateRouteTableInput{
		RouteTableId: awsPublicRouteTable.RouteTableId,
		SubnetId:     createSubnetOutput.Subnet.SubnetId,
	})
	if err != nil {
		return nil, err
	}

	return createSubnetOutput.Subnet, nil
}

func (c *AWSClient) getOrCreateInternetGateway(ctx context.Context, args v1.CreatePublicSubnetArgs) (*types.InternetGateway, error) {
	awsClient := ec2.NewFromConfig(c.awsConfig, func(o *ec2.Options) {
		o.Region = args.VPC.Location
	})

	// Find the Internet Gateway
	igwNameTag := fmt.Sprintf("%s-public", args.VPC.CloudID)

	describeInternetGatewaysOutput, err := awsClient.DescribeInternetGateways(ctx, &ec2.DescribeInternetGatewaysInput{
		Filters: []types.Filter{
			{
				Name:   aws.String("attachment.vpc-id"),
				Values: []string{args.VPC.CloudID},
			},
			{
				Name:   aws.String("tag:Name"),
				Values: []string{igwNameTag},
			},
		},
	})
	if err != nil {
		return nil, err
	}

	// If there are multiple internet gateways, return an error
	if len(describeInternetGatewaysOutput.InternetGateways) > 1 {
		return nil, fmt.Errorf("multiple internet gateways found for VPC %s", args.VPC.CloudID)
	}

	// If there is one internet gateway, return it
	if len(describeInternetGatewaysOutput.InternetGateways) == 1 {
		return &describeInternetGatewaysOutput.InternetGateways[0], nil
	}

	// If there is no internet gateway, create one
	createInternetGatewayOutput, err := awsClient.CreateInternetGateway(ctx, &ec2.CreateInternetGatewayInput{
		TagSpecifications: []types.TagSpecification{
			{
				ResourceType: types.ResourceTypeInternetGateway,
				Tags: makeTags(map[string]string{
					"Name": igwNameTag,
				}),
			},
		},
	})
	if err != nil {
		return nil, err
	}

	return createInternetGatewayOutput.InternetGateway, nil
}

func (c *AWSClient) getOrCreatePublicRouteTable(ctx context.Context, args v1.CreatePublicSubnetArgs) (*types.RouteTable, error) {
	awsClient := ec2.NewFromConfig(c.awsConfig, func(o *ec2.Options) {
		o.Region = args.VPC.Location
	})

	// Find the Public Route Table
	rtNameTag := fmt.Sprintf("%s-public", args.VPC.CloudID)

	describeRouteTablesOutput, err := awsClient.DescribeRouteTables(ctx, &ec2.DescribeRouteTablesInput{
		Filters: []types.Filter{
			{
				Name:   aws.String("vpc-id"),
				Values: []string{args.VPC.CloudID},
			},
			{
				Name:   aws.String("tag:Name"),
				Values: []string{rtNameTag},
			},
		},
	})
	if err != nil {
		return nil, err
	}

	// If there are multiple public route tables, return an error
	if len(describeRouteTablesOutput.RouteTables) > 1 {
		return nil, fmt.Errorf("multiple public route tables found for VPC %s", args.VPC.CloudID)
	}

	// If there is one public route table, return it
	if len(describeRouteTablesOutput.RouteTables) == 1 {
		return &describeRouteTablesOutput.RouteTables[0], nil
	}

	// If there is no public route table, create one
	createRouteTableOutput, err := awsClient.CreateRouteTable(ctx, &ec2.CreateRouteTableInput{
		VpcId: aws.String(args.VPC.CloudID),
		TagSpecifications: []types.TagSpecification{
			{
				ResourceType: types.ResourceTypeRouteTable,
				Tags: makeTags(map[string]string{
					"Name":                  rtNameTag,
					"brev-vpc-id":           args.VPC.CloudID,
					"brev-route-table-type": "public",
				}),
			},
		},
	})
	if err != nil {
		return nil, err
	}

	// Get or create the Internet Gateway
	awsInternetGateway, err := c.getOrCreateInternetGateway(ctx, args)
	if err != nil {
		return nil, err
	}

	// Create the route to the Internet Gateway
	_, err = awsClient.CreateRoute(ctx, &ec2.CreateRouteInput{
		RouteTableId:         createRouteTableOutput.RouteTable.RouteTableId,
		GatewayId:            awsInternetGateway.InternetGatewayId,
		DestinationCidrBlock: aws.String("0.0.0.0/0"),
	})
	if err != nil {
		return nil, err
	}

	return createRouteTableOutput.RouteTable, nil
}

func (c *AWSClient) createPrivateSubnet(ctx context.Context, args v1.CreatePrivateSubnetArgs) (*types.Subnet, error) {
	awsClient := ec2.NewFromConfig(c.awsConfig, func(o *ec2.Options) {
		o.Region = args.VPC.Location
	})

	// Find a public subnet in the same availability zone
	publicSubnets, err := awsClient.DescribeSubnets(ctx, &ec2.DescribeSubnetsInput{
		Filters: []types.Filter{
			{
				Name:   aws.String("vpc-id"),
				Values: []string{args.VPC.CloudID},
			},
			{
				Name:   aws.String("availability-zone"),
				Values: []string{args.Location},
			},
			{
				Name:   aws.String("tag:brev-subnet-type"),
				Values: []string{"public"},
			},
		},
	})
	if err != nil {
		return nil, err
	}

	if len(publicSubnets.Subnets) == 0 {
		return nil, fmt.Errorf("no public subnets found for VPC %s", args.VPC.CloudID)
	}

	// Arbitrarily choose the first public subnet
	publicSubnet := publicSubnets.Subnets[0]

	// Create the Private Subnet
	createSubnetOutput, err := awsClient.CreateSubnet(ctx, &ec2.CreateSubnetInput{
		VpcId:            aws.String(args.VPC.CloudID),
		CidrBlock:        aws.String(args.CidrBlock),
		AvailabilityZone: aws.String(args.Location),
	})
	if err != nil {
		return nil, err
	}

	// Get or create the NAT Gateway
	awsNatGateway, err := c.createNatGateway(ctx, args, publicSubnet)
	if err != nil {
		return nil, err
	}

	// Create a private route table
	createRouteTableOutput, err := awsClient.CreateRouteTable(ctx, &ec2.CreateRouteTableInput{
		VpcId: aws.String(args.VPC.CloudID),
		TagSpecifications: []types.TagSpecification{
			{
				ResourceType: types.ResourceTypeRouteTable,
				Tags: makeTags(map[string]string{
					"Name":             fmt.Sprintf("%s-private", args.VPC.CloudID),
					"brev-vpc-id":      args.VPC.CloudID,
					"brev-subnet-type": "private",
				}),
			},
		},
	})
	if err != nil {
		return nil, err
	}

	// Associate the Private Subnet with the Private Route Table
	_, err = awsClient.AssociateRouteTable(ctx, &ec2.AssociateRouteTableInput{
		RouteTableId: createRouteTableOutput.RouteTable.RouteTableId,
		SubnetId:     createSubnetOutput.Subnet.SubnetId,
	})
	if err != nil {
		return nil, err
	}

	// Create a route to the NAT Gateway
	_, err = awsClient.CreateRoute(ctx, &ec2.CreateRouteInput{
		RouteTableId:         createRouteTableOutput.RouteTable.RouteTableId,
		DestinationCidrBlock: aws.String("0.0.0.0/0"),
		GatewayId:            awsNatGateway.NatGatewayId,
	})
	if err != nil {
		return nil, err
	}

	return createSubnetOutput.Subnet, nil
}

func (c *AWSClient) createNatGateway(ctx context.Context, args v1.CreatePrivateSubnetArgs, subnet types.Subnet) (*types.NatGateway, error) {
	awsClient := ec2.NewFromConfig(c.awsConfig, func(o *ec2.Options) {
		o.Region = args.VPC.Location
	})

	// Allocate an Elastic IP address for the NAT Gateway
	allocateElasticIPOutput, err := awsClient.AllocateAddress(ctx, &ec2.AllocateAddressInput{
		Domain: types.DomainTypeVpc,
	})
	if err != nil {
		return nil, err
	}

	// Create the NAT Gateway in the provided subnet
	createNatGatewayOutput, err := awsClient.CreateNatGateway(ctx, &ec2.CreateNatGatewayInput{
		SubnetId:     subnet.SubnetId,
		AllocationId: allocateElasticIPOutput.AllocationId,
		TagSpecifications: []types.TagSpecification{
			{
				ResourceType: types.ResourceTypeNatgateway,
				Tags: makeTags(map[string]string{
					"Name":        fmt.Sprintf("%s-nat", args.VPC.CloudID),
					"brev-vpc-id": args.VPC.CloudID,
				}),
			},
		},
	})
	if err != nil {
		return nil, err
	}

	// Wait for the NAT Gateway to be available
	for {
		describeNatGatewaysOutput, err := awsClient.DescribeNatGateways(ctx, &ec2.DescribeNatGatewaysInput{
			NatGatewayIds: []string{*createNatGatewayOutput.NatGateway.NatGatewayId},
		})
		if err != nil {
			return nil, err
		}

		if describeNatGatewaysOutput.NatGateways[0].State == types.NatGatewayStateAvailable {
			break
		}

		time.Sleep(15 * time.Second)
	}

	return createNatGatewayOutput.NatGateway, nil
}

func (c *AWSClient) deleteSubnet(ctx context.Context, args v1.DeleteSubnetArgs) error {
	awsClient := ec2.NewFromConfig(c.awsConfig, func(o *ec2.Options) {
		o.Region = args.Subnet.Location
	})

	// Get the subnet
	describeSubnetsOutput, err := awsClient.DescribeSubnets(ctx, &ec2.DescribeSubnetsInput{
		SubnetIds: []string{args.Subnet.CloudID},
	})
	if err != nil {
		return err
	}
	if len(describeSubnetsOutput.Subnets) == 0 {
		return fmt.Errorf("no subnet found for %s", args.Subnet.CloudID)
	}
	subnet := describeSubnetsOutput.Subnets[0]

	// Find associated route tables
	describeRouteTablesOutput, err := awsClient.DescribeRouteTables(ctx, &ec2.DescribeRouteTablesInput{
		Filters: []types.Filter{
			{
				Name:   aws.String("association.subnet-id"),
				Values: []string{*subnet.SubnetId},
			},
		},
	})
	if err != nil {
		return err
	}

	// Delete associated route tables
	if len(describeRouteTablesOutput.RouteTables) > 0 {
		for _, routeTable := range describeRouteTablesOutput.RouteTables {
			// Delete the private route table
			_, err = awsClient.DeleteRouteTable(ctx, &ec2.DeleteRouteTableInput{
				RouteTableId: routeTable.RouteTableId,
			})
			if err != nil {
				return err
			}
		}
	}

	// Find associated NAT gateways
	describeNatGatewaysOutput, err := awsClient.DescribeNatGateways(ctx, &ec2.DescribeNatGatewaysInput{
		Filter: []types.Filter{
			{
				Name:   aws.String("subnet-id"),
				Values: []string{*subnet.SubnetId},
			},
		},
	})
	if err != nil {
		return err
	}

	// Delete associated NAT gateways
	if len(describeNatGatewaysOutput.NatGateways) > 0 {
		for _, natGateway := range describeNatGatewaysOutput.NatGateways {
			// Delete the NAT Gateway
			_, err = awsClient.DeleteNatGateway(ctx, &ec2.DeleteNatGatewayInput{
				NatGatewayId: natGateway.NatGatewayId,
			})
			if err != nil {
				return err
			}

			// Release the Elastic IP address
			_, err = awsClient.ReleaseAddress(ctx, &ec2.ReleaseAddressInput{
				AllocationId: natGateway.NatGatewayAddresses[0].AllocationId,
			})
			if err != nil {
				return err
			}
		}
	}

	// Delete the subnet
	_, err = awsClient.DeleteSubnet(ctx, &ec2.DeleteSubnetInput{
		SubnetId: subnet.SubnetId,
	})
	if err != nil {
		return err
	}

	return nil
}

func makeTags(tags map[string]string) []types.Tag {
	awsTags := make([]types.Tag, 0, len(tags))
	for key, value := range tags {
		awsTags = append(awsTags, types.Tag{Key: aws.String(key), Value: aws.String(value)})
	}
	return awsTags
}
