package routes

import (
	"github.com/gofiber/fiber/v2"

	"fiber-crud/controllers"
	"fiber-crud/middleware"
)

func ProductRoutes(app *fiber.App) {
	app.Post(
		"/products",
		middleware.Protected(),
		controllers.CreateProduct,
	)
	app.Get(
		"/products",
		middleware.Protected(),
		controllers.GetProducts,
	)
	app.Get(
		"/products/:id",
		middleware.Protected(),
		controllers.GetProductByID,
	)
	app.Put(
		"/products/:id",
		middleware.Protected(),
		controllers.UpdateProduct,
	)
	app.Delete(
		"/products/:id",
		middleware.Protected(),
		controllers.DeleteProduct,
	)
}

