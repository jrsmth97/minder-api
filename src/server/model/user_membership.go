package model

import (
	"time"
)

type UserMembership struct {
	BaseModel
	UserId       string    `json:"user_id"`
	MembershipId string    `json:"membership_id"`
	ValidUntil   time.Time `json:"valid_until"`

	User       User       `json:"user"`
	Membership Membership `json:"membership"`
}

var UserMemberships = []UserMembership{}
