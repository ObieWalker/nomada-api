package handlers

import (
  "strconv"
	"github.com/gofiber/fiber/v2"
  "github.com/obiewalker/nomada-api/handlers/crud"
  "github.com/obiewalker/nomada-api/pkg/database"
  "github.com/obiewalker/nomada-api/pkg/database/models"
)

func CreateStoppage(c *fiber.Ctx) ( error) { 
  req := new(model.StoppageRequest)

  user := c.Locals("user").(model.UserResponse)

  if err := c.BodyParser(req); err != nil {
    return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
  }

  newStoppage := &model.Stoppage{
    StoppageType: req.StoppageType,
    Description:  req.Description,    
    Location:     req.Location,
    LocationDesc: req.LocationDesc,
    UserID:       user.ID,
  }

  // update userGroup column stoppage to true as a transaction
	stoppage, err := crud.CreateStoppage(database.Instance.Db, newStoppage)

	if err != nil {
    return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
  }
  return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": fiber.Map{"stoppage": stoppage}})
}

func UpdateStoppage(c *fiber.Ctx) (error) { 
  id := c.Params("id")
  updateStoppage := new(model.StoppageRequest)

  if err := c.BodyParser(updateStoppage); err != nil {
    return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
  }

  i64, _ := strconv.ParseUint(id, 10, 32)
	uID := uint(i64)
	stoppage, err := crud.UpdateStoppage(database.Instance.Db, uID, updateStoppage)

	if err != nil {
    return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
  }
  return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": fiber.Map{"stoppage": stoppage}})
}

func GetStoppage(c *fiber.Ctx) (error) { 
  id := c.Params("id")
  i64, _ := strconv.ParseUint(id, 10, 32)
	uID := uint(i64)

  stoppage, err := crud.FindStoppage(database.Instance.Db, uID)
  if err != nil {
    return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": err.Error()})
  }

  return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": fiber.Map{"stoppage": stoppage}})
}

func GetUserStoppage(c *fiber.Ctx) (error) { 
  userId := c.Params("userId")

  stoppage, err := crud.FindUserStoppage(database.Instance.Db, userId)
  if err != nil {
    return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": err.Error()})
  }

  return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": fiber.Map{"stoppage": stoppage}})
}

func GetGroupStoppage(c *fiber.Ctx) (error) { 
  groupId := c.Params("groupId")

  stoppage, err := crud.FindGroupStoppage(database.Instance.Db, groupId)
  if err != nil {
    return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": err.Error()})
  }

  return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": fiber.Map{"stoppage": stoppage}})
}

func ResolveStoppage(c *fiber.Ctx) (error) { 
  id := c.Params("id")
  i64, _ := strconv.ParseUint(id, 10, 32)
	uID := uint(i64)
  user := c.Locals("user").(model.UserResponse)

  stoppage, err := crud.ResolveStoppage(database.Instance.Db, uID, user.ID)
  if err != nil {
    return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": err.Error()})
  }

  return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": fiber.Map{"stoppage": stoppage}})
}