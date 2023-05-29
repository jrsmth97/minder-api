package seed

import (
	"log"
	"minder/src/server/model"

	"gorm.io/gorm"
)

func SeedPrivilege(db *gorm.DB) {
	for _, privilege := range privileges {
		var privilegeExist []model.Privilege
		err := db.Where("privilege_name=?", privilege.PrivilegeName).Find(&privilegeExist).Error
		if err != nil {
			continue
		}

		if len(privilegeExist) > 0 {
			continue
		}

		err = db.Debug().Model(&model.Privilege{}).Create(&privilege).Error
		if err != nil {
			log.Fatalf("cannot seed privilege table: %v", err)
		}
	}
}

var privileges = []model.Privilege{
	{
		PrivilegeName: "Limited Swipes",
		ActionPath:    "/swipes",
		ActionLimit:   10,
		LimitInDay:    1,
	},
	{
		PrivilegeName: "Unlimited Swipes",
		ActionPath:    "/swipes",
		ActionLimit:   0,
		LimitInDay:    0,
	},
	{
		PrivilegeName: "Unlimited Likes",
		ActionPath:    "/swipes/like",
		ActionLimit:   0,
		LimitInDay:    0,
	},
	{
		PrivilegeName: "Unlimited Favourites",
		ActionPath:    "/swipes/favourite",
		ActionLimit:   0,
		LimitInDay:    0,
	},
	{
		PrivilegeName: "Verified Label",
		ActionPath:    "/user/verified",
		ActionLimit:   0,
		LimitInDay:    0,
	},
}
