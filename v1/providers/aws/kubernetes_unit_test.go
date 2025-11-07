package v1

import (
	"errors"
	"math"
	"reflect"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	ekstypes "github.com/aws/aws-sdk-go-v2/service/eks/types"

	v1 "github.com/brevdev/cloud/v1"
)

func TestValidateCreateNodeGroupArgs(t *testing.T) { //nolint:funlen // test ok
	tests := []struct {
		name        string
		args        v1.CreateNodeGroupArgs
		expectError error
	}{
		{
			name: "valid args",
			args: v1.CreateNodeGroupArgs{
				Name:         "test-node-group",
				RefID:        "test-ref",
				MinNodeCount: 1,
				MaxNodeCount: 3,
				InstanceType: "t3.medium",
				DiskSizeGiB:  20,
				ClusterID:    "cluster-123",
			},
			expectError: nil,
		},
		{
			name: "min node count less than 1",
			args: v1.CreateNodeGroupArgs{
				Name:         "test-node-group",
				RefID:        "test-ref",
				MinNodeCount: 0,
				MaxNodeCount: 3,
				InstanceType: "t3.medium",
				DiskSizeGiB:  20,
			},
			expectError: errNodeGroupMinNodeCountMustBeGreaterThan0,
		},
		{
			name: "max node count less than 1",
			args: v1.CreateNodeGroupArgs{
				Name:         "test-node-group",
				RefID:        "test-ref",
				MinNodeCount: 1,
				MaxNodeCount: 0,
				InstanceType: "t3.medium",
				DiskSizeGiB:  20,
			},
			expectError: errNodeGroupMaxNodeCountMustBeGreaterThan0,
		},
		{
			name: "max node count less than min node count",
			args: v1.CreateNodeGroupArgs{
				Name:         "test-node-group",
				RefID:        "test-ref",
				MinNodeCount: 5,
				MaxNodeCount: 3,
				InstanceType: "t3.medium",
				DiskSizeGiB:  20,
			},
			expectError: errNodeGroupMaxNodeCountMustBeGreaterThanOrEqualToMinNodeCount,
		},
		{
			name: "missing instance type",
			args: v1.CreateNodeGroupArgs{
				Name:         "test-node-group",
				RefID:        "test-ref",
				MinNodeCount: 1,
				MaxNodeCount: 3,
				InstanceType: "",
				DiskSizeGiB:  20,
			},
			expectError: errNodeGroupInstanceTypeIsRequired,
		},
		{
			name: "disk size too small",
			args: v1.CreateNodeGroupArgs{
				Name:         "test-node-group",
				RefID:        "test-ref",
				MinNodeCount: 1,
				MaxNodeCount: 3,
				InstanceType: "t3.medium",
				DiskSizeGiB:  10,
			},
			expectError: errNodeGroupDiskSizeGiBMustBeGreaterThanOrEqualTo20,
		},
		{
			name: "disk size exceeds max int32",
			args: v1.CreateNodeGroupArgs{
				Name:         "test-node-group",
				RefID:        "test-ref",
				MinNodeCount: 1,
				MaxNodeCount: 3,
				InstanceType: "t3.medium",
				DiskSizeGiB:  math.MaxInt32 + 1,
			},
			expectError: errNodeGroupDiskSizeGiBMustBeLessThanOrEqualToMaxInt32,
		},
		{
			name: "max node count exceeds max int32",
			args: v1.CreateNodeGroupArgs{
				Name:         "test-node-group",
				RefID:        "test-ref",
				MinNodeCount: 1,
				MaxNodeCount: math.MaxInt32 + 1,
				InstanceType: "t3.medium",
				DiskSizeGiB:  20,
			},
			expectError: errNodeGroupMaxNodeCountMustBeLessThanOrEqualToMaxInt32,
		},
		{
			name: "min node count exceeds max int32",
			args: v1.CreateNodeGroupArgs{
				Name:         "test-node-group",
				RefID:        "test-ref",
				MinNodeCount: math.MaxInt32 + 1,
				MaxNodeCount: math.MaxInt32 + 2,
				InstanceType: "t3.medium",
				DiskSizeGiB:  20,
			},
			expectError: errNodeGroupMinNodeCountMustBeLessThanOrEqualToMaxInt32,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateCreateNodeGroupArgs(tt.args)
			if err != nil && tt.expectError == nil {
				t.Fatalf("expected no error but got: %v", err)
			}
			if err != nil && tt.expectError != nil {
				if !errors.Is(err, tt.expectError) {
					t.Errorf("expected error %v, got %v", tt.expectError, err)
				}
			}
		})
	}
}

