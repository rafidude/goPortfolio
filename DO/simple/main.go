package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Enter your name: ")
	var name string
	fmt.Scanln(&name)
	name = strings.TrimSpace(name)
	fmt.Printf("Hello, %s! I'm Go!\n", name)

	sharks := []string{"Hammerhead", "Great White", "Tiger"}
	for _, shark := range sharks {
		fmt.Println(shark)
	}
}
