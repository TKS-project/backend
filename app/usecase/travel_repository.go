package usecase

import (
	"github.com/Kantaro0829/clean-architecture-in-go/domain"
)

type TravelRepository interface {
	Insert(t domain.Travel) (int, error)
}
