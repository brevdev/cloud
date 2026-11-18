package v1

import (
	"context"
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/brevdev/cloud/internal/errors"
	"github.com/brevdev/cloud/internal/sshbootstrap"
	v1 "github.com/brevdev/cloud/v1"
	"github.com/google/uuid"
)

func (c *ShadeformClient) instanceLaunchScript(attrs v1.CreateInstanceAttrs) (string, error) {
	firewall, err := c.GenerateFirewallScript(attrs.FirewallRules)
	if err != nil || attrs.SSHCertificateAuthority == nil {
		return firewall, err
	}
	setup, err := sshbootstrap.InstallForKeyOwners(*attrs.SSHCertificateAuthority)
	if err != nil {
		return "", errors.WrapAndTrace(err)
	}
	decoded, err := base64.StdEncoding.DecodeString(firewall)
	if err != nil {
		return "", errors.WrapAndTrace(err)
	}
	return base64.StdEncoding.EncodeToString([]byte(setup + "\n" + string(decoded))), nil
}

func (c *ShadeformClient) getSSHBootstrapKey(ctx context.Context, attrs v1.CreateInstanceAttrs) (string, error) {
	if attrs.SSHCertificateAuthority != nil {
		return c.getOrCreateSSHCAKey(ctx, attrs)
	}
	keyPairName := attrs.RefID
	if attrs.KeyPairName != nil {
		keyPairName = *attrs.KeyPairName
	}
	if keyPairName == "" {
		keyPairName = uuid.New().String()
		c.logger.Debug(ctx, "No key pair name provided, generating new one", v1.LogField("keyPairName", keyPairName))
	}
	if attrs.PublicKey == "" {
		return "", nil
	}
	id, err := c.addSSHKey(ctx, keyPairName, attrs.PublicKey)
	if err != nil && !strings.Contains(err.Error(), "name must be unique") {
		return "", errors.WrapAndTrace(fmt.Errorf("failed to add SSH key: %w", err))
	}
	return id, nil
}

func (c *ShadeformClient) getOrCreateSSHCAKey(ctx context.Context, attrs v1.CreateInstanceAttrs) (string, error) {
	publicKey, err := sshbootstrap.CanonicalPublicKey(attrs.SSHCertificateAuthority.PublicKey)
	if err != nil {
		return "", errors.WrapAndTrace(err)
	}
	name := "brev-sshca-" + attrs.RefID
	id, err := c.addSSHKey(ctx, name, publicKey)
	if err == nil {
		if id == "" {
			return "", errors.New("Shadeform returned an empty SSH CA key ID")
		}
		return id, nil
	}
	if !strings.Contains(err.Error(), "name must be unique") {
		return "", errors.WrapAndTrace(err)
	}
	keys, response, err := c.client.DefaultAPI.SshKeys(c.makeAuthContext(ctx)).Execute()
	if response != nil && response.Body != nil {
		defer func() { _ = response.Body.Close() }()
	}
	if err != nil {
		return "", errors.WrapAndTrace(err)
	}
	if keys != nil {
		for _, key := range keys.SshKeys {
			if key.Name != name {
				continue
			}
			existing, err := sshbootstrap.CanonicalPublicKey(key.PublicKey)
			if err != nil || existing != publicKey || key.Id == "" {
				return "", errors.Errorf("existing Shadeform SSH CA key %q does not match", name)
			}
			return key.Id, nil
		}
	}
	return "", errors.Errorf("Shadeform SSH CA key %q not found after name conflict", name)
}
