package model

type MembershipPrivilege struct {
	BaseModel
	MembershipId string `json:"membership_id"`
	PrivilegeId  string `json:"privilege_id"`

	Membership Membership `json:"membership"`
	Privilege  Privilege  `json:"privilege"`
}

var MembershipPrivileges = []MembershipPrivilege{}
