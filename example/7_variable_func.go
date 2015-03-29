package main

/* Write a function with one variadic parameter that 
finds the greatest number in a list of numbers.
*/

import "fmt"

func main() {
	fmt.Println(greatest(1,2,3,4,5,6,223,12312))
}

func greatest(args ... int) int {
	big:=args[0]
	for i:=1; i < len(args); i++{
		if big < args[i]{
			big = args[i]
		}
	}
	return big
}