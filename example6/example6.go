package main

import (
	"fmt"
	"sync"
)

func main() {
	tasks := []int{1, 2, 3, 4, 5}

	tasksChan := make(chan int)
	resultsChan := make(chan int)

	numWorkers := 3

	var wg sync.WaitGroup

	// Create worker goroutines
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for task := range tasksChan {
				result := processTask(task)
				resultsChan <- result
			}
		}()
	}

	// Enqueue tasks
	go func() {
		for _, task := range tasks {
			tasksChan <- task
		}
		close(tasksChan)
	}()

	// Collect results
	go func() {
		wg.Wait()
		close(resultsChan)
	}()

	// Process results
	for result := range resultsChan {
		fmt.Println("Processed result:", result)
	}
}

func processTask(task int) int {
	// Perform task processing
	return task * task
}
