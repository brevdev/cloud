package sshbootstrap

import (
	"crypto/ed25519"
	"crypto/rand"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	v1 "github.com/brevdev/cloud/v1"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/ssh"
)

func TestInstallForKeyOwnersUsesProviderAccountAndPreservesRotation(t *testing.T) {
	public, _, err := ed25519.GenerateKey(rand.Reader)
	require.NoError(t, err)
	key, err := ssh.NewPublicKey(public)
	require.NoError(t, err)
	ca := v1.SSHCertificateAuthority{
		PublicKey:       string(ssh.MarshalAuthorizedKey(key)),
		PrincipalPrefix: "brev:v1:vm:env-shadeform:login:",
	}
	script, err := InstallForKeyOwners(ca)
	require.NoError(t, err)
	for _, owner := range []string{"ubuntu", "shadeform", "no-matching-key"} {
		t.Run(owner, func(t *testing.T) {
			home := t.TempDir()
			original := map[string]string{}
			for _, user := range []string{"ubuntu", "shadeform", "root"} {
				original[user] = "unrelated-key\n"
				if user == owner {
					original[user] += strings.TrimSpace(ca.PublicKey) + " provider-comment\n"
				}
				if user == "root" {
					original[user] += `command="use the login user" ` + ca.PublicKey
				}
				require.NoError(t, os.MkdirAll(filepath.Join(home, user, ".ssh"), 0o700))
				require.NoError(t, os.WriteFile(filepath.Join(home, user, ".ssh", "authorized_keys"), []byte(original[user]), 0o600))
			}
			const utilities = `getent() {
    for user in ubuntu shadeform root; do
        printf '%s:x:1000:1000::%s/%s:/bin/bash\n' "$user" "$BOOTSTRAP_HOME" "$user"
    done
}
chown() { printf '%s:%s\n' "$1" "$2" >> "$BOOTSTRAP_HOME/ownership"; }
flock() { :; }
`
			run := func() ([]byte, error) {
				t.Helper()
				command := exec.CommandContext(t.Context(), "bash")
				command.Env = append(os.Environ(), "BOOTSTRAP_HOME="+home)
				command.Stdin = strings.NewReader(utilities + script + "printf continued")
				return command.CombinedOutput()
			}
			output, err := run()
			if owner == "no-matching-key" {
				require.Error(t, err, "missing provider bootstrap key must fail closed")
				return
			}
			require.NoError(t, err, string(output))
			require.Equal(t, "continued", string(output))
			for user, before := range original {
				contents, err := os.ReadFile(filepath.Join(home, user, ".ssh", "authorized_keys")) //nolint:gosec // test ok
				require.NoError(t, err)
				want := before
				if user == owner {
					want, err = AuthorizedKey(ca, owner)
					require.NoError(t, err)
					want += "\n"
				}
				require.Equal(t, want, string(contents), user)
			}
			path := filepath.Join(home, owner, ".ssh", "authorized_keys")
			ownership, err := os.ReadFile(filepath.Join(home, "ownership")) //nolint:gosec // test ok
			require.NoError(t, err)
			require.Contains(t, string(ownership), owner+":"+path+".lock\n", "later healing runs as the login user")
			lock, err := os.Stat(path + ".lock")
			require.NoError(t, err)
			require.Equal(t, os.FileMode(0o600), lock.Mode().Perm())
			require.NoError(t, os.WriteFile(path, []byte("rotated-ca\n"), 0o600))
			output, err = run()
			require.NoError(t, err, string(output))
			require.Equal(t, "continued", string(output))
			contents, err := os.ReadFile(path) //nolint:gosec // test ok
			require.NoError(t, err)
			require.Equal(t, "rotated-ca\n", string(contents))
		})
	}
}
