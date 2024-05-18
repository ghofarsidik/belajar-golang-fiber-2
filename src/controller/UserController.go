package controller

import (
	"gofiber_pijar/src/helpers"
	"gofiber_pijar/src/models"
	"os"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(c *fiber.Ctx) error {
	// var rawNewUser map[string]interface{}
	var newUser models.User
	if err := c.BodyParser(&newUser); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	// // xssmiddleware
	// rawNewUser = helpers.XSSMiddleware(rawNewUser)
	// var newUser models.User
	// mapstructure.Decode(rawNewUser, &newUser)

	//validate
	errors := helpers.ValidateStruct(newUser)
	if len(errors) > 0 {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(errors)
	}

	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	newUser.Password = string(hashPassword)

	models.PostUser(&newUser)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Register Successfully",
	})
}

func LoginUser(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	validateEmail := models.FindEmail(&user)
	if len(validateEmail) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Email is not Found",
		})
	}

	var passwordSecond string
	for _, user := range validateEmail {
		passwordSecond = user.Password
	}

	if err := bcrypt.CompareHashAndPassword([]byte(passwordSecond), []byte(user.Password)); err != nil || user.Password == "" {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Password invalid",
		})
	}

	jwtKey := os.Getenv("SECRETKEY")
	payload := map[string]interface{}{
		"email": user.Email,
	}

	token, err := helpers.GenerateToken(jwtKey, payload)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Coult not generate access token",
		})
	}

	refreshToken, err := helpers.GenerateRefreshToken(jwtKey, payload)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not generate refresh token",
		})
	}

	item := map[string]string{
		"Email":        user.Email,
		"Token":        token,
		"RefreshToken": refreshToken,
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Login successfully",
		"data":    item,
	})
}

func RefreshToken(c *fiber.Ctx) error {
	var input struct {
		RefreshToken string `json:"refresh_token"`
	}

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to parse request body",
		})
	}

	jwtKey := os.Getenv("SECRETKEY")
	token, err := helpers.GenerateToken(jwtKey, map[string]interface{}{"refreshToken": input.RefreshToken})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not generate access token",
		})
	}

	refreshToken, err := helpers.GenerateRefreshToken(jwtKey, map[string]interface{}{"refreshToken": input.RefreshToken})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not generate refresh token",
		})
	}

	item := map[string]string{
		"Token":        token,
		"RefreshToken": refreshToken,
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Refresh successfully",
		"data":    item,
	})
}
