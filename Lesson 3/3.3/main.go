package main

import "fmt"

const globalConst = 777

func main() {
	fmt.Println("Значение глобальной константы:", globalConst)

	const localConst = 555
	fmt.Println("Значение локальной константы:", localConst)

	// Затеняем глобальную константу локальной внутри функции main
	const globalConst = localConst
	fmt.Println("Значение глобальной константы внутри функции:", globalConst)
}
