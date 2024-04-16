package main

import "fmt"

func main() {
	animals := map[string]struct{}{
		"слон":    {},
		"бегемот": {},
		"носорог": {},
		"лев":     {},
	}

	delete(animals, "бегемот")
	fmt.Println(animals)
}
