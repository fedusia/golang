package main

import "fmt"

func main() {
	var total float64 = 0
	var value float64 = 0

	x := [5]float64{33, 35, 55, 66, 100}
	/*
	   We can use like this:
	   	for i:=0; i < len(x); i++ {
	   		total += x[i]
	   	}
	*/

	/*
	   If we would use like this the compiler will give us an error:

	   	for i, value = range x{
	   		total += value
	   	}
	*/

	/*
	   But we should use like this because a single _ (underscore) is used to tell the compiler that we don't need this. (In this case we don't need the iterator variable)
	*/
	for _, value = range x {
		total += value
	}
	fmt.Println(total / float64(len(x)))
}
