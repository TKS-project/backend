package usecase

import "github.com/Kantaro0829/clean-architecture-in-go/domain"

type TransportsInteractor struct {
	TransportsRepository TransportsRepository
}

func (interactor *TransportsInteractor) Add(b domain.Transportation) error {
	err := interactor.TransportsRepository.Insert(b)
	if err != nil {
		return err
	}
	return nil
}
