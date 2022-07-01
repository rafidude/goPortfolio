package main

import "time"

func main() {
	t := time.NewTimer(2 * time.Second)
	<-t.C
	println("Timer expired")

	t2 := time.NewTimer(1 * time.Second)
	go func() {
		<-t2.C
		println("Timer2 expired")
	}()
	t2.Stop()
	println("t2 stopped")
	time.Sleep(3 * time.Second)
}
