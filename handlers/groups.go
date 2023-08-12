package handlers

import (
	"github.com/gofiber/fiber/v2"
  "github.com/obiewalker/nomada-api/pkg/database"
  "github.com/obiewalker/nomada-api/handlers/crud"
  "github.com/obiewalker/nomada-api/pkg/database/models"
)

func CreateGroup(c *fiber.Ctx) ( error) { 
  req := new(model.GroupRequest)

  user := c.Locals("user").(model.UserResponse)

  if err := c.BodyParser(req); err != nil {
    return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
  }

  newGroup := &model.Group{
    Name: 		req.Name,
		UserID: 	user.ID,
  }

	group, err := crud.CreateGroup(database.Instance.Db, newGroup)

	if err != nil {
    return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
  }
  return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": fiber.Map{"group": group}})
}