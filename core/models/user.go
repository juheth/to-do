package models

type User struct {
	Id       int
	Name     string `json:"name" gorm:"primaryKey"`
	Email    string `json:"email"`
	Password string `json:"-"`
}
