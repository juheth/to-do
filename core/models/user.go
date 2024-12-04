package models

type User struct {
	ID       int    `gorm:"primary_key;auto_increment"`
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
