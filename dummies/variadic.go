package main

func main() {
	var a []int
	a = append(a, 1, 2, 3, 4)
	println(addNums(a...))
}

func addNums(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}
