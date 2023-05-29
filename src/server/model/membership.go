package model

type Membership struct {
	BaseModel
	MembershipName  string `json:"membership_name"`
	Description     string `json:"description"`
	Price           uint32 `json:"price"`
	DurationInMonth uint8  `json:"duration_in_month"`
}

var Memberships = []Membership{}
