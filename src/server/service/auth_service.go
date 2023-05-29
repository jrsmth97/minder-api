package service

import (
	"database/sql"
	"log"
	"minder/src/helper"
	"minder/src/server/enums"
	"minder/src/server/model"
	"minder/src/server/param"
	"minder/src/server/repository"
	"minder/src/server/view"
	"time"
)

type AuthServices struct {
	repo               repository.UserRepo
	locationRepo       repository.LocationRepo
	membershipRepo     repository.MembershipRepo
	userMembershipRepo repository.UserMembershipRepo
}

func NewAuthServices(repo repository.UserRepo, locationRepo repository.LocationRepo, membershipRepo repository.MembershipRepo, userMembershipRepo repository.UserMembershipRepo) *AuthServices {
	return &AuthServices{
		repo:               repo,
		locationRepo:       locationRepo,
		membershipRepo:     membershipRepo,
		userMembershipRepo: userMembershipRepo,
	}
}

func (u *AuthServices) Register(req *param.AuthRegister) *view.Response {
	user := req.ParseToModelRegister()

	_, err := u.locationRepo.GetLocationById(user.LocationId)
	if err != nil {
		return view.ErrBadRequest("Location not found")
	}

	emailExist, _ := u.repo.FindByEmail(user.Email)
	if emailExist != nil {
		return view.ErrBadRequest("Email exists")
	}

	hash, err := helper.GeneratePassword(user.Password)
	if err != nil {
		log.Printf("get error when try to generate password %v\n", "")
		return view.ErrInternalServer(err.Error())
	}

	user.Password = hash

	err = u.repo.Create(user)
	if err != nil {
		log.Printf("get error register user with error %v\n", "")
		return view.ErrInternalServer(err.Error())
	}

	// add free plan membership as default
	freeMembershipPlan, err := u.membershipRepo.GetMembershipByName(enums.FreePlanMembership)
	if err != nil {
		return view.ErrInternalServer(err.Error())
	}

	freePlanUserMembership := model.UserMembership{
		UserId:       user.ID.String(),
		MembershipId: freeMembershipPlan.ID.String(),
		ValidUntil:   time.Now().Add((time.Hour * 24 * 30) * 999),
	}

	err = u.userMembershipRepo.Create(&freePlanUserMembership)
	if err != nil {
		return view.ErrInternalServer(err.Error())
	}

	return view.SuccessRegister(view.NewUserProfileResponse(user, false))
}

func (u *AuthServices) Login(req *param.UserLogin) *view.Response {
	user, err := u.repo.FindByEmail(req.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return view.ErrUnauthorized()
		}
		return view.ErrUnauthorized()
	}

	err = helper.ValidatePassword(user.Password, req.Password)
	if err != nil {
		return view.ErrUnauthorized()
	}

	token := helper.Token{
		UserId: user.ID.String(),
		Email:  user.Email,
	}

	tokString, err := helper.CreateToken(&token)
	if err != nil {
		return view.ErrInternalServer(err.Error())
	}

	isVerified := false
	userMembership, _ := u.userMembershipRepo.FindByUserId(user.ID.String())

	membership := userMembership.Membership.MembershipName
	if membership == enums.SilverPlanMembership ||
		membership == enums.GoldPlanMembership ||
		membership == enums.DiamondPlanMembership {
		isVerified = true
	}

	return view.SuccessLogin(view.NewLoginResponse(user, tokString, isVerified))
}
