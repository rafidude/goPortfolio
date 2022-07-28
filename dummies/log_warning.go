package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	formattingLogging()
}

func formattingLogging() {
	fmt.Println("-----------formattingLogging-----------")
	var warning *log.Logger
	warning = log.New(
		os.Stdout,
		"WARNING: ",
		log.Ldate|log.Ltime|log.Lshortfile)
	warning.Println("This is a warning.")
	warning.Println("Something bad happened!")
}
