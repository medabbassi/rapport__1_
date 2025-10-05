// internal/auth/keys.go
package auth

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

// GenerateSecureToken creates a cryptographically secure random token
func GenerateSecureToken(length int) (string, error) {
	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		return "", fmt.Errorf("could not generate token: %w", err)
	}
	return base64.URLEncoding.EncodeToString(bytes), nil
}

// GenerateAPIKey creates a formatted API key with prefix
func GenerateAPIKey() (string, error) {
	token, err := GenerateSecureToken(32)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("ak_%s", token), nil
}

// GenerateAccessKey generates a secure random access key based on the project name, total length 32
func GenerateAccessKey(projectName string) (string, error) {
	const totalLength = 32
	nameLen := len(projectName)
	if nameLen >= totalLength {
		return projectName[:totalLength], nil
	}
	randomLen := totalLength - nameLen
	bytes := make([]byte, randomLen)
	if _, err := rand.Read(bytes); err != nil {
		return "", fmt.Errorf("could not generate random bytes: %w", err)
	}
	randomStr := base64.URLEncoding.EncodeToString(bytes)
	// Ensure the final string is exactly 32 chars
	accessKey := projectName + randomStr
	if len(accessKey) > totalLength {
		accessKey = accessKey[:totalLength]
	}
	return accessKey, nil
}
