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
	jsonStr := `{"number":1,"landlord":"ОстапБендер","tenat":"Шура Балаганов»}`

	c := contract{}
	err := json.Unmarshal([]byte(jsonStr), &c)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v\n", c)
}
