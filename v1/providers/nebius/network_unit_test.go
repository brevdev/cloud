package v1

import (
	"errors"
	"testing"

	nebiusvpc "github.com/nebius/gosdk/proto/nebius/vpc/v1"

	v1 "github.com/brevdev/cloud/v1"
)

func TestValidateCreateVPCArgs(t *testing.T) {
	tests := []struct {
		name        string
		args        v1.CreateVPCArgs
		expectError error
	}{
		{
			name: "valid args with large enough subnets",
			args: v1.CreateVPCArgs{
				RefID:     "test-vpc",
				Name:      "test-vpc",
				CidrBlock: "172.16.0.0/16",
				Subnets: []v1.CreateSubnetArgs{
					{CidrBlock: "172.16.0.0/19", Type: v1.SubnetTypePublic},
					{CidrBlock: "172.16.32.0/19", Type: v1.SubnetTypePrivate},
				},
			},
			expectError: nil,
		},
		{
			name: "invalid - subnet CIDR /24 (not larger than /24)",
			args: v1.CreateVPCArgs{
				RefID:     "test-vpc",
				Name:      "test-vpc",
				CidrBlock: "172.16.0.0/16",
				Subnets: []v1.CreateSubnetArgs{
					{CidrBlock: "172.16.0.0/24", Type: v1.SubnetTypePublic},
				},
			},
			expectError: errVPCSubnetCIDRBlockMustBeGreaterThan24,
		},
		{
			name: "invalid - subnet CIDR /28 (smaller than /24)",
			args: v1.CreateVPCArgs{
				RefID:     "test-vpc",
				Name:      "test-vpc",
				CidrBlock: "172.16.0.0/16",
				Subnets: []v1.CreateSubnetArgs{
					{CidrBlock: "172.16.0.0/28", Type: v1.SubnetTypePrivate},
				},
			},
			expectError: errVPCSubnetCIDRBlockMustBeGreaterThan24,
		},
		{
			name: "valid - subnet CIDR /23 (larger than /24)",
			args: v1.CreateVPCArgs{
				RefID:     "test-vpc",
				Name:      "test-vpc",
				CidrBlock: "172.16.0.0/16",
				Subnets: []v1.CreateSubnetArgs{
					{CidrBlock: "172.16.0.0/23", Type: v1.SubnetTypePublic},
				},
			},
			expectError: nil,
		},
		{
			name: "invalid - one subnet too small",
			args: v1.CreateVPCArgs{
				RefID:     "test-vpc",
				Name:      "test-vpc",
				CidrBlock: "172.16.0.0/16",
				Subnets: []v1.CreateSubnetArgs{
					{CidrBlock: "172.16.0.0/19", Type: v1.SubnetTypePublic},
					{CidrBlock: "172.16.32.0/28", Type: v1.SubnetTypePrivate},
				},
			},
			expectError: errVPCSubnetCIDRBlockMustBeGreaterThan24,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateCreateVPCArgs(tt.args)
			if err != nil && tt.expectError != nil {
				if !errors.Is(err, tt.expectError) {
					t.Fatalf("expected error %v, got %v", tt.expectError, err)
				}
			}
			if err == nil && tt.expectError != nil {
				t.Fatalf("expected error but got nil")
			}
		})
	}
}

func TestCidrBlockLargerThanMask(t *testing.T) { //nolint:funlen // test ok
	tests := []struct {
		name        string
		cidrBlock   string
		mask        int
		expected    bool
		expectError bool
	}{
		{
			name:        "/16 is larger than /24",
			cidrBlock:   "10.0.0.0/16",
			mask:        24,
			expected:    true,
			expectError: false,
		},
		{
			name:        "/19 is larger than /24",
			cidrBlock:   "10.0.0.0/19",
			mask:        24,
			expected:    true,
			expectError: false,
		},
		{
			name:        "/23 is larger than /24",
			cidrBlock:   "10.0.0.0/23",
			mask:        24,
			expected:    true,
			expectError: false,
		},
		{
			name:        "/24 is not larger than /24",
			cidrBlock:   "10.0.0.0/24",
			mask:        24,
			expected:    false,
			expectError: false,
		},
		{
			name:        "/28 is not larger than /24",
			cidrBlock:   "10.0.0.0/28",
			mask:        24,
			expected:    false,
			expectError: false,
		},
		{
			name:        "/32 is not larger than /24",
			cidrBlock:   "10.0.0.0/32",
			mask:        24,
			expected:    false,
			expectError: false,
		},
		{
			name:        "/8 is larger than /16",
			cidrBlock:   "10.0.0.0/8",
			mask:        16,
			expected:    true,
			expectError: false,
		},
		{
			name:        "/20 is not larger than /16",
			cidrBlock:   "10.0.0.0/20",
			mask:        16,
			expected:    false,
			expectError: false,
		},
		{
			name:        "invalid CIDR block",
			cidrBlock:   "invalid",
			mask:        24,
			expected:    false,
			expectError: true,
		},
		{
			name:        "invalid CIDR block - no mask",
			cidrBlock:   "10.0.0.0",
			mask:        24,
			expected:    false,
			expectError: true,
		},
		{
			name:        "IPv6 /48 is larger than /64",
			cidrBlock:   "2001:db8::/48",
			mask:        64,
			expected:    true,
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := cidrBlockLargerThanMask(tt.cidrBlock, tt.mask)

			if tt.expectError && err == nil {
				t.Errorf("expected error but got nil")
			}
			if !tt.expectError && err != nil {
				t.Errorf("expected no error but got: %v", err)
			}
			if !tt.expectError && result != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestParseNebiusNetworkStatus(t *testing.T) {
	tests := []struct {
		name           string
		status         *nebiusvpc.NetworkStatus
		expectedStatus v1.VPCStatus
	}{
		{
			name: "creating",
			status: &nebiusvpc.NetworkStatus{
				State: nebiusvpc.NetworkStatus_CREATING,
			},
			expectedStatus: v1.VPCStatusPending,
		},
		{
			name: "ready",
			status: &nebiusvpc.NetworkStatus{
				State: nebiusvpc.NetworkStatus_READY,
			},
			expectedStatus: v1.VPCStatusAvailable,
		},
		{
			name: "deleting",
			status: &nebiusvpc.NetworkStatus{
				State: nebiusvpc.NetworkStatus_DELETING,
			},
			expectedStatus: v1.VPCStatusDeleting,
		},
		{
			name: "unknown state",
			status: &nebiusvpc.NetworkStatus{
				State: nebiusvpc.NetworkStatus_STATE_UNSPECIFIED,
			},
			expectedStatus: v1.VPCStatusUnknown,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := parseNebiusNetworkStatus(tt.status)
			if result != tt.expectedStatus {
				t.Errorf("expected status %v, got %v", tt.expectedStatus, result)
			}
		})
	}
}
