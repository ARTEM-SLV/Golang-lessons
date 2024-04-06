package main

import "fmt"

func main() {
	var a int
	a = 2
	fmt.Println(a)

	change(&a)
	fmt.Println(a)
}

func change(v *int) {
	*v += 10
}