func TestValidateModifyNodeGroupArgs(t *testing.T) {
	tests := []struct {
		name        string
		args        v1.ModifyNodeGroupArgs
		expectError error
	}{
		{
			name: "valid args",
			args: v1.ModifyNodeGroupArgs{
				ID:           "node-group-123",
				ClusterID:    "cluster-123",
				MinNodeCount: 1,
				MaxNodeCount: 3,
			},
			expectError: nil,
		},
		{
			name: "min node count less than 1",
			args: v1.ModifyNodeGroupArgs{
				ID:           "node-group-123",
				ClusterID:    "cluster-123",
				MinNodeCount: 0,
				MaxNodeCount: 3,
			},
			expectError: errNodeGroupMinNodeCountMustBeGreaterThan0,
		},
		{
			name: "max node count less than 1",
			args: v1.ModifyNodeGroupArgs{
				ID:           "node-group-123",
				ClusterID:    "cluster-123",
				MinNodeCount: 1,
				MaxNodeCount: 0,
			},
			expectError: errNodeGroupMaxNodeCountMustBeGreaterThan0,
		},
		{
			name: "max node count less than min node count",
			args: v1.ModifyNodeGroupArgs{
				ID:           "node-group-123",
				ClusterID:    "cluster-123",
				MinNodeCount: 5,
				MaxNodeCount: 3,
			},
			expectError: errNodeGroupMaxNodeCountMustBeGreaterThanOrEqualToMinNodeCount,
		},
		{
			name: "min node count exceeds max int32",
			args: v1.ModifyNodeGroupArgs{
				ID:           "node-group-123",
				ClusterID:    "cluster-123",
				MinNodeCount: math.MaxInt32 + 1,
				MaxNodeCount: math.MaxInt32 + 2,
			},
			expectError: errNodeGroupMinNodeCountMustBeLessThanOrEqualToMaxInt32,
		},
		{
			name: "max node count exceeds max int32",
			args: v1.ModifyNodeGroupArgs{
				ID:           "node-group-123",
				ClusterID:    "cluster-123",
				MinNodeCount: 1,
				MaxNodeCount: math.MaxInt32 + 1,
			},
			expectError: errNodeGroupMaxNodeCountMustBeLessThanOrEqualToMaxInt32,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateModifyNodeGroupArgs(tt.args)
			if err != nil && tt.expectError == nil {
				t.Errorf("expected error but got nil")
			}
			if err != nil && tt.expectError != nil {
				if !errors.Is(err, tt.expectError) {
					t.Errorf("expected error %v, got %v", tt.expectError, err)
				}
			}
		})
	}
}

func TestValidatePutUserArgs(t *testing.T) {
	tests := []struct {
		name        string
		args        v1.SetClusterUserArgs
		expectError error
	}{
		{
			name: "valid args",
			args: v1.SetClusterUserArgs{
				Username:     "test-user",
				Role:         "cluster-admin",
				ClusterID:    "cluster-123",
				RSAPEMBase64: "base64encodedkey",
			},
			expectError: nil,
		},
		{
			name: "missing username",
			args: v1.SetClusterUserArgs{
				Username:     "",
				Role:         "cluster-admin",
				ClusterID:    "cluster-123",
				RSAPEMBase64: "base64encodedkey",
			},
			expectError: errUsernameIsRequired,
		},
		{
			name: "missing role",
			args: v1.SetClusterUserArgs{
				Username:     "test-user",
				Role:         "",
				ClusterID:    "cluster-123",
				RSAPEMBase64: "base64encodedkey",
			},
			expectError: errRoleIsRequired,
		},
		{
			name: "missing cluster ID",
			args: v1.SetClusterUserArgs{
				Username:     "test-user",
				Role:         "cluster-admin",
				ClusterID:    "",
				RSAPEMBase64: "base64encodedkey",
			},
			expectError: errClusterIDIsRequired,
		},
		{
			name: "missing RSA PEM base64",
			args: v1.SetClusterUserArgs{
				Username:     "test-user",
				Role:         "cluster-admin",
				ClusterID:    "cluster-123",
				RSAPEMBase64: "",
			},
			expectError: errRSAPEMBase64IsRequired,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validatePutUserArgs(tt.args)
			if err != nil && tt.expectError == nil {
				t.Errorf("expected error but got nil")
			}
			if err != nil && tt.expectError != nil {
				if !errors.Is(err, tt.expectError) {
					t.Errorf("expected error %v, got %v", tt.expectError, err)
				}
			}
		})
	}
}

