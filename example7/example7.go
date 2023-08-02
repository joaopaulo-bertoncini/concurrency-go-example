package main

import "fmt"

func main() {
	tasks := []int{1, 2, 3, 4, 5}

	// Stage 1: Generate tasks
	taskChan := make(chan int)
	go func() {
		for _, task := range tasks {
			taskChan <- task
		}
		close(taskChan)
	}()

	// Stage 2: Process tasks
	resultChan := make(chan int)
	go func() {
		for task := range taskChan {
			result := processTask(task)
			resultChan <- result
		}
		close(resultChan)
	}()

	// Stage 3: Consume results
	for result := range resultChan {
		fmt.Println("Processed result:", result)
	}
}

func processTask(task int) int {
	// Perform task processing
	return task * task
}
