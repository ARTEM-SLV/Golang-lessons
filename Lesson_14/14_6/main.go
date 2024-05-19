/*
Запущены две горутины, не считая главной горутины.
Первая читает из канала, вторая пишет в него.
Горутины завершаются, когда получают сигнал из канала stop (case <-stop:break OUT).

Первая горутина:
Читает из канала ch и выводит полученные значения.
Завершается при получении сигнала через канал stop.

Вторая горутина:
Генерирует целые числа и отправляет их в канал ch.
Завершается при получении сигнала через канал stop.

Главная горутина ждет 5 секунд,
что позволяет первой горутины считывать данные из канала
и выводить их на экран (1, 2, 3, ...)
*/

package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	stop := make(chan struct{}, 2)
	go func() {
	OUT:
		for {
			select {
			case <-stop:
				break OUT
			case v, ok := <-ch:
				if !ok {
					break OUT
				}
				fmt.Println(v)

			default:
				continue
			}
		}
		fmt.Println("завершение работы горутины_1")
	}()
	go func() {
		var i int
	OUT:
		for {
			i++
			select {
			case <-stop:
				break OUT
			default:
				time.Sleep(time.Second)
				if ch == nil {
					continue
				}
				ch <- i
			}
		}
		fmt.Println("завершение работы горутины_2")
	}()
	time.Sleep(5 * time.Second)
	stop <- struct{}{}
	stop <- struct{}{}
	time.Sleep(time.Second)

	fmt.Println("завершение работы главной горутины")
}
