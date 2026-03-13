package controllers

import (
	"product-crud/config"
	"product-crud/models"

	"github.com/gofiber/fiber/v3"
)

func CreateProduct(c fiber.Ctx) error {

	product := new(models.Product)

	if err := c.Bind().Body(product); err != nil {
		return err
	}

	config.DB.Create(&product)

	return c.JSON(product)
}

func GetProducts(c fiber.Ctx) error {

	var products []models.Product

	config.DB.Find(&products)

	return c.JSON(products)
}

func GetProduct(c fiber.Ctx) error {

	id := c.Params("id")

	var product models.Product

	config.DB.First(&product, id)

	return c.JSON(product)
}

func UpdateProduct(c fiber.Ctx) error {

	id := c.Params("id")

	var product models.Product

	config.DB.First(&product, id)

	if err := c.Bind().Body(&product); err != nil {
		return err
	}

	config.DB.Save(&product)

	return c.JSON(product)
}

func DeleteProduct(c fiber.Ctx) error {

	id := c.Params("id")

	config.DB.Delete(&models.Product{}, id)

	return c.JSON(fiber.Map{
		"message": "deleted success",
	})
}
