package repo_impl

import (
	"fmt"
	"minder/src/server/enums"
	"minder/src/server/model"
	"minder/src/server/repository"

	"gorm.io/gorm"
)

type purchaseRepo struct {
	db *gorm.DB
}

func NewPurchaseRepo(db *gorm.DB) repository.PurchaseRepo {
	return &purchaseRepo{
		db: db,
	}
}

func (r *purchaseRepo) GetAllPurchases() (*[]model.Purchase, error) {
	var purchases []model.Purchase

	err := r.db.Find(&purchases).Error
	if err != nil {
		return nil, err
	}

	return &purchases, nil
}

func (r *purchaseRepo) CreatePurchase(purchase *model.Purchase) error {
	err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(purchase).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

func (r *purchaseRepo) GetPendingPurchases() (*[]model.Purchase, error) {
	var purchases []model.Purchase
	err := r.db.Where("status = ?", enums.PaymentStatusPending).Find(&purchases).Error
	if err != nil {
		return nil, err
	}

	return &purchases, nil
}

func (r *purchaseRepo) GetFailedPurchases() (*[]model.Purchase, error) {
	var purchases []model.Purchase
	err := r.db.Where("status IN (?,?,?)", enums.PaymentStatusExpired, enums.PaymentStatusCancel, enums.PaymentStatusDeny).Find(&purchases).Error
	if err != nil {
		return nil, err
	}

	return &purchases, nil
}

func (r *purchaseRepo) GetSuccessPurchases() (*[]model.Purchase, error) {
	var purchases []model.Purchase
	err := r.db.Where("status = ?", enums.PaymentStatusSuccess).Find(&purchases).Error
	if err != nil {
		return nil, err
	}

	return &purchases, nil
}

func (r *purchaseRepo) GetPurchaseById(id string) (*model.Purchase, error) {
	var purchase model.Purchase
	err := r.db.Preload("User").Preload("Membership").Where("id = ?", id).First(&purchase).Error
	if err != nil {
		return nil, err
	}

	return &purchase, nil
}

func (r *purchaseRepo) GetPurchaseByIdWithoutInclude(id string) (*model.Purchase, error) {
	var purchase model.Purchase
	err := r.db.Where("id = ?", id).First(&purchase).Error
	if err != nil {
		return nil, err
	}

	return &purchase, nil
}

func (r *purchaseRepo) GetLastPurchaseByDate(date string) (*model.Purchase, error) {
	var lastPurchase model.Purchase
	err := r.db.Where(fmt.Sprintf("invoice_number LIKE 'INV-%sMI%s'", date, "%")).Order("created_at DESC").First(&lastPurchase).Error
	if err != nil {
		return nil, err
	}

	return &lastPurchase, nil

}

func (r *purchaseRepo) UpdatePurchase(id string, purchase *model.Purchase) error {
	return r.db.Where("id = ?", id).Updates(purchase).Error
}
