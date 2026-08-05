package main

import "fmt"

type Car struct {
	Brand        string
	Model        string
	Mileage      int
	PricePerHour float64
}

func PrintCarInfo(c Car) {
	fmt.Println("Brand:", c.Brand)
	fmt.Println("Model:", c.Model)
	fmt.Printf("Mileage: %d km\n", c.Mileage)
	fmt.Printf("Price: %.1f per hour", c.PricePerHour)
}

func main() {
	c := Car{
		Brand:        "Toyota",
		Model:        "Camry",
		Mileage:      45000,
		PricePerHour: 12.5,
	}
	PrintCarInfo(c)
}
