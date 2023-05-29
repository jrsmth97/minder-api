package service

import (
	"fmt"
	"io"
	"io/ioutil"
	"minder/src/server/view"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type MediaService struct{}

func NewMediaServices() *MediaService {
	return &MediaService{}
}

func (p *MediaService) Preparation(c *gin.Context) {
	_context = c
}

func (p *MediaService) UploadFile() *view.Response {
	request := _context.Request
	// max upload 1 MB
	request.Body = http.MaxBytesReader(_context.Writer, request.Body, 1*1024*1024)
	file, _, err := request.FormFile("file")
	if err != nil {
		return view.ErrBadRequest(err.Error())
	}

	defer file.Close()
	fileByte, err := ioutil.ReadAll(file)
	mimeType := http.DetectContentType(fileByte)

	if mimeType != "image/jpeg" && mimeType != "image/png" {
		return view.ErrBadRequest("file type not supported (png / jpg only) !")
	}

	uploadPath := "uploads/media/"
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return view.ErrInternalServer(err.Error())
	}

	dirPath := filepath.Join(dir, uploadPath)
	err = os.MkdirAll(dirPath, os.ModePerm)
	if err != nil {
		return view.ErrInternalServer(err.Error())
	}

	fileName := uuid.NewString() + ".jpg"
	fileLocation := filepath.Join(dirPath, fileName)
	targetFile, err := os.OpenFile(fileLocation, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return view.ErrInternalServer(err.Error())
	}
	defer targetFile.Close()

	if _, err := io.Copy(targetFile, file); err != nil {
		fmt.Println(err)
		return view.ErrInternalServer(err.Error())
	}

	return view.SuccessUpload(uploadPath + fileName)
}
