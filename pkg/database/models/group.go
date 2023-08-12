package model

import (
	"gorm.io/gorm"
	"github.com/google/uuid"
)

type Group struct {
	gorm.Model
	ID          			string 		`gorm:"type:uuid;primary_key;"`
	Name							string		`json:"group_name"  			gorm:"text;not null;default:null"`
	UserID						string		`gorm:"type:uuid;column:user_foreign_key;not null;"`
}

type GroupRequest struct {
	ID    	 string	`json:"id,omitempty"`
  Name     string `validate:"min=2,max=20"`
}

func (group *Group) BeforeCreate(tx *gorm.DB) (err error) {
	group.ID = uuid.New().String()
	return
}

func MigrateGroups(db *gorm.DB) error {
	err := db.AutoMigrate(&Group{})
	return err
}