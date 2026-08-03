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

func (acc account) outputPassword() {
	fmt.Printf("Ваш логин: %s, ваш пароль: %s, ваш URL: %s", acc.login, acc.password, acc.url)
}

var letterRunes = []rune("abcdefghijklmnoprstuvwxyzABCDEFGHIJKLMNOPRSTUVWXYZ0123456789-*")

func main() {
	fmt.Println(generatePassword(12))
	login := promptData("Введите логин: ")
	password := promptData("Введите пароль: ")
	url := promptData("Введите URL: ")

	account := account{
		password: password,
		login:    login,
		url:      url,
	}

	account.outputPassword()
}

func promptData(prompt string) string {
	fmt.Print(prompt)
	var res string
	fmt.Scan(&res)
	return res
}

func generatePassword(n int) string {
	res := make([]rune, n)
	for i := range res {
		res[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(res)
}
