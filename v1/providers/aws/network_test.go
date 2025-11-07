package v1

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"

	v1 "github.com/brevdev/cloud/v1"
)

func TestFilterSubnetArgs(t *testing.T) {
	tests := []struct {
		name            string
		subnets         []v1.CreateSubnetArgs
		subnetType      v1.SubnetType
		expectedSubnets []v1.CreateSubnetArgs
	}{
		{
			name: "filter public subnets",
			subnets: []v1.CreateSubnetArgs{
				{CidrBlock: "10.0.0.0/24", Type: v1.SubnetTypePublic},
				{CidrBlock: "10.0.1.0/24", Type: v1.SubnetTypePrivate},
				{CidrBlock: "10.0.2.0/24", Type: v1.SubnetTypePublic},
			},
			subnetType: v1.SubnetTypePublic,
			expectedSubnets: []v1.CreateSubnetArgs{
				{CidrBlock: "10.0.0.0/24", Type: v1.SubnetTypePublic},
				{CidrBlock: "10.0.2.0/24", Type: v1.SubnetTypePublic},
			},
		},
		{
			name: "filter private subnets",
			subnets: []v1.CreateSubnetArgs{
				{CidrBlock: "10.0.0.0/24", Type: v1.SubnetTypePublic},
				{CidrBlock: "10.0.1.0/24", Type: v1.SubnetTypePrivate},
				{CidrBlock: "10.0.2.0/24", Type: v1.SubnetTypePublic},
			},
			subnetType: v1.SubnetTypePrivate,
			expectedSubnets: []v1.CreateSubnetArgs{
				{CidrBlock: "10.0.1.0/24", Type: v1.SubnetTypePrivate},
			},
		},
		{
			name: "no matching subnets",
			subnets: []v1.CreateSubnetArgs{
				{CidrBlock: "10.0.0.0/24", Type: v1.SubnetTypePublic},
				{CidrBlock: "10.0.2.0/24", Type: v1.SubnetTypePublic},
			},
			subnetType:      v1.SubnetTypePrivate,
			expectedSubnets: []v1.CreateSubnetArgs{},
		},
		{
			name:            "empty input",
			subnets:         []v1.CreateSubnetArgs{},
			subnetType:      v1.SubnetTypePublic,
			expectedSubnets: []v1.CreateSubnetArgs{},
		},
		{
			name: "all subnets match",
			subnets: []v1.CreateSubnetArgs{
				{CidrBlock: "10.0.0.0/24", Type: v1.SubnetTypePublic},
				{CidrBlock: "10.0.1.0/24", Type: v1.SubnetTypePublic},
				{CidrBlock: "10.0.2.0/24", Type: v1.SubnetTypePublic},
			},
			subnetType: v1.SubnetTypePublic,
			expectedSubnets: []v1.CreateSubnetArgs{
				{CidrBlock: "10.0.0.0/24", Type: v1.SubnetTypePublic},
				{CidrBlock: "10.0.1.0/24", Type: v1.SubnetTypePublic},
				{CidrBlock: "10.0.2.0/24", Type: v1.SubnetTypePublic},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := filterSubnetArgs(tt.subnets, tt.subnetType)

			if len(result) != len(tt.expectedSubnets) {
				t.Errorf("expected %d subnets, got %d", len(tt.expectedSubnets), len(result))
			}

			for i, subnet := range result {
				if subnet.CidrBlock != tt.expectedSubnets[i].CidrBlock {
					t.Errorf("expected subnet CIDR %s, got %s", tt.expectedSubnets[i].CidrBlock, subnet.CidrBlock)
				}
				if subnet.Type != tt.expectedSubnets[i].Type {
					t.Errorf("expected subnet type %v, got %v", tt.expectedSubnets[i].Type, subnet.Type)
				}
			}
		})
	}
}

