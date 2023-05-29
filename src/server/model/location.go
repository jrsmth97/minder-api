package model

type Location struct {
	BaseModel
	LocationName string `gorm:"not null" json:"location_name"`
}

var Locations = []Location{}
