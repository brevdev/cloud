package v1

import (
	"errors"
	"testing"

	nebiusmk8s "github.com/nebius/gosdk/proto/nebius/mk8s/v1"

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
				InstanceType: "cpu-d3.4vcpu-16gb",
				DiskSizeGiB:  64,
				ClusterID:    "cluster-123",
			},
			expectError: nil,
		},
		{
			name: "missing name",
			args: v1.CreateNodeGroupArgs{
				Name:         "",
				RefID:        "test-ref",
				MinNodeCount: 1,
				MaxNodeCount: 3,
				InstanceType: "cpu-d3.4vcpu-16gb",
				DiskSizeGiB:  64,
			},
			expectError: errNodeGroupNameIsRequired,
		},
		{
			name: "missing refID",
			args: v1.CreateNodeGroupArgs{
				Name:         "test-node-group",
				RefID:        "",
				MinNodeCount: 1,
				MaxNodeCount: 3,
				InstanceType: "cpu-d3.4vcpu-16gb",
				DiskSizeGiB:  64,
			},
			expectError: errNodeGroupRefIDIsRequired,
		},
		{
			name: "min node count less than 1",
			args: v1.CreateNodeGroupArgs{
				Name:         "test-node-group",
				RefID:        "test-ref",
				MinNodeCount: 0,
				MaxNodeCount: 3,
				InstanceType: "cpu-d3.4vcpu-16gb",
				DiskSizeGiB:  64,
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
				InstanceType: "cpu-d3.4vcpu-16gb",
				DiskSizeGiB:  64,
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
				InstanceType: "cpu-d3.4vcpu-16gb",
				DiskSizeGiB:  64,
			},
			expectError: errNodeGroupMaxNodeCountMustBeGreaterThanOrEqualToMinNodeCount,
		},
		{
			name: "disk size less than 64",
			args: v1.CreateNodeGroupArgs{
				Name:         "test-node-group",
				RefID:        "test-ref",
				MinNodeCount: 1,
				MaxNodeCount: 3,
				InstanceType: "cpu-d3.4vcpu-16gb",
				DiskSizeGiB:  32,
			},
			expectError: errNodeGroupDiskSizeGiBMustBeGreaterThanOrEqualTo64,
		},
		{
			name: "missing instance type",
			args: v1.CreateNodeGroupArgs{
				Name:         "test-node-group",
				RefID:        "test-ref",
				MinNodeCount: 1,
				MaxNodeCount: 3,
				InstanceType: "",
				DiskSizeGiB:  64,
			},
			expectError: errNodeGroupInstanceTypeIsRequired,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateCreateNodeGroupArgs(tt.args)
			if err != nil && tt.expectError != nil {
				if !errors.Is(err, tt.expectError) {
					t.Errorf("expected error %v, got %v", tt.expectError, err)
				}
			}
			if err == nil && tt.expectError != nil {
				t.Errorf("expected error but got nil")
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
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateModifyNodeGroupArgs(tt.args)
			if err != nil && tt.expectError != nil {
				if !errors.Is(err, tt.expectError) {
					t.Errorf("expected error %v, got %v", tt.expectError, err)
				}
			}
			if err == nil && tt.expectError != nil {
				t.Errorf("expected error but got nil")
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
			if err != nil && tt.expectError != nil {
				if !errors.Is(err, tt.expectError) {
					t.Errorf("expected error %v, got %v", tt.expectError, err)
				}
			}
			if err == nil && tt.expectError != nil {
				t.Errorf("expected error but got nil")
			}
		})
	}
}

func TestParseNebiusClusterStatus(t *testing.T) { //nolint:dupl // false positive/
	tests := []struct {
		name           string
		status         *nebiusmk8s.ClusterStatus
		expectedStatus v1.ClusterStatus
	}{
		{
			name:           "nil status",
			status:         nil,
			expectedStatus: v1.ClusterStatusUnknown,
		},
		{
			name: "provisioning",
			status: &nebiusmk8s.ClusterStatus{
				State: nebiusmk8s.ClusterStatus_PROVISIONING,
			},
			expectedStatus: v1.ClusterStatusPending,
		},
		{
			name: "running",
			status: &nebiusmk8s.ClusterStatus{
				State: nebiusmk8s.ClusterStatus_RUNNING,
			},
			expectedStatus: v1.ClusterStatusAvailable,
		},
		{
			name: "deleting",
			status: &nebiusmk8s.ClusterStatus{
				State: nebiusmk8s.ClusterStatus_DELETING,
			},
			expectedStatus: v1.ClusterStatusDeleting,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := parseNebiusClusterStatus(tt.status)
			if result != tt.expectedStatus {
				t.Errorf("expected status %v, got %v", tt.expectedStatus, result)
			}
		})
	}
}

func TestParseNebiusNodeGroupStatus(t *testing.T) { //nolint:dupl // false positive
	tests := []struct {
		name           string
		status         *nebiusmk8s.NodeGroupStatus
		expectedStatus v1.NodeGroupStatus
	}{
		{
			name:           "nil status",
			status:         nil,
			expectedStatus: v1.NodeGroupStatusUnknown,
		},
		{
			name: "provisioning",
			status: &nebiusmk8s.NodeGroupStatus{
				State: nebiusmk8s.NodeGroupStatus_PROVISIONING,
			},
			expectedStatus: v1.NodeGroupStatusPending,
		},
		{
			name: "running",
			status: &nebiusmk8s.NodeGroupStatus{
				State: nebiusmk8s.NodeGroupStatus_RUNNING,
			},
			expectedStatus: v1.NodeGroupStatusAvailable,
		},
		{
			name: "deleting",
			status: &nebiusmk8s.NodeGroupStatus{
				State: nebiusmk8s.NodeGroupStatus_DELETING,
			},
			expectedStatus: v1.NodeGroupStatusDeleting,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := parseNebiusNodeGroupStatus(tt.status)
			if result != tt.expectedStatus {
				t.Errorf("expected status %v, got %v", tt.expectedStatus, result)
			}
		})
	}
}
