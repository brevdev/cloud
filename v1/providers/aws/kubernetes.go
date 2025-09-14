package v1

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/eks"
	ekstypes "github.com/aws/aws-sdk-go-v2/service/eks/types"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	iamtypes "github.com/aws/aws-sdk-go-v2/service/iam/types"

	v1 "github.com/brevdev/cloud/v1"
)

var _ v1.CloudMaintainKubernetes = &AWSClient{}

func (c *AWSClient) CreateCluster(ctx context.Context, args v1.CreateClusterArgs) (*v1.Cluster, error) {
	// Create the AWS client in the specified region
	awsClient := awsClient{
		eksClient: eks.NewFromConfig(c.awsConfig, func(o *eks.Options) {
			o.Region = args.Location
		}),
		iamClient: iam.NewFromConfig(c.awsConfig, func(o *iam.Options) {
			o.Region = args.Location
		}),
	}

	// Create the cluster
	awsCluster, err := createEKSCluster(ctx, awsClient, args)
	if err != nil {
		return nil, err
	}

	return &v1.Cluster{
		RefID:  args.RefID,
		Name:   *awsCluster.Name,
		Status: v1.ClusterStatusAvailable,
		// NodeGroups: args.NodeGroups,
	}, nil
}

func createEKSCluster(ctx context.Context, awsClient awsClient, args v1.CreateClusterArgs) (*ekstypes.Cluster, error) {
	serviceRoleARN, err := getOrCreateServiceRoleARN(ctx, awsClient, args)
	if err != nil {
		return nil, err
	}

	eksCluster, err := createCluster(ctx, awsClient, args, serviceRoleARN)
	if err != nil {
		return nil, err
	}

	err = installEKSAddons(ctx, awsClient, eksCluster)
	if err != nil {
		return nil, err
	}

	return eksCluster, nil
}

func getOrCreateServiceRoleARN(ctx context.Context, awsClient awsClient, args v1.CreateClusterArgs) (string, error) {
	serviceRoleName := fmt.Sprintf("%s-service-role", args.RefID)

	// Get and return the role if it exists
	serviceRole, err := awsClient.iamClient.GetRole(ctx, &iam.GetRoleInput{
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

	// Create the role
	input := &iam.CreateRoleInput{
		RoleName:    aws.String(serviceRoleName),
		Description: aws.String("Role for EKS cluster"),
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
		Tags: []iamtypes.Tag{
			{
				Key:   aws.String(tagBrevRefID),
				Value: aws.String(args.RefID),
			},
			{
				Key:   aws.String("CreatedBy"),
				Value: aws.String(tagBrevCloudSDK),
			},
		},
	}
	output, err := awsClient.iamClient.CreateRole(ctx, input)
	if err != nil {
		return "", err
	}

	// Attach the AmazonEKSClusterPolicy to the role
	_, err = awsClient.iamClient.AttachRolePolicy(ctx, &iam.AttachRolePolicyInput{
		RoleName:  aws.String(serviceRoleName),
		PolicyArn: aws.String("arn:aws:iam::aws:policy/AmazonEKSClusterPolicy"),
	})
	if err != nil {
		return "", err
	}

	return *output.Role.Arn, nil
}

func createCluster(ctx context.Context, awsClient awsClient, args v1.CreateClusterArgs, serviceRoleARN string) (*ekstypes.Cluster, error) {
	tags := map[string]string{
		"Name":       args.Name,
		tagBrevRefID: args.RefID,
		"CreatedBy":  tagBrevCloudSDK,
	}

	input := &eks.CreateClusterInput{
		Name:    aws.String(args.Name),
		Version: aws.String(args.KubernetesVersion),
		RoleArn: aws.String(serviceRoleARN),
		ResourcesVpcConfig: &ekstypes.VpcConfigRequest{
			SubnetIds: args.SubnetIDs,
		},
		Tags: tags,
	}

	output, err := awsClient.eksClient.CreateCluster(ctx, input)
	if err != nil {
		return nil, err
	}

	// Wait for the cluster to be active
	w := eks.NewClusterActiveWaiter(awsClient.eksClient, func(o *eks.ClusterActiveWaiterOptions) {
		o.MaxDelay = 30 * time.Second
		o.MinDelay = 10 * time.Second
	})
	err = w.Wait(ctx, &eks.DescribeClusterInput{
		Name: output.Cluster.Name,
	}, 10*time.Minute)
	if err != nil {
		return nil, err
	}

	return output.Cluster, nil
}

func installEKSAddons(ctx context.Context, awsClient awsClient, eksCluster *ekstypes.Cluster) error {
	err := installEKSAddon(ctx, awsClient, eksCluster, "vpc-cni")
	if err != nil {
		return err
	}

	err = installEKSAddon(ctx, awsClient, eksCluster, "eks-pod-identity-agent")
	if err != nil {
		return err
	}

	return nil
}

func installEKSAddon(ctx context.Context, awsClient awsClient, eksCluster *ekstypes.Cluster, addonName string) error {
	_, err := awsClient.eksClient.CreateAddon(ctx, &eks.CreateAddonInput{
		ClusterName: eksCluster.Name,
		AddonName:   aws.String(addonName),
	})
	if err != nil {
		return err
	}

	// Wait for the addon to be created
	w := eks.NewAddonActiveWaiter(awsClient.eksClient, func(o *eks.AddonActiveWaiterOptions) {
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
	return nil, nil
}

func (c *AWSClient) DeleteCluster(ctx context.Context, args v1.DeleteClusterArgs) error {
	return nil
}
