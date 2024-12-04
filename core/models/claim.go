package models

import "fmt"

// Claim representa los datos que estarán en el payload del JWT
type Claim struct {
	User  string `json:"user"`
	Email string `json:"email"`
	ID    int    `json:"id"`
	Exp   int64  `json:"exp"`
}

// Valid es el método requerido para que `Claim` implemente la interfaz `jwt.Claims`
// Este método debe devolver `nil` si el token es válido o un error si no lo es.
func (c Claim) Valid() error {
	// Aquí podrías agregar validaciones adicionales si las necesitas.
	if c.ID == 0 {
		return fmt.Errorf("invalid user ID")
	}
	return nil
}
