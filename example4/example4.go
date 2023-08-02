package main

import (
	"fmt"
	"time"
)

func main() {
	channel1 := make(chan string)
	channel2 := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		channel1 <- "Message from Channel 1"
	}()

	go func() {
		time.Sleep(1 * time.Second)
		channel2 <- "Message from Channel 2"
	}()

	select {
	case msg := <-channel1:
		fmt.Println("Received from Channel 1:", msg)
	case msg := <-channel2:
		fmt.Println("Received from Channel 2:", msg)
	}
}
