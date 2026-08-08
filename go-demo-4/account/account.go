package account

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"net/url"
	"time"
)

var letterRunes = []rune("abcdefghijklmnoprstuvwxyzABCDEFGHIJKLMNOPRSTUVWXYZ0123456789-*")

type Account struct {
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Login     string    `json:"login"`
	Password  string    `json:"password"`
	Url       string    `json:"url"`
}

func (acc *Account) OutputPassword() {
	fmt.Printf("Ваш логин: %s, ваш пароль: %s, ваш URL: %s", acc.Login, acc.Password, acc.Url)
}

func (acc *Account) ToBytes() ([]byte, error) {
	file, err := json.Marshal(acc)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func (acc *Account) generatePassword(n int) {
	res := make([]rune, n)
	for i := range res {
		res[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	acc.Password = string(res)
}

func NewAccount(login, password, urlString string) (*Account, error) {
	if login == "" {
		return nil, errors.New("Неверный login")
	}
	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, errors.New("Неверный формат URL")
	}
	createdAt := time.Now()
	updatedAt := time.Now()

	acc := &Account{
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		Login:     login,
		Url:       urlString,
		Password:  password,
	}
	if acc.Password == "" {
		acc.generatePassword(12)
	}
	return acc, nil
}
