package crud

import (
	"gorm.io/gorm"
  "github.com/gofiber/fiber/v2"
  "github.com/obiewalker/nomada-api/pkg/database/models"
  "github.com/obiewalker/nomada-api/pkg/database"
)

type EmailNotExistsError struct{}
type EmailAlreadyExistsError struct{}
type NotExistsError struct{}
type GetError struct{}

func (*EmailNotExistsError) Error() string {
  return "Email does not exist or password is wrong."
}

func (*EmailAlreadyExistsError) Error() string {
  return "This Email already being used."
}

func (*NotExistsError) Error() string {
  return "This does not exist."
}

func (*GetError) Error() string {
  return "There was an error retrieving this data."
}

func FindByEmail(db *gorm.DB, email string) (*model.User, error) {
  var user model.User
  result := db.Find(&user, &model.User{Email: email})
	if result.RowsAffected == 0 {
    return nil, &EmailNotExistsError{}
  }
  return &user, nil
}

func FindById(db *gorm.DB, ID string) (*model.User, error) {
  var user = model.User{ID: ID}
  result := db.First(&user)
	if result.RowsAffected != 1 {
    return nil, &NotExistsError{}
  }
  return &user, nil
}

func FindUsers(c *fiber.Ctx, db *gorm.DB) ([]model.User, error) {
  var users []model.User
  result := db.Scopes(database.Paginate(c)).Find(&users)

	if result.RowsAffected < 1 {
    return nil, &GetError{}
  }
  return users, nil
}

func FindBike(db *gorm.DB, Id string) (model.Bike, error) {
  var bike = model.Bike{ID: Id}
  result := db.First(&bike)
	if result.RowsAffected != 1 {
    return bike, &NotExistsError{}
  }
  return bike, nil
}

func PreloadUsers(db *gorm.DB, Id string, preload string) (model.User, error) {
  var user = model.User{ID: Id}
  db.Preload(preload, "not_in_use = ?", false).Find(&user)
  return user, nil
}