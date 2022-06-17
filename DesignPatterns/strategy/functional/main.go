package main

func main() {
	println("Functional Strategy Pattern!")
	println(execute(2, 3, add))
	println(execute(2, 3, multiply))
}

func add(a, b int) int {
	return a + b
}

func multiply(a, b int) int {
	return a * b
}

func execute(a, b int, f func(int, int) int) int {
	return f(a, b)
}
