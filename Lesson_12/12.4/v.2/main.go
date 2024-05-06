package main

import (
	"errors"
	"fmt"
)

// реализовать проверку на nil в методе Sing
type Bird interface {
	Sing() (string, error)
}
type Duck struct {
	voice string
}

func (d *Duck) Sing() (string, error) {
	if d != nil {
		return d.voice, nil
	}
	return "", errors.New("Ошибка пения!")
}

func main() {
	var d *Duck

	song, err := Sing(d)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(song)
}

func Sing(b Bird) (string, error) {
	return b.Sing()
}
