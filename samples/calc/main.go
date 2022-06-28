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
	pricePerHour := 1.25
	args := os.Args[1:]
	// if the command line has prince per hour information, use it
	if len(args) > 0 {
		pricePerHour, _ = strconv.ParseFloat(args[0], 64)
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
	sammy := map[string]string{"name": "Sammy",
		"animal": "shark", "color": "blue"}
	delete(sammy, "color")
	// loop through the map
	for key, value := range sammy {
		fmt.Printf("%s: %s\n", key, value)
	}

}
