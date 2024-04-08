package main

import "fmt"

type contacts struct {
	Addss, Phone string
}

type user struct {
	ID   int
	Name string
	contacts
}

type employee struct {
	ID   int
	Name string
	contacts
}

func main() {
	u := user{
		ID:   1,
		Name: "Ivan",
		contacts: contacts{
			Addss: "Moscow",
			Phone: "123456",
		},
	}

	e := employee{
		ID:   1,
		Name: "Petr",
		contacts: contacts{
			Addss: "Saint Petersburg",
			Phone: "654321",
		},
	}

	fmt.Println(u.Addss, u.Phone, e.Addss, e.Phone)
}
