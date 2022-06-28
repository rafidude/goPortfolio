package main

func main() {
	divideByZero()
	println("After recovering from panic: Hello, world!")
}

// function that divides by zero
func divideByZero() int {
	var a int = 1
	var b int = 0
	defer func() {
		if r := recover(); r != nil {
			println("Recovered in divideByZero: ", r)
		}
	}()
	return a / b
}
