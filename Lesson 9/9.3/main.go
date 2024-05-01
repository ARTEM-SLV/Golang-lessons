package main

import "fmt"

func main() {
	food := "огурец"
	checkFood(food)
}

func checkFood(f string) {
	switch {
	case f == "груша" || f == "яблоко" || f == "апельсин":
		fmt.Println("это фрукт")
	case f == "тыква" || f == "огурец" || f == "помидор":
		fmt.Println("это овощ")
	default:
		fmt.Println("что-то странное...")
	}
}
