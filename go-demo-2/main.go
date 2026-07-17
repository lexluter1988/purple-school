package main

import "fmt"

func main() {
	transactions := make([]string, 0, 2)
	fmt.Println(transactions, cap(transactions))
	transactions = append(transactions, "1")
	fmt.Println(transactions, cap(transactions))
	transactions = append(transactions, "2")
	fmt.Println(transactions, cap(transactions))
	transactions = append(transactions, "3")
	fmt.Println(transactions, cap(transactions))
	fmt.Println(transactions)

	tr := []float64{}
	for {
		transaction := getInputTransaction()
		if transaction == 0 {
			fmt.Println(tr)
			break
		}
		tr = append(tr, transaction)
	}
	balance := calculateBalance(tr)
	fmt.Printf("Ваш баланс: %.2f", balance)
}

func getInputTransaction() float64 {
	var transaction float64
	fmt.Println("Enter the transaction or n to exit: ")
	fmt.Scan(&transaction)
	return transaction
}

func calculateBalance(transactions []float64) float64 {
	balance := 0.0
	for _, value := range transactions {
		balance += value
	}
	return balance
}
