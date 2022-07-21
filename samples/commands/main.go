package main

import (
	"fmt"
	"os/exec"
	"runtime"
)

func execute_specific(command string) (output string, err error) {
	out, err := exec.Command(command).Output()
	if err != nil {
		fmt.Printf("%s", err)
		return "", err
	}
	fmt.Printf("%s command completed\n", command)
	output = string(out[:])
	fmt.Println(output)
	return output, err
}

func execute() {
	execute_specific("ls")
	execute_specific("pwd")
	// out, err = exec.Command("pwd").Output()
	// if err != nil {
	// 	fmt.Printf("%s", err)
	// }
	// fmt.Println("pwd command completed")
	// output = string(out[:])
	// fmt.Println(output)
}

func main() {
	if runtime.GOOS == "windows" {
		fmt.Println("Can't Execute this on a windows machine")
	} else {
		execute()
	}
}
