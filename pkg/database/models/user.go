package model

import (
	"fmt"
	"time"
	"gorm.io/gorm"
	"github.com/google/uuid"
)

type User struct {
	gorm.Model
	ID        		string 		`gorm:"type:uuid;primary_key;"`
	Firstname 		string		`json:"firstname"  			gorm:"text;not null;default:null"`
	Lastname 			string		`json:"lastname"  			gorm:"text;default:null"`
	Ridename			string		`json:"ridename"				gorm:"text;default:null"`
	IsDisabled 		bool 			`json:"is_disabled"	  	gorm:"default:false"`
	State 				string 		`json:"state"  					gorm:"text;default:null"`
	Country	  		string 		`json:"country"		  		gorm:"text;default:null"`
	Email		  		string 		`json:"email,omitempty" mapstructure:"email" csv:"Email,required" gorm:"unique;not null"`
	PasswordHash 	string		`json:"password"			  binding:"required"`
	Thumbnail			string 		`json:"photo"`
	Session 			Session
	Bikes					[]Bike
}

type LoginUserRequest struct {
  Email    string `validate:"required,min=5,max=20"` 
  Password string `validate:"required,min=4,max=36"` 
}

type ChangePasswordReqest struct {
  Password string `validate:"required,min=4,max=36"` 
}

type UserRequest struct {
  Email    			string 		`validate:"required,min=5,max=20"`
	Firstname 		string		`validate:"required,min=2,max=20"`
	Lastname 			string		`validate:"required,min=2,max=20"`
	Ridename			string		`validate:"required,min=1,max=20"`
	State 				string 		`validate:"required,min=2,max=20"`
	Country	 			string		`validate:"required,min=2,max=20"`
}

type UserResponse struct {
	ID        	string	  `json:"id,omitempty"`
	Firstname  	string    `json:"name,omitempty"`
	Ridename		string		`json:"ridename"`
	Email     	string    `json:"email,omitempty"`
	State 			string 		`json:"state,omitempty"`
	Country	  	string 		`json:"country,omitempty"`
	Thumbnail   string    `json:"photo,omitempty"`
	CreatedAt 	time.Time `json:"created_at"`
	UpdatedAt	  time.Time `json:"updated_at"`
}

func FilterUserRecord(user *User) UserResponse {
	return UserResponse{
		ID:        		user.ID,
		Firstname:  	user.Firstname,
		Ridename:			user.Ridename,
		Email:     		user.Email,
		State:				user.State,
		Country:			user.Country,
		Thumbnail: 		user.Thumbnail,
		CreatedAt: 		user.CreatedAt,
		UpdatedAt: 		user.UpdatedAt,
	}
}

func FilterGetUsers(users []User) []UserResponse {
	var filteredUsers []UserResponse
		for _, user := range users {
			currentUser := UserResponse{
			ID:        		user.ID,
			Firstname:  	user.Firstname,
			Ridename:			user.Ridename,
			Email:     		user.Email,
			State:				user.State,
			Country:			user.Country,
			Thumbnail: 		user.Thumbnail,
			CreatedAt: 		user.CreatedAt,
			UpdatedAt: 		user.UpdatedAt,
		}
	filteredUsers = append(filteredUsers, currentUser)
	}
	return filteredUsers
}

type Geo struct {
	Longtitude   float64 `gorm:"type:decimal(10,8)"`
	Latitude     float64 `gorm:"type:decimal(11,8)"`
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	user.ID = uuid.New().String()
	return
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
