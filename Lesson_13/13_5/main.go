package main

import (
	format "github.com/ARTEM-SLV/format_"
	format2 "github.com/ARTEM-SLV/format_/v2"
	"log"
)

func main() {
	err := format.Do(`E:\Go\Go-Lessons\Lesson_13\13_5\animals.json`, `E:\Go\Go-Lessons\Lesson_13\13_5\`)
	if err != nil {
		log.Println(err)
	}

	err = format2.Do(`E:\Go\Go-Lessons\Lesson_13\13_5\animals.xml`, `E:\Go\Go-Lessons\Lesson_13\13_5\`)
	if err != nil {
		log.Println(err)
	}
}
