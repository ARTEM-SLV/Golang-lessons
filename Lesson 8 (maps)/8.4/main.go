package main

import "fmt"

func main() {
	animals := map[string]struct{}{
		"слон":    {},
		"бегемот": {},
		"носорог": {},
		"лев":     {},
	}

	animals["выдра"] = struct{}{}

	fmt.Println(animals)
}
