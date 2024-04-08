package main

import "fmt"

type contract struct {
	id           int
	Number, Date string
}

// func (c contract) String() string {
// 	return fmt.Sprintf("{ID:%d Number:%s Date:%s}",
// 		c.id, c.Number, c.Date)
// }

func main() {
	c := newContract(1, `#000A\n101`, "2024-01-31")
	fmt.Printf("%#v", c)
}

func newContract(id int, Number, Date string) contract {
	c := contract{
		id:     id,
		Number: Number,
		Date:   Date,
	}

	return c
}
