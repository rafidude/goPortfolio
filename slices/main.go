package main

import (
	"errors"
	"fmt"
)

// insert item into a slice at a given index
func insert(slice []int, index int, item int) ([]int, error) {
	if index > len(slice) {
		return nil, errors.New("index out of bounds")
	}
	slice = append(slice[:index], append([]int{item},
		slice[index:]...)...)
	return slice, nil
}

//delete item from a slice at a given index
func delete(slice []int, index int) ([]int, error) {
	if index < 0 || index > len(slice) {
		return nil, errors.New("index out of bounds")
	}
	slice = append(slice[:index], slice[index+1:]...)
	return slice, nil
}

func main() {
	slice := []int{1, 2, 3, 4, 5}
	slice, _ = insert(slice, 2, 0)
	fmt.Println(slice)
	slice, _ = delete(slice, 4)
	fmt.Println(slice)
}