func TestMakeEC2Tags(t *testing.T) {
	tests := []struct {
		name     string
		tags     map[string]string
		expected []types.Tag
	}{
		{
			name:     "empty tags",
			tags:     map[string]string{},
			expected: []types.Tag{},
		},
		{
			name: "single tag",
			tags: map[string]string{
				"Name": "test-vpc",
			},
			expected: []types.Tag{
				{Key: aws.String("Name"), Value: aws.String("test-vpc")},
			},
		},
		{
			name: "multiple tags",
			tags: map[string]string{
				"Name":        "test-vpc",
				"Environment": "production",
				"Team":        "platform",
			},
			expected: []types.Tag{
				{Key: aws.String("Name"), Value: aws.String("test-vpc")},
				{Key: aws.String("Environment"), Value: aws.String("production")},
				{Key: aws.String("Team"), Value: aws.String("platform")},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ec2Tags := makeEC2Tags(tt.tags)

			if len(ec2Tags) != len(tt.expected) {
				t.Errorf("expected %d tags, got %d", tt.expected, len(ec2Tags))
			}

			// Verify tags are properly converted
			tagMap := make(map[string]string)
			for _, tag := range ec2Tags {
				tagMap[*tag.Key] = *tag.Value
			}

			for key, value := range tt.tags {
				if tagMap[key] != value {
					t.Errorf("expected tag %s=%s, got %s=%s", key, value, key, tagMap[key])
				}
			}
		})
	}
}

func TestAwsSubnetToCloudSubnet(t *testing.T) {
	tests := []struct {
		name       string
		awsSubnet  *types.Subnet
		subnetType v1.SubnetType
		vpc        *types.Vpc
	}{
		{
			name: "valid public subnet",
			awsSubnet: &types.Subnet{
				SubnetId:         aws.String("subnet-123"),
				CidrBlock:        aws.String("10.0.0.0/24"),
				AvailabilityZone: aws.String("us-east-1a"),
				Tags: []types.Tag{
					{Key: aws.String(tagBrevRefID), Value: aws.String("test-subnet")},
					{Key: aws.String(tagName), Value: aws.String("test-subnet-name")},
				},
			},
			subnetType: v1.SubnetTypePublic,
			vpc: &types.Vpc{
				VpcId: aws.String("vpc-123"),
			},
		},
		{
			name: "valid private subnet",
			awsSubnet: &types.Subnet{
				SubnetId:         aws.String("subnet-456"),
				CidrBlock:        aws.String("10.0.1.0/24"),
				AvailabilityZone: aws.String("us-east-1b"),
				Tags: []types.Tag{
					{Key: aws.String(tagBrevRefID), Value: aws.String("test-private-subnet")},
					{Key: aws.String(tagName), Value: aws.String("test-private-subnet-name")},
				},
			},
			subnetType: v1.SubnetTypePrivate,
			vpc: &types.Vpc{
				VpcId: aws.String("vpc-123"),
			},
		},
		{
			name: "subnet with minimal tags",
			awsSubnet: &types.Subnet{
				SubnetId:         aws.String("subnet-789"),
				CidrBlock:        aws.String("10.0.2.0/24"),
				AvailabilityZone: aws.String("us-east-1c"),
				Tags: []types.Tag{
					{Key: aws.String(tagBrevRefID), Value: aws.String("minimal-subnet")},
					{Key: aws.String(tagName), Value: aws.String("minimal-subnet-name")},
				},
			},
			subnetType: v1.SubnetTypePublic,
			vpc: &types.Vpc{
				VpcId: aws.String("vpc-123"),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := awsSubnetToCloudSubnet(tt.awsSubnet, tt.subnetType, tt.vpc)
			if err != nil {
				t.Fatalf("expected no error but got: %v", err)
			}
			if result != nil {
				if string(result.GetID()) != *tt.awsSubnet.SubnetId {
					t.Errorf("expected subnet ID %s, got %s", *tt.awsSubnet.SubnetId, result.GetID())
				}
				if result.GetCidrBlock() != *tt.awsSubnet.CidrBlock {
					t.Errorf("expected CIDR block %s, got %s", *tt.awsSubnet.CidrBlock, result.GetCidrBlock())
				}
				if result.GetSubnetType() != tt.subnetType {
					t.Errorf("expected subnet type %v, got %v", tt.subnetType, result.GetSubnetType())
				}
				if string(result.GetVPCID()) != *tt.vpc.VpcId {
					t.Errorf("expected VPC ID %s, got %s", *tt.vpc.VpcId, result.GetVPCID())
				}
			}
		})
	}
}
