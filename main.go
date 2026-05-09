package main

import (
	"fiber-crud/config"
	"fiber-crud/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {

	config.ConnectDB()
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	routes.ProductRoutes(app)
	routes.AuthRoutes(app)

	app.Listen(":3000")

}
