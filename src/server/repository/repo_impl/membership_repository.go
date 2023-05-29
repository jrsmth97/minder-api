package repo_impl

import (
	"minder/src/server/model"
	"minder/src/server/repository"

	"gorm.io/gorm"
)

type membershipRepo struct {
	db *gorm.DB
}

func NewMembershipRepo(db *gorm.DB) repository.MembershipRepo {
	return &membershipRepo{
		db: db,
	}
}

func (r *membershipRepo) GetMemberships() (*[]model.Membership, error) {
	var memberships []model.Membership

	err := r.db.Where("deleted_at IS NULL").Find(&memberships).Error
	if err != nil {
		return nil, err
	}

	return &memberships, nil
}

func (r *membershipRepo) GetMembershipById(membershipId string) (*model.Membership, error) {
	var membership model.Membership
	err := r.db.Where("memberships.id = ?", membershipId).Where("memberships.deleted_at IS NULL").First(&membership).Error
	if err != nil {
		return nil, err
	}

	return &membership, nil
}

func (r *membershipRepo) GetMembershipByName(membershipName string) (*model.Membership, error) {
	var membership model.Membership
	err := r.db.Where("memberships.membership_name = ?", membershipName).Where("memberships.deleted_at IS NULL").First(&membership).Error
	if err != nil {
		return nil, err
	}

	return &membership, nil
}

func (r *membershipRepo) CreateMembership(membership *model.Membership) error {
	err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(membership).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

func (r *membershipRepo) UpdateMembership(id string, membership *model.Membership) error {
	return r.db.Where("id = ?", id).Updates(membership).Error
}

func (r *membershipRepo) DeleteMembership(id string) error {
	err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Exec("DELETE from memberships WHERE id = ?", id).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}
