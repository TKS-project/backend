package database

import (
	"fmt"

	"github.com/Kantaro0829/clean-architecture-in-go/domain"
)

type TravelRepository struct {
	SqlHandler
}

func (db *TravelRepository) Insert(t domain.Travel) (int, error) {
	fmt.Println(t)

	err := db.Create(&t)
	if err != nil {
		return 0, err
	}
	travelId := t.ID
	return travelId, nil
}

func (db *TravelRepository) UpdateTravel(t domain.Travel) error {
	err := db.Update(&t)
	if err != nil {
		return err
	}
	return nil
}
