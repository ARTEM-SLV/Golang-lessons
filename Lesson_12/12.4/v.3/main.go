package main

import (
	"errors"
	"fmt"
)

type Bird interface {
	Sing() string
}
type Duck struct {
	voice string
}

func (d *Duck) Sing() string {
	return d.voice
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
	// явно полусчить тип Duck и првоерять его на nil
	d := b.(*Duck)
	if d != nil {
		return d.Sing(), nil
	}
	return "", errors.New("Ошибка пения!")
}
