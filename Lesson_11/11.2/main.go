package main

import (
	"errors"
	"fmt"
)

func main() {
	// Создаем цепочку ошибок
	err1 := errors.New("ошибка1")
	err2 := fmt.Errorf("ошибка2: %w", err1)
	err3 := fmt.Errorf("ошибка3: %w", err2)

	fmt.Println(err3)

	// Получаем внутреннюю ошибку
	err := errors.Unwrap(err3)
	if err != nil {
		fmt.Println(err)
	}
}
