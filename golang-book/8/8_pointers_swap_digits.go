package main

import "fmt"

/* Write a program that can swap two integers (x := 1; y := 2; swap(&x, &y) should give you x=2 and y=1). */

func main() {
	x:=31
	y:=2
	swap(&x,&y)
	fmt.Println("x=",x,"y=",y)
}

func swap(x *int, y *int) {
	xTmpPtr := new(int)
	*xTmpPtr = *x
	*x = *y
	*y = *xTmpPtr
}