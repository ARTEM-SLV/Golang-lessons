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
	for {
		select {
		case <-stop:
			fmt.Println("happy end")
			return
		case ch <- 1:
		}
	}

}
