package middleware

import (
	"minder/src/server/controller"
	"minder/src/server/view"

	"github.com/gin-gonic/gin"
)

const ADMIN_EMAIL = "admin@minder.com"

func (m *Middleware) AdminOnly(next gin.HandlerFunc) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userId := ctx.GetString("USER_ID")
		user := m.userSvc.FindById(userId)
		userDetail := user.Data.(*view.UserProfileResponse)

		if userDetail.Email != ADMIN_EMAIL {
			controller.WriteErrorJsonResponse(ctx, view.ErrUnauthorized())
			return
		}

		next(ctx)
	}
}
