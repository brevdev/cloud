package v1

import (
	"context"

	v1 "github.com/brevdev/cloud/v1"
)

func (c *NebiusClient) AddFirewallRulesToInstance(_ context.Context, _ v1.AddFirewallRulesToInstanceArgs) error {
	return v1.ErrNotImplemented
}

func (c *NebiusClient) RevokeSecurityGroupRules(_ context.Context, _ v1.RevokeSecurityGroupRuleArgs) error {
	return v1.ErrNotImplemented
}
