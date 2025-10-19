package v1

import "context"

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
	Subnets []Subnet
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
)

type CloudMaintainVPC interface {
	CreateVPC(ctx context.Context, args CreateVPCArgs) (*VPC, error)
	GetVPC(ctx context.Context, args GetVPCArgs) (*VPC, error)
	DeleteVPC(ctx context.Context, args DeleteVPCArgs) error
}

type CreateVPCArgs struct {
	Name      string
	RefID     string
	Location  string
	CidrBlock string
	Subnets   []CreateSubnetArgs
}

type CreateSubnetArgs struct {
	CidrBlock string
	Type      SubnetType
}

type GetVPCArgs struct {
	ID       CloudProviderResourceID
	Location string
}

type DeleteVPCArgs struct {
	ID CloudProviderResourceID
}
