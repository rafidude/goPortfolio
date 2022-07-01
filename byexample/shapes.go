package main

import "fmt"

type shape interface {
	area() float64
	perim() float64
}

type rectangle struct {
	width, height int
}

type circle struct {
	radius int
}

//area of rectangle
func (r rectangle) area() float64 {
	return float64(r.width * r.height)
}

//perimeter of rectangle
func (r rectangle) perim() float64 {
	return 2 * float64(r.width+r.height)
}

//area of circle
func (c circle) area() float64 {
	return 3.14 * float64(c.radius*c.radius)
}

//perimeter of circle
func (c circle) perim() float64 {
	return 2 * 3.14 * float64(c.radius)
}

func shapeStats(s shape) {
	fmt.Println("Area:", s.area())
	fmt.Println("Perimeter:", s.perim())
}

func main() {
	r := rectangle{width: 10, height: 5}
	c := circle{radius: 5}

	shapeStats(r)
	shapeStats(c)
}
