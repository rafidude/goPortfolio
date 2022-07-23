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

func TestAreaPerimeter(t *testing.T) {
	areaTests := []struct {
		shape     Shape
		area      float64
		perimeter float64
	}{
		{Rectangle{12, 6}, 72.0, 36.0},
		{Circle{10}, 314.1592, 62.8318},
		{Triangle{12, 6}, 36.0, 31.4164},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if math.Abs(got-tt.area) > delta {
			t.Errorf("Area got %g want %g", got, tt.area)
		}
		got2 := tt.shape.Perimeter()
		if math.Abs(got2-tt.perimeter) > delta {
			t.Errorf("Perimeter got %g want %g", got2, tt.perimeter)
		}
	}

}
