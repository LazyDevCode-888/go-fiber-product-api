package main

import (
	"go-fiber-product-api/config"
	"go-fiber-product-api/models"
	"go-fiber-product-api/routes"
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
)

func main() {

	app := fiber.New()

	// CORS
	app.Use(cors.New())

	// database
	config.ConnectDB()

	config.DB.AutoMigrate(&models.Product{})

	// routes
	routes.ProductRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
