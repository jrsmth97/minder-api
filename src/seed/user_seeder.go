package seed

import (
	"fmt"
	"log"
	"minder/src/helper"
	"minder/src/server/model"

	"gorm.io/gorm"
)

func SeedUser(db *gorm.DB, amount int) {
	var locations []model.Location
	db.Find(&locations)

	for i := 1; i <= amount; i++ {
		var userExist []model.User

		userEmail := fmt.Sprintf("user%v@mail.com", i)
		err := db.Where("email=?", userEmail).Find(&userExist).Error
		if len(userExist) > 0 {
			continue
		}

		user := &model.User{}
		randomLocationIndex := helper.RandomNumber(len(locations) - 1)
		user.LocationId = locations[randomLocationIndex].ID.String()
		genderRandNumber := helper.RandomNumber(2)
		user.Name = fmt.Sprintf("User %v", i)
		user.Email = userEmail
		user.Gender = uint8(genderRandNumber)
		user.Password = "user"
		user.BirthDate = "2000-01-01"

		hashPassword, err := helper.GeneratePassword(user.Password)
		user.Password = hashPassword
		err = db.Debug().Model(&model.User{}).Create(&user).Error
		if err != nil {
			log.Fatalf("cannot seed user table: %v", err)
		}
	}
}
