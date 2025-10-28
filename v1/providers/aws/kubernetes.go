package v1

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/eks"
	ekstypes "github.com/aws/aws-sdk-go-v2/service/eks/types"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	iamtypes "github.com/aws/aws-sdk-go-v2/service/iam/types"

	"github.com/brevdev/cloud/internal/errors"
	v1 "github.com/brevdev/cloud/v1"
)

var _ v1.CloudMaintainKubernetes = &AWSClient{}

const iamRolePathPrefix = "/brevcloudsdk/eks/clusters"

func (c *AWSClient) CreateCluster(ctx context.Context, args v1.CreateClusterArgs) (*v1.Cluster, error) {
	eksClient := eks.NewFromConfig(c.awsConfig)
	iamClient := iam.NewFromConfig(c.awsConfig)

	// Fetch the target VPC
	vpc, err := c.GetVPC(ctx, v1.GetVPCArgs{
		ID: args.VPCID,
	})
	if err != nil {
		return nil, errors.WrapAndTrace(err)
	}

	// Create a map of subnetID->subnet for this VPC so that we can find the target subnet
	subnetMap := make(map[string]*v1.Subnet)
	for _, subnet := range vpc.Subnets {
		subnetMap[string(subnet.ID)] = subnet
	}

	// Get the target subnets from the map
	subnets := make([]*v1.Subnet, len(args.SubnetIDs))
	for i, subnetID := range args.SubnetIDs {
		if _, ok := subnetMap[string(subnetID)]; !ok {
			return nil, errors.WrapAndTrace(fmt.Errorf("subnet ID %s does not match VPC %s", subnetID, vpc.ID))
		} else {
			subnets[i] = subnetMap[string(subnetID)]
		}
	}

	// Create the cluster
	awsCluster, err := createEKSCluster(ctx, eksClient, iamClient, subnets, args)
	if err != nil {
		return nil, errors.WrapAndTrace(err)
	}

	return &v1.Cluster{
		ID:                v1.CloudProviderResourceID(*awsCluster.Arn),
		Name:              *awsCluster.Name,
		RefID:             args.RefID,
		Provider:          CloudProviderID,
		Cloud:             CloudProviderID,
		Location:          c.region,
		VPCID:             vpc.ID,
		SubnetIDs:         args.SubnetIDs,
		KubernetesVersion: args.KubernetesVersion,
		Status:            v1.ClusterStatusPending,
		Tags:              args.Tags,
	}, nil
}

func createEKSCluster(ctx context.Context, eksClient *eks.Client, iamClient *iam.Client, subnets []*v1.Subnet, args v1.CreateClusterArgs) (*ekstypes.Cluster, error) {
	serviceRoleARN, err := getOrCreateServiceRoleARN(ctx, iamClient, args)
	if err != nil {
		return nil, errors.WrapAndTrace(err)
	}

	eksCluster, err := createCluster(ctx, eksClient, args, serviceRoleARN, subnets)
	if err != nil {
		return nil, errors.WrapAndTrace(err)
	}

	err = installEKSAddons(ctx, eksClient, eksCluster)
	if err != nil {
		return nil, errors.WrapAndTrace(err)
	}

	return eksCluster, nil
}

