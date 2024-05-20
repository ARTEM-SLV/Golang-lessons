package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Meteo interface {
	ReadTemp() string
	ChangeTemp(v string)
}

type Temp struct {
	value string
	mu    sync.RWMutex
}

func (t *Temp) ReadTemp() string {
	t.mu.RLock()
	defer t.mu.RUnlock()

	return t.value
}

func (t *Temp) ChangeTemp(v string) {
	t.mu.Lock()
	defer t.mu.Unlock()

	t.value = v
}

func main() {
	temp := Temp{
		value: "+20 C",
	}

	for i := 0; i < 10; i++ {
		n := rand.Intn(30-20+1) + 20

		go func() {
			temp.ChangeTemp(fmt.Sprintf("+%d C", n))
		}()
		go func() {
			fmt.Println(temp.ReadTemp())
		}()
	}

	time.Sleep(time.Second)
}
