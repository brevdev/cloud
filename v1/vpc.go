package v1

import (
	"context"
	"fmt"
	"time"
)

type VPC struct {
	// The name of the VPC, displayed on clients.
	Name string

	// The unique ID used to associate with this VPC.
	RefID string

	// The ID of the cloud credential used to create the VPC.
	CloudCredRefID string

	// The cloud provider that manages the VPC. Unless the provider is a broker to other clouds, this will be the same as
	// the Cloud field. For example, "aws".
	Provider string

	// The cloud that hosts the VPC. For example, "aws".
	Cloud string

	// The ID assigned by the cloud provider to the VPC.
	ID CloudProviderResourceID

	// The location of the VPC. For example, "us-east-1".
	Location string

	// The IPv4 network range for the VPC, in CIDR notation. For example, "10.0.0.0/16".
	CidrBlock string

	// The status of the VPC.
	Status VPCStatus

	// The subnets associated with the VPC.
	Subnets []*Subnet

	// The tags associated with the VPC.
	Tags Tags
}

type Subnet struct {
	// The name of the subnet, displayed on clients.
	Name string

	// The unique ID used to associate with this subnet.
	RefID string

	// The ID of the VPC that the subnet is associated with.
	VPCID CloudProviderResourceID

	// The ID assigned by the cloud provider to the subnet.
	ID CloudProviderResourceID

	// The location of the subnet. For example, "us-east-1".
	Location string

	// The IPv4 network range for the subnet, in CIDR notation. For example, "10.0.0.0/24".
	CidrBlock string

	// The type of the subnet.
	Type SubnetType

	// The tags associated with the subnet.
	Tags Tags
}

type SubnetType string

const (
	SubnetTypePublic  SubnetType = "public"
	SubnetTypePrivate SubnetType = "private"
)

type VPCStatus string

const (
	VPCStatusAvailable VPCStatus = "available"
	VPCStatusPending   VPCStatus = "pending"
	VPCStatusDeleting  VPCStatus = "deleting"
	VPCStatusUnknown   VPCStatus = "unknown"
)

type CloudMaintainVPC interface {
	// Create a new VPC.
	CreateVPC(ctx context.Context, args CreateVPCArgs) (*VPC, error)

	// Get a VPC identified by the provided args.
	GetVPC(ctx context.Context, args GetVPCArgs) (*VPC, error)

	// Delete a VPC identified by the provided args.
	DeleteVPC(ctx context.Context, args DeleteVPCArgs) error
}

type CreateVPCArgs struct {
	// The name of the VPC, displayed on clients.
	Name string

	// The unique ID used to associate with this VPC.
	RefID string

	// The location of the VPC.
	Location string

	// The IPv4 network range for the VPC, in CIDR notation. For example, "10.0.0.0/16".
	CidrBlock string

	// The subnets to create in the VPC.
	Subnets []CreateSubnetArgs

	// The tags to associate with the VPC.
	Tags Tags
}

type CreateSubnetArgs struct {
	// The IPv4 network range for the subnet, in CIDR notation. For example, "10.0.0.0/24".
	CidrBlock string

	// The type of the subnet.
	Type SubnetType

	// The tags to associate with the subnet.
	Tags Tags
}

type GetVPCArgs struct {
	// The ID of the VPC to get.
	ID CloudProviderResourceID

	// The location of the VPC.
	Location string
}

type DeleteVPCArgs struct {
	// The ID of the VPC to delete.
	ID CloudProviderResourceID
}

func ValidateCreateVPC(ctx context.Context, client CloudMaintainVPC, attrs CreateVPCArgs) (*VPC, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Minute)
	defer cancel()

	vpc, err := client.CreateVPC(ctx, attrs)
	if err != nil {
		return nil, err
	}

	if vpc.Name != attrs.Name {
		return nil, fmt.Errorf("VPC name does not match create args: '%s' != '%s'", vpc.Name, attrs.Name)
	}
	if vpc.RefID != attrs.RefID {
		return nil, fmt.Errorf("VPC refID does not match create args: '%s' != '%s'", vpc.RefID, attrs.RefID)
	}
	if vpc.Location != attrs.Location {
		return nil, fmt.Errorf("VPC location does not match create args: '%s' != '%s'", vpc.Location, attrs.Location)
	}
	if vpc.CidrBlock != attrs.CidrBlock {
		return nil, fmt.Errorf("VPC cidr block does not match create args: '%s' != '%s'", vpc.CidrBlock, attrs.CidrBlock)
	}
	if len(vpc.Subnets) != len(attrs.Subnets) {
		return nil, fmt.Errorf("VPC subnets does not match create args: '%d' != '%d'", len(vpc.Subnets), len(attrs.Subnets))
	}
	for key, value := range vpc.Tags {
		tagValue, ok := attrs.Tags[key]
		if !ok {
			return nil, fmt.Errorf("VPC tag does not match create args: '%s' not found", key)
		}
		if tagValue != value {
			return nil, fmt.Errorf("VPC tag does not match create args: '%s' != '%s'", key, value)
		}
	}

	cidrToSubnetMap := make(map[string]*Subnet)
	for _, subnet := range vpc.Subnets {
		cidrToSubnetMap[subnet.CidrBlock] = subnet
	}
	for _, subnet := range attrs.Subnets {
		subnetFromMap, ok := cidrToSubnetMap[subnet.CidrBlock]
		if !ok {
			return nil, fmt.Errorf("VPC subnet cidr block does not match create args: '%s' not found", subnet.CidrBlock)
		}
		if subnetFromMap.Type != subnet.Type {
			return nil, fmt.Errorf("VPC subnet type does not match create args: '%s' != '%s'", subnetFromMap.Type, subnet.Type)
		}
	}

	return vpc, nil
}

func ValidateGetVPC(ctx context.Context, client CloudMaintainVPC, attrs GetVPCArgs) (*VPC, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Minute)
	defer cancel()

	vpc, err := client.GetVPC(ctx, attrs)
	if err != nil {
		return nil, err
	}

	if vpc.ID != attrs.ID {
		return nil, fmt.Errorf("VPC ID does not match get args: '%s' != '%s'", vpc.ID, attrs.ID)
	}
	if attrs.Location != "" && vpc.Location != attrs.Location {
		return nil, fmt.Errorf("VPC location does not match get args: '%s' != '%s'", vpc.Location, attrs.Location)
	}

	return vpc, nil
}

func ValidateDeleteVPC(ctx context.Context, client CloudMaintainVPC, attrs DeleteVPCArgs) error {
	err := client.DeleteVPC(ctx, attrs)
	if err != nil {
		return err
	}
	return nil
}