func TestMakeIAMTags(t *testing.T) {
	tests := []struct {
		name     string
		tags     map[string]string
		expected int
	}{
		{
			name:     "empty tags",
			tags:     map[string]string{},
			expected: 0,
		},
		{
			name: "single tag",
			tags: map[string]string{
				"Name": "test-cluster",
			},
			expected: 1,
		},
		{
			name: "multiple tags",
			tags: map[string]string{
				"Name":        "test-cluster",
				"Environment": "production",
				"Team":        "platform",
			},
			expected: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			iamTags := makeIAMTags(tt.tags)

			if len(iamTags) != tt.expected {
				t.Errorf("expected %d tags, got %d", tt.expected, len(iamTags))
			}

			// Verify tags are properly converted
			tagMap := make(map[string]string)
			for _, tag := range iamTags {
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

func TestGetNodeGroupIAMRolePath(t *testing.T) {
	tests := []struct {
		name            string
		clusterRefID    string
		nodeGroupRefID  string
		expectedContain string
	}{
		{
			name:            "basic path",
			clusterRefID:    "my-cluster",
			nodeGroupRefID:  "my-nodegroup",
			expectedContain: "/brevcloudsdk/eks/clusters/my-cluster/nodegroups/my-nodegroup/",
		},
		{
			name:            "with special characters",
			clusterRefID:    "my-cluster-123",
			nodeGroupRefID:  "my-nodegroup-456",
			expectedContain: "/brevcloudsdk/eks/clusters/my-cluster-123/nodegroups/my-nodegroup-456/",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			path := getNodeGroupIAMRolePath(tt.clusterRefID, tt.nodeGroupRefID)

			if path != tt.expectedContain {
				t.Errorf("expected path to be %s, got %s", tt.expectedContain, path)
			}
		})
	}
}

func TestParseEKSClusterStatus(t *testing.T) {
	tests := []struct {
		name           string
		status         ekstypes.ClusterStatus
		expectedStatus v1.ClusterStatus
	}{
		{
			name:           "creating",
			status:         ekstypes.ClusterStatusCreating,
			expectedStatus: v1.ClusterStatusPending,
		},
		{
			name:           "active",
			status:         ekstypes.ClusterStatusActive,
			expectedStatus: v1.ClusterStatusAvailable,
		},
		{
			name:           "deleting",
			status:         ekstypes.ClusterStatusDeleting,
			expectedStatus: v1.ClusterStatusDeleting,
		},
		{
			name:           "failed",
			status:         ekstypes.ClusterStatusFailed,
			expectedStatus: v1.ClusterStatusFailed,
		},
		{
			name:           "updating",
			status:         ekstypes.ClusterStatusUpdating,
			expectedStatus: v1.ClusterStatusPending,
		},
		{
			name:           "pending",
			status:         ekstypes.ClusterStatusPending,
			expectedStatus: v1.ClusterStatusPending,
		},
		{
			name:           "unknown",
			status:         ekstypes.ClusterStatus("foobar"),
			expectedStatus: v1.ClusterStatusUnknown,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := parseEKSClusterStatus(tt.status)
			if result != tt.expectedStatus {
				t.Errorf("expected status %v, got %v", tt.expectedStatus, result)
			}
		})
	}
}

func TestParseEKSNodeGroupStatus(t *testing.T) {
	tests := []struct {
		name           string
		status         ekstypes.NodegroupStatus
		expectedStatus v1.NodeGroupStatus
	}{
		{
			name:           "creating",
			status:         ekstypes.NodegroupStatusCreating,
			expectedStatus: v1.NodeGroupStatusPending,
		},
		{
			name:           "active",
			status:         ekstypes.NodegroupStatusActive,
			expectedStatus: v1.NodeGroupStatusAvailable,
		},
		{
			name:           "deleting",
			status:         ekstypes.NodegroupStatusDeleting,
			expectedStatus: v1.NodeGroupStatusDeleting,
		},
		{
			name:           "create failed",
			status:         ekstypes.NodegroupStatusCreateFailed,
			expectedStatus: v1.NodeGroupStatusFailed,
		},
		{
			name:           "delete failed",
			status:         ekstypes.NodegroupStatusDeleteFailed,
			expectedStatus: v1.NodeGroupStatusFailed,
		},
		{
			name:           "unknown",
			status:         ekstypes.NodegroupStatus("foobar"),
			expectedStatus: v1.NodeGroupStatusUnknown,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := parseEKSNodeGroupStatus(tt.status)
			if result != tt.expectedStatus {
				t.Errorf("expected status %v, got %v", tt.expectedStatus, result)
			}
		})
	}
}

func TestGetClusterCACertificateBase64(t *testing.T) {
	tests := []struct {
		name     string
		cluster  *ekstypes.Cluster
		expected string
	}{
		{
			name:     "nil cluster",
			cluster:  nil,
			expected: "",
		},
		{
			name: "nil certificate authority",
			cluster: &ekstypes.Cluster{
				CertificateAuthority: nil,
			},
			expected: "",
		},
		{
			name: "nil certificate data",
			cluster: &ekstypes.Cluster{
				CertificateAuthority: &ekstypes.Certificate{
					Data: nil,
				},
			},
			expected: "",
		},
		{
			name: "valid certificate",
			cluster: &ekstypes.Cluster{
				CertificateAuthority: &ekstypes.Certificate{
					Data: aws.String("base64encodedcert"),
				},
			},
			expected: "base64encodedcert",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := getClusterCACertificateBase64(tt.cluster)
			if result != tt.expected {
				t.Errorf("expected %s, got %s", tt.expected, result)
			}
		})
	}
}

