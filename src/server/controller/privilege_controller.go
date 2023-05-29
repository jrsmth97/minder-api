package controller

import (
	"minder/src/server/param"
	"minder/src/server/service"
	"minder/src/server/view"

	"github.com/gin-gonic/gin"
)

type PrivilegeHandler struct {
	svc *service.PrivilegeService
}

func NewPrivilegeHandler(svc *service.PrivilegeService) *PrivilegeHandler {
	return &PrivilegeHandler{
		svc: svc,
	}
}

func (p *PrivilegeHandler) GetPrivileges(c *gin.Context) {
	resp := p.svc.GetPrivileges()

	WriteJsonResponse(c, resp)
}

func (m *PrivilegeHandler) CreatePrivilege(c *gin.Context) {
	var req param.PrivilegeCreate
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
	resp := m.svc.CreatePrivilege(&req)

	WriteJsonResponse(c, resp)
}

func (m *PrivilegeHandler) UpdatePrivilege(c *gin.Context) {
	privilegeId, isExist := c.Params.Get("privilegeId")
	if !isExist {
		WriteJsonResponse(c, view.ErrBadRequest("privilege not found"))
		return
	}

	var req param.PrivilegeUpdate
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

	resp := m.svc.UpdatePrivilege(privilegeId, &req)

	WriteJsonResponse(c, resp)
}

func (m *PrivilegeHandler) DeletePrivilege(c *gin.Context) {
	privilegeId, isExist := c.Params.Get("privilegeId")
	if !isExist {
		WriteJsonResponse(c, view.ErrBadRequest("privilege not found"))
		return
	}

	resp := m.svc.DeletePrivilege(privilegeId)

	WriteJsonResponse(c, resp)
}

func (p *PrivilegeHandler) GetPrivilegeById(c *gin.Context) {
	privilegeId, isExist := c.Params.Get("privilegeId")
	if !isExist {
		WriteJsonResponse(c, view.ErrBadRequest("privilege id not found in params"))
		return
	}

	resp := p.svc.GetPrivilegeById(privilegeId)
	WriteJsonResponse(c, resp)
}
