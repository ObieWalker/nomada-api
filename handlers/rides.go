package handlers

import (
	"strconv"
	"github.com/gofiber/fiber/v2"
  "github.com/obiewalker/nomada-api/handlers/crud"
  "github.com/obiewalker/nomada-api/pkg/database"
  "github.com/obiewalker/nomada-api/pkg/database/models"
)

func CreateRide(c *fiber.Ctx) ( error) { 
  r := new(model.RideRequest)

  user := c.Locals("user").(model.UserResponse)

  if err := c.BodyParser(r); err != nil {
    return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
  }

  newRide := &model.Ride{
    Name: r.Name,
    GroupID: r.GroupID,    
    CreatorID: user.ID,
    RoadCaptainID: r.RoadCaptainID,
    SweeperId: r.SweeperId,
		StartPos: r.StartPos,
		Destination: r.Destination,
  }

	ride, err := crud.CreateRide(database.Instance.Db, newRide)

	if err != nil {
    return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
  }
  return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": fiber.Map{"ride": ride}})
}

func UpdateRide(c *fiber.Ctx) ( error) { 
	id := c.Params("id")
  updateRide := new(model.RideRequest)

  if err := c.BodyParser(updateRide); err != nil {
    return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
  }

	ride, err := crud.UpdateRide(database.Instance.Db, id, updateRide)

	if err != nil {
    return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
  }
  return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": fiber.Map{"ride": ride}})
}

func GetRide(c *fiber.Ctx) (error) { 
  id := c.Params("id")
	i64, err := strconv.ParseUint(id, 10, 32)
	uID := uint(i64)

  ride, err := crud.FindRide(database.Instance.Db, uID)
  if err != nil {
    return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": err.Error()})
  }

  return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": fiber.Map{"ride": ride}})
}

func DeleteRide(c *fiber.Ctx) (error) { 
  id := c.Params("id")

  result := crud.DeleteRide(database.Instance.Db, id)

	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "This ride does not exist"})
	} else if result.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": result.Error})
	}
  return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "This ride has been deleted."})
}

func StartRide(c *fiber.Ctx) (error) { 
  id := c.Params("id")

  result := crud.DeleteRide(database.Instance.Db, id)

	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "This ride does not exist"})
	} else if result.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": result.Error})
	}
  return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "This ride has been deleted."})
}

func StopRide(c *fiber.Ctx) (error) { 
  id := c.Params("id")

  result := crud.DeleteRide(database.Instance.Db, id)

	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "This ride does not exist"})
	} else if result.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": result.Error})
	}
  return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "This ride has been deleted."})
}
