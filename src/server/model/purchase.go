package model

type Purchase struct {
	BaseModel
	UserId        string `json:"user_id"`
	InvoiceNumber string `json:"invoice_number"`
	MembershipId  string `json:"membership_id"`
	Amount        uint32 `json:"amount"`
	PaymentMethod string `json:"payment_method"`
	Status        string `json:"status"`
	VaNumber      string `json:"va_number"`

	User       User       `json:"user"`
	Membership Membership `json:"membership"`
}

var Purchases = []Purchase{}
