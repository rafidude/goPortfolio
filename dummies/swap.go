package main

func main() {
	x := 5
	y := 6
	swap(&x, &y)
	println(x, y)
}

func swap(a, b *int) {
	*a, *b = *b, *a
}
