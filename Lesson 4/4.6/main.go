package main

import "fmt"

func fibonacci(n, a, b int) {
	if n == 23 {
		return
	}

	fmt.Printf("%d ", a)
	fibonacci(n+1, b, a+b)
}

func main() {
	fmt.Println("Первые 23 числа Фибоначчи:")
	fibonacci(0, 0, 1)
}
