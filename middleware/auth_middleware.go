package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"

	"fiber-crud/config"
)

func Protected() fiber.Handler {

	return func(c *fiber.Ctx) error {

		var tokenString string

		cookieToken := c.Cookies("jwt")

		if cookieToken != "" {
			tokenString = cookieToken
		} else {
			authHeader := c.Get("Authorization")

			if authHeader == "" {
				return c.Status(401).JSON(fiber.Map{
					"message": "Missing Token",
				})
			}

			tokenString = strings.Replace(
				authHeader,
				"Bearer ",
				"",
				1,
			)
		}

		token, err := jwt.Parse(
			tokenString,
			func(token *jwt.Token) (interface{}, error) {
				return []byte(
					config.GetJWTSecret(),
				), nil
			},
		)

		if err != nil || !token.Valid {
			return c.Status(401).JSON(fiber.Map{
				"message": "Invalid Token",
			})
		}

		claims := token.Claims.(jwt.MapClaims)

		c.Locals(
			"user_id",
			uint(claims["user_id"].(float64)),
		)

		return c.Next()
	}
}