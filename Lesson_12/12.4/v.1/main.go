package main

import (
	"errors"
	"fmt"
)

// Ошибка из за того, что переменная объявлена как указатель (var d *Duck)
// в таком случае создается переменная d имеет значение nil
// затем в методе Sing() структуры Duck возвращаем занчение d.voice
// это приводит к ошибки, так как в методе Sing b <> nil,
// так как нулевое значение для интерфейса имеет и тип, и значение равное nil
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
	// ++ описание ошибки b не = nil, так как тип не nil
	fmt.Printf("Тип: %T, значение: %v", b, b)
	// -- описание ошибки

	// явно полусчить тип Duck и првоерять его на nil
	d := b.(*Duck)
	if d != nil {
		return d.Sing(), nil
	}
	return "", errors.New("Ошибка пения!")
}
