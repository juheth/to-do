package jwt

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/juheth/to-do/core/models"
)

func Token(tk string) (bool, error) {

	privateKey := os.Getenv("PRIVATE_KEY")
	if privateKey == "" {
		return false, errors.New("private key not found in environment variables")
	}

	privateKeyByte := []byte(privateKey)

	var claims models.Claim

	if !strings.HasPrefix(tk, "Bearer ") {
		return false, errors.New("Bearer token is missing or malformed")
	}

	tk = strings.TrimSpace(strings.TrimPrefix(tk, "Bearer "))

	token, err := jwt.ParseWithClaims(tk, claims, func(t *jwt.Token) (interface{}, error) {
		return privateKeyByte, nil
	})

	if err != nil {
		return false, fmt.Errorf("error parsing token: %w", err)
	}

	if !token.Valid {
		return false, errors.New("invalid token")
	}

	return true, nil
}
