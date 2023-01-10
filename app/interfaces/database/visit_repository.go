package database

import (
	"fmt"

	"github.com/Kantaro0829/clean-architecture-in-go/domain"
)

type VisitRepository struct {
	SqlHandler
}

func (db *VisitRepository) Insert(v domain.Visit) (int, error) {
	fmt.Println(v)

	err := db.Create(&v)
	if err != nil {
		return 0, err
	}
	visitId := v.ID
	return visitId, nil
}

func (db *VisitRepository) Update(v domain.Visit) error {
	sql := `
		UPDATE visits
			INNER JOIN budgets
				ON visits.budget_id = budgets.id
			INNER JOIN travel
				ON budgets.travel_id = travels.id
		SET visits.address
		Where travels.user_id = @userId AND visits.id = @id

	`
	arguments := map[string]interface{}{}
	_, err := db.RowSql(sql, arguments, &v)
	if err != nil {
		return err
	}
	return nil
}
