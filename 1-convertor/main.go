package main

import (
	"fmt"
	"strings"
)

var rates = map[string]map[string]float64{
	"USD": {"EUR": 0.92, "RUB": 90.5, "GBP": 0.79, "CNY": 7.24, "USD": 1},
	"EUR": {"USD": 1.09, "RUB": 98.3, "GBP": 0.86, "CNY": 7.87, "EUR": 1},
	"RUB": {"USD": 0.011, "EUR": 0.010, "GBP": 0.0087, "CNY": 0.080, "RUB": 1},
	"GBP": {"USD": 1.27, "EUR": 1.16, "RUB": 115.0, "CNY": 9.16, "GBP": 1},
	"CNY": {"USD": 0.138, "EUR": 0.127, "RUB": 12.5, "GBP": 0.109, "CNY": 1},
}

var supportedCurrencies = []string{"USD", "EUR", "RUB", "GBP", "CNY"}

func readCurrency(prompt string) string {
	hints := strings.Join(supportedCurrencies, ", ")
	for {
		fmt.Printf("%s (%s): ", prompt, hints)
		var input string
		fmt.Scan(&input)
		input = strings.ToUpper(strings.TrimSpace(input))
		for _, c := range supportedCurrencies {
			if c == input {
				return input
			}
		}
		fmt.Printf("Ошибка: валюта '%s' не поддерживается. Попробуйте снова.\n", input)
	}
}

func readAmount() float64 {
	for {
		fmt.Print("Введите сумму: ")
		var amount float64
		_, err := fmt.Scan(&amount)
		if err != nil || amount <= 0 {
			fmt.Println("Ошибка: введите положительное число. Попробуйте снова.")
			fmt.Scanln()
			continue
		}
		return amount
	}
}

func main() {
	fmt.Println("=== Конвертер валют ===")
	fmt.Println()

	from := readCurrency("Исходная валюта")
	amount := readAmount()
	to := readCurrency("Целевая валюта")

	result := convert(amount, from, to)

	fmt.Printf("\n%.2f %s = %.2f %s\n", amount, from, result, to)
}

func convert(amount float64, from, to string) float64 {
	switch from {
	case "USD", "EUR", "RUB", "GBP", "CNY":
		if rate, ok := rates[from][to]; ok {
			return amount * rate
		}
	}
	return 0
}
