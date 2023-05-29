package param

import (
	"minder/src/server/model"
)

type InterestCreate struct {
	InterestName string `validate:"required" json:"interest_name"`
}

type InterestUpdate struct {
	InterestCreate
}

func (p *InterestCreate) ParseToModel() *model.Interest {
	return &model.Interest{
		InterestName: p.InterestName,
	}
}
