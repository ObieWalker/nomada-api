package model

import (
	"gorm.io/gorm"
)

type Ride struct {
	gorm.Model
	ID							uint   `gorm:"primary_key" gorm:"AUTO_INCREMENT" json:"id,omitempty" mapstructure:"id" csv:"ID"`
	Name	 					string	`json:"name"  				gorm:"text;not null;default:null`
	Lastname 				string	`json:"lastname"  			gorm:"text;default:null`
	UserGroupID			int
	CreatedBy				string
	RoadCaptainID		string
	SweeperId				string
	RidersCount			int
	Done						bool
	StartPos				Geo
	Destination			Geo
	Stops						[]Geo
	Route						[]Geo
}