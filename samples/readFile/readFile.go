package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func readFilePrintString(filePath string) {
	// Read test.cmd file and output its contents
	file, err := os.Open(filePath)
	defer file.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	// Read the file into a byte slice
	b, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
		return
	}
	// Print the contents of the byte slice
	fmt.Println(string(b))
}

func readFilePrintLineByLine(filePath string) {
	// Read test.cmd file and output its contents
	file, err := os.Open(filePath)
	defer file.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	fileScanner := bufio.NewScanner(file)
	lineNum := 0
	for fileScanner.Scan() {
		fmt.Println(lineNum, fileScanner.Text())
		lineNum++
	}
}

func main() {
	readFilePrintString("./test.cmd")
	readFilePrintLineByLine("./test.cmd")
}
