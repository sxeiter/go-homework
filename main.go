package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Используйте: go run main.go<Операция(нужно выбрать одну из AVG, SUM или MED)><Числа через запятую>")
		return
	}
	operation := os.Args[1]
	numbersStr := os.Args[2]

	numberString := strings.Split(numbersStr, ",")
	var numbers []float64
	for _, numStr := range numberString {
		num, err := strconv.ParseFloat(strings.TrimSpace(numStr), 64)
		if err != nil {
			fmt.Println("Ошибка")
			return
		}
		numbers = append(numbers, num)
	}
	switch operation {
	case "AVG":
		fmt.Printf("Среднее: %.2f\n", AVG(numbers))
	case "SUM":
		fmt.Printf("Сумма: %.2f\n", SUM(numbers))
	case "MED":
		fmt.Printf("Медиана: %.2f\n", MED(numbers))
	default:
		fmt.Println("Недопустимая операция. Используйте AVG, SUM или MED.")
	}
}

func AVG(numbers []float64) (float64, error) {
	if len(numbers) == 0 {
		return 0, fmt.Errorf("вы ничего не передали")
	}
	sum := 0.0
	for _, num := range numbers {
		sum += num
	}
	return sum / float64(len(numbers)), nil
}

func SUM(numbers []float64) float64 {
	total := 0.0
	for _, num := range numbers {
		total += num
	}
	return total
}

func MED(numbers []float64) float64 {
	n := len(numbers)
	if n == 0 {
		return 0
	}
	sort.Float64s(numbers)
	if n%2 == 0 {
		return (numbers[n/2-1] + numbers[n/2]) / 2
	}
	return numbers[n/2]
}
