package v1

import (
	"encoding/base64"
	"fmt"

	v1 "github.com/brevdev/cloud/v1"
)

const (
	ufwForceReset           = "ufw --force reset"
	ufwDefaultDropIncoming  = "ufw default deny incoming"
	ufwDefaultAllowOutgoing = "ufw default allow outgoing"
	ufwDefaultAllowPort22   = "ufw allow 22/tcp"
	ufwDefaultAllowPort2222 = "ufw allow 2222/tcp"
	ufwForceEnable          = "ufw --force enable"

	// Clear DOCKER-USER policy.
	ipTablesResetDockerUserChain = "iptables -F DOCKER-USER"

	// Allow return traffic.
	ipTablesAllowDockerUserOutbound = "iptables -A DOCKER-USER -m conntrack --ctstate ESTABLISHED,RELATED -j ACCEPT"

	// Allow containers to initiate outbound traffic (default bridge + user-defined bridges).
	ipTablesAllowDockerUserOutboundInit0 = "iptables -A DOCKER-USER -i docker0 ! -o docker0 -j ACCEPT"
	ipTablesAllowDockerUserOutboundInit1 = "iptables -A DOCKER-USER -i br+     ! -o br+     -j ACCEPT"

	// Allow container-to-container on the same bridge.
	ipTablesAllowDockerUserDockerToDocker0 = "iptables -A DOCKER-USER -i docker0 -o docker0 -j ACCEPT"
	ipTablesAllowDockerUserDockerToDocker1 = "iptables -A DOCKER-USER -i br+     -o br+     -j ACCEPT"

	// Allow inbound traffic on the loopback interface.
	ipTablesAllowDockerUserInpboundLoopback = "iptables -A DOCKER-USER -i lo -j ACCEPT"

	// Drop everything else.
	ipTablesDropDockerUserInbound = "iptables -A DOCKER-USER -j DROP"
	ipTablesReturnDockerUser      = "iptables -A DOCKER-USER -j RETURN"
)

func (c *ShadeformClient) GenerateFirewallScript(firewallRules v1.FirewallRules) (string, error) {
	var commands []string
	commands = append(commands, c.getUFWCommands(firewallRules)...)
	commands = append(commands, c.getIPTablesCommands(firewallRules)...)

	script := "#!/bin/bash\nset -e\n"
	for _, command := range commands {
		script += fmt.Sprintf("%v\n", command)
	}

	encoded := base64.StdEncoding.EncodeToString([]byte(script))
	return encoded, nil
}

func (c *ShadeformClient) getUFWCommands(firewallRules v1.FirewallRules) []string {
	commands := []string{
		ufwForceReset,
		ufwDefaultDropIncoming,
		ufwDefaultAllowOutgoing,
		ufwDefaultAllowPort22,
		ufwDefaultAllowPort2222,
	}

	for _, firewallRule := range firewallRules.IngressRules {
		commands = append(commands, c.convertIngressFirewallRuleToUfwCommand(firewallRule)...)
	}

	for _, firewallRule := range firewallRules.EgressRules {
		commands = append(commands, c.convertEgressFirewallRuleToUfwCommand(firewallRule)...)
	}

	// Add the enable command
	commands = append(commands, ufwForceEnable)

	return commands
}

func (c *ShadeformClient) getIPTablesCommands(firewallRules v1.FirewallRules) []string {
	commands := []string{
		// Wait for Docker to be ready and DOCKER-USER chain to exist (max 5 minutes)
		`echo "Waiting for Docker and DOCKER-USER chain..."`,
		`for i in $(seq 1 150); do iptables -L DOCKER-USER -n >/dev/null 2>&1 && break || sleep 2; done`,
		`iptables -L DOCKER-USER -n >/dev/null 2>&1 || { echo "ERROR: DOCKER-USER chain not found after 5 minutes"; exit 1; }`,
		`echo "Docker is ready, applying firewall rules..."`,

		ipTablesResetDockerUserChain,
		ipTablesAllowDockerUserOutbound,
		ipTablesAllowDockerUserOutboundInit0,
		ipTablesAllowDockerUserOutboundInit1,
		ipTablesAllowDockerUserDockerToDocker0,
		ipTablesAllowDockerUserDockerToDocker1,
		ipTablesAllowDockerUserInpboundLoopback,
	}

	// Add ACCEPT rules for user-specified ingress ports BEFORE the DROP rule
	// This allows users to explicitly open ports for Docker containers
	for _, rule := range firewallRules.IngressRules {
		commands = append(commands, c.convertIngressFirewallRuleToIPTablesCommand(rule)...)
	}

	// Drop everything else and return
	commands = append(commands, ipTablesDropDockerUserInbound)
	commands = append(commands, ipTablesReturnDockerUser) // Expected by Docker

	return commands
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

// convertIngressFirewallRuleToIPTablesCommand generates iptables rules for DOCKER-USER chain
// to allow inbound traffic to Docker containers on specified ports.
// This is necessary because UFW rules only affect the INPUT chain (host traffic),
// not the FORWARD chain (container traffic).
func (c *ShadeformClient) convertIngressFirewallRuleToIPTablesCommand(firewallRule v1.FirewallRule) []string {
	cmds := []string{}

	// Generate port specification for iptables
	var portSpec string
	if firewallRule.FromPort == firewallRule.ToPort {
		portSpec = fmt.Sprintf("--dport %d", firewallRule.FromPort)
	} else {
		portSpec = fmt.Sprintf("--dport %d:%d", firewallRule.FromPort, firewallRule.ToPort)
	}

	if len(firewallRule.IPRanges) == 0 {
		// Allow from any source
		cmds = append(cmds, fmt.Sprintf("iptables -A DOCKER-USER -p tcp %s -j ACCEPT", portSpec))
		cmds = append(cmds, fmt.Sprintf("iptables -A DOCKER-USER -p udp %s -j ACCEPT", portSpec))
	} else {
		// Allow from specific IP ranges only
		for _, ipRange := range firewallRule.IPRanges {
			cmds = append(cmds, fmt.Sprintf("iptables -A DOCKER-USER -p tcp -s %s %s -j ACCEPT", ipRange, portSpec))
			cmds = append(cmds, fmt.Sprintf("iptables -A DOCKER-USER -p udp -s %s %s -j ACCEPT", ipRange, portSpec))
		}
	}

	return cmds
}
