package v2

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"net/http"
	"net/netip"
	"net/url"
	"slices"
	"sort"
	"strconv"
	"strings"
	"time"

	v1 "github.com/brevdev/cloud/v1"
	"github.com/google/uuid"
)

const firewallVersionHeader = "X-SFC-Firewall-Version"

type firewallRule struct {
	Direction string `json:"direction"`
	Protocol  string `json:"protocol"`
	Port      string `json:"port,omitempty"`
	Source    string `json:"source"`
}

type firewallResponse struct {
	ID    string         `json:"id"`
	Rules []firewallRule `json:"rules"`
}

func isAPIStatus(err error, status int) bool {
	var responseErr *apiError
	return errors.As(err, &responseErr) && responseErr.statusCode == status
}

func sshFirewallRules() []firewallRule {
	return []firewallRule{
		{Direction: "ingress", Protocol: "tcp", Port: "22", Source: "0.0.0.0/0"},
		{Direction: "ingress", Protocol: "tcp", Port: "2222", Source: "0.0.0.0/0"},
	}
}

func expandFirewallRules(rules v1.FirewallRules) ([]firewallRule, error) {
	for _, rule := range rules.EgressRules {
		if !isUnrestrictedEgress(rule) {
			return nil, fmt.Errorf("SFCompute does not support restricting outbound traffic")
		}
	}
	var result []firewallRule
	for _, rule := range rules.IngressRules {
		if rule.FromPort < 0 || rule.ToPort > 65535 || rule.FromPort > rule.ToPort {
			return nil, fmt.Errorf("invalid ingress port range %d-%d", rule.FromPort, rule.ToPort)
		}
		port := strconv.Itoa(int(rule.FromPort))
		if rule.FromPort != rule.ToPort {
			port += "-" + strconv.Itoa(int(rule.ToPort))
		}
		sources := rule.IPRanges
		if len(sources) == 0 {
			sources = []string{"0.0.0.0/0"}
		}
		for _, source := range sources {
			prefix, err := netip.ParsePrefix(source)
			if err != nil || !prefix.Addr().Is4() {
				return nil, fmt.Errorf("ingress source must be an IPv4 CIDR: %q", source)
			}
			// Brev's provider contract has no protocol field; a port rule covers TCP and UDP.
			for _, protocol := range []string{"tcp", "udp"} {
				result = append(result, firewallRule{
					Direction: "ingress", Protocol: protocol, Port: port, Source: prefix.Masked().String(),
				})
			}
		}
	}
	return result, nil
}

func isUnrestrictedEgress(rule v1.FirewallRule) bool {
	return rule.FromPort == 0 && rule.ToPort == 65535 &&
		(len(rule.IPRanges) == 0 || slices.Equal(rule.IPRanges, []string{"0.0.0.0/0"}))
}

func deduplicateRules(rules []firewallRule) []firewallRule {
	result := make([]firewallRule, 0, len(rules))
	seen := make(map[firewallRule]bool)
	for _, rule := range rules {
		if !seen[rule] {
			result = append(result, rule)
			seen[rule] = true
		}
	}
	return result
}

func (c *SFCClientV2) createInstanceFirewall(ctx context.Context, rules v1.FirewallRules) (*firewallResponse, error) {
	expanded, err := expandFirewallRules(rules)
	if err != nil {
		return nil, err
	}
	request := struct {
		Name      string         `json:"name"`
		Workspace string         `json:"workspace"`
		Rules     []firewallRule `json:"rules"`
	}{
		Name: "brev-" + uuid.NewString(), Workspace: c.GetWorkspaceResourcePath(),
		Rules: deduplicateRules(append(sshFirewallRules(), expanded...)),
	}
	var response firewallResponse
	_, err = c.client.doRequest(ctx, http.MethodPost, "/integrations/brev/v1/firewalls", nil, request, &response, nil)
	if err != nil {
		return nil, err
	}
	if response.ID == "" {
		return nil, fmt.Errorf("SFCompute returned a firewall without an ID")
	}
	return &response, nil
}

func (c *SFCClientV2) cleanupFirewall(ctx context.Context, id string, prior error) error {
	// A canceled create request must not cancel cleanup of its separate firewall.
	ctx, cancel := context.WithTimeout(context.WithoutCancel(ctx), 30*time.Second)
	defer cancel()
	_, err := c.client.doRequest(ctx, http.MethodDelete, "/integrations/brev/v1/firewalls/"+url.PathEscape(id), nil, nil, nil, nil)
	if isAPIStatus(err, http.StatusNotFound) {
		err = nil
	}
	if err != nil {
		err = fmt.Errorf("clean up firewall %s: %w", id, err)
	}
	return errors.Join(prior, err)
}

