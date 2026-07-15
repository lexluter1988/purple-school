package main

import (
	"fmt"
	"math"
)

func main() {
	userHeight, userKg := getUserInput()
	IMT := calculateResult(userHeight, userKg)
	outputResult(IMT)

}

func outputResult(result float64) {
	out := fmt.Sprintf("Ваш индекс массы тела: %.0f\n", result)
	fmt.Print(out)
}

func calculateResult(userHeight, userKg float64) float64 {
	IMT := userKg / math.Pow(userHeight/100, 2)
	return IMT
}

func getUserInput() (float64, float64) {
	var userHeight, userKg float64
	fmt.Println("___Калькулятор индекса массы тела___")
	fmt.Println("Введите ваш рост в сантиметрах")
	fmt.Scan(&userHeight)
	fmt.Println("Введите ваш вес")
	fmt.Scan(&userKg)
	return userHeight, userKg
}
