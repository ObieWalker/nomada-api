package model

import (
	"gorm.io/gorm"
)

type StoppageType string

const (
    Accident 	StoppageType = "accident"
    Fault 		StoppageType = "fault"
		None 			StoppageType = "none"	
)
type Stop struct {
	gorm.Model
	ID							uint   `gorm:"primary_key" gorm:"AUTO_INCREMENT" json:"id,omitempty" mapstructure:"id" csv:"ID"`
	StoppageType		StoppageType
	Description			string
	Location				Geo
	LocationDesc		string
	UserID					string
}

