package main

import "fmt"

func readInput() float64 {
	var amount float64
	fmt.Scan(&amount)
	return amount
}

func convert(amount float64, fromCurrency string, toCurrency string) float64 {
	return 0
}

func main() {
	amount := readInput()
	result := convert(amount, "USD", "EUR")
	fmt.Println(result)
}
