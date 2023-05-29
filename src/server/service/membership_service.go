package service

import (
	"database/sql"
	"minder/src/server/model"
	"minder/src/server/param"
	"minder/src/server/repository"
	"minder/src/server/view"

	"github.com/gin-gonic/gin"
)

type MembershipService struct {
	repo                    repository.MembershipRepo
	userRepo                repository.UserRepo
	privilegeRepo           repository.PrivilegeRepo
	membershipPrivilegeRepo repository.MembershipPrivilegeRepo
}

func NewMembershipServices(
	repo repository.MembershipRepo,
	userRepo repository.UserRepo,
	privilegeRepo repository.PrivilegeRepo,
	membershipPrivilegeRepo repository.MembershipPrivilegeRepo,
) *MembershipService {
	return &MembershipService{
		repo:                    repo,
		userRepo:                userRepo,
		privilegeRepo:           privilegeRepo,
		membershipPrivilegeRepo: membershipPrivilegeRepo,
	}
}

func (p *MembershipService) Preparation(c *gin.Context) {
	_context = c
}

func (p *MembershipService) GetMemberships() *view.Response {
	memberships, err := p.repo.GetMemberships()
	if err != nil {
		return view.ErrInternalServer(err.Error())
	}

	membershipPrivileges, err := p.membershipPrivilegeRepo.GetMembershipPrivileges()
	if err != nil {
		return view.ErrInternalServer(err.Error())
	}

	if err != nil {
		if err == sql.ErrNoRows {
			return view.ErrNotFound()
		}
		return view.ErrInternalServer(err.Error())
	}

	return view.SuccessFind(view.NewMembershipGetResponse(memberships, membershipPrivileges))
}

func (p *MembershipService) GetMembershipById(membershipId string) *view.Response {
	membership, err := p.repo.GetMembershipById(membershipId)
	if err != nil {
		return view.ErrInternalServer(err.Error())
	}

	return view.SuccessFind(membership)
}

func (p *MembershipService) CreateMembership(req *param.MembershipCreate) *view.Response {
	membership := req.ParseToModel()

	err := p.repo.CreateMembership(membership)
	if err != nil {
		return view.ErrInternalServer(err.Error())
	}

	var privileges []model.Privilege
	for _, privilegeId := range req.Privileges {
		privilege, err := p.privilegeRepo.GetPrivilegeById(privilegeId)
		if err != nil {
			return view.ErrBadRequest(err.Error())
		}

		membershipPrivilege := req.ParseToMembershipPrivilegeModel(membership.ID.String(), privilegeId)
		errCreateItem := p.membershipPrivilegeRepo.CreateMembershipPrivilege(membershipPrivilege)
		if errCreateItem != nil {
			return view.ErrInternalServer(errCreateItem.Error())
		}

		privileges = append(privileges, *privilege)
	}

	return view.SuccessCreated(view.NewMembershipCreateResponse(membership, &privileges))
}

func (p *MembershipService) UpdateMembership(membershipId string, req *param.MembershipUpdate) *view.Response {
	membership, errMembership := p.repo.GetMembershipById(membershipId)
	if errMembership != nil {
		return view.ErrBadRequest("membership doesn't exists")
	}

	membership.MembershipName = req.MembershipName
	membership.Description = req.Description
	membership.Price = req.Price
	membership.DurationInMonth = req.DurationInMonth
	err := p.repo.UpdateMembership(membershipId, membership)
	if err != nil {
		return view.ErrInternalServer(err.Error())
	}

	var memberPrivileges []model.MembershipPrivilege
	var privileges []model.Privilege
	for _, privilegeId := range req.Privileges {
		privilege, err := p.privilegeRepo.GetPrivilegeById(privilegeId)
		if err != nil {
			return view.ErrBadRequest(err.Error())
		}

		privileges = append(privileges, *privilege)
		memberPrivileges = append(memberPrivileges, model.MembershipPrivilege{
			MembershipId: membership.ID.String(),
			PrivilegeId:  privilege.ID.String(),
		})
	}

	errDeleteMemberPrivileges := p.membershipPrivilegeRepo.DeleteMembershipPrivilegeByMemberId(membershipId)
	if errDeleteMemberPrivileges != nil {
		return view.ErrInternalServer(err.Error())
	}

	for _, memberPrivilege := range memberPrivileges {
		errCreateMemberPrivilege := p.membershipPrivilegeRepo.CreateMembershipPrivilege(&memberPrivilege)
		if errCreateMemberPrivilege != nil {
			return view.ErrBadRequest(errCreateMemberPrivilege)
		}
	}

	return view.SuccessUpdated(view.NewMembershipUpdateResponse(membership, &privileges))
}

func (p *MembershipService) DeleteMembership(membershipId string) *view.Response {
	membership, err := p.repo.GetMembershipById(membershipId)
	if err != nil {
		return view.ErrBadRequest("membership doesn't exists")
	}

	err = p.membershipPrivilegeRepo.DeleteMembershipPrivilegeByMemberId(membershipId)
	if err != nil {
		return view.ErrInternalServer(err.Error())
	}

	err = p.repo.DeleteMembership(membershipId)
	if err != nil {
		return view.ErrInternalServer(err.Error())
	}

	return view.SuccessDeleted(view.NewMembershipUpdateResponse(membership, &[]model.Privilege{}))
}
