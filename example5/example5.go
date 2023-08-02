package main

import (
	"fmt"
	"sync"
)

func main() {
	tasks := []int{1, 2, 3, 4, 5}

	input := make(chan int)
	output := make(chan int)

	// Fan-out
	go func() {
		for _, task := range tasks {
			input <- task
		}
		close(input)
	}()

	// Worker goroutines
	var wg sync.WaitGroup
	numWorkers := 3

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for task := range input {
				result := processTask(task)
				output <- result
			}
		}()
	}

	// Fan-in
	go func() {
		wg.Wait()
		close(output)
	}()

	// Collect results
	for result := range output {
		fmt.Println("Processed result:", result)
	}
}

func processTask(task int) int {
	// Perform task processing
	return task * task
}
