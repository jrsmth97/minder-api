package seed

import (
	"log"
	"minder/src/server/model"

	"gorm.io/gorm"
)

func SeedMembership(db *gorm.DB) {
	for _, membership := range memberships {
		var membershipExist []model.Membership
		err := db.Where("membership_name=?", membership.MembershipName).Find(&membershipExist).Error
		if err != nil {
			continue
		}

		if len(membershipExist) > 0 {
			continue
		}

		err = db.Debug().Model(&model.Membership{}).Create(&membership).Error
		if err != nil {
			log.Fatalf("cannot seed membership table: %v", err)
		}
	}
}

var memberships = []model.Membership{
	{
		MembershipName:  "Free Plan",
		Price:           0,
		DurationInMonth: 0,
		Description:     "Limited Swipes (10 Swipes per Day)",
	},
	{
		MembershipName:  "Silver Plan",
		Price:           20000,
		DurationInMonth: 1,
		Description:     "Verified label, Unlimited swipes",
	},
	{
		MembershipName:  "Gold Plan",
		Price:           30000,
		DurationInMonth: 1,
		Description:     "Verified label, Unlimited swipes, Unlimited Likes",
	},
	{
		MembershipName:  "Diamond Plan",
		Price:           40000,
		DurationInMonth: 1,
		Description:     "Verified label, Unlimited swipes, Unlimited Likes, Unlimited Favourite",
	},
}
