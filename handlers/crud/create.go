package crud

import (
	"gorm.io/gorm"
	"github.com/obiewalker/nomada-api/pkg/postgres"
	"github.com/obiewalker/nomada-api/pkg/database/models"
) 

func CreateUser(db *gorm.DB, user *model.User) (*model.User, error) {
	err := db.Create(user).Error
	if err != nil {
		if postgres.IsUniqueConstraintError(err, model.UniqueConstraintEmail) {
			return user, &model.EmailDuplicateError{Email: user.Email}
		}
		return user, err
	}
	return user, nil
}

func CreateBike(db *gorm.DB, bike *model.Bike) (interface{}, error) {
	err := db.Create(bike).Error
	if err != nil {
		return "", err
	}
	return bike, nil
}

func CreateGroup(db *gorm.DB, group *model.Group) (*model.Group, error) {
	err := db.Create(group).Error
	if err != nil {
		return group, err
	}
	return group, nil
}