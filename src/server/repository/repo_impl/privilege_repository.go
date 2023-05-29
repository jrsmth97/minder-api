package repo_impl

import (
	"minder/src/server/model"
	"minder/src/server/repository"
	"strings"

	"gorm.io/gorm"
)

type privilegeRepo struct {
	db *gorm.DB
}

func NewPrivilegeRepo(db *gorm.DB) repository.PrivilegeRepo {
	return &privilegeRepo{
		db: db,
	}
}

func (r *privilegeRepo) GetPrivileges() (*[]model.Privilege, error) {
	var privileges []model.Privilege

	err := r.db.Where("deleted_at IS NULL").Find(&privileges).Error
	if err != nil {
		return nil, err
	}

	return &privileges, nil
}

func (r *privilegeRepo) GetPrivilegeById(privilegeId string) (*model.Privilege, error) {
	var privilege model.Privilege
	err := r.db.Where("privileges.id = ?", privilegeId).Where("privileges.deleted_at IS NULL").First(&privilege).Error
	if err != nil {
		return nil, err
	}

	return &privilege, nil
}

func (r *privilegeRepo) GetPrivilegeByBulkId(privilegeIds []string) (*[]model.Privilege, error) {
	var privileges []model.Privilege
	err := r.db.Where("privileges.id IN (?)", strings.Join(privilegeIds, ",")).Where("privileges.deleted_at IS NULL").Find(&privileges).Error
	if err != nil {
		return nil, err
	}

	return &privileges, nil
}

func (r *privilegeRepo) CreatePrivilege(privilege *model.Privilege) error {
	err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(privilege).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

func (r *privilegeRepo) UpdatePrivilege(id string, privilege *model.Privilege) error {
	return r.db.Where("id = ?", id).Updates(privilege).Error
}

func (r *privilegeRepo) DeletePrivilege(id string) error {
	err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Exec("DELETE from privileges WHERE id = ?", id).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}
