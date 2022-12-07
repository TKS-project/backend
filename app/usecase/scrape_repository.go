package usecase

import "github.com/Kantaro0829/clean-architecture-in-go/domain"

type ScrapeRepository interface {
	FindMeaning(word string) string

	FindTransports() string

	FindOneTransport(
		orvStation string,
		dnvStation string,
		year string,
		month string,
		day string,
		hour string,
		minute string,
	) domain.TransportsJson
}
