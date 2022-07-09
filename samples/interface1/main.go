package main

import "fmt"

type PowerDrawer interface {
	DrawPower(power int)
}

type socket struct{}

func (socket) Plug(device PowerDrawer, power int) {
	device.DrawPower(power)
}

type mobile struct {
	brand string
}

func (m mobile) DrawPower(power int) {
	fmt.Printf("%T -> brand: %s, power: %d\n", m, m.brand, power)
}

type laptop struct {
	cpu string
}

func (l laptop) DrawPower(power int) {
	fmt.Printf("%T -> cpu: %s, power: %d\n", l, l.cpu, power)
}

type toaster struct {
	amount int
}

func (t toaster) DrawPower(power int) {
	fmt.Printf("%T -> amount: %d, power: %d\n", t, t.amount, power)
}

type kettle struct {
	quantity string
}

func (k kettle) DrawPower(power int) {
	fmt.Printf("%T -> quantity: %s, power: %d\n", k, k.quantity, power)
}

func main() {
	m := mobile{"Apple"}
	l := laptop{"Intel i9"}
	t := toaster{4}
	k := kettle{"50%"}

	s := socket{}

	s.Plug(m, 10)
	s.Plug(l, 50)
	s.Plug(t, 30)
	s.Plug(k, 25)
}
