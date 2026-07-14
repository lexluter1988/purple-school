package main

import "fmt"

const usd_eur float64 = 0.88
const usd_rub float64 = 76.94

func main() {
	var eur_rub float64 = usd_rub / usd_eur
	fmt.Println("EUR to RUB =", eur_rub)
}
