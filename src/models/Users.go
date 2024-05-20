package models

import (
	"gofiber_pijar/src/configs"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
	Role     string `json:"role"`
}

func PostUser(item *User) error {
	result := configs.DB.Create(&item)
	return result.Error
}

func FindEmail(input *User) []User {
	items := []User{}
	configs.DB.Raw("SELECT * FROM users WHERE email = ?", input.Email).Scan(&items)
	return items
}
