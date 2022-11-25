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
