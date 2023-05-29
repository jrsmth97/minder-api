package view

import (
	"minder/src/server/model"
)

type MembershipCreateResponse struct {
	Id              string            `json:"id"`
	MembershipName  string            `json:"user_id"`
	Description     string            `json:"description"`
	Price           uint32            `json:"price"`
	DurationInMonth uint8             `json:"duration_in_month"`
	Privileges      []model.Privilege `json:"privileges"`
}

func NewMembershipCreateResponse(membership *model.Membership, privileges *[]model.Privilege) *MembershipCreateResponse {
	return &MembershipCreateResponse{
		Id:              membership.ID.String(),
		MembershipName:  membership.MembershipName,
		Description:     membership.Description,
		Price:           membership.Price,
		DurationInMonth: membership.DurationInMonth,
		Privileges:      *privileges,
	}
}

type MembershipUpdateResponse struct {
	MembershipCreateResponse
}

func NewMembershipUpdateResponse(membership *model.Membership, privileges *[]model.Privilege) *MembershipUpdateResponse {
	updateResponse := &MembershipUpdateResponse{}
	updateResponse.Id = membership.ID.String()
	updateResponse.MembershipName = membership.MembershipName
	updateResponse.Description = membership.Description
	updateResponse.Price = membership.Price
	updateResponse.DurationInMonth = membership.DurationInMonth
	updateResponse.Privileges = *privileges
	return updateResponse
}

type MembershipGetResponse struct {
	MembershipCreateResponse
}

func NewMembershipGetResponse(memberships *[]model.Membership, privileges *[]model.Privilege) *[]MembershipGetResponse {
	var membershipsResponse []MembershipGetResponse

	for _, membership := range *memberships {
		// var privileges []model.Privilege
		// for _, memberPrivilege := range *memberPrivileges {
		// 	if memberPrivilege.MembershipId != membership.ID.String() {
		// 		privileges = append(privileges, memberPrivilege.Privilege)
		// 	}
		// }

		response := &MembershipGetResponse{}
		response.Id = membership.ID.String()
		response.MembershipName = membership.MembershipName
		response.Description = membership.Description
		response.Price = membership.Price
		response.DurationInMonth = membership.DurationInMonth
		response.Privileges = *privileges

		membershipsResponse = append(membershipsResponse, *response)
	}

	return &membershipsResponse
}
