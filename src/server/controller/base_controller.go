package controller

import (
	"minder/src/server/view"

	"github.com/gin-gonic/gin"
)

func WriteJsonResponse(c *gin.Context, payload *view.Response) {
	c.JSON(payload.Status, payload)
}

func WriteErrorJsonResponse(c *gin.Context, payload *view.Response) {
	c.AbortWithStatusJSON(payload.Status, payload)
}
