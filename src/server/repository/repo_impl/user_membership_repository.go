package repo_impl

import (
	"minder/src/server/model"
	"minder/src/server/repository"

	"gorm.io/gorm"
)

type userMembershipRepo struct {
	db *gorm.DB
}

func NewUserMembershipRepo(db *gorm.DB) repository.UserMembershipRepo {
	return &userMembershipRepo{
		db: db,
	}
}

func (r *userMembershipRepo) Create(userMember *model.UserMembership) error {
	err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(userMember).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

func (r *userMembershipRepo) FindById(id string) (*model.UserMembership, error) {
	var userMember model.UserMembership
	err := r.db.Preload("User").Preload("Membership").Where("user_memberships.id = ?", id).Order("user_memberships.created_at DESC").First(&userMember).Error
	if err != nil {
		return nil, err
	}

	return &userMember, nil
}

func (r *userMembershipRepo) FindByUserId(id string) (*model.UserMembership, error) {
	var userMember model.UserMembership
	err := r.db.Preload("User").Preload("Membership").Where("user_memberships.user_id = ?", id).Having("user_memberships.valid_until >= NOW()").Order("user_memberships.created_at DESC").First(&userMember).Error
	if err != nil {
		return nil, err
	}

	return &userMember, nil
}
