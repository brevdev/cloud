package v1

import (
	"crypto/ed25519"
	"crypto/rand"
	"testing"

	v1 "github.com/brevdev/cloud/v1"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/ssh"
	"gopkg.in/yaml.v3"
)

func TestCloudInitSSHCertificateTrustPreservesFirewall(t *testing.T) {
	public, _, err := ed25519.GenerateKey(rand.Reader)
	require.NoError(t, err)
	key, err := ssh.NewPublicKey(public)
	require.NoError(t, err)
	publicKey := string(ssh.MarshalAuthorizedKey(key))
	args := v1.CreateInstanceAttrs{
		PublicKey: "legacy-user-key",
	}
	for _, certificate := range []bool{false, true} {
		if certificate {
			args.SSHCertificateAuthority = &v1.SSHCertificateAuthority{
				PublicKey:       publicKey,
				PrincipalPrefix: "brev:v1:vm:env-nebius:login:",
			}
		}
		userData, err := generateCloudInitUserData(args)
		require.NoError(t, err)
		var config struct {
			Keys     []string `yaml:"ssh_authorized_keys"`
			Commands []string `yaml:"runcmd"`
		}
		require.NoError(t, yaml.Unmarshal([]byte(userData), &config))
		require.Len(t, config.Keys, 1)
		require.Contains(t, config.Commands, "ufw allow 22/tcp")
		require.Contains(t, config.Commands, "sudo /usr/local/sbin/brev-apply-docker-firewall.sh || true")
		if !certificate {
			require.Equal(t, []string{"legacy-user-key"}, config.Keys)
			continue
		}
		parsed, _, options, rest, err := ssh.ParseAuthorizedKey([]byte(config.Keys[0]))
		require.NoError(t, err)
		require.Equal(t, key.Marshal(), parsed.Marshal())
		require.ElementsMatch(t, []string{"cert-authority", `principals="brev:v1:vm:env-nebius:login:ubuntu"`}, options)
		require.Empty(t, rest)
		require.NotContains(t, userData, "legacy-user-key")
		require.Contains(t, userData, dockerFirewallScriptPath)
		require.Contains(t, userData, dockerFirewallDropInPath)
	}
	args.SSHCertificateAuthority.PublicKey = "invalid-key"
	_, err = (&NebiusClient{}).CreateInstance(t.Context(), args)
	require.ErrorContains(t, err, "invalid SSH CA public key")
	args.SSHCertificateAuthority.PublicKey = publicKey
	args.SSHCertificateAuthority.PrincipalPrefix = ""
	_, err = generateCloudInitUserData(args)
	require.Error(t, err)
}
