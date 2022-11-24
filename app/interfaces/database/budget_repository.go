package database

import (
	"github.com/Kantaro0829/clean-architecture-in-go/domain"
)

type BudgetRepository struct {
	SqlHandler
}

func (db *BudgetRepository) Insert(b domain.Budget) (int, error) {
	err := db.Create(&b)
	if err != nil {
		return 0, err
	}
	budgetlId := b.ID
	return budgetlId, nil
}
