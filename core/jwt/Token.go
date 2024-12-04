package jwt

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/juheth/to-do/core/models"
)

// Token function parses and validates the JWT token
func Token(tk string) (bool, error) {
	// Load private key from environment variable
	privateKey := os.Getenv("PRIVATE_KEY")
	if privateKey == "" {
		return false, errors.New("private key not found in environment variables")
	}

	// Convert private key into bytes for signing verification
	privateKeyByte := []byte(privateKey)

	// Create a structure for claims that will be parsed
	var claims models.Claim

	// Validate Bearer token format
	if !strings.HasPrefix(tk, "Bearer ") {
		return false, errors.New("Bearer token is missing or malformed")
	}

	// Trim 'Bearer ' prefix and extra spaces from token string
	tk = strings.TrimSpace(strings.TrimPrefix(tk, "Bearer "))

	// Parse and validate the token using the claims and private key
	token, err := jwt.ParseWithClaims(tk, claims, func(t *jwt.Token) (interface{}, error) {
		return privateKeyByte, nil
	})

	// Error parsing token
	if err != nil {
		return false, fmt.Errorf("error parsing token: %w", err)
	}

	// Validate the token's validity
	if !token.Valid {
		return false, errors.New("invalid token")
	}

	// If everything is valid, return true
	return true, nil
}
