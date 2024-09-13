package middleware

import (
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func JWTMiddleware(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error":"Unauthorized"})
	}

	tokenString := strings.Split(authHeader,"")[1]

	token,err := jwt.Parse(tokenString,func(token *jwt.Token)(interface{},error){
		if _,ok :=token.Method.(*jwt.SigningMethodHMAC); !ok{
			return nil,fiber.NewError(fiber.StatusUnauthorized,"unexpected signing method")
		}
		return []byte(os.Getenv("JWT_SECRET")),nil
	})

	if err != nil || !token.Valid{
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error":"Unaunthorized"})
	}

	return c.Next()
}