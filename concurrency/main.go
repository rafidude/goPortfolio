package main

import "fmt"

func main() {
	ch := make(chan string)

	go func(x int) {
		fmt.Println(x)
	}(42)

	go push("moe", ch)
	go push("larry", ch)
	go push("curly", ch)
	fmt.Println(<-ch, <-ch, <-ch)
	// using an anonymous inner function in a goroutine
}

func push(name string, ch chan string) {
	msg := "Hello, " + name + "!\n"
	ch <- msg
}
