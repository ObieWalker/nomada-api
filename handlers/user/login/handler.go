package login

import (
	"gorm.io/gorm"
	"golang.org/x/crypto/bcrypt"
	"github.com/obiewalker/nomada-api/handlers/user"
  "github.com/obiewalker/nomada-api/pkg/database/models"
)

type Request struct {
  Email    string
  Password string
}

type Response struct {
  User *model.User
}

type PasswordMismatchError struct{}
func (e *PasswordMismatchError) Error() string {
  return "Email does not exist or password is wrong."
}

func Login(db *gorm.DB, req *Request) (*Response, error) { 

  user, err := user.FindByEmail(db, req.Email)
  if err != nil {
    return nil, err
  }

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password))
  if err != nil {
    return nil, &PasswordMismatchError{}
  }
  return &Response{User: user}, nil
}
