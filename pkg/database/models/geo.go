package model

type Geo struct {
	Longtitude   float64 `json:"lng" gorm:"lng;"`
	Latitude     float64 `json:"lat" gorm:"lat;"`
}
