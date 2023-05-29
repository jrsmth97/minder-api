package service

import (
	"fmt"
	"minder/src/helper"
	"minder/src/server/enums"
	"minder/src/server/model"
	"minder/src/server/repository"
	"minder/src/server/view"

	"github.com/gin-gonic/gin"
)

var _targetUser string

type UserSwipeServices struct {
	repo                    repository.UserSwipeRepo
	userRepo                repository.UserRepo
	userMembershipRepo      repository.UserMembershipRepo
	membershipPrivilegeRepo repository.MembershipPrivilegeRepo
	privilegeRepo           repository.PrivilegeRepo
}

func NewUserSwipeServices(
	repo repository.UserSwipeRepo,
	userRepo repository.UserRepo,
	userMembershipRepo repository.UserMembershipRepo,
	membershipPrivilegeRepo repository.MembershipPrivilegeRepo,
	privilegeRepo repository.PrivilegeRepo,
) *UserSwipeServices {
	return &UserSwipeServices{
		repo:                    repo,
		userRepo:                userRepo,
		userMembershipRepo:      userMembershipRepo,
		membershipPrivilegeRepo: membershipPrivilegeRepo,
		privilegeRepo:           privilegeRepo,
	}
}

func (p *UserSwipeServices) Preparation(c *gin.Context) {
	_context = c
}

func (u *UserSwipeServices) PassAction(targetUser string) *view.Response {
	_targetUser = targetUser
	_, err := u.userRepo.FindById(targetUser)
	if err != nil {
		return view.ErrBadRequest("user not found")
	}

	valid, errMessage := u.checkPrivilege(enums.PassSwipeAction)
	if !valid {
		return view.ErrBadRequest(errMessage)
	}

	userId := _context.GetString("USER_ID")

	if userId == targetUser {
		return view.ErrBadRequest("Oops, you cant pass yourself")
	}

	userSwipe := model.UserSwipe{
		UserId:   _context.GetString("USER_ID"),
		TargetId: targetUser,
		Action:   enums.PassSwipeAction,
	}

	err = u.repo.Create(&userSwipe)
	if err != nil {
		return view.ErrInternalServer(err.Error())
	}

	resp := view.UserSwipeResponse{
		UserTarget: targetUser,
		Match:      false,
	}

	return view.SuccessAction(resp)
}

func (u *UserSwipeServices) LikeAction(targetUser string) *view.Response {
	_targetUser = targetUser
	_, err := u.userRepo.FindById(targetUser)
	if err != nil {
		return view.ErrBadRequest("user not found")
	}

	valid, errMessage := u.checkPrivilege(enums.LikeSwipeAction)
	if !valid {
		return view.ErrBadRequest(errMessage)
	}

	userId := _context.GetString("USER_ID")
	if userId == targetUser {
		return view.ErrBadRequest("Loving yourself is good, but not in here")
	}
	userSwipe := model.UserSwipe{
		UserId:   userId,
		TargetId: targetUser,
		Action:   enums.LikeSwipeAction,
	}

	err = u.repo.Create(&userSwipe)
	if err != nil {
		return view.ErrInternalServer(err.Error())
	}

	//check matches
	userTargetLikeSwipes, err := u.repo.FindLike(targetUser)
	if err != nil {
		return view.ErrInternalServer(err.Error())
	}

	match := false
	for i := 0; i < len(*userTargetLikeSwipes); i++ {
		if (*userTargetLikeSwipes)[i].TargetId == userId {
			match = true
			break
		}
	}

	resp := view.UserSwipeResponse{
		UserTarget: targetUser,
		Match:      match,
	}

	return view.SuccessAction(resp)
}

