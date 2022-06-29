package main

func main() {
	c := make(chan int)
	go func() {
		c <- 1
	}()
	a := <-c
	println(a)
}
