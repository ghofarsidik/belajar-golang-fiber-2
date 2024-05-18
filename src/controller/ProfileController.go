package controller

import (
	"fmt"
	"gofiber_pijar/src/helpers"
	"gofiber_pijar/src/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/mitchellh/mapstructure"
)

// menampilkan semua produk
func GetAllProfiles(c *fiber.Ctx) error {
	profiles := models.SelectAllProfiles()
	return c.JSON(profiles)
}

// menampilkan 1 produk
func GetProfileById(c *fiber.Ctx) error {
	paramID := c.Params("id")
	id, _ := strconv.Atoi(paramID)

	foundProfile := models.SelectProfileByID(id)

	if foundProfile == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Profile not found",
		})
	}

	return c.JSON(foundProfile)
}

// menambahkan produk
func CreateProfile(c *fiber.Ctx) error {
	var rawNewProfile map[string]interface{}
	if err := c.BodyParser(&rawNewProfile); err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
		return err
	}

	// XSSMiddleware
	rawNewProfile = helpers.XSSMiddleware(rawNewProfile)
	var newProfile models.Profile
	mapstructure.Decode(rawNewProfile, &newProfile)

	// validate
	errors := helpers.ValidateStruct(newProfile)
	if len(errors) > 0 {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(errors)
	}

	models.PostProfile(&newProfile)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Profile created successfully",
	})
}

// memperbaharui data profile
func UpdateProfile(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	var rawUpdateProfile map[string]interface{}
	if err := c.BodyParser(&rawUpdateProfile); err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid Request Body",
		})
		return err
	}

	// XSSMiddleware
	rawUpdateProfile = helpers.XSSMiddleware(rawUpdateProfile)
	var updateProfile models.Profile
	mapstructure.Decode(rawUpdateProfile, &updateProfile)

	// validate
	errors := helpers.ValidateStruct(updateProfile)
	if len(errors) > 0 {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(errors)
	}

	err := models.UpdateProfile(id, &updateProfile)

	if err == nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": fmt.Sprintf("Profile with ID %d updated successfully", id),
			"profile": updateProfile,
		})
	} else {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": fmt.Sprintf("Profile with ID %d not found", id),
		})
	}
}

// delete profile
func DeleteProfile(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	err := models.DeleteProfile(id)

	if err == nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": fmt.Sprintf("Profile with ID %d deleted successfully", id),
		})
	} else {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": fmt.Sprintf("Profile with ID %d not found", id),
		})
	}
}
