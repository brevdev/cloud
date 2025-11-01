package rsa

import (
	"crypto/ed25519"
	"crypto/rsa"
	"crypto/x509"
	"fmt"

	"golang.org/x/crypto/ssh"
)

// BytesToRSAKey parses a byte slice into an RSA private key.
// It supports OpenSSH, PKCS8, and PKCS1 formats.
func BytesToRSAKey(keyBytes []byte) (any, error) {
	// The key may be in OpenSSH format
	key, err := ssh.ParseRawPrivateKey(keyBytes)
	if err == nil {
		// No error, so we can parse into a key object
		switch k := key.(type) {
		case *rsa.PrivateKey, *ed25519.PrivateKey:
			return k, nil
		default:
			return nil, fmt.Errorf("key is not an RSA private key")
		}
	}

	// The key may be in PKCS8 format
	key, err = x509.ParsePKCS8PrivateKey(keyBytes)
	if err == nil {
		// No error, so we can parse into a key object
		switch k := key.(type) {
		case *rsa.PrivateKey, *ed25519.PrivateKey:
			return k, nil
		default:
			return nil, fmt.Errorf("key is not an RSA private key")
		}
	}

	// The key may be in PKCS1 format
	key, err = x509.ParsePKCS1PrivateKey(keyBytes)
	if err == nil {
		return key, nil
	}

	// The key is not in any of the supported formats
	return nil, fmt.Errorf("key is not an RSA private key")
}
