package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func main() {
	var i int32
	go func() {
		atomic.AddInt32(&i, i+1)
	}()
	go func() {
		atomic.AddInt32(&i, i+1)
	}()
	go func() {
		atomic.AddInt32(&i, i+1)
	}()

	time.Sleep(time.Millisecond)
	fmt.Println(i)
}
