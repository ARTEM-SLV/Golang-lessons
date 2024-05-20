package main

import (
	"fmt"
	"sync"
	"time"
)

func start() {
	time.Sleep(2 * time.Second)
	fmt.Println("Функция start запущена")
}

func main() {
	var wg sync.WaitGroup
	var once sync.Once

	for i := 0; i < 10; i++ {
		i := i
		wg.Add(1)
		go func() {
			defer wg.Done()

			once.Do(start)
			fmt.Printf("Горутина %d запущена\n", i)
		}()
	}

	wg.Wait()
	fmt.Println("Все горутины завершены.")
}
