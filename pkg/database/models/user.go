package model

import (
	"fmt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UUID   				string  `gorm:"primaryKey"`
	Firstname 		string	`json:"firstname"  			gorm:"text;not null;default:null`
	Lastname 			string	`json:"lastname"  			gorm:"text;default:null`
	RideName			string	`json:"ridename"				gorm:"text;default:null`
	Deleted 			bool 		`json:"deleted"		  		gorm:"default:false`
	State 				string 	`json:"state"  					gorm:"text;default:null`
	Country	  		string 	`json:"country"		  		gorm:"text;default:null`
	Email		  		string 	`json:"email,omitempty" mapstructure:"email" csv:"Email,required" gorm:"unique;not null"`
	PasswordHash 	string	`json:"password" binding:"required"`
	Session 			Session
	Bikes					[]Bike
}

type Geo struct {
	Longtitude   float64 `gorm:"type:decimal(10,8)"`
	Latitude     float64 `gorm:"type:decimal(11,8)"`
}

func MigrateUsers(db *gorm.DB) error {
	err := db.AutoMigrate(&User{})
	return err
}

const (
  UniqueConstraintEmail    = "users_email_key"
)

type EmailDuplicateError struct {
  Email string
}

func (e *EmailDuplicateError) Error() string {
  return fmt.Sprintf("Email '%s' already exists", e.Email)
}
