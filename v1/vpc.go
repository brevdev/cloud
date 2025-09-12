package v1

import (
	"context"
	"fmt"

	"github.com/brevdev/cloud/internal/errors"
)

// VPC represents the complete specification of a Brev VPC.
type VPC struct {
	// The name of the VPC, displayed on clients.
	name string

	// The unique ID used to associate with this VPC.
	refID string

	// The cloud provider that manages the VPC. Unless the provider is a broker to other clouds, this will be the same as
	// the Cloud field. For example, "aws".
	provider string

	// The cloud that hosts the VPC. For example, "aws".
	cloud string

	// The ID assigned by the cloud provider to the VPC.
	id CloudProviderResourceID

	// The location of the VPC. For example, "us-east-1".
	location string

	// The IPv4 network range for the VPC, in CIDR notation. For example, "10.0.0.0/16".
	cidrBlock string

	// The status of the VPC.
	status VPCStatus

	// The subnets associated with the VPC.
	subnets []*Subnet

	// The tags associated with the VPC.
	tags Tags
}

type VPCStatus string

const (
	VPCStatusAvailable VPCStatus = "available"
	VPCStatusPending   VPCStatus = "pending"
	VPCStatusDeleting  VPCStatus = "deleting"
	VPCStatusUnknown   VPCStatus = "unknown"
)

func (v *VPC) GetName() string {
	return v.name
}

func (v *VPC) GetRefID() string {
	return v.refID
}

func (v *VPC) GetProvider() string {
	return v.provider
}

func (v *VPC) GetCloud() string {
	return v.cloud
}

func (v *VPC) GetID() CloudProviderResourceID {
	return v.id
}

func (v *VPC) GetLocation() string {
	return v.location
}

func (v *VPC) GetCidrBlock() string {
	return v.cidrBlock
}

func (v *VPC) GetStatus() VPCStatus {
	return v.status
}

func (v *VPC) GetSubnets() []*Subnet {
	return v.subnets
}

func (v *VPC) GetTags() Tags {
	return v.tags
}

// VPCSettings represents the settings for a VPC. This is the input to the NewVPC function.
type VPCSettings struct {
	// The name of the VPC, displayed on clients.
	Name string

	// The unique ID used to associate with this VPC.
	RefID string

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

func (s *VPCSettings) setDefaults() {
}

func (s *VPCSettings) validate() error {
	var errs []error
	if s.RefID == "" {
		errs = append(errs, fmt.Errorf("refID is required"))
	}
	if s.Name == "" {
		errs = append(errs, fmt.Errorf("name is required"))
	}
	if s.Provider == "" {
		errs = append(errs, fmt.Errorf("provider is required"))
	}
	if s.ID == "" {
		errs = append(errs, fmt.Errorf("id is required"))
	}
	if s.Location == "" {
		errs = append(errs, fmt.Errorf("location is required"))
	}
	if s.CidrBlock == "" {
		errs = append(errs, fmt.Errorf("cidrBlock is required"))
	}
	if s.Status == "" {
		errs = append(errs, fmt.Errorf("status is required"))
	}
	return errors.WrapAndTrace(errors.Join(errs...))
}

// NewVPC creates a new VPC from the provided settings.
func NewVPC(settings VPCSettings) (*VPC, error) {
	settings.setDefaults()
	err := settings.validate()
	if err != nil {
		return nil, errors.WrapAndTrace(err)
	}
	return &VPC{
		name:      settings.Name,
		refID:     settings.RefID,
		provider:  settings.Provider,
		cloud:     settings.Cloud,
		id:        settings.ID,
		location:  settings.Location,
		cidrBlock: settings.CidrBlock,
		status:    settings.Status,
		subnets:   settings.Subnets,
		tags:      settings.Tags,
	}, nil
}

// Subnet represents the complete specification of a Brev subnet.
type Subnet struct {
	// The name of the subnet, displayed on clients.
	name string

	// The unique ID used to associate with this subnet.
	refID string

	// The ID of the VPC that the subnet is associated with.
	vPCID CloudProviderResourceID

	// The ID assigned by the cloud provider to the subnet.
	id CloudProviderResourceID

	// The location of the subnet. For example, "us-east-1".
	location string

	// The IPv4 network range for the subnet, in CIDR notation. For example, "10.0.0.0/24".
	cidrBlock string

	// The type of the subnet.
	subnetType SubnetType

	// The tags associated with the subnet.
	tags Tags
}

type SubnetType string

const (
	SubnetTypePublic  SubnetType = "public"
	SubnetTypePrivate SubnetType = "private"
)

func (s *Subnet) GetName() string {
	return s.name
}

func (s *Subnet) GetRefID() string {
	return s.refID
}

func (s *Subnet) GetVPCID() CloudProviderResourceID {
	return s.vPCID
}

func (s *Subnet) GetID() CloudProviderResourceID {
	return s.id
}

func (s *Subnet) GetLocation() string {
	return s.location
}

func (s *Subnet) GetCidrBlock() string {
	return s.cidrBlock
}

func (s *Subnet) GetSubnetType() SubnetType {
	return s.subnetType
}

func (s *Subnet) GetTags() Tags {
	return s.tags
}

// SubnetSettings represents the settings for a subnet. This is the input to the NewSubnet function.
type SubnetSettings struct {
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

func (s *SubnetSettings) setDefaults() {
}

func (s *SubnetSettings) validate() error {
	var errs []error
	if s.RefID == "" {
		errs = append(errs, fmt.Errorf("refID is required"))
	}
	if s.Name == "" {
		errs = append(errs, fmt.Errorf("name is required"))
	}
	if s.VPCID == "" {
		errs = append(errs, fmt.Errorf("vPCID is required"))
	}
	if s.ID == "" {
		errs = append(errs, fmt.Errorf("id is required"))
	}

	return errors.WrapAndTrace(errors.Join(errs...))
}

// NewSubnet creates a new Subnet from the provided settings.
func NewSubnet(settings SubnetSettings) (*Subnet, error) {
	settings.setDefaults()
	err := settings.validate()
	if err != nil {
		return nil, errors.WrapAndTrace(err)
	}
	return &Subnet{
		name:       settings.Name,
		refID:      settings.RefID,
		vPCID:      settings.VPCID,
		id:         settings.ID,
		location:   settings.Location,
		cidrBlock:  settings.CidrBlock,
		subnetType: settings.Type,
		tags:       settings.Tags,
	}, nil
}

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

	// The IPv4 network range for the VPC, in CIDR notation. For example, "10.0.0.0/16".
	CidrBlock string

	// The subnets to create in the VPC.
	Subnets []CreateSubnetArgs

	// The tags to associate with the VPC.
	Tags Tags
}

type CreateSubnetArgs struct {
	// The unique ID used to associate with this subnet.
	RefID string

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
}

type DeleteVPCArgs struct {
	// The ID of the VPC to delete.
	ID CloudProviderResourceID
}
