package main

import (
	"encoding/xml"
	"fmt"
)

type Contract struct {
	Number   int    `xml:"number"`
	Landlord string `xml:"landlord"`
	Tenant   string `xml:"tenant"`
}

type Contracts struct {
	XMLName   xml.Name   `xml:"contracts"`
	Contracts []Contract `xml:"contract"`
}

func main() {
	c := Contract{
		Number:   1,
		Landlord: "Остап Бендер",
		Tenant:   "Шура Балаганов",
	}

	contracts := Contracts{
		Contracts: []Contract{c},
	}

	xmlData, err := xml.MarshalIndent(contracts, "", "    ")
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Println(string(xmlData))
}
