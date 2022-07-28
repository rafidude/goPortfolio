package main

func main() {
	max := 100
	for a, b := 0, 1; b <= max; a, b = b, a+b {
		println(b)
	}
}
