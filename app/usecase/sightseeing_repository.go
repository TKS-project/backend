package usecase

import (
	"github.com/Kantaro0829/clean-architecture-in-go/domain"
)

type SightseeingRepository interface {
	GetSightInfoByPrefectureId(preId int) ([]domain.Sightseeing, error)
}
