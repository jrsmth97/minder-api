package view

import (
	"minder/src/server/model"
)

type PrivilegeCreateResponse struct {
	Id            string `json:"id"`
	PrivilegeName string `json:"privilege_name"`
	ActionPath    string `json:"action_path"`
	ActionLimit   uint16 `json:"action_limit"`
	LimitInDay    uint16 `json:"limit_in_day"`
}

func NewPrivilegeCreateResponse(privilege *model.Privilege) *PrivilegeCreateResponse {
	return &PrivilegeCreateResponse{
		Id:            privilege.ID.String(),
		PrivilegeName: privilege.PrivilegeName,
		ActionPath:    privilege.ActionPath,
		ActionLimit:   privilege.ActionLimit,
		LimitInDay:    privilege.LimitInDay,
	}
}

type PrivilegeUpdateResponse struct {
	PrivilegeCreateResponse
}

func NewPrivilegeUpdateResponse(privilege *model.Privilege) *PrivilegeUpdateResponse {
	updateResponse := &PrivilegeUpdateResponse{}
	updateResponse.Id = privilege.ID.String()
	updateResponse.PrivilegeName = privilege.PrivilegeName
	updateResponse.ActionLimit = privilege.ActionLimit
	updateResponse.ActionLimit = privilege.ActionLimit
	updateResponse.LimitInDay = privilege.LimitInDay
	return updateResponse
}

type PrivilegeGetAllResponse struct {
	PrivilegeCreateResponse
}

func NewPrivilegeGetAllResponse(privileges *[]model.Privilege) *[]PrivilegeGetAllResponse {
	var privilegesResponse []PrivilegeGetAllResponse

	for _, privilege := range *privileges {
		response := &PrivilegeGetAllResponse{}
		response.Id = privilege.ID.String()
		response.PrivilegeName = privilege.PrivilegeName
		response.ActionPath = privilege.ActionPath
		response.ActionLimit = privilege.ActionLimit
		response.LimitInDay = privilege.LimitInDay
		privilegesResponse = append(privilegesResponse, PrivilegeGetAllResponse(*response))
	}

	return &privilegesResponse
}
