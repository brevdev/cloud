package v1

import (
	"context"
	"encoding/base64"
	"fmt"
	"math"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/eks"
	ekstypes "github.com/aws/aws-sdk-go-v2/service/eks/types"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	iamtypes "github.com/aws/aws-sdk-go-v2/service/iam/types"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	k8scmd "k8s.io/client-go/tools/clientcmd/api"
	"sigs.k8s.io/aws-iam-authenticator/pkg/token"

	"github.com/brevdev/cloud/internal/errors"
	cloudk8s "github.com/brevdev/cloud/internal/kubernetes"
	"github.com/brevdev/cloud/internal/rsa"
	v1 "github.com/brevdev/cloud/v1"
)

var (
	errUsernameIsRequired     = fmt.Errorf("username is required")
	errRoleIsRequired         = fmt.Errorf("role is required")
	errClusterIDIsRequired    = fmt.Errorf("cluster ID is required")
	errRSAPEMBase64IsRequired = fmt.Errorf("RSA PEM base64 is required")

	errNodeGroupMinNodeCountMustBeGreaterThan0                     = fmt.Errorf("node group minNodeCount must be greater than 0")
	errNodeGroupMaxNodeCountMustBeGreaterThan0                     = fmt.Errorf("node group maxNodeCount must be greater than 0")
	errNodeGroupMaxNodeCountMustBeGreaterThanOrEqualToMinNodeCount = fmt.Errorf("node group maxNodeCount must be greater than or equal to minNodeCount")
	errNodeGroupInstanceTypeIsRequired                             = fmt.Errorf("node group instanceType is required")
	errNodeGroupDiskSizeGiBMustBeGreaterThanOrEqualTo20            = fmt.Errorf("node group diskSizeGiB must be greater than or equal to 20")
	errNodeGroupDiskSizeGiBMustBeLessThanOrEqualToMaxInt32         = fmt.Errorf("node group diskSizeGiB must be less than or equal to %d", math.MaxInt32)
	errNodeGroupMaxNodeCountMustBeLessThanOrEqualToMaxInt32        = fmt.Errorf("node group maxNodeCount must be less than or equal to %d", math.MaxInt32)
	errNodeGroupMinNodeCountMustBeLessThanOrEqualToMaxInt32        = fmt.Errorf("node group minNodeCount must be less than or equal to %d", math.MaxInt32)
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
	for _, subnet := range vpc.GetSubnets() {
		subnetMap[string(subnet.GetID())] = subnet
	}

	// Get the target subnets from the map
	subnets := make([]*v1.Subnet, len(args.SubnetIDs))
	for i, subnetID := range args.SubnetIDs {
		if _, ok := subnetMap[string(subnetID)]; !ok {
			return nil, errors.WrapAndTrace(fmt.Errorf("subnet ID %s does not match VPC %s", subnetID, vpc.GetID()))
		} else {
			subnets[i] = subnetMap[string(subnetID)]
		}
	}

	// Create the cluster
	awsCluster, err := c.createEKSCluster(ctx, eksClient, iamClient, subnets, args)
	if err != nil {
		return nil, errors.WrapAndTrace(err)
	}

	brevCluster, err := v1.NewCluster(v1.ClusterSettings{
		ID: v1.CloudProviderResourceID(*awsCluster.Name),
		// ID:                v1.CloudProviderResourceID(*awsCluster.Arn), // todo: no API exists to fetch by ARN, so we may need to always use name as the ID
		Name:              *awsCluster.Name,
		RefID:             args.RefID,
		Provider:          CloudProviderID,
		Cloud:             CloudProviderID,
		Location:          c.region,
		VPCID:             vpc.GetID(),
		SubnetIDs:         args.SubnetIDs,
		KubernetesVersion: args.KubernetesVersion,
		Status:            v1.ClusterStatusPending,
		Tags:              args.Tags,
	})
	if err != nil {
		return nil, errors.WrapAndTrace(err)
	}
	return brevCluster, nil
}

