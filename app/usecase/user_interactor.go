package usecase

import (
	"fmt"

	"github.com/Kantaro0829/clean-architecture-in-go/domain"
	//"golang.org/x/crypto/bcrypt"
)

type UserInteractor struct {
	UserRepository UserRepository
	//以下試しに追加
	Tokenizer Tokenizer
}

func (interactor *UserInteractor) Add(u domain.User) error {
	//interactor.UserRepository.Store(u)
	err := interactor.UserRepository.SignUp(u)
	return err
}

func (interactor *UserInteractor) GetInfo() ([]domain.User, error) {
	obj, err := interactor.UserRepository.SelectAll()
	return obj, err
}

// func (interactor *UserInteractor) Update(u domain.User, json domain.User) error {
// 	result := interactor.UserRepository.Update(u, json)
// 	return result
// }

func (interactor *UserInteractor) Login(mail string, password string) (domain.Token, bool, error) {
	//regedPassword, err := interactor.UserRepository.GetPassword(mail)
	var token domain.Token
	user, err := interactor.UserRepository.GetMailNamePasswordByMail(mail)
	if err != nil {
		return token, false, err
	}
	regedPassword := user.Password
	//パスワード比較
	isValidated := ValitatePassword(regedPassword, password)

	if !isValidated {
		//パスワード不一致の時
		return token, isValidated, nil
	}
	fmt.Printf("id:%v, password:%v, mail:%v, name:%v", user.ID, user.Password, user.Mail, user.Name)
	//パスワード一致の時JWT発行
	token, err = interactor.Tokenizer.New(user)
	if err != nil {
		return token, isValidated, nil
	}

	return token, isValidated, nil
}

func (interactor *UserInteractor) Authenticate(token domain.Token) (int, error) {
	id, err := interactor.Tokenizer.Verify(token)
	if err != nil {

		return 0, err
	}
	return id, nil
}

func (interactor *UserInteractor) UpdateUser(userJson domain.User) (string, error, bool) {
	mail := userJson.Mail
	password := userJson.Password
	//name := userJson.Name
	//user, err := interactor.UserRepository.GetPasswordForUpdate(mail)
	regedPass, regedId, err := interactor.UserRepository.GetPassword(mail)
	if err != nil || regedId == 0 {
		return "パスワードorメールアドレスの入力間違い", err, false
	}
	//regedPassword := user.Password

	//パスワード比較
	//isValidated := ValitatePassword(regedPassword, password)
	isValidated := ValitatePassword(regedPass, password)
	fmt.Println("パスワード比較完了")
	if isValidated != true {
		return "パスワードorメールアドレスの入力間違い", nil, false
	}

	userJson.ID = regedId
	isUpdated := interactor.UserRepository.UpdateOne(userJson)

	//result := interactor.UserRepository.UpdateByMail(user)
	if isUpdated != nil {
		return "データ書き込み失敗", isUpdated, true
	}
	return "登録完了", nil, true
}

func (interactor *UserInteractor) NamePassword(mail string) (domain.NameAndPassword, error) {
	result, err := interactor.UserRepository.NameAndPassword(mail)
	if err != nil {
		return domain.NameAndPassword{}, err
	}
	return result, err
}

func ValitatePassword(regedPassword string, password string) bool {
	if regedPassword == password {
		return true
	}
	return false
}
