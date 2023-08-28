package main

import (
	"github.com/obiewalker/nomada-api/pkg/database"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/obiewalker/nomada-api/config"
)

func main() {

	database.ConnectDb()
	
	app := fiber.New()
	app.Use(cors.New(cors.Config{
    AllowOrigins:     config.GetEnvStr("CLIENT_ORIGIN"),
		AllowHeaders:     "Origin, Content-Type, Accept",
    AllowCredentials: true,
	}))
	setupRoutes(app)
	app.Listen(":3000")
}