func getOrCreateServiceRoleARN(ctx context.Context, iamClient *iam.Client, args v1.CreateClusterArgs) (string, error) {
	serviceRoleName := fmt.Sprintf("%s-service-role", args.RefID)

	// Get and return the role if it exists
	serviceRole, err := iamClient.GetRole(ctx, &iam.GetRoleInput{
		RoleName: aws.String(serviceRoleName),
	})
	if err == nil {
		return *serviceRole.Role.Arn, nil
	} else {
		// The error may be a NoSuchEntityException. If it is, we should ignore the error and create the role.
		var noSuchEntityError *iamtypes.NoSuchEntityException
		if !errors.As(err, &noSuchEntityError) {
			// The error is not a no such entity error, so we return the error
			return "", err
		}
	}

	// Convert the tags to AWS tags
	tags := make(map[string]string)
	for key, value := range args.Tags {
		tags[key] = value
	}

	// Add the required tags
	tags[tagName] = args.Name
	tags[tagBrevRefID] = args.RefID
	tags[tagCreatedBy] = tagBrevCloudSDK
	tags[tagBrevClusterID] = args.RefID

	iamTags := makeIAMTags(tags)

	iamPath := fmt.Sprintf("%s/%s/", iamRolePathPrefix, args.RefID)
	iamPath = strings.ReplaceAll(iamPath, "[^a-zA-Z0-9/]", "")

	// Create the role
	input := &iam.CreateRoleInput{
		RoleName:    aws.String(serviceRoleName),
		Description: aws.String("Role for EKS cluster"),
		Path:        aws.String(iamPath),
		AssumeRolePolicyDocument: aws.String(`{
		"Version": "2012-10-17",
		"Statement": [
			{
				"Effect": "Allow",
				"Principal": {
					"Service": "eks.amazonaws.com"
				},
				"Action": "sts:AssumeRole"
			}
		]
	}`),
		Tags: iamTags,
	}
	output, err := iamClient.CreateRole(ctx, input)
	if err != nil {
		return "", errors.WrapAndTrace(err)
	}

	// Attach the AmazonEKSClusterPolicy to the role
	_, err = iamClient.AttachRolePolicy(ctx, &iam.AttachRolePolicyInput{
		RoleName:  aws.String(serviceRoleName),
		PolicyArn: aws.String("arn:aws:iam::aws:policy/AmazonEKSClusterPolicy"),
	})
	if err != nil {
		return "", errors.WrapAndTrace(err)
	}

	return *output.Role.Arn, nil
}

func createCluster(ctx context.Context, eksClient *eks.Client, args v1.CreateClusterArgs, serviceRoleARN string, subnets []*v1.Subnet) (*ekstypes.Cluster, error) {
	// Convert the tags to AWS tags
	tags := make(map[string]string)
	for key, value := range args.Tags {
		tags[key] = value
	}

	// Add the required tags
	tags[tagName] = args.Name
	tags[tagBrevRefID] = args.RefID
	tags[tagCreatedBy] = tagBrevCloudSDK

	// Convert the subnets to subnet IDs
	subnetIDs := make([]string, len(subnets))
	for i, subnet := range subnets {
		subnetIDs[i] = string(subnet.ID)
	}

	input := &eks.CreateClusterInput{
		Name:    aws.String(args.Name),
		Version: aws.String(args.KubernetesVersion),
		RoleArn: aws.String(serviceRoleARN),
		ResourcesVpcConfig: &ekstypes.VpcConfigRequest{
			SubnetIds: subnetIDs,
		},
		Tags: tags,
	}

	output, err := eksClient.CreateCluster(ctx, input)
	if err != nil {
		return nil, errors.WrapAndTrace(err)
	}

	// Wait for the cluster to be active
	w := eks.NewClusterActiveWaiter(eksClient, func(o *eks.ClusterActiveWaiterOptions) {
		o.MaxDelay = 30 * time.Second
		o.MinDelay = 10 * time.Second
	})
	err = w.Wait(ctx, &eks.DescribeClusterInput{
		Name: output.Cluster.Name,
	}, 10*time.Minute)
	if err != nil {
		return nil, errors.WrapAndTrace(err)
	}

	return output.Cluster, nil
}

func installEKSAddons(ctx context.Context, eksClient *eks.Client, eksCluster *ekstypes.Cluster) error {
	err := installEKSAddon(ctx, eksClient, eksCluster, "vpc-cni")
	if err != nil {
		return err
	}

	err = installEKSAddon(ctx, eksClient, eksCluster, "eks-pod-identity-agent")
	if err != nil {
		return err
	}

	return nil
}

