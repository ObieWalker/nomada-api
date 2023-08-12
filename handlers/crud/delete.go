package crud

import (
	"gorm.io/gorm"
	"github.com/obiewalker/nomada-api/pkg/database/models"
) 

func DeleteBike(db *gorm.DB,  id string) (*gorm.DB) {

	result := db.Delete(&model.Bike{}, "id = ?", id)

	return result
}