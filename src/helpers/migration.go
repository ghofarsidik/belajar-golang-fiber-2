package helpers

import (
	"gofiber_pijar/src/configs"
	"gofiber_pijar/src/models"
)

func Migration() {
	configs.DB.AutoMigrate(&models.Product{}, &models.Profile{}, &models.Category{}, &models.User{})
}
