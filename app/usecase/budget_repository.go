package usecase

import (
	"github.com/Kantaro0829/clean-architecture-in-go/domain"
)

type BudgetRepository interface {
	Insert(b domain.Budget) (int, error)
	Update(b domain.BudgetUpdate, userId int) error
}
