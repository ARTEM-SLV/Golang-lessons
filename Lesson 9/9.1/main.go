package main

import "fmt"

func main() {
	fruit := ",бананы"
	count := fruitMarket(fruit)
	if count != -1 {
		fmt.Println("Количесвто", fruit, "=", count)
	}
}

func fruitMarket(fruit string) int {
	fruits := map[string]int{
		"апельсины": 5,
		"яблоки":    3,
		"сливы":     1,
		"груши":     0,
	}

	res, ok := fruits[fruit]
	if !ok {
		fmt.Printf("Фрукт %s отсутсвует в магазине", fruit)
		res = -1
	}

	return res
}
