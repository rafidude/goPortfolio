package main

import "fmt"

func main() {
	gen := fib()
	for i := 0; i < 10; i++ {
		fmt.Println(gen())
	}
}

// the beauty of the fib() function is that it generates
// Fibo- nacci numbers on demand — it doesn’t have to
// store all the numbers generated. It only stores the
// last two numbers of the Fibonacci numbers at any given time.
func fib() func() int {
	f1 := 0
	f2 := 1
	return func() int {
		f1, f2 = f2, (f1 + f2)
		return f1
	}
}
