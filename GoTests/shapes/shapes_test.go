package shapes

import (
	"math"
	"testing"
)

var delta float64 = 0.001

func TestPerimeter(t *testing.T) {
	r := Rectangle{10.0, 10.0}
	got := r.Perimeter()
	want := 40.0
	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}

	c := Circle{10.0}
	got = c.Perimeter()
	want = 62.8319
	if math.Abs(got-want) > delta {
		t.Errorf("got %.4f want %.4f", got, want)
	}
}

func TestArea(t *testing.T) {
	r := Rectangle{12.0, 6.0}
	got := r.Area()
	want := 72.0
	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}

	c := Circle{10.0}
	got = c.Area()
	want = 314.16
	if math.Abs(got-want) > delta {
		t.Errorf("got %.4f want %.4f", got, want)
	}
}
