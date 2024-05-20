package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)
	stop := make(chan struct{})
	go func() {
		<-ch
		stop <- struct{}{}
	}()
	ch <- 1
	<-stop
	fmt.Println("happy end")

}
