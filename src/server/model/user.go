package model

type User struct {
	BaseModel
	Name       string `gorm:"size:100;not null" json:"name"`
	Gender     uint8  `json:"gender"`
	Email      string `gorm:"size:100;unique" json:"email"`
	Password   string `gorm:"not null" json:"password"`
	BirthDate  string `json:"birth_date"`
	Phone      string `gorm:"size:100" json:"phone"`
	LocationId string `gorm:"not null" json:"location_id"`

	Location  Location       `json:"location"`
	Photos    []UserPhoto    `json:"photos"`
	Interests []UserInterest `json:"interests"`
}

var Users = []User{}
