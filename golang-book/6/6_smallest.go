package main

/* Task: find the smalest element in array */
import "fmt"

func main() {
	x := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 17, 9,
	}
	minimum := x[0]
	for i := 1; i < len(x); i++ {
		if x[i] < minimum {
			minimum = x[i]
		}
	}
	fmt.Println(minimum)
}
