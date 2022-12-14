package usecase

import (
	"github.com/Kantaro0829/clean-architecture-in-go/domain"
)

type SightseeingInteractor struct {
	SightseeingRepository SightseeingRepository
}

func (interactor *SightseeingInteractor) GetSightseeingInfo(prefectureId int) ([]domain.Sightseeing, error) {
	results, err := interactor.SightseeingRepository.GetSightInfoByPrefectureId(prefectureId)
	if err != nil {
		return results, err
	}
	return results, nil
}
