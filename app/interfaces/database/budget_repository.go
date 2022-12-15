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

func (db *BudgetRepository) Update(b domain.BudgetUpdate, userId int) error {
	sql := `
		UPDATE budgets
			INNER JOIN travels
				ON budgets.id = travels.id
		SET budgets.accommodation = @accommodation, budgets.sightseeing = @sightseeing, budgets.meal = @meal, budgets.sum = @sum, budgets.transportations = @transports
		WHERE travels.user_id = @userId AND budgets.id = @id
	`
	arguments := map[string]interface{}{
		"accommodation": b.Accommodation,
		"sightseeing":   b.Sightseeing,
		"meal":          b.Meal,
		"sum":           b.Sum,
		"transports":    b.Transportations,
		"userId":        userId,
		"id":            b.ID,
	}
	err := db.RowSql(sql, arguments, &b)
	if err != nil {
		return err
	}
	return nil
}
