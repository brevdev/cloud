package v1

import (
	"context"
	"fmt"
	"time"
)

type CloudModifyFirewall interface {
	AddFirewallRulesToInstance(ctx context.Context, args AddFirewallRulesToInstanceArgs) error
	RevokeSecurityGroupRules(ctx context.Context, args RevokeSecurityGroupRuleArgs) error
}

type AddFirewallRulesToInstanceArgs struct {
	InstanceID    CloudProviderInstanceID
	FirewallRules FirewallRules
}

type RevokeSecurityGroupRuleArgs struct {
	InstanceID           CloudProviderInstanceID
	SecurityGroupRuleIDs []string
}

type FirewallRules struct {
	IngressRules []FirewallRule
	EgressRules  []FirewallRule
}

type FirewallRule struct {
	ID       string // ignored when creating a new rule
	FromPort int32
	ToPort   int32
	IPRanges []string
}

type PortMapping struct {
	FromPort int
	ToPort   int
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

// WaitForVPCPredicate waits for the VPC to satisfy the predicate function. If the predicate returns true, the loop breaks.
type WaitForVPCPredicateOpts struct {
	Predicate func(vpc *VPC) bool
	Timeout   time.Duration
	Interval  time.Duration
}

func WaitForVPCPredicate(ctx context.Context, client CloudMaintainVPC, attrs GetVPCArgs, opts WaitForVPCPredicateOpts) error {
	ctx, cancel := context.WithTimeout(ctx, opts.Timeout)
	defer cancel()

	ticker := time.NewTicker(opts.Interval)
	defer ticker.Stop()

	fmt.Printf("Entering WaitForVPCPredicate, timeout: %s, interval: %s\n", opts.Timeout.String(), opts.Interval.String())
	for {
		vpc, err := client.GetVPC(ctx, attrs)
		if err != nil {
			return err
		}

		if opts.Predicate(vpc) {
			break
		}
		fmt.Printf("Waiting %s for VPC to satisfy predicate\n", opts.Interval.String())
		select {
		case <-ctx.Done():
			return fmt.Errorf("timeout waiting for VPC to satisfy predicate")
		case <-ticker.C:
			continue
		}
	}
	return nil
}

func ValidateDeleteVPC(ctx context.Context, client CloudMaintainVPC, attrs DeleteVPCArgs) error {
	err := client.DeleteVPC(ctx, attrs)
	if err != nil {
		return err
	}
	return nil
}
