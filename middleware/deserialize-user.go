package middleware

import (
	"fmt"
	"strings"
	"github.com/gofiber/fiber/v2"
  "github.com/golang-jwt/jwt/v5"

	"github.com/obiewalker/nomada-api/pkg/database"
  "github.com/obiewalker/nomada-api/pkg/database/models"
	"github.com/obiewalker/nomada-api/config"
)

func DeserializeUser(c *fiber.Ctx) error {
	var tokenString string
	authorization := c.Get("Authorization")

	if strings.HasPrefix(authorization, "Bearer ") {
		tokenString = strings.TrimPrefix(authorization, "Bearer ")
	} else if c.Cookies("token") != "" {
		tokenString = c.Cookies("token")
	}

	if tokenString == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "fail", "message": "You are not logged in"})
	}

	tokenByte, err := jwt.Parse(tokenString, func(jwtToken *jwt.Token) (interface{}, error) {
		if _, ok := jwtToken.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %s", jwtToken.Header["alg"])
		}

		return []byte(config.GetEnvStr("JWT_SECRET")), nil
	})
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "fail", "message": fmt.Sprintf("invalidate token: %v", err)})
	}

	claims, ok := tokenByte.Claims.(jwt.MapClaims)
	if !ok || !tokenByte.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "fail", "message": "Invalid token claim."})

	}

	var user model.User
	db := database.Instance.Db

	db.First(&user, "id = ?", fmt.Sprint(claims["sub"]))

	if user.ID != claims["sub"] {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"status": "fail", "message": "Unauthorised Token."})
	}

	c.Locals("user", model.FilterUserRecord(&user))

	return c.Next()
}
