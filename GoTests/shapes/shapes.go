package shapes

import "math"

type Shape interface {
	Area() float64
	Perimeter() float64
}

func Perimeter(s Shape) float64 {
	return s.Perimeter()
}

func Area(s Shape) float64 {
	return s.Area()
}

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

type Triangle struct {
	Base   float64
	Height float64
}

func (t Triangle) Perimeter() float64 {
	return t.Base + t.Height +
		math.Sqrt(t.Base*t.Base+t.Height*t.Height)
}

func (t Triangle) Area() float64 {
	return t.Base * t.Height * 0.5
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}
