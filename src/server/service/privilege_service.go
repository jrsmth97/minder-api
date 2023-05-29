package service

import (
	"database/sql"
	"minder/src/server/param"
	"minder/src/server/repository"
	"minder/src/server/view"

	"github.com/gin-gonic/gin"
)

type PrivilegeService struct {
	repo                    repository.PrivilegeRepo
	membershipPrivilegeRepo repository.MembershipPrivilegeRepo
}

func NewPrivilegeServices(
	repo repository.PrivilegeRepo,
	membershipPrivilegeRepo repository.MembershipPrivilegeRepo,
) *PrivilegeService {
	return &PrivilegeService{
		repo:                    repo,
		membershipPrivilegeRepo: membershipPrivilegeRepo,
	}
}

func (p *PrivilegeService) Preparation(c *gin.Context) {
	_context = c
}

func (p *PrivilegeService) GetPrivileges() *view.Response {
	privileges, err := p.repo.GetPrivileges()
	if err != nil {
		return view.ErrInternalServer(err.Error())
	}

	if err != nil {
		if err == sql.ErrNoRows {
			return view.ErrNotFound()
		}
		return view.ErrInternalServer(err.Error())
	}

	return view.SuccessFind(view.NewPrivilegeGetAllResponse(privileges))
}

func (p *PrivilegeService) GetPrivilegeById(privilegeId string) *view.Response {
	privilege, err := p.repo.GetPrivilegeById(privilegeId)
	if err != nil {
		return view.ErrNotFound(err.Error())
	}

	return view.SuccessFind(privilege)
}

func (p *PrivilegeService) CreatePrivilege(req *param.PrivilegeCreate) *view.Response {
	privilege := req.ParseToModel()

	err := p.repo.CreatePrivilege(privilege)
	if err != nil {
		return view.ErrInternalServer(err.Error())
	}

	return view.SuccessCreated(view.NewPrivilegeCreateResponse(privilege))
}

func (p *PrivilegeService) UpdatePrivilege(privilegeId string, req *param.PrivilegeUpdate) *view.Response {
	privilege, errPrivilege := p.repo.GetPrivilegeById(privilegeId)
	if errPrivilege != nil {
		return view.ErrBadRequest("privilege doesn't exists")
	}

	privilege.PrivilegeName = req.PrivilegeName
	privilege.ActionLimit = req.ActionLimit
	privilege.ActionPath = req.ActionPath
	privilege.LimitInDay = req.LimitInDay
	err := p.repo.UpdatePrivilege(privilegeId, privilege)
	if err != nil {
		return view.ErrInternalServer(err.Error())
	}

	return view.SuccessUpdated(view.NewPrivilegeUpdateResponse(privilege))
}

func (p *PrivilegeService) DeletePrivilege(privilegeId string) *view.Response {
	privilege, err := p.repo.GetPrivilegeById(privilegeId)
	if err != nil {
		return view.ErrBadRequest("privilege doesn't exists")
	}

	err = p.membershipPrivilegeRepo.DeleteMembershipPrivilegeByPrivilegeId(privilegeId)
	if err != nil {
		return view.ErrInternalServer(err.Error())
	}

	err = p.repo.DeletePrivilege(privilegeId)
	if err != nil {
		return view.ErrInternalServer(err.Error())
	}

	return view.SuccessDeleted(view.NewPrivilegeUpdateResponse(privilege))
}
