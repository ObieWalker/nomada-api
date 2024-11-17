package model

import (
	"gorm.io/gorm"
)

type Route struct {
	gorm.Model
	ID						uint   		`gorm:"primary_key" gorm:"AUTO_INCREMENT" json:"id,omitempty" mapstructure:"id" csv:"ID"`
	Name			 		string		`json:"name" gorm:"text;not null;"`
	Coords				Geo				`gorm:"embedded"`
	RideID				uint
}
