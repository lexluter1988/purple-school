package main

import (
	"errors"
	"fmt"
	"math/rand"
	"net/url"
	"time"
)

type account struct {
	login    string
	password string
	url      string
}

type accountWtimestamp struct {
	createdAt time.Time
	updatedAt time.Time
	account
}

func (acc *account) outputPassword() {
	fmt.Printf("Ваш логин: %s, ваш пароль: %s, ваш URL: %s", acc.login, acc.password, acc.url)
}

func (acc *account) generatePassword(n int) {
	res := make([]rune, n)
	for i := range res {
		res[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	acc.password = string(res)
}

func newAccount(login, password, urlString string) (*accountWtimestamp, error) {
	if login == "" {
		return nil, errors.New("Неверный login")
	}
	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, errors.New("Неверный формат URL")
	}
	createdAt := time.Now()
	updatedAt := time.Now()

	acc := &accountWtimestamp{
		createdAt: createdAt,
		updatedAt: updatedAt,
		account: account{
			login:    login,
			url:      urlString,
			password: password,
		},
	}
	if acc.password == "" {
		acc.generatePassword(12)
	}
	return acc, nil
}

var letterRunes = []rune("abcdefghijklmnoprstuvwxyzABCDEFGHIJKLMNOPRSTUVWXYZ0123456789-*")

func main() {
	login := promptData("Введите логин: ")
	password := promptData("Введите пароль: ")
	url := promptData("Введите URL: ")

	account, err := newAccount(login, password, url)
	if err != nil {
		fmt.Println(err)
		return
	}
	account.outputPassword()
	fmt.Println(account)
}

func promptData(prompt string) string {
	fmt.Print(prompt)
	var res string
	fmt.Scanln(&res)
	return res
}
