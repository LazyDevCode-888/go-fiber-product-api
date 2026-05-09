package controllers

import (
	"fiber-crud/config"
	"fiber-crud/models"

	"github.com/gofiber/fiber/v2"
)

func CreateProduct(c *fiber.Ctx) error {

	var product models.Product
	if err := c.BodyParser(&product); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid request",
		})
	}

	userID := c.Locals("user_id").(uint)
	product.UserID = userID

	result := config.DB.Create(&product)
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Cannot create product: " + result.Error.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Create Success",
		"data":    product,
	})
}

func GetProducts(c *fiber.Ctx) error {

	userID := c.Locals("user_id").(uint)

	var products []models.Product

	config.DB.Where("user_id = ?", userID).Find(&products)

	return c.JSON(fiber.Map{
		"message": "Success",
		"data":    products,
	})
}

func GetProductByID(c *fiber.Ctx) error {

	id := c.Params("id")
	userID := c.Locals("user_id").(uint)

	var product models.Product

	result := config.DB.Where(
		"id = ? AND user_id = ?",
		id,
		userID,
	).First(&product)

	if result.Error != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "Product not found",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Success",
		"data":    product,
	})
}

func UpdateProduct(c *fiber.Ctx) error {

	id := c.Params("id")
	userID := c.Locals("user_id").(uint)

	var product models.Product

	result := config.DB.Where(
		"id = ? AND user_id = ?",
		id,
		userID,
	).First(&product)

	if result.Error != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "Product not found",
		})
	}

	if err := c.BodyParser(&product); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid request",
		})
	}

	config.DB.Save(&product)

	return c.JSON(fiber.Map{
		"message": "Update Success",
		"data":    product,
	})
}

func DeleteProduct(c *fiber.Ctx) error {

	id := c.Params("id")
	userID := c.Locals("user_id").(uint)

	var product models.Product

	result := config.DB.Where(
		"id = ? AND user_id = ?",
		id,
		userID,
	).First(&product)

	if result.Error != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "Product not found",
		})
	}

	config.DB.Delete(&product)

	return c.JSON(fiber.Map{
		"message": "Delete Success",
	})
}