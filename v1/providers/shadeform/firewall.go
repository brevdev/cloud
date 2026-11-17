package v1

import (
	"encoding/base64"
	"fmt"
	"strings"

	v1 "github.com/brevdev/cloud/v1"
)

const (
	ufwForceReset           = "ufw --force reset"
	ufwDefaultDropIncoming  = "ufw default deny incoming"
	ufwDefaultAllowOutgoing = "ufw default allow outgoing"
	ufwDefaultAllowPort22   = "ufw allow 22/tcp"
	ufwDefaultAllowPort2222 = "ufw allow 2222/tcp"
	ufwForceEnable          = "ufw --force enable"

	// Ensure DOCKER-USER exists before clearing it. Docker normally creates this
	// chain, but firewall setup can run before Docker has initialized iptables.
	ipTablesCreateDockerUserChain = "iptables -N DOCKER-USER || true"

	// Clear DOCKER-USER policy.
	ipTablesResetDockerUserChain = "iptables -F DOCKER-USER"

	// Allow return traffic.
	ipTablesAllowDockerUserOutbound = "iptables -A DOCKER-USER -m conntrack --ctstate ESTABLISHED,RELATED -j ACCEPT"

	// Allow containers to initiate outbound traffic (default bridge + user-defined bridges).
	ipTablesAllowDockerUserOutboundInit0 = "iptables -A DOCKER-USER -i docker0 ! -o docker0 -j ACCEPT"
	ipTablesAllowDockerUserOutboundInit1 = "iptables -A DOCKER-USER -i br+     ! -o br+     -j ACCEPT"
	ipTablesAllowDockerUserOutboundInit2 = "iptables -A DOCKER-USER -i cni+    ! -o cni+    -j ACCEPT"
	ipTablesAllowDockerUserOutboundInit3 = "iptables -A DOCKER-USER -i cali+   ! -o cali+   -j ACCEPT"

	// Allow container-to-container on the same bridge.
	ipTablesAllowDockerUserDockerToDocker0 = "iptables -A DOCKER-USER -i docker0 -o docker0 -j ACCEPT"
	ipTablesAllowDockerUserDockerToDocker1 = "iptables -A DOCKER-USER -i br+     -o br+     -j ACCEPT"
	ipTablesAllowDockerUserDockerToDocker2 = "iptables -A DOCKER-USER -i cni+    -o cni+    -j ACCEPT"
	ipTablesAllowDockerUserDockerToDocker3 = "iptables -A DOCKER-USER -i cali+   -o cali+   -j ACCEPT"

	// Allow inbound traffic on the loopback interface.
	ipTablesAllowDockerUserInpboundLoopback = "iptables -A DOCKER-USER -i lo -j ACCEPT"

	// Allow inbound traffic on the netbird interface.
	ipTablesAllowDockerUserInpboundNetbird = "iptables -A DOCKER-USER -i wt0 -j ACCEPT"

	// Allow external inbound TCP traffic to any container port 22 (SSH)
	ipTablesAllowDockerUserContainerSSH = "iptables -A DOCKER-USER -p tcp --dport 22 -j ACCEPT"

	// Drop everything else.
	ipTablesDropDockerUserInbound = "iptables -A DOCKER-USER -j DROP"
	ipTablesReturnDockerUser      = "iptables -A DOCKER-USER -j RETURN"
)

func (c *ShadeformClient) GenerateBootstrapScript(cloud string, firewallRules v1.FirewallRules) (string, error) {
	var commands []string
	commands = append(commands, c.getUFWCommands(firewallRules)...)
	commands = append(commands, c.getIPTablesCommands()...)
	if shouldMountDataDisk(cloud) {
		commands = append(commands, mountDataDiskScript)
	}

	script := ""
	for _, command := range commands {
		script += fmt.Sprintf("%v\n", command)
	}

	encoded := base64.StdEncoding.EncodeToString([]byte(script))
	return encoded, nil
}

const cloudDenvr = "denvr"

const mountDataDiskScript = `
# --- Brev: mount extra data disk if a blank one is present ---
brev_mount_data_disk() {
  set -u
  MOUNT_DIR="/data"
  MIN_BYTES=$((100 * 1024 * 1024 * 1024))   # ignore disks < 100GiB

  if mountpoint -q "$MOUNT_DIR"; then
    echo "brev: $MOUNT_DIR already mounted; skipping"
    return 0
  fi

  ROOT_SRC="$(findmnt -no SOURCE / 2>/dev/null || true)"
  ROOT_DISK="$(lsblk -no PKNAME "$ROOT_SRC" 2>/dev/null | head -n1)"

  CAND=""; CAND_SIZE=0
  while read -r NAME SIZE TYPE FSTYPE; do
    [ "$TYPE" = "disk" ] || continue
    [ -n "$NAME" ] || continue
    [ "$NAME" = "$ROOT_DISK" ] && continue
    [ -n "$FSTYPE" ] && continue
    [ -n "$(lsblk -no NAME "/dev/$NAME" | tail -n +2)" ] && continue
    blkid "/dev/$NAME" >/dev/null 2>&1 && continue
    [ "$SIZE" -lt "$MIN_BYTES" ] && continue
    if [ "$SIZE" -gt "$CAND_SIZE" ]; then CAND="$NAME"; CAND_SIZE="$SIZE"; fi
  done <<EOF
$(lsblk -bdno NAME,SIZE,TYPE,FSTYPE)
EOF

  [ -n "$CAND" ] || { echo "brev: no blank data disk found; skipping"; return 0; }

  DEV="/dev/$CAND"
  echo "brev: formatting + mounting $DEV at $MOUNT_DIR"
  mkdir -p "$MOUNT_DIR"
  mkfs.ext4 -F "$DEV" || { echo "brev: mkfs failed on $DEV"; return 0; }
  mount "$DEV" "$MOUNT_DIR" || { echo "brev: mount failed on $DEV"; return 0; }
  chmod 777 "$MOUNT_DIR"

  UUID="$(blkid -s UUID -o value "$DEV")"
  if [ -n "$UUID" ] && ! grep -q "$UUID" /etc/fstab; then
    echo "UUID=$UUID $MOUNT_DIR ext4 defaults,nofail,x-systemd.device-timeout=5 0 2" >> /etc/fstab
  fi
  echo "brev: mounted $DEV ($MOUNT_DIR)"
}
brev_mount_data_disk || true
`

func shouldMountDataDisk(cloud string) bool {
	return strings.EqualFold(strings.TrimSpace(cloud), cloudDenvr)
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

func (c *ShadeformClient) getIPTablesCommands() []string {
	commands := []string{
		ipTablesCreateDockerUserChain,
		ipTablesResetDockerUserChain,
		ipTablesAllowDockerUserOutbound,
		ipTablesAllowDockerUserOutboundInit0,
		ipTablesAllowDockerUserOutboundInit1,
		ipTablesAllowDockerUserOutboundInit2, // TODO: add these back in when we have a way to test it
		ipTablesAllowDockerUserOutboundInit3,
		ipTablesAllowDockerUserDockerToDocker0,
		ipTablesAllowDockerUserDockerToDocker1,
		ipTablesAllowDockerUserDockerToDocker2,
		ipTablesAllowDockerUserDockerToDocker3,
		ipTablesAllowDockerUserInpboundLoopback,
		ipTablesAllowDockerUserInpboundNetbird,
		ipTablesAllowDockerUserContainerSSH,
		ipTablesDropDockerUserInbound,
		ipTablesReturnDockerUser, // Expected by Docker
	}
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
