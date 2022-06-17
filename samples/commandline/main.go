package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	argsWithoutProg := os.Args[1:]

	arg := os.Args[3]

	fmt.Println(argsWithoutProg)
	fmt.Println(arg)

	os.Setenv("FOO", "1")
	fmt.Println("FOO:", os.Getenv("FOO"))
	fmt.Println("BAR:", os.Getenv("BAR"))

	fmt.Println()
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		fmt.Println(pair[0])
	}

}
