package main

import "fmt"

func main() {
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
