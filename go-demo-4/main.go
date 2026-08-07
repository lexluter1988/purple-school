package main

import (
	"demo/password/account"
	"fmt"
)

func main() {
	login := promptData("Введите логин: ")
	password := promptData("Введите пароль: ")
	url := promptData("Введите URL: ")

	account, err := account.NewAccount(login, password, url)
	if err != nil {
		fmt.Println(err)
		return
	}
	account.OutputPassword()
	fmt.Println(account)
}

func promptData(prompt string) string {
	fmt.Print(prompt)
	var res string
	fmt.Scanln(&res)
	return res
}
