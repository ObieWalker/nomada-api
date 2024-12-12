package model

import (
	"gorm.io/gorm"
	"time"
	"github.com/google/uuid"
)

type Bike struct {
	gorm.Model
	ID        		string 		`gorm:"type:uuid;primary_key;"`
	BikePlate 		string		`json:"plate" gorm:"text;default:null"`
	BikeMake 			string		`json:"make" gorm:"text;default:null"`
	BikeModel	 	  string 		`json:"model" gorm:"text;default:null"`
	BikeYear 			string 		`json:"year" gorm:"text;default:null"`
	Thumbnail			string 		`json:"photo"`
	NotInUse			bool			`json:"not_in_use" gorm:"default:false"`
	UserID 				string	  `gorm:"type:uuid;column:user_foreign_key;not null;"`
}

type BikeRequest struct {
	ID        	 string	`json:"id,omitempty"`
  BikePlate    string `validate:"min=2,max=10"` 
  BikeMake 		 string `validate:"requird,min=2,max=12"` 
  BikeModel 	 string `validate:"min=4,max=20"` 
  BikeYear 		 string `validate:"min=2,max=6"`
}

type UserBikeResponse struct {
	ID        	string	  `json:"id,omitempty"`
	Firstname  	string    `json:"name,omitempty"`
	Ridename		string		`json:"ridename"`
	Email     	string    `json:"email,omitempty"`
	Thumbnail   string    `json:"photo,omitempty"`
	CreatedAt 	time.Time `json:"created_at"`
	UpdatedAt	  time.Time `json:"updated_at"`
	Bikes				[]BikeRequest
}

func FilterUserBike(user *User) UserBikeResponse {
	return UserBikeResponse{
		ID:        		user.ID,
		Firstname:  	user.Firstname,
		Ridename:			user.Ridename,
		Email:     		user.Email,
		Thumbnail: 		user.Thumbnail,
		CreatedAt: 		user.CreatedAt,
		UpdatedAt: 		user.UpdatedAt,
		Bikes:				filterBike(user.Bikes),
	}
}

func filterBike(bikes []Bike) []BikeRequest {
	var filteredBikes []BikeRequest
	for _, bike := range bikes {
		currentBike := BikeRequest{
			ID: 						bike.ID,
			BikePlate:			bike.BikePlate,
			BikeMake:				bike.BikeMake,
			BikeModel:			bike.BikeModel,
			BikeYear:				bike.BikeYear,
		}
		filteredBikes = append(filteredBikes, currentBike)
	}
	return filteredBikes
}

func (bike *Bike) BeforeCreate(tx *gorm.DB) (err error) {
	bike.ID = uuid.New().String()
	return
}

func MigrateBikes(db *gorm.DB) error {
	err := db.AutoMigrate(&Bike{})
	return err
}