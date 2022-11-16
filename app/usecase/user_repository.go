package usecase

import (
	"github.com/Kantaro0829/clean-architecture-in-go/domain"
)

type UserRepository interface {
	SignUp(u domain.User) error
	SelectAll() ([]domain.User, error)
	Delete(id string)
	DeleteByMail(user domain.User) error
	Update(u domain.User, name string)
	UpdateByMail(user domain.User) error
	GetPassword(mail string) (string, error)
	GetPasswordForUpdate(mail string) (domain.User, error)
	//以下追加
	GetMailNamePasswordByMail(mail string) (domain.User, error)

	NameAndPassword(mail string) (domain.NameAndPassword, error)
}
