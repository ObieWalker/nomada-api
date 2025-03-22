package signup

import (
	"github.com/obiewalker/nomada-api/handlers/crud"
	model "github.com/obiewalker/nomada-api/pkg/database/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type (
	Request struct {
		Email     string
		Password  string
		Firstname string
		Lastname  string
		Ridename  string
		State     string
		Country   string
	}
)

func Signup(db *gorm.DB, req *Request) (*model.User, error) {

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	newUser := &model.User{
		Email:        req.Email,
		PasswordHash: string(passwordHash),
		Firstname:    req.Firstname,
		Lastname:     req.Lastname,
		Ridename:     req.Ridename,
		State:        req.State,
		Country:      req.Country,
	}

	user, err := crud.CreateUser(db, newUser)
	if err != nil {
		return nil, err
	}
	return user, err
}
