package usecase

import (
	"template-ulamm-backend-go/pkg/model/entity"
	"template-ulamm-backend-go/pkg/repository"
)

type CityUsecase interface {
	FindAllUser() ([]entity.City, error)
}

type cityUsecase struct {
	repository repository.CityRepository
}

func NewCityUsecase(cityRepository repository.CityRepository) *cityUsecase {
	return &cityUsecase{cityRepository}
}

func (cu *cityUsecase) FindAllUser() ([]entity.City, error) {
	cities, err := cu.repository.FindAll()
	if err != nil {
		return cities, err
	}

	return cities, nil
}
