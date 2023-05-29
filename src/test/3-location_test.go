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

var _LOCATION_ID string

func TestGetAllLocations(t *testing.T) {
	client := httpclient.NewHttpClient(baseUrl)

	header := map[string]string{
		"Authorization": "Bearer " + _ACCESS_TOKEN,
	}
	client.SetHeader(header)

	resp, err := client.Get("locations")
	var parseResp view.Response

	err = json.Unmarshal(resp, &parseResp)
	expectCode := http.StatusOK
	actualCode := parseResp.Status

	assert.Equal(t, expectCode, actualCode, err)
}

func TestCreateLocation(t *testing.T) {
	client := httpclient.NewHttpClient(baseUrl)

	header := map[string]string{
		"Authorization": "Bearer " + _SPECIAL_TOKEN,
	}
	client.SetHeader(header)

	payload := param.LocationCreate{
		LocationName: "Test Location 1",
	}

	resp, err := client.Post("locations", payload)
	var parseResp view.Response

	err = json.Unmarshal(resp, &parseResp)
	expectCode := http.StatusCreated
	actualCode := parseResp.Status

	if parseResp.Status == http.StatusCreated {
		data := parseResp.Data.(map[string]interface{})
		_LOCATION_ID = data["id"].(string)
	}

	assert.Equal(t, expectCode, actualCode, err)
}

func TestGetLocation(t *testing.T) {
	client := httpclient.NewHttpClient(baseUrl)

	header := map[string]string{
		"Authorization": "Bearer " + _ACCESS_TOKEN,
	}
	client.SetHeader(header)

	resp, err := client.Get("locations/" + _LOCATION_ID)
	var parseResp view.Response

	err = json.Unmarshal(resp, &parseResp)
	expectCode := http.StatusOK
	actualCode := parseResp.Status

	assert.Equal(t, expectCode, actualCode, err)
}

func TestUpdateLocation(t *testing.T) {
	client := httpclient.NewHttpClient(baseUrl)

	header := map[string]string{
		"Authorization": "Bearer " + _SPECIAL_TOKEN,
	}
	client.SetHeader(header)

	payload := param.LocationUpdate{}
	payload.LocationName = "Test Location 1 updated"

	resp, err := client.Put("locations/"+_LOCATION_ID, payload)
	var parseResp view.Response

	err = json.Unmarshal(resp, &parseResp)
	expectCode := http.StatusOK
	actualCode := parseResp.Status

	assert.Equal(t, expectCode, actualCode, err)
}

func TestDeleteLocation(t *testing.T) {
	client := httpclient.NewHttpClient(baseUrl)

	header := map[string]string{
		"Authorization": "Bearer " + _SPECIAL_TOKEN,
	}
	client.SetHeader(header)

	resp, err := client.Delete("locations/" + _LOCATION_ID)
	var parseResp view.Response

	err = json.Unmarshal(resp, &parseResp)
	expectCode := http.StatusOK
	actualCode := parseResp.Status

	assert.Equal(t, expectCode, actualCode, err)
}
