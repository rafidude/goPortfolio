package main

import (
	"fmt"
	"strings"
)

const favColor = "blue"

func main() {
	var guess string

	for {
		fmt.Println("Guess my favorite color?")
		if _, err := fmt.Scanln(&guess); err != nil {
			fmt.Println("Error:", err)
			return
		}
		// trim spaces and to lowercase
		guess = strings.ToLower(strings.TrimSpace(guess))

		if guess == favColor {
			fmt.Printf("You are correct, %q is not my favorite color!\n", guess)
			return
		}
		fmt.Printf("Sorry, %q is not my favorite color. Guess again.\n", guess)
	}
}