func installEKSAddon(ctx context.Context, eksClient *eks.Client, eksCluster *ekstypes.Cluster, addonName string) error {
	_, err := eksClient.CreateAddon(ctx, &eks.CreateAddonInput{
		ClusterName: eksCluster.Name,
		AddonName:   aws.String(addonName),
	})
	if err != nil {
		return err
	}

	// Wait for the addon to be created
	w := eks.NewAddonActiveWaiter(eksClient, func(o *eks.AddonActiveWaiterOptions) {
		o.MaxDelay = 30 * time.Second
		o.MinDelay = 10 * time.Second
	})
	err = w.Wait(ctx, &eks.DescribeAddonInput{
		ClusterName: eksCluster.Name,
		AddonName:   aws.String(addonName),
	}, 10*time.Minute)
	if err != nil {
		return err
	}

	return nil
}

func (c *AWSClient) GetCluster(ctx context.Context, args v1.GetClusterArgs) (*v1.Cluster, error) {
	eksClient := eks.NewFromConfig(c.awsConfig)

	eksCluster, err := eksClient.DescribeCluster(ctx, &eks.DescribeClusterInput{
		Name: aws.String(string(args.ID)),
	})
	if err != nil {
		return nil, errors.WrapAndTrace(err)
	}

	subnetIDs := make([]v1.CloudProviderResourceID, 0, len(eksCluster.Cluster.ResourcesVpcConfig.SubnetIds))
	for i, subnetID := range eksCluster.Cluster.ResourcesVpcConfig.SubnetIds {
		subnetIDs[i] = v1.CloudProviderResourceID(subnetID)
	}

	nodeGroups, err := c.getClusterNodeGroups(ctx, eksClient, eksCluster.Cluster)
	if err != nil {
		return nil, errors.WrapAndTrace(err)
	}

	return &v1.Cluster{
		RefID:             eksCluster.Cluster.Tags[tagBrevRefID],
		ID:                v1.CloudProviderResourceID(*eksCluster.Cluster.Arn),
		Name:              *eksCluster.Cluster.Name,
		APIEndpoint:       *eksCluster.Cluster.Endpoint,
		KubernetesVersion: *eksCluster.Cluster.Version,
		Status:            parseEKSClusterStatus(eksCluster.Cluster.Status),
		VPCID:             v1.CloudProviderResourceID(*eksCluster.Cluster.ResourcesVpcConfig.VpcId),
		SubnetIDs:         subnetIDs,
		NodeGroups:        nodeGroups,
		Tags:              v1.Tags(eksCluster.Cluster.Tags),
	}, nil
}

func parseEKSClusterStatus(status ekstypes.ClusterStatus) v1.ClusterStatus {
	switch status {
	case ekstypes.ClusterStatusCreating:
		return v1.ClusterStatusPending
	case ekstypes.ClusterStatusActive:
		return v1.ClusterStatusAvailable
	case ekstypes.ClusterStatusDeleting:
		return v1.ClusterStatusDeleting
	case ekstypes.ClusterStatusFailed:
		return v1.ClusterStatusFailed
	case ekstypes.ClusterStatusUpdating:
		return v1.ClusterStatusPending
	case ekstypes.ClusterStatusPending:
		return v1.ClusterStatusPending
	}
	return v1.ClusterStatusUnknown
}

func (c *AWSClient) getClusterNodeGroups(ctx context.Context, eksClient *eks.Client, eksCluster *ekstypes.Cluster) ([]*v1.NodeGroup, error) {
	eksNodeGroupNames, err := eksClient.ListNodegroups(ctx, &eks.ListNodegroupsInput{
		ClusterName: eksCluster.Name,
	})
	if err != nil {
		return nil, errors.WrapAndTrace(err)
	}

	nodeGroups := make([]*v1.NodeGroup, 0, len(eksNodeGroupNames.Nodegroups))
	for _, eksNodeGroupName := range eksNodeGroupNames.Nodegroups {
		eksNodeGroup, err := eksClient.DescribeNodegroup(ctx, &eks.DescribeNodegroupInput{
			ClusterName:   eksCluster.Name,
			NodegroupName: aws.String(eksNodeGroupName),
		})
		if err != nil {
			return nil, errors.WrapAndTrace(err)
		}

		nodeGroups = append(nodeGroups, parseEKSNodeGroup(eksNodeGroup.Nodegroup))
	}

	return nodeGroups, nil
}

