package service

import (
	"database/sql"
	"fmt"
	"minder/src/helper"
	"minder/src/server/enums"
	"minder/src/server/model"
	"minder/src/server/param"
	"minder/src/server/repository"
	"minder/src/server/view"

	"github.com/gin-gonic/gin"
)

type UserServices struct {
	repo               repository.UserRepo
	locationRepo       repository.LocationRepo
	interestRepo       repository.InterestRepo
	userSwipeRepo      repository.UserSwipeRepo
	userMembershipRepo repository.UserMembershipRepo
}

func NewUserServices(repo repository.UserRepo, locationRepo repository.LocationRepo, interestRepo repository.InterestRepo, userSwipeRepo repository.UserSwipeRepo, userMembershipRepo repository.UserMembershipRepo) *UserServices {
	return &UserServices{
		repo:               repo,
		locationRepo:       locationRepo,
		interestRepo:       interestRepo,
		userSwipeRepo:      userSwipeRepo,
		userMembershipRepo: userMembershipRepo,
	}
}

func (s *UserServices) Preparation(c *gin.Context) {
	_context = c
}

func (s *UserServices) Explore() *view.Response {
	users, err := s.repo.Explore()

	maxExploreResults := 17
	userId := _context.GetString("USER_ID")
	userSwipesToday, errFindSwipes := s.userSwipeRepo.FindSwipesToday(userId)
	if errFindSwipes != nil {
		return view.ErrInternalServer(errFindSwipes.Error())
	}

	var userSwipedIds []string
	for _, userSwipe := range *userSwipesToday {
		userSwipedIds = append(userSwipedIds, userSwipe.TargetId)
	}

	var exploreUsers []model.User
	randomUserIndexRange := helper.MakeNumberRange(0, (len(*users) - 1))
	randomUserIndexRange = helper.ArrayIntShuffle(randomUserIndexRange)

	var usersIsVerified = make(map[string]bool)
	for _, index := range randomUserIndexRange {
		user := (*users)[index]
		if helper.StringContains(userSwipedIds, user.ID.String()) {
			continue
		}

		if user.ID.String() == userId {
			continue
		}

		usersIsVerified[user.ID.String()] = true
		userMembership, err := s.userMembershipRepo.FindByUserId(user.ID.String())
		if err != nil {
			continue
		}

		if userMembership.Membership.MembershipName == enums.FreePlanMembership {
			usersIsVerified[user.ID.String()] = false
		}

		exploreUsers = append(exploreUsers, user)
		if len(exploreUsers) == maxExploreResults {
			break
		}
	}

	if err != nil {
		if err == sql.ErrNoRows {
			return view.ErrNotFound()
		}
		return view.ErrInternalServer(err.Error())
	}

	return view.SuccessFind(view.NewUserFindResponse(&exploreUsers, usersIsVerified))
}

func (s *UserServices) FindById(id string) *view.Response {
	user, err := s.repo.FindById(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return view.ErrNotFound()
		}
		return view.ErrNotFound()
	}

	isVerified := false
	userMembership, _ := s.userMembershipRepo.FindByUserId(user.ID.String())

	membership := userMembership.Membership.MembershipName
	if membership == enums.SilverPlanMembership ||
		membership == enums.GoldPlanMembership ||
		membership == enums.DiamondPlanMembership {
		isVerified = true
	}

	return view.SuccessFind(view.NewUserProfileResponse(user, isVerified))
}

func (s *UserServices) FindByEmail(email string) *view.Response {
	user, err := s.repo.FindByEmail(email)
	if err != nil {
		if err == sql.ErrNoRows {
			return view.ErrNotFound()
		}
		return view.ErrNotFound()
	}
	fmt.Printf("userid => %v", user.ID)
	return view.SuccessFind(user)
}

func (s *UserServices) Update(id string, req *param.UserUpdate) *view.Response {
	user := req.ParseToModelUpdate()

	_, err := s.locationRepo.GetLocationById(user.LocationId)
	if err != nil {
		return view.ErrBadRequest("location doesn't exists")
	}

	_, err = s.repo.FindById(id)
	if err != nil {
		return view.ErrBadRequest("user doesn't exists")
	}

	for _, interest := range user.Interests {
		_, err = s.interestRepo.GetInterestById(interest.InterestId)
		if err != nil {
			return view.ErrBadRequest(fmt.Sprintf("interest %s doesn't exists", interest.InterestId))
		}
	}

	err = s.repo.Update(id, user)
	if err != nil {
		return view.ErrInternalServer(err.Error())
	}

	updatedUser := s.FindById(id)
	return view.SuccessUpdated(updatedUser)
}

func (s *UserServices) Delete(id string) *view.Response {
	user, err := s.repo.FindById(id)
	if err != nil {
		return view.ErrBadRequest("user doesn't exists")
	}

	err = s.repo.Delete(user.ID.String())
	if err != nil {
		return view.ErrInternalServer(err.Error())
	}

	return view.SuccessDeleted(view.NewUserProfileResponse(user, false))
}

func (s *UserServices) CountUsers() *view.Response {
	count, err := s.repo.CountUsers()
	if err != nil {
		return view.ErrInternalServer(err.Error())
	}

	return view.SuccessFind(count)
}
