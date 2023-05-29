package test

import (
	"encoding/json"
	"minder/src/server/param"
	"minder/src/server/pkg/httpclient"
	"minder/src/server/view"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

var _purchaseId string

func TestCreatePurchase(t *testing.T) {
	client := httpclient.NewHttpClient(baseUrl)

	header := map[string]string{
		"Authorization": "Bearer " + _ACCESS_TOKEN,
	}
	client.SetHeader(header)

	payload := param.CreatePurchase{
		MembershipId:  _PREMIUM_MEMBERSHIP_ID,
		PaymentMethod: "bca",
	}

	resp, err := client.Post("purchases", payload)
	var parseResp view.Response

	err = json.Unmarshal(resp, &parseResp)
	expectCode := http.StatusCreated
	actualCode := parseResp.Status

	if parseResp.Status == http.StatusCreated {
		data := parseResp.Data.(map[string]interface{})
		_purchaseId = data["id"].(string)
	}

	assert.Equal(t, expectCode, actualCode, err)
}

func TestGetPurchase(t *testing.T) {
	client := httpclient.NewHttpClient(baseUrl)

	header := map[string]string{
		"Authorization": "Bearer " + _ACCESS_TOKEN,
	}
	client.SetHeader(header)

	resp, err := client.Get("purchases/" + _purchaseId)
	var parseResp view.Response

	err = json.Unmarshal(resp, &parseResp)
	expectCode := http.StatusOK
	actualCode := parseResp.Status

	assert.Equal(t, expectCode, actualCode, err)
}

func TestCancelPurchase(t *testing.T) {
	client := httpclient.NewHttpClient(baseUrl)

	header := map[string]string{
		"Authorization": "Bearer " + _ACCESS_TOKEN,
	}
	client.SetHeader(header)

	resp, err := client.Post("purchases/cancel/"+_purchaseId, nil)
	var parseResp view.Response

	err = json.Unmarshal(resp, &parseResp)
	expectCode := http.StatusOK
	actualCode := parseResp.Status

	assert.Equal(t, expectCode, actualCode, err)
}

func TestSyncPurchases(t *testing.T) {
	client := httpclient.NewHttpClient(baseUrl)

	header := map[string]string{
		"Authorization": "Bearer " + _SPECIAL_TOKEN,
	}
	client.SetHeader(header)

	resp, err := client.Post("purchases/sync", nil)
	var parseResp view.Response

	err = json.Unmarshal(resp, &parseResp)
	expectCode := http.StatusOK
	actualCode := parseResp.Status

	_END_OF_TESTING = true
	assert.Equal(t, expectCode, actualCode, err)
}
