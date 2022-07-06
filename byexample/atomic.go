package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// non-atomic version
func addOne(s *int) {
	*s += 1
}

func main() {

	var ops uint64
	sum := 0

	var wg sync.WaitGroup

	for i := 0; i < 50; i++ {
		wg.Add(1)

		go func() {
			for c := 0; c < 1000; c++ {

				atomic.AddUint64(&ops, 1)
				addOne(&sum)
			}
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println("ops:", ops)
	fmt.Println("sum:", sum)
}
