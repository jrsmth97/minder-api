package test

import (
	"encoding/json"
	"fmt"
	"minder/src/helper"
	"minder/src/server/enums"
	"minder/src/server/pkg/httpclient"
	"minder/src/server/view"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var _exploredUsersForSpecialUser []string

func exploreForPaidUser() {
	client := httpclient.NewHttpClient(baseUrl)

	header := map[string]string{
		"Authorization": "Bearer " + _SPECIAL_TOKEN,
	}

	client.SetHeader(header)
	resp, _ := client.Get("users/explore")
	var parseResp view.Response

	_ = json.Unmarshal(resp, &parseResp)
	respCode := parseResp.Status

	if respCode == http.StatusOK {
		rawData, _ := json.Marshal(parseResp.Data)
		var users []view.UserFindResponse
		_ = json.Unmarshal(rawData, &users)

		for _, item := range users {
			userId := item.Id
			_exploredUsersForSpecialUser = append(_exploredUsersForSpecialUser, userId)
		}
	}
}

func TestLikeSwipe(t *testing.T) {
	client := httpclient.NewHttpClient(baseUrl)

	header := map[string]string{
		"Authorization": "Bearer " + _ACCESS_TOKEN,
	}
	client.SetHeader(header)

	resp, err := client.Post("swipes/like/"+_exploredUsers[0], nil)
	var parseResp view.Response

	err = json.Unmarshal(resp, &parseResp)
	expectCode := http.StatusOK
	actualCode := parseResp.Status

	assert.Equal(t, expectCode, actualCode, err)
}

func TestPassSwipe(t *testing.T) {
	client := httpclient.NewHttpClient(baseUrl)

	header := map[string]string{
		"Authorization": "Bearer " + _ACCESS_TOKEN,
	}
	client.SetHeader(header)

	resp, err := client.Post("swipes/pass/"+_exploredUsers[1], nil)
	var parseResp view.Response

	err = json.Unmarshal(resp, &parseResp)
	expectCode := http.StatusOK
	actualCode := parseResp.Status

	assert.Equal(t, expectCode, actualCode, err)
}

func TestFavouriteSwipeForFreeMembership(t *testing.T) {
	client := httpclient.NewHttpClient(baseUrl)

	header := map[string]string{
		"Authorization": "Bearer " + _ACCESS_TOKEN,
	}
	client.SetHeader(header)

	resp, err := client.Post("swipes/favourite/"+_exploredUsers[2], nil)
	var parseResp view.Response

	err = json.Unmarshal(resp, &parseResp)
	expectCode := http.StatusBadRequest
	actualCode := parseResp.Status

	assert.Equal(t, expectCode, actualCode, err)
}

func TestFavouriteSwipeForPaidMembership(t *testing.T) {
	exploreForPaidUser()

	client := httpclient.NewHttpClient(baseUrl)

	header := map[string]string{
		"Authorization": "Bearer " + _SPECIAL_TOKEN,
	}
	client.SetHeader(header)

	resp, err := client.Post("swipes/favourite/"+_exploredUsersForSpecialUser[1], nil)
	var parseResp view.Response

	err = json.Unmarshal(resp, &parseResp)
	expectCode := http.StatusOK
	actualCode := parseResp.Status

	assert.Equal(t, expectCode, actualCode, err)
}

func TestMoreThanDayLimitSwipeForFreeMembership(t *testing.T) {
	client := httpclient.NewHttpClient(baseUrl)

	header := map[string]string{
		"Authorization": "Bearer " + _ACCESS_TOKEN,
	}
	client.SetHeader(header)

	const COUNTSWIPE = 10
	var expects = []int{
		http.StatusOK,
		http.StatusOK,
		http.StatusOK,
		http.StatusOK,
		http.StatusOK,
		http.StatusOK,
		http.StatusOK,
		http.StatusOK,
		http.StatusBadRequest,
		http.StatusBadRequest,
	}

	var actuals []int
	var actions = []string{
		enums.LikeSwipeAction,
		enums.PassSwipeAction,
	}

	for i := 1; i <= COUNTSWIPE; i++ {
		randomIdx := helper.RandomNumber(1)
		resp, _ := client.Post("swipes/"+actions[randomIdx]+"/"+_exploredUsers[i+1], nil)
		var parseResp view.Response

		_ = json.Unmarshal(resp, &parseResp)
		actualCode := parseResp.Status
		actuals = append(actuals, actualCode)
		fmt.Println("")
		fmt.Printf("SWIPE FREE MEMBERSHIP %vx RES STATUS CODE => %v", i, actualCode)
		fmt.Println("")
		time.Sleep(time.Second / 4)
	}

	assert.Equal(t, expects, actuals, nil)
}

func TestMoreThanDayLimitSwipeForPaidMembership(t *testing.T) {
	client := httpclient.NewHttpClient(baseUrl)

	header := map[string]string{
		"Authorization": "Bearer " + _SPECIAL_TOKEN,
	}
	client.SetHeader(header)

	const COUNTSWIPE = 11

	var actualRespMessages []string
	var actions = []string{
		enums.LikeSwipeAction,
		enums.PassSwipeAction,
	}

	for i := 1; i <= COUNTSWIPE; i++ {
		randomIdx := helper.RandomNumber(1)
		resp, _ := client.Post("swipes/"+actions[randomIdx]+"/"+_exploredUsersForSpecialUser[i+1], nil)
		var parseResp view.Response

		_ = json.Unmarshal(resp, &parseResp)
		if parseResp.Error != nil {
			actualRespMsg := parseResp.Error.(string)
			actualRespMessages = append(actualRespMessages, actualRespMsg)
		}

		fmt.Println("")
		fmt.Printf("SWIPE PAID MEMBERSHIP %vx RES STATUS CODE => %v", i, parseResp.Status)
		fmt.Println("")

		time.Sleep(time.Second / 4)
	}

	swipeLimitErrMessage := "Swipe is limited 10 times per day for Free Plan membership"
	assert.NotContains(t, actualRespMessages, swipeLimitErrMessage, nil)
}
