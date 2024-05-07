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
	strJSON := `{"number":1,"landlord":"ОстапБендер","tenat":"Шура Балаганов»}`

	c := contract{}
	err := json.Unmarshal([]byte(strJSON), &c)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", c)
}
