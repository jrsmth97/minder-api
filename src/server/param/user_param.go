package param

import (
	"minder/src/server/model"
)

type AuthRegister struct {
	Name       string `validate:"required"`
	Gender     uint8  `validate:"required"`
	Email      string `validate:"required"`
	Password   string `validate:"required"`
	BirthDate  string `validate:"required" json:"birth_date"`
	Phone      string
	LocationId string `json:"location_id"`
}

type UserUpdate struct {
	Name       string `validate:"required"`
	Gender     uint8  `validate:"required"`
	Email      string `validate:"required"`
	BirthDate  string `validate:"required" json:"birth_date"`
	Phone      string
	LocationId string               `validate:"required" json:"location_id"`
	Photos     []model.UserPhoto    `validate:"required" json:"photos"`
	Interests  []model.UserInterest `validate:"required" json:"interests"`
}

type UserLogin struct {
	Email    string `validate:"required"`
	Password string `validate:"required"`
}

type ChangePassword struct {
	Email    string `validate:"required"`
	Password string `validate:"required"`
}

func (u *UserUpdate) ParseToModelUpdate() *model.User {
	return &model.User{
		Name:       u.Name,
		Gender:     u.Gender,
		Email:      u.Email,
		BirthDate:  u.BirthDate,
		Phone:      u.Phone,
		LocationId: u.LocationId,
		Photos:     u.Photos,
		Interests:  u.Interests,
	}
}

func (u *AuthRegister) ParseToModelRegister() *model.User {
	return &model.User{
		Name:       u.Name,
		Gender:     u.Gender,
		Email:      u.Email,
		Password:   u.Password,
		BirthDate:  u.BirthDate,
		Phone:      u.Phone,
		LocationId: u.LocationId,
	}
}
