package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	wg := sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		i := i
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					fmt.Println("стоп горутина: ", i)
					return
				default:
					fmt.Println("горутина работает: ", i)
				}
			}
		}()
	}
	time.Sleep(time.Second)

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("Stop")
		cancel()
	}()

	wg.Wait()
}
