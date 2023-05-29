package model

type UserSwipe struct {
	BaseModel
	UserId   string `json:"user_id"`
	Action   string `json:"action"`
	TargetId string `json:"target_id"`

	User User `json:"user"`
}

var UserSwipes = []UserSwipe{}
