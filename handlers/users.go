package handlers

import (
	"fmt"
  "encoding/json"
  "time"
  "strings"
  "strconv"
	"github.com/gofiber/fiber/v2"
  "github.com/golang-jwt/jwt/v5"
  "github.com/obiewalker/nomada-api/handlers/auth/signup"
  "github.com/obiewalker/nomada-api/pkg/database"
  "github.com/obiewalker/nomada-api/utils"
	"github.com/obiewalker/nomada-api/handlers/auth/login"
  "github.com/obiewalker/nomada-api/handlers/crud"
  "github.com/obiewalker/nomada-api/pkg/database/models"
  "github.com/obiewalker/nomada-api/config"
)

func Home(c *fiber.Ctx) error {
  return c.Status(fiber.StatusOK).JSON(fiber.Map{
    "status":  "success",
    "message": "Welcome!!!",
  })
}

func Signup(c *fiber.Ctx) (error) { 
  r := new(model.LoginUserRequest)

  if err := c.BodyParser(r); err != nil {
      return err
  }

  user, err := signup.Signup(database.Instance.Db, &signup.Request{
    Email: strings.ToLower(r.Email),
    Password: r.Password,
  })
	if err != nil {
    switch err.(type){
    case *model.EmailDuplicateError:
      return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
    default:
      return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": err.Error()})
    }
  }
  return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": fiber.Map{"user": user}})
}

func Login(c *fiber.Ctx) error {
  r := new(model.LoginUserRequest)
  if err := c.BodyParser(r); err != nil {
    return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
  }

  res, err := login.Login(database.Instance.Db, &login.Request{
    Email: r.Email, 
    Password: r.Password,
  })
  if err != nil {
    switch err.(type) {
    case *crud.EmailNotExistsError:
      return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
    case *login.PasswordMismatchError:
      return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
    default:
      return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": err.Error()})
    }
  }

  tokenByte := jwt.New(jwt.SigningMethodHS256)

  now := time.Now()
	claims := tokenByte.Claims.(jwt.MapClaims)
  jwtExpires, err := config.GetEnvTime("JWT_EXPIRED_IN")
	claims["sub"] = res.User.ID
  claims["exp"] = json.Number(strconv.FormatInt(time.Now().Add(jwtExpires).Unix(), 10))
	claims["iat"] = now.Unix()
	claims["nbf"] = now.Unix()

  tokenString, err := tokenByte.SignedString([]byte(config.GetEnvStr("JWT_SECRET")))

	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "fail", "message": fmt.Sprintf("generating JWT Token failed: %v", err)})
	}

  jwtMaxAge, err := config.GetEnvInt("JWT_MAXAGE")
	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    tokenString,
		Path:     "/",
		MaxAge:   jwtMaxAge * 60,
		Secure:   false,
		HTTPOnly: true,
		Domain:   "localhost",
	})

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "token": tokenString})
}

func GetUser(c *fiber.Ctx) (error) {
  res, err := crud.FindById(database.Instance.Db, c.Params("id"))
  if err != nil {
    switch err.(type){
    case *crud.NotExistsError:
      return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
    default:
      return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": err.Error()})
    }
  }
  return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": fiber.Map{"user": model.FilterUserRecord(res)}})
}

func GetUsers(c *fiber.Ctx) (error) {
  db := database.Instance.Db
  res, err := crud.FindUsers(c, db)
  if err != nil {
    switch err.(type){
    case *crud.NotExistsError:
      return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
    default:
      return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": err.Error()})
    }
  }
  return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": fiber.Map{"user": model.FilterGetUsers(res)}})
}

func DisableUser(c *fiber.Ctx) (error) {
  user := c.Locals("user").(model.UserResponse)
  res, err := crud.UpdateUser(database.Instance.Db, user.ID, "IsDisabled", true)
  if err != nil {
    switch err.(type){
    case *crud.NotExistsError:
      return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
    default:
      return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": err.Error()})
    }
  }
  return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "This user has been disabled.", "data": fiber.Map{"user": res}})
}

func ModifyUser(c *fiber.Ctx) (error) {
  user := c.Locals("user").(model.UserResponse)

  updateUser := new(model.UserRequest)

  if err := c.BodyParser(updateUser); err != nil {
    return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
  }

  res, err := crud.UpdateUserData(database.Instance.Db, user.ID, updateUser)
  if err != nil {
    return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
  }
  return c.Status(fiber.StatusOK).JSON(fiber.Map{
    "status": "success", 
    "data": fiber.Map{"user": model.FilterUserRecord(&res)}})
}

func Logout(c *fiber.Ctx) (error) {
	expired := time.Now().Add(-time.Hour * 24)
	c.Cookie(&fiber.Cookie{
		Name:    "token",
		Value:   "",
		Expires: expired,
	})
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success"})
}

// func RefreshToken(c *fiber.Ctx) (error) {

//   return c.JSON(fiber.Map{
//     "message": "Logged out successfully.",
//   })
// }

func GetCurrentUser(c *fiber.Ctx) (error) {
	user := c.Locals("user").(model.UserResponse)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": fiber.Map{"user": user}})
}

func ChangePassword(c *fiber.Ctx) (error) {
	user := c.Locals("user").(model.UserResponse)

  updatePassword := new(model.ChangePasswordReqest)

  if err := c.BodyParser(updatePassword); err != nil {
    return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
  }

  hashPassword, err := utils.HashPassword(updatePassword.Password)

  if err != nil {
    return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}
  _, err = crud.UpdateUser(database.Instance.Db, user.ID, "PasswordHash", hashPassword)
  if err != nil {
    return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": "Your password has been updated."})
}
