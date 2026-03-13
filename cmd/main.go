package main

import (
	"product-crud/config"
	"product-crud/models"
	"product-crud/routes"

	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v3"
)

func main() {

	app := fiber.New()

	app.Use(cors.New())

	config.ConnectDB()

	config.DB.AutoMigrate(&models.Product{})

	routes.ProductRoutes(app)

	app.Listen(":3000")
}
