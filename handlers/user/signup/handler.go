package signup

import (
	"gorm.io/gorm"
	"golang.org/x/crypto/bcrypt"
	"github.com/obiewalker/nomada-api/handlers/user"
  "github.com/obiewalker/nomada-api/pkg/database/models"

)

type (
  Request struct {
    Email    string
    Password string
  }

  Response struct {
    Id string
  }
)

func Signup(db *gorm.DB, req *Request) (*Response, error) { 

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
  if err != nil {
    return nil, err
  }
  newUser := &model.User{
		Email:        req.Email,
		PasswordHash: string(passwordHash),
  }

	id, err := user.Create(db, newUser)
  if err != nil {
    return nil, err
  }
  return &Response{Id: id}, err
}