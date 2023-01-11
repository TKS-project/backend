package database

import (
	"fmt"

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

func (db *PrefectureRepository) GetCitiesByPreId(prefectureId int) ([]domain.Citie, error) {
	//column
	//where prefecture code 指定
	//scan
	arguments := map[string]interface{}{
		"preId": prefectureId,
	}

	sql := `
	select * from cities where prefecture_id = @preId
	`
	result := []domain.Citie{}
	_, err := db.RowSql(sql, arguments, &result)
	if err != nil {
		return nil, err
	}
	fmt.Println(result)
	return result, nil
}

func (db *PrefectureRepository) FindAllDetailedCities() ([]domain.DetailedCitie, error) {
	detailedCities, err := db.GetAllDetailedCities()
	if err != nil {
		return []domain.DetailedCitie{}, err
	}
	return detailedCities, err
}

func (db *PrefectureRepository) GetDetailedCitiesByCityId(cityId int) ([]domain.DetailedCitie, error) {
	arguments := map[string]interface{}{
		"cityId": int16(cityId),
	}

	sql := `
	select * from detailed_cities where city_id = @cityId
	`
	result := []domain.DetailedCitie{}
	_, err := db.RowSql(sql, arguments, &result)
	if err != nil {
		return nil, err
	}
	fmt.Println(result)
	return result, nil
}

func (db *PrefectureRepository) IsDetailedCityExisting(cityId int) (bool, error) {
	fmt.Printf("cityId: %v", cityId)
	dCity := domain.DetailedCitie{}
	isExist, err := db.IsRecordExisting(&dCity, &domain.DetailedCitie{CityId: int16(cityId)})
	if err != nil {
		return false, err
	}
	return isExist, nil
}
