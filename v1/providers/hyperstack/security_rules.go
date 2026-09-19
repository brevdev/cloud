package hyperstack

import (
	"fmt"
	"net"

	virtualmachine "github.com/NexGenCloud/hyperstack-sdk-go/lib/virtual_machine"

	v1 "github.com/brevdev/cloud/v1"
)

func makeSecurityRules(rules v1.FirewallRules) ([]virtualmachine.CreateSecurityRulePayload, error) {
	securityRules := make([]virtualmachine.CreateSecurityRulePayload, 0, len(rules.IngressRules))
	seen := make(map[string]bool)

	addRule := func(ipRange, etherType string, fromPort, toPort int) {
		key := fmt.Sprintf("%s|%s|%d|%d", etherType, ipRange, fromPort, toPort)
		if seen[key] {
			return
		}
		seen[key] = true
		securityRules = append(securityRules, virtualmachine.CreateSecurityRulePayload{
			Direction:      "ingress",
			Ethertype:      etherType,
			Protocol:       virtualmachine.Tcp,
			RemoteIpPrefix: ipRange,
			PortRangeMin:   &fromPort,
			PortRangeMax:   &toPort,
		})
	}

	for _, rule := range rules.IngressRules {
		fromPort := int(rule.FromPort)
		toPort := int(rule.ToPort)
		if fromPort < 1 || toPort < fromPort || toPort > 65535 {
			return nil, fmt.Errorf("invalid hyperstack ingress port range %d-%d", fromPort, toPort)
		}
		for _, ipRange := range rule.IPRanges {
			ip, _, err := net.ParseCIDR(ipRange)
			if err != nil {
				return nil, fmt.Errorf("parse hyperstack ingress CIDR %q: %w", ipRange, err)
			}
			etherType := "IPv4"
			if ip.To4() == nil {
				etherType = "IPv6"
			}
			addRule(ipRange, etherType, fromPort, toPort)
		}
	}
	return securityRules, nil
}
