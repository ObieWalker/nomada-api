package crud

import (
	"github.com/gofiber/fiber/v2"
	"github.com/obiewalker/nomada-api/pkg/database"
	model "github.com/obiewalker/nomada-api/pkg/database/models"
	"gorm.io/gorm"
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
	var user model.User
	result := db.First(&user, "id = ?", ID)
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

func FindRide(db *gorm.DB, Id uint) (model.Ride, error) {
	var ride = model.Ride{ID: Id}
	result := db.First(&ride)
	if result.RowsAffected != 1 {
		return ride, &NotExistsError{}
	}
	return ride, nil
}

func FindStoppage(db *gorm.DB, Id uint) (model.Stoppage, error) {
	var stoppage = model.Stoppage{ID: Id}
	result := db.First(&stoppage)
	if result.RowsAffected != 1 {
		return stoppage, &NotExistsError{}
	}
	return stoppage, nil
}

func FindUserStoppage(db *gorm.DB, userId string) (model.Stoppage, error) {
	var stoppage model.Stoppage
	result := db.Where("userID = ? AND resolved =?", userId, false).Find(&stoppage)
	if result.RowsAffected != 1 {
		return stoppage, &NotExistsError{}
	}
	return stoppage, nil
}

func FindGroupStoppage(db *gorm.DB, groupId string) (model.UserGroup, error) {
	var userGroup model.UserGroup
	result := db.Where("groupID = ? AND stopppage =?", groupId, true).Preload("Users").Find(&userGroup)
	if result.RowsAffected != 1 {
		return userGroup, &NotExistsError{}
	}
	return userGroup, nil
}

func FindGroup(db *gorm.DB, Id string) (model.Group, error) {
	var group = model.Group{ID: Id}
	result := db.First(&group)
	if result.RowsAffected != 1 {
		return group, &NotExistsError{}
	}
	return group, nil
}

func PreloadUsersBikes(db *gorm.DB, Id string, preload string) (model.User, error) {
	var user = model.User{ID: Id}
	db.Preload(preload, "not_in_use = ?", false).Find(&user)
	return user, nil
}

func FindUsersGroups(db *gorm.DB, userId string) ([]model.Group, error) {
	var userGroups []model.Group
	var user = model.User{ID: userId}

	db.Model(&user).Association("Groups").Find(&userGroups)

	return userGroups, nil
}

func GetGroupUsers(db *gorm.DB, groupId string) (model.Group, error) {
	var group = model.Group{ID: groupId}
	err := db.Model(&model.Group{}).Preload("Members").Find(&group).Error
	return group, err
}
