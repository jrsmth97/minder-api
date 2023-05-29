package view

import (
	"minder/src/server/model"
)

type LocationCreateResponse struct {
	Id           string `json:"id"`
	LocationName string `json:"location_name"`
}

type LocationUpdateResponse struct {
	LocationCreateResponse
}

func NewLocationCreateResponse(location *model.Location) *LocationCreateResponse {
	return &LocationCreateResponse{
		Id:           location.ID.String(),
		LocationName: location.LocationName,
	}
}

func NewLocationUpdateResponse(location *model.Location) *LocationUpdateResponse {
	updateResponse := &LocationUpdateResponse{}
	updateResponse.Id = location.ID.String()
	updateResponse.LocationName = location.LocationName
	return updateResponse
}

type LocationGetAllResponse struct {
	LocationCreateResponse
}

func NewLocationGetAllResponse(locations *[]model.Location) *[]LocationGetAllResponse {
	var locationsResponse []LocationGetAllResponse

	for _, location := range *locations {
		response := &LocationGetAllResponse{}
		response.Id = location.ID.String()
		response.LocationName = location.LocationName
		locationsResponse = append(locationsResponse, *response)
	}

	return &locationsResponse
}
