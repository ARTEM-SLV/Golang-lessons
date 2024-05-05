package main

import (
	"errors"
	"fmt"
)

type ErrOne error
type ErrTwo error

func (e ErrOne) Error() string {
	return "ошибка1"
}

//func (e ErrTwo) Error() string {
//	return "ошибка2"
//}

func main() {
	var err ErrOne
	err := errors.New("")

	fmt.Println(err.Error())
}

//func checkInt1(v int) error {
//	var errData errOne
//	errData.err = checkInt2(v)
//	if errData.err != nil {
//		return fmt.Errorf("ошибка1:%w", errData.err)
//	}
//
//	return nil
//}
//
//func checkInt2(v int) error {
//	var errData errTwo
//
//	errData.err = checkInt3(v)
//	if errData.err != nil {
//		return fmt.Errorf("ошибка2:%w", errData.err)
//	}
//
//	return nil
//}
//
//func checkInt3(v int) error {
//	if v < 0 {
//		return errors.New("ошибка3")
//	}
//
//	return nil
//}
