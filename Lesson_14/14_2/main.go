/*
Create an unbuffered channel that will
accept a string value. Write “Hello,
string channel!” to the channel,
then read this value and output to stdout
*/

package main

import "fmt"

func main() {
	ch := make(chan string)
	go func() {
		ch <- "Привет, строковый канал!"
	}()

	fmt.Println(<-ch)
}
