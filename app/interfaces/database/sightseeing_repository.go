package database

import (
	"fmt"

	"github.com/Kantaro0829/clean-architecture-in-go/domain"
)

type SightseeingRepository struct {
	SqlHandler
}

func (db *SightseeingRepository) GetSightInfoByPrefectureId(preId int) ([]domain.Sightseeing, error) {

	where := map[string]interface{}{"prefecture_id": preId}
	results := []domain.Sightseeing{}
	_, err := db.Row("SELECT prefecture_id, name, image, description, city FROM sightseeings", where, &results)
	if err != nil {
		return results, err
	}
	fmt.Println(results)

	return results, nil
}
