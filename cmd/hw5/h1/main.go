// Напишите программу, которая запускает n потоков и дожидается завершения их всех

package main

import (
	"fmt"
	"sync"
)

const n = 100

func main() {
	var wg = &sync.WaitGroup{}

	for i := 1; i < n; i++ {
		wg.Add(1)
		go func(i int) {
			fmt.Println(i)
			defer wg.Done()
		}(i)
	}

	wg.Wait()
	fmt.Println("Done")
}
