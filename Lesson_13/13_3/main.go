package main

import (
	"encoding/xml"
	"fmt"
)

type Contract struct {
	Number   int    `xml:"number"`
	SignDate string `xml:"sign_date"`
	Landlord string `xml:"landlord"`
	Tenant   string `xml:"tenant"`
}

type Contracts struct {
	XMLName   xml.Name   `xml:"contracts"`
	Contracts []Contract `xml:"contract"`
}

func main() {
	xmlData := `
	<contracts>
		<contract>
			<number>1</number>
			<sign_date>2023-09-02</sign_date>
			<landlord>Остап</landlord>
			<tenant>Шура</tenant>
		</contract>
		<contract>
			<number>2</number>
			<sign_date>2023-09-03</sign_date>
			<landlord>Бендер</landlord>
			<tenant>Балаганов</tenant>
		</contract>
	</contracts>`

	var contracts Contracts
	if err := xml.Unmarshal([]byte(xmlData), &contracts); err != nil {
		fmt.Println("Ошибка при разборе XML:", err)
		return
	}

	fmt.Println("Контракты:")
	for _, c := range contracts.Contracts {
		fmt.Printf("Номер: %d, Дата подписания: %s, Арендодатель: %s, Арендатор: %s\n",
			c.Number, c.SignDate, c.Landlord, c.Tenant)
	}
}
