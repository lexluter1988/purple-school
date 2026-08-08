package main

import (
	"demo/password/account"
	"demo/password/files"
	"fmt"
)

func main() {
	createAccount()
}

func createAccount() {
	login := promptData("Введите логин: ")
	password := promptData("Введите пароль: ")
	url := promptData("Введите URL: ")

	account, err := account.NewAccount(login, password, url)
	if err != nil {
		fmt.Println(err)
		return
	}
	file, err := account.ToBytes()
	if err != nil {
		fmt.Println("Unable to marshal")
		return
	}
	files.WriteFile(file, "test.txt")
}

func promptData(prompt string) string {
	fmt.Print(prompt)
	var res string
	fmt.Scanln(&res)
	return res
}
