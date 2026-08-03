package main

import (
	"fmt"
	"math/rand"
)

type account struct {
	login    string
	password string
	url      string
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

func newAccount(login, url string) *account {
	return &account{
		login: login,
		url:   url,
	}
}

var letterRunes = []rune("abcdefghijklmnoprstuvwxyzABCDEFGHIJKLMNOPRSTUVWXYZ0123456789-*")

func main() {
	login := promptData("Введите логин: ")
	// password := promptData("Введите пароль: ")
	url := promptData("Введите URL: ")

	account := *newAccount(login, url)
	account.generatePassword(12)
	account.outputPassword()
	fmt.Println(account)
}

func promptData(prompt string) string {
	fmt.Print(prompt)
	var res string
	fmt.Scan(&res)
	return res
}
