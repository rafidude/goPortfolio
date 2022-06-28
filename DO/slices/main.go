package main

import (
	"errors"
	"fmt"
	"time"
)

func slicesFunc() {
	sl := []string{"a", "b", "c"}
	sl2 := []string{"d", "e", "f"}
	sl = append(sl, sl2...)
	fmt.Println(sl)
	// remove 3rd element from slice sl
	sl = append(sl[:2], sl[3:]...)
	fmt.Println(sl)
}

func errorsSimple() {
	err := errors.New("My error")
	fmt.Println(err)
	// formatted error
	err = fmt.Errorf("error occurred at: %v", time.Now())
	fmt.Println(err)
}

func simulateError() error {
	return errors.New("something went wrong")
}

func main() {
	slicesFunc()
	errorsSimple()

	err := simulateError()
	if err != nil {
		fmt.Println(err)
	}

	sl1 := []int{1, 2, 3}
	fmt.Println(sl1[3])
}
