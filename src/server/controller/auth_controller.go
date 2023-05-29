package controller

import (
	"fmt"
	"minder/src/server/param"
	"minder/src/server/service"
	"minder/src/server/view"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	svc *service.AuthServices
}

func NewAuthHandler(svc *service.AuthServices) *AuthHandler {
	return &AuthHandler{
		svc: svc,
	}
}

func (a *AuthHandler) Register(c *gin.Context) {
	var req param.AuthRegister
	err := c.ShouldBindJSON(&req)
	if err != nil {
		resp := view.ErrBadRequest(err.Error())
		WriteJsonResponse(c, resp)
		return
	}

	err = param.Validate(req)
	if err != nil {
		resp := view.ErrBadRequest(err.Error())
		WriteJsonResponse(c, resp)
		return
	}

	resp := a.svc.Register(&req)
	fmt.Println(resp)
	WriteJsonResponse(c, resp)
}

func (a *AuthHandler) Login(c *gin.Context) {
	var req param.UserLogin
	err := c.ShouldBindJSON(&req)
	if err != nil {
		resp := view.ErrBadRequest(err.Error())
		WriteJsonResponse(c, resp)
		return
	}

	err = param.Validate(req)
	if err != nil {
		resp := view.ErrBadRequest(err.Error())
		WriteJsonResponse(c, resp)
		return
	}

	resp := a.svc.Login(&req)
	WriteJsonResponse(c, resp)
}
