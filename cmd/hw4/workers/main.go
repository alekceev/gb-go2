// С помощью пула воркеров написать программу, которая запускает 1000 горутин,
// каждая из которых увеличивает число на 1.
// Дождаться завершения всех горутин и убедиться,
// что при каждом запуске программы итоговое число равно 1000.

package main

import (
	"fmt"
	"time"
)

const WorkerCount = 8

func main() {
	// пул воркеров
	var workers = make(chan struct{}, WorkerCount)
	// канал для получения результата
	var ch1 = make(chan int)
	var ch2 = make(chan int)
	var res = make(chan int)

	// добавим начальное значение в канал
	go func() {
		ch1 <- 1
	}()

	// читатель, нужно значение записываем в канал res
	go func() {
		for val := range ch2 {
			fmt.Println(val)
			if val == 1000 {
				res <- val
				// закрываем канал при нужном значении
				close(ch1)
			} else {
				ch1 <- val
			}
		}
	}()

	for k := 1; k < 1000; k++ {
		// берём воркер в работу засылая в канал данные
		workers <- struct{}{}

		// горутина обработчик - в конце читаем с канала воркера
		// для его освобождения
		go func() {
			defer func() {
				<-workers
			}()

			i, ok := <-ch1
			if !ok {
				return
			}

			// имитация задерки при выполнении полезной работы
			time.Sleep(1 * time.Millisecond)
			// типо инкрементим
			ch2 <- i + 1
		}()
	}

	// ждём результата
	i := <-res
	fmt.Println("Num:", i)
}