func (c *AWSClient) createEKSCluster(ctx context.Context, eksClient *eks.Client, iamClient *iam.Client, subnets []*v1.Subnet, args v1.CreateClusterArgs) (*ekstypes.Cluster, error) {
	serviceRole, err := c.createServiceRole(ctx, iamClient, args)
	if err != nil {
		return nil, errors.WrapAndTrace(err)
	}

	eksCluster, err := c.createCluster(ctx, eksClient, args, serviceRole, subnets)
	if err != nil {
		return nil, errors.WrapAndTrace(err)
	}

	err = c.installEKSAddons(ctx, eksClient, eksCluster)
	if err != nil {
		return nil, errors.WrapAndTrace(err)
	}

	return eksCluster, nil
}

func (c *AWSClient) createServiceRole(ctx context.Context, iamClient *iam.Client, args v1.CreateClusterArgs) (*iamtypes.Role, error) {
	serviceRoleName := fmt.Sprintf("%s-service-role", args.RefID)

	c.logger.Debug(ctx, "creating service role", v1.Field{Key: "name", Value: serviceRoleName})

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
		return nil, errors.WrapAndTrace(err)
	}

	// Attach the AmazonEKSClusterPolicy to the role
	_, err = iamClient.AttachRolePolicy(ctx, &iam.AttachRolePolicyInput{
		RoleName:  aws.String(serviceRoleName),
		PolicyArn: aws.String("arn:aws:iam::aws:policy/AmazonEKSClusterPolicy"),
	})
	if err != nil {
		return nil, errors.WrapAndTrace(err)
	}

	return output.Role, nil
}

func (c *AWSClient) createCluster(ctx context.Context, eksClient *eks.Client, args v1.CreateClusterArgs, serviceRole *iamtypes.Role, subnets []*v1.Subnet) (*ekstypes.Cluster, error) {
	c.logger.Debug(ctx, "creating cluster", v1.Field{Key: "name", Value: args.Name})

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
		subnetIDs[i] = string(subnet.GetID())
	}

	c.logger.Debug(ctx, "creating cluster",
		v1.Field{Key: "clusterName", Value: args.Name},
		v1.Field{Key: "kubernetesVersion", Value: args.KubernetesVersion},
		v1.Field{Key: "serviceRoleARN", Value: *serviceRole.Arn},
		v1.Field{Key: "subnetIDs", Value: subnetIDs},
		v1.Field{Key: "tags", Value: tags},
	)
	input := &eks.CreateClusterInput{
		Name:    aws.String(args.Name),
		Version: aws.String(args.KubernetesVersion),
		RoleArn: aws.String(*serviceRole.Arn),
		ResourcesVpcConfig: &ekstypes.VpcConfigRequest{
			SubnetIds: subnetIDs,
		},
		AccessConfig: &ekstypes.CreateAccessConfigRequest{
			AuthenticationMode: ekstypes.AuthenticationModeApiAndConfigMap,
		},
		Tags: tags,
	}

	output, err := eksClient.CreateCluster(ctx, input)
	if err != nil {
		return nil, errors.WrapAndTrace(err)
	}

	return output.Cluster, nil
}

func (c *AWSClient) installEKSAddons(ctx context.Context, eksClient *eks.Client, eksCluster *ekstypes.Cluster) error {
	err := c.installEKSAddon(ctx, eksClient, eksCluster, "vpc-cni")
	if err != nil {
		return err
	}

	err = c.installEKSAddon(ctx, eksClient, eksCluster, "eks-pod-identity-agent")
	if err != nil {
		return err
	}

	return nil
}

