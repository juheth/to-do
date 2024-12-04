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

	// Cargar las variables del archivo .env
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error al cargar el archivo .env: %v", err)
	}

	// Obtener la clave privada desde el archivo .env
	PrivateKey := os.Getenv("PRIVATE_KEY")
	privateKeyByte := []byte(PrivateKey)

	// Crear el payload del token
	payload := jwt.MapClaims{
		"user":     user.Name,
		"email":    user.Email,
		"id":       user.ID,
		"Date_Exp": time.Now().Add(time.Hour * 24).Unix(), // Expira en 24 horas
	}

	// Crear el token con el algoritmo de firma HS256
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	// Firmar el token con la clave privada
	tokenStr, err := token.SignedString(privateKeyByte)
	if err != nil {
		log.Printf("Error al firmar el token: %v", err)
		return "", err
	}

	// Retornar el token firmado
	return tokenStr, nil
}
