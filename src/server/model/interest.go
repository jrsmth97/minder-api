package model

type Interest struct {
	BaseModel
	InterestName string `gorm:"not null" json:"interest_name"`
}

var Interests = []Interest{}
