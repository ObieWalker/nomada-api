package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/obiewalker/nomada-api/handlers/crud"
	"github.com/obiewalker/nomada-api/handlers/transactions"
	"github.com/obiewalker/nomada-api/pkg/database"
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
    OwnerID:  user.ID,
  }

  group, err := transaction.CreateGroupTransaction(database.Instance.Db, newGroup)
	if err != nil {
    return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
  }

	if err != nil {
    return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
  }

  // return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": fiber.Map{"group": model.FilterGroup(group)}})
  return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": fiber.Map{"group": group}})
}

func UpdateGroupName(c *fiber.Ctx) ( error) {
  id := c.Params("id")
  req := new(model.GroupRequest)

  if err := c.BodyParser(req); err != nil {
    return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
  }

  _, err := crud.UpdateGroupData(database.Instance.Db, id, req)
  if err != nil {
    return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
  }

  return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "This group name has been updated." })
}

func AddUserToGroup(c *fiber.Ctx) (error) { 
  groupId := c.Params("groupId")
  userId := c.Params("userId")

  user, group, err := transaction.AddUserTransaction(database.Instance.Db, groupId, userId)
  if err != nil {
    return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": err.Error()})
  }

  return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": fiber.Map{"userId": user.ID, "groupId": group.ID}})
}

func JoinGroup(c *fiber.Ctx) (error) { 
  groupId := c.Params("groupId")
  user := c.Locals("user").(model.UserResponse)

  fmt.Println("---------user---------------")
  fmt.Print(user.ID)  
  fmt.Println("---------group id---------------")
  fmt.Print(groupId)
  userData, groupData, err := transaction.AddUserTransaction(database.Instance.Db, groupId, user.ID)
  if err != nil {
    return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": err.Error()})
  }

  fmt.Println("-----------------userdata id")
  fmt.Println(userData.ID)
  fmt.Println("-----------------groupdata id")
  fmt.Println(groupData.ID)
  return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": fiber.Map{"userId": userData.ID, "groupId": groupData.ID}})
}

func GetGroup(c *fiber.Ctx) (error) { 
  id := c.Params("id")

  bike, err := crud.FindGroup(database.Instance.Db, id)
  if err != nil {
    return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": err.Error()})
  }
  return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": fiber.Map{"bike": bike}})
}

func DeleteGroup(c *fiber.Ctx) (error) { 
  groupId := c.Params("groupId")
  user := c.Locals("user").(model.UserResponse)

  // check that current user that wants to delete is part of the user group
  // modify this feature when user group is added
  result := crud.DeleteGroup(database.Instance.Db, groupId, user.ID)

	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "This Group does not exist"})
	} else if result.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": result.Error})
	}
  return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "This group has been deleted."})
}

func DeleteUserFromGroup(c *fiber.Ctx) (error) { 
  groupId := c.Params("groupId")
  userId := c.Params("userId")
  currentUser := c.Locals("user").(model.UserResponse)

  err := transaction.RemoveFromGroup(database.Instance.Db, groupId, currentUser.ID, userId)

  if err != nil {
    return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": err.Error()})
  }
  return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "User has been removed."})
}

func LeaveGroup(c *fiber.Ctx) (error) { 
  groupId := c.Params("groupId")
  currentUser := c.Locals("user").(model.UserResponse)

  err := transaction.RemoveFromGroup(database.Instance.Db, groupId, currentUser.ID, currentUser.ID)

  if err != nil {
    return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": err.Error()})
  }
  return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "User has been removed."})
}
