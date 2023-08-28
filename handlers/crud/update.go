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

func UpdateGroupData(db *gorm.DB,  id string, groupData *model.GroupRequest) (model.Group, error) {
	var group model.Group

	db.Model(&group).Where("id = ?", id).Updates(groupData)
	return group, nil
}

func UpdateBike(db *gorm.DB,  id string, updatedBike *model.BikeRequest) (model.Bike, error) {
	var bike model.Bike

	db.Model(&bike).Where("id = ?", id).Updates(updatedBike)
	return bike, nil
}

func UpdateRide(db *gorm.DB,  id string, updatedRide *model.RideRequest) (model.Ride, error) {
	var ride model.Ride

	db.Model(&ride).Where("id = ?", id).Updates(updatedRide)
	return ride, nil
}

func UpdateStoppage(db *gorm.DB,  id uint, updatedStoppage *model.StoppageRequest) (model.Stoppage, error) {
	var stoppage model.Stoppage

	db.Model(&stoppage).Where("id = ?", id).Updates(updatedStoppage)
	return stoppage, nil
}

func ResolveStoppage(db *gorm.DB,  stoppageID uint, userId string) (model.Stoppage, error) {
	var stoppage = model.Stoppage{ID: stoppageID}

	db.Model(&stoppage).Update("resolved", true)
	db.Model(&model.UserGroup{}).Where("userID = ? AND stoppage = ?", userId, true).Update("stoppage", false)
	return stoppage, nil
}

func UpdateBikeStatus(db *gorm.DB, ID string) (model.Bike, error) {
	var bike = model.Bike{ID: ID}
	err := db.Model(&bike).Update("NotInUse", true).Error
	if err != nil {
		return bike, err
	}
	return bike, nil
}

func AddUserToGroup(db *gorm.DB, groupId string, userId string) (model.User, model.Group, error) {
	var user = model.User{ID: userId}
	var group = model.Group{ID: groupId}

	if err := db.First(&user).Error; err != nil {
		return user,group, err
	}

	if err := db.First(&group).Error; err != nil {
		return user, group, err
	}

	user.Groups = append(user.Groups, &group)
	if err := db.Save(&user).Error; err != nil {
		return user, group, err
	}

	group.Members = append(group.Members, &user)
	if err := db.Save(&group).Error; err != nil {
		return user, group, err
	}
	return user, group, nil
}
