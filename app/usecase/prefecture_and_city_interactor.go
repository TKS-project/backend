package usecase

import (
	"github.com/Kantaro0829/clean-architecture-in-go/domain"
)

type PrefectureAndCityInteractor struct {
	PrefectureRepository PrefectureRepository
}

func (interactor *PrefectureAndCityInteractor) FindAllPrefecture() ([]domain.Prefecture, error) {

	prefectures, err := interactor.PrefectureRepository.FindAll()
	if err != nil {
		return []domain.Prefecture{}, err
	}
	return prefectures, err
}

func (interactor *PrefectureAndCityInteractor) FindAllCities() ([]domain.Citie, error) {
	cities, err := interactor.PrefectureRepository.FindAllCities()
	if err != nil {
		return []domain.Citie{}, err
	}
	return cities, err
}

func (interactor *PrefectureAndCityInteractor) FindAllDetailedCities() ([]domain.DetailedCitie, error) {
	detailedCities, err := interactor.PrefectureRepository.FindAllDetailedCities()
	if err != nil {
		return []domain.DetailedCitie{}, err
	}
	return detailedCities, err
}

func (interactor *PrefectureAndCityInteractor) FindCitiesByPrefectureId(prefectureId int) ([]domain.Citie, error) {
	cities, err := interactor.PrefectureRepository.GetCitiesByPreId(prefectureId)
	if err != nil {
		return nil, err
	}
	return cities, nil
}

func (interactor *PrefectureAndCityInteractor) FindDetailedCitiesByCityId(cityId int) ([]domain.DetailedCitie, error) {
	dCities, err := interactor.PrefectureRepository.GetDetailedCitiesByCityId(cityId)
	if err != nil {
		return nil, err
	}
	return dCities, nil
}

func (interactor *PrefectureAndCityInteractor) IsDetailedCityExist(cityId int) (bool, error) {
	isExist, err := interactor.PrefectureRepository.IsDetailedCityExisting(cityId)
	if err != nil {
		return false, err
	}
	return isExist, nil
}
