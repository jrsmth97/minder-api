package service

import (
	"database/sql"
	"fmt"
	"minder/src/helper"
	"minder/src/server/enums"
	"minder/src/server/model"
	"minder/src/server/param"
	"minder/src/server/repository"
	"minder/src/server/view"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/midtrans/midtrans-go"
)

type PurchaseService struct {
	repo               repository.PurchaseRepo
	membershipRepo     repository.MembershipRepo
	userMembershipRepo repository.UserMembershipRepo
	userRepo           repository.UserRepo
}

func NewPurchaseServices(
	repo repository.PurchaseRepo,
	membershipRepo repository.MembershipRepo,
	userMembershipRepo repository.UserMembershipRepo,
	userRepo repository.UserRepo,
) *PurchaseService {
	return &PurchaseService{
		repo:               repo,
		membershipRepo:     membershipRepo,
		userMembershipRepo: userMembershipRepo,
		userRepo:           userRepo,
	}
}

func (p *PurchaseService) Preparation(c *gin.Context) {
	_context = c
}

func (p *PurchaseService) GetPurchases() *view.Response {
	purchases, err := p.repo.GetAllPurchases()

	if err != nil {
		if err == sql.ErrNoRows {
			return view.ErrNotFound()
		}
		return view.ErrInternalServer(err.Error())
	}

	return view.SuccessFind(purchases)
}

func (p *PurchaseService) GetPurchaseById(purchaseId string) *view.Response {
	purchase, err := p.repo.GetPurchaseById(purchaseId)
	if err != nil {
		return view.ErrNotFound(err.Error())
	}

	return view.SuccessFind(view.NewFindPurchaseResponse(purchase))
}

func (p *PurchaseService) CreatePurchase(req *param.CreatePurchase) *view.Response {
	purchase := req.ParseToModel()

	var validBankMethods = []string{
		string(midtrans.BankBca),
		string(midtrans.BankBri),
		string(midtrans.BankBni),
	}

	if !helper.StringContains(validBankMethods, purchase.PaymentMethod) {
		return view.ErrBadRequest("BCA / BRI / BNI payment method only")
	}

	purchase.InvoiceNumber = helper.GenerateInvoiceNumber(p.repo)
	fmt.Println("InvoiceNumber=> " + purchase.InvoiceNumber)
	membership, err := p.membershipRepo.GetMembershipById(purchase.MembershipId)
	if err != nil {
		return view.ErrBadRequest(err.Error())
	}

	userId := _context.GetString("USER_ID")
	user, err := p.userRepo.FindById(userId)
	if err != nil {
		return view.ErrBadRequest(err.Error())
	}

	purchase.Amount = membership.Price
	purchase.Status = enums.PaymentStatusPending
	purchase.UserId = user.ID.String()

	chargeParam := helper.GenerateChargeRequest(purchase, membership, user)
	midtransResp, err := helper.ChargeTransaction(chargeParam)
	if midtransResp.StatusCode != "201" {
		return view.ErrInternalServer(err.Error())
	}

	fmt.Println("midtrans response => " + midtransResp.StatusMessage)
	purchase.VaNumber = midtransResp.VaNumbers[0].VANumber
	err = p.repo.CreatePurchase(purchase)
	if err != nil {
		return view.ErrInternalServer(err.Error())
	}

	return view.SuccessCreated(view.NewFindPurchaseResponse(purchase))
}

func (p *PurchaseService) CancelPurchase(purchaseId string) *view.Response {
	purchase, errPurchase := p.repo.GetPurchaseById(purchaseId)
	if errPurchase != nil {
		return view.ErrBadRequest("purchase doesn't exists")
	}

	midtransResp, err := helper.CancelTransaction(purchase.InvoiceNumber)
	if midtransResp.StatusCode != "201" {
		return view.ErrInternalServer(err.Error())
	}

	purchase.Status = enums.PaymentStatusCancel
	errUpdate := p.repo.UpdatePurchase(purchaseId, purchase)
	if errUpdate != nil {
		return view.ErrInternalServer(err.Error())
	}

	return view.SuccessUpdated(purchase)
}

func (p *PurchaseService) SyncPurchase() *view.Response {
	pendingPurchases, err := p.repo.GetPendingPurchases()
	if err != nil {
		return view.ErrBadRequest("pending purchase doesn't exists")
	}

	var syncPurchases []model.Purchase
	for i := 0; i < len(*pendingPurchases); i++ {
		purchase := (*pendingPurchases)[i]
		getMidtransStatus, err := helper.GetTransactionStatus(purchase.InvoiceNumber)
		if err != nil {
			continue
		}

		if getMidtransStatus.TransactionStatus == enums.PaymentStatusSuccess {
			membership, err := p.membershipRepo.GetMembershipById(purchase.MembershipId)
			if err != nil {
				view.ErrBadRequest("Membership not found")
			}

			validUntilDate := time.Now().Add((time.Hour * 24 * 30) * time.Duration(membership.DurationInMonth))
			userMember := model.UserMembership{
				UserId:       purchase.UserId,
				MembershipId: purchase.MembershipId,
				ValidUntil:   validUntilDate,
			}

			err = p.userMembershipRepo.Create(&userMember)
			if err != nil {
				return view.ErrInternalServer(err.Error())
			}
		}

		sync := false
		if purchase.Status != getMidtransStatus.TransactionStatus {
			sync = true
		}

		purchase.Status = getMidtransStatus.TransactionStatus
		if sync {
			syncPurchases = append(syncPurchases, purchase)
		}

		errUpdate := p.repo.UpdatePurchase(purchase.ID.String(), &purchase)
		if errUpdate != nil {
			return view.ErrInternalServer(err.Error())
		}
	}

	return view.SuccessUpdated(view.NewSyncPurchaseResponse(&syncPurchases))
}
