// Реализуйте функцию для разблокировки мьютекса с помощью ​defer
package main

import (
	"fmt"
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

const count = 100

func main() {
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
