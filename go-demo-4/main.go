package main

import "fmt"

type account struct {
	login    string
	password string
	url      string
}

func main() {
	login := promptData("Введите логин: ")
	password := promptData("Введите пароль: ")
	url := promptData("Введите URL: ")

	account := account{
		password: password,
		login:    login,
		url:      url,
	}

	outputPassword(&account)
}

func promptData(prompt string) string {
	fmt.Print(prompt)
	var res string
	fmt.Scan(&res)
	return res
}

func outputPassword(acc *account) {
	fmt.Printf("Ваш логин: %s, ваш пароль: %s, ваш URL: %s", (*acc).login, acc.password, acc.url)
}
