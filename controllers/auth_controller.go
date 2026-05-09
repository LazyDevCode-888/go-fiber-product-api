package controllers

import (
	"fiber-crud/config"
	"fiber-crud/models"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *fiber.Ctx) error {

	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid request",
		})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(user.Password),
		14,
	)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Error hashing password",
		})
	}

	user.Password = string(hashedPassword)

	result := config.DB.Create(&user)
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Cannot create user: " + result.Error.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Register Success",
		"data": fiber.Map{
			"id":       user.ID,
			"username": user.Username,
			"email":    user.Email,
			"fullname": user.Fullname,
			"age":      user.Age,
		},
	})
}

func Login(c *fiber.Ctx) error {

	type LoginInput struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var input LoginInput

	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid request",
		})
	}

	var user models.User
	result := config.DB.Where(
		"email = ?",
		input.Email,
	).First(&user)

	if result.Error != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "User not found",
		})
	}

	err := bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(input.Password),
	)

	if err != nil {
		return c.Status(401).JSON(fiber.Map{
			"message": "Invalid Password",
		})
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"user_id": user.ID,
			"email":   user.Email,
			"exp": time.Now().Add(
				time.Hour * 24,
			).Unix(),
		},
	)

	tokenString, err := token.SignedString(
		[]byte(config.GetJWTSecret()),
	)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Cannot create token",
		})
	}

	c.Cookie(&fiber.Cookie{
		Name:     "jwt",
		Value:    tokenString,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,  // JavaScript เข้าถึงไม่ได้ (ป้องกัน XSS)
		Secure:   false, // เปลี่ยนเป็น true เมื่อใช้ HTTPS
		SameSite: "Lax",
	})

	return c.JSON(fiber.Map{
		"message": "Login Success",
		"token":   tokenString,
	})
}

func Logout(c *fiber.Ctx) error {
	c.Cookie(&fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour), // หมดอายุทันที
		HTTPOnly: true,
	})

	return c.JSON(fiber.Map{
		"message": "Logout Success",
	})
}