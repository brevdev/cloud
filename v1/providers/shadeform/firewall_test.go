package v1

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShadeformIPTablesCommandsCreateDockerUserChainBeforeFlush(t *testing.T) {
	client := &ShadeformClient{}
	commands := strings.Join(client.getIPTablesCommands(), "\n")

	createChainIndex := strings.Index(commands, "iptables -N DOCKER-USER")
	flushChainIndex := strings.Index(commands, "iptables -F DOCKER-USER")

	assert.Greater(t, createChainIndex, -1)
	assert.Greater(t, flushChainIndex, createChainIndex)
}
