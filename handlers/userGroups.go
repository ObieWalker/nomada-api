package handlers

import (
	"github.com/gofiber/fiber/v2"
  "github.com/obiewalker/nomada-api/pkg/database"
  "github.com/obiewalker/nomada-api/handlers/crud"
  "github.com/obiewalker/nomada-api/pkg/database/models"
)

func GetOwnGroups(c *fiber.Ctx) (error) { 
  user := c.Locals("user").(model.UserResponse)

  usersGroups, err := crud.FindUsersGroups(database.Instance.Db, user.ID)
  if err != nil {
    return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": err.Error()})
  }

  // return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": fiber.Map{"userGroups": model.FilterUsersGroup(&usersGroups)}})
  return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": fiber.Map{"userGroups": &usersGroups}})
}

func GetUsersGroups(c *fiber.Ctx) (error) { 
  userId := c.Params("userId")
  usersGroups, err := crud.FindUsersGroups(database.Instance.Db, userId)
  if err != nil {
    return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": err.Error()})
  }

  // return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": fiber.Map{"userGroups": model.FilterUsersGroup(&usersGroups)}})
  return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": fiber.Map{"userGroups": &usersGroups}})
}

func GetGroupUsers(c *fiber.Ctx) ( error) {
  // user := c.Locals("user").(model.UserResponse)
	groupId := c.Params("groupId")
	
	groupUser, err := crud.GetGroupUsers(database.Instance.Db, groupId)

	if err != nil {
    return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
  }
  return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": fiber.Map{"groupUser": groupUser}})
}

