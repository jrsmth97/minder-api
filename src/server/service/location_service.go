package service

import (
	"database/sql"
	"minder/src/server/param"
	"minder/src/server/repository"
	"minder/src/server/view"

	"github.com/gin-gonic/gin"
)

type LocationService struct {
	repo repository.LocationRepo
}

func NewLocationServices(
	repo repository.LocationRepo,
) *LocationService {
	return &LocationService{
		repo: repo,
	}
}

func (p *LocationService) Preparation(c *gin.Context) {
	_context = c
}

func (p *LocationService) GetLocations() *view.Response {
	locations, err := p.repo.GetLocations()

	if err != nil {
		if err == sql.ErrNoRows {
			return view.ErrNotFound()
		}
		return view.ErrInternalServer(err.Error())
	}

	return view.SuccessFind(view.NewLocationGetAllResponse(locations))
}

func (p *LocationService) GetLocationById(locationId string) *view.Response {
	location, err := p.repo.GetLocationById(locationId)
	if err != nil {
		return view.ErrInternalServer(err.Error())
	}

	return view.SuccessFind(location)
}

func (p *LocationService) CreateLocation(req *param.LocationCreate) *view.Response {
	location := req.ParseToModel()

	err := p.repo.CreateLocation(location)
	if err != nil {
		return view.ErrInternalServer(err.Error())
	}

	return view.SuccessCreated(view.NewLocationCreateResponse(location))
}

func (p *LocationService) UpdateLocation(locationId string, req *param.LocationUpdate) *view.Response {
	location, errLocation := p.repo.GetLocationById(locationId)
	if errLocation != nil {
		return view.ErrBadRequest("location doesn't exists")
	}

	location.LocationName = req.LocationName
	err := p.repo.UpdateLocation(locationId, location)
	if err != nil {
		return view.ErrInternalServer(err.Error())
	}

	return view.SuccessUpdated(view.NewLocationUpdateResponse(location))
}

func (p *LocationService) DeleteLocation(locationId string) *view.Response {
	location, err := p.repo.GetLocationById(locationId)
	if err != nil {
		return view.ErrBadRequest("location doesn't exists")
	}

	err = p.repo.DeleteLocation(locationId, location)
	if err != nil {
		return view.ErrInternalServer(err.Error())
	}

	return view.SuccessDeleted(view.NewLocationUpdateResponse(location))
}
