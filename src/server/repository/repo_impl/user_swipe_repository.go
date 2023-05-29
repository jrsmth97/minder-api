package repo_impl

import (
	"minder/src/server/model"
	"minder/src/server/repository"

	"gorm.io/gorm"
)

type userSwipeRepo struct {
	db *gorm.DB
}

func NewUserSwipeRepo(db *gorm.DB) repository.UserSwipeRepo {
	return &userSwipeRepo{
		db: db,
	}
}

func (r *userSwipeRepo) Create(userSwipe *model.UserSwipe) error {
	err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(userSwipe).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

func (r *userSwipeRepo) FindSwipes(userId string) (*[]model.UserSwipe, error) {
	var userSwipes []model.UserSwipe
	err := r.db.Where("user_swipes.user_id = ?", userId).Find(&userSwipes).Error
	if err != nil {
		return nil, err
	}
	return &userSwipes, nil
}

func (r *userSwipeRepo) FindSwipesToday(userId string) (*[]model.UserSwipe, error) {
	var userSwipes []model.UserSwipe
	err := r.db.Where("(user_swipes.created_at BETWEEN CONCAT(CURDATE(), ' 00:00:00') AND CONCAT(CURDATE(), ' 23:59:59'))").Where("user_swipes.user_id = ?", userId).Find(&userSwipes).Error
	if err != nil {
		return nil, err
	}
	return &userSwipes, nil
}

func (r *userSwipeRepo) FindLike(userId string) (*[]model.UserSwipe, error) {
	var userSwipes []model.UserSwipe
	err := r.db.Where("user_swipes.action = 'like'").Where("user_swipes.user_id = ?", userId).Find(&userSwipes).Error
	if err != nil {
		return nil, err
	}
	return &userSwipes, nil
}

func (r *userSwipeRepo) FindLikeToday(userId string) (*[]model.UserSwipe, error) {
	var userSwipes []model.UserSwipe
	err := r.db.Where("(user_swipes.created_at BETWEEN CONCAT(CURDATE(), ' 00:00:00') AND CONCAT(CURDATE(), ' 23:59:59'))").Where("user_swipes.action = 'like'").Where("user_swipes.user_id = ?", userId).Find(&userSwipes).Error
	if err != nil {
		return nil, err
	}
	return &userSwipes, nil
}

func (r *userSwipeRepo) FindPass(userId string) (*[]model.UserSwipe, error) {
	var userSwipes []model.UserSwipe
	err := r.db.Where("user_swipes.action = 'pass'").Where("user_swipes.user_id = ?", userId).Find(&userSwipes).Error
	if err != nil {
		return nil, err
	}
	return &userSwipes, nil
}

func (r *userSwipeRepo) FindPassToday(userId string) (*[]model.UserSwipe, error) {
	var userSwipes []model.UserSwipe
	err := r.db.Where("(user_swipes.created_at BETWEEN CONCAT(CURDATE(), ' 00:00:00') AND CONCAT(CURDATE(), ' 23:59:59'))").Where("user_swipes.action = 'pass'").Where("user_swipes.user_id = ?", userId).Find(&userSwipes).Error
	if err != nil {
		return nil, err
	}
	return &userSwipes, nil
}

func (r *userSwipeRepo) FindFavourite(userId string) (*[]model.UserSwipe, error) {
	var userSwipes []model.UserSwipe
	err := r.db.Where("user_swipes.action = 'pass'").Where("user_swipes.user_id = ?", userId).Find(&userSwipes).Error
	if err != nil {
		return nil, err
	}
	return &userSwipes, nil
}

func (r *userSwipeRepo) FindFavouriteToday(userId string) (*[]model.UserSwipe, error) {
	var userSwipes []model.UserSwipe
	err := r.db.Where("(user_swipes.created_at BETWEEN CONCAT(CURDATE(), ' 00:00:00') AND CONCAT(CURDATE(), ' 23:59:59'))").Where("user_swipes.action = 'favourite'").Where("user_swipes.user_id = ?", userId).Find(&userSwipes).Error
	if err != nil {
		return nil, err
	}
	return &userSwipes, nil
}

func (r *userSwipeRepo) FindById(id string) (*model.UserSwipe, error) {
	var userSwipe model.UserSwipe
	err := r.db.Where("id=?", id).First(&userSwipe).Error
	if err != nil {
		return nil, err
	}
	return &userSwipe, nil
}
