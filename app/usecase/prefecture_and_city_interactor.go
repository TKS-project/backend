package usecase

import "github.com/Kantaro0829/clean-architecture-in-go/domain"

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
