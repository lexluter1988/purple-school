package main

import (
	"fmt"
)

const usd_eur float64 = 0.88
const usd_rub float64 = 76.94

func main() {
	amount, currency_from, currency_to := getInput()
	calcExchange(amount, currency_from, currency_to)

}

func getInput() (amount float64, currency_from string, currency_to string) {
	fmt.Print("Введите количество валюты: ")
	fmt.Scan(&amount)
	fmt.Print("Введите название ваше валюты: ")
	fmt.Scan(&currency_from)
	fmt.Print("Введите название желаемой валюты: ")
	fmt.Scan(&currency_to)
	return
}

func calcExchange(amount float64, currency_from string, currency_to string) float64 {
	return 0.0
}
