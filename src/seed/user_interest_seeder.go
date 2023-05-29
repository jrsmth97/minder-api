package seed

import (
	"log"
	"minder/src/helper"
	"minder/src/server/model"

	"gorm.io/gorm"
)

func SeedUserInterest(db *gorm.DB) {
	var users []model.User
	db.Find(&users)

	var interests []model.Interest
	db.Find(&interests)

	for _, user := range users {
		var userInterestExist []model.UserInterest
		err := db.Where("user_id=?", user.ID.String()).Find(&userInterestExist).Error
		if err != nil {
			continue
		}

		if len(userInterestExist) > 0 {
			continue
		}

		randomNumber := helper.RandomNumber(len(interests) - 1)
		maxRandomRange := randomNumber + 2
		if maxRandomRange > len(interests)-1 {
			maxRandomRange = len(interests) - 1
			randomNumber = maxRandomRange - 2
		}

		randomRange := helper.MakeNumberRange(randomNumber, maxRandomRange)
		randomRange = helper.ArrayIntShuffle(randomRange)

		for _, index := range randomRange {
			userInterest := model.UserInterest{
				UserId:     user.ID.String(),
				InterestId: interests[index].ID.String(),
			}

			err = db.Debug().Model(&model.UserInterest{}).Create(&userInterest).Error
			if err != nil {
				log.Fatalf("cannot seed user interest table: %v", err)
			}
		}
	}
}
