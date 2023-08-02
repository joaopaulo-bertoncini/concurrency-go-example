package main

import (
	"fmt"
	"sync"
)

func main() {
	var mutex sync.Mutex
	var wg sync.WaitGroup

	counter := 0

	// Increment the counter concurrently
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mutex.Lock()
			counter++
			mutex.Unlock()
		}()
	}

	wg.Wait()

	fmt.Println("Counter:", counter)
}
