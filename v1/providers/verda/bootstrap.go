package verda

import (
	"fmt"
	"net"
	"strings"

	v1 "github.com/brevdev/cloud/v1"
)

const (
	dockerFirewallScriptPath = "/usr/local/sbin/brev-apply-docker-firewall.sh"
	dockerFirewallDropInPath = "/etc/systemd/system/docker.service.d/10-brev-firewall.conf"
)

func buildStartupScript(rules v1.FirewallRules, publicKey string) (string, error) {
	ufwRules, dockerRules, err := firewallRuleCommands(rules.IngressRules)
	if err != nil {
		return "", err
	}

	var script strings.Builder
	script.WriteString(`#!/bin/bash
set -u

if ! id brev >/dev/null 2>&1; then
  useradd --create-home --shell /bin/bash brev
fi
install -d -m 0700 -o brev -g brev /home/brev/.ssh
cat > /home/brev/.ssh/authorized_keys <<'BREV_AUTHORIZED_KEYS'
`)
	script.WriteString(publicKey)
	script.WriteString(`
BREV_AUTHORIZED_KEYS
chown brev:brev /home/brev/.ssh/authorized_keys
chmod 0600 /home/brev/.ssh/authorized_keys
echo 'brev ALL=(ALL) NOPASSWD:ALL' > /etc/sudoers.d/brev
chmod 0440 /etc/sudoers.d/brev

if ! command -v ufw >/dev/null 2>&1; then
  apt-get update -y
  DEBIAN_FRONTEND=noninteractive apt-get install -y ufw iptables
fi

mkdir -p /usr/local/sbin /etc/systemd/system/docker.service.d
cat > ` + dockerFirewallScriptPath + ` <<'BREV_FIREWALL'
#!/bin/sh
iptables -N DOCKER-USER 2>/dev/null || true
iptables -F DOCKER-USER || true
iptables -A DOCKER-USER -m conntrack --ctstate ESTABLISHED,RELATED -j ACCEPT
iptables -A DOCKER-USER -i docker0 ! -o docker0 -j ACCEPT
iptables -A DOCKER-USER -i br+ ! -o br+ -j ACCEPT
iptables -A DOCKER-USER -i cni+ ! -o cni+ -j ACCEPT
iptables -A DOCKER-USER -i cali+ ! -o cali+ -j ACCEPT
iptables -A DOCKER-USER -i docker0 -o docker0 -j ACCEPT
iptables -A DOCKER-USER -i br+ -o br+ -j ACCEPT
iptables -A DOCKER-USER -i cni+ -o cni+ -j ACCEPT
iptables -A DOCKER-USER -i cali+ -o cali+ -j ACCEPT
iptables -A DOCKER-USER -i lo -j ACCEPT
iptables -A DOCKER-USER -i wt0 -j ACCEPT
`)
	for _, command := range dockerRules {
		script.WriteString(command)
		script.WriteByte('\n')
	}
	script.WriteString(`iptables -A DOCKER-USER -j DROP
exit 0
BREV_FIREWALL
chmod 0755 ` + dockerFirewallScriptPath + `

cat > ` + dockerFirewallDropInPath + ` <<'BREV_DROP_IN'
[Service]
ExecStartPost=-` + dockerFirewallScriptPath + `
BREV_DROP_IN

systemctl daemon-reload || true
ufw --force reset
ufw default deny incoming
ufw default allow outgoing
ufw allow 22/tcp
`)
	for _, command := range ufwRules {
		script.WriteString(command)
		script.WriteByte('\n')
	}
	script.WriteString(`ufw --force enable
` + dockerFirewallScriptPath + ` || true
`)
	return script.String(), nil
}

func firewallRuleCommands(rules []v1.FirewallRule) ([]string, []string, error) {
	var ufwCommands []string
	var dockerCommands []string
	for _, rule := range rules {
		if rule.FromPort < 1 || rule.ToPort > 65535 || rule.FromPort > rule.ToPort {
			return nil, nil, fmt.Errorf(
				"invalid firewall port range %d-%d",
				rule.FromPort,
				rule.ToPort,
			)
		}

		sources := rule.IPRanges
		if len(sources) == 0 {
			sources = []string{"0.0.0.0/0"}
		}
		for _, source := range sources {
			ip, network, err := net.ParseCIDR(source)
			if err != nil {
				return nil, nil, fmt.Errorf("invalid firewall CIDR %q: %w", source, err)
			}
			if ip.To4() == nil {
				return nil, nil, fmt.Errorf("IPv6 firewall CIDR %q is not supported", source)
			}
			source = network.String()

			if rule.FromPort == rule.ToPort {
				ufwCommands = append(ufwCommands, fmt.Sprintf(
					"ufw allow from %s to any port %d",
					source,
					rule.FromPort,
				))
			} else {
				for _, protocol := range []string{"tcp", "udp"} {
					ufwCommands = append(ufwCommands, fmt.Sprintf(
						"ufw allow from %s to any port %d:%d proto %s",
						source,
						rule.FromPort,
						rule.ToPort,
						protocol,
					))
				}
			}

			portSpec := fmt.Sprintf("%d", rule.FromPort)
			if rule.FromPort != rule.ToPort {
				portSpec = fmt.Sprintf("%d:%d", rule.FromPort, rule.ToPort)
			}
			for _, protocol := range []string{"tcp", "udp"} {
				dockerCommands = append(dockerCommands, fmt.Sprintf(
					"iptables -A DOCKER-USER -s %s -p %s --dport %s -j ACCEPT",
					source,
					protocol,
					portSpec,
				))
			}
		}
	}
	return ufwCommands, dockerCommands, nil
}
