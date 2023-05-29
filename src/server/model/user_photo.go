package model

type UserPhoto struct {
	BaseModel
	UserId string `json:"user_id"`
	Asset  string `json:"asset"`
}

var UserPhotos = []UserPhoto{}
