package infrastructure

import (
	"fmt"

	"github.com/Kantaro0829/clean-architecture-in-go/domain"
	"github.com/Kantaro0829/clean-architecture-in-go/interfaces/database"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type SqlHandler struct {
	db *gorm.DB
}

func NewSqlHandler() database.SqlHandler {
	dsn := "root:ecc@tcp(db:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error)
	}
	sqlHandler := new(SqlHandler)
	sqlHandler.db = db
	return sqlHandler
}

//基本的に引数はinterface{}で受け取るDIPしたときにどんな型の構造体でも受け入れれる
//	&はここではつけない
func (handler *SqlHandler) Create(obj interface{}) error {
	if err := handler.db.Create(obj).Error; err != nil {
		return err
	}
	return nil
}

func (handler *SqlHandler) FindAll(obj interface{}) (interface{}, error) {
	if err := handler.db.Find(obj).Error; err != nil {
		return nil, err
	}
	return obj, nil
}

func (handler *SqlHandler) FindOne(obj interface{}, columns []string, where interface{}) (interface{}, error) {
	if err := handler.db.Select(columns).Where(where).First(obj).Error; err != nil {
		return nil, err
	}
	return obj, nil

}

func (handler *SqlHandler) RowSql(s string, arguments interface{}, scan interface{}) (interface{}, error) {
	if err := handler.db.Raw(s, arguments).Scan(scan).Error; err != nil {
		return nil, err
	}
	return scan, nil
}

func (handler *SqlHandler) Row(row string, where interface{}, scan interface{}) (interface{}, error) {
	if err := handler.db.Raw(row).Where(where).Scan(scan).Error; err != nil {
		return nil, err
	}
	return scan, nil
}

func (handler *SqlHandler) Update(obj interface{}) error {
	if err := handler.db.Save(obj).Error; err != nil {
		return err
	}
	return nil
}

func (handler *SqlHandler) DeleteById(obj interface{}, id string) {
	handler.db.Delete(obj, id)
}

func (handler *SqlHandler) DeleteOne(user domain.User) error {
	if err := handler.db.Delete(&user).Error; err != nil {
		return err
	}
	return nil
}

func (handler *SqlHandler) UpdateName(user domain.User) error {
	fmt.Println(user)
	if err := handler.db.Model(&user).Update("name", user.Name).Error; err != nil {
		return err
	}
	return nil
}

func (handler *SqlHandler) GetIdMailNamePasswordByMail(mail string) (domain.User, error) {
	user := domain.User{}
	if err := handler.db.Select("id", "mail", "name", "password").Where("mail = ?", mail).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (handler *SqlHandler) GetPasswordAndId(mail string) (domain.User, error) {
	user := domain.User{}
	if err := handler.db.Select("password, id").Where("mail = ?", mail).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (handler *SqlHandler) GetAllPrefectures() ([]domain.Prefecture, error) {
	prefecture := []domain.Prefecture{}
	if err := handler.db.Find(&prefecture).Error; err != nil {
		return prefecture, err
	}
	fmt.Println(prefecture)
	return prefecture, nil
}

func (handler *SqlHandler) GetAllCities() ([]domain.Citie, error) {
	cities := []domain.Citie{}
	if err := handler.db.Find(&cities).Error; err != nil {
		return cities, err
	}
	fmt.Println(cities)
	return cities, nil
}

func (handler *SqlHandler) GetAllDetailedCities() ([]domain.DetailedCitie, error) {
	detailedCities := []domain.DetailedCitie{}
	if err := handler.db.Find(&detailedCities).Error; err != nil {
		return detailedCities, err
	}
	fmt.Println(detailedCities)
	return detailedCities, nil
}

func (handler *SqlHandler) GetPasswordByMail(mail string) (string, error) {
	user := domain.User{}
	if err := handler.db.Select("password").Where("mail = ?", mail).First(&user).Error; err != nil {
		return "", err
	}
	return user.Password, nil
}

func (handler *SqlHandler) IsRecordExisting(obj interface{}, where interface{}) (bool, error) {
	result := handler.db.Where(where).First(obj)
	fmt.Println(result)
	isExist := result.RowsAffected
	fmt.Println(isExist)
	fmt.Println(obj)
	if isExist < 1 {
		return false, nil
	}
	return true, nil
}
