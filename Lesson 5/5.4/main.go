package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var (
		str     string  = "Hello"
		num     int     = 42
		flt     float64 = 3.14
		boolean bool    = true
	)

	fmt.Printf("str:   value = %s, address = %p\n", str, &str)
	fmt.Printf("num:   value = %d, address = %p\n", num, &num)
	fmt.Printf("flt:   value = %f, address = %p\n", flt, &flt)
	fmt.Printf("array: value = %v, address = %p\n", boolean, &boolean)

	// Размер переменных в байтах
	fmt.Printf("Size of string:  %d bytes\n", unsafe.Sizeof(str))
	fmt.Printf("Size of int:     %d bytes\n", unsafe.Sizeof(num))
	fmt.Printf("Size of float64: %d bytes\n", unsafe.Sizeof(flt))
	fmt.Printf("Size of array:   %d bytes\n", unsafe.Sizeof(boolean))

	/*
		Каждая переменная хранится по определенному адресу памяти,
		и разные переменные распологаются в разных местах памяти.
		Адрес переменной получается путем присвоения этой переменной места в памяти.
		Когда мы объявляем переменную, Go выделяет память для этой переменной,
		и каждая переменная получает свой собственный уникальный адрес в памяти.
	*/

	// Для примера сравним указателей на одинаковые переменные
	str2 := str
	fmt.Printf("&str == &str2 %v\n", &str == &str2)
	fmt.Printf("str:   value = %s, address = %p\n", str, &str)
	fmt.Printf("str:   value = %s, address = %p\n", str2, &str2)
	// В этом примере в переменную str2 присваиваем значение переменной str.
	// Go создает разные адреса для каждой переменной

	/*
		Если нужно получить ссылку на тот же адрес пременной,
		то нужно передать ссылку на переменную,	а не значение,
		тогда переменная будет объявлена как указатель
	*/
	str3 := &str
	fmt.Printf("&str == &str2 %v\n", &str == str3)
	fmt.Printf("str:   value = %s, address = %p\n", str, &str)
	fmt.Printf("str:   value = %s, address = %p\n", *str3, str3)
}
