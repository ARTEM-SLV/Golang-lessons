package main

import "fmt"

func main() {
	// Создание карты с ключами и пустыми структурами в качестве значений
	animals := map[string]struct{}{
		"слон":    {},
		"бегемот": {},
		"носорог": {},
		"лев":     {},
	}
	
	fmt.Println("Содержимое карты:")
	for key := range animals {
		fmt.Println(key)
	}
}
