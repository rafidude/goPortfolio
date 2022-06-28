package main

import (
	"fmt"
	"os"
	"strconv"
)

// this program uses command line arguments
// to set the price per hour and calculate
// the total monthly and yearly prices
func main() {
	args := os.Args[1:]
	var pricePerHour float64
	// convert args[0] to float64
	pricePerHour, err := strconv.ParseFloat(args[0], 64)
	if err != nil {
		pricePerHour = 1.25
	}
	fmt.Printf("pricePerHour: %f\n", pricePerHour)
	// Calculate number of hours in a month and year
	hoursInDay := 24
	hoursInMonth := hoursInDay * 30
	hoursInYear := hoursInMonth * 12
	totalMonthlyPrice := float64(hoursInMonth) * pricePerHour
	totalYearlyPrice := float64(hoursInYear) * pricePerHour
	fmt.Printf("Total monthly price: %.2f\n", totalMonthlyPrice)
	fmt.Printf("Total yearly price: %.2f\n", totalYearlyPrice)
}
