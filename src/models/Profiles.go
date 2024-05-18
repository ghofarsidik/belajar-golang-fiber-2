package models

import (
	"gofiber_pijar/src/configs"

	"gorm.io/gorm"
)

type Profile struct {
	gorm.Model
	Name   string `json:"name" validate:"required"`
	Photo  string `json:"photo"`
	Email  string `json:"email" validate:"required,email"`
	Phone  string `json:"phone" validate:"required,number"`
	Gender string `json:"gender" validate:"required"`
}

func SelectAllProfiles() []*Profile {
	var items []*Profile
	configs.DB.Find(&items)
	return items
}

func SelectProfileByID(id int) *Profile {
	var item Profile
	configs.DB.First(&item, "id = ?", id)
	return &item
}

func PostProfile(newProfile *Profile) error {
	result := configs.DB.Create(&newProfile)
	return result.Error
}

func UpdateProfile(id int, item *Profile) error {
	result := configs.DB.Model(&Profile{}).Where("id = ?", id).Updates(item)
	return result.Error
}

func DeleteProfile(id int) error {
	result := configs.DB.Delete(&Profile{}, "id = ?", id)
	return result.Error
}
