package main

import (
	"errors"
	"fmt"
)

const usd_eur float64 = 0.88
const usd_rub float64 = 76.94

func main() {
	currency_from, amount, currency_to := getInput()
	r, err := calcExchange(currency_from, amount, currency_to)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("result ", r)

}

func getInput() (currency_from string, amount float64, currency_to string) {
	step := 0
	for {
		switch step {
		case 0:
			fmt.Print("Введите название вашей валюты: ")
			fmt.Scan(&currency_from)
			r, err := validateSourceCurrency(currency_from)
			if !r {
				fmt.Println(err)
				continue
			}
			step += 1
		case 1:
			fmt.Print("Введите количество валюты: ")
			fmt.Scan(&amount)
			r, err := validateAmount(amount)
			if !r {
				fmt.Println(err)
				continue
			}
			step += 1
		case 2:
			fmt.Print("Введите название желаемой валюты: ")
			fmt.Scan(&currency_to)
			r, err := validateDestinationCurrency(currency_from, currency_to)
			if !r {
				fmt.Println(err)
				continue
			}
			step += 1
		}
		if step == 3 {
			break
		}
	}
	return
}

func validateSourceCurrency(currency string) (bool, error) {
	if currency != "USD" && currency != "EUR" && currency != "RUB" {
		return false, errors.New("Не поддерживаемая валюта (USD|RUB|EUR)")
	}
	return true, nil
}

func validateDestinationCurrency(currency_from string, currency_to string) (bool, error) {
	message := ""
	if currency_to != "USD" && currency_to != "EUR" && currency_to != "RUB" {
		return false, errors.New("Не поддерживаемая валюта (USD|RUB|EUR)")
	} else if currency_from == currency_to {
		switch currency_from {
		case "USD":
			message = "Не выберите другую валюту (RUB|EUR)"
		case "EUR":
			message = "Не выберите другую валюту (RUB|USD)"
		case "RUB":
			message = "Не выберите другую валюту (EUR|USD)"
		}
		return false, errors.New(message)
	}
	return true, nil
}

func validateAmount(amount float64) (bool, error) {
	if amount <= 0 {
		return false, errors.New("Количество валюты не может быть негативным")
	}
	return true, nil
}

func calcExchange(currency_from string, amount float64, currency_to string) (float64, error) {
	var eur_rub float64 = usd_rub / usd_eur
	switch currency_from {
	case "USD":
		switch currency_to {
		case "RUB":
			return amount * usd_rub, nil
		case "EUR":
			return amount * usd_eur, nil
		}
	case "EUR":
		switch currency_to {
		case "RUB":
			return amount * eur_rub, nil
		case "USD":
			return amount / usd_eur, nil
		}
	case "RUB":
		switch currency_to {
		case "USD":
			return amount / usd_rub, nil
		case "EUR":
			return amount / eur_rub, nil
		}
	}
	return 0, errors.ErrUnsupported
}
