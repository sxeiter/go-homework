package main

import (
	"fmt"
)

const USD float64 = 1

var currencyRates = map[string]*float64{
	"USD": new(float64),
	"EUR": new(float64),
	"RUB": new(float64),
}

func init() {
	*currencyRates["USD"] = 1
	*currencyRates["EUR"] = 0.9
	*currencyRates["RUB"] = 80
}

func main() {
	fromCurrency, amount, toCurrency := getUserInput()
	result := calculate(amount, fromCurrency, toCurrency)
	outputResult(result, toCurrency)
}

func getUserInput() (string, float64, string) {
	var fromCurrency string
	var amount float64
	var toCurrency string
	for {
		fmt.Println("Введите исходную валюту (USD, RUB, EUR):")
		fmt.Scan(&fromCurrency)
		if isValidCurrency(fromCurrency) {
			break
		} else {
			fmt.Println("Неверная валюта. Пожалуйста, попробуйте снова.")
		}
	}
	for {
		fmt.Println("Введите количество денег:")
		_, err := fmt.Scan(&amount)
		if err == nil && amount > 0 {
			break
		} else {
			fmt.Println("Неверное число. Пожалуйста, введите положительное число.")
			var temp string
			fmt.Scan(&temp)
		}
	}
	for {
		fmt.Println("Введите валюту, в которую будем конвертировать (USD, RUB, EUR):")
		fmt.Scan(&toCurrency)
		if isValidCurrency(toCurrency) {
			break
		} else {
			fmt.Println("Неверная валюта. Пожалуйста, попробуйте снова.")
		}
	}
	return fromCurrency, amount, toCurrency
}

func isValidCurrency(currency string) bool {
	return currency == "USD" || currency == "RUB" || currency == "EUR"
}

func calculate(amount float64, fromCurrency string, toCurrency string) float64 {
	if fromCurrency == toCurrency {
		return amount
	}
	fromRate := *currencyRates[fromCurrency]
	toRate := *currencyRates[toCurrency]
	return amount * (toRate / fromRate)
}

func outputResult(result float64, currency string) {
	fmt.Printf("Результат: %.2f %s\n", result, currency)
}
