package main

import (
	"fmt"
	"math"
)

func main() {
	for i := 1; i <= 3; i++ {
		fmt.Println(i)
		j := float64(i)/100.0 + 1.0
		fmt.Println(i, math.Pow(j, 30))
		fmt.Println(i, math.Pow(j, 100))
		fmt.Println(i, math.Pow(j, 365))
	}
}