func TestGetClusterAPIEndpoint(t *testing.T) {
	tests := []struct {
		name     string
		cluster  *ekstypes.Cluster
		expected string
	}{
		{
			name:     "nil cluster",
			cluster:  nil,
			expected: "",
		},
		{
			name: "nil endpoint",
			cluster: &ekstypes.Cluster{
				Endpoint: nil,
			},
			expected: "",
		},
		{
			name: "valid endpoint",
			cluster: &ekstypes.Cluster{
				Endpoint: aws.String("https://example.com"),
			},
			expected: "https://example.com",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := getClusterAPIEndpoint(tt.cluster)
			if result != tt.expected {
				t.Errorf("expected %s, got %s", tt.expected, result)
			}
		})
	}
}

func TestParseEKSNodeGroup(t *testing.T) { //nolint:gocognit // test ok
	tests := []struct {
		name      string
		nodeGroup *ekstypes.Nodegroup
	}{
		{
			name: "valid node group",
			nodeGroup: &ekstypes.Nodegroup{
				NodegroupName: aws.String("test-nodegroup"),
				ScalingConfig: &ekstypes.NodegroupScalingConfig{
					MinSize: aws.Int32(1),
					MaxSize: aws.Int32(3),
				},
				InstanceTypes: []string{"t3.medium"},
				DiskSize:      aws.Int32(20),
				Status:        ekstypes.NodegroupStatusActive,
				Tags: map[string]string{
					tagBrevRefID: "test-ref",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := parseEKSNodeGroup(tt.nodeGroup)
			if err != nil {
				t.Fatalf("expected no error but got: %v", err)
			}
			if result == nil {
				t.Fatalf("expected valid node group but got nil")
			}
			if result.GetName() != *tt.nodeGroup.NodegroupName {
				t.Errorf("expected name %s, got %s", *tt.nodeGroup.NodegroupName, result.GetName())
			}
			if result.GetMinNodeCount() != int(*tt.nodeGroup.ScalingConfig.MinSize) {
				t.Errorf("expected min node count %d, got %d", *tt.nodeGroup.ScalingConfig.MinSize, result.GetMinNodeCount())
			}
			if result.GetMaxNodeCount() != int(*tt.nodeGroup.ScalingConfig.MaxSize) {
				t.Errorf("expected max node count %d, got %d", *tt.nodeGroup.ScalingConfig.MaxSize, result.GetMaxNodeCount())
			}
			if result.GetInstanceType() != tt.nodeGroup.InstanceTypes[0] {
				t.Errorf("expected instance type %s, got %s", tt.nodeGroup.InstanceTypes[0], result.GetInstanceType())
			}
			if result.GetDiskSizeGiB() != int(*tt.nodeGroup.DiskSize) {
				t.Errorf("expected disk size %d, got %d", *tt.nodeGroup.DiskSize, result.GetDiskSizeGiB())
			}
			if result.GetStatus() != parseEKSNodeGroupStatus(tt.nodeGroup.Status) {
				t.Errorf("expected status %v, got %v", parseEKSNodeGroupStatus(tt.nodeGroup.Status), result.GetStatus())
			}
			if result.GetRefID() != tt.nodeGroup.Tags[tagBrevRefID] {
				t.Errorf("expected ref ID %s, got %s", tt.nodeGroup.Tags[tagBrevRefID], result.GetRefID())
			}
			if !reflect.DeepEqual(result.GetTags(), v1.Tags(tt.nodeGroup.Tags)) {
				t.Errorf("expected tags %v, got %v", v1.Tags(tt.nodeGroup.Tags), result.GetTags())
			}
		})
	}
}
