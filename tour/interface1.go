package main

type I interface {
	M()
}

type T struct {
	S string
}

func (t T) M() {
	println(t.S)
}

func main() {
	var i I = T{"Hello"}
	i.M()
}
