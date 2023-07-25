package model

import (
	"gorm.io/gorm"
)

type UserGroup struct {
	gorm.Model
	ID   					int  		`gorm:"primaryKey"`
	UserID 				string
	GroupID				uint
	HasStopped		bool		`gorm:"default:false`
}