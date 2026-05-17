package main

import "fmt"

func main() {
	const usdToEur = 0.92
	const usdToRub = 81.50

	eurToRub := usdToRub / usdToEur

	fmt.Println(eurToRub)
}
