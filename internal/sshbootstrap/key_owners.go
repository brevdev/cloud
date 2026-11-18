package sshbootstrap

import (
	"crypto/sha256"
	"fmt"
	"strings"

	v1 "github.com/brevdev/cloud/v1"
)

// InstallForKeyOwners discovers the provider's login account from the unique CA
// public key supplied to its create API, then replaces raw keys with scoped trust.
func InstallForKeyOwners(authority v1.SSHCertificateAuthority) (string, error) {
	if authority.PrincipalPrefix == "" || strings.ContainsAny(authority.PrincipalPrefix, "\"\\,\r\n") {
		return "", fmt.Errorf("invalid SSH CA principal prefix")
	}
	publicKey, err := CanonicalPublicKey(authority.PublicKey)
	if err != nil {
		return "", fmt.Errorf("invalid SSH CA public key: %w", err)
	}
	return fmt.Sprintf(installForKeyOwnersScript, shellQuote(publicKey), shellQuote(authority.PrincipalPrefix),
		sha256.Sum256([]byte(authority.PrincipalPrefix+publicKey))), nil
}

func shellQuote(value string) string {
	return "'" + strings.ReplaceAll(value, "'", "'\"'\"'") + "'"
}

// A CA-specific marker prevents a replayed launch script from undoing rotation.
// Use the health probe's authorized_keys lock and publish each file atomically.
const installForKeyOwnersScript = `#!/bin/bash
(
set -eu
umask 077
CA_PUBLIC_KEY=%s
PRINCIPAL_PREFIX=%s
CA_TYPE=${CA_PUBLIC_KEY%%%% *}
CA_DATA=${CA_PUBLIC_KEY#* }
INSTALLED=0
while IFS=: read -r SSH_USER _ _ _ _ HOME_DIR _; do
    AUTH_KEYS="$HOME_DIR/.ssh/authorized_keys"
    MARKER="$HOME_DIR/.ssh/.brev-ca-bootstrap-%x"
    if [ -f "$MARKER" ]; then
        INSTALLED=1
        continue
    fi
    [ -f "$AUTH_KEYS" ] || continue
    if ! awk -v type="$CA_TYPE" -v key="$CA_DATA" '$1 == type && $2 == key { found=1 } END { exit !found }' "$AUTH_KEYS"; then
        continue
    fi
    (
        exec 9>"$AUTH_KEYS.lock"
        flock 9
        chown "$SSH_USER" "$AUTH_KEYS.lock"
        chmod 600 "$AUTH_KEYS.lock"
        [ ! -f "$MARKER" ] || exit 0
        TMP_KEYS=$(mktemp "$AUTH_KEYS.tmp.XXXXXX")
        trap 'rm -f "$TMP_KEYS"' EXIT
        printf 'cert-authority,principals="%%s%%s" %%s\n' "$PRINCIPAL_PREFIX" "$SSH_USER" "$CA_PUBLIC_KEY" > "$TMP_KEYS"
        chmod 600 "$TMP_KEYS"
        chown "$SSH_USER" "$TMP_KEYS"
        mv "$TMP_KEYS" "$AUTH_KEYS"
        touch "$MARKER"
    )
    INSTALLED=1
done < <(getent passwd)
[ "$INSTALLED" -eq 1 ] || { echo "No account contains the SSH CA bootstrap key" >&2; exit 1; }
)
BREV_CA_STATUS=$?
[ "$BREV_CA_STATUS" -eq 0 ] || exit "$BREV_CA_STATUS"
`
