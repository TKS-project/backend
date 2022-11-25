package usecase

import "github.com/Kantaro0829/clean-architecture-in-go/domain"

type VisitInteractor struct {
	VisitRepository VisitRepository
}

func (interactor *VisitInteractor) Add(v domain.Visit) (int, error) {
	visitId, err := interactor.VisitRepository.Insert(v)
	if err != nil {
		return 0, err
	}
	return visitId, nil
}
