package database

import (
	"github.com/Kantaro0829/clean-architecture-in-go/domain"
)

type PrefectureRepository struct {
	SqlHandler
}

//DB操作のための関数
//同階層の./interface/database/user_repogitory.goから呼び出している

//本来ここで詳細な操作をするべき？ 構造体の指定などもここ？
func (db *PrefectureRepository) FindAll() ([]domain.Prefecture, error) {
	result, err := db.GetAllPrefectures()
	if err != nil {
		return []domain.Prefecture{}, err
	}
	return result, err
}

func (db *PrefectureRepository) FindAllCities() ([]domain.Citie, error) {
	cities, err := db.GetAllCities()
	if err != nil {
		return []domain.Citie{}, err
	}
	return cities, err
}

func (db *PrefectureRepository) FindAllDetailedCities() ([]domain.DetailedCitie, error) {
	detailedCities, err := db.GetAllDetailedCities()
	if err != nil {
		return []domain.DetailedCitie{}, err
	}
	return detailedCities, err
}
