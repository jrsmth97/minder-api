package param

import (
	"minder/src/server/model"
)

type PrivilegeCreate struct {
	PrivilegeName string `validate:"required" json:"privilege_name"`
	ActionPath    string `validate:"required" json:"action_path"`
	ActionLimit   uint16 `validate:"required" json:"action_limit"`
	LimitInDay    uint16 `validate:"required" json:"limit_in_day"`
}

type PrivilegeUpdate struct {
	PrivilegeCreate
}

func (l *PrivilegeCreate) ParseToModel() *model.Privilege {
	return &model.Privilege{
		PrivilegeName: l.PrivilegeName,
		ActionPath:    l.ActionPath,
		ActionLimit:   l.ActionLimit,
		LimitInDay:    l.LimitInDay,
	}
}
