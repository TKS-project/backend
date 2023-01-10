package usecase

import "github.com/Kantaro0829/clean-architecture-in-go/domain"

type TravelInteractor struct {
	TravelRepository TravelRepository
	Tokenizer        Tokenizer
}

func (interactor *TravelInteractor) Authenticate(token domain.Token) (int, error) {
	id, err := interactor.Tokenizer.Verify(token)
	if err != nil {

		return 0, err
	}
	return id, nil
}

func (interactor *TravelInteractor) Add(t domain.Travel) (int, error) {
	travelId, err := interactor.TravelRepository.Insert(t)
	return travelId, err
}

func (interactor *TravelInteractor) Update(t domain.Travel) error {
	err := interactor.TravelRepository.UpdateTravel(t)
	return err
}
