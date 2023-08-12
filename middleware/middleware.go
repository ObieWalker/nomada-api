package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"github.com/obiewalker/nomada-api/config"
)

func AuthReq() func(*fiber.Ctx) error {
	cfg := basicauth.Config{
			Users: map[string]string{
				config.GetEnvStr("USERNAME"): config.GetEnvStr("PASSWORD"),
			},
		}
	err := basicauth.New(cfg);
	return err
}
