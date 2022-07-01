package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case t := <-ticker.C:
				fmt.Println("tick at", t)
			case <-done:
				ticker.Stop()
				println("Ticker stopped")
				return
			}
		}
	}()

	time.Sleep(1600 * time.Millisecond)
	done <- true
}
