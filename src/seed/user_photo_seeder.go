package seed

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"minder/src/helper"
	"minder/src/server/enums"
	"minder/src/server/model"
	"os"
	"path/filepath"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func SeedUserPhoto(db *gorm.DB) {
	var users []model.User
	db.Find(&users)

	for _, user := range users {
		var userPhotoExist []model.UserPhoto
		err := db.Where("user_id=?", user.ID.String()).Find(&userPhotoExist).Error
		if len(userPhotoExist) > 0 {
			continue
		}

		dir, err := os.Getwd()
		if err != nil {
			panic(err)
		}

		staticSrcFile := "/src/static/images/"
		randomNumber := helper.RandomNumberV2(2)
		var fileName string
		switch user.Gender {
		case enums.MaleGender:
			fileName = staticSrcFile + "male-" + fmt.Sprint(randomNumber) + ".jpg"
		case enums.FemaleGender:
			fileName = staticSrcFile + "female-" + fmt.Sprint(randomNumber) + ".jpg"
		default:
			fileName = staticSrcFile + "binary.jpg"
		}

		srcFile := dir + fileName
		pathFile := moveFile(srcFile)

		userPhoto := model.UserPhoto{
			UserId: user.ID.String(),
			Asset:  pathFile,
		}

		err = db.Debug().Model(&model.UserPhoto{}).Create(&userPhoto).Error
		if err != nil {
			log.Fatalf("cannot seed user photo table: %v", err)
		}
	}
}

func moveFile(srcFile string) string {
	uploadPath := "uploads/media/"
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	readFileBytes, err := os.ReadFile(srcFile)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	readFile := bytes.NewReader(readFileBytes)

	dirPath := filepath.Join(dir, uploadPath)
	err = os.MkdirAll(dirPath, os.ModePerm)
	if err != nil {
		panic(err)
	}

	fileName := uuid.NewString() + ".jpg"
	fileLocation := filepath.Join(dirPath, fileName)
	targetFile, err := os.OpenFile(fileLocation, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer targetFile.Close()

	if _, err := io.Copy(targetFile, readFile); err != nil {
		fmt.Println(err)
		panic(err)
	}

	return uploadPath + fileName
}
