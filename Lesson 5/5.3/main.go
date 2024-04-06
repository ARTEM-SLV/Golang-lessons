package main

import "fmt"

func main() {
	var a int
	var p *int

	p = &a

	a = 10
	fmt.Printf("Значение а до изменения %d\n", a)

	*p = 11
	fmt.Printf("Значение а после изменения %d", a)
}
