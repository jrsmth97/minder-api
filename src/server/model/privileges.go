package model

type Privilege struct {
	BaseModel
	PrivilegeName string `json:"privilege_name"`
	ActionPath    string `json:"action_path"`
	ActionLimit   uint16 `json:"action_limit"`
	LimitInDay    uint16 `json:"limit_iy_day"`
}

var Privileges = []Privilege{}
