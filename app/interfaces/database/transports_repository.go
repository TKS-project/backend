package database

import (
	"github.com/Kantaro0829/clean-architecture-in-go/domain"
)

type TransportsRepository struct {
	SqlHandler
}

func (db *TransportsRepository) Insert(b domain.Transportation) error {
	err := db.Create(&b)
	if err != nil {
		return err
	}
	return nil
}
