package main

import "fmt"

func main() {
	const USD float64 = 1
	const USDtoEUR float64 = USD * 0.9
	const USDtoRUB float64 = USD * 80
	const EURtoRUB float64 = USDtoRUB / USDtoEUR
}

func inputData() float64 {
	var userInput float64
	fmt.Println("Введите какое то значение")
	fmt.Scan(&userInput)
	return userInput
}

func calculate(amount float64, fromCurrency string, toCurrency string) float64 {
	return 0
}
