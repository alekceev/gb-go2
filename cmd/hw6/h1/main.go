// Написать программу, которая использует мьютекс для безопасного доступа к данным
// из нескольких потоков. Выполните трассировку программы

// go tool trace trace.out

package main

import (
	"fmt"
	"os"
	"runtime/trace"
	"sync"
)

type Counter struct {
	sync.Mutex
	Cnt int
}

func (c *Counter) Inc() {
	c.Lock()
	defer c.Unlock()
	c.Cnt += 1
}

const count = 1000

func main() {
	trace.Start(os.Stderr)
	defer trace.Stop()

	var wg = sync.WaitGroup{}
	counter := &Counter{}

	for i := 0; i < count; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Inc()
		}()
	}

	wg.Wait()

	fmt.Printf("counter: %d\n", counter.Cnt)
}
