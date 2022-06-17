package main

import "fmt"

func main() {
	println("Strategy Pattern")
	c := calculator{
		2, 3, &adder{},
	}
	fmt.Println("Add:", c.op.execute(2, 3))
	c.setOperation(&multiplier{})
	fmt.Println("Multiply:", c.op.execute(2, 3))
}

type operation interface {
	execute(int, int) int
}

type adder struct{}

type multiplier struct{}

func (op *adder) execute(a, b int) int {
	return a + b
}

func (op *multiplier) execute(a, b int) int {
	return a * b
}

type calculator struct {
	a, b int
	op   operation
}

func (c *calculator) setOperation(op operation) {
	c.op = op
}