func (c *AWSClient) installEKSAddon(ctx context.Context, eksClient *eks.Client, eksCluster *ekstypes.Cluster, addonName string) error {
	c.logger.Debug(ctx, "installing EKS addon",
		v1.Field{Key: "clusterName", Value: *eksCluster.Name},
		v1.Field{Key: "name", Value: addonName},
	)

	_, err := eksClient.CreateAddon(ctx, &eks.CreateAddonInput{
		ClusterName: eksCluster.Name,
		AddonName:   aws.String(addonName),
	})
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
		var noSuchEntityError *ekstypes.ResourceNotFoundException
		if errors.As(err, &noSuchEntityError) {
			return nil, v1.ErrResourceNotFound
		}
		return nil, errors.WrapAndTrace(err)
	}

	subnetIDs := make([]v1.CloudProviderResourceID, 0, len(eksCluster.Cluster.ResourcesVpcConfig.SubnetIds))
	for _, subnetID := range eksCluster.Cluster.ResourcesVpcConfig.SubnetIds {
		subnetIDs = append(subnetIDs, v1.CloudProviderResourceID(subnetID))
	}

	nodeGroups, err := c.getClusterNodeGroups(ctx, eksClient, eksCluster.Cluster)
	if err != nil {
		return nil, errors.WrapAndTrace(err)
	}

	// List all addons and use their status to determine if the cluster is ready
	addonNames, err := eksClient.ListAddons(ctx, &eks.ListAddonsInput{
		ClusterName: eksCluster.Cluster.Name,
	})
	if err != nil {
		return nil, errors.WrapAndTrace(err)
	}

	inactiveAddons := 0
	for _, name := range addonNames.Addons {
		addon, err := eksClient.DescribeAddon(ctx, &eks.DescribeAddonInput{
			ClusterName: eksCluster.Cluster.Name,
			AddonName:   aws.String(name),
		})
		if err != nil {
			return nil, errors.WrapAndTrace(err)
		}
		if addon.Addon.Status != ekstypes.AddonStatusActive {
			inactiveAddons++
		}
	}

	var clusterStatus v1.ClusterStatus
	if inactiveAddons > 0 {
		clusterStatus = v1.ClusterStatusPending
	} else {
		clusterStatus = parseEKSClusterStatus(eksCluster.Cluster.Status)
	}

	brevCluster, err := v1.NewCluster(v1.ClusterSettings{
		RefID: eksCluster.Cluster.Tags[tagBrevRefID],
		ID:    v1.CloudProviderResourceID(*eksCluster.Cluster.Name),
		// ID:                v1.CloudProviderResourceID(*eksCluster.Cluster.Arn), // todo: no API exists to fetch by ARN, so we may need to always use name as the ID
		Name:                       *eksCluster.Cluster.Name,
		KubernetesVersion:          *eksCluster.Cluster.Version,
		Status:                     clusterStatus,
		VPCID:                      v1.CloudProviderResourceID(*eksCluster.Cluster.ResourcesVpcConfig.VpcId),
		SubnetIDs:                  subnetIDs,
		NodeGroups:                 nodeGroups,
		ClusterCACertificateBase64: getClusterCACertificateBase64(eksCluster.Cluster),
		APIEndpoint:                getClusterAPIEndpoint(eksCluster.Cluster),
		Provider:                   CloudProviderID,
		Tags:                       v1.Tags(eksCluster.Cluster.Tags),
	})
	if err != nil {
		return nil, errors.WrapAndTrace(err)
	}
	return brevCluster, nil
}

func getClusterCACertificateBase64(cluster *ekstypes.Cluster) string {
	if cluster == nil ||
		cluster.CertificateAuthority == nil ||
		cluster.CertificateAuthority.Data == nil {
		return ""
	}
	return *cluster.CertificateAuthority.Data
}

func getClusterAPIEndpoint(cluster *ekstypes.Cluster) string {
	if cluster == nil ||
		cluster.Endpoint == nil {
		return ""
	}
	return *cluster.Endpoint
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
	// First fetch the names of all of the cluster's node groups
	eksNodeGroupNames, err := eksClient.ListNodegroups(ctx, &eks.ListNodegroupsInput{
		ClusterName: eksCluster.Name,
	})
	if err != nil {
		return nil, errors.WrapAndTrace(err)
	}

	// Then fetch the details of each node group
	nodeGroups := make([]*v1.NodeGroup, 0, len(eksNodeGroupNames.Nodegroups))
	for _, eksNodeGroupName := range eksNodeGroupNames.Nodegroups {
		eksNodeGroup, err := eksClient.DescribeNodegroup(ctx, &eks.DescribeNodegroupInput{
			ClusterName:   eksCluster.Name,
			NodegroupName: aws.String(eksNodeGroupName),
		})
		if err != nil {
			return nil, errors.WrapAndTrace(err)
		}

		brevNodeGroup, err := parseEKSNodeGroup(eksNodeGroup.Nodegroup)
		if err != nil {
			return nil, errors.WrapAndTrace(err)
		}
		nodeGroups = append(nodeGroups, brevNodeGroup)
	}

	return nodeGroups, nil
}

