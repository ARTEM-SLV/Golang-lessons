package main

import (
	"fmt"
)

func main() {
	a := 1
	do(a)
}
func do(v any) {
	a := 10
	fmt.Println(a)

	if intV, ok := v.(int); ok {
		a += intV
		fmt.Println("Результат сложения:", a)
	} else {
		fmt.Println("Не удалось извлечь значение типа int из интерфейса")
	}

}
