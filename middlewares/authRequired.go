package middlewares

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func AuthRequired(c *fiber.Ctx) error {
	// รับ token จาก Authorization header (Bearer Token)
	tokenString := c.Get("Authorization")
	if tokenString == "" {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	tokenString = tokenString[len("Bearer "):]

	jwtSecret := os.Getenv("JWT_SECRET")

	token, err := jwt.ParseWithClaims(tokenString, jwt.MapClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})

	if err != nil || !token.Valid {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	return c.Next()
}
