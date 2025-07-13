package main

import (
	"fmt"
)

const (
	USD      float64 = 1
	USDtoEUR float64 = USD * 0.9
	USDtoRUB float64 = USD * 80
	EURtoRUB float64 = USDtoRUB / USDtoEUR
)

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
	switch fromCurrency {
	case "USD":
		switch toCurrency {
		case "EUR":
			return amount * USDtoEUR
		case "RUB":
			return amount * USDtoRUB
		}
	case "EUR":
		switch toCurrency {
		case "USD":
			return amount / USDtoEUR
		case "RUB":
			return amount * EURtoRUB
		}
	case "RUB":
		switch toCurrency {
		case "USD":
			return amount / USDtoRUB
		case "EUR":
			return amount / EURtoRUB
		}
	}
	return 0
}

func outputResult(result float64, currency string) {
	fmt.Printf("Результат: %.2f %s\n", result, currency)
}
