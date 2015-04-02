package main

import "fmt"

func main() {
	fmt.Println("What is your name?")
	var name string
	fmt.Scanf("%s", &name)
	fmt.Println("Hi", name, "!!!")
}
