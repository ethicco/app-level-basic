package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

func sum(nums []float64) float64 {
	total := 0.0

	for _, num := range nums {
		total += num
	}

	return total
}

func avg(nums []float64) float64 {
	return sum(nums) / float64(len(nums))
}

func median(nums []float64) float64 {
	sorted := make([]float64, len(nums))
	copy(sorted, nums)
	sort.Float64s(sorted)

	mid := len(sorted) / 2

	if len(sorted)%2 == 0 {
		return math.Round((sorted[mid-1]+sorted[mid])/2*100) / 100
	}

	return sorted[mid]
}

func main() {
	var operation string
	fmt.Print("Введите операцию (AVG, SUM, MED): ")
	fmt.Scan(&operation)
	operation = strings.ToUpper(strings.TrimSpace(operation))

	var input string
	fmt.Print("Введите числа через запятую: ")
	fmt.Scan(&input)

	parts := strings.Split(input, ",")
	nums := make([]float64, 0, len(parts))

	for _, p := range parts {
		p = strings.TrimSpace(p)

		if p == "" {
			continue
		}

		n, err := strconv.ParseFloat(p, 64)

		if err != nil {
			fmt.Printf("Ошибка: '%s' не является числом\n", p)
			return
		}

		nums = append(nums, n)
	}

	if len(nums) == 0 {
		fmt.Println("Ошибка: нет чисел для расчёта")
		return
	}

	switch operation {
	case "SUM":
		fmt.Printf("Сумма: %g\n", sum(nums))
	case "AVG":
		fmt.Printf("Среднее: %g\n", avg(nums))
	case "MED":
		fmt.Printf("Медиана: %g\n", median(nums))
	default:
		fmt.Printf("Неизвестная операция: %s\n", operation)
	}
}
