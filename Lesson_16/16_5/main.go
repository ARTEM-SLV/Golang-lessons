package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}
	stop := make(chan struct{})

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for {
				select {
				case <-stop:
					fmt.Printf("stop горутина: %d\n", id)
					return
				default:
					time.Sleep(time.Second)
					fmt.Printf("сложные вычисления горутины: %d\n", id)
				}
			}
		}(i)
	}

	go func() {
		time.Sleep(3 * time.Second)
		fmt.Println("ой, всё!")
		close(stop)
	}()

	wg.Wait()
}
