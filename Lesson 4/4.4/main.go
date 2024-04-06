package main

import "fmt"

func main() {
	f := hell()
	f()
}

func hell() func() {
	f := func() {
		fmt.Println("hello, Go")
	}

	return f
}