func (u *UserSwipeServices) FavouriteAction(targetUser string) *view.Response {
	_targetUser = targetUser
	_, err := u.userRepo.FindById(targetUser)
	if err != nil {
		return view.ErrBadRequest("user not found")
	}

	valid, errMessage := u.checkPrivilege(enums.FavouriteSwipeAction)
	if !valid {
		return view.ErrBadRequest(errMessage)
	}

	userId := _context.GetString("USER_ID")
	if userId == targetUser {
		return view.ErrBadRequest("Loving yourself is good, but not in here")
	}

	userSwipe := model.UserSwipe{
		UserId:   userId,
		TargetId: targetUser,
		Action:   enums.FavouriteSwipeAction,
	}

	err = u.repo.Create(&userSwipe)
	if err != nil {
		return view.ErrInternalServer(err.Error())
	}

	resp := view.UserSwipeResponse{
		UserTarget: targetUser,
		Match:      false,
	}

	return view.SuccessAction(resp)
}

func (c *UserSwipeServices) checkPrivilege(action string) (bool, string) {
	userId := _context.GetString("USER_ID")
	userMembership, err := c.userMembershipRepo.FindByUserId(userId)
	if err != nil {
		return false, "User dont have membership plan !"
	}

	memberPrivileges, err := c.membershipPrivilegeRepo.GetMembershipPrivilegeByMemberId(userMembership.MembershipId)
	if err != nil {
		return false, err.Error()
	}

	var userPrivileges []model.Privilege
	for _, memberPrivilege := range *memberPrivileges {
		userPrivileges = append(userPrivileges, memberPrivilege.Privilege)
	}

	userSwipesToday, err := c.repo.FindSwipesToday(userId)
	if err != nil {
		return false, err.Error()
	}

	for _, swiped := range *userSwipesToday {
		if swiped.TargetId == _targetUser {
			return false, "Cant swipe same user in a day"
		}
	}

	userLikesToday, err := c.repo.FindLikeToday(userId)
	if err != nil {
		return false, err.Error()
	}

	userFavsToday, err := c.repo.FindFavouriteToday(userId)
	if err != nil {
		return false, err.Error()
	}

	// check swipe privilege
	if action == enums.PassSwipeAction || action == enums.LikeSwipeAction {
		if userMembership.Membership.MembershipName == enums.FreePlanMembership {
			freePlanIdx := c.searchPrivilegeIdxByName(&userPrivileges, enums.LimitedSwipesPrivilege)
			freePlanPrivilege := userPrivileges[freePlanIdx]

			swipeLimit := freePlanPrivilege.ActionLimit
			if len(*userSwipesToday) >= int(swipeLimit) {
				return false, fmt.Sprintf("Swipe is limited %v times per day for %s membership", swipeLimit, enums.FreePlanMembership)
			}
		}
	}

	// check swipe privilege
	if action == enums.LikeSwipeAction {
		if userMembership.Membership.MembershipName != enums.GoldPlanMembership &&
			userMembership.Membership.MembershipName != enums.DiamondPlanMembership {
			const likeActionLimit = 100
			if len(*userLikesToday) >= int(likeActionLimit) {
				return false, fmt.Sprintf("Swipe is limited %v times per day for %s membership", likeActionLimit, userMembership.Membership.MembershipName)
			}
		}
	}

	// check fav privilege
	if action == enums.FavouriteSwipeAction {
		if userMembership.Membership.MembershipName == enums.FreePlanMembership {
			return false, "Favourite Swipe action only for paid memberships"
		}

		if userMembership.Membership.MembershipName != enums.DiamondPlanMembership {
			const favActionLimit = 10
			if len(*userFavsToday) >= int(favActionLimit) {
				return false, fmt.Sprintf("Favourite Swipe is limited %v times per day for %s membership", favActionLimit, userMembership.Membership.MembershipName)
			}
		}
	}

	return true, ""
}

func (c *UserSwipeServices) searchPrivilegeIdxByName(privileges *[]model.Privilege, name string) int {
	return helper.FindIndex(*privileges, func(value interface{}) bool {
		return value.(model.Privilege).PrivilegeName == name
	})
}