func parseEKSNodeGroup(eksNodeGroup *ekstypes.Nodegroup) (*v1.NodeGroup, error) {
	brevNodeGroup, err := v1.NewNodeGroup(v1.NodeGroupSettings{
		ID: v1.CloudProviderResourceID(*eksNodeGroup.NodegroupName),
		// ID:           v1.CloudProviderResourceID(*eksNodeGroup.NodegroupArn), // todo: no API exists to fetch by ARN, so we may need to always use name as the ID
		RefID:        eksNodeGroup.Tags[tagBrevRefID],
		Name:         *eksNodeGroup.NodegroupName,
		MinNodeCount: int(*eksNodeGroup.ScalingConfig.MinSize),
		MaxNodeCount: int(*eksNodeGroup.ScalingConfig.MaxSize),
		InstanceType: eksNodeGroup.InstanceTypes[0], // todo: handle multiple instance types
		DiskSize:     v1.NewBytes(v1.BytesValue(*eksNodeGroup.DiskSize), v1.Gibibyte),
		Status:       parseEKSNodeGroupStatus(eksNodeGroup.Status),
		Tags:         v1.Tags(eksNodeGroup.Tags),
	})
	if err != nil {
		return nil, errors.WrapAndTrace(err)
	}
	return brevNodeGroup, nil
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
	err := validateCreateNodeGroupArgs(args)
	if err != nil {
		return nil, errors.WrapAndTrace(err)
	}

	eksClient := eks.NewFromConfig(c.awsConfig)
	iamClient := iam.NewFromConfig(c.awsConfig)

	// Fetch the target cluster
	cluster, err := c.GetCluster(ctx, v1.GetClusterArgs{
		ID: args.ClusterID,
	})
	if err != nil {
		return nil, errors.WrapAndTrace(err)
	}

	// Convert the target cluster's subnet IDs to AWS subnet IDs
	subnetIDs := make([]string, len(cluster.GetSubnetIDs()))
	for i, subnetID := range cluster.GetSubnetIDs() {
		subnetIDs[i] = string(subnetID)
	}

	// Create the node role that will be attached to all nodes in the node group
	nodeRoleARN, err := c.createNodeRole(ctx, iamClient, cluster, args)
	if err != nil {
		return nil, errors.WrapAndTrace(err)
	}

	// Convert the tags to AWS tags
	tags := make(map[string]string)
	for key, value := range args.Tags {
		tags[key] = value
	}
	tags[tagName] = args.Name
	tags[tagBrevRefID] = args.RefID
	tags[tagCreatedBy] = tagBrevCloudSDK

	// Create the node group
	c.logger.Debug(ctx, "creating node group",
		v1.Field{Key: "clusterName", Value: cluster.GetName()},
		v1.Field{Key: "nodeGroupName", Value: args.Name},
	)
	output, err := eksClient.CreateNodegroup(ctx, &eks.CreateNodegroupInput{
		ClusterName:   aws.String(cluster.GetName()),
		NodegroupName: aws.String(args.Name),
		NodeRole:      aws.String(nodeRoleARN),
		ScalingConfig: &ekstypes.NodegroupScalingConfig{
			MinSize: aws.Int32(int32(args.MinNodeCount)), //nolint:gosec // checked in input validation
			MaxSize: aws.Int32(int32(args.MaxNodeCount)), //nolint:gosec // checked in input validation
		},
		DiskSize: aws.Int32(int32(args.DiskSize.ByteCount() / v1.Gibibyte.ByteCount())), //nolint:gosec // checked in input validation
		Subnets:  subnetIDs,
		InstanceTypes: []string{
			args.InstanceType,
		},
		Tags: tags,
	})
	if err != nil {
		return nil, errors.WrapAndTrace(err)
	}

	brevNodeGroup, err := parseEKSNodeGroup(output.Nodegroup)
	if err != nil {
		return nil, errors.WrapAndTrace(err)
	}
	return brevNodeGroup, nil
}

