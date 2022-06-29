package main

import "fmt"

func main() {
	var i interface{}
	describe(i)
	i = "hello"
	describe(i)
	i = 3.14
	describe(i)
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
