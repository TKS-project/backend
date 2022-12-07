package usecase

import "github.com/Kantaro0829/clean-architecture-in-go/domain"

type GetMeaningInteractor struct {
	ScrapeRepository ScrapeRepository
	// //以下試しに追加
	// //Tokenizer token.Tokenizer
	// Tokenizer Tokenizer
}

func (interactor *GetMeaningInteractor) GetTranslatedMeaning(word string) string {

	meaning := interactor.ScrapeRepository.FindMeaning(word)
	return meaning
}

func (interactor *GetMeaningInteractor) GetTransports() string {
	text := interactor.ScrapeRepository.FindTransports()
	return text
}

func (interactor *GetMeaningInteractor) GetOneTransports(
	transport domain.TransportsRequest,
) domain.TransportsJson {

	res := interactor.ScrapeRepository.FindOneTransport(
		transport.OrvStation,
		transport.DnvStation,
		transport.Year,
		transport.Month,
		transport.Day,
		transport.Hour,
		transport.Minute,
	)

	return res
}
