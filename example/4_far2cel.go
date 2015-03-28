package main

import "fmt"

var Fahrenheite int
var Celsius int

/* Example from book how to convert temprature from Fahrenheite to Celsius*/
func main() {
	fmt.Println("Enter temprature in Fahrenheite:")

	fmt.Scanf("%d", &Fahrenheite)
	Celsius = (Fahrenheite - 32) * 5 / 9
	fmt.Println("Celsius is:", Celsius)
}
