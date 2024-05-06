package main

import (
	"errors"
	"fmt"
)

type myFirstError struct {
	msg string
}

func (e myFirstError) Error() string {
	return e.msg
}

func main() {
	firstErr := myFirstError{msg: "Случилась моя первая ошибка"}

	err1 := errors.New("ошибка1")
	err2 := fmt.Errorf("ошибка2:%w", err1)
	err3 := fmt.Errorf("ошибка3:%w", err2)
	fmt.Println(err3)

	if errors.As(err3, &firstErr) {
		fmt.Printf("Цепочка ошибок содержит ошибку типа %T\n", firstErr)
		fmt.Println(firstErr)
	} else {
		fmt.Printf("Цепочка ошибок не содержит ошибку типа %T\n", firstErr)
	}
}
