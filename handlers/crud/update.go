package crud

import (
	"gorm.io/gorm"
	"github.com/obiewalker/nomada-api/pkg/postgres"
	"github.com/obiewalker/nomada-api/pkg/database/models"
) 

func UpdateUser(db *gorm.DB,  ID string, column string, value interface{}) (string, error) {
	var user model.User
	err := db.Model(&user).Where("id = ?", ID).Update(column, value).Error
	if err != nil {
		return "", err
	}
	return ID, nil
}

func UpdateUserData(db *gorm.DB,  id string, updatedUser *model.UserRequest) (model.User, error) {
	var user model.User

	err := db.Model(&user).Where("id = ?", id).Updates(updatedUser).Error
	if err != nil {
		if postgres.IsUniqueConstraintError(err, model.UniqueConstraintEmail) {
			return user, &model.EmailDuplicateError{Email: user.Email}
		}
		return user, err
	}
	return user, nil
}

func UpdateBike(db *gorm.DB,  id string, updatedBike *model.BikeRequest) (model.Bike, error) {
	var bike model.Bike

	db.Model(&bike).Where("id = ?", id).Updates(updatedBike)
	return bike, nil
}

func UpdateBikeStatus(db *gorm.DB,  ID string) (model.Bike, error) {
	var bike = model.Bike{ID: ID}
	err := db.Model(&bike).Update("NotInUse", true).Error
	if err != nil {
		return bike, err
	}
	return bike, nil
}