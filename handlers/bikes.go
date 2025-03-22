package handlers

import (
	"github.com/gofiber/fiber/v2"
  "github.com/obiewalker/nomada-api/handlers/crud"
  "github.com/obiewalker/nomada-api/pkg/database"
  "github.com/obiewalker/nomada-api/pkg/database/models"
)

func CreateBike(c *fiber.Ctx) ( error) { 
  r := new(model.BikeRequest)

  user := c.Locals("user").(model.UserResponse)

  if err := c.BodyParser(r); err != nil {
    return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
  }

  newBike := &model.Bike{
    BikePlate: r.BikePlate,
    BikeMake: r.BikeMake,    
    BikeModel: r.BikeModel,
    BikeYear: r.BikeYear,
    UserID: user.ID,
  }

	bike, err := crud.CreateBike(database.Instance.Db, newBike)

	if err != nil {
    return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
  }
  return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": fiber.Map{"bike": bike}})
}

func UpdateBike(c *fiber.Ctx) (error) {
  id := c.Params("id")
  updateBike := new(model.BikeRequest)

  if err := c.BodyParser(updateBike); err != nil {
    return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
  }

  res, err := crud.FindBike(database.Instance.Db, id)
  if err != nil {
    return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": err.Error()})
  }

  user := c.Locals("user").(model.UserResponse)
  if user.ID != res.UserID {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "fail", "message":  "You are  not authorized to update this bike information."})
  }

	bike, err := crud.UpdateBike(database.Instance.Db, id, updateBike)

	if err != nil {
    return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
  }
  return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": fiber.Map{"bike": bike}})
}

func GetBike(c *fiber.Ctx) (error) { 
  id := c.Params("id")

  bike, err := crud.FindBike(database.Instance.Db, id)
  if err != nil {
    return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": err.Error()})
  }

  return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": fiber.Map{"bike": bike}})
}

func GetUsersBike(c *fiber.Ctx) (error) { 
  id := c.Params("userId")
  user, err := crud.PreloadUsersBikes(database.Instance.Db, id, "Bikes")
  if err != nil {
    return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": err.Error()})
  }
  return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": fiber.Map{"user": model.FilterUserBike(&user)}})
}

func GetOwnBike(c *fiber.Ctx) (error) {
  user := c.Locals("user").(model.UserResponse)

  userData, err := crud.PreloadUsersBikes(database.Instance.Db, user.ID, "Bikes")
  if err != nil {
    return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": err.Error()})
  }
  return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": fiber.Map{"user": model.FilterUserBike(&userData)}})
}

func DisableBike(c *fiber.Ctx) (error) { 
  id := c.Params("id")

  res, err := crud.FindBike(database.Instance.Db, id)
  if err != nil {
    return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": err.Error()})
  }

  user := c.Locals("user").(model.UserResponse)
  if user.ID != res.UserID {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "fail", "message":  "You are  not authorized to update this bike status."})
  }

  bike, err := crud.UpdateBikeStatus(database.Instance.Db, id)
  if err != nil {
    return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": err.Error()})
  }
  return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "This bike has been deleted.", "bikeId": bike.ID})
}

func DeleteBike(c *fiber.Ctx) (error) { 
  id := c.Params("id")

  result := crud.DeleteBike(database.Instance.Db, id)

	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "This bike does not exist"})
	} else if result.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": result.Error})
	}
  return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "This bike has been deleted."})
}