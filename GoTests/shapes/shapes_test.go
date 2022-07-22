package shapes

import (
	"math"
	"testing"
)

var delta float64 = 0.001

func TestPerimeter(t *testing.T) {
	perimeter := func(s Shape, want float64, t *testing.T) {
		got := Perimeter(s)
		if math.Abs(got-want) > delta {
			t.Errorf("got %.4f want %.4f", got, want)
		}
	}

	r := Rectangle{10.0, 10.0}
	perimeter(r, 40.0, t)

	c := Circle{10.0}
	perimeter(c, 62.8319, t)
}

func TestArea(t *testing.T) {
	area := func(s Shape, want float64, t *testing.T) {
		got := Area(s)
		if math.Abs(got-want) > delta {
			t.Errorf("got %.4f want %.4f", got, want)
		}
	}

	r := Rectangle{12.0, 6.0}
	area(r, 72.0, t)

	c := Circle{10.0}
	area(c, 314.1592, t)
}
