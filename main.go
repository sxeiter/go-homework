package main

import "fmt"

func main() {
	const USD float64 = 1
	const USDtoEUR float64 = USD * 1.1
	const USDtoRUB float64 = USD * 80
	const EURtoRUB float64 = USDtoRUB / USDtoEUR
}

func inputData() float64 {
	var userInput float64
	fmt.Println("Введите какое то значение")
	fmt.Scan(&userInput)
	return userInput
}

func calculate(money float64, usd float64, usdConvertTo float64) {

}
