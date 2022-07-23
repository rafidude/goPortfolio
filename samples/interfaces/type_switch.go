package main

import "fmt"

type User struct {
	Name string
}

func checkType(a interface{}) {
	switch a.(type) {
	case string:
		fmt.Println("String", a.(string))
	case int:
		fmt.Println("Int", a.(int))
	case User:
		fmt.Println("User", a.(User))
	default:
		fmt.Println("Unknown")
	}
}

func main() {
	checkType("Hello")
	checkType(5)
	checkType(User{Name: "John"})
}
