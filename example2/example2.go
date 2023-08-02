package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		// Simulating a long-running operation
		time.Sleep(2 * time.Second)
		cancel() // Cancel the operation
	}()

	select {
	case <-ctx.Done():
		fmt.Println("Operation cancelled")
	}
}
