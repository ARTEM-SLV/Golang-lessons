package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	err := checkInt1(-1)
	if err != nil {
		log.Fatal(err)
	}
}

func checkInt1(v int) error {
	err := checkInt2(v)
	if err != nil {
		return fmt.Errorf("ошибка1:%w", err)
	}

	return nil
}

func checkInt2(v int) error {
	var err error

	err = checkInt3(v)
	if err != nil {
		return fmt.Errorf("ошибка2:%w", err)
	}

	return nil
}

func checkInt3(v int) error {
	if v < 0 {
		return errors.New("ошибка3")
	}

	return nil
}
