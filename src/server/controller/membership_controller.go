package controller

import (
	"minder/src/server/param"
	"minder/src/server/service"
	"minder/src/server/view"

	"github.com/gin-gonic/gin"
)

type MembershipHandler struct {
	svc *service.MembershipService
}

func NewMembershipHandler(svc *service.MembershipService) *MembershipHandler {
	return &MembershipHandler{
		svc: svc,
	}
}

func (p *MembershipHandler) GetMemberships(c *gin.Context) {
	resp := p.svc.GetMemberships()

	WriteJsonResponse(c, resp)
}

func (m *MembershipHandler) CreateMembership(c *gin.Context) {
	var req param.MembershipCreate
	err := c.ShouldBindJSON(&req)
	if err != nil {
		WriteJsonResponse(c, view.ErrBadRequest(err.Error()))
		return
	}

	err = param.Validate(req)
	if err != nil {
		resp := view.ErrBadRequest(err.Error())
		WriteJsonResponse(c, resp)
		return
	}

	m.svc.Preparation(c)
	resp := m.svc.CreateMembership(&req)

	WriteJsonResponse(c, resp)
}

func (m *MembershipHandler) UpdateMembership(c *gin.Context) {
	membershipId, isExist := c.Params.Get("membershipId")
	if !isExist {
		WriteJsonResponse(c, view.ErrBadRequest("membership not found"))
		return
	}

	var req param.MembershipUpdate
	err := c.ShouldBindJSON(&req)
	if err != nil {
		WriteJsonResponse(c, view.ErrBadRequest(err.Error()))
		return
	}

	err = param.Validate(req)
	if err != nil {
		resp := view.ErrBadRequest(err.Error())
		WriteJsonResponse(c, resp)
		return
	}

	resp := m.svc.UpdateMembership(membershipId, &req)

	WriteJsonResponse(c, resp)
}

func (m *MembershipHandler) DeleteMembership(c *gin.Context) {
	membershipId, isExist := c.Params.Get("membershipId")
	if !isExist {
		WriteJsonResponse(c, view.ErrBadRequest("membership not found"))
		return
	}

	resp := m.svc.DeleteMembership(membershipId)

	WriteJsonResponse(c, resp)
}

func (p *MembershipHandler) GetMembershipById(c *gin.Context) {
	membershipId, isExist := c.Params.Get("membershipId")
	if !isExist {
		WriteJsonResponse(c, view.ErrBadRequest("membership id not found in params"))
		return
	}

	resp := p.svc.GetMembershipById(membershipId)
	WriteJsonResponse(c, resp)
}
