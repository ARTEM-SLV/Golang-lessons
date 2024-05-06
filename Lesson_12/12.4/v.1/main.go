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
	// ошибка из за того, что переменная объявлена как указатель (var d *Duck)
	// в таком случае создается переменная d имеет значение nil
	// затем в методе Sing() структуры Duck возвращаем занчение d.voice, что приводит к ошибки
	// вместо создания указателя создаю структуру Duck
	var d Duck //var d *Duck
	d.voice = "кря-кря-кря"

	song, err := Sing(&d)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(song)
}

func Sing(b Bird) (string, error) {
	if b != nil {
		return b.Sing(), nil
	}
	return "", errors.New("Ошибка пения!")
}
