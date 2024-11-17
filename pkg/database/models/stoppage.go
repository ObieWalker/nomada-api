package model

import "gorm.io/gorm"

type StoppageType string

const (
    Accident 	StoppageType = "accident"
    Fault 		StoppageType = "fault"
		None 			StoppageType = "none"	
)
type Stoppage struct {
	gorm.Model
	ID							uint   				`gorm:"primary_key" gorm:"AUTO_INCREMENT" json:"id,omitempty" mapstructure:"id" csv:"ID"`
	StoppageType		StoppageType
	Description			string				`json:"description" gorm:"text;default:null"`
	Location				Geo						`gorm:"embedded"`
	LocationDesc		string				`json:"location_desc" gorm:"text;default:null"`
	UserID					string
	User						User
	Resolved				bool					`json:"resolved" gorm:"default:false"`
}

type StoppageRequest struct {
	ID        	 		string	`json:"id,omitempty"`
  StoppageType    StoppageType 
  Description  		string `validate:"requird,min=2,max=40"` 
  Location 	 			Geo
  LocationDesc 	  string `validate:"min=2,max=40"`
}

