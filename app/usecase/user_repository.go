package usecase

import (
	"github.com/Kantaro0829/clean-architecture-in-go/domain"
)

type UserRepository interface {
	SignUp(u domain.User) error
	SelectAll() ([]domain.User, error)
	UpdateOne(obj domain.User) error
	Delete(id string)
	//Update(u domain.User, name string)
	//UpdateByMail(user domain.User) error
	GetPassword(mail string) (string, int, error)
	//以下追加
	GetMailNamePasswordByMail(mail string) (domain.User, error)

	NameAndPassword(mail string) (domain.NameAndPassword, error)
}
