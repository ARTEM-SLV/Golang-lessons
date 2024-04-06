package main

import "fmt"

func main() {
	hello()
	defer fmt.Println("завершение работы")
}

func hello() {
	fmt.Println("Hello, Go")
}
