package repository

import (
	"template-ulamm-backend-go/pkg/datasource"
	"template-ulamm-backend-go/pkg/model/entity"
)

type CityRepository interface {
	FindAll() ([]entity.City, error)
}

type cityRepository struct {
	Datasource *datasource.Datasource
}

func NewCityRepository(datasource *datasource.Datasource) CityRepository {
	return &cityRepository{
		Datasource: datasource,
	}
}

func (cr *cityRepository) FindAll() ([]entity.City, error) {
	var cities []entity.City
	if err := cr.Datasource.GormDB.Find(&cities).Error; err != nil {
		return nil, err
	}
	return cities, nil
}
