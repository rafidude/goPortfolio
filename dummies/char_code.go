package main

import (
	"fmt"
)

func main() {
	for pos, char := range "Hello, world!" {
		fmt.Println(pos, char)
	}

	for pos, char := range "Hello, world!" {
		fmt.Printf("%d %c\n", pos, char)
	}
}
