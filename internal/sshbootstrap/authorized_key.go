// Package sshbootstrap renders resource-scoped SSH CA trust for provider bootstrapping.
package sshbootstrap

import (
	"fmt"
	"strings"

	v1 "github.com/brevdev/cloud/v1"
	"golang.org/x/crypto/ssh"
)

// AuthorizedKey binds CA trust to the resource and its actual SSH login account.
func AuthorizedKey(authority v1.SSHCertificateAuthority, user string) (string, error) {
	if authority.PrincipalPrefix == "" || user == "" {
		return "", fmt.Errorf("SSH CA bootstrap requires a principal prefix and SSH user")
	}
	publicKey, err := CanonicalPublicKey(authority.PublicKey)
	if err != nil {
		return "", fmt.Errorf("invalid SSH CA public key: %w", err)
	}
	return fmt.Sprintf("cert-authority,principals=%q %s", authority.PrincipalPrefix+user, publicKey), nil
}

// CanonicalPublicKey discards authorized-key options, comments, and trailing content.
func CanonicalPublicKey(raw string) (string, error) {
	key, _, _, _, err := ssh.ParseAuthorizedKey([]byte(raw))
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(ssh.MarshalAuthorizedKey(key))), nil
}
