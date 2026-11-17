package v1

import (
	"encoding/base64"
	"strings"
	"testing"

	v1 "github.com/brevdev/cloud/v1"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestShadeformIPTablesCommandsCreateDockerUserChainBeforeFlush(t *testing.T) {
	client := &ShadeformClient{}
	commands := strings.Join(client.getIPTablesCommands(), "\n")

	createChainIndex := strings.Index(commands, "iptables -N DOCKER-USER")
	flushChainIndex := strings.Index(commands, "iptables -F DOCKER-USER")

	assert.Greater(t, createChainIndex, -1)
	assert.Greater(t, flushChainIndex, createChainIndex)
}

func decodeBootstrapScript(t *testing.T, cloud string) string {
	t.Helper()
	client := &ShadeformClient{}
	encoded, err := client.GenerateBootstrapScript(cloud, v1.FirewallRules{})
	require.NoError(t, err)
	decoded, err := base64.StdEncoding.DecodeString(encoded)
	require.NoError(t, err)
	return string(decoded)
}

func TestShadeformBootstrapScriptMountsDataDiskForDenvr(t *testing.T) {
	script := decodeBootstrapScript(t, "denvr")

	// Firewall is always applied.
	assert.Contains(t, script, "ufw --force enable")
	// Denvr also gets the data-disk mount logic.
	assert.Contains(t, script, "brev_mount_data_disk")
	assert.Contains(t, script, "mkfs.ext4 -F")
	assert.Contains(t, script, "/data")
	assert.Contains(t, script, "mountpoint -q")
}

func TestShadeformBootstrapScriptSkipsDataDiskForOtherClouds(t *testing.T) {
	script := decodeBootstrapScript(t, "hyperstack")

	// Firewall is still applied.
	assert.Contains(t, script, "ufw --force enable")
	// No data-disk mount logic for non-denvr clouds.
	assert.NotContains(t, script, "brev_mount_data_disk")
}
