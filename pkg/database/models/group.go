package model

import (
	"gorm.io/gorm"
)

type Group struct {
	gorm.Model
	ID          			uint 		`gorm:"primaryKey"`
	Name							string
	UserID						string
}
