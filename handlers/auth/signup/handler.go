package signup

import (
	"gorm.io/gorm"
	"golang.org/x/crypto/bcrypt"
  "github.com/obiewalker/nomada-api/handlers/crud"
  "github.com/obiewalker/nomada-api/pkg/database/models"

)

type (
  Request struct {
    Email    string
    Password string
  }

  Response struct {
    Email string
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

	user, err := crud.CreateUser(db, newUser)
  if err != nil {
    return nil, err
  }
  return &Response{Email: user.Email}, err
}