package main

import (
	"encoding/json"
	"fmt"
)

type contract struct {
	Number   int
	Landlord string
	Tenat    string
}

func main() {
	c := contract{
		Number:   2,
		Landlord: "Остап",
		Tenat:    "Шура",
	}
	fmt.Println(c)

	jsonStr, err := json.Marshal(c)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(jsonStr))
}
