package main

import (
	"fmt"
	"strings"
)

func main() {
	// strings.Fields splits a string into a slice of strings
	data := "username 	password   email 		date"
	fields := strings.Fields(data)
	fmt.Printf("%q\n", fields)
}
