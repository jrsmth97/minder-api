package view

import (
	"minder/src/server/enums"
	"minder/src/server/model"
	"strings"
	"time"
)

type UserFindResponse struct {
	Id           string         `json:"id"`
	Name         string         `json:"name"`
	Gender       uint8          `json:"gender"`
	GenderName   string         `json:"gender_name"`
	LocationId   string         `json:"location_id"`
	LocationName string         `json:"location_name"`
	BirthDate    string         `json:"birth_date"`
	IsVerified   bool           `json:"is_verified"`
	Age          uint8          `json:"age"`
	Photos       []Photo        `json:"photos"`
	Interests    []UserInterest `json:"interests"`
}

type UserProfileResponse struct {
	Id           string         `json:"id"`
	Name         string         `json:"name"`
	Gender       uint8          `json:"gender"`
	LocationId   string         `json:"location_id"`
	LocationName string         `json:"location_name"`
	BirthDate    string         `json:"birth_date"`
	Email        string         `json:"email"`
	Phone        string         `json:"phone"`
	IsVerified   bool           `json:"is_verified"`
	Photos       []Photo        `json:"photos"`
	Interests    []UserInterest `json:"interests"`
}

type Photo struct {
	Id      string `json:"id"`
	MediaId string `json:"media_id"`
}

type UserInterest struct {
	Id           string `json:"id"`
	InterestId   string `json:"interest_id"`
	InterestName string `json:"interest_name"`
}

type LoginResponse struct {
	User        UserProfileResponse `json:"user"`
	AccessToken string              `json:"access_token"`
}

func NewUserFindResponse(users *[]model.User, isVerifiedUsers map[string]bool) []UserFindResponse {
	var usersFind []UserFindResponse
	for _, user := range *users {
		usersFind = append(usersFind, *ParseModelToUserFind(&user, isVerifiedUsers))
	}
	return usersFind
}

func NewUserProfileResponse(user *model.User, isVerified bool) *UserProfileResponse {
	return &UserProfileResponse{
		Id:           user.ID.String(),
		Name:         user.Name,
		Gender:       user.Gender,
		LocationId:   user.LocationId,
		LocationName: user.Location.LocationName,
		BirthDate:    user.BirthDate,
		Email:        user.Email,
		Phone:        user.Phone,
		IsVerified:   isVerified,
		Photos:       ParseModelToPhotos(user.Photos),
		Interests:    ParseModelToInterests(user.Interests),
	}
}

func NewLoginResponse(user *model.User, token string, isVerified bool) *LoginResponse {
	userProfile := NewUserProfileResponse(user, isVerified)

	return &LoginResponse{
		User:        *userProfile,
		AccessToken: token,
	}
}

func ParseModelToUserFind(user *model.User, isVerifiedUsers map[string]bool) *UserFindResponse {
	var genderName string

	switch user.Gender {
	case enums.MaleGender:
		genderName = enums.MaleGenderName
	case enums.FemaleGender:
		genderName = enums.FemaleGenderName
	default:
		genderName = enums.BinaryGenderName
	}

	parseDateFormat := "2006-01-02"
	parsedBirthDate, _ := time.Parse(parseDateFormat, user.BirthDate)
	difference := time.Since(parsedBirthDate)
	age := difference.Hours() / 24 / 365

	isVerified := false
	if isVerifiedUsers != nil {
		isVerified = isVerifiedUsers[user.ID.String()]
	}

	return &UserFindResponse{
		Id:           user.ID.String(),
		Name:         user.Name,
		Gender:       user.Gender,
		GenderName:   genderName,
		LocationId:   user.LocationId,
		LocationName: user.Location.LocationName,
		BirthDate:    user.BirthDate,
		Age:          uint8(age),
		IsVerified:   isVerified,
		Photos:       ParseModelToPhotos(user.Photos),
		Interests:    ParseModelToInterests(user.Interests),
	}
}

func ParseModelToPhotos(photos []model.UserPhoto) []Photo {
	var userPhotos []Photo
	for _, photo := range photos {
		splitPath := strings.Split(photo.Asset, "/")
		mediaId := strings.Split(splitPath[len(splitPath)-1], ".")[0]
		photo := Photo{
			Id:      photo.ID.String(),
			MediaId: mediaId,
		}

		userPhotos = append(userPhotos, photo)
	}

	return userPhotos
}

func ParseModelToInterests(interests []model.UserInterest) []UserInterest {
	var userInterests []UserInterest
	for _, interest := range interests {
		userInterest := UserInterest{
			Id:           interest.ID.String(),
			InterestId:   interest.InterestId,
			InterestName: interest.Interest.InterestName,
		}

		userInterests = append(userInterests, userInterest)
	}

	return userInterests
}
