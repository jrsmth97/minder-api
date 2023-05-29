package test

import (
	"encoding/json"
	"minder/src/helper"
	"minder/src/server/param"
	"minder/src/server/pkg/httpclient"
	"minder/src/server/view"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

// for admin routing only
const _SPECIAL_TOKEN = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3N1ZWQiOiIyMTM3LTA2LTI2VDA0OjQ1OjM1Ljg0MTg5NyswNzowMCIsInBheWxvYWQiOnsidXNlcl9pZCI6IjBiOGNlN2YzLTUyMGQtNDcxZi1iNzEwLTUxYzcxMzZlOWNmMyIsImVtYWlsIjoiYWRtaW5AbWluZGVyLmNvbSJ9fQ.BTufVrhMfx0PrIzYbpnpZeEZm6jjw6q6zIUFr8Y1btc"

var _MEMBERSHIP_ID string
var _PREMIUM_MEMBERSHIP_ID string

func TestGetAllMemberships(t *testing.T) {
	client := httpclient.NewHttpClient(baseUrl)

	header := map[string]string{
		"Authorization": "Bearer " + _ACCESS_TOKEN,
	}
	client.SetHeader(header)

	resp, err := client.Get("memberships")
	var parseResp view.Response

	err = json.Unmarshal(resp, &parseResp)
	expectCode := http.StatusOK
	actualCode := parseResp.Status

	if actualCode == http.StatusOK {
		rawData, _ := json.Marshal(parseResp.Data)
		var memberships []view.MembershipGetResponse
		_ = json.Unmarshal(rawData, &memberships)

		premiumMemberIdx := helper.FindIndex(memberships, func(value interface{}) bool {
			return value.(view.MembershipGetResponse).MembershipName == "Diamond Plan"
		})
		_PREMIUM_MEMBERSHIP_ID = memberships[premiumMemberIdx].Id
	}

	assert.Equal(t, expectCode, actualCode, err)
}

func TestCreateMembership(t *testing.T) {
	client := httpclient.NewHttpClient(baseUrl)

	header := map[string]string{
		"Authorization": "Bearer " + _SPECIAL_TOKEN,
	}
	client.SetHeader(header)

	payload := param.MembershipCreate{
		MembershipName:  "Test Membership 1",
		Price:           10000,
		DurationInMonth: 1,
		Description:     "This is test membership",
		Privileges:      []string{"f456b60c-e4d7-4c2c-b843-b25f37c9b4bd"},
	}

	resp, err := client.Post("memberships", payload)
	var parseResp view.Response

	err = json.Unmarshal(resp, &parseResp)
	expectCode := http.StatusCreated
	actualCode := parseResp.Status

	if parseResp.Status == http.StatusCreated {
		data := parseResp.Data.(map[string]interface{})
		_MEMBERSHIP_ID = data["id"].(string)
	}

	assert.Equal(t, expectCode, actualCode, err)
}

func TestGetMembership(t *testing.T) {
	client := httpclient.NewHttpClient(baseUrl)

	header := map[string]string{
		"Authorization": "Bearer " + _ACCESS_TOKEN,
	}
	client.SetHeader(header)

	resp, err := client.Get("memberships/" + _MEMBERSHIP_ID)
	var parseResp view.Response

	err = json.Unmarshal(resp, &parseResp)
	expectCode := http.StatusOK
	actualCode := parseResp.Status

	assert.Equal(t, expectCode, actualCode, err)
}

func TestUpdateMembership(t *testing.T) {
	client := httpclient.NewHttpClient(baseUrl)

	header := map[string]string{
		"Authorization": "Bearer " + _SPECIAL_TOKEN,
	}
	client.SetHeader(header)

	payload := param.MembershipCreate{
		MembershipName:  "Test Membership 1",
		Price:           15000,
		DurationInMonth: 1,
		Description:     "This is test membership updated",
		Privileges:      []string{"f456b60c-e4d7-4c2c-b843-b25f37c9b4bd"},
	}

	resp, err := client.Put("memberships/"+_MEMBERSHIP_ID, payload)
	var parseResp view.Response

	err = json.Unmarshal(resp, &parseResp)
	expectCode := http.StatusOK
	actualCode := parseResp.Status

	assert.Equal(t, expectCode, actualCode, err)
}

func TestDeleteMembership(t *testing.T) {
	client := httpclient.NewHttpClient(baseUrl)

	header := map[string]string{
		"Authorization": "Bearer " + _SPECIAL_TOKEN,
	}
	client.SetHeader(header)

	resp, err := client.Delete("memberships/" + _MEMBERSHIP_ID)
	var parseResp view.Response

	err = json.Unmarshal(resp, &parseResp)
	expectCode := http.StatusOK
	actualCode := parseResp.Status

	assert.Equal(t, expectCode, actualCode, err)
}