func validateCreateNodeGroupArgs(args v1.CreateNodeGroupArgs) error {
	errs := []error{}
	if args.MinNodeCount < 1 {
		errs = append(errs, errNodeGroupMinNodeCountMustBeGreaterThan0)
	}
	if args.MaxNodeCount < 1 {
		errs = append(errs, errNodeGroupMaxNodeCountMustBeGreaterThan0)
	}
	if args.MaxNodeCount < args.MinNodeCount {
		errs = append(errs, errNodeGroupMaxNodeCountMustBeGreaterThanOrEqualToMinNodeCount)
	}
	if args.InstanceType == "" {
		errs = append(errs, errNodeGroupInstanceTypeIsRequired)
	}
	if args.DiskSize.LessThan(v1.NewBytes(20, v1.Gibibyte)) {
		errs = append(errs, errNodeGroupDiskSizeGiBMustBeGreaterThanOrEqualTo20)
	}
	if args.DiskSize.GreaterThan(v1.NewBytes(math.MaxInt32, v1.Gibibyte)) {
		errs = append(errs, errNodeGroupDiskSizeGiBMustBeLessThanOrEqualToMaxInt32)
	}
	if args.MaxNodeCount > math.MaxInt32 {
		errs = append(errs, errNodeGroupMaxNodeCountMustBeLessThanOrEqualToMaxInt32)
	}
	if args.MinNodeCount > math.MaxInt32 {
		errs = append(errs, errNodeGroupMinNodeCountMustBeLessThanOrEqualToMaxInt32)
	}
	return errors.WrapAndTrace(errors.Join(errs...))
}

