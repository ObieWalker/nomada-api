package user

import (
	"gorm.io/gorm"
  "github.com/obiewalker/nomada-api/pkg/database/models"
)

type EmailNotExistsError struct{}

func (*EmailNotExistsError) Error() string {
  return "Email does not exist or password is wrong."
}

func FindByEmail(db *gorm.DB, email string) (*model.User, error) {
  var user model.User
  result := db.Find(&user, &model.User{Email: email})
	if result.RowsAffected == 0 {
    return nil, &EmailNotExistsError{}
  }
  return &user, nil
}