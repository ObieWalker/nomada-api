package handlers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
  "github.com/obiewalker/nomada-api/handlers/user/signup"
  "github.com/obiewalker/nomada-api/pkg/database"
	"github.com/obiewalker/nomada-api/handlers/user/login"
	"github.com/obiewalker/nomada-api/handlers/user"
  "github.com/obiewalker/nomada-api/pkg/database/models"
)

type Request struct {
  Email    string `validate:"required,min=5,max=20"` 
  Password string `validate:"required,min=4,max=36"` 
}

type Response struct {
  Id string
}

func Home(c *fiber.Ctx) error {
  return c.SendString("Home!!!")
}

func Signup(c *fiber.Ctx) ( error) { 
  r := new(Request)

  if err := c.BodyParser(r); err != nil {
      return err
  }

  res, err := signup.Signup(database.Instance.Db, &signup.Request{
    Email: r.Email,
    Password: r.Password,
  })
	if err != nil {
    switch err.(type){
    case *model.EmailDuplicateError:
			msg := fmt.Sprintf("Bad Request: , %s ", err.Error())
      return c.SendString(msg)
    default:
			msg := fmt.Sprintf("Internal Server Error: , %s ", err.Error())
      return c.SendString(msg)
    }
  }
	msg := fmt.Sprintf("Created: , %s", res)
  return c.SendString(msg)
}

func Login(c *fiber.Ctx) ( error) {
  r := new(Request)
  if err := c.BodyParser(r); err != nil {
    return err
  }

  res, err := login.Login(database.Instance.Db, &login.Request{
    Email: r.Email, 
    Password: r.Password,
  })
  if err != nil {
    switch err.(type) {
    case *user.EmailNotExistsError:
      msg := fmt.Sprintf("Bad Request: , %s ", err.Error())
      return c.SendString(msg)
    case *login.PasswordMismatchError:
      msg := fmt.Sprintf("Bad Request: , %s ", err.Error())
      return c.SendString(msg)
    default:
      msg := fmt.Sprintf("Internal Server Error: , %s ", err.Error())
      return c.SendString(msg)
    }
  }
  msg:= fmt.Sprintf("Ok: User '%s' logged in", res.User)
  return c.SendString(msg)
}
