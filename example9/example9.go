package main

import (
	"errors"
	"fmt"
)

func main() {
	resultChan := make(chan int, 1)
	errorChan := make(chan error, 1)

	go func() {
		result, err := performTask()
		if err != nil {
			errorChan <- err // Propagate the error
			return
		}
		resultChan <- result
	}()

	for {
		select {
		case result := <-resultChan:
			fmt.Println("Task result:", result)
			return
		case err := <-errorChan:
			fmt.Println("Error occurred:", err)
			return
		}
	}
}

func performTask() (int, error) {
	// Perform the task and handle errors
	return 42, errors.New("something went wrong")
}
