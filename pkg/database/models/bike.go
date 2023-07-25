package model

import (
	"gorm.io/gorm"
)

type Bike struct {
	gorm.Model
	BikePlate 		string		`json:"plate"  		gorm:"text;default:null`
	BikeMake 			string		`json:"make"  		gorm:"text;default:null`
	BikeModel	 	  string 		`json:"model"			gorm:"text;default:null`
	BikeYear 			string 		`json:"year"		  gorm:"text;default:null`
	UserID 				string
}

func MigrateBikes(db *gorm.DB) error {
	err := db.AutoMigrate(&Bike{})
	return err
}