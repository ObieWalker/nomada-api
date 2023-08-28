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

func CreateBike(db *gorm.DB, bike *model.Bike) (*model.Bike, error) {
	err := db.Create(bike).Error
	if err != nil {
		return nil, err
	}
	return bike, nil
}

func CreateGroup(db *gorm.DB, group *model.Group) (*model.Group, error) {
	if err := db.Create(&group).Error; err != nil {
		return nil, err
	}
	return group, nil
}

func CreateUserGroup(db *gorm.DB, userGroup *model.UserGroup) (*model.UserGroup, error) {
	err := db.Create(&userGroup).Error
	if err != nil {
		return userGroup, err
	}
	return userGroup, nil
}

func CreateRide(db *gorm.DB, ride *model.Ride) (*model.Ride, error) {
	err := db.Create(ride).Error
	if err != nil {
		return nil, err
	}
	return ride, nil
}

func CreateStoppage(db *gorm.DB, stoppage *model.Stoppage) (*model.Stoppage, error) {
	err := db.Create(stoppage).Error
	if err != nil {
		return nil, err
	}
	return stoppage, nil
}