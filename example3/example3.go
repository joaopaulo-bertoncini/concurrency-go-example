package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	var counter int64

	// Increment the counter concurrently
	for i := 0; i < 10; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
		}()
	}

	// Wait for goroutines to complete
	// (not required in this specific example)

	fmt.Println("Counter:", atomic.LoadInt64(&counter))
}
