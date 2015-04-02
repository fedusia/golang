package main

/* Write a function which takes an integer and halves it and
returns true if it was even or false if it was odd. For example half(1)
should return (0, false) and half(2) should return (1, true).
*/

import "fmt"

func main() {
	dig := 67
	fmt.Println(halv(dig))
}

func halv(x int) (int, bool) {
	if x%2 == 0 {
		return 1, true
	} else {
		return 0, false
	}
}
