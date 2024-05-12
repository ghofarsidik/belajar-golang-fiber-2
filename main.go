package main

import (
	"gofiber_pijar/src/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	routes.Router(app)
	app.Listen(":3000")
}
