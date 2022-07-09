package main

import "fmt"

type User struct {
	Name string
}

func main() {
	var m = map[string]User{
		"a": User{Name: "Peter"},
		"b": User{Name: "John"},
	}
	m["c"] = User{"Paul"}

	fmt.Println(m)

	c, ok := m["c"]
	fmt.Println("Key c:", c, ok)

	d, ok := m["d"]
	fmt.Println("Key d:", d, ok)

	m["d"] = User{"Mary"}
	delete(m, "b")
	m["a"] = User{"Ron"}
	fmt.Println(m)

}