func parseEKSNodeGroup(eksNodeGroup *ekstypes.Nodegroup) *v1.NodeGroup {
	return &v1.NodeGroup{
		ID:           v1.CloudProviderResourceID(*eksNodeGroup.NodegroupArn),
		RefID:        eksNodeGroup.Tags[tagBrevRefID],
		Name:         *eksNodeGroup.NodegroupName,
		MinNodeCount: int(*eksNodeGroup.ScalingConfig.MinSize),
		MaxNodeCount: int(*eksNodeGroup.ScalingConfig.MaxSize),
		InstanceType: eksNodeGroup.InstanceTypes[0], // todo: handle multiple instance types
		DiskSizeGiB:  int(*eksNodeGroup.DiskSize),
		Status:       parseEKSNodeGroupStatus(eksNodeGroup.Status),
		Tags:         v1.Tags(eksNodeGroup.Tags),
	}
}

func parseEKSNodeGroupStatus(status ekstypes.NodegroupStatus) v1.NodeGroupStatus {
	switch status {
	case ekstypes.NodegroupStatusCreating:
		return v1.NodeGroupStatusPending
	case ekstypes.NodegroupStatusActive:
		return v1.NodeGroupStatusAvailable
	case ekstypes.NodegroupStatusDeleting:
		return v1.NodeGroupStatusDeleting
	case ekstypes.NodegroupStatusCreateFailed:
		return v1.NodeGroupStatusFailed
	case ekstypes.NodegroupStatusDeleteFailed:
		return v1.NodeGroupStatusFailed
	}
	return v1.NodeGroupStatusUnknown
}

func (c *AWSClient) CreateNodeGroup(ctx context.Context, args v1.CreateNodeGroupArgs) (*v1.NodeGroup, error) {
	eksClient := eks.NewFromConfig(c.awsConfig)
	iamClient := iam.NewFromConfig(c.awsConfig)

	cluster, err := c.GetCluster(ctx, v1.GetClusterArgs{
		ID: args.ClusterID,
	})
	if err != nil {
		return nil, errors.WrapAndTrace(err)
	}

	subnetIDs := make([]string, len(cluster.SubnetIDs))
	for i, subnetID := range cluster.SubnetIDs {
		subnetIDs[i] = string(subnetID)
	}

	nodeRoleARN, err := createNodeRole(ctx, iamClient, cluster, args)
	if err != nil {
		return nil, errors.WrapAndTrace(err)
	}

	tags := make(map[string]string)
	for key, value := range args.Tags {
		tags[key] = value
	}
	tags[tagName] = args.Name
	tags[tagBrevRefID] = args.RefID
	tags[tagCreatedBy] = tagBrevCloudSDK

	output, err := eksClient.CreateNodegroup(ctx, &eks.CreateNodegroupInput{
		ClusterName:   aws.String(cluster.Name),
		NodegroupName: aws.String(args.Name),
		NodeRole:      aws.String(nodeRoleARN),
		ScalingConfig: &ekstypes.NodegroupScalingConfig{
			MinSize: aws.Int32(int32(args.MinNodeCount)),
			MaxSize: aws.Int32(int32(args.MaxNodeCount)),
		},
		DiskSize: aws.Int32(int32(args.DiskSizeGiB)),
		Subnets:  subnetIDs,
		InstanceTypes: []string{
			args.InstanceType,
		},
		Tags: tags,
	})
	if err != nil {
		return nil, errors.WrapAndTrace(err)
	}

	return parseEKSNodeGroup(output.Nodegroup), nil
}

