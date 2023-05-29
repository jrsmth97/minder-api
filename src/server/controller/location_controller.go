package controller

import (
	"minder/src/server/param"
	"minder/src/server/service"
	"minder/src/server/view"

	"github.com/gin-gonic/gin"
)

type LocationHandler struct {
	svc *service.LocationService
}

func NewLocationHandler(svc *service.LocationService) *LocationHandler {
	return &LocationHandler{
		svc: svc,
	}
}

func (p *LocationHandler) GetLocations(c *gin.Context) {
	resp := p.svc.GetLocations()

	WriteJsonResponse(c, resp)
}

func (m *LocationHandler) CreateLocation(c *gin.Context) {
	var req param.LocationCreate
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
	resp := m.svc.CreateLocation(&req)

	WriteJsonResponse(c, resp)
}

func (m *LocationHandler) UpdateLocation(c *gin.Context) {
	locationId, isExist := c.Params.Get("locationId")
	if !isExist {
		WriteJsonResponse(c, view.ErrBadRequest("location not found"))
		return
	}

	var req param.LocationUpdate
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

	resp := m.svc.UpdateLocation(locationId, &req)

	WriteJsonResponse(c, resp)
}

func (m *LocationHandler) DeleteLocation(c *gin.Context) {
	locationId, isExist := c.Params.Get("locationId")
	if !isExist {
		WriteJsonResponse(c, view.ErrBadRequest("location not found"))
		return
	}

	resp := m.svc.DeleteLocation(locationId)

	WriteJsonResponse(c, resp)
}

func (p *LocationHandler) GetLocationById(c *gin.Context) {
	locationId, isExist := c.Params.Get("locationId")
	if !isExist {
		WriteJsonResponse(c, view.ErrBadRequest("location id not found in params"))
		return
	}

	resp := p.svc.GetLocationById(locationId)
	WriteJsonResponse(c, resp)
}
