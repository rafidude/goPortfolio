package main

import (
	"fmt"
	"time"
)

// Implement timeouts functionality using channels
func main() {
	// channel string buffer 1
	ch := make(chan string, 1)
	// send a message to channel after sleep for 2 seconds
	go func() {
		time.Sleep(2 * time.Second)
		ch <- "hello1"
	}()

	select {
	case msg := <-ch:
		println(msg)
	case <-time.After(1 * time.Second):
		println("timeout 1")
	case <-time.After(3 * time.Second):
		fmt.Println("timeout 2")
	}

	// channel string buffer 1
	ch2 := make(chan string, 1)
	// send a message to channel after sleep for 2 seconds
	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "hello2"
	}()
	select {
	case msg2 := <-ch2:
		println(msg2)
	case msg1 := <-ch:
		println(msg1)
	case <-time.After(3 * time.Second):
		fmt.Println("timeout 2")
	}

}