func createNodeRole(ctx context.Context, iamClient *iam.Client, cluster *v1.Cluster, args v1.CreateNodeGroupArgs) (string, error) {
	roleName := fmt.Sprintf("%s-node-role", args.RefID)

	// Convert the tags to AWS tags
	tags := make(map[string]string)
	for key, value := range args.Tags {
		tags[key] = value
	}

	// Add the required tags
	tags[tagName] = args.Name
	tags[tagBrevRefID] = args.RefID
	tags[tagCreatedBy] = tagBrevCloudSDK

	iamTags := makeIAMTags(tags)

	iamPath := fmt.Sprintf("%s/%s/nodegroups/%s/", iamRolePathPrefix, cluster.RefID, args.RefID)
	iamPath = strings.ReplaceAll(iamPath, "[^a-zA-Z0-9/]", "")

	// Create the role
	input := &iam.CreateRoleInput{
		RoleName:    aws.String(roleName),
		Description: aws.String("Role for EKS node group"),
		Path:        aws.String(iamPath),
		AssumeRolePolicyDocument: aws.String(`{
		"Version": "2012-10-17",
		"Statement": [
			{
				"Effect": "Allow",
				"Principal": {
					"Service": "ec2.amazonaws.com"
				},
				"Action": "sts:AssumeRole"
			}
		]
	}`),
		Tags: iamTags,
	}
	output, err := iamClient.CreateRole(ctx, input)
	if err != nil {
		return "", errors.WrapAndTrace(err)
	}

	// Attach the required managed policies to the role
	managedPolicies := []string{
		"arn:aws:iam::aws:policy/AmazonEC2ContainerRegistryReadOnly",
		"arn:aws:iam::aws:policy/AmazonEKSWorkerNodePolicy",
		"arn:aws:iam::aws:policy/AmazonEKS_CNI_Policy",
	}
	for _, policyArn := range managedPolicies {
		_, err = iamClient.AttachRolePolicy(ctx, &iam.AttachRolePolicyInput{
			RoleName:  aws.String(roleName),
			PolicyArn: aws.String(policyArn),
		})
		if err != nil {
			return "", errors.WrapAndTrace(err)
		}
	}

	return *output.Role.Arn, nil
}

func (c *AWSClient) PutUser(_ context.Context, _ v1.PutUserArgs) (*v1.PutUserResponse, error) {
	panic("unimplemented")
}

func (c *AWSClient) DeleteCluster(ctx context.Context, args v1.DeleteClusterArgs) error {
	eksClient := eks.NewFromConfig(c.awsConfig)
	iamClient := iam.NewFromConfig(c.awsConfig)

	cluster, err := c.GetCluster(ctx, v1.GetClusterArgs{
		ID: args.ID,
	})
	if err != nil {
		return errors.WrapAndTrace(err)
	}

	_, err = eksClient.DeleteCluster(ctx, &eks.DeleteClusterInput{
		Name: aws.String(cluster.Name),
	})
	if err != nil {
		return errors.WrapAndTrace(err)
	}

	// Get the roles associated with the cluster
	roles, err := iamClient.ListRoles(ctx, &iam.ListRolesInput{
		PathPrefix: aws.String(iamRolePathPrefix),
	})
	if err != nil {
		return errors.WrapAndTrace(err)
	}

	for _, role := range roles.Roles {
		_, err = iamClient.DeleteRole(ctx, &iam.DeleteRoleInput{
			RoleName: role.RoleName,
		})
		if err != nil {
			return errors.WrapAndTrace(err)
		}
	}
	return nil
}

func makeIAMTags(tags map[string]string) []iamtypes.Tag {
	iamTags := make([]iamtypes.Tag, 0, len(tags))
	for key, value := range tags {
		iamTags = append(iamTags, iamtypes.Tag{Key: aws.String(key), Value: aws.String(value)})
	}
	return iamTags
}
