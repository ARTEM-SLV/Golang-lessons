package main

import "fmt"

func main() {
	slice1 := []int{1, 2, 3}
	slice2 := []int{4, 5, 6}
	slice3 := make([]int, 0, 3)

	slice3 = append(slice3, slice1...)
	slice3 = append(slice3, slice2...)

	fmt.Println(slice3)
}
