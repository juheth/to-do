package models

import "fmt"

type Claim struct {
	User  string `json:"user"`
	Email string `json:"email"`
	ID    int    `json:"id"`
	Exp   int64  `json:"exp"`
}

func (c Claim) Valid() error {
	if c.ID == 0 {
		return fmt.Errorf("invalid user ID")
	}
	return nil
}
