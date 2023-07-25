package model

import (
	"gorm.io/gorm"
	"time"
)

type Session struct {
	gorm.Model
	UUID   						string 		 `gorm:"primaryKey"`
	SessionToken			string
	VerifiedToken			string
	SessionTokenExp		time.Time
	VerifiedTokenExp	time.Time
	UserID 						string
}