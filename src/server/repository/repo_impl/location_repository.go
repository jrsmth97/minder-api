package repo_impl

import (
	"minder/src/server/model"
	"minder/src/server/repository"

	"gorm.io/gorm"
)

type locationRepo struct {
	db *gorm.DB
}

func NewLocationRepo(db *gorm.DB) repository.LocationRepo {
	return &locationRepo{
		db: db,
	}
}

func (r *locationRepo) GetLocations() (*[]model.Location, error) {
	var locations []model.Location

	err := r.db.Where("deleted_at IS NULL").Find(&locations).Error
	if err != nil {
		return nil, err
	}

	return &locations, nil
}

func (r *locationRepo) CreateLocation(location *model.Location) error {
	err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(location).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

func (r *locationRepo) UpdateLocation(id string, location *model.Location) error {
	return r.db.Where("id = ?", id).Updates(location).Error
}

func (r *locationRepo) DeleteLocation(id string, location *model.Location) error {
	err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Exec("DELETE from locations WHERE id = ?", id).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

func (r *locationRepo) GetLocationById(locationId string) (*model.Location, error) {
	var location model.Location
	err := r.db.Where("locations.id=?", locationId).Where("locations.deleted_at IS NULL").First(&location).Error
	if err != nil {
		return nil, err
	}

	return &location, nil
}
