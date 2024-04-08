package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var (
		num     int     = 42
		str     string  = "Hello"
		flt     float64 = 3.14
		boolean bool    = true
	)

	fmt.Printf("num:   value = %d, address = %p\n", num, &num)
	fmt.Printf("str:   value = %s, address = %p\n", str, &str)
	fmt.Printf("flt:   value = %f, address = %p\n", flt, &flt)
	fmt.Printf("array: value = %v, address = %p\n", boolean, &boolean)

	// Размер переменных в байтах
	fmt.Printf("Size of int:     %d bytes\n", unsafe.Sizeof(num))
	fmt.Printf("Size of string:  %d bytes\n", unsafe.Sizeof(str))
	fmt.Printf("Size of float64: %d bytes\n", unsafe.Sizeof(flt))
	fmt.Printf("Size of array:   %d bytes\n", unsafe.Sizeof(boolean))
}
