package view

import (
	"minder/src/server/model"
)

type InterestCreateResponse struct {
	Id           string `json:"id"`
	InterestName string `json:"interest_name"`
}

func NewInterestCreateResponse(interest *model.Interest) *InterestCreateResponse {
	return &InterestCreateResponse{
		Id:           interest.ID.String(),
		InterestName: interest.InterestName,
	}
}

type InterestUpdateResponse struct {
	InterestCreateResponse
}

func NewInterestUpdateResponse(interest *model.Interest) *InterestUpdateResponse {
	updateResponse := &InterestUpdateResponse{}
	updateResponse.Id = interest.ID.String()
	updateResponse.InterestName = interest.InterestName
	return updateResponse
}

type InterestGetAllResponse struct {
	InterestCreateResponse
}

func NewInterestGetAllResponse(interests *[]model.Interest) *[]InterestGetAllResponse {
	var interestsResponse []InterestGetAllResponse

	for _, interest := range *interests {
		response := &InterestGetAllResponse{}
		response.Id = interest.ID.String()
		response.InterestName = interest.InterestName
		interestsResponse = append(interestsResponse, InterestGetAllResponse(*response))
	}

	return &interestsResponse
}
