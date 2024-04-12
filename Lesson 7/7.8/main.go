package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(slice)

	index := 3
	// Срезирование среза для удаления элемента
	slice = append(slice[:index], slice[index+1:]...)
	fmt.Println(slice)
}
