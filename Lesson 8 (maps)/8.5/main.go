package main

import "fmt"

func main() {
	animals := map[string]int{
		"слон":    3,
		"бегемот": 0,
		"носорог": 5,
		"лев":     1,
	}

	animals["бегемот"] = 2
	fmt.Println(animals)
}
