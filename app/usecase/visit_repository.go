package usecase

import (
	"github.com/Kantaro0829/clean-architecture-in-go/domain"
)

type VisitRepository interface {
	Insert(v domain.Visit) (int, error)
}
