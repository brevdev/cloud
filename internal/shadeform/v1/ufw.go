package v1

import (
	"encoding/base64"
	"fmt"
	v1 "github.com/brevdev/cloud/pkg/v1"
)

const (
	ufwForceReset           = "ufw --force reset"
	ufwDefaultDropIncoming  = "ufw default deny incoming"
	ufwDefaultAllowOutgoing = "ufw default allow outgoing"
	ufwDefaultAllowPort22   = "ufw allow 22/tcp"
	ufwDefaultAllowPort2222 = "ufw allow 2222/tcp"
	ufwForceEnable          = "ufw --force enable"
)

func (c *ShadeformClient) GenerateFirewallScript(firewallRules v1.FirewallRules) (string, error) {
	commands := []string{ufwForceReset, ufwDefaultDropIncoming, ufwDefaultAllowOutgoing, ufwDefaultAllowPort22, ufwDefaultAllowPort2222}

	for _, firewallRule := range firewallRules.IngressRules {
		commands = append(commands, c.convertIngressFirewallRuleToUfwCommand(firewallRule)...)
	}

	for _, firewallRule := range firewallRules.EgressRules {
		commands = append(commands, c.convertEgressFirewallRuleToUfwCommand(firewallRule)...)
	}

	// Add the enable command
	commands = append(commands, ufwForceEnable)

	script := ""
	for _, command := range commands {
		script = script + fmt.Sprintf("%v\n", command)
	}

	encoded := base64.StdEncoding.EncodeToString([]byte(script))
	return encoded, nil
}

func (c *ShadeformClient) convertIngressFirewallRuleToUfwCommand(firewallRule v1.FirewallRule) []string {
	cmds := []string{}
	portSpecs := []string{}
	if firewallRule.FromPort == firewallRule.ToPort {
		portSpecs = append(portSpecs, fmt.Sprintf("port %d", firewallRule.FromPort))
	} else {
		// port ranges require two separate rules for tcp and udp
		portSpecs = append(portSpecs, fmt.Sprintf("port %d:%d proto tcp", firewallRule.FromPort, firewallRule.ToPort))
		portSpecs = append(portSpecs, fmt.Sprintf("port %d:%d proto udp", firewallRule.FromPort, firewallRule.ToPort))
	}

	if len(firewallRule.IPRanges) == 0 {
		for _, portSpec := range portSpecs {
			cmds = append(cmds, fmt.Sprintf("ufw allow in from any to any %s", portSpec))
		}
	}

	for _, ipRange := range firewallRule.IPRanges {
		for _, portSpec := range portSpecs {
			cmds = append(cmds, fmt.Sprintf("ufw allow in from %s to any %s", ipRange, portSpec))
		}
	}
	return cmds
}

func (c *ShadeformClient) convertEgressFirewallRuleToUfwCommand(firewallRule v1.FirewallRule) []string {
	cmds := []string{}
	portSpecs := []string{}
	if firewallRule.FromPort == firewallRule.ToPort {
		portSpecs = append(portSpecs, fmt.Sprintf("port %d", firewallRule.FromPort))
	} else {
		// port ranges require two separate rules for tcp and udp
		portSpecs = append(portSpecs, fmt.Sprintf("port %d:%d proto tcp", firewallRule.FromPort, firewallRule.ToPort))
		portSpecs = append(portSpecs, fmt.Sprintf("port %d:%d proto udp", firewallRule.FromPort, firewallRule.ToPort))
	}

	if len(firewallRule.IPRanges) == 0 {
		for _, portSpec := range portSpecs {
			cmds = append(cmds, fmt.Sprintf("ufw allow out to any %s", portSpec))
		}
	}

	for _, ipRange := range firewallRule.IPRanges {
		for _, portSpec := range portSpecs {
			cmds = append(cmds, fmt.Sprintf("ufw allow out to %s %s", ipRange, portSpec))
		}
	}
	return cmds
}
