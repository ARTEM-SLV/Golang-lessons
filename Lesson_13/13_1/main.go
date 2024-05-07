package main

import (
	"encoding/json"
	"fmt"
)

type contract struct {
	Number   int    `json:"number"`
	Landlord string `json:"landlord"`
	Tenant   string `json:"tenant"`
}

func main() {
	var c contract
	jsonStr := `{"number":1,"landlord":"ОстапБендер","tenant":"Шура Балаганов"}`

	err := json.Unmarshal([]byte(jsonStr), &c)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Printf("%+v\n", c)
}
