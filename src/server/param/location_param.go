package param

import (
	"minder/src/server/model"
)

type LocationCreate struct {
	LocationName string `validate:"required" json:"location_name"`
}

type LocationUpdate struct {
	LocationCreate
}

func (l *LocationCreate) ParseToModel() *model.Location {
	return &model.Location{
		LocationName: l.LocationName,
	}
}
