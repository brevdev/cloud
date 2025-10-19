package v1

import "context"

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

func ValidateCreateVPC(ctx context.Context, client CloudMaintainVPC, attrs CreateVPCArgs) error {
	_, err := client.CreateVPC(ctx, attrs)
	if err != nil {
		return err
	}
	return nil
}

func ValidateGetVPC(ctx context.Context, client CloudMaintainVPC, attrs GetVPCArgs) error {
	_, err := client.GetVPC(ctx, attrs)
	if err != nil {
		return err
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