func (c *SFCClientV2) getFirewall(ctx context.Context, id string) (*firewallResponse, string, error) {
	var response firewallResponse
	headers, err := c.client.doRequest(ctx, http.MethodGet, "/integrations/brev/v1/firewalls/"+url.PathEscape(id), nil, nil, &response, nil)
	if err != nil {
		return nil, "", err
	}
	return &response, headers.Get(firewallVersionHeader), nil
}

func ingressRuleID(rule firewallRule) string {
	hash := sha256.Sum256([]byte(rule.Port + "|" + rule.Source))
	return "sfc-ingress-" + hex.EncodeToString(hash[:])
}

func (c *SFCClientV2) loadInstanceFirewall(ctx context.Context, source *instanceResponse, instance *v1.Instance) error {
	if source.Firewall == "" || source.Status == instanceStatusTerminated {
		return nil
	}
	firewall, _, err := c.getFirewall(ctx, source.Firewall)
	if err != nil {
		return err
	}
	rules := make(map[string]v1.FirewallRule)
	for _, rule := range firewall.Rules {
		if rule.Direction != "ingress" || rule.Port == "" || slices.Contains(sshFirewallRules(), rule) {
			continue
		}
		ports := strings.SplitN(rule.Port, "-", 2)
		from, err := strconv.ParseInt(ports[0], 10, 32)
		if err != nil {
			return fmt.Errorf("invalid firewall port from SFCompute: %w", err)
		}
		to := from
		if len(ports) == 2 {
			to, err = strconv.ParseInt(ports[1], 10, 32)
			if err != nil {
				return fmt.Errorf("invalid firewall port from SFCompute: %w", err)
			}
		}
		id := ingressRuleID(rule)
		rules[id] = v1.FirewallRule{ID: id, FromPort: int32(from), ToPort: int32(to), IPRanges: []string{rule.Source}}
	}
	ids := make([]string, 0, len(rules))
	for id := range rules {
		ids = append(ids, id)
	}
	sort.Strings(ids)
	for _, id := range ids {
		instance.FirewallRules.IngressRules = append(instance.FirewallRules.IngressRules, rules[id])
	}
	return nil
}

func (c *SFCClientV2) AddFirewallRulesToInstance(ctx context.Context, args v1.AddFirewallRulesToInstanceArgs) error {
	rules, err := expandFirewallRules(args.FirewallRules)
	if err != nil {
		return err
	}
	return c.updateInstanceFirewall(ctx, args.InstanceID, func(current []firewallRule) []firewallRule {
		return deduplicateRules(append(current, rules...))
	})
}

func (c *SFCClientV2) RevokeSecurityGroupRules(ctx context.Context, args v1.RevokeSecurityGroupRuleArgs) error {
	return c.updateInstanceFirewall(ctx, args.InstanceID, func(current []firewallRule) []firewallRule {
		return slices.DeleteFunc(current, func(rule firewallRule) bool {
			return rule.Direction == "ingress" && !slices.Contains(sshFirewallRules(), rule) &&
				slices.Contains(args.SecurityGroupRuleIDs, ingressRuleID(rule))
		})
	})
}

func (c *SFCClientV2) updateInstanceFirewall(ctx context.Context, id v1.CloudProviderInstanceID, change func([]firewallRule) []firewallRule) error {
	instance, err := c.client.getInstance(ctx, string(id))
	if err != nil {
		return err
	}
	if !instance.EnablePublicIPv4 || instance.Firewall == "" || instance.Status == instanceStatusTerminated {
		return fmt.Errorf("configurable firewall requires a live instance created with public networking: %w", v1.ErrNotImplemented)
	}
	for range 5 {
		firewall, version, err := c.getFirewall(ctx, instance.Firewall)
		if err != nil {
			return err
		}
		if version == "" {
			return fmt.Errorf("SFCompute API does not support conditional firewall updates")
		}
		request := struct {
			Rules []firewallRule `json:"rules"`
		}{Rules: change(firewall.Rules)}
		_, err = c.client.doRequest(ctx, http.MethodPut, "/integrations/brev/v1/firewalls/"+url.PathEscape(instance.Firewall), nil,
			request, nil, http.Header{firewallVersionHeader: []string{version}})
		if !isAPIStatus(err, http.StatusConflict) {
			return err
		}
	}
	return fmt.Errorf("firewall changed during five update attempts; retry the operation")
}
