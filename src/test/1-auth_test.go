package test

import (
	"encoding/json"
	"fmt"
	"minder/src/helper"
	"minder/src/server/param"
	"minder/src/server/pkg/httpclient"
	"minder/src/server/view"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const baseUrl = "http://localhost:4444"

var USER_NAME = ""
var USER_EMAIL = ""
var USER_PASS = "user"

var _ACCESS_TOKEN string
var _locationIds []string

func fetchCountUsers() {
	client := httpclient.NewHttpClient(baseUrl)

	header := map[string]string{
		"Authorization": "Bearer " + _SPECIAL_TOKEN,
	}

	client.SetHeader(header)
	resp, _ := client.Get("users/count")
	var parseResp view.Response

	_ = json.Unmarshal(resp, &parseResp)
	respCode := parseResp.Status

	if respCode == http.StatusOK {
		countUsers := parseResp.Data.(float64)
		USER_NAME = fmt.Sprintf("User %v", countUsers)
		USER_EMAIL = fmt.Sprintf("user%v@mail.com", countUsers)
	}
}

func fetchLocations() {
	client := httpclient.NewHttpClient(baseUrl)

	header := map[string]string{
		"Authorization": "Bearer " + _SPECIAL_TOKEN,
	}

	client.SetHeader(header)
	resp, _ := client.Get("locations")
	var parseResp view.Response

	_ = json.Unmarshal(resp, &parseResp)
	respCode := parseResp.Status

	if respCode == http.StatusOK {
		rawData, _ := json.Marshal(parseResp.Data)
		var locations []view.LocationGetAllResponse
		_ = json.Unmarshal(rawData, &locations)

		for _, item := range locations {
			locationId := item.Id
			_locationIds = append(_locationIds, locationId)
		}
	}
}

func TestRegister(t *testing.T) {
	fetchCountUsers()
	time.Sleep(time.Second / 2)
	fetchLocations()
	time.Sleep(time.Second / 2)

	client := httpclient.NewHttpClient(baseUrl)

	locationId := _locationIds[helper.RandomNumber(len(_locationIds)-1)]
	payload := param.AuthRegister{
		Name:       USER_NAME,
		Gender:     1,
		Email:      USER_EMAIL,
		Password:   USER_PASS,
		BirthDate:  "1999-09-09",
		Phone:      "08193938933",
		LocationId: locationId,
	}

	resp, err := client.Post("auth/register", payload)
	var parseResp view.Response

	err = json.Unmarshal(resp, &parseResp)
	expectCode := http.StatusCreated
	actualCode := parseResp.Status

	if actualCode == http.StatusCreated {
		fmt.Println("")
		fmt.Printf("USER CREATED:")
		fmt.Println("")
		fmt.Printf("NAME => %s", USER_NAME)
		fmt.Println("")
		fmt.Printf("EMAIL => %s", USER_EMAIL)
		fmt.Println("")
	}

	assert.Equal(t, expectCode, actualCode, err)
}

func TestRegisterInvalidLocation(t *testing.T) {
	client := httpclient.NewHttpClient(baseUrl)

	payload := param.AuthRegister{
		Name:       USER_NAME,
		Gender:     1,
		Email:      USER_EMAIL,
		Password:   USER_PASS,
		BirthDate:  "1999-09-09",
		Phone:      "08193938933",
		LocationId: "120bbdb1-50b8-41f9-868b-4bbf98d583ae", // invalid location id
	}

	resp, err := client.Post("auth/register", payload)
	var parseResp view.Response

	err = json.Unmarshal(resp, &parseResp)
	expectCode := http.StatusBadRequest
	actualCode := parseResp.Status

	assert.Equal(t, expectCode, actualCode, err)
}

func TestLogin(t *testing.T) {
	client := httpclient.NewHttpClient(baseUrl)

	payload := param.UserLogin{
		Email:    USER_EMAIL,
		Password: USER_PASS,
	}

	resp, err := client.Post("auth/login", payload)
	var parseResp view.Response

	err = json.Unmarshal(resp, &parseResp)
	expectCode := http.StatusOK
	actualCode := parseResp.Status

	if actualCode == http.StatusOK {
		data := parseResp.Data.(map[string]interface{})
		_ACCESS_TOKEN = data["access_token"].(string)
	}

	assert.Equal(t, expectCode, actualCode, err)
}

func TestInvalidLogin(t *testing.T) {
	client := httpclient.NewHttpClient(baseUrl)

	payload := param.UserLogin{
		Email:    "wronguser@mail.com",
		Password: "wrongpassword",
	}

	resp, err := client.Post("auth/login", payload)
	var parseResp view.Response

	err = json.Unmarshal(resp, &parseResp)
	expectCode := http.StatusUnauthorized
	actualCode := parseResp.Status

	assert.Equal(t, expectCode, actualCode, err)
}
