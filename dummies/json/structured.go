package main

import (
	"encoding/json"
	"fmt"
)

type People struct {
	Firstname string
	Lastname  string
	Details   struct {
		Height int
		Weight float32
	}
}

func main() {
	var persons []People
	jsonString := `[
	{"firstname":"Wei-Meng",
        "lastname":"Lee",
                "details": {
                    "height":175,
                    "weight":70.0
                }
	},
    {"firstname":"Mickey",
        "lastname":"Mouse",
		"details": {
			"height":105,
			"weight":85.5
		}}
	]`
	err := json.Unmarshal([]byte(jsonString), &persons)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, person := range persons {
		fmt.Println(person.Firstname, person.Lastname)
		fmt.Println(person.Details.Height, person.Details.Weight)
	}
}
