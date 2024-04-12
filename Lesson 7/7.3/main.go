package main

import "fmt"

func main() {
	//вариант 1
	fmt.Println("variant 1")
	var arr = [4]string{"яблоко", "груша", "помидор", "абрикос"}
	fmt.Println(arr)

	arr[2] = "персик"
	fmt.Println(arr)

	//вариант 2
	fmt.Println("variant 2")
	var arr2 = [4]string{"яблоко", "груша", "помидор", "абрикос"}
	fmt.Println(arr2)
	for i := range arr2 {
		if arr2[i] == "помидор" {
			arr2[i] = "персик"
		}
	}
	fmt.Println(arr2)
}
