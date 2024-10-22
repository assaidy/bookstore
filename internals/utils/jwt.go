package utils

import (
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateJwtToken(id int, username string) (string, error) {
	claims := jwt.MapClaims{
		"id":    id,
		"admin": false,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}
	return tokenStr, nil
}

func GetUserIdFromContext(c *fiber.Ctx) (int, bool) {
	claims, ok := c.Locals("user").(*jwt.Token).Claims.(jwt.MapClaims)
	if !ok {
		return 0, false
	}
	id, ok := claims["id"].(float64)
	if !ok {
		return 0, false
	}
	return int(id), true
}
