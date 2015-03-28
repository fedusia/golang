package main

import "fmt"

const ft float64 = 0.3048

var distance float64
var meters float64

func main() {
	fmt.Println("Enter distance in ft:")
	fmt.Scanf("%f", &distance)

	meters = ft * distance
	fmt.Println("Distance in meters:", meters)
}