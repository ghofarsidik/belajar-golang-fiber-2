package routes

import (
	"gofiber_pijar/src/controller"

	"github.com/gofiber/fiber/v2"
)

func Router(app *fiber.App) {
	app.Get("/products", controller.GetAllProducts)
	app.Get("/products/:id", controller.GetProductById)
	app.Post("/product", controller.CreateProduct)
	app.Put("/product/:id", controller.UpdateProduct)
	app.Delete("/product/:id", controller.DeleteProduct)
}
