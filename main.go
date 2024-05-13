package main

import (
	"gofiber_pijar/src/configs"
	"gofiber_pijar/src/helpers"
	"gofiber_pijar/src/routes"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app := fiber.New()
	configs.InitDB()
	helpers.Migration()
	routes.Router(app)

	app.Listen(":3000")
}
