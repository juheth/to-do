package jwt

import (
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
	"github.com/juheth/to-do/core/models"
)

func JWT(user *models.User) (string, error) {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error al cargar el archivo .env: %v", err)
	}

	PrivateKey := os.Getenv("PRIVATE_KEY")
	privateKeyByte := []byte(PrivateKey)

	payload := jwt.MapClaims{
		"user":     user.Name,
		"email":    user.Email,
		"id":       user.ID,
		"Date_Exp": time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	tokenStr, err := token.SignedString(privateKeyByte)
	if err != nil {
		log.Printf("Error al firmar el token: %v", err)
		return "", err
	}

	return tokenStr, nil
}
