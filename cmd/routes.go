package main

import (
    "github.com/gofiber/fiber/v2"
		"github.com/obiewalker/nomada-api/handlers"
		"github.com/gofiber/fiber/v2/middleware/logger"
		"github.com/obiewalker/nomada-api/middleware"
)

func setupRoutes(app *fiber.App) {
	api := app.Group("/api", logger.New(), middleware.AuthReq())  

	api.Get("/", handlers.Home)
	api.Post("/login", middleware.ValidateCredentials, handlers.Login)
	api.Post("/signup", middleware.ValidateCredentials, handlers.Signup)
	// api.Get("/user/:id", handlers.GetUser)
	// api.Get("/user", handlers.GetUsers)
	// api.Delete("/user/:id", handlers.DisableUser)

	// api.Post("/bike", handlers.CreateBike)
	// api.Put("/bike/:id", handlers.UpdateBike)
	// api.Get("/bike/:id", handlers.GetBike)
	// api.Delete("/bike/:id", handlers.DisableBike)

	// api.Post("/group", handlers.CreateGroup)
	// api.Put("/group", handlers.UpdateGroup)
	// api.Get("/group/:id", handlers.GetGroup)
	// api.Delete("/group", handlers.DeleteGroup)

	// api.Post("/userergroup", handlers.CreateUsererGroup)
	// api.Put("/usergroup", handlers.UpdateUserGroup)
	// api.Get("/usergroup/:id", handlers.GetGroupUsers)
	// api.Delete("/usergroup", handlers.DeleteUserGroup)

	// api.Post("/ride", handlers.CreateRide)
	// api.Put("/ride", handlers.UpdateRide)
	// api.Get("/ride/:id", handlers.GetRide)
	// api.Delete("/ride", handlers.DeleteRide)

	// api.Post("/fault", handlers.CreateFault)
	// api.Put("/fault", handlers.UpdateFault)
	// api.Get("/fault/:id", handlers.GetFault)
	// api.Delete("/fault", handlers.SoftDeleteFault)

	// api.Post("/route", handlers.CreateRoute)
	// api.Put("/route", handlers.UpdateRoute)
	// api.Get("/route/:id", handlers.GetRoute)
	// api.Delete("/route", handlers.DeleteRoute)
}
