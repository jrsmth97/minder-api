package model

type UserInterest struct {
	BaseModel
	UserId     string `json:"user_id"`
	InterestId string `json:"interest_id"`

	User     User     `json:"user"`
	Interest Interest `json:"interest"`
}

var UserInterests = []UserInterest{}
