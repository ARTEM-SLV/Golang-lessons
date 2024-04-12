package main

import "fmt"

func main() {
	slice := make([]int, 0, 10)
	fmt.Println("slice:", slice, "длина:", len(slice), "ёмкость:", cap(slice))
}
