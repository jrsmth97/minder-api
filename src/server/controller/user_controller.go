package controller

import (
	"fmt"
	"minder/src/server/param"
	"minder/src/server/service"
	"minder/src/server/view"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	svc *service.UserServices
}

func NewUserHandler(svc *service.UserServices) *UserHandler {
	return &UserHandler{
		svc: svc,
	}
}

func (a *UserHandler) Profile(c *gin.Context) {
	fmt.Println("Log from ", c.GetString("USER_EMAIL"))
	resp := a.svc.FindById(c.GetString("USER_ID"))

	WriteJsonResponse(c, resp)
}

func (a *UserHandler) Explore(c *gin.Context) {
	fmt.Println("Log from ", c.GetString("USER_EMAIL"))

	a.svc.Preparation(c)
	resp := a.svc.Explore()

	WriteJsonResponse(c, resp)
}

func (u *UserHandler) FindByEmail(c *gin.Context) {
	fmt.Println("Log from ", c.GetString("USER_EMAIL"))
	email, isExist := c.Params.Get("email")
	if !isExist {
		WriteJsonResponse(c, view.ErrBadRequest("email not found in params"))
		return
	}

	resp := u.svc.FindByEmail(email)

	WriteJsonResponse(c, resp)
}

func (u *UserHandler) Update(c *gin.Context) {
	userId := c.GetString("USER_ID")

	var req param.UserUpdate

	fmt.Println((req))
	err := c.ShouldBindJSON(&req)
	if err != nil {
		WriteJsonResponse(c, view.ErrBadRequest(err.Error()))
		return
	}

	err = param.Validate(req)
	if err != nil {
		user := view.ErrBadRequest(err.Error())
		WriteJsonResponse(c, user)
		return
	}

	resp := u.svc.Update(userId, &req)

	WriteJsonResponse(c, resp)
}

func (m *UserHandler) DeleteUser(c *gin.Context) {
	userId := c.GetString("USER_ID")
	resp := m.svc.Delete(userId)

	WriteJsonResponse(c, resp)
}

func (m *UserHandler) CountUsers(c *gin.Context) {
	resp := m.svc.CountUsers()

	WriteJsonResponse(c, resp)
}
