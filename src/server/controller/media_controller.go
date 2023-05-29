package controller

import (
	"fmt"
	"minder/src/server/service"
	"minder/src/server/view"
	"os"

	"github.com/gin-gonic/gin"
)

type MediaHandler struct {
	svc *service.MediaService
}

func NewMediaHandler(svc *service.MediaService) *MediaHandler {
	return &MediaHandler{
		svc: svc,
	}
}

func (p *MediaHandler) UploadMedia(c *gin.Context) {
	p.svc.Preparation(c)
	path := p.svc.UploadFile()

	WriteJsonResponse(c, path)
}

func (p *MediaHandler) GetMedia(c *gin.Context) {
	mediaId, isExist := c.Params.Get("mediaId")
	if !isExist {
		WriteJsonResponse(c, view.ErrBadRequest("media not found"))
		return
	}

	dir, err := os.Getwd()
	if err != nil {
		WriteErrorJsonResponse(c, view.ErrInternalServer(err.Error()))
		return
	}

	mediaPath := dir + "/uploads/media/" + mediaId + ".jpg"
	_, err = os.Stat(mediaPath)
	if os.IsNotExist(err) {
		WriteErrorJsonResponse(c, view.ErrNotFound())
		return
	}

	fmt.Println("mediaPath => " + mediaPath)
	c.File(mediaPath)
}

func (m *MediaHandler) DeleteMedia(c *gin.Context) {
	mediaId, isExist := c.Params.Get("mediaId")
	if !isExist {
		WriteJsonResponse(c, view.ErrBadRequest("media not found"))
		return
	}

	dir, err := os.Getwd()
	if err != nil {
		WriteErrorJsonResponse(c, view.ErrInternalServer(err.Error()))
		return
	}

	mediaPath := dir + "/uploads/media/" + mediaId + ".jpg"
	_, err = os.Stat(mediaPath)
	if os.IsNotExist(err) {
		WriteErrorJsonResponse(c, view.ErrNotFound())
		return
	}

	err = os.Remove(mediaPath)
	if err != nil {
		WriteErrorJsonResponse(c, view.ErrInternalServer(err.Error()))
		return
	}

	WriteJsonResponse(c, view.SuccessDeleted(nil))
}
