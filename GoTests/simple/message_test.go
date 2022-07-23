package simple

import "testing"

func TestHello(t *testing.T) {
	want := "Hello, world!"
	if got := hello(); got != want {
		t.Errorf("hello() = %q, want %q", got, want)
	}
}

func TestMorning(t *testing.T) {
	want := "Good morning, world!"
	if got := morning(); got != want {
		t.Errorf("morning() = %q, want %q", got, want)
	}
}
