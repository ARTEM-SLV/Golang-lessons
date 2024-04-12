package main

import "fmt"

func main() {
	slice := make([]int, 0, 10)
	fmt.Println("slice:", slice, "длина:", len(slice), "ёмкость:", cap(slice))

	slice = append(slice, 4, 1, 8, 9)

	fmt.Println("slice:", slice, "длина:", len(slice), "ёмкость:", cap(slice))
}
