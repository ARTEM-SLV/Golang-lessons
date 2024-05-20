package main

import (
	format "github.com/ARTEM-SLV/format_"
	format2 "github.com/ARTEM-SLV/format_/v2"
	"log"
)

func main() {
	err := format.Do(`.\animals.json`, `.`)
	if err != nil {
		log.Println(err)
	}

	err = format2.Do(`.\animals.xml`, `.\`)
	if err != nil {
		log.Println(err)
	}
}
