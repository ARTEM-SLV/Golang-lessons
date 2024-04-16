package main

import "fmt"

func main() {
	animals := map[string]int{
		"слон":    3,
		"бегемот": 0,
		"носорог": 5,
		"лев":     1,
	}

	printAnimal(animals, "слон")
	printAnimal(animals, "бегемот")
	printAnimal(animals, "осьминог")
}

func printAnimal(animals map[string]int, key string) {
	v, ok := animals[key]
	fmt.Printf("Животное: %s, количество: %d (есть в карте: %v).\n", key, v, ok)
}
