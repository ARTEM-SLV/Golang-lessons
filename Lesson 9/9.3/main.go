package main

import "fmt"

func main() {

}

func checkFood(food string) {
	switch food {
	case "груша":
		fmt.Println("это фрукт")
	case "яблоко":
		fmt.Println("это фрукт")
	case "апельсин":
		fmt.Println("это фрукт")
	case "тыква":
		fmt.Println("это овощ")
	case "огурец":
		fmt.Println("это овощ")
	case "помидор":
		fmt.Println("это овощ")
	default:
		fmt.Println("что-то странное...")
	}
}
