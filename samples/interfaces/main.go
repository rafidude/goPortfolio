package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
}

type rect struct {
	width, height float64
}
type circle struct {
	radius float64
}

func (r rect) area() float64 {
	return r.width * r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
}

type Body struct {
	Msg interface{}
}

func main() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}

	measure(r)
	measure(c)

	b := Body{"Hello there"}
	fmt.Printf("%#v %T\n", b.Msg, b.Msg)

	b.Msg = 5
	fmt.Printf("%#v %T\n", b.Msg, b.Msg)
}
