package test

import (
	"encoding/json"
	"minder/src/server/model"
	"minder/src/server/param"
	"minder/src/server/pkg/httpclient"
	"minder/src/server/view"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

var _USER_ID string
var _exploredUsers []string

func TestGetProfile(t *testing.T) {
	client := httpclient.NewHttpClient(baseUrl)

	header := map[string]string{
		"Authorization": "Bearer " + _ACCESS_TOKEN,
	}
	client.SetHeader(header)

	resp, err := client.Get("users/me")
	var parseResp view.Response

	err = json.Unmarshal(resp, &parseResp)
	expectCode := http.StatusOK
	actualCode := parseResp.Status

	if actualCode == http.StatusOK {
		data := parseResp.Data.(map[string]interface{})
		_USER_ID = data["id"].(string)
	}

	assert.Equal(t, expectCode, actualCode, err)
}

func TestExplore(t *testing.T) {
	client := httpclient.NewHttpClient(baseUrl)

	header := map[string]string{
		"Authorization": "Bearer " + _ACCESS_TOKEN,
	}
	client.SetHeader(header)

	resp, err := client.Get("users/explore")
	var parseResp view.Response

	err = json.Unmarshal(resp, &parseResp)
	expectCode := http.StatusOK
	actualCode := parseResp.Status

	if actualCode == http.StatusOK {
		rawData, _ := json.Marshal(parseResp.Data)
		var users []view.UserFindResponse
		err = json.Unmarshal(rawData, &users)

		for _, item := range users {
			userId := item.Id
			_exploredUsers = append(_exploredUsers, userId)
		}
	}

	assert.Equal(t, expectCode, actualCode, err)
}

func TestUpdateProfile(t *testing.T) {
	client := httpclient.NewHttpClient(baseUrl)

	header := map[string]string{
		"Authorization": "Bearer " + _ACCESS_TOKEN,
	}
	client.SetHeader(header)

	payload := param.UserUpdate{
		Name:       USER_NAME,
		Gender:     2,
		Email:      USER_EMAIL,
		BirthDate:  "2000-09-19",
		Phone:      "",
		LocationId: "61d2116c-f33a-4d0b-80ab-f8fc7aa8ae48",
		Photos:     []model.UserPhoto{},
		Interests:  []model.UserInterest{},
	}

	resp, err := client.Put("users/me", payload)
	var parseResp view.Response

	err = json.Unmarshal(resp, &parseResp)
	expectCode := http.StatusOK
	actualCode := parseResp.Status

	assert.Equal(t, expectCode, actualCode, err)
}
