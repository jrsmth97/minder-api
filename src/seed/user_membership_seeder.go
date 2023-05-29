package seed

import (
	"log"
	"minder/src/helper"
	"minder/src/server/model"
	"time"

	"gorm.io/gorm"
)

func SeedUserMembership(db *gorm.DB) {
	var users []model.User
	db.Find(&users)

	var memberships []model.Membership
	db.Find(&memberships)

	for _, user := range users {
		var userMembershipExis []model.UserMembership
		err := db.Where("user_id=?", user.ID.String()).Find(&userMembershipExis).Error
		if len(userMembershipExis) > 0 {
			continue
		}

		var specialUsers = []string{
			"Super Admin",
			"User 1",
			"User 2",
			"User 3",
		}

		randomMembershipIndex := helper.RandomNumber(len(memberships) - 1)
		membership := memberships[randomMembershipIndex]
		membershipId := membership.ID.String()

		if helper.StringContains(specialUsers, user.Name) {
			idx := helper.FindIndex(memberships, func(value interface{}) bool {
				return value.(model.Membership).MembershipName == "Diamond Plan"
			})
			membership = memberships[idx]
			membershipId = membership.ID.String()
		}

		userMembership := model.UserMembership{
			UserId:       user.ID.String(),
			MembershipId: membershipId,
			ValidUntil:   time.Now().Add((time.Hour * 24 * 30) * time.Duration(membership.DurationInMonth)),
		}

		err = db.Debug().Model(&model.UserMembership{}).Create(&userMembership).Error
		if err != nil {
			log.Fatalf("cannot seed user membership table: %v", err)
		}

	}
}