func (c *AWSClient) createNodeRole(ctx context.Context, iamClient *iam.Client, cluster *v1.Cluster, args v1.CreateNodeGroupArgs) (string, error) {
	roleName := fmt.Sprintf("%s-node-role", args.RefID)

	c.logger.Debug(ctx, "creating node role",
		v1.Field{Key: "clusterName", Value: cluster.GetName()},
		v1.Field{Key: "roleName", Value: roleName},
	)

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
	iamPath := getNodeGroupIAMRolePath(cluster.GetRefID(), args.RefID)

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

func (c *AWSClient) GetNodeGroup(ctx context.Context, args v1.GetNodeGroupArgs) (*v1.NodeGroup, error) {
	eksClient := eks.NewFromConfig(c.awsConfig)

	eksNodeGroup, err := eksClient.DescribeNodegroup(ctx, &eks.DescribeNodegroupInput{
		ClusterName:   aws.String(string(args.ClusterID)),
		NodegroupName: aws.String(string(args.ID)),
	})
	if err != nil {
		var noSuchEntityError *ekstypes.ResourceNotFoundException
		if errors.As(err, &noSuchEntityError) {
			return nil, v1.ErrResourceNotFound
		}
		return nil, errors.WrapAndTrace(err)
	}

	brevNodeGroup, err := parseEKSNodeGroup(eksNodeGroup.Nodegroup)
	if err != nil {
		return nil, errors.WrapAndTrace(err)
	}
	return brevNodeGroup, nil
}

func (c *AWSClient) ModifyNodeGroup(ctx context.Context, args v1.ModifyNodeGroupArgs) error {
	eksClient := eks.NewFromConfig(c.awsConfig)

	err := validateModifyNodeGroupArgs(args)
	if err != nil {
		return errors.WrapAndTrace(err)
	}

	cluster, err := c.GetCluster(ctx, v1.GetClusterArgs{
		ID: args.ClusterID,
	})
	if err != nil {
		return errors.WrapAndTrace(err)
	}

	nodeGroup, err := c.GetNodeGroup(ctx, v1.GetNodeGroupArgs{
		ClusterID: cluster.GetID(),
		ID:        args.ID,
	})
	if err != nil {
		return errors.WrapAndTrace(err)
	}

	_, err = eksClient.UpdateNodegroupConfig(ctx, &eks.UpdateNodegroupConfigInput{
		ClusterName:   aws.String(cluster.GetName()),
		NodegroupName: aws.String(nodeGroup.GetName()),
		ScalingConfig: &ekstypes.NodegroupScalingConfig{
			DesiredSize: aws.Int32(int32(args.MinNodeCount)), //nolint:gosec // checked in input validation
			MinSize:     aws.Int32(int32(args.MinNodeCount)), //nolint:gosec // checked in input validation
			MaxSize:     aws.Int32(int32(args.MaxNodeCount)), //nolint:gosec // checked in input validation
		},
	})
	if err != nil {
		return errors.WrapAndTrace(err)
	}
	return nil
}

func validateModifyNodeGroupArgs(args v1.ModifyNodeGroupArgs) error {
	errs := []error{}
	if args.MinNodeCount < 1 {
		errs = append(errs, errNodeGroupMinNodeCountMustBeGreaterThan0)
	}
	if args.MaxNodeCount < 1 {
		errs = append(errs, errNodeGroupMaxNodeCountMustBeGreaterThan0)
	}
	if args.MaxNodeCount < args.MinNodeCount {
		errs = append(errs, errNodeGroupMaxNodeCountMustBeGreaterThanOrEqualToMinNodeCount)
	}
	if args.MinNodeCount > math.MaxInt32 {
		errs = append(errs, errNodeGroupMinNodeCountMustBeLessThanOrEqualToMaxInt32)
	}
	if args.MaxNodeCount > math.MaxInt32 {
		errs = append(errs, errNodeGroupMaxNodeCountMustBeLessThanOrEqualToMaxInt32)
	}
	return errors.WrapAndTrace(errors.Join(errs...))
}

func (c *AWSClient) DeleteNodeGroup(ctx context.Context, args v1.DeleteNodeGroupArgs) error {
	eksClient := eks.NewFromConfig(c.awsConfig)
	iamClient := iam.NewFromConfig(c.awsConfig)

	// Fetch the target cluster
	cluster, err := c.GetCluster(ctx, v1.GetClusterArgs{
		ID: args.ClusterID,
	})
	if err != nil {
		return errors.WrapAndTrace(err)
	}

	// Fetch the target node group
	nodeGroup, err := c.GetNodeGroup(ctx, v1.GetNodeGroupArgs{
		ClusterID: cluster.GetID(),
		ID:        args.ID,
	})
	if err != nil {
		return errors.WrapAndTrace(err)
	}

	// Get the roles associated with the node group
	iamPath := getNodeGroupIAMRolePath(cluster.GetRefID(), nodeGroup.GetRefID())
	roles, err := iamClient.ListRoles(ctx, &iam.ListRolesInput{
		PathPrefix: aws.String(iamPath),
	})
	if err != nil {
		return errors.WrapAndTrace(err)
	}

	c.logger.Debug(ctx, fmt.Sprintf("found %d roles associated with node group", len(roles.Roles)),
		v1.Field{Key: "clusterName", Value: cluster.GetName()},
		v1.Field{Key: "nodeGroupName", Value: nodeGroup.GetName()},
	)

	// Delete the roles associated with the node group
	for _, role := range roles.Roles {
		c.logger.Debug(ctx, "removing role from EKS access entries", v1.Field{Key: "roleName", Value: *role.RoleName})

		// Remove roles from EKS access entries
		_, err = eksClient.DeleteAccessEntry(ctx, &eks.DeleteAccessEntryInput{
			ClusterName:  aws.String(cluster.GetName()),
			PrincipalArn: role.Arn,
		})
		if err != nil {
			return errors.WrapAndTrace(err)
		}

		// Delete the role
		err = c.deleteRole(ctx, iamClient, role)
		if err != nil {
			return errors.WrapAndTrace(err)
		}
	}

	// Delete the node group
	c.logger.Debug(ctx, "deleting node group",
		v1.Field{Key: "clusterName", Value: cluster.GetName()},
		v1.Field{Key: "nodeGroupName", Value: nodeGroup.GetName()},
	)
	_, err = eksClient.DeleteNodegroup(ctx, &eks.DeleteNodegroupInput{
		ClusterName:   aws.String(cluster.GetName()),
		NodegroupName: aws.String(nodeGroup.GetName()),
	})
	if err != nil {
		return errors.WrapAndTrace(err)
	}

	return nil
}

// TODO: AWS EKS only supports IAM or OIDC authentication.
func (c *AWSClient) SetClusterUser(ctx context.Context, args v1.SetClusterUserArgs) (*v1.ClusterUser, error) {
	err := validatePutUserArgs(args)
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
	certificateAuthorityData, err := base64.StdEncoding.DecodeString(cluster.GetClusterCACertificateBase64())
	if err != nil {
		return nil, errors.WrapAndTrace(fmt.Errorf("failed to decode certificate authority data: %w", err))
	}

	// Generate the complete kubeconfig
	kubeconfigBytes, err := clientcmd.Write(k8scmd.Config{
		Kind:       "Config",
		APIVersion: "v1",
		Clusters: map[string]*k8scmd.Cluster{
			cluster.GetRefID(): {
				Server:                   cluster.GetAPIEndpoint(),
				CertificateAuthorityData: certificateAuthorityData,
			},
		},
		AuthInfos: map[string]*k8scmd.AuthInfo{
			cluster.GetRefID(): {
				ClientCertificateData: signedCertificate,
				ClientKeyData:         privateKeyBytes,
			},
		},
	})
	if err != nil {
		return nil, errors.WrapAndTrace(fmt.Errorf("failed to write kubeconfig: %w", err))
	}

	brevClusterUser, err := v1.NewClusterUser(v1.ClusterUserSettings{
		ClusterName:                           cluster.GetRefID(),
		ClusterCertificateAuthorityDataBase64: cluster.GetClusterCACertificateBase64(),
		ClusterServerURL:                      cluster.GetAPIEndpoint(),
		Username:                              args.Username,
		UserClientCertificateDataBase64:       base64.StdEncoding.EncodeToString(signedCertificate),
		UserClientKeyDataBase64:               base64.StdEncoding.EncodeToString(privateKeyBytes),
		KubeconfigBase64:                      base64.StdEncoding.EncodeToString(kubeconfigBytes),
	})
	if err != nil {
		return nil, errors.WrapAndTrace(fmt.Errorf("failed to create cluster user: %w", err))
	}
	return brevClusterUser, nil
}

func validatePutUserArgs(args v1.SetClusterUserArgs) error {
	errs := []error{}
	if args.Username == "" {
		errs = append(errs, errUsernameIsRequired)
	}
	if args.Role == "" {
		errs = append(errs, errRoleIsRequired)
	}
	if args.ClusterID == "" {
		errs = append(errs, errClusterIDIsRequired)
	}
	if args.RSAPEMBase64 == "" {
		errs = append(errs, errRSAPEMBase64IsRequired)
	}
	return errors.WrapAndTrace(errors.Join(errs...))
}

func (c *AWSClient) DeleteCluster(ctx context.Context, args v1.DeleteClusterArgs) error {
	eksClient := eks.NewFromConfig(c.awsConfig)
	iamClient := iam.NewFromConfig(c.awsConfig)

	// Fetch the target cluster
	cluster, err := c.GetCluster(ctx, v1.GetClusterArgs{ //nolint:staticcheck // prefer explicit struct literal
		ID: args.ID,
	})
	if err != nil {
		return errors.WrapAndTrace(err)
	}

	// Delete the cluster
	c.logger.Debug(ctx, "deleting cluster", v1.Field{Key: "clusterName", Value: cluster.GetName()})
	_, err = eksClient.DeleteCluster(ctx, &eks.DeleteClusterInput{
		Name: aws.String(cluster.GetName()),
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

	c.logger.Debug(ctx, fmt.Sprintf("found %d roles associated with cluster", len(roles.Roles)),
		v1.Field{Key: "clusterName", Value: cluster.GetName()},
	)

	// Delete the roles associated with the cluster
	for _, role := range roles.Roles {
		err = c.deleteRole(ctx, iamClient, role)
		if err != nil {
			return errors.WrapAndTrace(err)
		}
	}

	return nil
}

func (c *AWSClient) deleteRole(ctx context.Context, iamClient *iam.Client, role iamtypes.Role) error {
	// Get the instance profiles associated with the role -- these are created as a side effect of attachment to a node (EC2 instance)
	instanceProfiles, err := iamClient.ListInstanceProfilesForRole(ctx, &iam.ListInstanceProfilesForRoleInput{
		RoleName: role.RoleName,
	})
	if err != nil {
		return errors.WrapAndTrace(err)
	}

	c.logger.Debug(ctx, fmt.Sprintf("found %d instance profiles associated with role", len(instanceProfiles.InstanceProfiles)),
		v1.Field{Key: "roleName", Value: *role.RoleName},
	)
	// Remove the role from the instance profiles
	for _, instanceProfile := range instanceProfiles.InstanceProfiles {
		c.logger.Debug(ctx, "removing role from instance profile",
			v1.Field{Key: "instanceProfileName", Value: *instanceProfile.InstanceProfileName},
			v1.Field{Key: "roleName", Value: *role.RoleName},
		)
		_, err = iamClient.RemoveRoleFromInstanceProfile(ctx, &iam.RemoveRoleFromInstanceProfileInput{
			InstanceProfileName: instanceProfile.InstanceProfileName,
			RoleName:            role.RoleName,
		})
		if err != nil {
			return errors.WrapAndTrace(err)
		}

		// Delete the instance profile
		c.logger.Debug(ctx, "deleting instance profile",
			v1.Field{Key: "instanceProfileName", Value: *instanceProfile.InstanceProfileName},
			v1.Field{Key: "roleName", Value: *role.RoleName},
		)
		_, err = iamClient.DeleteInstanceProfile(ctx, &iam.DeleteInstanceProfileInput{
			InstanceProfileName: instanceProfile.InstanceProfileName,
		})
		if err != nil {
			return errors.WrapAndTrace(err)
		}
	}

	// Detach the policies from the role
	attachedPolicies, err := iamClient.ListAttachedRolePolicies(ctx, &iam.ListAttachedRolePoliciesInput{
		RoleName: role.RoleName,
	})
	if err != nil {
		return errors.WrapAndTrace(err)
	}

	c.logger.Debug(ctx, fmt.Sprintf("found %d policies associated with role", len(attachedPolicies.AttachedPolicies)),
		v1.Field{Key: "roleName", Value: *role.RoleName},
	)
	for _, policy := range attachedPolicies.AttachedPolicies {
		c.logger.Debug(ctx, "detaching policy from role",
			v1.Field{Key: "policyArn", Value: *policy.PolicyArn},
			v1.Field{Key: "roleName", Value: *role.RoleName},
		)
		_, err = iamClient.DetachRolePolicy(ctx, &iam.DetachRolePolicyInput{
			RoleName:  role.RoleName,
			PolicyArn: policy.PolicyArn,
		})
		if err != nil {
			return errors.WrapAndTrace(err)
		}
	}

	// Delete the role
	c.logger.Debug(ctx, "deleting role", v1.Field{Key: "roleName", Value: *role.RoleName})
	_, err = iamClient.DeleteRole(ctx, &iam.DeleteRoleInput{
		RoleName: role.RoleName,
	})
	if err != nil {
		return errors.WrapAndTrace(err)
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

func getNodeGroupIAMRolePath(clusterRefID string, nodeGroupRefID string) string {
	iamPath := fmt.Sprintf("%s/%s/nodegroups/%s/", iamRolePathPrefix, clusterRefID, nodeGroupRefID)
	iamPath = strings.ReplaceAll(iamPath, "[^a-zA-Z0-9/]", "")
	return iamPath
}

func (c *AWSClient) newK8sClient(ctx context.Context, cluster *v1.Cluster) (*kubernetes.Clientset, error) {
	newK8sConfig, err := c.newK8sConfig(ctx, cluster)
	if err != nil {
		return nil, errors.WrapAndTrace(fmt.Errorf("failed to create k8s config: %w", err))
	}

	// Create a clientset to interact with the cluster using the bearer token and CA certificate
	clientset, err := kubernetes.NewForConfig(newK8sConfig)
	if err != nil {
		return nil, errors.WrapAndTrace(fmt.Errorf("failed to create clientset: %w", err))
	}

	return clientset, nil
}

func (c *AWSClient) newK8sConfig(ctx context.Context, cluster *v1.Cluster) (*rest.Config, error) {
	// Decode the cluster CA certificate
	clusterCACertificate, err := base64.StdEncoding.DecodeString(cluster.GetClusterCACertificateBase64())
	if err != nil {
		return nil, errors.WrapAndTrace(fmt.Errorf("failed to decode cluster CA certificate: %w", err))
	}

	// Get a bearer token to authenticate to the cluster
	forwardSessionName := true
	cache := false
	tokenGenerator, err := token.NewGenerator(forwardSessionName, cache)
	if err != nil {
		return nil, errors.WrapAndTrace(fmt.Errorf("failed to create token generator: %w", err))
	}

	token, err := tokenGenerator.GetWithOptions(ctx, &token.GetTokenOptions{
		ClusterID: cluster.GetName(),
	})
	if err != nil {
		return nil, errors.WrapAndTrace(fmt.Errorf("failed to generate token: %w", err))
	}

	return &rest.Config{
		Host:        cluster.GetAPIEndpoint(),
		BearerToken: token.Token,
		TLSClientConfig: rest.TLSClientConfig{
			CAData: clusterCACertificate,
		},
	}, nil
}
