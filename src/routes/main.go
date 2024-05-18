package routes

import (
	"gofiber_pijar/src/controller"
	"gofiber_pijar/src/middlewares"

	"github.com/gofiber/fiber/v2"
)

func Router(app *fiber.App) {
	//product
	app.Get("/products", controller.GetAllProducts)
	app.Get("/products/:id", controller.GetProductById)
	app.Post("/product", controller.CreateProduct)
	app.Put("/product/:id", controller.UpdateProduct)
	app.Delete("/product/:id", controller.DeleteProduct)

	//Profiles
	app.Get("/profiles", controller.GetAllProfiles)
	app.Get("/profiles/:id", controller.GetProfileById)
	app.Post("/profile", controller.CreateProfile)
	app.Put("/profile/:id", controller.UpdateProfile)
	app.Delete("/profile/:id", controller.DeleteProfile)

	//category
	// app.Get("/categories", controller.GetAllCategories)
	app.Get("/categories", middlewares.JwtMiddleware(), controller.GetAllCategories)

	app.Get("/categories/:id", controller.GetCategoryById)
	app.Post("/category", controller.CreateCategory)
	app.Put("/category/:id", controller.UpdateCategory)
	app.Delete("/category/:id", controller.DeleteCategory)

	//user
	app.Post("/register", controller.RegisterUser)
	app.Post("/login", controller.LoginUser)
	app.Post("/refreshToken", controller.RefreshToken)
}
