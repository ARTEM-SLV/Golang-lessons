package main

import (
	"errors"
	"fmt"
	"reflect"
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
		fmt.Println("Цепочка ошибок содержит ошибку типа", reflect.TypeOf(firstErr))
		fmt.Println(firstErr)
	} else {
		fmt.Println("Цепочка ошибок не содержит ошибку типа", reflect.TypeOf(firstErr))
	}
}
