package v1

import (
	"context"
	"fmt"
	"time"
)

// ValidateCreateVPC validates that the CreateVPC functionality works correctly.
func ValidateCreateVPC(ctx context.Context, client CloudMaintainVPC, attrs CreateVPCArgs) (*VPC, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Minute)
	defer cancel()

	vpc, err := client.CreateVPC(ctx, attrs)
	if err != nil {
		return nil, err
	}

	if vpc.GetName() != attrs.Name {
		return nil, fmt.Errorf("VPC name does not match create args: '%s' != '%s'", vpc.GetName(), attrs.Name)
	}
	if vpc.GetRefID() != attrs.RefID {
		return nil, fmt.Errorf("VPC refID does not match create args: '%s' != '%s'", vpc.GetRefID(), attrs.RefID)
	}
	if vpc.GetCidrBlock() != attrs.CidrBlock {
		return nil, fmt.Errorf("VPC cidr block does not match create args: '%s' != '%s'", vpc.GetCidrBlock(), attrs.CidrBlock)
	}
	if len(vpc.GetSubnets()) != len(attrs.Subnets) {
		return nil, fmt.Errorf("VPC subnets does not match create args: '%d' != '%d'", len(vpc.GetSubnets()), len(attrs.Subnets))
	}
	for key, value := range attrs.Tags {
		tagValue, ok := vpc.GetTags()[key]
		if !ok {
			return nil, fmt.Errorf("VPC tag does not match create args: '%s' not found", key)
		}
		if tagValue != value {
			return nil, fmt.Errorf("VPC tag does not match create args: '%s' != '%s'", key, value)
		}
	}

	cidrToSubnetMap := make(map[string]*Subnet)
	for _, subnet := range vpc.GetSubnets() {
		cidrToSubnetMap[subnet.GetCidrBlock()] = subnet
	}
	for _, subnetAttrs := range attrs.Subnets {
		subnetFromMap, ok := cidrToSubnetMap[subnetAttrs.CidrBlock]
		if !ok {
			return nil, fmt.Errorf("VPC subnet cidr block does not match create args: '%s' not found", subnetAttrs.CidrBlock)
		}
		if subnetFromMap.GetSubnetType() != subnetAttrs.Type {
			return nil, fmt.Errorf("VPC subnet type does not match create args: '%s' != '%s'", subnetFromMap.GetSubnetType(), subnetAttrs.Type)
		}
	}

	return vpc, nil
}

// ValidateGetVPC validates that the GetVPC functionality works correctly.
func ValidateGetVPC(ctx context.Context, client CloudMaintainVPC, attrs GetVPCArgs) (*VPC, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Minute)
	defer cancel()

	vpc, err := client.GetVPC(ctx, attrs)
	if err != nil {
		return nil, err
	}

	if vpc.GetID() != attrs.ID {
		return nil, fmt.Errorf("VPC ID does not match get args: '%s' != '%s'", vpc.GetID(), attrs.ID)
	}

	return vpc, nil
}

// ValidateDeleteVPC validates that the DeleteVPC functionality works correctly.
func ValidateDeleteVPC(ctx context.Context, client CloudMaintainVPC, attrs DeleteVPCArgs) error {
	err := client.DeleteVPC(ctx, attrs)
	if err != nil {
		return err
	}
	return nil
}
