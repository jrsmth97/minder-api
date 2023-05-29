package repo_impl

import (
	"fmt"
	"minder/src/helper"
	"minder/src/server/model"
	"minder/src/server/repository"

	"gorm.io/gorm"
)

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) repository.UserRepo {
	return &userRepo{
		db: db,
	}
}

func (r *userRepo) Create(user *model.User) error {
	err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(user).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

func (r *userRepo) Explore() (*[]model.User, error) {
	var users []model.User

	exploreLimit := 100
	var countUsers int
	r.db.Raw("SELECT count(*) FROM users").Scan(&countUsers)
	fmt.Println("we have ", countUsers, " user rows")

	randomOffset := 0
	if countUsers > exploreLimit {
		randomOffset = helper.RandomNumber(countUsers - exploreLimit)
	}

	err := r.dbQuery().Find(&users).Offset(randomOffset).Limit(exploreLimit).Error

	if err != nil {
		return nil, err
	}
	return &users, nil
}

func (r *userRepo) CountUsers() (*int, error) {
	var countUsers int
	err := r.db.Raw("SELECT count(*) FROM users").Scan(&countUsers).Error
	fmt.Println("we have ", countUsers, " user rows")

	if err != nil {
		return nil, err
	}

	return &countUsers, nil
}

func (r *userRepo) FindById(id string) (*model.User, error) {
	var user model.User
	err := r.dbQuery().Where("users.id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepo) FindByEmail(email string) (*model.User, error) {
	var user model.User
	err := r.dbQuery().Where("users.email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepo) Update(id string, user *model.User) error {
	err := r.db.Transaction(func(tx *gorm.DB) error {
		photos := user.Photos
		interests := user.Interests

		user.Photos = nil
		user.Interests = nil
		if err := tx.Table("users").Where("id = ?", id).Updates(user).Error; err != nil {
			return err
		}

		if err := tx.Exec("DELETE from user_photos WHERE user_id = ?", id).Error; err != nil {
			return err
		}

		for _, photo := range photos {
			photo.UserId = id
			if err := tx.Create(&photo).Error; err != nil {
				return err
			}
		}

		if err := tx.Exec("DELETE from user_interests WHERE user_id = ?", id).Error; err != nil {
			return err
		}

		for _, interest := range interests {
			interest.UserId = id
			if err := tx.Table("user_interests").Create(&interest).Error; err != nil {
				return err
			}
		}

		return nil
	})
	return err
}

func (r *userRepo) Delete(id string) error {
	err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Exec("DELETE from user_photos WHERE user_id = ?", id).Error; err != nil {
			return err
		}

		if err := tx.Exec("DELETE from user_memberships WHERE user_id = ?", id).Error; err != nil {
			return err
		}

		if err := tx.Exec("DELETE from user_interests WHERE user_id = ?", id).Error; err != nil {
			return err
		}

		if err := tx.Exec("DELETE from user_swipes WHERE user_id = ?", id).Error; err != nil {
			return err
		}

		if err := tx.Exec("DELETE from purchases WHERE user_id = ?", id).Error; err != nil {
			return err
		}

		if err := tx.Exec("DELETE from users WHERE id = ?", id).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

func (r *userRepo) dbQuery() *gorm.DB {
	return r.db.Preload("Location").Preload("Interests").Preload("Interests.Interest").Preload("Photos")
}
