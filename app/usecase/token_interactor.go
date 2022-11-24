package usecase

import "github.com/Kantaro0829/clean-architecture-in-go/domain"

type TokenInteractor struct {
	Tokenizer Tokenizer
}

func (interactor *TokenInteractor) Create(u domain.User) (domain.Token, error) {
	token, err := interactor.Tokenizer.New(u)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (interactor *TokenInteractor) GetId(t domain.Token) (int, error) {
	id, err := interactor.Tokenizer.Verify(t)
	if err != nil {
		return 0, err
	}
	return id, nil
}
