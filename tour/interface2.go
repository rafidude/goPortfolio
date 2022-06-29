package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	println(t.S)
}

type F float64

func (f F) M() {
	println(f)
}

func main() {
	var i I = &T{"Hello"}
	i.M()
	describe(i)

	var f F = 3.14
	f.M()
	describe(f)
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
