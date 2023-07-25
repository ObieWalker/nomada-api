package model

import (
	"gorm.io/gorm"
)

type FaultType string

const (
    Accident 	FaultType = "accident"
    Fault 		FaultType = "fault"
)
type Stop struct {
	gorm.Model
	ID							uint   `gorm:"primary_key" gorm:"AUTO_INCREMENT" json:"id,omitempty" mapstructure:"id" csv:"ID"`
	FaultType				FaultType
	Description			string
	Location				Geo
	LocationDesc		string
	UserID					string
}

