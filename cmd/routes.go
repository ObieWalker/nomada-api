package main

import (
    "github.com/gofiber/fiber/v2"
		"github.com/obiewalker/nomada-api/handlers"
		"github.com/gofiber/fiber/v2/middleware/logger"
		"github.com/obiewalker/nomada-api/middleware"
)

func setupRoutes(app *fiber.App) {
	api := app.Group("/api", logger.New())  

	api.Get("/", handlers.Home)
	api.Get("/me", middleware.DeserializeUser, handlers.GetCurrentUser)
	api.Post("/login", middleware.ValidateCredentials, handlers.Login)
	api.Post("/signup", middleware.ValidateCredentials, handlers.Signup)
	api.Get("/user/:id", handlers.GetUser)
	api.Get("/user", handlers.GetUsers)
	api.Put("/user", middleware.DeserializeUser, handlers.ModifyUser)
	api.Patch("/user", middleware.DeserializeUser, handlers.ChangePassword)
	api.Delete("/user", middleware.DeserializeUser, handlers.DisableUser)
	api.Get("/auth/logout", middleware.DeserializeUser, handlers.Logout)
	// api.Get("/auth/refresh", handlers.RefreshToken)

	api.Post("/bike", middleware.DeserializeUser, handlers.CreateBike)
	api.Put("/bike/:id/thumbnail", middleware.DeserializeUser, handlers.UpdateThumbnail)
	api.Put("/bike/:id", middleware.DeserializeUser, handlers.UpdateBike)
	api.Get("/bike/:id", handlers.GetBike)
	api.Get("/user/:userId/bike", handlers.GetUsersBike)
	api.Get("/user/bike", handlers.GetOwnBike)
	api.Delete("/bike/:id", middleware.DeserializeUser, handlers.DisableBike)
	// api.Delete("/bike/:id", middleware.DeserializeUser, handlers.DeleteBike)

	api.Post("/group", middleware.DeserializeUser, handlers.CreateGroup)
	api.Put("/group/:id", middleware.DeserializeUser, handlers.UpdateGroupName)
	api.Get("/group/:id", handlers.GetGroup)
	api.Get("/users/:userId/groups", middleware.DeserializeUser, handlers.GetUsersGroups)
	api.Get("/user/groups", middleware.DeserializeUser, handlers.GetOwnGroups)
	api.Delete("/group/:groupId", middleware.DeserializeUser, handlers.DeleteGroup)

	api.Post("/group/:groupId/user", middleware.DeserializeUser, handlers.JoinGroup)
	api.Get("/group/:groupId/users", handlers.GetGroupUsers)
	api.Put("/group/:groupId/user/:userId", middleware.DeserializeUser, handlers.AddUserToGroup)
	api.Delete("/group/:groupId/user/:userId",  middleware.DeserializeUser, handlers.DeleteUserFromGroup)
	api.Delete("/group/:groupId/user",  middleware.DeserializeUser, handlers.LeaveGroup)

	api.Post("/ride",  middleware.DeserializeUser, handlers.CreateRide)
	api.Put("/ride/:id/start",  middleware.DeserializeUser, handlers.StartRide)
	api.Put("/ride/:id/stop",  middleware.DeserializeUser, handlers.StopRide)
	api.Put("/ride/:id",  middleware.DeserializeUser, handlers.UpdateRide)
	api.Get("/ride/:id",  middleware.DeserializeUser, handlers.GetRide)
	api.Delete("/ride/:id",  middleware.DeserializeUser, handlers.DeleteRide)

	api.Post("/stoppage", handlers.CreateStoppage)
	api.Put("/stoppage", handlers.UpdateStoppage)
	api.Get("/stoppage/:id", handlers.GetStoppage)
	api.Get("/user/:userId/stoppage", handlers.GetUserStoppage)
	api.Get("/group/:groupId/stoppage", handlers.GetGroupStoppage)
	api.Delete("/stoppage/:id", handlers.ResolveStoppage)

	// api.Post("/route", handlers.SaveRoute)
	// api.Put("/route", handlers.UpdateRoute)
	// api.Get("/route/:id", handlers.GetRoute)
	// api.Delete("/route", handlers.DeleteRoute)
}
