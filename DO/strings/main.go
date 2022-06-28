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
	
	a := "some string"
	b := []byte(a)
	fmt.Println(b)

	sammy := map[string]string{"name": "Sammy", 
		"animal": "shark", "color": "blue"}

	for k, v := range sammy {
		fmt.Println(k, v)
	}
}
