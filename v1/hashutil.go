package v1

import (
	"crypto/sha256"
	"encoding/hex"
)

func HashSensitiveString(s string) (string, error) {
	// Preprocess the password with SHA-256
	sha256Hash := sha256.Sum256([]byte(s))
	sha256HashString := hex.EncodeToString(sha256Hash[:])
	return sha256HashString, nil
}
