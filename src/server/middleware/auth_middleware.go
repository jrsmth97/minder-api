package middleware

import (
	"minder/src/helper"
	"minder/src/server/controller"
	"minder/src/server/view"
	"strings"

	"github.com/gin-gonic/gin"
)

func (m *Middleware) Auth(c *gin.Context) {
	bearerToken := c.GetHeader("Authorization")
	tokenArr := strings.Split(bearerToken, "Bearer ")

	if len(tokenArr) != 2 {
		c.Set("ERROR", "no token")
		controller.WriteErrorJsonResponse(c, view.ErrUnauthorized())
		return
	}

	token, err := helper.VerifyToken(tokenArr[1])
	if err != nil {
		c.Set("ERROR", err.Error())
		controller.WriteErrorJsonResponse(c, view.ErrUnauthorized())
		return
	}

	user := m.userSvc.FindById(token.UserId)
	userDetail := user.Data.(*view.UserProfileResponse)
	if userDetail == nil {
		c.Set("ERROR", err.Error())
		controller.WriteErrorJsonResponse(c, view.ErrUnauthorized())
		return
	}

	c.Set("USER_ID", userDetail.Id)
	c.Set("USER_EMAIL", userDetail.Email)
	c.Set("USER_NAME", userDetail.Name)

	c.Next()

}
