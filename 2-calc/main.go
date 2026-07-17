package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	operation, numbers := getInput()
	result := calculate(operation, numbers)
	fmt.Printf("Результат операции %s = %.2f ", operation, result)
}

func getInput() (operation string, numbers []float64) {
	step := 0
	for {
		switch step {
		case 0:
			fmt.Print("Выберите операцию (AVG-среднее, SUM-сумма, MED-медиана): ")
			fmt.Scan(&operation)
			r, err := validateOperation(operation)
			if !r {
				fmt.Println(err)
				continue
			}
			step += 1
		case 1:
			scanner := bufio.NewScanner(os.Stdin)
			fmt.Print("Введите числа через запятую: ")
			scanner.Scan()
			numbers_str := scanner.Text()
			r, err := validateArray(numbers_str)
			if err != nil {
				fmt.Println(err)
				continue
			}
			step += 1
			numbers = r
		}
		if step == 2 {
			break
		}
	}
	return
}

func validateOperation(operation string) (bool, error) {
	if operation != "AVG" && operation != "SUM" && operation != "MED" {
		return false, errors.New("Не операция (AVG|SUM|MED)")
	}
	return true, nil
}

func validateArray(numbers_str string) ([]float64, error) {
	parts := strings.Split(numbers_str, ",")
	result := []float64{}

	for _, part := range parts {
		cleaned := strings.TrimSpace(part)
		num, err := strconv.ParseFloat(cleaned, 64)
		if err != nil {
			return nil, errors.New("Не смогли определить число, попробуйте снова")
		}
		result = append(result, num)
	}
	return result, nil
}

func calculate(operation string, numbers []float64) (result float64) {
	switch operation {
	case "AVG":
		for _, value := range numbers {
			result += value
		}
		result = result / float64(len(numbers))
	case "SUM":
		for _, value := range numbers {
			result += value
		}
	case "MED":
		sorted := numbers
		slices.Sort(sorted)
		mid := len(sorted) / 2

		if len(sorted)%2 == 0 {
			result = sorted[mid-1]
		} else {
			result = sorted[mid]
		}
	}
	return
}
