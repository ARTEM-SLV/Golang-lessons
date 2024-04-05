package main

import "fmt"

func main() {
	const (
		const1 = 1 << iota
		const2
		const3
		const4
		const5
	)

	fmt.Println("Значение константы const1:", const1)
	fmt.Println("Значение константы const2:", const2)
	fmt.Println("Значение константы const3:", const3)
	fmt.Println("Значение константы const4:", const4)
	fmt.Println("Значение константы const5:", const5)
}
