package nebius

import (
	"context"

	"github.com/brevdev/cloud"
)

func (c *NebiusClient) AddFirewallRulesToInstance(_ context.Context, _ cloud.AddFirewallRulesToInstanceArgs) error {
	return cloud.ErrNotImplemented
}

func (c *NebiusClient) RevokeSecurityGroupRules(_ context.Context, _ cloud.RevokeSecurityGroupRuleArgs) error {
	return cloud.ErrNotImplemented
}
