package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(slice)

	index := 3
	// Срезирование среза для удаления элемента
	slice = deleteElement(slice, index)
	fmt.Println(slice)
}

func deleteElement(slice []int, index int) []int {
	return append(slice[:index], slice[index+1:]...)
}
