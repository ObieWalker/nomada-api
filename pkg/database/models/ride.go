package model

import (
	"gorm.io/gorm"
)

type Ride struct {
	gorm.Model
	ID							uint   	`gorm:"primary_key" gorm:"AUTO_INCREMENT" json:"id,omitempty" mapstructure:"id" csv:"ID"`
	Name	 					string	`json:"name"  				gorm:"text;not null;default:null`
	GroupID					string
	Route						Route
	CreatorID				string
	RoadCaptainID		string
	SweeperId				string
	RidersCount			int
	StartPos				Geo			`gorm:"embedded"`
	Destination			Geo			`gorm:"embedded"`
	Stops						Geo			`gorm:"embedded"`
	Started					bool
	Ended						bool
}

type RideRequest struct {
	ID        	 		string	`json:"id,omitempty"`
  Name    		 		string `validate:"min=2,max=30"` 
  GroupID 		 		string 
  CreatorID 	 		string 
  RoadCaptainID 	string  
	SweeperId 		 	string 
  StartPos 	 			Geo 
  Destination 		Geo
}

func MigrateRides(db *gorm.DB) error {
	err := db.AutoMigrate(&Ride{})
	return err
}
