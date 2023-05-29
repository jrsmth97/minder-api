package helper

import (
	"minder/src/server/model"
	"os"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
)

func setupGlobalMidtransConfigApi() {
	midtrans.ServerKey = os.Getenv("MIDTRANS_SERVER_KEY")
	midtrans.Environment = midtrans.Sandbox
}

func GenerateChargeRequest(purchase *model.Purchase, membership *model.Membership, user *model.User) *coreapi.ChargeReq {
	phone := "08123456789"
	if purchase.User.Phone != "" {
		phone = purchase.User.Phone
	}

	return &coreapi.ChargeReq{
		PaymentType: coreapi.PaymentTypeBankTransfer,
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  purchase.InvoiceNumber,
			GrossAmt: int64(purchase.Amount),
		},
		Items: &[]midtrans.ItemDetails{
			{
				ID:    purchase.MembershipId,
				Price: int64(purchase.Amount),
				Qty:   1,
				Name:  membership.MembershipName,
			},
		},
		CustomerDetails: &midtrans.CustomerDetails{
			FName: user.Name,
			LName: "",
			Email: user.Email,
			Phone: phone,
		},
		BankTransfer: &coreapi.BankTransferDetails{
			Bank: midtrans.Bank(purchase.PaymentMethod),
		},
	}
}

func ChargeTransaction(chargeReq *coreapi.ChargeReq) (*coreapi.ChargeResponse, *midtrans.Error) {
	setupGlobalMidtransConfigApi()
	return coreapi.ChargeTransaction(chargeReq)
}

func CancelTransaction(orderId string) (*coreapi.ChargeResponse, *midtrans.Error) {
	setupGlobalMidtransConfigApi()
	return coreapi.CancelTransaction(orderId)
}

func GetTransactionStatus(orderId string) (*coreapi.TransactionStatusResponse, *midtrans.Error) {
	setupGlobalMidtransConfigApi()
	return coreapi.CheckTransaction(orderId)
}
