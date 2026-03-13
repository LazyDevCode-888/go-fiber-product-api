package routes

import (
	"product-crud/controllers"

	"github.com/gofiber/fiber/v3"
)

func ProductRoutes(app *fiber.App) {

	api := app.Group("/api")

	api.Post("/products", controllers.CreateProduct)
	api.Get("/products", controllers.GetProducts)
	api.Get("/products/:id", controllers.GetProduct)
	api.Put("/products/:id", controllers.UpdateProduct)
	api.Delete("/products/:id", controllers.DeleteProduct)

}
