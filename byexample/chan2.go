package main

import "time"

//create a worker
func worker(done chan bool) {
	println("working...")
	// wait for a second
	time.Sleep(time.Second)
	println("done")
	done <- true
}

func main() {
	// create a channel of type bool
	c := make(chan bool)
	// Start the worker
	go worker(c)
	// wait for the worker to finish
	<-c
}
