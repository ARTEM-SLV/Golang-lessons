package main

import "fmt"

func main() {
	dividend := 16
	divisor := 3

	result := dividend / divisor
	remainder := dividend % divisor

	fmt.Printf("Результат: %d, остаток от деления: %d, тип результата: %T\n", result, remainder, result)
}
