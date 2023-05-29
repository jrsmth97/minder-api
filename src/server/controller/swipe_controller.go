package controller

import (
	"fmt"
	"minder/src/server/service"
	"minder/src/server/view"

	"github.com/gin-gonic/gin"
)

type SwipeHandler struct {
	svc *service.UserSwipeServices
}

func NewSwipeHandler(svc *service.UserSwipeServices) *SwipeHandler {
	return &SwipeHandler{
		svc: svc,
	}
}

func (a *SwipeHandler) Like(c *gin.Context) {
	fmt.Println("Log from ", c.GetString("USER_EMAIL"))
	targetId, isExist := c.Params.Get("targetId")
	if !isExist {
		WriteJsonResponse(c, view.ErrBadRequest("target user not found"))
		return
	}

	a.svc.Preparation(c)
	resp := a.svc.LikeAction(targetId)

	WriteJsonResponse(c, resp)
}

func (a *SwipeHandler) Pass(c *gin.Context) {
	fmt.Println("Log from ", c.GetString("USER_EMAIL"))
	targetId, isExist := c.Params.Get("targetId")
	if !isExist {
		WriteJsonResponse(c, view.ErrBadRequest("target user not found"))
		return
	}

	a.svc.Preparation(c)
	resp := a.svc.PassAction(targetId)

	WriteJsonResponse(c, resp)
}

func (a *SwipeHandler) Favourite(c *gin.Context) {
	fmt.Println("Log from ", c.GetString("USER_EMAIL"))
	targetId, isExist := c.Params.Get("targetId")
	if !isExist {
		WriteJsonResponse(c, view.ErrBadRequest("target user not found"))
		return
	}

	a.svc.Preparation(c)
	resp := a.svc.FavouriteAction(targetId)

	WriteJsonResponse(c, resp)
}
