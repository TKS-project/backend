package usecase

import (
	"github.com/Kantaro0829/clean-architecture-in-go/domain"
)

type ApiRequestRepository interface {
	FindRoom() domain.VacantHotels
}
