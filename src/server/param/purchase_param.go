package param

import (
	"minder/src/server/model"
)

type CreatePurchase struct {
	MembershipId  string `validate:"required" json:"membership_id"`
	PaymentMethod string `validate:"required" json:"payment_method"`
}

func (p *CreatePurchase) ParseToModel() *model.Purchase {
	return &model.Purchase{
		MembershipId:  p.MembershipId,
		PaymentMethod: p.PaymentMethod,
	}
}
