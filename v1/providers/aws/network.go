package v1

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/smithy-go"

	"github.com/brevdev/cloud/internal/errors"
	v1 "github.com/brevdev/cloud/v1"
)

func (c *AWSClient) CreateVPC(ctx context.Context, args v1.CreateVPCArgs) (*v1.VPC, error) {
	// Validate the inputs
	publicSubnetArgs := filterSubnetArgs(args.Subnets, v1.SubnetTypePublic)
	privateSubnetArgs := filterSubnetArgs(args.Subnets, v1.SubnetTypePrivate)

	// If there are no public subnets but there are private subnets, return an error, as we need at least one
	// public subnet to create NAT gateways for private subnets.
	if len(publicSubnetArgs) == 0 && len(privateSubnetArgs) > 0 {
		return nil, errors.WrapAndTrace(fmt.Errorf("VPC creation with private subnets requires at least one public subnet, but no public subnets were provided for VPC %s", args.RefID))
	}

	// Create the AWS client in the specified region
	awsClient := ec2.NewFromConfig(c.awsConfig)

	// Create the VPC and subnets
	awsVPC, subnets, err := createCompleteVPC(ctx, awsClient, args, publicSubnetArgs, privateSubnetArgs)
	if err != nil {
		return nil, errors.WrapAndTrace(err)
	}

	tags := make(map[string]string)
	for _, tag := range awsVPC.Tags {
		tags[*tag.Key] = *tag.Value
	}

	return &v1.VPC{
		RefID:          args.RefID,
		Name:           args.Name,
		Location:       c.region,
		CloudCredRefID: c.GetReferenceID(),
		Provider:       CloudProviderID,
		Cloud:          CloudProviderID,
		ID:             v1.CloudProviderResourceID(*awsVPC.VpcId),
		CidrBlock:      *awsVPC.CidrBlock,
		Status:         v1.VPCStatusPending,
		Subnets:        subnets,
		Tags:           v1.Tags(tags),
	}, nil
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

func createCompleteVPC(ctx context.Context, awsClient *ec2.Client, args v1.CreateVPCArgs, publicSubnetArgs []v1.CreateSubnetArgs, privateSubnetArgs []v1.CreateSubnetArgs) (*types.Vpc, []*v1.Subnet, error) {
	// Create the VPC
	vpc, err := createVPC(ctx, awsClient, args)
	if err != nil {
		return nil, nil, errors.WrapAndTrace(err)
	}

	// Enable DNS support for the VPC
	err = enableVPCDNSSupport(ctx, awsClient, vpc)
	if err != nil {
		return nil, nil, errors.WrapAndTrace(err)
	}

	// Enable DNS hostnames for the VPC
	err = enableVPCDNSHostnames(ctx, awsClient, vpc)
	if err != nil {
		return nil, nil, errors.WrapAndTrace(err)
	}

	// Create an Internet Gateway for the VPC
	_, err = createInternetGateway(ctx, awsClient, vpc, args)
	if err != nil {
		return nil, nil, errors.WrapAndTrace(err)
	}

	var subnets []*v1.Subnet

	// Create public subnets
	var publicSubnets []*types.Subnet
	for _, subnetArgs := range publicSubnetArgs {
		// Create the public subnet
		publicSubnet, err := createPublicSubnet(ctx, awsClient, vpc, subnetArgs, args)
		if err != nil {
			return nil, nil, errors.WrapAndTrace(err)
		}
		publicSubnets = append(publicSubnets, publicSubnet)

		subnets = append(subnets, awsSubnetToCloudSubnet(publicSubnet, vpc))
	}

	for i := range privateSubnetArgs {
		// Choose a public subnet for the NAT gateway
		natGatewaySubnet := publicSubnets[i%len(publicSubnets)]
		subnetArgs := privateSubnetArgs[i]

		// Create the private subnet
		privateSubnet, err := createPrivateSubnet(ctx, awsClient, vpc, natGatewaySubnet, subnetArgs, args)
		if err != nil {
			return nil, nil, errors.WrapAndTrace(err)
		}

		subnets = append(subnets, awsSubnetToCloudSubnet(privateSubnet, vpc))
	}

	return vpc, subnets, nil
}

func createVPC(ctx context.Context, awsClient *ec2.Client, args v1.CreateVPCArgs) (*types.Vpc, error) {
	// Convert the tags to AWS tags
	tags := make(map[string]string)
	for key, value := range args.Tags {
		tags[key] = value
	}

	// Add the required tags
	tags[tagName] = args.Name
	tags[tagBrevRefID] = args.RefID
	tags[tagCreatedBy] = tagBrevCloudSDK

	awsTags := makeEC2Tags(tags)

	input := &ec2.CreateVpcInput{
		CidrBlock: &args.CidrBlock,
		TagSpecifications: []types.TagSpecification{
			{
				ResourceType: types.ResourceTypeVpc,
				Tags:         awsTags,
			},
		},
	}

	output, err := awsClient.CreateVpc(ctx, input)
	if err != nil {
		return nil, errors.WrapAndTrace(err)
	}

	return output.Vpc, nil
}

func enableVPCDNSSupport(ctx context.Context, awsClient *ec2.Client, vpc *types.Vpc) error {
	input := &ec2.ModifyVpcAttributeInput{
		VpcId: vpc.VpcId,
		EnableDnsSupport: &types.AttributeBooleanValue{
			Value: aws.Bool(true),
		},
	}

	_, err := awsClient.ModifyVpcAttribute(ctx, input)
	if err != nil {
		return errors.WrapAndTrace(err)
	}
	return nil
}

func enableVPCDNSHostnames(ctx context.Context, awsClient *ec2.Client, vpc *types.Vpc) error {
	input := &ec2.ModifyVpcAttributeInput{
		VpcId: vpc.VpcId,
		EnableDnsHostnames: &types.AttributeBooleanValue{
			Value: aws.Bool(true),
		},
	}

	_, err := awsClient.ModifyVpcAttribute(ctx, input)
	if err != nil {
		return errors.WrapAndTrace(err)
	}
	return nil
}

func createInternetGateway(ctx context.Context, awsClient *ec2.Client, vpc *types.Vpc, args v1.CreateVPCArgs) (*types.InternetGateway, error) {
	tags := make(map[string]string)
	for key, value := range args.Tags {
		tags[key] = value
	}

	tags[tagName] = fmt.Sprintf("%s-public", *vpc.VpcId)
	tags[tagBrevVPCID] = *vpc.VpcId
	tags[tagCreatedBy] = tagBrevCloudSDK

	awsTags := makeEC2Tags(tags)

	createInput := &ec2.CreateInternetGatewayInput{
		TagSpecifications: []types.TagSpecification{
			{
				ResourceType: types.ResourceTypeInternetGateway,
				Tags:         awsTags,
			},
		},
	}
	// Create an Internet Gateway for the VPC
	createOutput, err := awsClient.CreateInternetGateway(ctx, createInput)
	if err != nil {
		return nil, errors.WrapAndTrace(err)
	}

	internetGateway := createOutput.InternetGateway

	// Attach the Internet Gateway to the VPC
	_, err = awsClient.AttachInternetGateway(ctx, &ec2.AttachInternetGatewayInput{
		InternetGatewayId: internetGateway.InternetGatewayId,
		VpcId:             vpc.VpcId,
	})
	if err != nil {
		return nil, errors.WrapAndTrace(err)
	}

	return internetGateway, nil
}

func awsSubnetToCloudSubnet(awsSubnet *types.Subnet, vpc *types.Vpc) *v1.Subnet {
	tags := make(map[string]string)
	for _, tag := range awsSubnet.Tags {
		tags[*tag.Key] = *tag.Value
	}

	return &v1.Subnet{
		ID:        v1.CloudProviderResourceID(*awsSubnet.SubnetId),
		Name:      tags[tagName],
		VPCID:     v1.CloudProviderResourceID(*vpc.VpcId),
		Location:  *awsSubnet.AvailabilityZone,
		CidrBlock: *awsSubnet.CidrBlock,
		Type:      v1.SubnetTypePublic,
		Tags:      v1.Tags(tags),
	}
}

func createPublicSubnet(ctx context.Context, awsClient *ec2.Client, vpc *types.Vpc, createSubnetArgs v1.CreateSubnetArgs, createVPCArgs v1.CreateVPCArgs) (*types.Subnet, error) {
	tags := make(map[string]string)
	for key, value := range createSubnetArgs.Tags {
		tags[key] = value
	}

	tags[tagName] = fmt.Sprintf("%s-public", *vpc.VpcId)
	tags[tagBrevVPCID] = *vpc.VpcId
	tags[tagBrevSubnetType] = string(createSubnetArgs.Type)
	tags[tagCreatedBy] = tagBrevCloudSDK

	awsTags := makeEC2Tags(tags)

	input := &ec2.CreateSubnetInput{
		VpcId:     vpc.VpcId,
		CidrBlock: aws.String(createSubnetArgs.CidrBlock),
		TagSpecifications: []types.TagSpecification{
			{
				ResourceType: types.ResourceTypeSubnet,
				Tags:         awsTags,
			},
		},
	}
	output, err := awsClient.CreateSubnet(ctx, input)
	if err != nil {
		return nil, errors.WrapAndTrace(err)
	}

	subnet := output.Subnet

	// Get or create the Public Route Table for the VPC
	publicRouteTable, err := getOrCreatePublicRouteTable(ctx, awsClient, vpc, createVPCArgs)
	if err != nil {
		return nil, errors.WrapAndTrace(err)
	}

	// Associate the Public Subnet with the Public Route Table
	_, err = awsClient.AssociateRouteTable(ctx, &ec2.AssociateRouteTableInput{
		RouteTableId: publicRouteTable.RouteTableId,
		SubnetId:     subnet.SubnetId,
	})
	if err != nil {
		return nil, errors.WrapAndTrace(err)
	}

	return subnet, nil
}

func getOrCreatePublicRouteTable(ctx context.Context, awsClient *ec2.Client, vpc *types.Vpc, args v1.CreateVPCArgs) (*types.RouteTable, error) {
	// Find the Public Route Table
	rtNameTag := fmt.Sprintf("%s-public", *vpc.VpcId)

	describeRouteTablesOutput, err := awsClient.DescribeRouteTables(ctx, &ec2.DescribeRouteTablesInput{
		Filters: []types.Filter{
			{
				Name:   aws.String("vpc-id"),
				Values: []string{*vpc.VpcId},
			},
			{
				Name:   aws.String("tag:Name"),
				Values: []string{rtNameTag},
			},
		},
	})
	if err != nil {
		return nil, errors.WrapAndTrace(err)
	}

	// If there are multiple public route tables, return an error
	if len(describeRouteTablesOutput.RouteTables) > 1 {
		return nil, fmt.Errorf("multiple public route tables found for VPC %s", *vpc.VpcId)
	}

	// If there is one public route table, return it
	if len(describeRouteTablesOutput.RouteTables) == 1 {
		return &describeRouteTablesOutput.RouteTables[0], nil
	}

	// If there is no public route table, create one
	tags := make(map[string]string)
	for key, value := range args.Tags {
		tags[key] = value
	}

	tags[tagName] = rtNameTag
	tags[tagBrevVPCID] = *vpc.VpcId
	tags[tagCreatedBy] = tagBrevCloudSDK

	awsTags := makeEC2Tags(tags)
	input := &ec2.CreateRouteTableInput{
		VpcId: aws.String(*vpc.VpcId),
		TagSpecifications: []types.TagSpecification{
			{
				ResourceType: types.ResourceTypeRouteTable,
				Tags:         awsTags,
			},
		},
	}
	output, err := awsClient.CreateRouteTable(ctx, input)
	if err != nil {
		return nil, errors.WrapAndTrace(err)
	}

	routeTable := output.RouteTable

	// Get or create the Internet Gateway
	internetGateway, err := getOrCreateInternetGateway(ctx, awsClient, vpc, args)
	if err != nil {
		return nil, errors.WrapAndTrace(err)
	}

	// Create the route to the Internet Gateway
	_, err = awsClient.CreateRoute(ctx, &ec2.CreateRouteInput{
		RouteTableId:         routeTable.RouteTableId,
		GatewayId:            internetGateway.InternetGatewayId,
		DestinationCidrBlock: aws.String("0.0.0.0/0"),
	})
	if err != nil {
		return nil, errors.WrapAndTrace(err)
	}

	return routeTable, nil
}

func getOrCreateInternetGateway(ctx context.Context, awsClient *ec2.Client, vpc *types.Vpc, args v1.CreateVPCArgs) (*types.InternetGateway, error) {
	// Find the Internet Gateway
	igwNameTag := fmt.Sprintf("%s-public", *vpc.VpcId)

	describeInternetGatewaysOutput, err := awsClient.DescribeInternetGateways(ctx, &ec2.DescribeInternetGatewaysInput{
		Filters: []types.Filter{
			{
				Name:   aws.String("attachment.vpc-id"),
				Values: []string{*vpc.VpcId},
			},
			{
				Name:   aws.String("tag:Name"),
				Values: []string{igwNameTag},
			},
		},
	})
	if err != nil {
		return nil, errors.WrapAndTrace(err)
	}

	// If there are multiple internet gateways, return an error
	if len(describeInternetGatewaysOutput.InternetGateways) > 1 {
		return nil, fmt.Errorf("multiple internet gateways found for VPC %s", *vpc.VpcId)
	}

	// If there is one internet gateway, return it
	if len(describeInternetGatewaysOutput.InternetGateways) == 1 {
		return &describeInternetGatewaysOutput.InternetGateways[0], nil
	}

	// If there is no internet gateway, create one
	tags := make(map[string]string)
	for key, value := range args.Tags {
		tags[key] = value
	}

	tags[tagName] = igwNameTag
	tags[tagBrevVPCID] = *vpc.VpcId
	tags[tagCreatedBy] = tagBrevCloudSDK

	awsTags := makeEC2Tags(tags)

	input := &ec2.CreateInternetGatewayInput{
		TagSpecifications: []types.TagSpecification{
			{
				ResourceType: types.ResourceTypeInternetGateway,
				Tags:         awsTags,
			},
		},
	}
	output, err := awsClient.CreateInternetGateway(ctx, input)
	if err != nil {
		return nil, errors.WrapAndTrace(err)
	}

	return output.InternetGateway, nil
}

func createPrivateSubnet(ctx context.Context, awsClient *ec2.Client, vpc *types.Vpc, natGatewaySubnet *types.Subnet, createSubnetArgs v1.CreateSubnetArgs, createVPCArgs v1.CreateVPCArgs) (*types.Subnet, error) {
	tags := make(map[string]string)
	for key, value := range createSubnetArgs.Tags {
		tags[key] = value
	}

	tags[tagName] = fmt.Sprintf("%s-private", *vpc.VpcId)
	tags[tagBrevVPCID] = *vpc.VpcId
	tags[tagBrevSubnetType] = string(createSubnetArgs.Type)
	tags[tagCreatedBy] = tagBrevCloudSDK

	awsTags := makeEC2Tags(tags)

	createSubnetInput := &ec2.CreateSubnetInput{
		VpcId:            vpc.VpcId,
		CidrBlock:        aws.String(createSubnetArgs.CidrBlock),
		AvailabilityZone: aws.String(*natGatewaySubnet.AvailabilityZone),
		TagSpecifications: []types.TagSpecification{
			{
				ResourceType: types.ResourceTypeSubnet,
				Tags:         awsTags,
			},
		},
	}
	createSubnetOutput, err := awsClient.CreateSubnet(ctx, createSubnetInput)
	if err != nil {
		return nil, errors.WrapAndTrace(err)
	}

	// Get or create the NAT Gateway
	natGateway, err := createNatGateway(ctx, awsClient, vpc, natGatewaySubnet, createSubnetArgs)
	if err != nil {
		return nil, errors.WrapAndTrace(err)
	}

	// Create a private route table
	tags = make(map[string]string)
	for key, value := range createSubnetArgs.Tags {
		tags[key] = value
	}

	tags[tagName] = fmt.Sprintf("%s-private", *vpc.VpcId)
	tags[tagBrevVPCID] = *vpc.VpcId
	tags[tagCreatedBy] = tagBrevCloudSDK

	awsTags = makeEC2Tags(tags)

	createRouteTableInput := &ec2.CreateRouteTableInput{
		VpcId: aws.String(*vpc.VpcId),
		TagSpecifications: []types.TagSpecification{
			{
				ResourceType: types.ResourceTypeRouteTable,
				Tags:         awsTags,
			},
		},
	}
	createRouteTableOutput, err := awsClient.CreateRouteTable(ctx, createRouteTableInput)
	if err != nil {
		return nil, errors.WrapAndTrace(err)
	}

	routeTable := createRouteTableOutput.RouteTable

	// Associate the Private Subnet with the Private Route Table
	_, err = awsClient.AssociateRouteTable(ctx, &ec2.AssociateRouteTableInput{
		RouteTableId: routeTable.RouteTableId,
		SubnetId:     createSubnetOutput.Subnet.SubnetId,
	})
	if err != nil {
		return nil, errors.WrapAndTrace(err)
	}

	// Create a route to the NAT Gateway
	_, err = awsClient.CreateRoute(ctx, &ec2.CreateRouteInput{
		RouteTableId:         routeTable.RouteTableId,
		DestinationCidrBlock: aws.String("0.0.0.0/0"),
		GatewayId:            natGateway.NatGatewayId,
	})
	if err != nil {
		return nil, errors.WrapAndTrace(err)
	}

	return createSubnetOutput.Subnet, nil
}

func createNatGateway(ctx context.Context, awsClient *ec2.Client, vpc *types.Vpc, subnet *types.Subnet, args v1.CreateSubnetArgs) (*types.NatGateway, error) {
	// Allocate an Elastic IP address for the NAT Gateway
	allocateElasticIPOutput, err := awsClient.AllocateAddress(ctx, &ec2.AllocateAddressInput{
		Domain: types.DomainTypeVpc,
	})
	if err != nil {
		return nil, errors.WrapAndTrace(err)
	}

	// Create the NAT Gateway in the provided subnet
	tags := make(map[string]string)
	for key, value := range args.Tags {
		tags[key] = value
	}

	tags[tagName] = fmt.Sprintf("%s-nat", *vpc.VpcId)
	tags[tagBrevVPCID] = *vpc.VpcId
	tags[tagCreatedBy] = tagBrevCloudSDK

	awsTags := makeEC2Tags(tags)

	createNatGatewayInput := &ec2.CreateNatGatewayInput{
		SubnetId:     subnet.SubnetId,
		AllocationId: allocateElasticIPOutput.AllocationId,
		TagSpecifications: []types.TagSpecification{
			{
				ResourceType: types.ResourceTypeNatgateway,
				Tags:         awsTags,
			},
		},
	}
	createNatGatewayOutput, err := awsClient.CreateNatGateway(ctx, createNatGatewayInput)
	if err != nil {
		return nil, errors.WrapAndTrace(err)
	}

	natGateway := createNatGatewayOutput.NatGateway

	// Wait for the NAT Gateway to be available
	w := ec2.NewNatGatewayAvailableWaiter(awsClient, func(o *ec2.NatGatewayAvailableWaiterOptions) {
		o.MaxDelay = 10 * time.Second
		o.MinDelay = 10 * time.Second
	})
	err = w.Wait(ctx, &ec2.DescribeNatGatewaysInput{
		NatGatewayIds: []string{*natGateway.NatGatewayId},
	}, 10*time.Minute)
	if err != nil {
		return nil, errors.WrapAndTrace(err)
	}

	return natGateway, nil
}

// GetVPC retrieves a VPC from AWS
func (c *AWSClient) GetVPC(ctx context.Context, args v1.GetVPCArgs) (*v1.VPC, error) {
	awsClient := ec2.NewFromConfig(c.awsConfig)

	awsVPC, err := getVPC(ctx, awsClient, args)
	if err != nil {
		return nil, errors.WrapAndTrace(err)
	}

	tags := make(map[string]string)
	for _, tag := range awsVPC.Tags {
		tags[*tag.Key] = *tag.Value
	}
	brevVPCName := tags[tagName]
	brevRefID := tags[tagBrevRefID]

	status, err := getVPCStatus(ctx, awsClient, awsVPC)
	if err != nil {
		return nil, err
	}

	subnets, err := getVPCSubnets(ctx, awsClient, awsVPC)
	if err != nil {
		return nil, errors.WrapAndTrace(err)
	}

	return &v1.VPC{
		RefID:          brevRefID,
		Name:           brevVPCName,
		Location:       c.region,
		ID:             v1.CloudProviderResourceID(*awsVPC.VpcId),
		CloudCredRefID: c.GetReferenceID(),
		Provider:       CloudProviderID,
		Cloud:          CloudProviderID,
		CidrBlock:      *awsVPC.CidrBlock,
		Status:         status,
		Subnets:        subnets,
		Tags:           v1.Tags(tags),
	}, nil
}

func getVPC(ctx context.Context, awsClient *ec2.Client, args v1.GetVPCArgs) (*types.Vpc, error) {
	describeVPCsOutput, err := awsClient.DescribeVpcs(ctx, &ec2.DescribeVpcsInput{
		VpcIds: []string{string(args.ID)},
	})
	if err != nil {
		var apiErr smithy.APIError
		if errors.As(err, &apiErr) && apiErr.ErrorCode() == "InvalidVpcID.NotFound" {
			return nil, v1.ErrResourceNotFound
		}
		return nil, errors.WrapAndTrace(err)
	}

	if len(describeVPCsOutput.Vpcs) == 0 {
		return nil, nil
	}

	return &describeVPCsOutput.Vpcs[0], nil
}

func getVPCStatus(ctx context.Context, awsClient *ec2.Client, awsVPC *types.Vpc) (v1.VPCStatus, error) {
	if awsVPC.State == types.VpcStatePending {
		return v1.VPCStatusPending, nil
	}

	// The VPC is available if all NAT gateways are available
	natGateways, err := awsClient.DescribeNatGateways(ctx, &ec2.DescribeNatGatewaysInput{
		Filter: []types.Filter{
			{
				Name:   aws.String("vpc-id"),
				Values: []string{*awsVPC.VpcId},
			},
		},
	})
	if err != nil {
		return v1.VPCStatusAvailable, errors.WrapAndTrace(err)
	}

	for _, natGateway := range natGateways.NatGateways {
		if natGateway.State != types.NatGatewayStateAvailable {
			return v1.VPCStatusPending, nil
		}
	}

	return v1.VPCStatusAvailable, nil
}

func getVPCSubnets(ctx context.Context, awsClient *ec2.Client, awsVPC *types.Vpc) ([]*v1.Subnet, error) {
	describeSubnetsOutput, err := awsClient.DescribeSubnets(ctx, &ec2.DescribeSubnetsInput{
		Filters: []types.Filter{
			{
				Name:   aws.String("vpc-id"),
				Values: []string{*awsVPC.VpcId},
			},
		},
	})
	if err != nil {
		return nil, errors.WrapAndTrace(err)
	}

	subnets := make([]*v1.Subnet, 0)
	for _, subnet := range describeSubnetsOutput.Subnets {
		tags := make(map[string]string)
		for _, tag := range subnet.Tags {
			tags[*tag.Key] = *tag.Value
		}

		brevSubnetName := tags[tagName]
		brevSubnetType := v1.SubnetType(tags[tagBrevSubnetType])

		subnets = append(subnets, &v1.Subnet{
			ID:        v1.CloudProviderResourceID(*subnet.SubnetId),
			VPCID:     v1.CloudProviderResourceID(*awsVPC.VpcId),
			Location:  *subnet.AvailabilityZone,
			CidrBlock: *subnet.CidrBlock,
			Type:      brevSubnetType,
			Name:      brevSubnetName,
			Tags:      v1.Tags(tags),
		})
	}
	return subnets, nil
}

func (c *AWSClient) DeleteVPC(ctx context.Context, args v1.DeleteVPCArgs) error {
	awsClient := ec2.NewFromConfig(c.awsConfig)

	err := deleteVPC(ctx, awsClient, string(args.ID))
	if err != nil {
		return errors.WrapAndTrace(err)
	}
	return nil
}

func deleteVPC(ctx context.Context, awsClient *ec2.Client, vpcID string) error {
	err := deleteNATGateways(ctx, awsClient, vpcID)
	if err != nil {
		return errors.WrapAndTrace(err)
	}

	err = deleteInternetGateways(ctx, awsClient, vpcID)
	if err != nil {
		return errors.WrapAndTrace(err)
	}

	err = deleteSubnets(ctx, awsClient, vpcID)
	if err != nil {
		return errors.WrapAndTrace(err)
	}

	err = deleteRouteTables(ctx, awsClient, vpcID)
	if err != nil {
		return errors.WrapAndTrace(err)
	}

	// Delete the VPC
	_, err = awsClient.DeleteVpc(ctx, &ec2.DeleteVpcInput{
		VpcId: aws.String(vpcID),
	})
	if err != nil {
		return errors.WrapAndTrace(err)
	}

	return nil
}

func deleteNATGateways(ctx context.Context, awsClient *ec2.Client, vpcID string) error {
	// Find associated NAT gateways
	describeNatGatewaysOutput, err := awsClient.DescribeNatGateways(ctx, &ec2.DescribeNatGatewaysInput{
		Filter: []types.Filter{
			{
				Name:   aws.String("vpc-id"),
				Values: []string{vpcID},
			},
		},
	})
	if err != nil {
		return errors.WrapAndTrace(err)
	}

	// Delete associated NAT gateways
	for _, natGateway := range describeNatGatewaysOutput.NatGateways {
		if natGateway.State == types.NatGatewayStateDeleting || natGateway.State == types.NatGatewayStateDeleted {
			continue
		}

		// Delete the NAT Gateway
		_, err = awsClient.DeleteNatGateway(ctx, &ec2.DeleteNatGatewayInput{
			NatGatewayId: natGateway.NatGatewayId,
		})
		if err != nil {
			return errors.WrapAndTrace(err)
		}

		// Wait until the NAT Gateway is deleted
		w := ec2.NewNatGatewayDeletedWaiter(awsClient, func(o *ec2.NatGatewayDeletedWaiterOptions) {
			o.MaxDelay = 10 * time.Second
			o.MinDelay = 10 * time.Second
		})
		err = w.Wait(ctx, &ec2.DescribeNatGatewaysInput{
			NatGatewayIds: []string{*natGateway.NatGatewayId},
		}, 10*time.Minute)
		if err != nil {
			return errors.WrapAndTrace(err)
		}

		// Release the Elastic IP address
		_, err = awsClient.ReleaseAddress(ctx, &ec2.ReleaseAddressInput{
			AllocationId: natGateway.NatGatewayAddresses[0].AllocationId,
		})
		if err != nil {
			return errors.WrapAndTrace(err)
		}
	}

	return nil
}

func deleteInternetGateways(ctx context.Context, awsClient *ec2.Client, vpcID string) error {
	// Find all Internet Gateways
	describeInternetGatewaysOutput, err := awsClient.DescribeInternetGateways(ctx, &ec2.DescribeInternetGatewaysInput{
		Filters: []types.Filter{
			{
				Name:   aws.String("attachment.vpc-id"),
				Values: []string{vpcID},
			},
		},
	})
	if err != nil {
		return errors.WrapAndTrace(err)
	}

	for _, internetGateway := range describeInternetGatewaysOutput.InternetGateways {
		// Detach the Internet Gateway from the VPC
		_, err = awsClient.DetachInternetGateway(ctx, &ec2.DetachInternetGatewayInput{
			InternetGatewayId: internetGateway.InternetGatewayId,
			VpcId:             aws.String(vpcID),
		})
		if err != nil {
			return errors.WrapAndTrace(err)
		}

		// Delete the Internet Gateway
		_, err = awsClient.DeleteInternetGateway(ctx, &ec2.DeleteInternetGatewayInput{
			InternetGatewayId: internetGateway.InternetGatewayId,
		})
		if err != nil {
			return errors.WrapAndTrace(err)
		}
	}

	return nil
}

func deleteSubnets(ctx context.Context, awsClient *ec2.Client, vpcID string) error {
	// Find all subnets
	describeSubnetsOutput, err := awsClient.DescribeSubnets(ctx, &ec2.DescribeSubnetsInput{
		Filters: []types.Filter{
			{
				Name:   aws.String("vpc-id"),
				Values: []string{vpcID},
			},
		},
	})
	if err != nil {
		return errors.WrapAndTrace(err)
	}

	// Delete all subnets
	for _, subnet := range describeSubnetsOutput.Subnets {
		// Delete the subnet
		_, err = awsClient.DeleteSubnet(ctx, &ec2.DeleteSubnetInput{
			SubnetId: subnet.SubnetId,
		})
		if err != nil {
			return errors.WrapAndTrace(err)
		}
	}

	return nil
}

func deleteRouteTables(ctx context.Context, awsClient *ec2.Client, vpcID string) error {
	// Find all route tables
	describeRouteTablesOutput, err := awsClient.DescribeRouteTables(ctx, &ec2.DescribeRouteTablesInput{
		Filters: []types.Filter{
			{
				Name:   aws.String("vpc-id"),
				Values: []string{vpcID},
			},
			{
				Name:   aws.String("tag:" + tagBrevVPCID), // ensure we do not select the default route table
				Values: []string{vpcID},
			},
		},
	})
	if err != nil {
		return errors.WrapAndTrace(err)
	}

	// Delete all route tables
	for _, routeTable := range describeRouteTablesOutput.RouteTables {
		_, err = awsClient.DeleteRouteTable(ctx, &ec2.DeleteRouteTableInput{
			RouteTableId: routeTable.RouteTableId,
		})
		if err != nil {
			return errors.WrapAndTrace(err)
		}
	}

	return nil
}

func makeEC2Tags(tags map[string]string) []types.Tag {
	awsTags := make([]types.Tag, 0, len(tags))
	for key, value := range tags {
		awsTags = append(awsTags, types.Tag{Key: aws.String(key), Value: aws.String(value)})
	}
	return awsTags
}
