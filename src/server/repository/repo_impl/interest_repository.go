package repo_impl

import (
	"minder/src/server/model"
	"minder/src/server/repository"

	"gorm.io/gorm"
)

type interestRepo struct {
	db *gorm.DB
}

func NewInterestRepo(db *gorm.DB) repository.InterestRepo {
	return &interestRepo{
		db: db,
	}
}

func (r *interestRepo) GetInterests() (*[]model.Interest, error) {
	var interests []model.Interest

	err := r.db.Find(&interests).Error
	if err != nil {
		return nil, err
	}

	return &interests, nil
}

func (r *interestRepo) CreateInterest(interest *model.Interest) error {
	err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(interest).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

func (r *interestRepo) UpdateInterest(id string, interest *model.Interest) error {
	return r.db.Where("id = ?", id).Updates(interest).Error
}

func (r *interestRepo) DeleteInterest(id string) error {
	err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Exec("DELETE from interests WHERE id = ?", id).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

func (r *interestRepo) GetInterestById(interestId string) (*model.Interest, error) {
	var interest model.Interest
	err := r.db.Where("interests.id=?", interestId).First(&interest).Error
	if err != nil {
		return nil, err
	}

	return &interest, nil
}
