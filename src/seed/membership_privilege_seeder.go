package seed

import (
	"log"
	"minder/src/helper"
	"minder/src/server/enums"
	"minder/src/server/model"

	"gorm.io/gorm"
)

func SeedMembershipPrivilege(db *gorm.DB) {
	var memberships []model.Membership
	db.Find(&memberships)

	var privileges []model.Privilege
	db.Find(&privileges)

	for _, member := range memberships {
		var memberPrivilegesExist []model.MembershipPrivilege
		err := db.Where("membership_id=?", member.ID.String()).Find(&memberPrivilegesExist).Error
		if len(memberPrivilegesExist) > 0 {
			continue
		}

		switch member.MembershipName {
		case enums.FreePlanMembership:
			privilegeIdx := searchPrivilegeIdxByName(privileges, enums.LimitedSwipesPrivilege)
			freePlanPrivilege := model.MembershipPrivilege{
				MembershipId: member.ID.String(),
				PrivilegeId:  privileges[privilegeIdx].ID.String(),
			}

			err = db.Debug().Model(&model.MembershipPrivilege{}).Create(&freePlanPrivilege).Error
		case enums.SilverPlanMembership:
			privilege1Idx := searchPrivilegeIdxByName(privileges, enums.UnlimitedSwipesPrivilege)
			privilege2Idx := searchPrivilegeIdxByName(privileges, enums.VerifiedLabelPrivilege)
			silverPlanPrivileges := []model.MembershipPrivilege{
				{
					MembershipId: member.ID.String(),
					PrivilegeId:  privileges[privilege1Idx].ID.String(),
				},
				{
					MembershipId: member.ID.String(),
					PrivilegeId:  privileges[privilege2Idx].ID.String(),
				},
			}

			for _, privilege := range silverPlanPrivileges {
				err = db.Debug().Model(&model.MembershipPrivilege{}).Create(&privilege).Error
			}
		case enums.GoldPlanMembership:
			privilege1Idx := searchPrivilegeIdxByName(privileges, enums.UnlimitedLikesPrivilege)
			privilege2Idx := searchPrivilegeIdxByName(privileges, enums.UnlimitedSwipesPrivilege)
			privilege3Idx := searchPrivilegeIdxByName(privileges, enums.VerifiedLabelPrivilege)
			goldPlanPrivileges := []model.MembershipPrivilege{
				{
					MembershipId: member.ID.String(),
					PrivilegeId:  privileges[privilege1Idx].ID.String(),
				},
				{
					MembershipId: member.ID.String(),
					PrivilegeId:  privileges[privilege2Idx].ID.String(),
				},
				{
					MembershipId: member.ID.String(),
					PrivilegeId:  privileges[privilege3Idx].ID.String(),
				},
			}

			for _, privilege := range goldPlanPrivileges {
				err = db.Debug().Model(&model.MembershipPrivilege{}).Create(&privilege).Error
			}
		case enums.DiamondPlanMembership:
			privilege1Idx := searchPrivilegeIdxByName(privileges, enums.UnlimitedLikesPrivilege)
			privilege2Idx := searchPrivilegeIdxByName(privileges, enums.UnlimitedSwipesPrivilege)
			privilege3Idx := searchPrivilegeIdxByName(privileges, enums.UnlimitedFavouritesPrivilege)
			privilege4Idx := searchPrivilegeIdxByName(privileges, enums.VerifiedLabelPrivilege)
			diaomondPlanPrivileges := []model.MembershipPrivilege{
				{
					MembershipId: member.ID.String(),
					PrivilegeId:  privileges[privilege1Idx].ID.String(),
				},
				{
					MembershipId: member.ID.String(),
					PrivilegeId:  privileges[privilege2Idx].ID.String(),
				},
				{
					MembershipId: member.ID.String(),
					PrivilegeId:  privileges[privilege3Idx].ID.String(),
				},
				{
					MembershipId: member.ID.String(),
					PrivilegeId:  privileges[privilege4Idx].ID.String(),
				},
			}

			for _, privilege := range diaomondPlanPrivileges {
				err = db.Debug().Model(&model.MembershipPrivilege{}).Create(&privilege).Error
			}
		}

		if err != nil {
			log.Fatalf("cannot seed membership privileges table: %v", err)
		}

	}
}

func searchPrivilegeIdxByName(privileges []model.Privilege, name string) int {
	return helper.FindIndex(privileges, func(value interface{}) bool {
		return value.(model.Privilege).PrivilegeName == name
	})
}
