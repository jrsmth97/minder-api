package param

import (
	"minder/src/server/model"
)

type MembershipCreate struct {
	MembershipName  string   `validate:"required" json:"membership_name"`
	Price           uint32   `validate:"required" json:"price"`
	DurationInMonth uint8    `validate:"required" json:"duration_in_month"`
	Description     string   `json:"description"`
	Privileges      []string `validate:"required" json:"privileges"`
}

type MembershipUpdate struct {
	MembershipCreate
}

func (p *MembershipCreate) ParseToModel() *model.Membership {
	return &model.Membership{
		MembershipName:  p.MembershipName,
		Price:           p.Price,
		DurationInMonth: p.DurationInMonth,
		Description:     p.Description,
	}
}

func (p *MembershipCreate) ParseToMembershipPrivilegeModel(membershipId string, privilegeId string) *model.MembershipPrivilege {
	return &model.MembershipPrivilege{
		MembershipId: membershipId,
		PrivilegeId:  privilegeId,
	}
}
