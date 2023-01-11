package database

import (
	"github.com/Kantaro0829/clean-architecture-in-go/domain"
)

type SqlHandler interface {
	Create(object interface{}) error
	FindAll(object interface{}) (interface{}, error)
	FindOne(obj interface{}, columns []string, where interface{}) (interface{}, error)
	RowSql(sql string, args interface{}, result interface{}) (interface{}, error)
	IsRecordExisting(obj interface{}, where interface{}) (bool, error)

	Update(object interface{}) error
	DeleteById(object interface{}, id string)
	UpdateName(user domain.User) error
	GetPasswordByMail(mail string) (string, error)
	GetIdMailNamePasswordByMail(mail string) (domain.User, error)
	GetPasswordAndId(mail string) (domain.User, error)
	DeleteOne(user domain.User) error
	GetAllPrefectures() ([]domain.Prefecture, error)
	GetAllCities() ([]domain.Citie, error)
	GetAllDetailedCities() ([]domain.DetailedCitie, error)

	Row(row string, where interface{}, scan interface{}) (interface{}, error)
}
