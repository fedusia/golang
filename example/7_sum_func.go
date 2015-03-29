package main

/* Task: sum is a function which takes a slice of numbers and adds them together. What would its function signature look like in Go? */

import "fmt"

func main() {
	v := []int{1,2,4}
	fmt.Println(sum(v))
}

func sum (slice []int) int{
	total := 0
	for _, v := range slice {
		total += v
	}
	return total
}