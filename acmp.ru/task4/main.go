package main

import (
	"io/ioutil"
	"log"
	"fmt"
	"strconv"
)

var (
	input = "./INPUT.TXT"
	output = "./OUTPUT.TXT"
)

func main() {
	bytes, err := ioutil.ReadFile(input)
	if err != nil {
		log.Fatalln(err)
	}
	line := string(bytes)
	a, err := strconv.Atoi(line)
	//5 594
	b := 9
	c := b - a
	//d := 
	fmt.Println(a,b,c)
}