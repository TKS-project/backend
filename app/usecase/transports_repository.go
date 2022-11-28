package usecase

import (
	"github.com/Kantaro0829/clean-architecture-in-go/domain"
)

type TransportsRepository interface {
	Insert(b domain.Transportation) error
}
