package hyperstack

import (
	"context"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/NexGenCloud/hyperstack-sdk-go/lib/keypair"
	"golang.org/x/crypto/ssh"

	v1 "github.com/brevdev/cloud/v1"
)

type resolvedKeyPair struct {
	name      string
	managedID int
}

func (c *HyperstackClient) resolveKeyPair(ctx context.Context, attrs v1.CreateInstanceAttrs, environmentName string) (resolvedKeyPair, error) {
	keyPairName := attrs.RefID
	if attrs.KeyPairName != nil {
		keyPairName = *attrs.KeyPairName
	}
	keyPairName = strings.TrimSpace(keyPairName)
	if strings.TrimSpace(attrs.PublicKey) == "" {
		return resolvedKeyPair{name: keyPairName}, nil
	}

	publicKey, err := normalizeSSHPublicKey(attrs.PublicKey)
	if err != nil {
		return resolvedKeyPair{}, err
	}
	keyID, err := c.findKeyPair(ctx, keyPairName, environmentName, publicKey)
	if err != nil {
		return resolvedKeyPair{}, err
	}
	if keyID != 0 {
		resolved := resolvedKeyPair{name: keyPairName}
		if attrs.KeyPairName == nil {
			resolved.managedID = keyID
		}
		return resolved, nil
	}

	response, err := c.keypairs.ImportKeyPairWithResponse(ctx, keypair.ImportKeypairPayload{
		EnvironmentName: environmentName,
		Name:            keyPairName,
		PublicKey:       publicKey,
	})
	if err != nil {
		return resolvedKeyPair{}, wrapTransportError("import key pair", err)
	}
	if response.StatusCode() != http.StatusOK {
		return resolvedKeyPair{}, responseError("import key pair", response.StatusCode(), response.Body, nil)
	}
	if response.JSON200 == nil || response.JSON200.Keypair == nil || response.JSON200.Keypair.Id == nil || *response.JSON200.Keypair.Id <= 0 {
		return resolvedKeyPair{}, errors.New("hyperstack import key pair response did not contain a keypair ID")
	}
	return resolvedKeyPair{name: keyPairName, managedID: *response.JSON200.Keypair.Id}, nil
}

func (c *HyperstackClient) findKeyPair(ctx context.Context, keyName, environmentName, publicKey string) (int, error) {
	pageSize := strconv.Itoa(defaultPageSize)
	for page := 1; ; page++ {
		pageNumber := strconv.Itoa(page)
		response, err := c.keypairs.ListKeyPairsWithResponse(ctx, &keypair.ListKeyPairsParams{
			Page:     &pageNumber,
			PageSize: &pageSize,
			Search:   &keyName,
		})
		if err != nil {
			return 0, wrapTransportError("list key pairs", err)
		}
		if response.StatusCode() != http.StatusOK {
			return 0, responseError("list key pairs", response.StatusCode(), response.Body, nil)
		}
		if response.JSON200 == nil || response.JSON200.Keypairs == nil {
			return 0, errors.New("hyperstack list key pairs response did not contain data")
		}

		providerKeys := *response.JSON200.Keypairs
		keyID, matchErr := matchingKeyPair(providerKeys, keyName, environmentName, publicKey)
		if matchErr != nil || keyID != 0 {
			return keyID, matchErr
		}
		if len(providerKeys) < defaultPageSize {
			return 0, nil
		}
	}
}

func matchingKeyPair(providerKeys []keypair.KeypairFields, keyName, environmentName, publicKey string) (int, error) {
	for _, providerKey := range providerKeys {
		if stringValue(providerKey.Name) != keyName || providerKey.Environment == nil || stringValue(providerKey.Environment.Name) != environmentName {
			continue
		}
		existingKey, err := normalizeSSHPublicKey(stringValue(providerKey.PublicKey))
		if err != nil || existingKey != publicKey {
			return 0, fmt.Errorf("hyperstack key pair %q already exists with a different public key", keyName)
		}
		if providerKey.Id == nil || *providerKey.Id <= 0 {
			return 0, fmt.Errorf("hyperstack key pair %q did not contain an ID", keyName)
		}
		return *providerKey.Id, nil
	}
	return 0, nil
}

func (c *HyperstackClient) deleteManagedKeyPair(ctx context.Context, keyPairID int) error {
	response, err := c.keypairs.DeleteKeyPairWithResponse(ctx, keyPairID)
	if err != nil {
		return wrapTransportError("delete managed key pair", err)
	}
	if response.StatusCode() == http.StatusNotFound {
		return nil
	}
	if response.StatusCode() != http.StatusOK {
		return responseError("delete managed key pair", response.StatusCode(), response.Body, nil)
	}
	return nil
}

func normalizeSSHPublicKey(publicKey string) (string, error) {
	publicKey = strings.TrimSpace(publicKey)
	if key, _, _, _, err := ssh.ParseAuthorizedKey([]byte(publicKey)); err == nil {
		return strings.TrimSpace(string(ssh.MarshalAuthorizedKey(key))), nil
	}

	block, _ := pem.Decode([]byte(publicKey))
	if block == nil {
		return "", errors.New("hyperstack public key must be OpenSSH or PEM encoded")
	}
	parsedKey, pkixErr := x509.ParsePKIXPublicKey(block.Bytes)
	if pkixErr != nil {
		rsaKey, pkcs1Err := x509.ParsePKCS1PublicKey(block.Bytes)
		if pkcs1Err != nil {
			return "", fmt.Errorf("parse hyperstack PEM public key: %w", errors.Join(pkixErr, pkcs1Err))
		}
		parsedKey = rsaKey
	}
	key, err := ssh.NewPublicKey(parsedKey)
	if err != nil {
		return "", fmt.Errorf("convert hyperstack public key to OpenSSH: %w", err)
	}
	return strings.TrimSpace(string(ssh.MarshalAuthorizedKey(key))), nil
}
