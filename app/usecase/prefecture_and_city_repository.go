package usecase

import (
	"github.com/Kantaro0829/clean-architecture-in-go/domain"
)

type PrefectureRepository interface {
	FindAll() ([]domain.Prefecture, error)
	FindAllCities() ([]domain.Citie, error)
	FindAllDetailedCities() ([]domain.DetailedCitie, error)
	GetCitiesByPreId(prefectureId int) ([]domain.Citie, error)
	IsDetailedCityExisting(int) (bool, error)
	GetDetailedCitiesByCityId(cityId int) ([]domain.DetailedCitie, error)
}
