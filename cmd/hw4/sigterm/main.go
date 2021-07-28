// Написать программу, которая при получении в канал сигнала SIGTERM останавливается не позднее,
// чем за одну секунду (установить таймаут).

package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	// ctx := context.Background()
	ctx, cancelFunc := context.WithDeadline(context.Background(), time.Now().Add(60*time.Second))
	defer cancelFunc()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGTERM, syscall.SIGINT)

	doneCh := make(chan error)

	// прерывание работы по сигналу
	go func(ctx context.Context) {
		sig := <-sigCh
		fmt.Println("\nSIG:", sig)

		time.Sleep(time.Second)
		// отмена контекста
		cancelFunc()
		// в канал тоже можем писать ошибку
		// doneCh <- fmt.Errorf("Canceled...")
	}(ctx)

	// полезная работа тут
	go Work(ctx, doneCh)

	// ждём данных из канала
	err := <-doneCh

	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Println("Done")
	}
}

func Work(ctx context.Context, ch chan<- error) {
	i := 0
	for {
		select {
		case <-ctx.Done():
			ch <- ctx.Err()
			return
		default:
		}
		rand.Seed(time.Now().UTC().UnixNano())
		sec := rand.Intn(10)
		fmt.Println("working...", sec)
		time.Sleep(time.Duration(sec) * time.Second)
		// ch <- fmt.Errorf("Failed...")
		i++
		if i > 10 {
			break
		}

	}

	ch <- nil
}
