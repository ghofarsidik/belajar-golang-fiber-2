package helpers

import (
	"fmt"
	"gofiber_pijar/src/models"
	"os"
	"regexp"

	"github.com/go-playground/validator/v10"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type ErrorResponse struct {
	FailedField string `json:"failed_field"`
	Tag         string `json:"tag"`
	Value       string `json:"value"`
}

func ValidateStruct(param any) []*ErrorResponse {
	var errors []*ErrorResponse
	validate := validator.New()

	err := validate.Struct(param)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			element := ErrorResponse{
				FailedField: err.StructNamespace(),
				Tag:         err.Tag(),
				Value:       err.Param(),
			}
			errors = append(errors, &element)
		}
	}

	if user, ok := param.(models.User); ok {
		// Validasi password jika param adalah struct User
		passwordErrors := ValidatePassword(user.Password)
		errors = append(errors, passwordErrors...)
	}

	if user, ok := param.(models.User); ok {
		// Validasi password jika param adalah struct User
		emailErrors := ValidateEmail(user.Email)
		errors = append(errors, emailErrors...)
	}

	return errors
}

func ValidatePassword(password string) []*ErrorResponse {
	var errors []*ErrorResponse

	// validasi untuk memastikan tidak ada spasi
	noSpace := regexp.MustCompile(`^\S+$`)
	if !noSpace.MatchString(password) {
		errors = append(errors, &ErrorResponse{
			FailedField: "Password",
			Tag:         "password",
			Value:       "Password must not contain spaces",
		})
		return errors
	}

	// validasi untuk memastikan setidaknya satu huruf kecil
	hasLower := regexp.MustCompile(`[a-z]`)
	if !hasLower.MatchString(password) {
		errors = append(errors, &ErrorResponse{
			FailedField: "Password",
			Tag:         "password",
			Value:       "Password must contain at least one lowercase letter",
		})
		return errors
	}

	// validasi untuk memastikan setidaknya satu huruf besar
	hasUpper := regexp.MustCompile(`[A-Z]`)
	if !hasUpper.MatchString(password) {
		errors = append(errors, &ErrorResponse{
			FailedField: "Password",
			Tag:         "password",
			Value:       "Password must contain at least one uppercase letter",
		})
		return errors
	}

	// validasi untuk memastikan setidaknya satu angka
	hasNumber := regexp.MustCompile(`\d`)
	if !hasNumber.MatchString(password) {
		errors = append(errors, &ErrorResponse{
			FailedField: "Password",
			Tag:         "password",
			Value:       "Password must contain at least one number",
		})
		return errors
	}

	// validasi untuk memastikan setidaknya satu karakter spesial
	hasSpecial := regexp.MustCompile(`[\W_]`)
	if !hasSpecial.MatchString(password) {
		errors = append(errors, &ErrorResponse{
			FailedField: "Password",
			Tag:         "password",
			Value:       "Password must contain at least one special character",
		})
		return errors
	}

	return errors
}

func ValidateEmail(email string) []*ErrorResponse {
	var errors []*ErrorResponse

	db, err := InitDB()
	if err != nil {
		errors = append(errors, &ErrorResponse{
			FailedField: "Email",
			Tag:         "email",
			Value:       "Error connecting to database",
		})
		return errors
	}

	if !EmailUnique(db, email) {
		errors = append(errors, &ErrorResponse{
			FailedField: "Email",
			Tag:         "unique",
			Value:       "Email already exists",
		})
		return errors
	}

	return errors
}

func InitDB() (*gorm.DB, error) {
	dbURL := os.Getenv("URL")

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	return db, nil
}

func EmailUnique(db *gorm.DB, email string) bool {
	var user models.User

	if err := db.Where("email =?", email).First(&user).Error; err != nil {
		return true
	}
	return false
}
