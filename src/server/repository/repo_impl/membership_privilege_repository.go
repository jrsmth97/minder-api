package repo_impl

import (
	"fmt"
	"minder/src/server/model"
	"minder/src/server/repository"
	"strings"

	"gorm.io/gorm"
)

type membershipPrivilegeRepo struct {
	db *gorm.DB
}

func NewMembershipPrivilegeRepo(db *gorm.DB) repository.MembershipPrivilegeRepo {
	return &membershipPrivilegeRepo{
		db: db,
	}
}

func (r *membershipPrivilegeRepo) GetMembershipPrivileges() (*[]model.MembershipPrivilege, error) {
	var membershipPrivileges []model.MembershipPrivilege

	err := r.db.Preload("Membership").Preload("Privilege").Where("deleted_at IS NULL").Find(&membershipPrivileges).Error
	if err != nil {
		return nil, err
	}

	return &membershipPrivileges, nil
}

func (r *membershipPrivilegeRepo) GetMembershipPrivilegeById(membershipPrivilegeId string) (*model.MembershipPrivilege, error) {
	var membershipPrivilege model.MembershipPrivilege
	err := r.db.Preload("Membership").Preload("Privilege").Where("membership_privileges.id = ?", membershipPrivilegeId).Where("membership_privileges.deleted_at IS NULL").First(&membershipPrivilege).Error
	if err != nil {
		return nil, err
	}

	return &membershipPrivilege, nil
}

func (r *membershipPrivilegeRepo) GetMembershipPrivilegeByMemberId(membershipPrivilegeId string) (*[]model.MembershipPrivilege, error) {
	var membershipPrivileges []model.MembershipPrivilege
	err := r.db.Preload("Membership").Preload("Privilege").Where("membership_privileges.membership_id = ?", membershipPrivilegeId).Where("membership_privileges.deleted_at IS NULL").Find(&membershipPrivileges).Error
	if err != nil {
		return nil, err
	}

	return &membershipPrivileges, nil
}

func (r *membershipPrivilegeRepo) GetMembershipPrivilegeByBulkMemberId(membershipPrivilegeIds []string) (*[]model.MembershipPrivilege, error) {
	var membershipPrivileges []model.MembershipPrivilege
	bulkId := strings.Join(membershipPrivilegeIds[:], ",")
	fmt.Println("bulkId => " + bulkId)
	err := r.db.Preload("Membership").Preload("Privilege").Where("membership_privileges.membership_id IN (?)", bulkId).Where("membership_privileges.deleted_at IS NULL").Find(&membershipPrivileges).Error
	if err != nil {
		return nil, err
	}

	return &membershipPrivileges, nil
}

func (r *membershipPrivilegeRepo) CreateMembershipPrivilege(membershipPrivilege *model.MembershipPrivilege) error {
	err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(membershipPrivilege).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

func (r *membershipPrivilegeRepo) UpdateMembershipPrivilege(id string, membershipPrivilege *model.MembershipPrivilege) error {
	return r.db.Where("id = ?", id).Updates(membershipPrivilege).Error
}

func (r *membershipPrivilegeRepo) DeleteMembershipPrivilege(id string) error {
	err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Exec("DELETE from membership_privileges WHERE id = ?", id).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

func (r *membershipPrivilegeRepo) DeleteMembershipPrivilegeByMemberId(id string) error {
	err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Exec("DELETE from membership_privileges WHERE membership_id = ?", id).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

func (r *membershipPrivilegeRepo) DeleteMembershipPrivilegeByPrivilegeId(id string) error {
	err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Exec("DELETE from membership_privileges WHERE privilege_id = ?", id).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}
