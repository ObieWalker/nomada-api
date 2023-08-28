package crud

import (
	"gorm.io/gorm"
	"github.com/obiewalker/nomada-api/pkg/database/models"
) 

func DeleteBike(db *gorm.DB,  id string) (*gorm.DB) {
	result := db.Delete(&model.Bike{}, "id = ?", id)
	return result
}

func DeleteGroup(db *gorm.DB,  id string, userID string) (*gorm.DB) {
	result := db.Delete(&model.Group{}, "id = ?", id)
	return result
}

func DeleteRide(db *gorm.DB, id string) (*gorm.DB) {
	result := db.Delete(&model.Ride{},  "id = ?", id)
	return result
}